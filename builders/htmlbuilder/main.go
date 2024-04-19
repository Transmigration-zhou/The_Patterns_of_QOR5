package main

import (
	"context"
	. "github.com/theplant/htmlgo"
	"os"
)

func main() {
	var html = HTML(
		Head(
			Title("XML encoding with Go"),
		),
		Body(
			H1("XML encoding with Go"),
			P().Text("this format can be used as an alternative markup to XML"),
			A().
				Href("http://golang.org").
				Text("Go"),
			P(
				Text("this is some"),
				B("mixed"),
				Text("text. For more see the"),
				A().
					Href("http://golang.org").
					Text("Go"),
				Text("project"),
			),
			P().Text("some text"),
		),
	)

	_ = Fprint(os.Stdout, html, context.TODO())

}
