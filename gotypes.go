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
*/
package gotypes
