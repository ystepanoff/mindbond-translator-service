package openai

import (
	"context"
	"fmt"
	go_openai "github.com/sashabaranov/go-openai"
)

type Client struct {
	Client *go_openai.Client
}

func Init(token string) Client {
	client := Client{go_openai.NewClient(token)}
	return client
}

func (c *Client) Translate(message string, fromLang string, toLang string) (string, error) {
	response, err := c.Client.CreateChatCompletion(
		context.Background(),
		go_openai.ChatCompletionRequest{
			Model: go_openai.GPT4,
			Messages: []go_openai.ChatCompletionMessage{
				{
					Role:    go_openai.ChatMessageRoleUser,
					Content: fmt.Sprintf("Translate the following text from %s to %s: \"%s\"", fromLang, toLang, message),
				},
			},
		},
	)
	if err != nil {
		return "", err
	}
	return response.Choices[0].Message.Content, nil
}
