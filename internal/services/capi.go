package services

import (
	"article-content-analysis/internal/models"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
)

func GetArticleFieldsFromCapi(path string, apiKey string) (*models.ArticleFields, error) {
	var articleFields = new(models.ArticleFields)
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
	fields := gjson.Get(string(body), "response.content.fields").Raw
	fieldsBytes := []byte(fields)
	articleFieldsError := json.Unmarshal(fieldsBytes, &articleFields)
	if articleFieldsError != nil {
		return nil, errors.Wrap(articleFieldsError, "could not parse response from CAPI")
	}
	return articleFields, nil
}
