package pagination

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaginationFilter struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

func NewPaginationFilterFromGinContext(ctx *gin.Context) PaginationFilter {
	return NewPaginationFilter(
		getPaginationFromGinCtx(ctx),
	)
}

func NewPaginationFilter(page, size int) PaginationFilter {
	return PaginationFilter{
		Page:     page,
		PageSize: size,
	}
}

func getPaginationFromGinCtx(ctx *gin.Context) (page, size int) {
	var err error
	if page, err = strconv.Atoi(ctx.Query("page")); err != nil {
		page = 1
	}

	if size, err = strconv.Atoi(ctx.Query("pageSize")); err != nil {
		size = 100
	}

	return
}

func (p *PaginationFilter) ToDatabaseFormat() (offset, limit int) {
	offset = max(0, p.Page-1)

	if p.PageSize <= 0 {
		limit = 100
	} else {
		limit = p.PageSize
	}

	return
}
