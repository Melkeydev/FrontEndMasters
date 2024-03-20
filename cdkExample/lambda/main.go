package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

const (
	invalidRequest = "Invalid Request"
)

type MyEvent struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

func HandleRequest(ctx context.Context, event MyEvent) (string, error) {
	if event.Name == "" {
		return invalidRequest, fmt.Errorf("name cannot be empty in request")
	}

	return fmt.Sprintf("Successful call - thank you %s", event.Name), nil
}

func main() {
	lambda.Start(HandleRequest)
}
