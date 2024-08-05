# Go-Types

[![Go Reference](https://pkg.go.dev/badge/github.com/cramanan/go-types.svg)](https://pkg.go.dev/github.com/cramanan/go-types)

A collection of "upgraded" Golang types.

## Overview

-   [Go-Types](#go-types)
-   [Overview](#overview)
-   [Description](#description)
    -   [Informations](#informations)
-   [Import](#import)
-   [Types](#types)
    -   [String](#string)
    -   [Slice](#slice)

## Description

Influenced by JavaScript & Rust, the Go-Types package offers Object-Oriented wrappers of [Golang types](https://go.dev/ref/spec#Types) with built-in methods that native types doesn't implement.
They also convert standard library functions into methods.

### Informations

#### Runtime errors

This package only provides wrappers and do not handle panic. Errors such as out of range, nil pointer dereference or deadlock error will still panic.

#### Your code, your rules

The Go-Types project was designed for any type of project. The use of generics and interfaces (any) for the slices.Unordered type completly overthrows the [comparable](https://go.dev/blog/comparable) and [Ordered](https://pkg.go.dev/cmp#Ordered) interfaces.

If you wish to use this package for structs, native slices or maps that cannot be compared with [comparison operators](https://go.dev/ref/spec#Comparison_operators), you will have to use functions that use your own comparison rules:

-   Equality rule: Defines whether a value is equal to another.
-   Ordering rules: Defines which of 2 values is greater/lower than the other.

Example:

```golang
// Restore Equality for comparable
func Equals[T comparable](left T,right T) bool { return left == right }

// Does Nothing. Use case: Mapping function
func Nothing[T comparable](from T) T { return from }
```

These type of function can be used in a lot of functions/methods ending with "Func".

Methods expecting custom rules/functions have a name ending in "Func". e.g: (slice Slice[T]).ContainsFunc <!--add link that tracks line -->

## Import

To add this package to your project. You must have Golang 1.19 or above and
use the `go get` command:

```
go get github.com/cramanan/go-types
```

## Types

Most packages imports 2 function that returns their respective types: `New()` and `From()`

Every functions return shallow copies and never modify the original value in any way. <sub>(if so, please notify me)</sub>

### String

The String type comes bundled with every functions from the standard strings library as methods. e.g:

```golang
foo := strings.New() <=> var foo strings.String = ""

foo = strings.From("foo") <=> foo = "foo"

fmt.Println(foo.ToUpper()) // returns "FOO"
```

### Slice

The Slice type implements common Array manipulation functions and methods like `Filter`, `Map`, `Reduce` and even more. e.g:

```golang
bar := slices.New(1, 2, 3) <=> var bar slices.Slice[int]{1, 2, 3}

baz := slices.From([]byte("baz")) <=> var baz = Slice[byte]("baz")

// Methods
baz = baz.Append('f','o', 'o') // bar is now Slice[byte]("foobaz")

//functions
slices.Map(bar, func(i int) int { return i *2 }) // return Slice[int]{2,4,6}
```

Author: [C. Ramananjaona](https://github.com/cramanan)
