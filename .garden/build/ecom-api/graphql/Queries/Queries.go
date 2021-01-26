package Queries

import (
	"github.com/GMRiley/ecom/ecom-api/graphql/Types"
	"github.com/graphql-go/graphql"
)

var QueryObject = graphql.NewObject(graphql.ObjectConfig{
	Name: "Queries",
	Fields: graphql.Fields{
		"Items": &graphql.Field{
			Type: ItemQueries,
		},
		"Categories": &graphql.Field{
			Type: CategoryQueries,
		},
	},
})

var ItemQueries = graphql.NewObject(graphql.ObjectConfig{
	Name: "Items",
	Fields: graphql.Fields{
		"ItemByID": &graphql.Field{
			Type: Types.ItemType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var hello = "HelloWorld"
				return func() (interface{}, error) {
					return hello, nil
				}, nil
			},
		},
	},
})

var CategoryQueries = graphql.NewObject(graphql.ObjectConfig{
	Name: "Categories",
	Fields: graphql.Fields{
		"CategoryByID": &graphql.Field{
			Type: Types.ItemType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var hello = "HelloWorld"
				return func() (interface{}, error) {
					return hello, nil
				}, nil
			},
		},
	},
})
