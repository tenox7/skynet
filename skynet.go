package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/chzyer/readline"
	"github.com/fatih/color"
	openai "github.com/sashabaranov/go-openai"
)

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("Please set the OPENAI_API_KEY environment variable.")
	}

	modelName := os.Getenv("OPENAI_MODEL_NAME")
	if modelName == "" {
		log.Fatal("Please set the OPENAI_MODEL_NAME environment variable.")
	}

	client := openai.NewClient(apiKey)
	promptColor := color.New(color.FgRed).SprintFunc()
	responseColor := color.New(color.FgGreen).SprintFunc()

	rl, err := readline.NewEx(&readline.Config{
		Prompt:          promptColor(modelName + "> "),
		HistoryFile:     "/tmp/readline.tmp",
		AutoComplete:    nil,
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	})
	if err != nil {
		log.Fatalf("Failed to initialize readline: %v", err)
	}
	defer rl.Close()

	fmt.Println("Interactive OpenAI API Client (type 'exit' or press Ctrl+D to quit)")

	for {
		line, err := rl.Readline()
		if err != nil {
			if err == readline.ErrInterrupt { // Handle Ctrl+C
				if len(line) == 0 {
					break
				} else {
					continue
				}
			} else if err == io.EOF { // Handle Ctrl+D
				break
			} else {
				log.Fatalf("Readline error: %v", err)
			}
		}

		if line == "exit" {
			break
		}

		if line == "" {
			continue
		}

		message := openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: line,
		}

		req := openai.ChatCompletionRequest{
			Model:    modelName,
			Messages: []openai.ChatCompletionMessage{message},
		}

		resp, err := client.CreateChatCompletion(context.Background(), req)
		if err != nil {
			fmt.Printf("API error: %v\n", err)
			continue
		}

		fmt.Println(responseColor(resp.Choices[0].Message.Content))
	}
}
