package main

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/gqltesting"
	"github.com/stretchr/testify/assert"
	"testing"
)

func schema(t *testing.T) *graphql.Schema {
	schemaDef, err := getSchema("./schema.graphql")
	assert.Nil(t, err)
	return graphql.MustParseSchema(schemaDef, &Resolver{}, graphql.UseStringDescriptions())
}

func runTest(t *testing.T, query string, expected string) {
	gqltesting.RunTest(t, &gqltesting.Test{
		Schema:         schema(t),
		Query:          query,
		ExpectedResult: expected,
	})
}

func TestEvaluate(t *testing.T) {
	testCases := map[string]string{
		// test case
		`
			{
				evaluate(expr:"2 liters + 4 ml in ml") {
					value
					units {
						name
						pluralName
						measureOf
						partOf
					}
				}
			}
		`: `
			{
				"evaluate": {
					"value": 2004,
					"units": {
						"name": "milliliter",
						"pluralName": "milliliters",
						"measureOf": "volume",
						"partOf": "metric"
					}
				}
			}
		`,
		// test case
		`
			{
				evaluate(expr:"2 liters + 4 ml in ml") {
					value
					units {
						pluralName
					}
				}
			}
		`: `
			{
				"evaluate": {
					"value": 2004,
					"units": {
						"pluralName": "milliliters"
					}
				}
			}
		`,
	}
	for query, expected := range testCases {
		runTest(t, query, expected)
	}
}

func TestUnitsByName(t *testing.T) {
	testCases := map[string]string{
		// test case
		`
			{
				unitByName(name:"liters") {
					name
					pluralName
						measureOf
						partOf
					}
			}
		`: `
			{
				"unitByName": {
					"name": "liter",
					"pluralName": "liters",
					"measureOf": "volume",
					"partOf": "none"
				}
			}
		`,
	}
	for query, expected := range testCases {
		runTest(t, query, expected)
	}
}
