package model

type WordModel struct {
	ID           int    `json:"id"`
	Word         string `json:"word"`
	MeaningVN    string `json:"meaning_vn"`
	MeaningDE    string `json:"meaning_de"`
	WordTypeName string `json:"word_type_name"`
}

type Word struct {
	ID         int    `json:"id"`
	Word       string `json:"word"`
	MeaningVN  string `json:"meaning_vn"`
	MeaningDE  string `json:"meaning_de"`
	WordTypeID int    `json:"word_type_id"`
}
