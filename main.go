package main

import (
	"log"
	"net/http"

	"github.com/davidthorpe71/go-grphql-test/schema"
	"github.com/graphql-go/handler"
)

func main() {

	h := handler.New(&handler.Config{
		Schema:     &schema.MySchema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})

	http.Handle("/graphql", h)
	log.Print("Ready: listening on port 4040...\n")
	http.ListenAndServe(":4040", nil)
}
