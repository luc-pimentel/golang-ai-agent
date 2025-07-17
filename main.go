package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

type Chatbot struct {
	client  *openai.Client
	history []openai.ChatCompletionMessage
}

func NewChatbot(apiKey string) *Chatbot {
	return &Chatbot{
		client:  openai.NewClient(apiKey),
		history: make([]openai.ChatCompletionMessage, 0),
	}
}

func (c *Chatbot) AddMessage(role, content string) {
	message := openai.ChatCompletionMessage{
		Role:    role,
		Content: content,
	}
	c.history = append(c.history, message)
}

func (c *Chatbot) Chat(userInput string) (string, error) {
	c.AddMessage(openai.ChatMessageRoleUser, userInput)

	resp, err := c.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: c.history,
		},
	)

	if err != nil {
		return "", err
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no response from OpenAI")
	}

	assistant_response := resp.Choices[0].Message.Content
	c.AddMessage(openai.ChatMessageRoleAssistant, assistant_response)

	return assistant_response, nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY not found in environment")
	}

	chatbot := NewChatbot(apiKey)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("AI Chatbot initialized. Type 'quit' to exit.")
	fmt.Print("> ")

	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		
		if input == "quit" {
			break
		}
		
		if input == "" {
			fmt.Print("> ")
			continue
		}

		response, err := chatbot.Chat(input)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Printf("Bot: %s\n", response)
		}

		fmt.Print("> ")
	}

	fmt.Println("Goodbye!")
}