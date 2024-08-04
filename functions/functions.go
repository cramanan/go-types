// The functions package provides helper function and function types
// matching some go-types callback functions.
package functions

// The OrderedFunc restores the Ordered operators for functions ending in Funcs
type OrderedFunc[T Ordered] func(T, T) bool

// '==' as a function.
func Equal[T comparable](a T, b T) bool { return a == b }

// '!=' as a function.
func NotEqual[T comparable](a T, b T) bool { return a != b }

// '>' as a function.
func Greater[T Ordered](a T, b T) bool { return a > b }

// '>=' as a function.
func GreaterOrEqual[T Ordered](a T, b T) bool { return a >= b }

// '<' as a function.
func Less[T Ordered](a T, b T) bool { return a < b }

// '<=' as a function.
func LessOrEqual[T Ordered](a T, b T) bool { return a <= b }

// The CallbackFn is a function used in slice iterations.
// The generic value T match the value and the integer match the index
type CallbackFn[T any] func(T, int)

type CompareFn[T any] func(T, T) int

func Ascending[T Ordered](t1, t2 T) int {
	if t1 > t2 {
		return 1
	}

	if t2 > t1 {
		return -1
	}

	return 0

}

func Descending[T Ordered](t1, t2 T) int {
	if t1 < t2 {
		return 1
	}

	if t2 < t1 {
		return -1
	}

	return 0
}
