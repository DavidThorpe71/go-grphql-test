package schema

import (
	"errors"

	"github.com/graphql-go/graphql"
)

type Article struct {
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	RelatedArticles []Article `json:"relatedArticles"`
}

var ArticleType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Article",
	Description: "Article object type",
	Fields: graphql.Fields{
		"title": &graphql.Field{
			Type:    graphql.String,
			Resolve: ArticleResolver,
		},
		"description": &graphql.Field{
			Type:    graphql.String,
			Resolve: ArticleResolver,
		},
	},
})

func AdditionalArticleFields() {
	ArticleType.AddFieldConfig("relatedArticles", &graphql.Field{
		Type:    graphql.NewList(ArticleType),
		Resolve: ArticleResolver,
	})
}

func ArticleResolver(p graphql.ResolveParams) (interface{}, error) {
	data, ok := p.Source.(Article)

	if ok {

		fieldName := p.Info.FieldName

		switch fieldName {
		case "title":
			return data.Title, nil
		case "description":
			return data.Title, nil
		case "relatedArticles":
			return data.RelatedArticles, nil
		default:
			return nil, errors.New(`No field with name ` + fieldName)
		}
	}
	return nil, nil
}
