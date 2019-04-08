package models

import "github.com/aws/aws-sdk-go/service/comprehend"

type ArticleFields struct {
	Headline string `json:"headline"`
	Byline   string `json:"byline"`
	BodyText string `json:"bodyText"`
}

type Gender int
const (
	Male	Gender = 0
	Female	Gender = 1
)

type Byline struct {
	Name string
	Gender Gender
}

type Person struct {
	comprehend.Entity
	Gender Gender
}

type ContentAnalysis struct {
	Path string
	Headline string
	BodyText string
	Bylines []Byline
	People []Person
	Locations []comprehend.Entity
	Organisations []comprehend.Entity
	CreativeWorkTitles []comprehend.Entity
	CommercialItem []comprehend.Entity
	Events []comprehend.Entity
}
