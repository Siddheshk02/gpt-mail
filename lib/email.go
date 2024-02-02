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

	fmt.Print("> Enter the Purpose/Information for the Email : ")
	Purpose, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	prompt := fmt.Sprintf("Subject: %s Audience: %s Purpose/Information: %s\n", Subject, Audience, Purpose)

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatalln("Invalid API Key!!")
	}

	prompt_prefix := os.Getenv("Prompt_prefix")

	client := openai.NewClient(apiKey)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    "user",
					Content: prompt_prefix + prompt,
				},
			},
		})

	if err != nil {
		fmt.Printf("Completion error: %v\n", err)
	}

	res := resp.Choices[0].Message.Content

	return res, nil
}
