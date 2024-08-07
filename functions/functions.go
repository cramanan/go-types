// The functions package provides helper function and function types
// matching some go-types callback functions.
package functions

import "golang.org/x/exp/constraints"

// compare returns
//
//	-1 if x is less than y,
//	 0 if x equals y,
//	+1 if x is greater than y.
//
// For floating-point types, a NaN is considered less than any non-NaN,
// a NaN is considered equal to a NaN, and -0.0 is equal to 0.0.
func compare[T constraints.Ordered](x, y T) int {
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
func isNaN[T constraints.Ordered](x T) bool { return x != x }

// '==' as a function.
func Equal[T comparable](x T, y T) bool { return x == y }

// '!=' as a function.
func NotEqual[T comparable](x T, y T) bool { return x != y }

// '>' as a function.
func Greater[T constraints.Ordered](x T, y T) bool { return compare(x, y) == 1 }

// '>=' as a function.
func GreaterOrEqual[T constraints.Ordered](x, y T) bool { return compare(x, y) >= 0 }

// '<' as a function.
func Less[T constraints.Ordered](x T, y T) bool { return compare(x, y) == -1 }

// '<=' as a function.
func LessOrEqual[T constraints.Ordered](x T, y T) bool { return compare(x, y) <= 0 }

// Satisfy returns a function that checks if a value is equal to the target value.
// The returned function can be used with functions like strings.FieldsFunc.
func Satisfy[T comparable](target T) func(T) bool { return func(x T) bool { return x == target } }

// Ascending compares two values of type T and returns an integer indicating their order.
func Ascending[T constraints.Ordered](x, y T) (order int) { return compare(x, y) }

// Descending compares two values of type T and returns an integer indicating their order in reverse.
func Descending[T constraints.Ordered](x, y T) (order int) { return -Ascending(x, y) }
