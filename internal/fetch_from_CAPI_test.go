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

func TestGetArticleFieldsFromPath(t *testing.T) {
	res := GetArticleFieldsFromPath("/film/2019/apr/04/amazon-claims-woody-allen-sabotaged-films-with-metoo-comments", "test")
	fmt.Print(res)
}

func TestGetEntitiesForPath(t *testing.T) {
	GetEntitiesForPath("/film/2019/apr/04/amazon-claims-woody-allen-sabotaged-films-with-metoo-comments")
}
