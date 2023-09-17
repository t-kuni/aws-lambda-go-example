package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"os"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	input, err := json.Marshal(name)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Input: %s\n", input)

	for _, env := range os.Environ() {
		log.Println(env)
	}
	return fmt.Sprintf("Hello %s!", name.Name), nil
}

func main() {
	lambda.Start(HandleRequest)
}
