package services

import "article-content-analysis/internal"

func GetContentAnalysisFromS3(path string) (*internal.ContentAnalysis, error) {
	var contentAnalysis = new(internal.ContentAnalysis)
	return contentAnalysis, nil
}

func StoreContentAnalysisInS3(contentAnalysis *internal.ContentAnalysis) (error) {
	return nil
}
