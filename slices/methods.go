package slices

// Prepend adds the given values to the beginning of the Slice.
//
// This method returns a new Slice containing the prepended values followed by the original elements.
//
// The original Slice is not modified.
//
// Example:
//
//	s := Slice[int]{1, 2, 3}
//
// s = s.Prepend(0) // returns a Slice[int] containing [0, 1, 2, 3]
func (slice Slice[T]) Prepend(values ...T) Slice[T] {
	return append(values, slice...)
}

// Append adds the given values to the end of the Slice.
//
// This method returns a new Slice containing the original elements followed by the appended values.
// The original Slice is not modified.
//
// Example:
//
//	s := Slice[int]{1, 2, 3}
//
// s = s.Append(4, 5) // returns a Slice[int] containing [1, 2, 3, 4, 5]
func (slice Slice[T]) Append(values ...T) Slice[T] {
	return append(slice, values...)
}

// At returns the element at the specified index in the Slice.
//
// The index can be either positive (from the start of the Slice) or negative (from the end of the Slice).
//
// If the index is negative, it is treated as an offset from the end of the Slice.
//
// Example:
//
//	s := Slice[int]{1, 2, 3, 4, 5}
//
//	s.At(0)  // returns 1
//	s.At(-1) // returns 5
//	s.At(2)  // returns 3
//
// Note: This method does not perform bounds checking for performance reasons.
func (slice Slice[T]) At(n int) T {
	if n < 0 {
		n = len(slice) + n
	}
	return slice[n]
}

// Len returns the length of the Slice.
//
// Example:
//
// s := Slice[int]{1, 2, 3, 4, 5}
// s.Len()  // returns 5
func (slice Slice[T]) Len() int {
	return len(slice)
}

// ForEach iterates over the elements of the Slice and calls the provided callback function for each element.
//
// The callback function is called with two arguments: the current element's value and its index in the Slice.
//
// Example:
//
//	s := Slice[int]{1, 2, 3}
//
//	s.ForEach(func(value int, index int) {
//	    fmt.Printf("Element at index %d: %d\n", index, value)
//	})
//
// This would output:
//
//	Element at index 0: 1
//	Element at index 1: 2
//	Element at index 2: 3
//
// Note: The callback function is called for each element in the Slice, in order.
//
// If the callback function panics, the iteration is stopped and the panic is propagated.
func (slice Slice[T]) ForEach(callbackFn func(value T, index int)) {
	for idx, v := range slice {
		callbackFn(v, idx)
	}
}

// Filter returns a new Slice containing only the elements for which the provided callback function returns true.
//
// The callback function is called with two arguments: the current element's value and its index in the Slice.
//
// It should return true if the element should be included in the filtered Slice, and false otherwise.
//
// Example:
//
//	s := Slice[int]{1, 2, 3, 4, 5}
//
//	evenNumbers := s.Filter(func(element int, index int) bool {
//	    return element % 2 == 0
//	})
//
//	// evenNumbers is now Slice[int]{2, 4}
//
// Note: The callback function is called for each element in the Slice, in order.
// The filtered Slice is created by appending elements to a new Slice, so the order of elements is preserved.
func (slice Slice[T]) Filter(callbackFn func(element T, index int) bool) (filtered Slice[T]) {
	filtered = make(Slice[T], 0, len(slice))
	for idx, value := range slice {
		if callbackFn(value, idx) {
			filtered = append(filtered, value)
		}
	}
	return filtered
}

// Some returns true if at least one element in the Slice satisfies the provided callback function.
//
// The callback function is called with two arguments: the current element's value and its index in the Slice.
//
// It should return true if the element satisfies the condition, and false otherwise.
//
// Example:
//
//	s := Slice[int]{1, 2, 3, 4, 5}
//
//	hasEvenNumber := s.Some(func(element int, index int) bool {
//	    return element % 2 == 0
//	})
//	// hasEvenNumber is now true
//
// Note: The iteration stops as soon as the callback function returns true for any element.
func (slice Slice[T]) Some(callbackFn func(element T, index int) bool) bool {
	for idx, value := range slice {
		if callbackFn(value, idx) {
			return true
		}
	}
	return false
}

// Every returns true if all elements in the Slice satisfy the provided callback function.
//
// The callback function is called with two arguments: the current element's value and its index in the Slice.
//
// It should return true if the element satisfies the condition, and false otherwise.
//
// Example:
//
//	s := Slice[int]{2, 4, 6, 8, 10}
//
//	allEvenNumbers := s.Every(func(element int, index int) bool {
//	    return element % 2 == 0
//	})
//	// allEvenNumbers is now true
//
// Note: The iteration stops as soon as the callback function returns false for any element.
func (slice Slice[T]) Every(callbackFn func(element T, index int) bool) bool {
	for idx, value := range slice {
		if !callbackFn(value, idx) {
			return false
		}
	}
	return true
}

// Reverse returns a new Slice with the elements in reverse order.
//
// Example:
//
//	s := Slice[int]{1, 2, 3, 4, 5}
//
//	reversed := s.Reverse()
//	// reversed is now Slice[int]{5, 4, 3, 2, 1}
//
// Note: The original Slice is not modified. A new Slice is created with the elements in reverse order.
func (slice Slice[T]) Reverse() Slice[T] {
	reversed := make(Slice[T], len(slice))
	for i := len(slice) - 1; i >= 0; i-- {
		reversed[len(slice)-1-i] = slice[i]
	}
	return reversed

}

// IndexFunc returns the index of the first element in the Slice that satisfies the provided callback function,
//
// or -1 if no such element is found.
//
// The callback function is called with one argument: the current element's value.
//
// It should return true if the element satisfies the condition, and false otherwise.
//
// Example:
//
//	s := Slice[int]{1, 2, 3, 4, 5}
//
//	index := s.IndexFunc(func(element int) bool {
//	    return element == 3
//	})
//	// index is now 2
//
// Note: The iteration stops as soon as the callback function returns true for any element.
func (slice Slice[T]) IndexFunc(f func(T) bool) int {
	for i := range slice {
		if f(slice[i]) {
			return i
		}
	}
	return -1
}

// ContainsFunc returns true if the Slice contains at least one element that satisfies the provided callback function,
//
// and false otherwise.
//
// The callback function is called with one argument: the current element's value.
//
// It should return true if the element satisfies the condition, and false otherwise.
//
// Example:
//
//	s := Slice[int]{1, 2, 3, 4, 5}
//
//	contains := s.ContainsFunc(func(element int) bool {
//	    return element == 3
//	})
//	// contains is now true
//
// Note: This method is a convenience wrapper around IndexFunc, and has the same iteration behavior.
func (slice Slice[T]) ContainsFunc(f func(T) bool) bool {
	return slice.IndexFunc(f) >= 0
}

// EqualFunc compares two slices elements for equality using a custom comparison function.
//
// Example:
//
//	s1 := Slice[int]{1,2,3,4,5}
//	s2 := Slice[int]{1,2,3,4,5}
//
// eq := func(e1, e2 int) bool{ return e1 == e2 }
// equals := s1.EqualFunc(s2, eq)
// //equals is now true
func (s1 Slice[T]) EqualFunc(s2 Slice[T], eq func(T, T) bool) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if !eq(s1[i], s2[i]) {
			return false
		}
	}
	return true
}

// Swap replace the elements at index i and j at the same time.
//
// This method is used by the data.Sort interface.
func (s Slice[T]) Swap(i, j int) {
	if i < 0 {
		i = len(s) + i
	}

	if j < 0 {
		j = len(s) + i
	}

	s[i], s[j] = s[j], s[i]
}

func (s Slice[T]) SortFunc(cmp func(T, T) int) Slice[T] {
	copy := Clone(s)
	SortFunc(copy, cmp)
	return copy
}
