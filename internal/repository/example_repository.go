package repository

import "database/sql"

type IExampleRepository interface {
}

type ExampleRepository struct {
	db *sql.DB
}

func NewExampleRepository(db *sql.DB) IExampleRepository {
	return &ExampleRepository{db: db}
}
