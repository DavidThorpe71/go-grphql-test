package pkg

import (
	"testing"

	"github.com/davidthorpe71/go-grphql-test/pkg/mocks"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestFail(t *testing.T) {

	Convey(`GIVEN SomeService's SomeOperation is called`, t, func() {

		mockController := gomock.NewController(t)
		mockSomeService := mocks.NewMockSomeService(mockController)

		mockSomeService.EXPECT().SomeOperation().Return(true)

		Convey(`THEN the return value is true`, func() {

			So(DummyFunction(mockSomeService), ShouldEqual, true)

			mockController.Finish()
		})
	})
}

var fetchArticleByIdQuery = `
	query getArticleById {
		getArticleById(id: "2") {
			id
			intro
		}
	}
`

func TestGetArticleByIdQuery(t *testing.T) {

	initQL(&TestGraphQLResolvers{})
	Convey("Returns an article with the correct id", t, func() {

		Convey("the article ID should equal 1", nil)

		Convey("the article should have a name property", nil)

	})

}

func DummyFunction(someService *mocks.MockSomeService) bool {

	someService.SomeOperation()

	return false
}
