package main

import (
	"fmt"
	"github.com/bcicen/go-units"
)

// Property A physical property that is described by a unit of measure.
type Property string

const (
	Length      Property = "length"
	Mass        Property = "mass"
	Volume      Property = "volume"
	Density     Property = "density"
	Time        Property = "time"
	Temperature Property = "temperature"
	Pressure    Property = "pressure"
	Bytes       Property = "bytes"
)

func NewProperty(name string) (Property, error) {
	var prop Property
	switch name {
	case Length.String():
		prop = Length
	case Mass.String():
		prop = Mass
	case Volume.String():
		prop = Volume
	case Density.String():
		prop = Density
	case Time.String():
		prop = Time
	case Temperature.String():
		prop = Temperature
	case Pressure.String():
		prop = Pressure
	case Bytes.String():
		prop = Bytes
	default:
		return prop, fmt.Errorf("no physical property named '%s'", name)
	}
	return prop, nil
}

func (p Property) String() string {
	return string(p)
}

// System All units are part of a system of measurement.
type System string

const (
	None     = "none"
	Metric   = "metric"
	Imperial = "imperial"
	US       = "us"
)

func NewSystem(name string) (System, error) {
	var system System
	switch name {
	case Metric:
		system = Metric
	case Imperial:
		system = Imperial
	case US:
		system = US
	case None, "":
		system = None
	default:
		return system, fmt.Errorf("no system named '%s'", name)
	}
	return system, nil
}

func (s System) String() string {
	return string(s)
}

// Unit A unit of measurement that describes some physical property.
type Unit struct {
	name       string
	pluralName string
	measureOf  Property
	partOf     System
}

func (u Unit) Name() string {
	return u.name
}

func (u Unit) PluralName() string {
	return u.pluralName
}

func (u Unit) MeasureOf() *string {
	measureOf := u.measureOf.String()
	return &measureOf
}

func (u Unit) PartOf() *string {
	partOf := u.partOf.String()
	return &partOf
}

func UnitFromString(input string) (Unit, error) {
	var result Unit
	u, err := units.Find(input)
	if err != nil {
		return result, nil
	}
	name := u.Name
	if name == "" {
		return result, fmt.Errorf("invalid unit name '%s'", name)
	}
	pluralName := u.PluralName()
	if pluralName == "" {
		return result, fmt.Errorf("invalid plural unit name '%s'", pluralName)
	}
	property, err := NewProperty(u.Quantity)
	if err != nil {
		return result, err
	}
	system, err := NewSystem(u.System())
	if err != nil {
		return result, err
	}
	result = Unit{name, pluralName, property, system}
	return result, nil
}

// Result The result of evaluating an expression.
type Result struct {
	value float64
	units Unit
}

func (r Result) Value() *float64 {
	return &r.value
}

func (r Result) Units() *Unit {
	return &r.units
}
