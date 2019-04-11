package utils

import (
	"fmt"
	"testing"
)

func TestFormatDate(t *testing.T) {
	res, err := FormatDate("2019-04-11T13:35:13Z")

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(res)
	}
}
