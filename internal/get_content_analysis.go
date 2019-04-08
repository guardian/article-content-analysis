package internal

import (
	"article-content-analysis/internal/services"
	"github.com/aws/aws-sdk-go/service/comprehend"
	"github.com/pkg/errors"
)

type Gender int
const (
	Male	Gender = 0
	Female	Gender = 1
)

type Byline struct {
	Name string
	Gender Gender
}

type Person struct {
	comprehend.Entity
	Gender Gender
}

type ContentAnalysis struct {
	Path string
	Headline string
	BodyText string
	Bylines []Byline
	People []Person
	Locations []comprehend.Entity
	Organisations []comprehend.Entity
	CreativeWorkTitles []comprehend.Entity
	CommercialItem []comprehend.Entity
	Events []comprehend.Entity
}

func ConstructContentAnalysis(articleFields *services.ArticleFields, entities *[]comprehend.Entity) (*ContentAnalysis) {
	return new(ContentAnalysis)
}

func GetContentAnalysis(path string) (*ContentAnalysis, error) {
	var contentAnalysis = new(ContentAnalysis)
	contentAnalysis, err := services.GetContentAnalysisFromS3(path)
	if err != nil {
		return contentAnalysis, errors.Wrap(err, "Could'nt get article fields for given article")
	}

	if contentAnalysis != nil {
		return contentAnalysis, nil
	}

	articleFields, err := services.GetArticleFieldsFromCapi(path)
	if err != nil {
		return contentAnalysis, err
	}

	entities, err := services.GetEntitiesFromComprehend(articleFields.BodyText)

	contentAnalysis = ConstructContentAnalysis(articleFields, entities)

	storeContentAnalysisInS3Error := services.StoreContentAnalysisInS3(contentAnalysis)

	if storeContentAnalysisInS3Error != nil {
		panic("Could not store in S3")
	}

	return contentAnalysis, nil
}

