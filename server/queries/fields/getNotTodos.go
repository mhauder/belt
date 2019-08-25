package queries

import (
	"context"
	"github.com/graphql-go/graphql"
	mongo "github.com/mhauder/belt/server/data"
	"go.mongodb.org/mongo-driver/bson"
	"time"

	"github.com/mhauder/belt/server/types"
)

type todoStruct struct {
	NAME        string `json:"name"`
	DESCRIPTION string `json:"description"`
}

var GetNotTodos = &graphql.Field{
	Type:        graphql.NewList(types.NotTodo),
	Description: "Get all not todos",
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		notTodoCollection := mongo.Client.Database("medium-app").Collection("Not_Todos")
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
		todos, err := notTodoCollection.Find(ctx, bson.D{})

		if err != nil {
			panic(err)
		}
		var todosList []todoStruct
		for todos.Next(context.Background()) {
			keys, err := todos.Current.Elements()

			if err != nil {
				panic(err)
			}

			todo := todoStruct{}
			for _, key := range keys {
				keyString := key.Key()
				valueString := key.Value().String()

				switch keyString {
				case "name":
					todo.NAME = valueString
				case "description":
					todo.DESCRIPTION = valueString
				default:
				}
			}
			todosList = append(todosList, todo)
		}

		return todosList, nil
	},
}
