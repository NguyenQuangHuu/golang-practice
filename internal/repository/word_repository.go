package repository

import (
	"awesomeProject/internal/model"
	"database/sql"
)

// IWordRepository Repository for db layer
type IWordRepository interface {
	GetWordByID(id int) (*model.WordModel, error)
	GetAllWords() ([]*model.WordModel, error)
	AddWord(word *model.Word) error
	UpdateWordByID(word *model.Word) error
	FindByWord(word string) ([]*model.WordModel, error)
}

type WordRepository struct {
	db *sql.DB
}

func NewWordRepository(db *sql.DB) *WordRepository {
	return &WordRepository{db}
}

func (w *WordRepository) GetWordByID(id int) (*model.WordModel, error) {
	var word model.WordModel
	result := w.db.QueryRow("select wi.id,wi.word,wi.meaning_vn,wi.meaning_de,wt.name "+
		"from word_information as wi"+
		" inner join word_type as wt on wi.word_type_id = wt.id where wi.id = $1", id)
	if err := result.Scan(&word.ID, &word.Word, &word.MeaningVN, &word.MeaningDE, &word.WordTypeName); err != nil {
		return nil, err
	}
	return &word, nil
}

func (w *WordRepository) GetAllWords() ([]*model.WordModel, error) {
	var words []*model.WordModel
	result, err := w.db.Query("select wi.id,wi.word,wi.meaning_vn,wi.meaning_de,wt.name " +
		"from word_information as wi inner join word_type as wt on wi.word_type_id = wt.id")
	defer result.Close()
	if err != nil {
		return nil, err
	}
	for result.Next() {
		var word model.WordModel
		if err := result.Scan(&word.ID, &word.Word, &word.MeaningVN, &word.MeaningDE, &word.WordTypeName); err != nil {
			return nil, err
		}
		words = append(words, &word)
	}
	return words, err
}

func (w *WordRepository) AddWord(word *model.Word) error {
	_, err := w.db.Exec(
		"insert into word_information"+
			"(word,meaning_vn,meaning_de,word_type_id)"+
			" values ($1,$2,$3,$4) ", word.Word, word.MeaningVN, word.MeaningDE, word.WordTypeID)
	if err != nil {
		return err
	}
	return nil
}

func (w *WordRepository) UpdateWordByID(word *model.Word) error {
	_, err := w.db.Exec("update word_information "+
		"set word = $1, meaning_vn = $2, meaning_de = $3, word_type_id = $4 where id = $5",
		&word.Word, &word.MeaningVN, &word.MeaningDE, &word.WordTypeID, &word.ID)
	if err != nil {
		return err
	}
	return nil
}

func (w *WordRepository) FindByWord(word string) ([]*model.WordModel, error) {
	var query = "select wi.id,wi.word,wi.meaning_vn,wi.meaning_de,wt.name " +
		"from word_information as wi " +
		"inner join word_type as wt on wi.word_type_id = wt.id where wi.word like $1"
	rows, err := w.db.Query(query, word)
	defer rows.Close()
	var words []*model.WordModel
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var word model.WordModel
		if err := rows.Scan(&word.ID, &word.Word, &word.MeaningVN, &word.MeaningDE, &word.WordTypeName); err != nil {
			return nil, err
		}
		words = append(words, &word)
	}

	return words, nil
}
