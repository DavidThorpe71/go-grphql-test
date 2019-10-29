package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// Article scontains information about a single article
type Article struct {
	ID           string `json:"id"`
	AmbientVideo string `json:"ambientVideo"`
	StoryType    string `json:"storyType"`
	Intro        string `json:"intro"`
}

var articles = []Article{
	{
		ID:           "1",
		AmbientVideo: "AmbientVideo Test 1",
		StoryType:    "StoryType Test 1",
		Intro:        "Intro Test 1",
	},
	{
		ID:           "2",
		AmbientVideo: "AmbientVideo Test 2",
		StoryType:    "StoryType Test 2",
		Intro:        "Intro Test 2",
	},
	{
		ID:           "3",
		AmbientVideo: "AmbientVideo Test 3",
		StoryType:    "StoryType Test 3",
		Intro:        "Intro Test 3",
	},
}

var articleType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Article",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"ambientVideo": &graphql.Field{
				Type: graphql.String,
			},
			"storyType": &graphql.Field{
				Type: graphql.String,
			},
			"author": &graphql.Field{
				Type: authorType,
			},
			"intro": &graphql.Field{
				Type: graphql.String,
			},
			"businessUnit": &graphql.Field{
				Type: graphql.String,
			},
			"bodyCards": &graphql.Field{
				Type: graphql.NewList(cardType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// article := p.Source
					article, ok := p.Source.(interface{})
					if ok {
						return article, nil
					}
					// for _, item := range article.Card
					return nil, nil
				},
			},
		},
	},
)

var cardType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Card",
		Fields: graphql.Fields{
			"name": {
				Type: graphql.String,
			},
			"cardType": {
				Type: graphql.String,
			},
			"article": {
				Type: graphql.ID,
			},
		},
	},
)

var authorType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Author",
		Fields: graphql.Fields{
			// "assetImage": {
			// 	Type: imageType,
			// },
			"description": {
				Type: graphql.String,
			},
			"email": {
				Type: graphql.String,
			},
			// "links": {
			// 	Type: linkType,
			// },
			"name": {
				Type: graphql.String,
			},
			"primaryVertical": {
				Type: graphql.String,
			},
			"twitter": {
				Type: graphql.String,
			},
			// "meta": {
			// 	Type: metaType,
			// },
		},
	},
)

func getArticleByIdResolver(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(string)
	if ok {
		// Find product
		for _, article := range articles {
			if article.ID == id {
				return article, nil
			}
		}
	}
	return nil, nil
}

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			/*
				Get a single article by id
			*/
			"getArticleById": &graphql.Field{
				Type:        articleType,
				Description: "Get article by ID",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: getArticleByIdResolver,
			},
			"getAllArticles": &graphql.Field{
				Type:        graphql.NewList(articleType),
				Description: "Get all articles in database",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return articles, nil
				},
			},
		},
	},
)

func main() {
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: queryType,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})

	http.Handle("/graphql", h)
	http.ListenAndServe(":8080", nil)
}
