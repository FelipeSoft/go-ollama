package rest

import (
	"awesomeProject/rest/prompt"
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/prompt/ollama", prompt.PromptHandler).Methods("POST")

	return r
}
