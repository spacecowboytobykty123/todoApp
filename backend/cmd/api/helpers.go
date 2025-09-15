package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
	"todoapp/backend/pkg/data"
	"todoapp/backend/pkg/validator"
)

type envelope map[string]any

func (app *application) writeJSON(w http.ResponseWriter, status int, data any, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t") // JSON делает приятнее на вид но тяжелее
	if err != nil {
		return err
	}

	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	maxBytes := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError
		var maxBytesError *http.MaxBytesError
		switch {

		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)

		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")

		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)

		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return fmt.Errorf("body contains unknown key %s", fieldName)

		case errors.As(err, &maxBytesError):
			return fmt.Errorf("body must not be larger than %d bytes", maxBytesError.Limit)

		case errors.As(err, &invalidUnmarshalError):
			panic(err)
		default:
			return err
		}
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only contain a single JSON value")
	}

	return nil
}

func (app *application) readIDParam(r *http.Request) (uuid.UUID, error) {
	params := httprouter.ParamsFromContext(r.Context())
	idParam := params.ByName("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		return uuid.Nil, errors.New("invalid UUID parameter")
	}

	return id, nil
}

func (app *application) readDateRange(qs url.Values, startKey, endKey string, v *validator.Validator) (*time.Time, *time.Time) {
	var startDate *time.Time
	var endDate *time.Time

	// Parse start date
	if s := qs.Get(startKey); s != "" {
		t, err := time.Parse("2006-01-02", s) // strict YYYY-MM-DD
		if err != nil {
			v.AddError(startKey, "must be in format YYYY-MM-DD")
		} else {
			startDate = &t
		}
	}

	// Parse end date
	if e := qs.Get(endKey); e != "" {
		t, err := time.Parse("2006-01-02", e)
		if err != nil {
			v.AddError(endKey, "must be in format YYYY-MM-DD")
		} else {
			endDate = &t
		}
	}

	if startDate != nil && endDate != nil && endDate.Before(*startDate) {
		v.AddError(endKey, "must be after start date")
	}

	return startDate, endDate
}

func (app *application) readTaskStatus(qs url.Values, key string, defaultValue data.TaskStatus) data.TaskStatus {
	s := qs.Get(key)

	if s == "" {
		return defaultValue
	}

	statuses := map[string]data.TaskStatus{
		"Назначена":   data.Assigned,
		"В работе":    data.InProgress,
		"Выполненный": data.Completed,
		"Отклонена":   data.Rejected,
	}

	if v, ok := statuses[s]; ok {
		return v
	}

	return defaultValue
}
