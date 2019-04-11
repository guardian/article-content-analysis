package services

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetArticleFieldsFromCapi(t *testing.T) {
	res, err := GetArticleFieldsFromCapi("/commentisfree/2019/apr/08/workers-rights-survive-brexit-labour-demand-more", "test")
	if err != nil {
		t.Error(err)
	} else {
		res, err := json.Marshal(res)
		if err != nil {
			t.Error(err)
		} else {
			fmt.Println(string(res))
		}
	}
}

func TestGetArticleFieldsFromCapiForDateRange(t *testing.T) {
	res, err := GetArticleFieldsFromCapiForDateRange("2018-01-02", "2018-01-02", "test")
	if err != nil {
		t.Error(err)
	} else {
		res, err := json.Marshal(res)
		if err != nil {
			t.Error(err)
		} else {
			fmt.Println(string(res))
		}
	}
}
