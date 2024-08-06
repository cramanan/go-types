package functions

// ComparisonFunc is a function type that takes two values of type T and returns an
//
// integer indicating their relative order. The return value is:
//   - positive if the first value is greater than the second
//   - negative if the first value is less than the second
//   - zero if the values are equal
type ComparisonFunc[T any] func(T, T) int

// CallbackFunc is a function type that takes a value of type T and an integer index
// as arguments. It is often used as a callback function to process each element in
// a collection, such as an array or slice, where the integer index represents the
// position of the element being processed.
type CallbackFunc[T any] func(T, int)

// SatisfyFunc returns a predicate function for a given value of type T.
type SatisfyFunc[T comparable] func(T) func(T) bool

type CallbackFuncReturns[O, I any] func(I, int) O
