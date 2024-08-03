# Go-Types

[![Go Reference](https://pkg.go.dev/badge/github.com/cramanan/go-types.svg)](https://pkg.go.dev/github.com/cramanan/go-types)

A collection of "upgraded" Golang types.

## Overview

-   [Go-Types](#go-types)
-   [Overview](#overview)
-   [Description](#description)
-   [Import](#import)
-   [Types](#types)
    -   [String](#string)
    -   [Slice](#slice)

## Description

The Go-Types package contains Object-Oriented types of [Golang types](https://go.dev/ref/spec#Types) with built-in methods that native types doesn't implement.
They also convert standard library functions into methods.

## Import

To add this package to your project. Use the `go get` command:

```
go get github.com/cramanan/go-types
```

## Types

### String

The [strings](/strings/strings.go) package imports 2 functions to create the String type. `New()` and `From()`.

The String type comes bundled with every functions from the standard strings library as methods. e.g: `fromString.ToUpper() => "FOO"`

Note that these methods returns shallow copies and do not modify the string in any way.

### Slice
