package core_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/atulanand206/find/server/core"
)

var MONGO_CLIENT_ID = "mongodb://localhost:27017/"
var GAME_DATABASE = "binquiz"

func TestGetActiveQuizzes(t *testing.T) {
	var jsonStr = []byte(`{"player_id":"104441109454181166545"}`)
	req, err := http.NewRequest(http.MethodPost, "/quizzes/active", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3MiOnRydWUsImVtYWlsIjoiYXR1bC5hbmFuZEB6ZW1vc29sYWJzLmNvbSIsImV4cGlyZXNBdCI6MTYzNTg2NTU4MCwiaXNzdWVkQXQiOiIyMDIxLTExLTAyVDE1OjM2OjIwLjM4OTU1MiswNTozMCIsIm5hbWUiOiJBdHVsIEFuYW5kIiwidXNlcklkIjoiMTEyMDAxNjg5NzQ4MDk5ODAxNzI0In0.1TbtA3SBpUest2vYodXReII11k7k1dq3_mVIgMas4ck")
	fmt.Printf("req: %v\n", req)
	rr := httptest.NewRecorder()
	handler := core.Routes(MONGO_CLIENT_ID, GAME_DATABASE)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	// expected := `[{"id":1,"first_name":"Krish","last_name":"Bhanushali","email_address":"krishsb@g.com","phone_number":"0987654321"},{"id":2,"first_name":"xyz","last_name":"pqr","email_address":"xyz@pqr.com","phone_number":"1234567890"},{"id":6,"first_name":"FirstNameSample","last_name":"LastNameSample","email_address":"lr@gmail.com","phone_number":"1111111111"}]`
	var actual []core.Game
	json.NewDecoder(rr.Body).Decode(&actual)

}
