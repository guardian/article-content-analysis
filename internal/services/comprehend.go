package services

import (
    "fmt"
    "github.com/aws/aws-sdk-go/service/comprehend"
    "github.com/pkg/errors"
    "math"
    "sync"
)

func GetComprehendClient(profile string) (*comprehend.Comprehend, error) {
    sess, err := GetAwsSession(profile, "eu-west-1")

    if err != nil {
        return nil, errors.Wrap(err, "unable to create new sessions")
    }

    return comprehend.New(sess), nil
}

type ComprehendResult struct {
    Result *comprehend.DetectEntitiesOutput
    Error error
}

const ComprehendMaxChars = 5000

func GetEntitiesFromBodyText(bodyText string) ([]*comprehend.Entity, error) {
    client, err := GetComprehendClient("developerPlayground")

    if err != nil {
        return nil, errors.Wrap(err, "couldn't create client")
    }

    // Use a separate goroutine to request each chunk, and wait for each to write to the channel
    comprehendResults := make(chan ComprehendResult)
    var wg sync.WaitGroup
    wg.Add(int(math.Ceil( float64(len(bodyText)) / float64(ComprehendMaxChars) )))

    for i := 0; i < len(bodyText); {
        var end = i + ComprehendMaxChars-1

		if end >= len(bodyText) {
			//final chunk
			end = len(bodyText)-1
		} else if bodyText[end] != ' ' {
			//Avoid splitting on a word
			for j := end - 1; j >= i; j-- {
				if bodyText[j] == ' ' {
					end = j
					break
				}
			}
		}

        var chunk = bodyText[i:end]

        go func(text string) {
            defer wg.Done()

            input := &comprehend.DetectEntitiesInput{}
            input.SetText(text)
            input.SetLanguageCode("en")
            result, err := client.DetectEntities(input)
            if err != nil {
                fmt.Println("Comprehend request error", err)
            }

            comprehendResults <- ComprehendResult{result, err}
        }(chunk)

        i = end+1
    }

    go func() {
        wg.Wait()
        close(comprehendResults)
    }()

    results := make([]*comprehend.Entity, 0)
    for response := range comprehendResults {
        results = append(results, response.Result.Entities...)
    }

    return results, nil
}

func GetEntitiesFromPath(path string) ([]*comprehend.Entity, error) {
    articleFields, err := GetArticleFieldsFromCapi(path, "test")
    if err != nil {
        return nil, errors.Wrap(err, "Couldn't get article fields from CAPI for given path")
    }

    entities, err := GetEntitiesFromBodyText(articleFields.Fields.BodyText)

    if err != nil {
        return nil, errors.Wrap(err, "Error retrieving entities from body text")
    }

    return entities, nil
}
