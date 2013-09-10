package gounit

import (
	"reflect"
	"testing"
)

func Equals(a, b interface{}) bool {
	typeA := reflect.TypeOf(a)
	typeB := reflect.TypeOf(b)

	if typeA != typeB {
		return false
	}

	return a == b
}

func AssertEquals(t *testing.T, expected, result interface{}) {
	if !Equals(expected, result) {
		t.Errorf("expected %v, received %v", expected, result)
	}
}
