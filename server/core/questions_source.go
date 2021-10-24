package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func fileName(name string) string {
	return fmt.Sprintf(filePath, name)
}

func ReadIndexes() (index IndexStore, err error) {
	byteValue, err := ioutil.ReadFile(fileName(indexFile))
	json.Unmarshal(byteValue, &index)
	return
}

func ReadQuestionsForIndex(index Index) (questions []NewQuestion) {
	byteValue, _ := ioutil.ReadFile(fileName(index.Tag))
	var questionBank QuestionBank
	json.Unmarshal(byteValue, &questionBank)
	questions = questionBank.Questions
	return
}
