package main

import (
	"github.com/bcicen/go-units"
	calc "github.com/nickwallen/quick-calc"
)

// Resolver The base resolver required for graphql resolution.
type Resolver struct{}

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

	unitName := amount.Units.String()
	unit, err := NewUnit(unitName)
	if err != nil {
		return result, err
	}

	result = Result{amount.Value, unit}
	return result, nil
}
