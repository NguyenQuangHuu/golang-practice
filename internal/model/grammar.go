package model

type Grammar struct {
	ID             int              `json:"id"`
	Title          string           `json:"title"`
	GrammarFormula []GrammarFormula `json:"grammar_formula"`
	UsageCase      []UsageCase      `json:"usage_case"`
	Tips           []Tips           `json:"tips"`
}
