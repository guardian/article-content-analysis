package internal

import (
	"article-content-analysis/internal/models"
	"article-content-analysis/internal/services"
	"github.com/aws/aws-sdk-go/service/comprehend"
	"github.com/pkg/errors"
)

func ConstructContentAnalysis(articleFields *models.ArticleFields, entities []*comprehend.Entity) *models.ContentAnalysis {
	return new(models.ContentAnalysis)
}

func GetContentAnalysis(path string) (*models.ContentAnalysis, error) {
	var contentAnalysis = new(models.ContentAnalysis)
	contentAnalysis, err := services.GetContentAnalysisFromS3(path)
	if err != nil {
		return contentAnalysis, errors.Wrap(err, "Could'nt get article fields for given article")
	}

	if contentAnalysis != nil {
		return contentAnalysis, nil
	}

	articleFields, err := services.GetArticleFieldsFromCapi(path, "test")

	if err != nil {
		return contentAnalysis, errors.Wrap(err, "Couldn't get article fields from CAPI for given path")
	}

	entities, err := services.GetEntitiesFromPath(path)

	contentAnalysis = ConstructContentAnalysis(articleFields, entities)

	storeContentAnalysisInS3Error := services.StoreContentAnalysisInS3(contentAnalysis)

	if storeContentAnalysisInS3Error != nil {
		panic("Could not store in S3")
	}

	return contentAnalysis, nil
}
