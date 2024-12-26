package service

import (
	"awesomeProject/internal/repository"
)

type IUsageService interface {
}

type UsageService struct {
	usageRepo repository.IUsageRepository
}

func NewUsageService(usageRepo repository.IUsageRepository) repository.IUsageRepository {
	return &UsageService{usageRepo: usageRepo}
}
