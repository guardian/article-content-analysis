package main

import (
	"article-content-analysis/internal"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/pkg/errors"
	"net/http"
)

type Input struct {
	Path string `json:"path"`
	CapiKey string `json:"capiKey"`
}

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var res events.APIGatewayProxyResponse
	var input = new(Input)
	var requestBodyError = json.Unmarshal([]byte(request.Body), &input)

	if requestBodyError != nil {
		return res, errors.Wrap(requestBodyError, "Could not parse json body")
	}
	contentAnalysis, err := internal.GetContentAnalysis(input.Path, input.CapiKey)
	if err != nil {
		return res, errors.Wrap(err, "Did not manage to get entities for path")
	}

	body, err := json.Marshal(contentAnalysis)
	if err != nil {
		return res, errors.Wrap(err, "Could not marshall body")
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(body),
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
