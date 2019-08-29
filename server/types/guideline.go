package types

import (
	"github.com/graphql-go/graphql"
)

var Guideline = graphql.NewObject(graphql.ObjectConfig{
	Name: "Guideline",
	Fields: graphql.Fields{
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"data": &graphql.Field{
			Type: graphql.String,
		},
		"beltId": &graphql.Field{
			Type: graphql.Int,
		},
		"type": &graphql.Field{
			Type: graphql.Int,
		},
		"rationale": &graphql.Field{
			Type: graphql.String,
		},
		"solution": &graphql.Field{
			Type: graphql.String,
		},
		"priority": &graphql.Field{
			Type: graphql.Int,
		},
		"severity": &graphql.Field{
			Type: graphql.Int,
		},
		"archived": &graphql.Field{
			Type: graphql.Boolean,
		},
		"testMode": &graphql.Field{
			Type: graphql.Boolean,
		},
		"resources": &graphql.Field{
			Type: graphql.String,
		},
		"createdByUserId": &graphql.Field{
			Type: graphql.Int,
		},
		"createdAt": &graphql.Field{
			Type: graphql.DateTime,
		},
		"lastUpdatedByUserId": &graphql.Field{
			Type: graphql.Int,
		},
		"updatedAt": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})
