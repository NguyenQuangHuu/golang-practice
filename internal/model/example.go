package model

type Example struct {
	ID            int          `json:"id"`
	Word          string       `json:"word"`
	ExampleByWord string       `json:"example_by_word"`
	IsCorrect     bool         `json:"is_correct"`
	SentenceType  SentenceType `json:"sentence_type"`
}

type SentenceType struct {
	ID          int    `json:"id"`
	TypeName    string `json:"type_name"`
	Description string `json:"description"`
}
