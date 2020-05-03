package main

import (
	"github.com/bcicen/go-units"
	calc "github.com/nickwallen/quick-calc"
)

type Resolver struct{}

func (r *Resolver) UnitByName(args struct{ Name string }) (*Unit, error) {
	var result Unit
	result, err := UnitFromString(args.Name)
	if err != nil {
		return &result, err
	}
	return &result, nil
}

func (r *Resolver) Units() (*[]*Unit, error) {
	var result []*Unit
	for _, theirUnit := range units.All() {
		ourUnit, err := UnitFromString(theirUnit.Name)
		if err != nil {
			return &result, err
		}
		result = append(result, &ourUnit)
	}
	return &result, nil
}

func (r *Resolver) Evaluate(args struct{ Expr string }) (Result, error) {
	var result Result
	amount, err := calc.CalculateAmount(args.Expr)
	if err != nil {
		return result, err
	}

	unitName := amount.Units.String()
	unit, err := UnitFromString(unitName)
	if err != nil {
		return result, err
	}

	result = Result{amount.Value, unit}
	return result, nil
}
