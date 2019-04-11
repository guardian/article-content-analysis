package models

import "github.com/aws/aws-sdk-go/service/comprehend"

type ContentFields struct {
    Headline string `json:"headline"`
    Byline   string `json:"byline"`
    BodyText string `json:"bodyText"`
}
type Content struct {
    WebPublicationDate string `json:"webPublicationDate"`
    Section            string `json:"sectionId"`
    Fields             ContentFields `json:"fields"`
}
type Gender int

const (
    Unknown Gender = 0
    Female  Gender = 1
    Male    Gender = 2
)

type Byline struct {
    Name   string `json:"name"`
    Gender Gender `json:"gender"`
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
    Events             []*comprehend.Entity `json:"events"`
    CacheHit           bool                 `json:"cacheHit"`
    WebPublicationDate string               `json:"webPublicationDate"`
    Section            string               `json:"section"`
}
