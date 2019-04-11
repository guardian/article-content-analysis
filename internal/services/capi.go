package services

import (
	"article-content-analysis/internal/models"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetArticleFieldsFromCapi(path string, apiKey string) (*models.Content, error) {
	var articleFields = new(models.Content)
	urlPrefix := "https://content.guardianapis.com"
	urlSuffix := "?api-key=" + apiKey + "&show-fields=byline,bodyText,headline"
	url := urlPrefix + path + urlSuffix
	resp, err := http.Get(url)
	if err != nil {
		return articleFields, errors.Wrap(err, "no response from CAPI")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return articleFields, errors.Wrap(err, "could not read body test")
	}

	//TODO: validate response
	fields := gjson.Get(string(body), "response.content").Raw
	fieldsBytes := []byte(fields)
	articleFieldsError := json.Unmarshal(fieldsBytes, &articleFields)
	if articleFieldsError != nil {
		return nil, errors.Wrap(articleFieldsError, "could not parse response from CAPI")
	}
	return articleFields, nil
}

func GetArticleFieldsFromCapiForDateRange(fromDate string, endDate string, apiKey string) (*models.CapiSearchResponse, error) {
	var fullResults = new(models.CapiSearchResponse)
	var reachedEndOfResults = false
	var pageIndex = 1

	for reachedEndOfResults == false {
		var capiReponse = new(models.CapiSearchResponse)

		urlPrefix := "https://content.guardianapis.com/search"
		urlSuffix := "?api-key=" + apiKey + "&from-date=" + fromDate + "&to-date=" + endDate + "&page-size=200" + "&page=" + strconv.Itoa(pageIndex)
		url := urlPrefix + urlSuffix
		fmt.Println(url)
		resp, err := http.Get(url)
		if err != nil {
			return nil, errors.Wrap(err, "no response from CAPI")
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.Wrap(err, "could not read body test")
		}

		//TODO: validate response
		fields := gjson.Get(string(body), "response").Raw
		fieldsBytes := []byte(fields)
		capiResponseError := json.Unmarshal(fieldsBytes, &capiReponse)
		if capiResponseError != nil {
			return nil, errors.Wrap(capiResponseError, "could not parse response from CAPI")
		}
		if capiReponse.Status == "error" {
			reachedEndOfResults = true
		} else {
			for _, result := range capiReponse.Results {
				result.Id = "/" + result.Id
				fullResults.Results = append(fullResults.Results, result)
				pageIndex++

			}
		}

	}
	return fullResults, nil
}
