package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type notifierImpl struct {
}

type Notifier interface {
	Post(fullURL string, payload any, response any) error
}

func NewNotifierService() Notifier {
	return &notifierImpl{}
}

func (m *notifierImpl) Post(fullURL string, payload any, response any) error {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	// Create the request first
	req, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Set Content-Type header
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Log response body untuk debugging
	bodyBytes, _ := io.ReadAll(resp.Body)
	log.Printf("[Core API] Response Status: %s", resp.Status)
	log.Printf("[Core API] Response Body: %s", string(bodyBytes))

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("received non-2xx response status: %s", resp.Status)
	}

	// Define a wrapper struct to extract the "data" field from the response.
	// json.RawMessage is used to delay the unmarshaling of the nested data.
	var coreResponse struct {
		Data json.RawMessage `json:"data"`
	}

	// First, unmarshal the full response into the wrapper.
	if err := json.Unmarshal(bodyBytes, &coreResponse); err != nil {
		return fmt.Errorf("failed to decode core response wrapper: %w", err)
	}

	// Then, unmarshal the content of the "data" field into the final response struct.
	if err := json.Unmarshal(coreResponse.Data, response); err != nil {
		return fmt.Errorf("failed to decode nested 'data' from response: %w", err)
	}

	return nil
}
