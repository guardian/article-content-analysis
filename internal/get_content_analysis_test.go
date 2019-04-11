package internal

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetContentAnalysis(t *testing.T) {
	res, err := GetContentAnalysis("/uk-news/2019/apr/11/meghan-and-harry-want-to-celebrate-birth-of-baby-in-private", "test")

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
