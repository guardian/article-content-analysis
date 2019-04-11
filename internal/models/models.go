package models

import "github.com/aws/aws-sdk-go/service/comprehend"

type ContentFields struct {
	Headline string `json:"headline"`
	Byline   string `json:"byline"`
	BodyText string `json:"bodyText"`
}

type GenderAnalysis struct {
	People []struct {
		Text        string `json:"text"`
		Normal      string `json:"normal	"`
		FirstName   string `json:"firstName"`
		MiddleName  string `json:"middleName"`
		NickName    string `json:"nickName"`
		LastName    string `json:"lastName"`
		GenderGuess string `json:"genderGuess"`
		Pronoun     string `json:"pronoun"`
	} `json:"people"`
}
type Content struct {
	WebPublicationDate string        `json:"webPublicationDate"`
	Section            string        `json:"sectionId"`
	Fields             ContentFields `json:"fields"`
	Id                 string        `json:"id"`
}

type CapiSearchResponse struct {
	Status  string
	Results []Content
}

type Gender string

type Byline struct {
	Name   string `json:"name"`
	Gender Gender `json:"gender"`
}

type Person struct {
	comprehend.Entity
	Gender Gender
}

//Note - entity offsets may not be accurate for articles longer than 5000 chars
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
