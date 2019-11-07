package schema

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/graphql-go/graphql"
)

var MySchema graphql.Schema

func init() {

	AdditionalArticleFields()

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"getArticleById": &graphql.Field{
				Type:        ArticleType,
				Description: "Get article by ID",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, _ := p.Args["id"].(string)
					fmt.Printf(`GraphQL query arg: %q`, id)

					//TODO: Add call to real Database here
					data, _ := ioutil.ReadFile("mock.json")

					var newArticle Article
					err := json.Unmarshal(data, &newArticle)

					if err != nil {
						fmt.Printf(`Unable to read input json file %v`, err)
					}

					return newArticle, nil
				},
			},
		},
	})

	MySchema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
}
