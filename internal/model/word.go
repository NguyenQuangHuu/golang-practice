package model

type WordModel struct {
	ID           int    `json:"id"`
	Word         string `json:"word"`
	MeaningVN    string `json:"meaning_vn"`
	MeaningDE    string `json:"meaning_de"`
	WordTypeName string `json:"word_type_name"`
}

type Word struct {
	ID        int      `json:"id"`
	Word      string   `json:"word"`
	MeaningVN string   `json:"meaning_vn"`
	MeaningDE string   `json:"meaning_de"`
	WordType  WordType `json:"word_type"`
}

type WordType struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
