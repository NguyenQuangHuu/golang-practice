package service

import (
	"awesomeProject/internal/repository"
)

type IFormulaService interface{}

type FormulaService struct {
	formulaRepo repository.IFormulaRepository
}

func NewFormulaService(formulaRepo repository.IFormulaRepository) IFormulaService {
	return &FormulaService{formulaRepo: formulaRepo}
}
