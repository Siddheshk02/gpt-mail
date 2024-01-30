package lib

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

func Email() (string, error) {

	godotenv.Load()

	fmt.Print("> Enter the Subject for the Email : ")
	reader := bufio.NewReader(os.Stdin)
	Subject, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("> Enter the Audience for the Email : ")
	Audience, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("> Enter the Purpose for the Email : ")
	Purpose, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	prompt := fmt.Sprintf("Subject: %s Audience: %s Purpose: %s\n", Subject, Audience, Purpose)

	apiKey := os.Getenv("OPENAI_API_KEY")
	fmt.Println(apiKey)
	if apiKey == "" {
		log.Fatalln("Invalid API Key!!")
	}

	prompt_prefix := os.Getenv("Prompt_prefix")
	client := openai.NewClient(apiKey)

	// resp, err := client.CreateCompletion(
	// 	context.Background(),
	// 	openai.CompletionRequest{
	// 		Model:     openai.GPT3Dot5Turbo,
	// 		MaxTokens: 5,
	// 		Prompt:    prompt_prefix + prompt,
	// 	},
	// )

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			MaxTokens: 5,
			Model:     openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt_prefix + prompt,
				},
			},
		})

	if err != nil {
		fmt.Printf("Completion error: %v\n", err)
	}

	fmt.Println(resp.Choices[0].Message)

	return "", nil
}
