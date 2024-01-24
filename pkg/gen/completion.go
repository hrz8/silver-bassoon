package gen

import (
	"context"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
)

var DDLCached string

func AppropriateDDL(ddl string, insert string) (string, error) {
	apiKey := os.Getenv("OPEN_AI_KEY")
	if apiKey == "" {
		return ddl + insert, nil
	}

	client := openai.NewClient(apiKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: fmt.Sprintf("given this postgres DDL and unformatted INSERT INTO script: %s - %s", ddl, insert),
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "create a DDL along with the formatted INSERT INTO with the appropriate data types of each column based on the seed data. use table relation with one another (REFERENCES). consider the data type of primary key as well as the UNIQUE and NULLABLE field, replace it by the appropriate one if possible. return the result of DDL (CREATE TABLE ...) and INSERT INTO. make sure the response is pure sql with no extra comments or phrase, skip any comments",
				},
			},
		},
	)

	if err != nil {
		return "", fmt.Errorf("completion error: %v", err)
	}

	DDLCached = resp.Choices[0].Message.Content

	return DDLCached, nil
}
