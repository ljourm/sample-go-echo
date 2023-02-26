package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	echoAdapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
)

func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	router := NewRouter()

	echoLambda := echoAdapter.NewV2(router)

	return echoLambda.ProxyWithContext(ctx, req)
}

func main() {
	// router := NewRouter()
	// router.Logger.Fatal(router.Start(":1323"))

	lambda.Start(Handler)
}
