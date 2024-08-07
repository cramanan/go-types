# Go-Types

[![Go Reference](https://pkg.go.dev/badge/github.com/cramanan/go-types.svg)](https://pkg.go.dev/github.com/cramanan/go-types)

A collection of advanced Golang types and generic wrappers.

## Overview

-   [Go-Types](#go-types)
-   [Overview](#overview)
-   [Description](#description)
-   [Import](#import)
-   [Packages](#packages)
    -   [Boolean](#boolean)
    -   [String](#string)
    -   [Slice](#slice)
    -   [Functions](#functions)
    -   [Constants](#constants)
-   [Informations](#informations)

## Description

Influenced by JavaScript and Python, the Go-Types package offers Object-Oriented wrappers of [Golang types](https://go.dev/ref/spec#Types) with built-in methods that native types doesn't implement.

They also convert standard library functions into generics methods.

## Import

To add this package to your project, use the `go get` command:

```
go get github.com/cramanan/go-types
```

## Packages

### Boolean

The Boolean type is a wrapper that adds logical operators as methods for booleans. e.g:

The package provides comparison functions for all boolean types.

It also implement a Non-Zero / truthy function.

```golang
booleans.False.NOT() // return true

booleans.NOT(isTrue) // returns false

booleans.IsTruthy("") // returns false
```

### String

The String type comes bundled with every functions from the standard strings library as methods. e.g:

```golang
import "github.com/cramanan/go-types/strings"

foo := strings.New() // returns "" as strings.String

foo = strings.From("foo") //returns "" foo as strings.String

fmt.Println(foo.ToUpper()) // returns "FOO"
```

Since `gotypes/strings` overwrites the standard `strings` library, the Go-Types package imports, elevates with generics then exports every `strings` function.

```golang
type IString interface { ~string | ~[]byte | ~[]rune }
```

Every exported functions use the IString interface as a parameter. This allows standard `strings` manipulation function to work with all the above type.

### Slice

The Slice type implements common Array manipulation functions and methods like `Filter`, `Map`, `Reduce` and more. e.g:

```golang
bar := slices.New(1, 2, 3) // returns slices.Slice[int]{1, 2, 3}

baz := slices.From([]byte("baz")) // returns Slice[byte]{ 'b', 'a', 'z'}

// Methods
baz = baz.Prepend('f', 'o', 'o') // bar is now Slice[byte]("foobaz")

// Functions
slices.Map(bar, func(i int) int { return i *2 }) // return Slice[int]{2, 4, 6}
```

The `gotypes/slices` also overwrites the standard [`slices`](https://pkg.go.dev/slices) library or [`golang.org/x/exp/slices`](https://pkg.go.dev/slices) (depending on your version).

```
type Ordered Slice[constraints.Ordered]
```

For simple data types that can be compared, it is better to use the Ordened type. The Slice type is the one used for unordered types. [See more](#your-code-your-rules)

### Functions

The functions package provides some [callback functions](/functions/functions.go) for Funcs arguments. It also provides [types](/functions/types.go) from these functions to define arguments, types, methods...

### Constants

At that time the `constants` package doesn't include any useful nor significant value. Feel free to propose any.

### Informations

#### Runtime errors

This package only provides wrappers and do not handle panics.
Errors such as out of range, nil pointer dereference or deadlock errors will still panic.

#### Your code, your rules

For the slices package.

If you wish to use this package and compare custom or complex data types that cannot be compared with [comparison operators](https://go.dev/ref/spec#Comparison_operators), you will have to use the Slice type with functions that use your own comparison functions:

-   Ordering rules: Defines which of 2 values is greater/lower than the other.
-   Equality rule: Defines whether a value is equal to another.

Methods expecting custom rules/functions often have a name ending in "Func". e.g: (slice Slice[T]).ContainsFunc

Author: [C. Ramananjaona](https://github.com/cramanan)
