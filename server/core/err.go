package core

const (
	Err_RequestNotDecoded = "Request can't be decoded."
	Err_IndexNotDecoded   = "Index can't be decoded."
	Err_MatchNotDecoded   = "Match information can't be decoded."
	Err_PlayerNotDecoded  = "Player information can't be decoded."

	Err_CollectionsNotCreated  = "Collections can't be created."
	Err_DeckDtosNotCreated     = "Deck dtos can't be created."
	Err_MatchRequestNotCreated = "New match request can't be created."
	Err_MatchNotCreated        = "New match can't be created."
	Err_PlayerNotCreated       = "Player can't be created."
	Err_QuestionNotCreated     = "Question can't be created."
	Err_AnswerNotCreated       = "Answer can't be created."

	Err_IndexNotSeeded     = "Index can't be seeded."
	Err_QuestionsNotSeeded = "Questions can't be seeded."
	Err_AnswersNotSeeded   = "Answers can't be seeded."

	Err_IndexNotPresent    = "Index does not exist."
	Err_MatchNotPresent    = "Match with the given info does not exist."
	Err_PlayerNotPresent   = "Player with the given info does not exist."
	Err_QuestionNotPresent = "Question does not exist."
	Err_AnswerNotPresent   = "Answer does not exist."

	Err_MatchNotUpdated = "Match can't be updated."

	Err_CollectionsNotDropped = "Collections can't be dropped."

	Err_QuizmasterCantPlay = "Quizmaster can't join the match as a player."
	Err_WaitingForPlayers  = "Waiting for more players to join."
)
