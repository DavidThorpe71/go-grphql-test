package schema

import (
	"errors"
	"testing"

	"github.com/graphql-go/graphql"
	. "github.com/smartystreets/goconvey/convey"
)

func TestArticleResolver(t *testing.T) {
	Convey(`ARTICLE FIELD RESOLVER TESTS`, t, func() {

		m := ArticleType.Fields()

		for _, i := range m {
			mockResolveParams := graphql.ResolveParams{
				Source: Article{
					Title:       "test title",
					Description: "test description",
					RelatedArticles: []Article{
						{
							Title:       "test related article title",
							Description: "test related article description",
						},
					},
				},
				Info: graphql.ResolveInfo{
					FieldName: i.Name,
				},
			}
			switch i.Name {
			case "title":
				Convey(`Expect title field to resolve`, func() {
					have, _ := ArticleResolver(mockResolveParams)
					want := `test title`
					So(have, ShouldResemble, want)
				})
			case "description":
				Convey(`Expect description field to resolve`, func() {
					have, _ := ArticleResolver(mockResolveParams)
					want := `test description`
					So(have, ShouldResemble, want)
				})
			case "relatedArticles":
				Convey(`Expect relatedArticles field to resolve`, func() {
					have, _ := ArticleResolver(mockResolveParams)
					want := []Article{
						{
							Title:       "test related article title",
							Description: "test related article description",
						},
					}
					So(have, ShouldResemble, want)
				})
			}

		}

		Convey(`No field found`, func() {
			mockResolveParams := graphql.ResolveParams{
				Source: Article{
					Title:       "test title",
					Description: "test description",
					RelatedArticles: []Article{
						{
							Title:       "test related article title",
							Description: "test related article description",
						},
					},
				},
				Info: graphql.ResolveInfo{
					FieldName: "fakeField",
				},
			}
			_, have := ArticleResolver(mockResolveParams)
			want := errors.New(`No field with name fakeField`)
			So(have, ShouldResemble, want)
		})
	})
}
