package repository

import (
	"database/sql"
)

type IGrammarRepository interface {
}

type GrammarRepository struct {
	db *sql.DB
}

func NewGrammarRepository(db *sql.DB) IGrammarRepository {
	return &GrammarRepository{db: db}
}

func (r *GrammarRepository) AddGrammar() {

}
