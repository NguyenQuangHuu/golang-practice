package model

type UsageCase struct {
	ID      int       `json:"id"`
	Content string    `json:"content"`
	Example []Example `json:"example"`
}
