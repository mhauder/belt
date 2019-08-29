package queries

import (
	"context"
	"github.com/graphql-go/graphql"
	mongo "github.com/mhauder/belt/server/data"
	"go.mongodb.org/mongo-driver/bson"
	"time"

	"github.com/mhauder/belt/server/types"
)

type guidelineStruct struct {
	TITLE       string `json:"name"`
	DESCRIPTION string `json:"title"`
}

var GetGuidelines = &graphql.Field{
	Type:        graphql.NewList(types.Guideline),
	Description: "Get all guidelines",
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		notTodoCollection := mongo.Client.Database("medium-app").Collection("Not_Todos")
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
		todos, err := notTodoCollection.Find(ctx, bson.D{})

		if err != nil {
			panic(err)
		}
		var guidelineList []guidelineStruct
		for todos.Next(context.Background()) {
			keys, err := todos.Current.Elements()

			if err != nil {
				panic(err)
			}

			todo := guidelineStruct{}
			for _, key := range keys {
				keyString := key.Key()
				valueString := key.Value().String()

				switch keyString {
				case "title":
					todo.TITLE = valueString
				case "description":
					todo.DESCRIPTION = valueString
				default:
				}
			}
			guidelineList = append(guidelineList, todo)
		}

		return guidelineList, nil
	},
}
