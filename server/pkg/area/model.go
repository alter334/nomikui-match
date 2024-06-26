package area

import "github.com/jmoiron/sqlx"

type AreaService struct {
	db *sqlx.DB
}

func NewAreaService(_db *sqlx.DB) *AreaService {
	return &AreaService{db: _db}
}
