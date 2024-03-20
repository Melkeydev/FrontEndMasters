package main

import (
	"fmt"
	"lambda-func/app"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Username string `json:"username"`
}

// Define our Handler
func HandleRequest(event MyEvent) (string, error) {
	if event.Username == "" {
		return "Invalid Request", fmt.Errorf("username cannot be empty")
	}

	return fmt.Sprintf("Successfully called by - %s", event.Username), nil
}

func main() {
	lambdaApp := app.NewApp()
	lambda.Start(lambdaApp.ApiHandler.RegisterUser)
}
