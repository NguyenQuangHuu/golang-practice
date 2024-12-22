package model

type Word struct {
	ID               int              `json:"id"`
	Word             string           `json:"word"`
	Pronounce        string           `json:"pronounce"`
	MeaningVN        string           `json:"meaning_vn"`
	MeaningDE        string           `json:"meaning_de"`
	VocabularyLesson VocabularyLesson `json:"vocabulary_lesson"`
	WordType         WordType         `json:"word_type"`
}

type WordType struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
