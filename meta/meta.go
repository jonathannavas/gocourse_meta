package meta

import (
	"log"
	"strconv"
)

type Meta struct {
	TotalCount int `json:"total_count"`
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	PageCount  int `json:"page_count"`
}

func New(page, perPage, total int, pageLimitDefault string) (*Meta, error) {

	if perPage <= 0 {
		var err error
		perPage, err = strconv.Atoi(pageLimitDefault)
		if err != nil {
			return nil, err
		}
	}

	pageCount := 0

	if total >= 0 {
		pageCount = (total + perPage - 1) / perPage
		if page > pageCount {
			page = pageCount
		}
	}

	if page < 1 {
		page = 1
	}

	return &Meta{
		Page:       page,
		TotalCount: total,
		PerPage:    perPage,
		PageCount:  pageCount,
	}, nil
}

func (p *Meta) Offset() int {
	log.Println(p)
	return (p.Page - 1) * p.PerPage
}

func (p *Meta) Limit() int {
	return p.PerPage
}
