package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"github.com/GMRiley/ecom/ecom-api/graphql/Queries"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"k8s.io/klog"
)

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: Queries.QueryObject,
		//Mutation: graphql.NewObject(Mutations.BuildMutations()),
	},
)

func health(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "healthy\n")
}

func main() {
	klog.InitFlags(nil)
	flag.Set("log_file", "ecom-api.log")
	flag.Parse()

	klog.Infof("creating graphql handler")

	var resultHandler handler.ResultCallbackFn = func(ctx context.Context, params *graphql.Params, result *graphql.Result, responseBody []byte) {

	}

	// Attempt GraphQL Endpoint
	h := handler.New(&handler.Config{
		Schema:           &schema,
		Pretty:           true,
		GraphiQL:         false,
		Playground:       true,
		ResultCallbackFn: resultHandler,
	})

	http.Handle("/graphql", h)
	http.HandleFunc("/health", health)
	http.ListenAndServe(":8080", nil)

	klog.Flush()
}
