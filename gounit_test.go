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
	struct1 := testStruct{
		1,
		"string",
		2.9,
	}

	struct2 := testStruct{
		1,
		"string",
		2.9,
	}

	if !Equals(struct1, struct2) {
		t.Fail()
	}
}

func TestEqualsStructRecursive(t *testing.T) {
	childStruct := testStruct{
		1,
		"string",
		2.9,
	}

	struct1 := testStructRecursive{
		childStruct,
		"string",
	}

	struct2 := testStructRecursive{
		childStruct,
		"string",
	}

	if !Equals(struct1, struct2) {
		t.Fail()
	}
}

func TestEqualsMap(t *testing.T) {
	map1 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	map2 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	if !Equals(map1, map2) {
		t.Fail()
	}
}

func TestEqualsArray(t *testing.T) {
	array1 := [...]int32{1, 2, 3}
	array2 := [...]int32{1, 2, 3}

	if !Equals(array1, array2) {
		t.Fail()
	}
}

func TestEqualsSlice(t *testing.T) {
	slice1 := []int32{1, 2, 3}
	slice2 := []int32{1, 2, 3}

	if !Equals(slice1, slice2) {
		t.Fail()
	}
}
