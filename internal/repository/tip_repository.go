package repository

import "database/sql"

type ITipRepository interface {
}

type TipRepository struct {
	db *sql.DB
}

func NewTipRepository(db *sql.DB) ITipRepository {
	return &TipRepository{db: db}
}
