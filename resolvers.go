package service

import (
	"fmt"
	"github.com/bcicen/go-units"
	"github.com/graph-gophers/graphql-go"
	calc "github.com/nickwallen/quick-calc"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"path/filepath"
)

// Resolver The base resolver required for graphql resolution.
type Resolver struct{}

// Schema Returns the graphQL schema to use.
func Schema() (schema *graphql.Schema, err error) {
	bytes, err := readFile("./schema.graphql")
	if err != nil {
		return schema, err
	}
	schemaDef := string(bytes)
	schema = graphql.MustParseSchema(schemaDef, &Resolver{}, graphql.UseStringDescriptions())
	return schema, nil
}

func getSchema(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// UnitByName Finds a unit by its name. Required for graphql resolution.
func (r *Resolver) UnitByName(args struct{ Name string }) (*Unit, error) {
	var result Unit
	result, err := NewUnit(args.Name)
	if err != nil {
		return &result, err
	}
	return &result, nil
}

// Units Finds all known units. Required for graphql resolution.
func (r *Resolver) Units() (*[]*Unit, error) {
	var result []*Unit
	for _, theirUnit := range units.All() {
		ourUnit, err := NewUnit(theirUnit.Name)
		if err != nil {
			return &result, err
		}
		result = append(result, &ourUnit)
	}
	return &result, nil
}

// Evaluate an expression. Required for graphql resolution.
func (r *Resolver) Evaluate(args struct{ Expr string }) (Result, error) {
	var result Result
	amount, err := calc.CalculateAmount(args.Expr)
	if err != nil {
		return result, err
	}

	unitName := amount.Units
	unit, err := NewUnit(unitName)
	if err != nil {
		return result, err
	}

	result = Result{amount.Value, unit, args.Expr}
	log.Info(fmt.Sprintf("evaluate(%s) = %.2f %s", args.Expr, result.value, result.units.pluralName))
	return result, nil
}

func readFile(relativePath string) (contents []byte, err error) {
	path, _ := filepath.Abs(relativePath)
	if err != nil {
		return contents, err
	}
	page, err := ioutil.ReadFile(path)
	if err != nil {
		return contents, err
	}
	return page, nil
}
