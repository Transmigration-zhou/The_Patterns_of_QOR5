package main

import (
	"fmt"
	"github.com/huandu/go-sqlbuilder"
)

func main() {
	sb := sqlbuilder.NewSelectBuilder()

	sb.Select("id").From("user").
		Where(
			sb.In("status", 1, 2, 5),
			sb.Or(
				sb.Equal("name", "foo"),
				sb.Like("email", "foo@%"),
			),
		)

	sql, args := sb.Build()
	fmt.Println(sql)
	fmt.Println(args)

	// Output:
	// SELECT id FROM user WHERE status IN (?, ?, ?) AND (name = ? OR email LIKE ?)
	// [1 2 5 foo foo@%]
}
