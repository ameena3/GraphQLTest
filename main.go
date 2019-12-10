package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

var ct []customType

type customType struct {
	Id   int
	Name string
}

var customType1 = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "customType",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.Int,
				Description: "This is the ID of the name",
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

func main() {

	//Initializing fake dataset here

	ct = []customType{customType{Id: 1, Name: "Anubhav"}, customType{Id: 2, Name: "Anubhav2"}, customType{Id: 3, Name: "Anubhav3"}, customType{Id: 4, Name: "Anubhav4"}}

	fields := graphql.Fields{
		"hello": &graphql.Field{
			Name:        "hello",
			Type:        graphql.String,
			Description: "Hey there Anubhav this is just a simple hello field",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
		"me": &graphql.Field{
			Type:        customType1,
			Description: "This is a custom me type field",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, ok := p.Args["id"].(int)
				if ok {
					for _, v := range ct {
						if v.Id == id {
							return v, nil
						}
					}
				}
				return ct, nil
			},
		},
		"list": &graphql.Field{
			Name: "listProducts",
			Type: graphql.NewList(customType1),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return ct, nil
			},
		},
	}

	rootquery := graphql.NewObject(
		graphql.ObjectConfig{
			Name:   "RootQuery",
			Fields: fields,
		})

	schconfig := graphql.SchemaConfig{
		Query: rootquery,
	}
	sch, err := graphql.NewSchema(schconfig)
	checkerr(err)

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("query")
		result := graphql.Do(graphql.Params{
			Schema:        sch,
			RequestString: query,
		})
		json.NewEncoder(w).Encode(result)
	})

	http.Handle("/", handler.New(
		&handler.Config{
			Schema:     &sch,
			Pretty:     true,
			GraphiQL:   true,
			Playground: true,
		}))

	fmt.Println("Now server is running on port 8080")
	fmt.Println("Test with Get      : curl -g 'http://localhost:8080/graphql?query={hello}'")
	http.ListenAndServe(":8080", nil)
}

func checkerr(err error) {
	if err != nil {
		log.Println("There was an error generating the schema")
		log.Fatalf(err.Error())
	}
}
