package service

import "awesomeProject/internal/repository"

type IExampleService interface {
}

type ExampleService struct {
	exampleRepo repository.IExampleRepository
}

func NewExampleService(exampleRepo repository.IExampleRepository) IExampleService {
	return &ExampleService{exampleRepo: exampleRepo}
}
