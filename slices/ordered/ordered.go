package ordered

import (
	"cmp"

	"github.com/cramanan/go-types/slices"
)

type Ordered[O cmp.Ordered] slices.Slice[O]

// Len returns the length of the Ordered.
//
// Example:
//
// s := Ordered[int]{1, 2, 3, 4, 5}
// s.Len()  // returns 5
func (s Ordered[T]) Len() int { return len(s) }

func (s Ordered[O]) Less(i, j int) bool { return s[i] < s[j] }

// Swap replace the elements at index i and j at the same time.
//
// This method is used by the data.Sort interface.
func (s Ordered[T]) Swap(i, j int) {
	if i < 0 {
		i = len(s) + i
	}

	if j < 0 {
		j = len(s) + i
	}

	s[i], s[j] = s[j], s[i]
}

// NewOrdered returns a new slices.ISlice of type T, optionally populated with the provided values.
//
// If no values are provided, an empty ISlice is returned.
//
// Example:
//
//	s := slices.New(1, 2, 3) // returns a ISlice[int] containing [1, 2, 3]
//
// s := slices.New[string]() // returns an empty ISlice[string]
func New[T cmp.Ordered](values ...T) Ordered[T] {
	return append(*new(Ordered[T]), values...)
}

func From[O cmp.Ordered](s Ordered[O]) Ordered[O] {
	return Ordered[O](s)
}

// Prepend adds the given values to the beginning of the Ordered.
//
// This method returns a new Ordered containing the prepended values followed by the original elements.
//
// The original Ordered is not modified.
//
// Example:
//
//	s := Ordered[int]{1, 2, 3}
//
// s = s.Prepend(0) // returns a Ordered[int] containing [0, 1, 2, 3]
func (slice Ordered[T]) Prepend(values ...T) Ordered[T] {
	return append(values, slice...)
}

// At returns the element at the specified index in the Ordered.
//
// The index can be either positive (from the start of the Ordered) or negative (from the end of the Ordered).
//
// If the index is negative, it is treated as an offset from the end of the Ordered.
//
// Example:
//
//	s := Ordered[int]{1, 2, 3, 4, 5}
//
//	s.At(0)  // returns 1
//	s.At(-1) // returns 5
//	s.At(2)  // returns 3
//
// Note: This method does not perform bounds checking for performance reasons.
func (slice Ordered[T]) At(n int) T {
	if n < 0 {
		n = len(slice) + n
	}
	return slice[n]
}

// Append adds the given values to the end of the Ordered.
//
// This method returns a new Ordered containing the original elements followed by the appended values.
// The original Ordered is not modified.
//
// Example:
//
//	s := Ordered[int]{1, 2, 3}
//
// s = s.Append(4, 5) // returns a Ordered[int] containing [1, 2, 3, 4, 5]
func (slice Ordered[T]) Append(values ...T) Ordered[T] {
	return append(slice, values...)
}

// ForEach iterates over the elements of the Ordered and calls the provided callback function for each element.
//
// The callback function is called with two arguments: the current element's value and its index in the Ordered.
//
// Example:
//
//	s := Ordered[int]{1, 2, 3}
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
// Note: The callback function is called for each element in the Ordered, in order.
//
// If the callback function panics, the iteration is stopped and the panic is propagated.
func (slice Ordered[T]) ForEach(callbackFn func(value T, index int)) {
	for idx, v := range slice {
		callbackFn(v, idx)
	}
}

// Filter returns a new Ordered containing only the elements for which the provided callback function returns true.
//
// The callback function is called with two arguments: the current element's value and its index in the Ordered.
//
// It should return true if the element should be included in the filtered Ordered, and false otherwise.
//
// Example:
//
//	s := Ordered[int]{1, 2, 3, 4, 5}
//
//	evenNumbers := s.Filter(func(element int, index int) bool {
//	    return element % 2 == 0
//	})
//
//	// evenNumbers is now Ordered[int]{2, 4}
//
// Note: The callback function is called for each element in the Ordered, in order.
// The filtered Ordered is created by appending elements to a new Ordered, so the order of elements is preserved.
func (slice Ordered[T]) Filter(callbackFn func(element T, index int) bool) (filtered Ordered[T]) {
	filtered = make(Ordered[T], 0, len(slice))
	for idx, value := range slice {
		if callbackFn(value, idx) {
			filtered = append(filtered, value)
		}
	}
	return filtered
}

// Some returns true if at least one element in the Ordered satisfies the provided callback function.
//
// The callback function is called with two arguments: the current element's value and its index in the Ordered.
//
// It should return true if the element satisfies the condition, and false otherwise.
//
// Example:
//
//	s := Ordered[int]{1, 2, 3, 4, 5}
//
//	hasEvenNumber := s.Some(func(element int, index int) bool {
//	    return element % 2 == 0
//	})
//	// hasEvenNumber is now true
//
// Note: The iteration stops as soon as the callback function returns true for any element.
func (slice Ordered[T]) Some(callbackFn func(element T, index int) bool) bool {
	for idx, value := range slice {
		if callbackFn(value, idx) {
			return true
		}
	}
	return false
}

// Every returns true if all elements in the Ordered satisfy the provided callback function.
//
// The callback function is called with two arguments: the current element's value and its index in the Ordered.
//
// It should return true if the element satisfies the condition, and false otherwise.
//
// Example:
//
//	s := Ordered[int]{2, 4, 6, 8, 10}
//
//	allEvenNumbers := s.Every(func(element int, index int) bool {
//	    return element % 2 == 0
//	})
//	// allEvenNumbers is now true
//
// Note: The iteration stops as soon as the callback function returns false for any element.
func (slice Ordered[T]) Every(callbackFn func(element T, index int) bool) bool {
	for idx, value := range slice {
		if !callbackFn(value, idx) {
			return false
		}
	}
	return true
}

func (slice Ordered[O]) Map(callbackFn func(O, int) O) (mapped Ordered[O]) {
	for i, value := range slice {
		mapped = append(mapped, callbackFn(value, i))
	}
	return mapped
}

// Equal reports whether two slices are equal: the same length and all
// elements equal. If the lengths are different, Equal returns false.
// Otherwise, the elements are compared in increasing index order, and the
// comparison stops at the first unequal pair.
// Floating point NaNs are not considered equal.
func (s1 Ordered[O]) Equal(s2 Ordered[O]) bool {
	return slices.Equal(s1, s2)
}

// Compare compares the elements of s1 and s2, using [cmp.Compare] on each pair
// of elements. The elements are compared sequentially, starting at index 0,
// until one element is not equal to the other.
// The result of comparing the first non-matching elements is returned.
// If both slices are equal until one of them ends, the shorter slice is
// considered less than the longer one.
// The result is 0 if s1 == s2, -1 if s1 < s2, and +1 if s1 > s2.
func (s1 Ordered[O]) Compare(s2 Ordered[O]) int {
	return slices.Compare(s1, s2)
}

// Index returns the index of the first occurrence of v in s,
// or -1 if not present.
func (s Ordered[O]) Index(v O) int {
	return slices.Index(s, v)
}

// Contains reports whether v is present in s.
func (s Ordered[O]) Contains(v O) bool {
	return slices.Contains(s, v)
}

// Clone returns a copy of the slice.
// The elements are copied using assignment, so this is a shallow clone.
func (s Ordered[O]) Clone() Ordered[O] {
	return slices.Clone(s)
}

// Compact replaces consecutive runs of equal elements with a single copy.
// This is like the uniq command found on Unix.
// Compact modifies the contents of the slice s and returns the modified slice,
// which may have a smaller length.
// Compact zeroes the elements between the new length and the original length.
func (s Ordered[O]) Compact() Ordered[O] {
	return slices.Compact(s)

}

// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run $GOROOT/src/sort/gen_sort_variants.go -generic
//sort.go

// IsSorted reports whether x is sorted in ascending order.
func (s Ordered[O]) IsSorted() bool {
	return slices.IsSorted(s)
}

// BinarySearch searches for target in a sorted slice and returns the position
// where target is found, or the position where target would appear in the
// sort order; it also returns a bool saying whether the target is really found
// in the slice. The slice must be sorted in increasing order.
func (s Ordered[O]) BinarySearch(target O) (int, bool) {
	return slices.BinarySearch(s, target)
}

// Max returns the maximal value in x. It panics if x is empty.
// For floating-point E, Max propagates NaNs (any NaN value in x
// forces the output to be NaN).
func (s Ordered[O]) Max(x Ordered[O]) O {
	return slices.Max(x)
}

// Min returns the minimal value in x. It panics if x is empty.
// For floating-point numbers, Min propagates NaNs (any NaN value in x
// forces the output to be NaN).
func (s Ordered[O]) Min() O {
	return slices.Min(s)
}

// Sort sorts a slice of any ordered type in ascending order.
// When sorting floating-point numbers, NaNs are ordered before other values.
func (s Ordered[O]) Sort() {
	slices.Sort(s)
}

func (s Ordered[O]) BinarySearchFunc(target O, cmp func(O, O) int) (int, bool) {
	return slices.BinarySearchFunc(s, target, cmp)
}
