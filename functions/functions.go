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
