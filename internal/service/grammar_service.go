package service

import (
	"awesomeProject/internal/repository"
)

type IGrammarService interface {
}

type GrammarService struct {
	grammarRepo repository.IGrammarRepository
}

func NewGrammarService(grammarRepo repository.IGrammarRepository) repository.IGrammarRepository {
	return &GrammarService{grammarRepo: grammarRepo}
}
