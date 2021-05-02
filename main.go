package main

import (
	"net/http"

	"github.com/sh-git-repo/Cortnus-v0-backend/api_v1/welcome"
)

func main() {
	http.HandleFunc("/", welcome.WelcomeIndex)
	http.HandleFunc("/api/v1/welcome", welcome.WelcomePlayer)
	http.ListenAndServe(":8080", nil)
}
