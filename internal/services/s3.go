package services

import (
	"article-content-analysis/internal/models"
)

func GetContentAnalysisFromS3(path string) (*models.ContentAnalysis, error) {
	var contentAnalysis *models.ContentAnalysis = nil

	return contentAnalysis, nil
}

func StoreContentAnalysisInS3(contentAnalysis *models.ContentAnalysis) error {
	return nil
}
