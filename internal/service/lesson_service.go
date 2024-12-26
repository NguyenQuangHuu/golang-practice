package service

import (
	"awesomeProject/internal/repository"
)

type ILessonService interface {
}

type LessonService struct {
	lessonRepo repository.ILessonRepository
}

func NewLessonService(lessonRepo repository.ILessonRepository) repository.ILessonRepository {
	return &LessonService{lessonRepo: lessonRepo}
}
