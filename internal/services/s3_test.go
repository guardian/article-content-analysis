package services_test

import (
	"fmt"
	"testing"
	"article-content-analysis/internal/services"
)

func TestGetContentAnalysisFromS3(t *testing.T) {
	res, err := services.GetContentAnalysisFromS3("/commentisfree/2019/apr/08/workers-rights-survive-brexit-labour-demand-more")
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(res.Path)
	}
}