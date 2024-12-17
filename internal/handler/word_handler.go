package handler

import (
	"awesomeProject/internal/model"
	"awesomeProject/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type IWordHandler interface {
	GetWordByID(c *gin.Context)
	GetAllWords(c *gin.Context)
	AddWord(c *gin.Context)
	UpdateWordByID(c *gin.Context)
	FindByWord(c *gin.Context)
}

type WordHandle struct {
	service service.IWordService
}

func NewWordHandleRequest(wordService service.IWordService) *WordHandle {
	return &WordHandle{service: wordService}
}

func (w *WordHandle) GetWordByID(c *gin.Context) {
	param := c.Params.ByName("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	word, err := w.service.GetWordByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, word)
}

func (w *WordHandle) GetAllWords(c *gin.Context) {
	words, err := w.service.GetAllWords()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, words)
}

func (w *WordHandle) AddWord(c *gin.Context) {
	var word model.Word
	if err := c.ShouldBindJSON(&word); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := w.service.AddWord(&word)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, word)
}

func (w *WordHandle) UpdateWordByID(c *gin.Context) {
	var word model.Word
	param := c.Params.ByName("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err = c.ShouldBindJSON(&word); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if id != word.ID {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "wrong id"})
		return
	}
	err = w.service.UpdateWordByID(&word)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, word)
}

func (w *WordHandle) FindByWord(c *gin.Context) {
	searchWord := c.DefaultQuery("query", "")
	result, err := w.service.FindByWord(searchWord)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}
