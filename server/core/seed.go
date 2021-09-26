package core

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
		Id  string `json:"id" bson:"_id"`
		Tag string `json:"tag" bson:"tag"`
	}

	IndexWrapper struct {
		Indexes []Index `json:"indexes" bson:"indexes"`
	}

	IndexStore struct {
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
	filePath  = "G:\\binge\\binge-questions\\%s.json"
	indexFile = "index"
)

func HandlerAddQuestion(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var requestBody NewQuestion
	err := decoder.Decode(&requestBody)
	if err != nil {
		http.Error(w, Err_RequestNotDecoded, http.StatusInternalServerError)
		return
	}
	var question Question
	question.Id = uuid.New().String()
	question.Statements = requestBody.Statements
	requestDto, _ := mongo.Document(question)
	insertedQuestionId, err := mongo.Write(Database, QuestionCollection, *requestDto)
	if err != nil {
		http.Error(w, Err_QuestionNotCreated, http.StatusInternalServerError)
		return
	}
	var answer Answer
	answer.Id = uuid.New().String()
	answer.QuestionId = fmt.Sprint(insertedQuestionId.InsertedID)
	answer.Answer = requestBody.Answer
	requestDto, _ = mongo.Document(answer)
	insertedAnswerId, err := mongo.Write(Database, AnswerCollection, *requestDto)
	if err != nil {
		http.Error(w, Err_AnswerNotCreated, http.StatusInternalServerError)
		return
	}
	var response AddQuestionResponse
	response.QuestionId = fmt.Sprint(insertedQuestionId.InsertedID)
	response.AnswerId = fmt.Sprint(insertedAnswerId.InsertedID)
	json.NewEncoder(w).Encode(response)
}

func HandlerSeedQuestions(w http.ResponseWriter, r *http.Request) {
	filePath = os.Getenv("SEED_FILES_PATH")

	err := mongo.DropCollections(Database, []string{IndexCollection, QuestionCollection, AnswerCollection})
	if err != nil {
		http.Error(w, Err_CollectionsNotDropped, http.StatusInternalServerError)
		return
	}

	err = mongo.CreateCollections(Database, []string{IndexCollection, QuestionCollection, AnswerCollection})
	if err != nil {
		http.Error(w, Err_CollectionsNotCreated, http.StatusInternalServerError)
		return
	}

	indexes, questions, answers, err := createDeckDtos()
	if err != nil {
		http.Error(w, Err_DeckDtosNotCreated, http.StatusInternalServerError)
		return
	}

	err = seedIndexes(indexes)
	if err != nil {
		http.Error(w, Err_IndexNotSeeded, http.StatusInternalServerError)
		return
	}

	err = seedQuestions(questions)
	if err != nil {
		http.Error(w, Err_QuestionsNotSeeded, http.StatusInternalServerError)
		return
	}

	err = seedAnswers(answers)
	if err != nil {
		http.Error(w, Err_AnswersNotSeeded, http.StatusInternalServerError)
		return
	}
}

func createDeckDtos() (indexs []Index, questions []Question, answers []Answer, err error) {
	indexsWrapper, err := readIndexes()
	if err != nil {
		return
	}
	indexs = make([]Index, 0)
	questions = make([]Question, 0)
	answers = make([]Answer, 0)
	for _, index := range indexsWrapper.Indexes {
		idx := initNewIndex(index)
		indexs = append(indexs, idx)
		newQuestions := readQuestionsForIndex(idx)
		for _, newQuestion := range newQuestions {
			question := initNewQuestion(idx, newQuestion)
			questions = append(questions, question)
			answers = append(answers, initNewAnswer(question, newQuestion))
		}
	}
	return
}

func seedIndexes(indexes []Index) error {
	var indexesDto []interface{}
	for _, t := range indexes {
		indexesDto = append(indexesDto, t)
	}
	_, err := mongo.WriteMany(Database, IndexCollection, indexesDto)
	return err
}

func seedQuestions(questions []Question) error {
	var questionsDto []interface{}
	for _, t := range questions {
		questionsDto = append(questionsDto, t)
	}
	_, err := mongo.WriteMany(Database, QuestionCollection, questionsDto)
	return err
}

func seedAnswers(answers []Answer) error {
	var answersDto []interface{}
	for _, answer := range answers {
		answersDto = append(answersDto, answer)
	}
	_, err := mongo.WriteMany(Database, AnswerCollection, answersDto)
	return err
}

func readIndexes() (index IndexStore, err error) {
	byteValue, err := ioutil.ReadFile(fileName(indexFile))
	json.Unmarshal(byteValue, &index)
	return
}

func readQuestionsForIndex(index Index) (questions []NewQuestion) {
	byteValue, _ := ioutil.ReadFile(fileName(index.Tag))
	var questionBank QuestionBank
	json.Unmarshal(byteValue, &questionBank)
	questions = questionBank.Questions
	return
}

func fileName(name string) string {
	return fmt.Sprintf(filePath, name)
}

func initNewIndex(tag string) (index Index) {
	index.Tag = tag
	index.Id = uuid.New().String()
	return
}

func initNewQuestion(index Index, newQuestion NewQuestion) (question Question) {
	question.Statements = newQuestion.Statements
	question.Tag = index.Id
	question.Id = uuid.New().String()
	return
}

func initNewAnswer(question Question, newQuestion NewQuestion) (answer Answer) {
	answer.Answer = newQuestion.Answer
	answer.QuestionId = question.Id
	answer.Id = uuid.New().String()
	return
}

func mapSansTags(tags []string) (tagsMap map[string]bool) {
	tagsMap = make(map[string]bool)
	for _, tag := range tags {
		tagsMap[tag] = true
	}
	return
}
