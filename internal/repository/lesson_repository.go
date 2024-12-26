package repository

import "database/sql"

type ILessonRepository interface {
}

type LessonRepository struct {
	db *sql.DB
}

func NewLessonRepository(db *sql.DB) ILessonRepository {
	return &LessonRepository{db: db}
}
