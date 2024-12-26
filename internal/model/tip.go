package model

type Tips struct {
	ID      int       `json:"id"`
	Content string    `json:"content"`
	Example []Example `json:"example"`
	TipType TipType   `json:"tip_type"`
}

type TipType struct {
	ID                 int    `json:"id"`
	TipTypeName        string `json:"tip_type_name"`
	TipTypeDescription string `json:"tip_type_description"`
}
