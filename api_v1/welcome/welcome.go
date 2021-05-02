package welcome

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sh-git-repo/Cortnus-v0-backend/api_v1/loggingutils"
)

func WelcomeIndex(writer http.ResponseWriter, request *http.Request) {
	/*
		Welcome page for the index route.
	*/
	fmt.Fprintf(
		writer,
		"Welcome to home of Cortnus the local time is %s",
		time.Now(),
	)

	loggingutils.LogResponse("WelcomeIndex", 200)
}

func WelcomePlayer(writer http.ResponseWriter, request *http.Request) {
	/*
		Welcome page for the player.
	*/
	query := request.URL.Query()
	player_name_array, present := query["player"]

	var player_name string
	player_name = player_name_array[0]

	var status int
	var message string

	if (!present) || (len(player_name) < 1) {

		status = http.StatusBadRequest
		message = `{"data": "Player name not in query."}`

	} else {

		status = http.StatusOK
		message = fmt.Sprintf(`{"data": "Welcome %s, Cortnus awaits."}`, player_name)

	}

	loggingutils.LogResponse("WelcomePlayer", status)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	writer.Write([]byte(message))
}
