package internal

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetContentAnalysis(t *testing.T) {
	res, err := GetContentAnalysisForPath("/lifeandstyle/2019/apr/01/five-ways-to-do-squats-safely", "test")

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

func TestGetContentAnalysisForDateRange(t *testing.T) {
	res, err := GetContentAnalysisForDateRange("2019-04-10", "2019-04-11", "test")

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
