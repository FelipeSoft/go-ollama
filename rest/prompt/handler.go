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
	// Set headers for SSE (Server-Sent Events)
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Ensure we can flush the response immediately
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	// Read the request body (prompt)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Cannot read request body", http.StatusBadRequest)
		return
	}

	// Unmarshal the input to get the prompt
	var input Input
	err = json.Unmarshal(body, &input)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Process the prompt and stream output to the client
	response, err := PromptWorker(input.Prompt, w, flusher)
	if err != nil {
		http.Error(w, "Failed to process prompt", http.StatusInternalServerError)
		return
	}

	// Optionally print full response (for debugging or logging purposes)
	fmt.Print(response)
}
