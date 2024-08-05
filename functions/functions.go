// The functions package provides helper function and function types
// matching some go-types callback functions.
package functions

import (
	"cmp"
)

// '==' as a function.
func Equal[T comparable](x T, y T) bool { return x == y }

// '!=' as a function.
func NotEqual[T comparable](x T, y T) bool { return x != y }

// '>' as a function.
func Greater[T cmp.Ordered](x T, y T) bool { return cmp.Compare(x, y) == 1 }

// '>=' as a function.
func GreaterOrEqual[T cmp.Ordered](x, y T) bool { return cmp.Compare(x, y) >= 0 }

// '<' as a function.
func Less[T cmp.Ordered](x T, y T) bool { return cmp.Compare(x, y) == -1 }

// '<=' as a function.
func LessOrEqual[T cmp.Ordered](x T, y T) bool { return cmp.Compare(x, y) <= 0 }

// Satisfy returns a function that checks if a value is equal to the target value.
// The returned function can be used with functions like strings.IndexFunc to find
// the index of the target value in a slice.
func Satisfy[T comparable](target T) func(T) bool {
	return func(x T) bool {
		return x == target
	}
}

// Ascending compares two values of type T and returns an integer indicating their order.
func Ascending[T cmp.Ordered](x, y T) (order int) {
	return cmp.Compare(x, y)
}

// Descending compares two values of type T and returns an integer indicating their order in reverse.
func Descending[T cmp.Ordered](x, y T) (order int) {
	order = Ascending(x, y)
	return -order
}
