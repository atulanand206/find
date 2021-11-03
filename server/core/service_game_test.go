package core_test

import (
	"testing"

	"github.com/atulanand206/find/server/core"
)

func testPlayerJoin(quizId string) (core.GameResponse, error) {
	service := core.Service{}
	player, _ := core.PlayerService{}.FindOrCreatePlayer(testPlayer())
	request := core.Request{
		QuizId: quizId,
		Person: player,
	}
	return service.GenerateEnterGameResponse(request)
}

func TestCreateGame(t *testing.T) {
	teardown := Setup(t)
	defer teardown(t)

	service := core.Service{}

	quizmaster := testPlayer()
	specs := testSpecs()
	response, _ := service.GenerateCreateGameResponse(quizmaster, specs)

	if response.Quiz.QuizMaster.Email != quizmaster.Email {
		t.Fatalf("player email does not match")
	}
	if response.Quiz.Specs.Name != specs.Name {
		t.Fatalf("name does not match")
	}
	if response.Quiz.Specs.Points != specs.Points {
		t.Fatalf("points does not match")
	}
	if response.Quiz.Specs.Questions != specs.Questions {
		t.Fatalf("question count does not match")
	}
	if response.Quiz.Specs.Rounds != specs.Rounds {
		t.Fatalf("round count does not match")
	}
	if response.Quiz.Specs.Teams != specs.Teams {
		t.Fatalf("team count does not match")
	}
	if response.Quiz.Specs.Players != specs.Players {
		t.Fatalf("player count does not match")
	}
}

func TestEnterGameQuizmaster(t *testing.T) {
	teardown := Setup(t)
	defer teardown(t)

	service := core.Service{}

	quizmaster := testPlayer()
	specs := testSpecs()
	response, _ := service.GenerateCreateGameResponse(quizmaster, specs)
	quiz := response.Quiz
	request := core.Request{
		QuizId: quiz.Id,
		Person: quizmaster,
	}
	res, _ := service.GenerateEnterGameResponse(request)
	if res.Quiz.Id != quiz.Id {
		t.Fatalf("quiz id does not match")
	}
}

func TestEnterGameNewPlayer(t *testing.T) {
	teardown := Setup(t)
	defer teardown(t)

	service := core.Service{}

	quizmaster, _ := core.PlayerService{}.FindOrCreatePlayer(testPlayer())
	specs := testSpecs()
	response, _ := service.GenerateCreateGameResponse(quizmaster, specs)
	quiz := response.Quiz
	player, _ := core.PlayerService{}.FindOrCreatePlayer(testPlayer())
	request := core.Request{
		QuizId: quiz.Id,
		Person: player,
	}
	res, _ := service.GenerateEnterGameResponse(request)
	if res.Quiz.Id != quiz.Id {
		t.Fatalf("quiz id does not match")
	}
	for _, team := range res.Snapshot.Roster {
		for _, plr := range team.Players {
			if plr.Email == player.Email {
				return
			}
		}
	}
	t.Fatalf("player not added")
}

func TestEnterGameAlreadyAddedPlayer(t *testing.T) {
	teardown := Setup(t)
	defer teardown(t)

	service := core.Service{}

	quizmaster, _ := core.PlayerService{}.FindOrCreatePlayer(testPlayer())
	specs := testSpecs()
	response, _ := service.GenerateCreateGameResponse(quizmaster, specs)
	quiz := response.Quiz
	player, _ := core.PlayerService{}.FindOrCreatePlayer(testPlayer())
	request := core.Request{
		QuizId: quiz.Id,
		Person: player,
	}
	res, _ := service.GenerateEnterGameResponse(request)
	res, _ = service.GenerateEnterGameResponse(request)
	if res.Quiz.Id != quiz.Id {
		t.Fatalf("quiz id does not match")
	}
	for _, team := range res.Snapshot.Roster {
		for _, plr := range team.Players {
			if plr.Email == player.Email {
				return
			}
		}
	}
	t.Fatalf("player not added")
}

func TestEnterGameFull(t *testing.T) {
	teardown := Setup(t)
	defer teardown(t)

	service := core.Service{}

	quizmaster, _ := core.PlayerService{}.FindOrCreatePlayer(testPlayer())
	specs := testSpecs()
	response, _ := service.GenerateCreateGameResponse(quizmaster, specs)
	quiz := response.Quiz

	for i := 0; i < specs.Players*specs.Teams; i++ {
		testPlayerJoin(quiz.Id)
	}
	_, err := testPlayerJoin(quiz.Id)
	if err == nil {
		t.Fatalf("test failed")
	}
}
