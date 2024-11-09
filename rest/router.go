package rest

import (
	"awesomeProject/rest/prompt"
	"net/http"
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	r := mux.NewRouter()
	// Add the handler for the prompt endpoint
	r.HandleFunc("/prompt/ollama", prompt.PromptHandler).Methods("POST")
	r.HandleFunc("/", home)

	return r
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../static/index.html")
}
