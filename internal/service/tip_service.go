package service

import (
	"awesomeProject/internal/repository"
)

type ITipService interface {
}

type TipService struct {
	tipRepo repository.ITipRepository
}

func NewTipService(tipRepo repository.ITipRepository) repository.ITipRepository {
	return &TipService{tipRepo: tipRepo}
}
