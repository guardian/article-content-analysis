package models

import "github.com/aws/aws-sdk-go/service/comprehend"

type ArticleFields struct {
	Headline string `json:"headline"`
	Byline   string `json:"byline"`
	BodyText string `json:"bodyText"`
}

type Gender int

const (
	Unknown Gender = 0
	Male    Gender = 1
	Female  Gender = 2
)

type Byline struct {
	Name   string
	Gender Gender
}

type Person struct {
	comprehend.Entity
	Gender Gender
}

type ContentAnalysis struct {
	Path               string               `json:"path"`
	Headline           string               `json:"headline"`
	BodyText           string               `json:"bodyText"`
	Bylines            []*Byline            `json:"bylines"`
	People             []*Person            `json:"people"`
	Locations          []*comprehend.Entity `json:"locations"`
	Organisations      []*comprehend.Entity `json:"organisations"`
	CreativeWorkTitles []*comprehend.Entity `json:"creativeWorkTitles"`
	CommercialItems    []*comprehend.Entity `json:"commercialItems"`
	Events             []*comprehend.Entity `json:"evnets"`
}
