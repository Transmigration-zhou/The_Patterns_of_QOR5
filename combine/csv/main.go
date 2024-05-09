package main

import (
	"encoding/csv"
	"net/http"
	"strings"

	. "github.com/theplant/htmlgo"
)

func CSVSearchFuncWrapper(csvpath string, in SearchFunc) SearchFunc {
	r := csv.NewReader(strings.NewReader(csvContent))
	results, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	return func(keyword string) ([][]string, error) {
		return results, nil
	}
}

func main() {
	b := New().
		SearchFunc(
			CSVSearchFuncWrapper("test.csv", nil),
		)
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
			tr.AppendChildren(Td(RawHTML(col)).Style("border: 1px solid #eeeeee;"))
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

var csvContent = `
transaction_id,date,amount,category
1,2/16/2022,186,transportation
2,5/17/2022,61,transportation
3,5/24/2022,211,groceries
4,8/28/2022,172,entertainment
5,5/13/2022,122,dining
6,5/24/2022,268,bills
7,5/18/2022,154,bills
8,8/2/2022,219,entertainment
9,1/22/2022,37,bills
10,5/25/2022,72,dining
11,3/20/2022,58,bills
12,12/16/2022,149,dining
13,12/15/2022,14,entertainment
14,10/7/2022,240,bills
15,9/26/2022,99,bills
16,1/10/2022,62,entertainment
17,4/9/2022,127,bills
18,5/1/2022,241,dining
19,10/18/2022,271,bills
20,10/9/2022,232,dining
21,12/9/2022,298,transportation
22,6/8/2022,111,transportation
23,10/29/2022,10,transportation
24,2/1/2022,186,entertainment
25,10/9/2022,24,bills
26,11/13/2022,263,groceries
27,9/8/2022,281,groceries
28,2/4/2022,176,groceries
29,3/1/2022,188,dining
30,10/22/2022,242,dining
31,11/12/2022,177,bills
32,5/8/2022,171,transportation
33,4/2/2022,167,dining
34,3/22/2022,190,transportation
35,12/28/2022,199,entertainment
36,6/30/2022,49,transportation
37,4/14/2022,259,groceries
38,7/12/2022,111,dining
39,9/21/2022,128,dining
40,3/17/2022,113,entertainment
41,10/13/2022,211,transportation
42,12/11/2022,150,entertainment
43,9/17/2022,148,bills
44,11/25/2022,238,bills
45,5/16/2022,171,groceries
46,8/9/2022,81,entertainment
47,11/14/2022,171,bills
48,1/23/2022,62,entertainment
49,9/4/2022,276,groceries
50,12/9/2022,219,bills
51,9/11/2022,214,entertainment
52,11/26/2022,135,groceries
53,12/18/2022,143,dining
54,12/9/2022,291,dining
55,10/24/2022,212,groceries
56,8/25/2022,202,bills
57,5/23/2022,109,transportation
58,2/13/2022,248,bills
59,2/7/2022,151,dining
60,8/28/2022,288,bills
61,1/22/2022,149,entertainment
62,6/30/2022,227,bills
63,3/30/2022,245,entertainment
64,8/22/2022,250,transportation
65,8/16/2022,213,bills
66,5/1/2022,171,dining
67,5/10/2022,90,bills
68,6/5/2022,194,entertainment
69,7/10/2022,238,bills
70,10/18/2022,26,transportation
71,2/1/2022,193,bills
72,9/24/2022,69,transportation
73,6/3/2022,247,groceries
74,5/7/2022,111,bills
75,5/12/2022,140,dining
76,8/11/2022,41,bills
77,11/27/2022,118,dining
78,3/6/2022,178,entertainment
79,10/15/2022,242,groceries
80,12/1/2022,182,bills
81,3/8/2022,236,bills
82,6/26/2022,169,dining
83,2/8/2022,183,entertainment
84,11/24/2022,81,groceries
85,4/3/2022,22,entertainment
86,3/12/2022,7,bills
87,10/10/2022,233,dining
88,1/31/2022,73,entertainment
89,7/29/2022,37,groceries
90,4/4/2022,214,entertainment
91,8/5/2022,29,transportation
92,4/26/2022,86,dining
93,7/5/2022,32,bills
94,6/5/2022,195,bills
95,2/21/2022,220,entertainment
96,7/27/2022,178,transportation
97,9/18/2022,150,transportation
98,2/12/2022,148,dining
99,11/22/2022,96,transportation
100,8/21/2022,243,bills
`
