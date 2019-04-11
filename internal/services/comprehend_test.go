package services

import (
    "fmt"
    "testing"
)

func TestGetEntitiesFromPath(t *testing.T) {
    res, err := GetEntitiesFromPath("/uk-news/2019/apr/11/julian-assange-arrested-at-ecuadorian-embassy-wikileaks")

    if err != nil {
        t.Error(err)
    } else {
        for _, entity := range res {
            fmt.Println(entity.GoString())
        }
    }
}
