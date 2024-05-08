package main

import (
	"context"
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

	b, _ := j.Marshal(context.TODO())
	fmt.Println(string(b))
}

type Component interface {
	Marshal(ctx context.Context) ([]byte, error)
}

type Object struct {
}

type Field struct {
}

type Array struct {
}

func F(name string, value interface{}) *Field {
	return nil
}

func O(fs ...*Field) *Object {
	return nil
}

func A(vs ...interface{}) *Array {
	return nil
}

func (o *Object) Marshal(ctx context.Context) ([]byte, error) {
	return nil, nil
}
