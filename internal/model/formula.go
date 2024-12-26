package model

type GrammarFormula struct {
	ID      int       `json:"id"`
	Formula string    `json:"formula"`
	Example []Example `json:"example"`
}
