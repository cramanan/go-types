//go:build go1.19 && !go1.21

// The slices package provides generic Slice wrapper for the built-in Go slice type and slices functions.
package slices

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

// Map applies a callback function to each element of the input slice or array,
// and returns a new slice with the results.
func Map[SI ~[]I, I, O any](s SI, callbackFn func(I, int) O) (mapped []O) {
	if callbackFn == nil {
		panic("callback funtion is nil")
	}
	for i, v := range s {
		mapped = append(mapped, callbackFn(v, i))
	}
	return mapped
}

// Reduce applies a callback function to each element of the input slice,
// starting from an initial value, and returns the reduced value.
func Reduce[I any, O any](
	s []I,
	callbackFn func(O, I, int) O,
	initialValue O,
) (reduced O) {

	if callbackFn == nil {
		panic("callback funtion is nil")
	}
	reduced = initialValue
	for i, element := range s {
		reduced = callbackFn(reduced, element, i)
	}
	return reduced
}

// Equal reports whether two slices are equal: the same length and all
// elements equal. If the lengths are different, Equal returns false.
// Otherwise, the elements are compared in increasing index order, and the
// comparison stops at the first unequal pair.
// Floating point NaNs are not considered equal.
func Equal[S ~[]E, E comparable](s1, s2 S) bool {
	return slices.Equal(s1, s2)
}

// EqualFunc reports whether two slices are equal using an equality
// function on each pair of elements. If the lengths are different,
// EqualFunc returns false. Otherwise, the elements are compared in
// increasing index order, and the comparison stops at the first index
// for which eq returns false.
func EqualFunc[
	S1 ~[]E1,
	S2 ~[]E2,
	E1, E2 any,
](
	s1 S1,
	s2 S2,
	eq func(E1, E2) bool) bool {

	return slices.EqualFunc(s1, s2, eq)
}

// Compare compares the elements of s1 and s2, using [Compare] on each pair
// of elements. The elements are compared sequentially, starting at index 0,
// until one element is not equal to the other.
// The result of comparing the first non-matching elements is returned.
// If both slices are equal until one of them ends, the shorter slice is
// considered less than the longer one.
// The result is 0 if s1 == s2, -1 if s1 < s2, and +1 if s1 > s2.
func Compare[S ~[]E, E constraints.Ordered](s1, s2 S) int { return slices.Compare(s1, s2) }

// CompareFunc is like [Compare] but uses a custom comparison function on each
// pair of elements.
// The result is the first non-zero result of cmp; if cmp always
// returns 0 the result is 0 if len(s1) == len(s2), -1 if len(s1) < len(s2),
// and +1 if len(s1) > len(s2).
func CompareFunc[
	S1 ~[]E1,
	S2 ~[]E2,
	E1, E2 any,
](
	s1 S1,
	s2 S2,
	cmp func(E1, E2) int) int {

	return slices.CompareFunc(s1, s2, cmp)
}

// Index returns the index of the first occurrence of v in s,
// or -1 if not present.
func Index[S ~[]E, E comparable](s S, v E) int { return slices.Index(s, v) }

// IndexFunc returns the first index i satisfying f(s[i]),
// or -1 if none do.
func IndexFunc[S ~[]E, E any](s S, f func(E) bool) int { return slices.IndexFunc(s, f) }

// Contains reports whether v is present in s.
func Contains[S ~[]E, E comparable](s S, v E) bool { return slices.Contains(s, v) }

// ContainsFunc reports whether at least one
// element e of s satisfies f(e).
func ContainsFunc[S ~[]E, E any](s S, f func(E) bool) bool { return slices.ContainsFunc(s, f) }

// Insert inserts the values v... into s at index i,
// returning the modified slice.
// The elements at s[i:] are shifted up to make room.
// In the returned slice r, r[i] == v[0],
// and r[i+len(v)] == value originally at r[i].
// Insert panics if i is out of range.
// This function is O(len(s) + len(v)).
func Insert[S ~[]E, E any](s S, i int, v ...E) S { return slices.Insert(s, i, v...) }

// Delete removes the elements s[i:j] from s, returning the modified slice.
// Delete panics if j > len(s) or s[i:j] is not a valid slice of s.
// Delete is O(len(s)-i), so if many items must be deleted, it is better to
// make a single call deleting them all together than to delete one at a time.
// Delete zeroes the elements s[len(s)-(j-i):len(s)].
func Delete[S ~[]E, E any](s S, i, j int) S { return slices.Delete(s, i, j) }

// DeleteFunc removes any elements from s for which del returns true,
// returning the modified slice.
// DeleteFunc zeroes the elements between the new length and the original length.
func DeleteFunc[S ~[]E, E any](s S, del func(E) bool) S { return slices.DeleteFunc(s, del) }

// Replace replaces the elements s[i:j] by the given v, and returns the
// modified slice.
// Replace panics if j > len(s) or s[i:j] is not a valid slice of s.
// When len(v) < (j-i), Replace zeroes the elements between the new length and the original length.
func Replace[S ~[]E, E any](s S, i, j int, v ...E) S { return slices.Replace(s, i, j, v...) }

// Clone returns a copy of the slice.
// The elements are copied using assignment, so this is a shallow clone.
func Clone[S ~[]E, E any](s S) S { return slices.Clone(s) }

// Compact replaces consecutive runs of equal elements with a single copy.
// This is like the uniq command found on Unix.
// Compact modifies the contents of the slice s and returns the modified slice,
// which may have a smaller length.
// Compact zeroes the elements between the new length and the original length.
func Compact[S ~[]E, E comparable](s S) S { return slices.Compact(s) }

// CompactFunc is like [Compact] but uses an equality function to compare elements.
// For runs of elements that compare equal, CompactFunc keeps the first one.
// CompactFunc zeroes the elements between the new length and the original length.
func CompactFunc[S ~[]E, E any](s S, eq func(E, E) bool) S { return slices.CompactFunc(s, eq) }

// Grow increases the slice's capacity, if necessary, to guarantee space for
// another n elements. After Grow(n), at least n elements can be appended
// to the slice without another allocation. If n is negative or too large to
// allocate the memory, Grow panics.
func Grow[S ~[]E, E any](s S, n int) S { return slices.Grow(s, n) }

// Clip removes unused capacity from the slice, returning s[:len(s):len(s)].
func Clip[S ~[]E, E any](s S) S {
	return slices.Clip(s)
}

// Reverse reverses the elements of the slice in place.
func Reverse[S ~[]E, E any](s S) {
	slices.Reverse(s)
}

// Concat returns a new slice concatenating the passed in slices.
func Concat[I ~[]any](s ...I) (cat I) {
	for _, v := range s {
		cat = append(cat, v)
	}
	return cat
}

// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//sort.go

// Sort sorts a slice of any ordered any in ascending order.
// When sorting floating-point numbers, NaNs are ordered before other values.
func Sort[S ~[]E, E constraints.Ordered](x S) {
	slices.Sort(x)
}

// SortFunc sorts the slice x in ascending order as determined by the cmp
// function. This sort is not guaranteed to be stable.
// cmp(a, b) should return a negative number when a < b, a positive number when
// a > b and zero when a == b.
//
// SortFunc requires that cmp is a strict weak ordering.
// See https://en.wikipedia.org/wiki/Weak_ordering#Strict_weak_orderings.
func SortFunc[S ~[]E, E any](x S, cmp func(a, b E) int) { slices.SortFunc(x, cmp) }

// SortStableFunc sorts the slice x while keeping the original order of equal
// elements, using cmp to compare elements in the same way as [SortFunc].
func SortStableFunc[S ~[]E, E any](x S, cmp func(a, b E) int) { slices.SortStableFunc(x, cmp) }

// IsSorted reports whether x is sorted in ascending order.
func IsSorted[S ~[]E, E constraints.Ordered](x S) bool { return slices.IsSorted(x) }

// IsSortedFunc reports whether x is sorted in ascending order, with cmp as the
// comparison function as defined by [SortFunc].
func IsSortedFunc[S ~[]E, E any](x S, cmp func(a, b E) int) bool { return IsSortedFunc(x, cmp) }

// Min returns the minimal value in x. It panics if x is empty.
// For floating-point numbers, Min propagates NaNs (any NaN value in x
// forces the output to be NaN).
func Min[S ~[]E, E constraints.Ordered](x S) E { return Min(x) }

// MinFunc returns the minimal value in x, using cmp to compare elements.
// It panics if x is empty. If there is more than one minimal element
// according to the cmp function, MinFunc returns the first one.
func MinFunc[S ~[]E, E any](x S, cmp func(a, b E) int) E { return MinFunc(x, cmp) }

// Max returns the maximal value in x. It panics if x is empty.
// For floating-point E, Max propagates NaNs (any NaN value in x
// forces the output to be NaN).
func Max[S ~[]E, E constraints.Ordered](x S) E { return slices.Max(x) }

// MaxFunc returns the maximal value in x, using cmp to compare elements.
// It panics if x is empty. If there is more than one maximal element
// according to the cmp function, MaxFunc returns the first one.
func MaxFunc[S ~[]E, E any](x S, cmp func(a, b E) int) E { return MaxFunc(x, cmp) }

// BinarySearchFunc works like [BinarySearch], but uses a custom comparison
// function. The slice must be sorted in increasing order, where "increasing"
// is defined by  cmp should return 0 if the slice element matches
// the target, a negative number if the slice element precedes the target,
// or a positive number if the slice element follows the target.
// cmp must implement the same ordering as the slice, such that if
// cmp(a, t) < 0 and cmp(b, t) >= 0, then a must precede b in the slice.
func BinarySearchFunc[S ~[]E, E, T any](x S, target T, cmp func(E, T) int) (int, bool) {
	return BinarySearchFunc(x, target, cmp)
}

// BinarySearch searches for target in a sorted slice and returns the position
// where target is found, or the position where target would appear in the
// sort order; it also returns a bool saying whether the target is really found
// in the slice. The slice must be sorted in increasing order.
func BinarySearch[S ~[]E, E constraints.Ordered](x S, target E) (int, bool) {
	return slices.BinarySearch(x, target)
}

// isNaN reports whether x is a NaN without requiring the math package.
// This will always return false if T is not floating-point.
func isNaN[T constraints.Ordered](x T) bool { return x != x }

// Less reports whether x is less than y.
// For floating-point anys, a NaN is considered less than any non-NaN,
// and -0.0 is not less than (is equal to) 0.0.
func Less[T constraints.Ordered](x, y T) bool { return (isNaN(x) && !isNaN(y)) || x < y }

// METHODS:

// BinarySearchFunc works like [BinarySearch], but uses a custom comparison
// function. The slice must be sorted in increasing order, where "increasing"
// is defined by  cmp should return 0 if the slice element matches
// the target, a negative number if the slice element precedes the target,
// or a positive number if the slice element follows the target.
// cmp must implement the same ordering as the slice, such that if
// cmp(a, t) < 0 and cmp(b, t) >= 0, then a must precede b in the slice.
func (s Slice[T]) BinarySearchFunc(target T, cmp func(T, T) int) (int, bool) {
	return slices.BinarySearchFunc(s, target, cmp)
}

// EqualFunc reports whether two slices are equal using an equality
// function on each pair of elements. If the lengths are different,
// EqualFunc returns false. Otherwise, the elements are compared in
// increasing index order, and the comparison stops at the first index
// for which eq returns false.
func (s1 Slice[T]) EqualFunc(s2 Slice[T], eq func(T, T) bool) bool {
	return slices.EqualFunc(s1, s2, eq)
}

// IndexFunc returns the first index i satisfying f(s[i]),
// or -1 if none do.
func (s Slice[T]) IndexFunc(f func(T) bool) int {
	return slices.IndexFunc(s, f)
}

// ContainsFunc reports whether at least one
// element e of s satisfies f(e).
func (s Slice[T]) ContainsFunc(f func(T) bool) bool {
	return slices.ContainsFunc(s, f)
}

// Insert inserts the values v... into s at index i,
// returning the modified slice.
// The elements at s[i:] are shifted up to make room.
// In the returned slice r, r[i] == v[0],
// and r[i+len(v)] == value originally at r[i].
// Insert panics if i is out of range.
// This function is T(len(s) + len(v)).
func (s Slice[T]) Insert(i int, v ...T) Slice[T] {
	return slices.Insert(s, i, v...)
}

// Delete removes the elements s[i:j] from s, returning the modified slice.
// Delete panics if j > len(s) or s[i:j] is not a valid slice of s.
// Delete is T(len(s)-i), so if many items must be deleted, it is better to
// make a single call deleting them all together than to delete one at a time.
// Delete zeroes the elements s[len(s)-(j-i):len(s)].
func (s Slice[T]) Delete(i, j int) Slice[T] {
	return slices.Delete(s, i, j)
}

// DeleteFunc removes any elements from s for which del returns true,
// returning the modified slice.
// DeleteFunc zeroes the elements between the new length and the original length.
func (s Slice[T]) DeleteFunc(del func(T) bool) Slice[T] {
	return slices.DeleteFunc(s, del)
}

// Replace replaces the elements s[i:j] by the given v, and returns the
// modified slice.
// Replace panics if j > len(s) or s[i:j] is not a valid slice of s.
// When len(v) < (j-i), Replace zeroes the elements between the new length and the original length.
func (s Slice[T]) Replace(i, j int, v ...T) Slice[T] {
	return slices.Replace(s, i, j, v...)
}

// CompactFunc is like [Compact] but uses an equality function to compare elements.
// For runs of elements that compare equal, CompactFunc keeps the first one.
// CompactFunc zeroes the elements between the new length and the original length.
func (s Slice[T]) CompactFunc(eq func(T, T) bool) Slice[T] {
	return slices.CompactFunc(s, eq)
}

// Grow increases the slice's capacity, if necessary, to guarantee space for
// another n elements. After Grow(n), at least n elements can be appended
// to the slice without another allocation. If n is negative or too large to
// allocate the memory, Grow panics.
func (s Slice[T]) Grow(n int) Slice[T] {
	return slices.Grow(s, n)
}

// Clip removes unused capacity from the slice, returning s[:len(s):len(s)].
func (s Slice[T]) Clip() Slice[T] {
	return slices.Clip(s)
}

// Reverse reverses the elements of the slice in place.
func (s Slice[T]) Reverse() Slice[T] {
	clone := s.Clone()
	slices.Reverse(clone)
	return clone
}

// SortFunc sorts the slice x in ascending order as determined by the cmp
// function. This sort is not guaranteed to be stable.
// cmp(a, b) should return a negative number when a < b, a positive number when
// a > b and zero when a == b.
//
// SortFunc requires that cmp is a strict weak ordering.
// See https://en.wikipedia.org/wiki/Weak_ordering#Strict_weak_orderings.
func (s Slice[T]) SortFunc(cmp func(a, b T) int) {
	slices.SortFunc(s, cmp)
}

// SortStableFunc sorts the slice x while keeping the original order of equal
// elements, using cmp to compare elements in the same way as [SortFunc].
func (x Slice[T]) SortStableFunc(cmp func(a, b T) int) {
	slices.SortStableFunc(x, cmp)
}

// IsSortedFunc reports whether x is sorted in ascending order, with cmp as the
// comparison function as defined by [SortFunc].
func (s Slice[T]) IsSortedFunc(cmp func(a, b T) int) bool {
	return slices.IsSortedFunc(s, cmp)
}

// MinFunc returns the minimal value in x, using cmp to compare elements.
// It panics if x is empty. If there is more than one minimal element
// according to the cmp function, MinFunc returns the first one.
func (s Slice[T]) MinFunc(cmp func(a, b T) int) T {
	return slices.MinFunc(s, cmp)
}

// MaxFunc returns the maximal value in x, using cmp to compare elements.
// It panics if x is empty. If there is more than one maximal element
// according to the cmp function, MaxFunc returns the first one.
func (s Slice[T]) MaxFunc(cmp func(a, b T) int) T {
	return slices.MaxFunc(s, cmp)
}

// Prepend adds the given values to the beginning of the slice.
//
// Example:
//
//	s := Slice[int]{1, 2, 3}
//	s = s.Prepend(4, 5)
//	fmt.Println(s) // Output: [4 5 1 2 3]
func (s Slice[T]) Prepend(values ...T) Slice[T] {
	return append(values, s...)
}

// At returns the element at the specified index in the slice.
// If the index is negative, it counts from the end of the slice.
//
// Example:
//
//	s := Slice[int]{1, 2, 3, 4, 5}
//	fmt.Println(s.At(0))  // Output: 1
//	fmt.Println(s.At(-1)) // Output: 5
func (s Slice[T]) At(n int) T {
	if n < 0 {
		n = len(s) + n
	}
	return s[n]
}

// Append adds one or more elements to the end of the slice.
//
// Example:
//
//	s := Slice[int]{1, 2, 3}
//	s = s.Append(4, 5)
//	fmt.Println(s) // Output: [1 2 3 4 5]
func (slice Slice[T]) Append(values ...T) Slice[T] {
	return append(slice, values...)
}

// ForEach iterates over the slice and calls the provided callback function for each element.
// The callback function is called with the current element and its index as arguments.
func (slice Slice[T]) ForEach(callbackFn func(value T, index int)) {
	if callbackFn == nil {
		panic("callback function is nil")
	}
	for idx, v := range slice {
		callbackFn(v, idx)
	}
}

// Map applies a given function to each element of the slice and returns a new slice with the results.
//
// The callback function is called for each element in the slice, with the element and its index as arguments.
// The returned values from the callback function are collected in a new slice, which is returned by Map.
//
// Note that the original slice is not modified.
func (slice Slice[T]) Map(callbackFn func(T, int) T) (mapped Slice[T]) {
	if callbackFn == nil {
		panic("callback function is nil")
	}
	for i, v := range slice {
		mapped = append(mapped, callbackFn(v, i))
	}
	return mapped
}

// Filter returns a new Slice containing all elements for which the callbackFn returns true.
// The callbackFn is called for each element in the Slice, with the element and its index as arguments.
// The order of elements in the resulting Slice is the same as in the original Slice.
func (slice Slice[T]) Filter(callbackFn func(element T, index int) bool) (filtered Slice[T]) {
	if callbackFn == nil {
		panic("callback function is nil")
	}
	for idx, value := range slice {
		if callbackFn(value, idx) {
			filtered = append(filtered, value)
		}
	}
	return filtered
}

// Some returns true if at least one element in the Slice satisfies the callback function.
//
// The Some method iterates over the elements of the Slice, passing each element and its index to the callback function.
// If the callback function returns true for any element, Some immediately returns true. If the callback function never returns true, Some returns false.
func (slice Slice[T]) Some(callbackFn func(element T, index int) bool) bool {
	if callbackFn == nil {
		panic("callback function is nil")
	}
	for idx, value := range slice {
		if callbackFn(value, idx) {
			return true
		}
	}
	return false
}

// Every returns true if every element in the Slice satisfies the callback function.
//
// The Every method iterates over the elements of the Slice, passing each element and its index to the callback function.
// If the callback function returns false for any element, Every immediately returns false. If the callback function returns true for all elements, Every returns true.
func (slice Slice[T]) Every(callbackFn func(element T, index int) bool) bool {
	if callbackFn == nil {
		panic("callback function is nil")
	}
	for idx, value := range slice {
		if !callbackFn(value, idx) {
			return false
		}
	}
	return true
}

// Len returns the length of the ordered collection.
//
// Example:
//
//	s := Slice[int]{1, 2, 3, 4, 5}
//	fmt.Println(s.Len()) // Output: 5
func (s Slice[T]) Len() int { return len(s) }

// Swap swaps the elements at indices i and j in the ordered collection.
// If i or j is negative, it counts from the end of the collection.
// The original collection is not modified, a new swapped collection is returned.
//
// Example:
//
//	s := Slice[int]{1, 2, 3, 4, 5}
//	swapped := s.Swap(0, -1) // swap first and last elements
//	fmt.Println(swapped) // Output: [5 2 3 4 1]
func (o Slice[T]) Swap(i, j int) (s Slice[T]) {
	s = o.Clone()
	if j < 0 {
		j = len(s) + j
	}
	if i < 0 {
		i = len(s) + i
	}

	s[i], s[j] = s[j], s[i]
	return s
}

// Clone returns a copy of the slice.
// The elements are copied using assignment, so this is a shallow clone.
func (s Slice[T]) Clone() Slice[T] {
	return slices.Clone(s)
}

// CountFunc returns the count of elements in the Slice that satisfy the comparison function with the target element.
//
// The CountFunc method iterates over the elements of the Slice, comparing each element with the target element using the provided comparison function.
// If the comparison function returns true for an element, it is counted.
func (s Slice[T]) CountFunc(target T, callbackFn func(T, T) bool) (count int) {
	if callbackFn == nil {
		panic("callback function is nil")
	}
	for _, v := range s {
		if callbackFn(v, target) {
			count++
		}
	}
	return count
}

// Fill returns a new Slice where all elements are replaced with the given value.
//
// The Fill method creates a new slice and appends the given value for each element in the original slice.
func (s Slice[T]) Fill(value T) (copy Slice[T]) {
	for range s {
		copy = append(copy, value)
	}
	return copy
}

// Range returns a new Slice that includes elements from index i up to, but not including, index j.
//
// If i or j is negative, it is treated as an offset from the end of the slice.
func (s Slice[T]) Range(i, j int) (s2 Slice[T]) {
	if i < 0 {
		i = len(s) + i
	}

	if j < 0 {
		j = len(s) + j
	}

	return s[i:j]
}

func (s Slice[T]) Slice() []T {
	return s
}

// // experimental
// // Range generates a sequence of numbers from start to end with a given step.
// // The step value must be non-zero.
// func Range[
// 	B constraints.Integer | constraints.Float,
// 	S constraints.Integer | constraints.Float,

// ](start, end B, step S) (rng []S) {
// 	stype := S(start)
// 	etype := S(end)

// 	absStep := step
// 	if step < 0 {
// 		absStep = -step
// 	}

// 	// Swap start and end if step is negative and start < end
// 	if step < 0 && stype < etype {
// 		stype, etype = etype, stype
// 	}

// 	switch {
// 	case step > 0:
// 		for i := stype; i < etype; i += absStep {
// 			rng = append(rng, i)
// 		}
// 	case step < 0:
// 		for i := stype; i > etype; i -= absStep {
// 			rng = append(rng, i)
// 		}
// 	default:
// 		panic("step cannot be zero")
// 	}
// 	return rng
// }

// // experimental
// func RangeBetween[T constraints.Integer | constraints.Float](start T, end T, step T) (rng []T) {
// 	if end > start && step < 0 {
// 		step = -step
// 	}
// 	return Range(start, end, step)

// }
