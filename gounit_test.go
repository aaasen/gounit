package gounit

import (
	"testing"
)

func TestEqualsString(t *testing.T) {
	if !Equals("hello", "hello") {
		t.Fail()
	}
}

func TestEqualsNumber(t *testing.T) {
	if !Equals(1, 1) {
		t.Fail()
	}
}

func TestEqualsNumberDiffType(t *testing.T) {
	if !Equals(1, 1.0) {
		t.Fail()
	}
}

func TestEqualsFloatThreshold(t *testing.T) {
	if !Equals(1, 0.999999) {
		t.Fail()
	}
}

type testStruct struct {
	a int
	b string
	c float64
}

type testStructRecursive struct {
	a testStruct
	b string
}

func TestEqualsStruct(t *testing.T) {
	structA := testStruct{
		1,
		"string",
		2.9,
	}

	structB := testStruct{
		1,
		"string",
		2.9,
	}

	if !Equals(structA, structB) {
		t.Fail()
	}
}

func TestEqualsStructRecursive(t *testing.T) {
	childStruct := testStruct{
		1,
		"string",
		2.9,
	}

	structA := testStructRecursive{
		childStruct,
		"string",
	}

	structB := testStructRecursive{
		childStruct,
		"string",
	}

	if !Equals(structA, structB) {
		t.Fail()
	}
}
