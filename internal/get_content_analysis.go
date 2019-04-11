package internal

import (
	"article-content-analysis/internal/models"
	"article-content-analysis/internal/services"
	"fmt"
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
		Section:            content.Section,
		WebPublicationDate: content.WebPublicationDate,
	}

	return &contentAnalysis
}

func AddGenderToContentAnalysis(contentAnalysis *models.ContentAnalysis) (*models.ContentAnalysis, error) {
	for _, person := range contentAnalysis.People {
		genderAnalysis, err := services.GetGenderAnalysis(*person.Text)

		if err != nil {
			return nil, errors.Wrap(err, "Error getting gender analysis for "+*person.Text)
		}

		if len(genderAnalysis.People) > 0 {
			if genderAnalysis.People[0].GenderGuess == "Female" {
				person.Gender = "Female"
			}
			if genderAnalysis.People[0].GenderGuess == "Male" {
				person.Gender = "Male"
			}
		}
	}

	for _, person := range contentAnalysis.Bylines {
		genderAnalysis, err := services.GetGenderAnalysis(person.Name)

		if err != nil {
			return nil, errors.Wrap(err, "Error getting gender analysis for byline "+person.Name)
		}

		if len(genderAnalysis.People) > 0 {
			if genderAnalysis.People[0].GenderGuess == "Female" {
				person.Gender = "Female"
			}
			if genderAnalysis.People[0].GenderGuess == "Male" {
				person.Gender = "Male"
			}
		}
	}

	return contentAnalysis, nil

}

func GetContentAnalysisForPath(path string, capiKey string) (*models.ContentAnalysis, error) {
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

	contentAnalysisWithGender, err := AddGenderToContentAnalysis(contentAnalysis)

	if err != nil {
		return nil, errors.Wrap(err, "Error adding gender to people")
	}

	storeContentAnalysisInS3Error := services.StoreContentAnalysisInS3(contentAnalysisWithGender)

	if storeContentAnalysisInS3Error != nil {
		return nil, errors.Wrap(storeContentAnalysisInS3Error, "Could not store in S3")
	}

	return contentAnalysis, nil
}

func GetContentAnalysisForDateRange(fromDate string, endDate string, apiKey string) ([]*models.ContentAnalysis, error) {
	articlesInRange, err := services.GetArticleFieldsFromCapiForDateRange(fromDate, endDate, apiKey)

	if err != nil {
		return nil, errors.Wrap(err, "Could not get articles from CAPI")
	}

	var results []*models.ContentAnalysis = nil

	for _, article := range articlesInRange.Results {
		analysis, err := GetContentAnalysisForPath(article.Id, apiKey)

		if err != nil {
			fmt.Println("did not work for article " + article.Id)
		} else {
			fmt.Println("worked for article " + article.Id)
			results = append(results, analysis)
		}
	}

	return results, nil

}
