package services

import (
	"fmt"
	"testing"
)

func TestGetEntitiesFromPath(t *testing.T) {
	res, err := GetEntitiesFromPath("/commentisfree/2019/apr/08/workers-rights-survive-brexit-labour-demand-more")

	if err != nil {
		t.Error(err)
	} else {
		for _, entity := range res {
			fmt.Println(entity.GoString())
		}
	}
}
