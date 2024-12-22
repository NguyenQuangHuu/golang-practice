package repository

import (
	"awesomeProject/internal/model"
	"database/sql"
)

// IWordRepository Repository for communicate with db layer
type IWordRepository interface {
	GetWordByID(id int) (*model.Word, error)
	GetAllWords() ([]*model.Word, error)
	AddWord(word *model.Word) error
	UpdateWordByID(word *model.Word) error
	FindByWord(word string) ([]*model.Word, error)
}

type WordRepository struct {
	db *sql.DB
}

// NewWordRepository constructor
func NewWordRepository(db *sql.DB) *WordRepository {
	return &WordRepository{db}
}

// GetWordByID function use to get word by id
func (w *WordRepository) GetWordByID(id int) (*model.Word, error) {
	var word model.Word
	result := w.db.QueryRow("select wi.id,wi.word,wi.meaning_vn,wi.meaning_de,wt.id,wt.name,wt.description "+
		"from word_information as wi"+
		" inner join word_type as wt on wi.word_type_id = wt.id where wi.id = $1", id)
	var wordTypeID int
	if err := result.Scan(&word.ID, &word.Word, &word.MeaningVN, &word.MeaningDE, &wordTypeID, &word.WordType.Name, &word.WordType.Description); err != nil {
		return nil, err
	}
	return &word, nil
}

func (w *WordRepository) GetAllWords() ([]*model.Word, error) {
	var words []*model.Word
	result, err := w.db.Query("select wi.id,wi.word,wi.meaning_vn,wi.meaning_de,wt.id,wt.name,wt.description " +
		"from word_information as wi inner join word_type as wt on wi.word_type_id = wt.id")
	defer result.Close()
	if err != nil {
		return nil, err
	}
	for result.Next() {
		var word model.Word
		var wordTypeID int
		if err := result.Scan(&word.ID, &word.Word, &word.MeaningVN, &word.MeaningDE, &wordTypeID, &word.WordType.Name, &word.WordType.Description); err != nil {
			return nil, err
		}
		words = append(words, &word)
	}
	return words, err
}

func (w *WordRepository) AddWord(word *model.Word) error {
	_, err := w.db.Exec("insert into word_information "+
		"(word,meaning_vn,meaning_de,word_type_id)"+
		" values ($1,$2,$3,$4) ", word.Word, word.MeaningVN, word.MeaningDE, word.WordType.ID)
	if err != nil {
		return err
	}
	return nil
}

func (w *WordRepository) UpdateWordByID(word *model.Word) error {
	_, err := w.db.Exec("update word_information "+
		"set word = $1, meaning_vn = $2, meaning_de = $3, word_type_id = $4 where id = $5",
		&word.Word, &word.MeaningVN, &word.MeaningDE, &word.WordType.ID, &word.ID)
	if err != nil {
		return err
	}
	return nil
}

func (w *WordRepository) FindByWord(word string) ([]*model.Word, error) {
	var query = "" +
		"select word_information.id,word_information.word,word_information.meaning_vn,word_information.meaning_de," +
		"word_type.id,word_type.name,word_type.description from word_information inner join word_type on word_information.word_type_id = word_type.id where word like $1"
	rows, err := w.db.Query(query, word+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var words []*model.Word
	for rows.Next() {
		var word model.Word
		var wordTypeID int
		if err := rows.Scan(&word.ID, &word.Word, &word.MeaningVN, &word.MeaningDE, &wordTypeID, &word.WordType.Name, &word.WordType.Description); err != nil {
			return nil, err
		}
		word.WordType.ID = wordTypeID
		words = append(words, &word)
	}
	return words, nil
}
