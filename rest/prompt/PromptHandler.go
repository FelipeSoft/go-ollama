package prompt

import "net/http"

func PromptHandler(w http.ResponseWriter, r *http.Request) {
	response := "response from ollama"
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}
