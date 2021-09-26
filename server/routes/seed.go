package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/atulanand206/go-mongo"
	"github.com/google/uuid"
)

type (
	Index struct {
		Indexes []string `json:"indexes" bson:"indexes"`
	}

	QuestionBank struct {
		Questions []NewQuestion `json:"questions" bson:"questions"`
	}

	NewQuestion struct {
		Statements []string `json:"statements"`
		Answer     string   `json:"answer"`
	}

	AddQuestionResponse struct {
		QuestionId string `json:"question_id"`
		AnswerId   string `json:"answer_id"`
	}
)

var (
	filePath  string
	indexFile = "index"
)

func HandlerSeedQuestions(w http.ResponseWriter, r *http.Request) {
	filePath = os.Getenv("SEED_FILES_PATH")
	questions, answers := seedQuestions()

	var questionsDto []interface{}
	for _, t := range questions {
		questionsDto = append(questionsDto, t)
	}
	_, err := mongo.WriteMany(Database, QuestionCollection, questionsDto)
	if err != nil {
		http.Error(w, Err_QuestionsNotSeeded, http.StatusInternalServerError)
		return
	}

	var answersDto []interface{}
	for _, answer := range answers {
		answersDto = append(answersDto, answer)
	}
	_, err = mongo.WriteMany(Database, AnswerCollection, answersDto)
	if err != nil {
		http.Error(w, Err_AnswersNotSeeded, http.StatusInternalServerError)
		return
	}
}

func seedQuestions() (questions []Question, answers []Answer) {
	indexes := indexes()
	questions = make([]Question, 0)
	answers = make([]Answer, 0)
	for _, index := range indexes.Indexes {
		newQuestions := indexQuestion(index)
		for _, newQuestion := range newQuestions {
			question := mapNewQuestion(index, newQuestion)
			questions = append(questions, question)
			answers = append(answers, mapNewAnswer(question, newQuestion))
		}
	}
	return
}

func indexes() (index Index) {
	byteValue, _ := ioutil.ReadFile(fileName(indexFile))
	json.Unmarshal(byteValue, &index)
	return
}

func indexQuestion(index string) (questions []NewQuestion) {
	byteValue, _ := ioutil.ReadFile(fileName(index))
	var questionBank QuestionBank
	json.Unmarshal(byteValue, &questionBank)
	questions = questionBank.Questions
	return
}

func fileName(name string) string {
	return fmt.Sprintf(filePath, name)
}

func mapNewQuestion(index string, newQuestion NewQuestion) (question Question) {
	question.Id = uuid.New().String()
	question.Statements = newQuestion.Statements
	question.Tag = index
	return
}

func mapNewAnswer(question Question, newQuestion NewQuestion) (answer Answer) {
	answer.Id = uuid.New().String()
	answer.Answer = newQuestion.Answer
	answer.QuestionId = question.Id
	return
}
