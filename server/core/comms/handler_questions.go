package comms

import (
	"net/http"
	"os"

	"github.com/atulanand206/find/server/core/errors"
	"github.com/atulanand206/find/server/core/models"
	"github.com/atulanand206/find/server/core/utils"
)

var (
	FilePath  = "G:\\binge\\binge-questions\\%s.json"
	IndexFile = "index"
)

func HandlerAddQuestion(w http.ResponseWriter, r *http.Request) {
	requestBody, err := models.DecodeAddQuestionRequest(r)
	if err != nil {
		http.Error(w, errors.Err_RequestNotDecoded, http.StatusInternalServerError)
		return
	}

	if err = Controller.QuestionService.AddQuestion(requestBody.Tag, requestBody.Questions); err != nil {
		http.Error(w, errors.Err_QuestionsNotSeeded, http.StatusInternalServerError)
		return
	}
}

func HandlerSeedQuestions(w http.ResponseWriter, r *http.Request) {
	FilePath = os.Getenv("SEED_FILES_PATH")

	var err error
	if err = Db.DropCollections(); err != nil {
		http.Error(w, errors.Err_CollectionsNotDropped, http.StatusInternalServerError)
		return
	}

	if err = Db.CreateCollections(); err != nil {
		http.Error(w, errors.Err_CollectionsNotCreated, http.StatusInternalServerError)
		return
	}

	indexes, questions, answers, err := utils.CreateDeckDtos()
	if err != nil {
		http.Error(w, errors.Err_DeckDtosNotCreated, http.StatusInternalServerError)
		return
	}

	if err = Controller.QuestionService.Crud.SeedIndexes(indexes); err != nil {
		http.Error(w, errors.Err_IndexNotSeeded, http.StatusInternalServerError)
		return
	}

	if err = Controller.QuestionService.Crud.SeedQuestions(questions); err != nil {
		http.Error(w, errors.Err_QuestionsNotSeeded, http.StatusInternalServerError)
		return
	}

	if err = Controller.QuestionService.Crud.SeedAnswers(answers); err != nil {
		http.Error(w, errors.Err_AnswersNotSeeded, http.StatusInternalServerError)
		return
	}
}
