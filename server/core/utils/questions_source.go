package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/atulanand206/find/core/models"
)

func fileName(name string) string {
	return fmt.Sprintf("filePath", name)
}

func ReadIndexes() (index models.IndexStore, err error) {
	byteValue, err := ioutil.ReadFile(fileName("indexFile"))
	json.Unmarshal(byteValue, &index)
	return
}

func ReadQuestionsForIndex(index models.Index) (questions []models.NewQuestion) {
	byteValue, _ := ioutil.ReadFile(fileName(index.Tag))
	var questionBank models.QuestionBank
	json.Unmarshal(byteValue, &questionBank)
	questions = questionBank.Questions
	return
}
