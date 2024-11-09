package prompt

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// OllamaBodyRequest structure for sending prompt to the backend service
type OllamaBodyRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

// Output structure for decoding the response from the backend
type Output struct {
	Response string `json:"response"`
}

// PromptWorker processes the prompt and streams the output using SSE
func PromptWorker(prompt string, w http.ResponseWriter, flusher http.Flusher) (string, error) {
	url := "http://localhost:11434/api/generate"

	// Prepare the request body
	body, err := json.Marshal(OllamaBodyRequest{
		Model:  "llama3.2",
		Prompt: prompt,
	})
	if err != nil {
		return "", fmt.Errorf("cannot parse struct to JSON: %w", err)
	}

	// Make the HTTP POST request
	response, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("cannot complete the request for Ollama: %w", err)
	}
	defer response.Body.Close()

	// Check if the response status is OK
	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error with response from Ollama, status code: %d", response.StatusCode)
	}

	// Use bufio.Scanner to read the response body line-by-line
	scanner := bufio.NewScanner(response.Body)
	var fullOutput string

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Received line:", line)

		// Decode the JSON response into Output struct
		var output Output
		if err := json.Unmarshal([]byte(line), &output); err != nil {
			return "", fmt.Errorf("error decoding response: %w", err)
		}

		// Log the decoded response
		fmt.Println("Decoded response:", output.Response)

		// Write the response to the client using SSE format
		_, err := w.Write([]byte("data: " + output.Response + "\n\n"))
		if err != nil {
			return "", fmt.Errorf("failed to write data to response: %w", err)
		}

		// Flush the response to ensure it's sent immediately
		flusher.Flush()

		// Concatenate the response to fullOutput for later return
		fullOutput += output.Response
	}

	// Check if there was an error while reading the response
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading response: %w", err)
	}

	// Return the full output after processing all lines
	return fullOutput, nil
}
