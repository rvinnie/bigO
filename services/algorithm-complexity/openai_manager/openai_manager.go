package openai_manager

import (
	"context"
	"github.com/sashabaranov/go-openai"
)

const (
	Temperature = 0
	MaxTokens   = 128
)

type OpenAIManager struct {
	client *openai.Client
	model  string
}

func NewOpenAIManager(client *openai.Client, model string) *OpenAIManager {
	return &OpenAIManager{
		client: client,
		model:  model,
	}
}

func (s *OpenAIManager) MakeRequest(userMessage, systemMessage string) (string, error) {
	req := openai.ChatCompletionRequest{
		Model: s.model,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: userMessage,
			},
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: systemMessage,
			},
		},
		Temperature: Temperature,
		MaxTokens:   MaxTokens,
	}

	resp, err := s.client.CreateChatCompletion(context.Background(), req)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
