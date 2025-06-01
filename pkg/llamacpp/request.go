package llamacpp

import (
	"fmt"
)

type LlamaRequest struct {
	MaxTokens int64            `json:"max_tokens"`
	Messages  []RequestMessage `json:"messages"`
}

type RequestMessage struct {
	Role    string    `json:"role"`
	Content []Content `json:"content"`
}

type Content struct {
	Type     string    `json:"type"`
	Text     *string   `json:"text,omitempty"`
	ImageURL *ImageURL `json:"image_url,omitempty"`
}

type ImageURL struct {
	URL string `json:"url"`
}

func NewRecognizerLlamaRequest(maxTokens int64, category string, base64image string) *LlamaRequest {
	recognizerPrompt := fmt.Sprintf("Response with a single word to the question, without additional symbols or information. What type of %s is the main subject in the image?", category)

	return &LlamaRequest{
		MaxTokens: maxTokens,
		Messages: []RequestMessage{
			{
				Role: "user",
				Content: []Content{
					{
						Type: "text",
						Text: &recognizerPrompt,
					},
					{
						Type: "image_url",
						ImageURL: &ImageURL{
							URL: base64image,
						},
					},
				},
			},
		},
	}
}

func NewCounterLlamaRequest(maxTokens int64, category string, base64image string) *LlamaRequest {
	counterPrompt := fmt.Sprintf("Count the number of elements of type %s which appears in the image. Is very important to be precise with this number. The response should be a digit, not a word. How many %ss appears in the image?", category, category)

	return &LlamaRequest{
		MaxTokens: maxTokens,
		Messages: []RequestMessage{
			{
				Role: "user",
				Content: []Content{
					{
						Type: "text",
						Text: &counterPrompt,
					},
					{
						Type: "image_url",
						ImageURL: &ImageURL{
							URL: base64image,
						},
					},
				},
			},
		},
	}
}
