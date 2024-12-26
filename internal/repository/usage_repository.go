package repository

import "database/sql"

type IUsageRepository interface {
}

type UsageRepository struct {
	db *sql.DB
}

func NewUsageRepository(db *sql.DB) IUsageRepository {
	return &UsageRepository{db: db}
}
