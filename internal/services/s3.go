package services

import (
	"article-content-analysis/internal/models"
)

func GetContentAnalysisFromS3(path string) (*models.ContentAnalysis, error) {
	var contentAnalysis = new(models.ContentAnalysis)
	return contentAnalysis, nil
}

func StoreContentAnalysisInS3(contentAnalysis *models.ContentAnalysis) (error) {
	return nil
}
