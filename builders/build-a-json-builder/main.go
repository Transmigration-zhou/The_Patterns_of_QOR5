package main

import (
	"fmt"
)

func main() {
	j := O(
		F("name", "John"),
		F("age", 1),
		F("cities", A("Hangzhou", "Shanghai", "Beijing")),
		F("isMarried", true),
		F("hobbies",
			O(
				F("reading", true),
				F("swimming", false),
			),
		),
		F("friends",
			A(
				O(
					F("name", "Alice"),
					F("age", 25),
				),
				O(
					F("name", "Bob"),
					F("age", 30),
				),
			),
		),
	)
	fmt.Println(j)
}

type Object struct {
}

type Field struct {
}

type Array struct {
}

type VT interface {
	string | int | int64 | bool | float64 | *Object | *Array
}

func F[V VT](name string, value V) *Field {
	return &Field{}
}

func O(fs ...*Field) *Object {
	return &Object{}
}

func A[V VT](vs ...V) *Array {
	return &Array{}
}
