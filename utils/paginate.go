package utils

import "gorm.io/gorm"

type Paginate struct {
	limit int
	page  int
}

const defaultOfLimit = 10

func NewPaginate(page, limit int) *Paginate {
	if page <= 0 {
		page = 1
	}
	if limit <= defaultOfLimit {
		limit = defaultOfLimit
	}
	return &Paginate{limit: limit, page: page}
}

func (p *Paginate) PaginatedResult(db *gorm.DB) *gorm.DB {
	offset := (p.page - 1) * p.limit
	return db.Offset(offset).
		Limit(p.limit)
}
