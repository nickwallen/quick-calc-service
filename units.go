package service

import (
	"fmt"
	"github.com/bcicen/go-units"
)

// Property A physical property that is described by a unit of measure.
type Property string

// Physical properties that can be measured.
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

// NewProperty Create a new property.
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

// String Stringify a Property.
func (p Property) String() string {
	return string(p)
}

// System All units are part of a system of measurement.
type System string

// The systems of measurement.
const (
	None     = "none"
	Metric   = "metric"
	Imperial = "imperial"
	US       = "us"
)

// NewSystem Create a new system of measurement.
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

// String Stringify a System.
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

// NewUnit Creates a Unit from the name of the unit.
func NewUnit(input string) (Unit, error) {
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

// Name The name of a Unit. Required for resolution with graphql.
func (u Unit) Name() string {
	return u.name
}

// PluralName the plural name of a Unit. Required for resolution with graphql.
func (u Unit) PluralName() string {
	return u.pluralName
}

// MeasureOf The unit measures which physical property? Required for resolution with graphql.
func (u Unit) MeasureOf() *string {
	measureOf := u.measureOf.String()
	return &measureOf
}

// PartOf The unit is 'part-of' what system of measurement? Required for resolution with graphql.
func (u Unit) PartOf() *string {
	partOf := u.partOf.String()
	return &partOf
}

// Result The result of evaluating an expression.
type Result struct {
	value float64
	units Unit
}

// Value The value of the result. Required for resolution with graphql.
func (r Result) Value() *float64 {
	return &r.value
}

// Units The units of the result. Required for resolution with graphql.
func (r Result) Units() *Unit {
	return &r.units
}
