package main

import (
	"article-entity-analysis/internal"
	"encoding/json"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"net/http"
)

type Path struct {
	Path string `json:"path"`
}

type APIGatewayProxyResponse struct {
	StatusCode      int               `json:"statusCode"`
	Headers         map[string]string `json:"headers"`
	Body            string            `json:"body"`
	IsBase64Encoded bool              `json:"isBase64Encoded,omitempty"`
}

func HandleRequest(ctx context.Context, path Path) (APIGatewayProxyResponse, error) {
	var resString string
	var res APIGatewayProxyResponse
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

	return APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(body),
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
