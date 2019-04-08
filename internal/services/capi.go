package services

type ArticleFields struct {
	Headline string `json:"headline"`
	Byline   string `json:"byline"`
	BodyText string `json:"bodyText"`
}

func GetArticleFieldsFromCapi(path string) (*ArticleFields, error) {
	var articleFields = new(ArticleFields)
	return articleFields, nil
}