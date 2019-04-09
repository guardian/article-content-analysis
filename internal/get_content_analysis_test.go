package internal

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetContentAnalysis(t *testing.T) {
	res, err := GetContentAnalysis("/politics/2019/apr/09/tory-islamophobia-sajid-javid-facebook-anti-muslim-posts-party-members", "test")

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
