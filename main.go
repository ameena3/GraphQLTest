package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"syscall"

	database "github.com/ameena3/test/Database"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"golang.org/x/crypto/ssh/terminal"
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

var customType2 = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "customType2",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.Int,
				Description: "This is the ID of the name",
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"firstname": &graphql.Field{
				Type: graphql.String,
			},
			"lastname": &graphql.Field{
				Type: graphql.String,
			},
			"apikey": &graphql.Field{
				Type: graphql.String,
			},
			"createdat": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

var ComplianceComputerType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ComplianceComputerType",
		Fields: graphql.Fields{
			"ComplianceComputerID": &graphql.Field{
				Type:        graphql.Int,
				Description: "Unique identifier for Compliance Computer",
			},
			"ComputerName": &graphql.Field{
				Type:        graphql.String,
				Description: "Name of the device as found",
			},
			"InventoryAgent": &graphql.Field{
				Type:        graphql.String,
				Description: "The source where the inventory of the agent came from",
			},
			"AssetID": &graphql.Field{
				Type:        graphql.Int,
				Description: "The asset ID associated with this device",
			},
		},
	})

func main() {
	fmt.Println("Enter password for the database")
	bytepassword, err := terminal.ReadPassword(int(syscall.Stdin))
	checkerr(err)
	password := string(bytepassword)
	//Initializing fake dataset here
	data := database.Data{}
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

		"listdb": &graphql.Field{
			Name: "listProducts2",
			Type: graphql.NewList(customType2),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				data.ConnectToDb(password)
				_, users, err := data.GetUsers()
				checkerr(err)
				return users, nil
			},
		},

		"ComplianceComputer": &graphql.Field{
			Name: "ComplianceComputer",
			Type: ComplianceComputerType,
			Args: graphql.FieldConfigArgument{
				"ComplianceComputerID": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				ComplianceComputerID, ok := p.Args["ComplianceComputerID"].(int)
				if ok {
					err := data.ConnectToDb(password)
					checkerr(err)
					count, retdata, err := data.GetComplianceComputerByComplianceComputerID(ComplianceComputerID)
					checkerr(err)
					if count == 0 {
						log.Printf("There was no record of this id")
						return nil, nil
					}

					return retdata, nil
				}
				return nil, nil

			},
		},
		"ListComplianceComputer": &graphql.Field{
			Name: "ComplianceComputer",
			Type: graphql.NewList(ComplianceComputerType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {

				err := data.ConnectToDb(password)
				checkerr(err)
				count, retdata, err := data.GetListOfComplianceComputer()
				checkerr(err)
				log.Printf("returned rowcount is %d \n", count)
				return retdata, nil

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
		log.Printf(err.Error())
	}
}
