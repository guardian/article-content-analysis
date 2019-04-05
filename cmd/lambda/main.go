package main

import (
	"article-entity-analysis/internal"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/pkg/errors"
	"net/http"
)

type Path struct {
	Path string `json:"path"`
}

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var resString string
	var res events.APIGatewayProxyResponse
	var path = new(Path)
	var requestBodyError = json.Unmarshal([]byte(request.Body), &path)
	if requestBodyError != nil {
		return res, errors.Wrap(requestBodyError, "Could not parse json body")
	}
	entities, err := internal.GetEntitiesForPath(path.Path)
	if err != nil {
		return res, errors.Wrap(err, "Did not manage to get entities for path")
	}
	for _, entity := range entities {
		resString += entity.GoString()
	}

	body, err := json.Marshal(resString)
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
