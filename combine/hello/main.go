package main

import (
	"net/http"

	. "github.com/theplant/htmlgo"
)

func main() {
	b := New()
	http.Handle("/", b)
	http.ListenAndServe(":8080", nil)
}

type SearchFunc func(keyword string) ([][]string, error)

type Builder struct {
	Searcher SearchFunc
}

func New() *Builder {
	return &Builder{
		Searcher: func(keyword string) ([][]string, error) {
			return [][]string{
				{"hello", "world"},
			}, nil
		},
	}
}

func (b *Builder) SearchFunc(v SearchFunc) *Builder {
	b.Searcher = v
	return b
}

func (b *Builder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	keyword := r.FormValue("keyword")
	rows, err := b.Searcher(keyword)
	if err != nil {
		panic(err)
	}

	t := Table().Style("width: 100%; border-collapse: collapse; font-family: monospace;")
	for _, row := range rows {
		tr := Tr()
		for _, col := range row {
			tr.AppendChildren(Td().Text(col).Style("border: 1px solid #eeeeee;"))
		}
		t.AppendChildren(tr)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = Fprint(w, HTML(Body(
		Form(
			Input("keyword").Type("text").Placeholder("Search...").Value(keyword),
			Button("Search").Type("submit"),
		).Action("/").Method("GET"),
		t,
	)), r.Context())
	if err != nil {
		panic(err)
	}
}
