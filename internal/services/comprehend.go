package services

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/comprehend"
	"github.com/pkg/errors"
	"time"
)

func GetComprehendClient(profile string) (*comprehend.Comprehend, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1"),
	})

	if err != nil {
		return nil, errors.Wrap(err, "unable to create new sessions")
	}

	sess.Config.Credentials = credentials.NewChainCredentials(
		[]credentials.Provider{
			&credentials.EnvProvider{},
			&ec2rolecreds.EC2RoleProvider{
				Client:       ec2metadata.New(sess),
				ExpiryWindow: 5 * time.Minute,
			},
			&credentials.SharedCredentialsProvider{
				Profile: profile,
			},
		},
	)

	if _, err := sess.Config.Credentials.Get(); err != nil {
		return nil, errors.Wrap(err, "unable to get credentials")
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

	entities, err := GetEntitiesFromBodyText(articleFields.BodyText)

	if err != nil {
		return nil, errors.Wrap(err, "Error retrieving entities from body text")
	}

	return entities, nil
}
