package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

var (
	ecsClient   *ecs.Client
	serviceName string
	clusterName string
)

type Body struct {
	Name string `json:"name"`
}

func router(ctx context.Context, event events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {

	authorization := event.Headers["Authorization"]

	payload := strings.TrimPrefix(authorization, "Basic ")

	decoded, err := base64.StdEncoding.DecodeString(payload)

	if err != nil {
		return events.LambdaFunctionURLResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       http.StatusText(http.StatusInternalServerError),
		}, nil
	}

	authParts := strings.SplitN(string(decoded), ":", 2)

	if authParts[0] != "edu" || authParts[1] != "edu" {
		return events.LambdaFunctionURLResponse{
			StatusCode: http.StatusUnauthorized,
			Body:       http.StatusText(http.StatusUnauthorized),
		}, nil
	}

	var b Body

	if err := json.Unmarshal([]byte(event.Body), &b); err != nil {
		return events.LambdaFunctionURLResponse{
			StatusCode: http.StatusBadRequest,
			Body:       http.StatusText(http.StatusBadRequest),
		}, nil
	}

	if event.RequestContext.HTTP.Path == "/test" {
		return events.LambdaFunctionURLResponse{
			StatusCode: http.StatusOK,
			Body:       fmt.Sprintf("testPath%s", b.Name),
		}, nil
	}

	return events.LambdaFunctionURLResponse{
		StatusCode: http.StatusOK,
		Body:       b.Name,
	}, nil
}

func main() {
	lambda.Start(router)
}
