package core

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type PaginationMeta struct {
	TotalItems   int `json:"totalItems"`
	ItemCount    int `json:"itemCount"`
	ItemsPerPage int `json:"itemsPerPage"`
	TotalPages   int `json:"totalPages"`
	CurrentPage  int `json:"currentPage"`
}

type PaginationQuery struct {
	Page    int    // default: 1
	Limit   int    // default: 10
	OrderBy string // default: "id"
	SortBy  string // default: "asc"
	Search  string // default: ""
}

func ParsePaginationQuery(c *gin.Context) PaginationQuery {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	return PaginationQuery{
		Page:    page,
		Limit:   limit,
		Search:  c.DefaultQuery("search", ""),
		OrderBy: c.DefaultQuery("orderBy", "id"),
		SortBy:  c.DefaultQuery("sortBy", "asc"),
	}
}
