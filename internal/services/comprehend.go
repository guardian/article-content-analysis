package services

import "github.com/aws/aws-sdk-go/service/comprehend"

func GetEntitiesFromComprehend(bodyText string) (*[]comprehend.Entity, error) {
	var entities = new([]comprehend.Entity)
	return entities, nil
}