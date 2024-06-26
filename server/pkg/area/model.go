package area

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type AreaService struct {
	db *sqlx.DB
}

type Area struct {
	Areaname string    `db:"areaname"`
	Id       uuid.UUID `db:"id"`
}

func NewAreaService(_db *sqlx.DB) *AreaService {
	return &AreaService{db: _db}
}
