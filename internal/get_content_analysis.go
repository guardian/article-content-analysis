package internal

import (
	"article-content-analysis/internal/models"
	"article-content-analysis/internal/services"
	"github.com/aws/aws-sdk-go/service/comprehend"
	"github.com/pkg/errors"
)

func ConstructContentAnalysis(path string, content *models.Content, entities []*comprehend.Entity, cacheHit bool) *models.ContentAnalysis {
	byline := models.Byline{
		Name: content.Fields.Byline,
	}

	var people []*models.Person = nil
	var locations []*comprehend.Entity = nil
	var organisations []*comprehend.Entity = nil
	var creativeWorkTitles []*comprehend.Entity = nil
	var commercialItems []*comprehend.Entity = nil
	var events []*comprehend.Entity = nil

	for _, entity := range entities {
		if *entity.Type == "PERSON" {
			people = append(people, &models.Person{Entity: *entity})
		}
		if *entity.Type == "LOCATION" {
			locations = append(locations, entity)
		}
		if *entity.Type == "ORGANIZATION" {
			organisations = append(organisations, entity)
		}
		if *entity.Type == "COMMERCIAL_ITEM" {
			commercialItems = append(commercialItems, entity)
		}
		if *entity.Type == "TITLE" {
			creativeWorkTitles = append(creativeWorkTitles, entity)
		}
		if *entity.Type == "EVENT" {
			events = append(events, entity)
		}

	}

	contentAnalysis := models.ContentAnalysis{
		Path:               path,
		Headline:           content.Fields.Headline,
		BodyText:           content.Fields.BodyText,
		Bylines:            []*models.Byline{&byline},
		People:             people,
		Locations:          locations,
		Organisations:      organisations,
		CreativeWorkTitles: creativeWorkTitles,
		CommercialItems:    commercialItems,
		Events:             events,
		CacheHit:           cacheHit,
	}

	return &contentAnalysis
}

func GetContentAnalysis(path string, capiKey string) (*models.ContentAnalysis, error) {
	contentAnalysis, err := services.GetContentAnalysisFromS3(path) //will return error if object is not in s3

	if contentAnalysis != nil {
		contentAnalysis.CacheHit = true
		return contentAnalysis, nil
	}

	articleFields, err := services.GetArticleFieldsFromCapi(path, capiKey)

	if err != nil {
		return nil, errors.Wrap(err, "Couldn't get article fields from CAPI for given path")
	}

	entities, err := services.GetEntitiesFromPath(path)

	if err != nil {
		return nil, errors.Wrap(err, "Couldn't get entities for given path")
	}

	contentAnalysis = ConstructContentAnalysis(path, articleFields, entities, false)

	storeContentAnalysisInS3Error := services.StoreContentAnalysisInS3(contentAnalysis)

	if storeContentAnalysisInS3Error != nil {
		return nil, errors.Wrap(storeContentAnalysisInS3Error, "Could not store in S3")
	}

	return contentAnalysis, nil
}
