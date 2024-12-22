package model

type Notes struct {
	ID            int           `json:"id"`
	Content       string        `json:"content"`
	GrammarDetail GrammarDetail `json:"grammar_detail"`
	NoteType      NoteType      `json:"note_type"`
	Example       []Example     `json:"example"`
}

type NoteType struct {
	ID      int    `json:"id"`
	Name    string `json:"name"` // Warning, Hints
	Content string `json:"content"`
}

type Example struct {
	ID      int    `json:"id"`
	Word    string `json:"word"`
	Example string `json:"example"`
}
