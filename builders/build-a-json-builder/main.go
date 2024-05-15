package main

import (
	"bytes"
	"context"
	"encoding/json"
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
	fields []*Field
}

type Field struct {
	k string
	v interface{}
}

type Array struct {
	values []interface{}
}

func F(name string, value interface{}) *Field {
	return &Field{name, value}
}

func O(fs ...*Field) *Object {
	return &Object{fs}
}

func A(vs ...interface{}) *Array {
	return &Array{vs}
}

func (o *Object) Marshal(ctx context.Context) ([]byte, error) {
	buf := bytes.NewBuffer([]byte("{"))
	for i, field := range o.fields {
		if i > 0 {
			buf.WriteByte(',')
		}
		b, err := field.Marshal(ctx)
		if err != nil {
			return nil, err
		}
		buf.Write(b)
	}
	buf.WriteByte('}')
	return buf.Bytes(), nil
}

func (f *Field) Marshal(ctx context.Context) ([]byte, error) {
	buf := bytes.NewBuffer([]byte("\""))
	buf.WriteString(f.k)
	buf.WriteByte('"')
	buf.WriteByte(':')
	switch v := f.v.(type) {
	case Component:
		b, err := v.Marshal(ctx)
		if err != nil {
			return nil, err
		}
		buf.Write(b)
	default:
		b, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		buf.Write(b)
	}
	return buf.Bytes(), nil
}

func (a *Array) Marshal(ctx context.Context) ([]byte, error) {
	buf := bytes.NewBuffer([]byte("["))
	for i, value := range a.values {
		if i > 0 {
			buf.WriteByte(',')
		}
		switch v := value.(type) {
		case Component:
			b, err := v.Marshal(ctx)
			if err != nil {
				return nil, err
			}
			buf.Write(b)
		default:
			b, err := json.Marshal(v)
			if err != nil {
				return nil, err
			}
			buf.Write(b)
		}
	}
	buf.WriteByte(']')
	return buf.Bytes(), nil
}
