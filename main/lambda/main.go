package main

import (
	"article-entity-analysis/internal"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(func() string {
		internal.GetEntitiesForPath("/film/2019/apr/04/amazon-claims-woody-allen-sabotaged-films-with-metoo-comments")
		return "test"
	})
}
