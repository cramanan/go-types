package gotypes

/*
# Types

## String

	package main

	import "github.com/cramanan/go-types/strings"

	func main() {
		// Initiate an empty string
		newString := strings.New() // <=> var newString strings.String = ""

		// From converts any string into strings.String
		fromString := strings.From("Hello World !")

		// From can also convert any string aliases, []rune and []byte
		fromString = strings.From([]byte{'F','o','o'})
	}
*/