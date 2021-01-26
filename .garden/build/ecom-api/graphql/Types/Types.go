package Types

import (
	"github.com/graphql-go/graphql"
)

var ItemType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:   "Item",
		Fields: ItemFields,
	},
)

var CategoryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:   "Category",
		Fields: CategoryFields,
	},
)

var ItemFields = graphql.Fields{
	"id": &graphql.Field{Type: graphql.String},
}

var CategoryFields = graphql.Fields{
	"id": &graphql.Field{Type: graphql.String},
}
