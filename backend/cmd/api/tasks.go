package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
	"todoapp/backend/pkg/data"
	"todoapp/backend/pkg/validator"
)

func (app *application) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var inputTask struct {
		Name     string          `json:"name"`
		Desc     string          `json:"desc"`
		Status   data.TaskStatus `json:"status"`
		Deadline *time.Time      `json:"deadline"`
	}

	err := app.readJSON(w, r, &inputTask)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	task := &data.Task{
		Name:     inputTask.Name,
		Desc:     inputTask.Desc,
		Status:   inputTask.Status,
		Deadline: inputTask.Deadline,
	}

	v := validator.New()

	if data.ValidateTask(v, task); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Tasks.InsertTask(task)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/task/%d", task.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"task": task}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getTaskHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	task, err := app.models.Tasks.GetTask(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"task": task}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	task, err := app.models.Tasks.GetTask(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var taskinput struct {
		Name     *string          `json:"name"`
		Desc     *string          `json:"desc"`
		Status   *data.TaskStatus `json:"status"`
		Deadline *time.Time       `json:"deadline"`
	}

	err = app.readJSON(w, r, &taskinput)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if taskinput.Name != nil {
		task.Name = *taskinput.Name
	}

	if taskinput.Desc != nil {
		task.Desc = *taskinput.Desc
	}

	if taskinput.Status != nil {
		task.Status = *taskinput.Status
	}

	if taskinput.Deadline != nil {
		task.Deadline = taskinput.Deadline
	}

	v := validator.New()

	if data.ValidateTask(v, task); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Tasks.UpdateSubscription(task)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return

	}

	err = app.writeJSON(w, http.StatusOK, envelope{"task": task}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
	}

	err = app.models.Tasks.DeleteTask(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)

		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "успешно удалено"}, nil)
}

func (app *application) listTasksHandler(w http.ResponseWriter, r *http.Request) {
	var taskInput struct {
		taskStatus data.TaskStatus
		startDate  *time.Time
		endDate    *time.Time
		data.Filters
	}

	v := validator.New()

	qs := r.URL.Query()

	taskInput.startDate, taskInput.endDate = app.readDateRange(qs, "start", "end", v)
	taskInput.taskStatus = app.readTaskStatus(qs, "status", "")

	var taskStatus *data.TaskStatus
	if taskInput.taskStatus != "" {
		taskStatus = &taskInput.taskStatus
	}

	tasks, metadata, err := app.models.Tasks.GetAllTasks(taskInput.startDate, taskInput.endDate, taskStatus, taskInput.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"tasks": tasks, "metadata": metadata}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
