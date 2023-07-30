package openai_manager

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

//type OpenAI interface {
//	MakeRequest(content string) (string, error)
//}

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

func (s *OpenAIManager) MakeRequest(language, code string) (string, error) {
	systemMessage := fmt.Sprintf("You will be provided with %s code, and your task is to calculate its time complexity.", language)

	req := openai.ChatCompletionRequest{
		Model: s.model,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: code,
			},
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: systemMessage,
			},
		},
		Temperature: 0,
		MaxTokens:   256,
	}

	resp, err := s.client.CreateChatCompletion(context.Background(), req)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
