# Go-Types

[![Go Reference](https://pkg.go.dev/badge/github.com/cramanan/go-types.svg)](https://pkg.go.dev/github.com/cramanan/go-types)

A collection of "upgraded" Golang types.

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
-   [Informations](#informations)

## Description

Influenced by JavaScript, Python and Rust, the Go-Types package offers Object-Oriented wrappers of [Golang types](https://go.dev/ref/spec#Types) with built-in methods that native types doesn't implement.

They also convert standard library functions into generics methods.

## Import

To add this package to your project, use the `go get` command:

```
go get github.com/cramanan/go-types
```

## Packages

Most packages imports 2 function that returns their respective types: `New()` and `From()`

Every methods return shallow copies and never modify the original value in any way. <sub>(if so, please report issues.)</sub>

### Boolean

The Boolean type is a wrapper that adds Logical operators as methods for booleans. e.g:

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

### Functions

The functions package provides some [callback functions](/functions/functions.go) for Funcs arguments. It also provides [types](/functions/types.go) from these functions to define arguments, types, methods...

### Informations

#### Runtime errors

This package only provides wrappers and do not handle panic. Errors such as out of range, nil pointer dereference or deadlock errors will still panic.

#### Your code, your rules

For the slices package.

The Go-Types project was designed for any type of project. The use of generics and interfaces (any) type completly overthrows the [comparable](https://go.dev/blog/comparable) and [Ordered](https://pkg.go.dev/cmp#Ordered) interfaces.

If you wish to use this package to compare custom or complex data types that cannot be compared with [comparison operators](https://go.dev/ref/spec#Comparison_operators), you will have to use the ordened sub-package with functions that use your own comparison rules:

-   Ordering rules: Defines which of 2 values is greater/lower than the other.
-   Equality rule (optionnal): Defines whether a value is equal to another.

Methods expecting custom rules/functions often have a name ending in "Func". e.g: (slice Slice[T]).ContainsFunc <!--add link that tracks line -->

Author: [C. Ramananjaona](https://github.com/cramanan)
