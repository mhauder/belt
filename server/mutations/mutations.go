package mutations

import (
	fields "github.com/mhauder/belt/server/mutations/fields"

	"github.com/graphql-go/graphql"
)

var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"createNotTodo": fields.CreateNotTodo,
	},
})
