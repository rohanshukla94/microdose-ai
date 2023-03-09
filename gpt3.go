package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/0x9ef/openai-go"
	"log"
	"microdose-ai/summary"
)

func MakeRequest(question string, link string) (CsvData, error) {

	apiInstance := openai.New(summary.Gpt3ApiKey())

	ctx := context.Background()

	completionResp, err := apiInstance.Completion(ctx, &openai.CompletionOptions{
		// Choose model, you can see list of available models in models.go file
		Model: openai.ModelTextDavinci003,
		// Number of completion tokens to generate response. By default - 1024
		MaxTokens: 1200,
		// Text to completion
		Prompt: []string{question},
	})

	gptSummary := completionResp.Choices[0].Text

	if err != nil {
		fmt.Println(err)
	}

	respImage, err := apiInstance.ImageCreate(context.Background(), &openai.ImageCreateOptions{
		Prompt: "Generate an Image describing what the link is trying to say https://www.1mg.com/drugs/augmentin-625-duo-tablet-138629",
		Size:   openai.SizeMedium,
	})
	if err != nil {
		log.Fatal(err)
	}
	b, err := json.MarshalIndent(respImage, "", "  ")

	if err != nil {
		log.Fatal(err)
	}
	//else {
	//	//log.Println(string(b))
	//}
	if err != nil {
		fmt.Println(err)
	}

	csvData := CsvData{
		ID:              "1234",
		ProductLink:     link,
		Prompt:          question,
		MetaDescription: gptSummary,
		HindiTranslate:  gptSummary,
		Image:           string(b),
	}

	return csvData, nil

}
