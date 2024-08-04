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

This package only provides wrappers and do not handle any of these recurring problems:

-   panics : panic situations are always propagated when they occurs.
-   concurrency : if you use these types for concurrent jobs you will have to add your own handlers.
-   nil pointer dereferences : the provided functions & methods does not protect from nil pointer dereferences.

#### Your code, your rules

The Go-Types project was designed for any type of project. The use of generics and interfaces completly overthrow the [comparable](https://go.dev/blog/comparable) and Ordered interfaces.

If you wish to use this package for structs, native slices or maps that cannot be compared with [comparison operators](https://go.dev/ref/spec#Comparison_operators), you will have to use functions that use your own comparison rules:

-   Equality rule: Defines whether a value is equal to another.
-   Ordering rules: Defines which of 2 values is greater/lower than the other.

Methods expecting custom rules/functions have a name ending in "Func". e.g: (slice Slice[T]).ContainsFunc <!--add link that tracks line -->

## Import

To add this package to your project. Use the `go get` command:

```
go get github.com/cramanan/go-types
```

## Types

### String

The [strings](/strings/strings.go) package imports 2 functions to create the String type. `New()` and `From()`. <!--add links that track lines-->

The String type comes bundled with every functions from the standard strings library as methods. e.g:

```golang
foo := strings.New() // <=> var foo strings.String = ""

foo = strings.From("foo") // <=> foo = "foo"

fmt.Println(foo.ToUpper()) // returns "FOO"
```

Note that these methods returns shallow copies and do not modify the string in any way.

### Slice
