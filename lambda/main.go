package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type  MyEvent struct {
	Name string `json:"what is your name?`
	Age int `json: "age"?`
}

type MyResponse struct {
	Message string `json:"Ans"`
}

func HandleLambda(event MyEvent)(MyResponse, error) {
	return MyResponse{Message: fmt.Sprintf("%s is %d old", event.Name, event.Age)}, nil
} 

func main() {
	lambda.Start(HandleLambda)
}