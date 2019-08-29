package mutations

import (
	"context"
	mongo "github.com/mhauder/belt/server/data"
	"github.com/mhauder/belt/server/types"

	"github.com/graphql-go/graphql"
)

type guidelineStruct struct {
	TITLE       string `json:"name"`
	DESCRIPTION string `json:"description"`
}

var CreateGuideline = &graphql.Field{
	Type:        types.Guideline,
	Description: "Create a guideline",
	Args: graphql.FieldConfigArgument{
		"title": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"description": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},

	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		// get our params
		title, _ := params.Args["title"].(string)
		description, _ := params.Args["description"].(string)
		notTodoCollection := mongo.Client.Database("medium-app").Collection("Not_Todos")
		_, err := notTodoCollection.InsertOne(context.Background(), map[string]string{"title": title, "description": description})

		if err != nil {
			panic(err)
		}

		return guidelineStruct{title, description}, nil
	},
}
