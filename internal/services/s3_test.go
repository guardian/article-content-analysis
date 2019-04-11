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
	res, err := services.GetContentAnalysisFromS3("/commentisfree/2019/apr/08/wall-street-socialism-jpmorgan-jamie-dimon-bailout")
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(res.Path)
	}
}

func TestStoreContentAnalysisInS3(t *testing.T) {
	contentFields := models.ContentFields{"test_headline", "test_byline", "test_body"}
	content := models.Content{"2019-04-11T13:35:13Z", "football", contentFields, "test"}
	var events []*comprehend.Entity = nil
	contentAnalysis := internal.ConstructContentAnalysis(
		"/commentisfree/2019/apr/08/wall-street-socialism-jpmorgan-jamie-dimon-bailout",
		&content,
		events,
		false,
	)

	err := services.StoreContentAnalysisInS3(contentAnalysis)
	if err != nil {
		t.Error(err)
	}
}
