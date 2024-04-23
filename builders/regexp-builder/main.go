package main

import (
	"fmt"
	. "github.com/hedhyw/rex/pkg/rex"
)

func main() {

	r := New(
		Helper.Email(),
	).MustCompile()
	fmt.Println(r)

	r = New(
		Chars.
			Range('A', 'Z').
			Repeat().
			OneOrMore(),
		Chars.
			Digits().
			Repeat().
			ZeroOrOne(),
	).MustCompile()
	fmt.Println(r)

	// We can define a set of characters and reuse the block.
	customCharacters := Common.Class(
		Chars.Range('a', 'z'), // `[a-z]`
		Chars.Upper(),         // `[A-Z]`
		Chars.Single('-'),     // `\x2D`
		Chars.Digits(),        // `[0-9]`
	) // `[a-zA-Z-0-9]`

	re := New(
		Chars.Begin(), // `^`
		customCharacters.Repeat().OneOrMore(),

		// Email delimeter.
		Chars.Single('@'), // `@`

		// Allow dot after delimter.
		Common.Class(
			Chars.Single('.'), // \.
			customCharacters,
		).
			Repeat().
			OneOrMore(),

		// Email should contain at least one dot.
		Chars.Single('.'), // `\.`
		Chars.
			Alphanumeric().
			Repeat().
			Between(2, 3),

		Chars.End(), // `$`
	).MustCompile()
	fmt.Println(re)
}
