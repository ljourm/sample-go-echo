package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	echoAdapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/joho/godotenv"
)

func initEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Print(".env is not found")
	}
}

func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	router := NewRouter()

	echoLambda := echoAdapter.NewV2(router)

	return echoLambda.ProxyWithContext(ctx, req)
}

func main() {
	initEnv()

	if(os.Getenv("IS_LOCAL") == "true") {
		router := NewRouter()
		router.Logger.Fatal(router.Start(":1323"))
	} else {
		lambda.Start(Handler)
	}
}
