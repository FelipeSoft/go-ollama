package prompt

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Input struct {
	Prompt string `json:"prompt"`
}

func PromptHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Print("PromptHandler Error: Cannot read JSON.")
	}

	var input Input
	err = json.Unmarshal(body, &input)
	if err != nil {
		fmt.Print("PromptHandler Error: Cannot unmarshal JSON.")
	}
	response := fmt.Sprintf("Received prompt: %s", input.Prompt)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}
