package main

import (
	"article-entity-analysis/internal"
	"github.com/aws/aws-lambda-go/lambda"
	"golang.org/x/net/context"
)

type Path struct {
	Path string `json:"path"`
}

func HandleRequest(ctx context.Context, path Path) (string, error) {
	entities, err := internal.GetEntitiesForPath(path.Path)
	if err != nil {
		return "", err
	}
	var res = ""
	for _, entity := range entities {
		res += entity.GoString()
	}
	return res, nil
}

func main() {
	lambda.Start(HandleRequest)
}
