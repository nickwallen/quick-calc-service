package service

import (
	"github.com/bcicen/go-units"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewProperty(t *testing.T) {
	cases := map[string]Property{
		"meters":      Length,
		"kilometers":  Length,
		"pounds":      Mass,
		"stones":      Mass,
		"milliliters": Volume,
		"gallons":     Volume,
		"hours":       Time,
		"seconds":     Time,
		"f":           Temperature,
		"celsius":     Temperature,
	}
	for name, phyProp := range cases {
		meters, err := units.Find(name)
		result, err := NewProperty(meters.Quantity)
		assert.Equal(t, phyProp, result)
		assert.Nil(t, err)
	}
}

func TestNewProperty_Error(t *testing.T) {
	_, err := NewProperty("invalidUnit")
	assert.NotNil(t, err)
}

func TestNewSystem(t *testing.T) {
	cases := map[string]System{
		"meters":      Metric,
		"kilometers":  Metric,
		"pounds":      Imperial,
		"stones":      Imperial,
		"milliliters": Metric,
		"gallons":     Imperial,
		"hours":       None,
		"seconds":     None,
		"f":           US,
		"celsius":     Metric,
	}
	for name, system := range cases {
		units, err := units.Find(name)
		result, err := NewSystem(units.System())
		assert.Equal(t, system, result, "For '%s' expected '%s', but got '%s'", name, system, result)
		assert.Nil(t, err)
	}
}

func TestNewSystem_Error(t *testing.T) {
	_, err := NewSystem("invalid")
	assert.NotNil(t, err)
}
