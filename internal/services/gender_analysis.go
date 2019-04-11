package services

import (
	"article-content-analysis/internal/models"
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

func GetGenderAnalysis(name string) (*models.GenderAnalysis, error) {
	values := map[string]string{"name": name}

	jsonValue, _ := json.Marshal(values)

	resp, err := http.Post("https://wr0ih1jbs3.execute-api.eu-west-1.amazonaws.com/PROD/getEntities", "application/json", bytes.NewBuffer(jsonValue))
	defer resp.Body.Close()

	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshall s3 data")
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, errors.Wrap(err, "could not read body test")
	}
	var genderAnalysis = new(models.GenderAnalysis)
	//TODO: validate response
	genderAnalysisError := json.Unmarshal(body, &genderAnalysis)
	if genderAnalysisError != nil {
		return nil, errors.Wrap(genderAnalysisError, "could not parse response from CAPI")
	}
	return genderAnalysis, nil
}
