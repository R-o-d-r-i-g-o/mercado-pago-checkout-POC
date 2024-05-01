package pagination

import (
	"math"
)

type Page[T any] struct {
	Data        []T   `json:"data"`
	CurrentPage int   `json:"currentPage"`
	Count       int   `json:"count"`
	Limit       int   `json:"limit"`
	TotalCount  int64 `json:"totalCount"`
	TotalPages  int64 `json:"totalPages"`
}

func (p Page[T]) CurrentOn(pageNumber int) Page[T] {
	p.CurrentPage = pageNumber
	return p
}

func (p Page[T]) LimitedTo(limit int) Page[T] {
	p.Limit = limit
	return p
}

func (p Page[T]) TotalOf(total int64) Page[T] {
	p.TotalCount = total
	return p
}

func (p Page[T]) With(data []T) Page[T] {
	p.Data = data
	p.Count = len(data)
	p.TotalPages = int64(max(1, math.Ceil(float64(p.TotalCount)/float64(p.Limit))))
	return p
}
