package graphqlunit

import (
	"github.com/graphql-go/graphql"
	graphqltype "github.com/k197781/go-graphql/graphql/type"
	"github.com/k197781/go-graphql/models"
)

// data of User
var data = map[string]models.User{
	"1": {Id: 1, Name: "Dan"},
	"2": {Id: 2, Name: "Lee"},
	"3": {Id: 3, Name: "Nick"},
}

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: graphqltype.UserType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, isOK := p.Args["id"].(string)
					if isOK {
						return data[idQuery], nil
					}
					return nil, nil
				},
			},
		},
	})

// Schema is response schema
var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
	},
)
