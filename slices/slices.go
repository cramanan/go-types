package slices

type Slice[T any] []T

/*
New returns a new slices.Slice of type T, optionally populated with the provided values.

If no values are provided, an empty Slice is returned.

Example:

	s := slices.New(1, 2, 3) // returns a Slice[int] containing [1, 2, 3]

	s := slices.New[string]() // returns an empty Slice[string]
*/
func New[T any](values ...T) Slice[T] {
	return append(Slice[T]{}, values...)
}

/*
From converts a standard Go slice to a Slice of the same type.

This function does not create a copy of the original slice, but rather wraps it in a Slice.
Modifications made to the returned Slice will affect the original slice.

Example:

	s := []int{1, 2, 3}

	wrapped := From(s) // returns a Slice[int] containing [1, 2, 3]
*/
func From[T any](slice []T) Slice[T] {
	return slice
}

/*
Prepend adds the given values to the beginning of the Slice.

This method returns a new Slice containing the prepended values followed by the original elements.

The original Slice is not modified.

Example:

	s := Slice[int]{1, 2, 3}

	s = s.Prepend(0) // returns a Slice[int] containing [0, 1, 2, 3]
*/
func (slice Slice[T]) Prepend(values ...T) Slice[T] {
	return append(values, slice...)
}

/*
Append adds the given values to the end of the Slice.

This method returns a new Slice containing the original elements followed by the appended values.
The original Slice is not modified.

Example:

	s := Slice[int]{1, 2, 3}

	s = s.Append(4, 5) // returns a Slice[int] containing [1, 2, 3, 4, 5]
*/
func (slice Slice[T]) Append(values ...T) Slice[T] {
	return append(slice, values...)
}

/*
At returns the element at the specified index in the Slice.

The index can be either positive (from the start of the Slice) or negative (from the end of the Slice).

If the index is negative, it is treated as an offset from the end of the Slice.

Example:

	s := Slice[int]{1, 2, 3, 4, 5}

	s.At(0)  // returns 1
	s.At(-1) // returns 5
	s.At(2)  // returns 3

Panics if the index is out of bounds.

Note: This method does not perform bounds checking for performance reasons.
*/
func (slice Slice[T]) At(n int) T {
	if n < 0 {
		n = len(slice) + n
	}
	return slice[n]
}

/*
Length returns the length of the Slice.

Example:

	s := Slice[int]{1, 2, 3, 4, 5}
	s.Len()  // returns 5
*/
func (slice Slice[T]) Len() int {
	return len(slice)
}

/*
ForEach iterates over the elements of the Slice and calls the provided callback function for each element.

The callback function is called with two arguments: the current element's value and its index in the Slice.

Example:

	s := Slice[int]{1, 2, 3}

	s.ForEach(func(value int, index int) {
	    fmt.Printf("Element at index %d: %d\n", index, value)
	})

This would output:

	Element at index 0: 1
	Element at index 1: 2
	Element at index 2: 3

Note: The callback function is called for each element in the Slice, in order.

If the callback function panics, the iteration is stopped and the panic is propagated.
*/
func (slice Slice[T]) ForEach(callbackFn func(value T, index int)) {
	for idx, v := range slice {
		callbackFn(v, idx)
	}
}

/*
Filter returns a new Slice containing only the elements for which the provided callback function returns true.

The callback function is called with two arguments: the current element's value and its index in the Slice.

It should return true if the element should be included in the filtered Slice, and false otherwise.

Example:

	s := Slice[int]{1, 2, 3, 4, 5}

	evenNumbers := s.Filter(func(element int, index int) bool {
	    return element % 2 == 0
	})

	// evenNumbers is now Slice[int]{2, 4}

Note: The callback function is called for each element in the Slice, in order.
The filtered Slice is created by appending elements to a new Slice, so the order of elements is preserved.
*/
func (slice Slice[T]) Filter(callbackFn func(element T, index int) bool) (filtered Slice[T]) {
	filtered = make(Slice[T], 0, len(slice))
	for idx, value := range slice {
		if callbackFn(value, idx) {
			filtered = append(filtered, value)
		}
	}
	return filtered
}

/*
Some returns true if at least one element in the Slice satisfies the provided callback function.

The callback function is called with two arguments: the current element's value and its index in the Slice.

It should return true if the element satisfies the condition, and false otherwise.

Example:

	s := Slice[int]{1, 2, 3, 4, 5}

	hasEvenNumber := s.Some(func(element int, index int) bool {
	    return element % 2 == 0
	})
	// hasEvenNumber is now true

Note: The iteration stops as soon as the callback function returns true for any element.
*/
func (slice Slice[T]) Some(callbackFn func(element T, index int) bool) bool {
	for idx, value := range slice {
		if callbackFn(value, idx) {
			return true
		}
	}
	return false
}

/*
Every returns true if all elements in the Slice satisfy the provided callback function.

The callback function is called with two arguments: the current element's value and its index in the Slice.

It should return true if the element satisfies the condition, and false otherwise.

Example:

	s := Slice[int]{2, 4, 6, 8, 10}

	allEvenNumbers := s.Every(func(element int, index int) bool {
	    return element % 2 == 0
	})
	// allEvenNumbers is now true

Note: The iteration stops as soon as the callback function returns false for any element.
*/
func (slice Slice[T]) Every(callbackFn func(element T, index int) bool) bool {
	for idx, value := range slice {
		if !callbackFn(value, idx) {
			return false
		}
	}
	return true
}

/*
Reverse returns a new Slice with the elements in reverse order.

Example:

	s := Slice[int]{1, 2, 3, 4, 5}

	reversed := s.Reverse()
	// reversed is now Slice[int]{5, 4, 3, 2, 1}

Note: The original Slice is not modified. A new Slice is created with the elements in reverse order.
*/
func (slice Slice[T]) Reverse() Slice[T] {
	reversed := make(Slice[T], len(slice))
	for i := len(slice) - 1; i >= 0; i-- {
		reversed[len(slice)-1-i] = slice[i]
	}
	return reversed

}

/*
IndexFunc returns the index of the first element in the Slice that satisfies the provided callback function,

or -1 if no such element is found.

The callback function is called with one argument: the current element's value.

It should return true if the element satisfies the condition, and false otherwise.

Example:

	s := Slice[int]{1, 2, 3, 4, 5}

	index := s.IndexFunc(func(element int) bool {
	    return element == 3
	})
	// index is now 2

Note: The iteration stops as soon as the callback function returns true for any element.
*/
func (slice Slice[T]) IndexFunc(f func(T) bool) int {
	for i := range slice {
		if f(slice[i]) {
			return i
		}
	}
	return -1
}

/*
ContainsFunc returns true if the Slice contains at least one element that satisfies the provided callback function,

and false otherwise.

The callback function is called with one argument: the current element's value.

It should return true if the element satisfies the condition, and false otherwise.

Example:

	s := Slice[int]{1, 2, 3, 4, 5}

	contains := s.ContainsFunc(func(element int) bool {
	    return element == 3
	})
	// contains is now true

Note: This method is a convenience wrapper around IndexFunc, and has the same iteration behavior.
*/
func (slice Slice[T]) ContainsFunc(f func(T) bool) bool {
	return slice.IndexFunc(f) >= 0
}

/*
EqualFunc compares two slices for equality using a custom comparison function.

Example:

	s1 := Slice[int]{1,2,3,4,5}
	s2 := Slice[int]{1,2,3,4,5}

	eq := func(e1, e2 int) bool{ return e1 == e2 }
	equals := s1.EqualFunc(s2, eq)
	//equals is now true
*/
func (s1 Slice[T]) EqualFunc(s2 Slice[T], eq func(T, T) bool) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, v1 := range s1 {
		if !eq(v1, s2[i]) {
			return false
		}
	}
	return true
}

/*
Map applies a transformation function to each element of the input Slice[From] and returns a new Slice[To] with the results.

The transformation function is called with one argument: the current element's value.

It should return the transformed value of type To.

Example:

	s := Slice[int]{1, 2, 3, 4, 5}

	double := func(x int) int { return x * 2 }

	doubled := Map(s, double)

	// doubled is now Slice[int]{2, 4, 6, 8, 10}

Note: The order of elements in the output Slice is the same as in the input Slice.
*/
func Map[From, To any](s Slice[From], f func(From) To) (mapped Slice[To]) {
	for _, v := range s {
		mapped = append(mapped, f(v))
	}
	return mapped
}

/*
Reduce applies a reduction function to each element of the input Slice[From] and returns a single value of type To.

The reduction function is called with two arguments: the accumulator (initially set to initialValue) and the current element's value.

It should return the new accumulator value.

The reduction process starts with the initialValue and iterates over the input Slice, applying the reduction function to each element.

The final accumulator value is returned as the result.

Example:

	s := Slice[int]{1, 2, 3, 4, 5}

	sum := func(acc int, current int) int { return acc + current }

	result := Reduce(s, sum, 0)
	// result is now 15 (the sum of all elements in the Slice)

Note: If the input Slice is empty, the initialValue is returned as the result.
*/
func Reduce[From, To any](s Slice[From], callbackFn func(To, From) To, initialValue To) (reduced To) {
	reduced = initialValue
	for _, element := range s {
		reduced = callbackFn(reduced, element)
	}
	return reduced
}
