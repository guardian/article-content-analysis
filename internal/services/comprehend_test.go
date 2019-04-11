package services

import (
	"fmt"
	"testing"
)

func TestGetEntitiesFromPath(t *testing.T) {
	res, err := GetEntitiesFromPath("/artanddesign/2019/feb/21/from-hepworth-to-rodin-uk-sculpture-collection-to-be-catalogued-online")

	if err != nil {
		t.Error(err)
	} else {
		for _, entity := range res {
			fmt.Println(entity.GoString())
		}
	}
}
