package internal

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	res := HelloWorld()
	if res != "Hello world" {
		t.Error(res)
	}
}

func TestGetEntitiesForPath(t *testing.T) {
	res, err := GetEntitiesForPath("/film/2019/apr/04/amazon-claims-woody-allen-sabotaged-films-with-metoo-comments")
	if err != nil {
		t.Error(err)
	}

	for _, entity := range res {
		fmt.Println(entity.GoString())
	}
}
