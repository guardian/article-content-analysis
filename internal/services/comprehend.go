package services

import (
	"github.com/aws/aws-sdk-go/service/comprehend"
	"github.com/pkg/errors"
)

func GetComprehendClient(profile string) (*comprehend.Comprehend, error) {
	sess, err := GetAwsSession(profile, "eu-west-1")

	if err != nil {
		return nil, errors.Wrap(err, "unable to create new sessions")
	}

	return comprehend.New(sess), nil
}

func GetEntitiesFromBodyText(bodyText string) ([]*comprehend.Entity, error) {
	client, err := GetComprehendClient("developerPlayground")

	if err != nil {
		return nil, errors.Wrap(err, "couldn't create client")
	}

	input := &comprehend.DetectEntitiesInput{}
	input.SetText(bodyText)
	input.SetLanguageCode("en")
	result, err := client.DetectEntities(input)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get entities")
	}
	return result.Entities, nil
}

func GetEntitiesFromPath(path string) ([]*comprehend.Entity, error) {
	articleFields, err := GetArticleFieldsFromCapi(path, "test")
	if err != nil {
		return nil, errors.Wrap(err, "Couldn't get article fields from CAPI for given path")
	}

	var bodyText = articleFields.BodyText

	// hack to stop it failing on long articles
	if len(articleFields.BodyText) > 4999 {
		bodyText = articleFields.BodyText[0:4999]
	}
	entities, err := GetEntitiesFromBodyText(bodyText)

	if err != nil {
		return nil, errors.Wrap(err, "Error retrieving entities from body text")
	}

	return entities, nil
}
