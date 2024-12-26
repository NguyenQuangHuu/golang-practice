package repository

import "database/sql"

type IFormulaRepository interface{}

type FormulaRepository struct {
	db *sql.DB
}

func NewFormulaRepository(db *sql.DB) IFormulaRepository {
	return &FormulaRepository{db: db}
}
