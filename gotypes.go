/*
# Types

# String

Documentation: https://pkg.go.dev/github.com/cramanan/go-types/strings

	package main

	import "github.com/cramanan/go-types/strings"

	func main() {
		// Initiate an empty string
		newString := strings.New() // <=> var newString strings.String = ""

		// From converts any string into strings.String
		fromString := strings.From("Hello World !")

		// From can also convert any string aliases, []rune and []byte
		fromString = strings.From([]byte{'F','o','o'})

		// Every standard strings functions are available as methods
		upper := fromString.ToUpper()
		lower := fromString.ToLower()

		// And even more
		last := fromString.At(-1) // "o"
	}

# Slice

Documentation : https://pkg.go.dev/github.com/cramanan/go-types/slices

	package main

	import (

		"github.com/cramanan/go-types/slices"

	)

	func main() {
		// Initiate an slice (replace type by wanted type)
		var newSlice slices.Slice[int] // <=> var newString Slice[type]
		newSlice = slices.New(1, 2, 3)

		// From converts any slice into slices.Slice
		fromSlice := slices.From([]rune("Hello"))

		// slices.Slice comes with special methods
		fromSlice.At(-1)
		fromSlice.Append(0, 1)
	}
*/
package gotypes
