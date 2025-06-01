package llamacpp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Client struct {
	url       string
	maxTokens int64
	client    *http.Client
}

func NewClient(baseUrl string, maxTokens int64) *Client {
	return &Client{
		url:       composeURL(baseUrl),
		maxTokens: maxTokens,
		client:    &http.Client{},
	}
}

func (lc *Client) RecognizeSubject(category string, base64Image string) (*string, error) {
	if category == "" {
		return nil, fmt.Errorf("category cannot be empty")
	}
	if base64Image == "" {
		return nil, fmt.Errorf("base64Image cannot be empty")
	}

	payload := NewRecognizerLlamaRequest(lc.maxTokens, category, base64Image)

	serverResponse, err := lc.sendRequest(payload)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}

	var llamaResponse LlamaResponse
	if err := json.Unmarshal(*serverResponse, &llamaResponse); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %v", err)
	}

	if len(llamaResponse.Choices) == 0 {
		return nil, fmt.Errorf("no choices in response")
	}

	content := llamaResponse.Choices[0].Message.Content
	cleanContent := strings.ToUpper(strings.ReplaceAll(content, ".", ""))

	return &cleanContent, nil
}

func (lc *Client) CountSubjects(category string, base64Image string) (*string, error) {
	if category == "" {
		return nil, fmt.Errorf("category cannot be empty")
	}
	if base64Image == "" {
		return nil, fmt.Errorf("base64Image cannot be empty")
	}

	payload := NewCounterLlamaRequest(lc.maxTokens, category, base64Image)

	serverResponse, err := lc.sendRequest(payload)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}

	var llamaResponse LlamaResponse
	if err := json.Unmarshal(*serverResponse, &llamaResponse); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %v", err)
	}

	if len(llamaResponse.Choices) == 0 {
		return nil, fmt.Errorf("no choices in response")
	}

	content := llamaResponse.Choices[0].Message.Content
	cleanContent := strings.ToUpper(strings.ReplaceAll(content, ".", ""))

	return &cleanContent, nil
}

func (lc *Client) sendRequest(payload *LlamaRequest) (*[]byte, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error marshaling JSON: %v", err)
	}

	req, err := http.NewRequest("POST", lc.url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := lc.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error executing request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return &body, nil
}

func composeURL(baseURL string) string {
	return fmt.Sprintf("%s%s", baseURL, "/v1/chat/completions")
}
