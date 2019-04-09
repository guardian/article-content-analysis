package services_test

import (
	"article-content-analysis/internal"
	"article-content-analysis/internal/models"
	"article-content-analysis/internal/services"
	"fmt"
	"github.com/aws/aws-sdk-go/service/comprehend"
	"testing"
)

func TestGetContentAnalysisFromS3(t *testing.T) {
	res, err := services.GetContentAnalysisFromS3("/commentisfree/2019/apr/08/workers-rights-survive-brexit-labour-demand-more")
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(res.Path)
	}
}

func TestStoreContentAnalysisInS3(t *testing.T) {
	articleFields := models.ArticleFields{"test_headline","test_byline","test_body"}
	var events []*comprehend.Entity = nil
	contentAnalysis := internal.ConstructContentAnalysis(
		"/commentisfree/2019/apr/08/workers-rights-survive-brexit-labour-demand-more",
		&articleFields,
		events,
	)

	err := services.StoreContentAnalysisInS3(contentAnalysis)
	if err != nil {
		t.Error(err)
	}
}
