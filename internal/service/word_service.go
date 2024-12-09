package service

import (
	"awesomeProject/internal/model"
	"awesomeProject/internal/repository"
)

type IWordService interface {
	GetWordByID(id int) (*model.WordModel, error)
	GetAllWords() ([]*model.WordModel, error)
	AddWord(word *model.Word) error
	UpdateWordByID(word *model.Word) error
	FindByWord(word string) ([]*model.WordModel, error)
}

type WordService struct {
	wordRepository repository.IWordRepository
}

func (w *WordService) FindByWord(word string) ([]*model.WordModel, error) {
	//TODO implement me
	result, err := w.wordRepository.FindByWord(word)
	if err != nil {
		return nil, err
	}
	if len(result) > 0 {
		return result, nil
	}
	return nil, nil
}

func NewWordService(wordRepository repository.IWordRepository) IWordService {
	return &WordService{wordRepository: wordRepository}
}

func (w *WordService) GetWordByID(id int) (*model.WordModel, error) {
	result, err := w.wordRepository.GetWordByID(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (w *WordService) GetAllWords() ([]*model.WordModel, error) {
	result, err := w.wordRepository.GetAllWords()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (w *WordService) AddWord(word *model.Word) error {
	err := w.wordRepository.AddWord(word)

	if err != nil {
		return err
	}
	return nil
}

func (w *WordService) UpdateWordByID(word *model.Word) error {
	err := w.wordRepository.UpdateWordByID(word)
	if err != nil {
		return err
	}
	return nil
}
