package services

import (
	"fmt"
	"testing"
)

func TestGetGenderAnalysis(t *testing.T) {
	res, err := GetGenderAnalysis("Jonny Rankin")

	if err != nil {
		t.Error(err)
	} else {
		for _, people := range res.People {
			fmt.Println(people)
		}
	}
}
