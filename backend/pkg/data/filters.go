package data

import (
	"math"
	"todoapp/backend/pkg/validator"
)

type Filters struct {
	Page         int
	PageSize     int
	Sort         string
	SortSafelist []string
}

type Metadata struct {
	CurrentPage  int `json:"current_page,omitempty"`
	PageSize     int `json:"page_size,omitempty"`
	FirstPage    int `json:"first_page,omitempty"`
	LastPage     int `json:"last_page,omitempty"`
	TotalRecords int `json:"total_records,omitempty"`
}

func calculateMetadata(totalRecords, page, pageSize int) Metadata {
	if totalRecords == 0 {
		return Metadata{}
	}

	return Metadata{
		CurrentPage:  page,
		PageSize:     pageSize,
		FirstPage:    1,
		LastPage:     int(math.Ceil(float64(totalRecords) / float64(pageSize))),
		TotalRecords: totalRecords,
	}
}

func ValidateFilters(v *validator.Validator, f Filters) {
	v.Check(f.Page > 0, "page", "должно быть больше нуля!")
	v.Check(f.Page <= 10_000_000, "page", "Максимум 10 миллионов")
	v.Check(f.PageSize > 0, "page_size", "должно быть больше нуля!")
	v.Check(f.PageSize <= 20, "page", "Максимум 20")

}

func (f Filters) limit() int {
	if f.PageSize <= 0 {
		return 20 // дефолт, например
	}
	return f.PageSize
}

func (f Filters) offset() int {
	if f.Page <= 0 {
		return 0
	}
	return (f.Page - 1) * f.PageSize
}
