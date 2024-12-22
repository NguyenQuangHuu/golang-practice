package model

import "time"

type Grammar struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	LastModified time.Time `json:"last_modified"`
}

type GrammarDetail struct {
	ID         int     `json:"id"`
	Title      string  `json:"title"`
	Definition string  `json:"definition"`
	Example    string  `json:"example"`
	Grammar    Grammar `json:"grammar"`
}

type SentenceExample struct {
	ID           int          `json:"id"`
	Example      string       `json:"example"`
	Grammar      Grammar      `json:"grammar"`
	SentenceType SentenceType `json:"sentence_type"`
}

type SentenceType struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Usage struct {
	ID            int           `json:"id"`
	Situation     string        `json:"situation"`
	Example       Example       `json:"example"`
	GrammarDetail GrammarDetail `json:"grammar_detail"`
}
