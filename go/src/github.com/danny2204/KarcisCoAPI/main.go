package main

import (
	"github.com/danny2204/KarcisCoAPI/connection"
	"github.com/danny2204/KarcisCoAPI/gql/mutation"
	"github.com/danny2204/KarcisCoAPI/gql/queries"
	"github.com/danny2204/KarcisCoAPI/middleware"
	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"log"
	"net/http"
)
func main() {

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:        queries.GetRoot(),
		Mutation:     mutation.GetRoot(),
	})

	if err != nil {
		panic(err)
	}

	h := handler.New(&handler.Config{
		Schema:           &schema,
		Pretty:           true,
		GraphiQL:         true,
		Playground:       true,
	})

	wrapped := middleware.CorsMiddleware(h)

	r := mux.NewRouter()
	r.Use(middleware.LogMiddleware)
	r.Handle("/api", wrapped)

	log.Fatal(http.ListenAndServe(":" + connection.ApiPort, r))

}