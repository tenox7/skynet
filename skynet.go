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
	// Read API key and model name from environment variables
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("Please set the OPENAI_API_KEY environment variable.")
	}

	modelName := os.Getenv("OPENAI_MODEL_NAME")
	if modelName == "" {
		log.Fatal("Please set the OPENAI_MODEL_NAME environment variable.")
	}

	// Create OpenAI client
	client := openai.NewClient(apiKey)

	// Set up color functions
	promptColor := color.New(color.FgRed).SprintFunc()
	responseColor := color.New(color.FgGreen).SprintFunc()

	// Set up readline for interactive prompt with history
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
		// Read user input with prompt editing and history
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

		// Skip empty input
		if line == "" {
			continue
		}

		// Prepare the chat message
		message := openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: line,
		}

		// Create chat completion request
		req := openai.ChatCompletionRequest{
			Model:    modelName,
			Messages: []openai.ChatCompletionMessage{message},
		}

		// Send request to OpenAI API
		resp, err := client.CreateChatCompletion(context.Background(), req)
		if err != nil {
			fmt.Printf("API error: %v\n", err)
			continue
		}

		// Display the assistant's reply
		fmt.Println(responseColor(resp.Choices[0].Message.Content))
	}
}
