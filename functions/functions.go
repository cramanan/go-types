// The functions package provides helper function and function types
// matching some go-types callback functions.
package functions

// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package cmp provides types and functions related to comparing
// ordered values.
// package cmp

// Ordered is a constraint that permits any ordered type: any type
// that supports the operators < <= >= >.
// If future releases of Go add new ordered types,
// this constraint will be modified to include them.
//
// Note that floating-point types may contain NaN ("not-a-number") values.
// An operator such as == or < will always report false when
// comparing a NaN value with any other value, NaN or not.
// See the [Compare] function for a consistent way to compare NaN values.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// // Less reports whether x is less than y.
// // For floating-point types, a NaN is considered less than any non-NaN,
// // and -0.0 is not less than (is equal to) 0.0.
// func Less[T Ordered](x, y T) bool {
// 	return (isNaN(x) && !isNaN(y)) || x < y
// }

// Compare returns
//
//	-1 if x is less than y,
//	 0 if x equals y,
//	+1 if x is greater than y.
//
// For floating-point types, a NaN is considered less than any non-NaN,
// a NaN is considered equal to a NaN, and -0.0 is equal to 0.0.
func Compare[T Ordered](x, y T) int {
	xNaN := isNaN(x)
	yNaN := isNaN(y)
	if xNaN && yNaN {
		return 0
	}
	if xNaN || x < y {
		return -1
	}
	if yNaN || x > y {
		return +1
	}
	return 0
}

// isNaN reports whether x is a NaN without requiring the math package.
// This will always return false if T is not floating-point.
func isNaN[T Ordered](x T) bool {
	return x != x
}

// '==' as a function.
func Equal[T comparable](x T, y T) bool { return x == y }

// '!=' as a function.
func NotEqual[T comparable](x T, y T) bool { return x != y }

// '>' as a function.
func Greater[T Ordered](x T, y T) bool { return Compare(x, y) == 1 }

// '>=' as a function.
func GreaterOrEqual[T Ordered](x, y T) bool { return Compare(x, y) >= 0 }

// '<' as a function.
func Less[T Ordered](x T, y T) bool { return Compare(x, y) == -1 }

// '<=' as a function.
func LessOrEqual[T Ordered](x T, y T) bool { return Compare(x, y) <= 0 }

// Satisfy returns a function that checks if a value is equal to the target value.
// The returned function can be used with functions like strings.IndexFunc to find
// the index of the target value in a slice.
func Satisfy[T comparable](target T) func(T) bool {
	return func(x T) bool {
		return x == target
	}
}

// Ascending compares two values of type T and returns an integer indicating their order.
func Ascending[T Ordered](x, y T) (order int) {
	return Compare(x, y)
}

// Descending compares two values of type T and returns an integer indicating their order in reverse.
func Descending[T Ordered](x, y T) (order int) {
	order = Ascending(x, y)
	return -order
}
