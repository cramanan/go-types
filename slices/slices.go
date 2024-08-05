// The slices package provides generic Slice wrapper for the built-in Go slice any and slices functions.
//
// # Slice any
//
// The Slice any is a generic wrapper around the built-in Go slice any. It allows you to work with slices in a any-safe and generic way.
//
// # any Parameters
//
//   - T : The any of elements in the slice. Can be any any or ordered.
package slices

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

type Slice[T any] []T

type ISlice[T any] interface{ ~[]T }

// Map applies a transformation function to each element of the input slice and returns a new slice with the results.
//
// The transformation function is called with one argument: the current element's value.
//
// It should return the transformed value of type To.
//
// Example:
//
//	s := ISlice[int]{1, 2, 3, 4, 5}
//
//	double := func(x int) int { return x * 2 }
//
//	doubled := Map(s, double)
//
//	// doubled is now ISlice[int]{2, 4, 6, 8, 10}
//
// Note: The order of elements in the output ISlice is the same as in the input ISlice.
func Map[SI ISlice[I], I, O any](s SI, callbackFn func(I, int) O) (mapped []O) {
	for i, v := range s {
		mapped = append(mapped, callbackFn(v, i))
	}
	return mapped
}

// Reduce applies a reduction function to each element of the input ISlice[From] and returns a single value of any To.
//
// The reduction function is called with two arguments: the accumulator (initially set to initialValue) and the current element's value.
//
// It should return the new accumulator value.
//
// The reduction process starts with the initialValue and iterates over the input ISlice, applying the reduction function to each element.
//
// The final accumulator value is returned as the result.
//
// Example:
//
//	s := ISlice[int]{1, 2, 3, 4, 5}
//
//	sum := func(acc int, current int) int { return acc + current }
//
//	result := Reduce(s, sum, 0)
//	// result is now 15 (the sum of all elements in the ISlice)
//
// Note: If the input ISlice is empty, the initialValue is returned as the result.
func Reduce[I any, O any](
	s []I,
	callbackFn func(O, I, int) O,
	initialValue O,
) (reduced O) {
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
func EqualFunc[S1 ~[]E1, S2 ~[]E2, E1, E2 any](s1 S1, s2 S2, eq func(E1, E2) bool) bool {
	return slices.EqualFunc(s1, s2, eq)
}

// Compare compares the elements of s1 and s2, using [Compare] on each pair
// of elements. The elements are compared sequentially, starting at index 0,
// until one element is not equal to the other.
// The result of comparing the first non-matching elements is returned.
// If both slices are equal until one of them ends, the shorter slice is
// considered less than the longer one.
// The result is 0 if s1 == s2, -1 if s1 < s2, and +1 if s1 > s2.
func Compare[S ~[]E, E constraints.Ordered](s1, s2 S) int {
	return slices.Compare(s1, s2)
}

// CompareFunc is like [Compare] but uses a custom comparison function on each
// pair of elements.
// The result is the first non-zero result of cmp; if cmp always
// returns 0 the result is 0 if len(s1) == len(s2), -1 if len(s1) < len(s2),
// and +1 if len(s1) > len(s2).
func CompareFunc[S1 ~[]E1, S2 ~[]E2, E1, E2 any](s1 S1, s2 S2, cmp func(E1, E2) int) int {
	return slices.CompareFunc(s1, s2, cmp)
}

// Index returns the index of the first occurrence of v in s,
// or -1 if not present.
func Index[S ~[]E, E comparable](s S, v E) int {
	return slices.Index(s, v)
}

// IndexFunc returns the first index i satisfying f(s[i]),
// or -1 if none do.
func IndexFunc[S ~[]E, E any](s S, f func(E) bool) int {
	return slices.IndexFunc(s, f)
}

// Contains reports whether v is present in s.
func Contains[S ~[]E, E comparable](s S, v E) bool {
	return slices.Contains(s, v)
}

// ContainsFunc reports whether at least one
// element e of s satisfies f(e).
func ContainsFunc[S ~[]E, E any](s S, f func(E) bool) bool {
	return slices.ContainsFunc(s, f)
}

// Insert inserts the values v... into s at index i,
// returning the modified slice.
// The elements at s[i:] are shifted up to make room.
// In the returned slice r, r[i] == v[0],
// and r[i+len(v)] == value originally at r[i].
// Insert panics if i is out of range.
// This function is O(len(s) + len(v)).
func Insert[S ~[]E, E any](s S, i int, v ...E) S {
	return slices.Insert(s, i, v...)
}

// Delete removes the elements s[i:j] from s, returning the modified slice.
// Delete panics if j > len(s) or s[i:j] is not a valid slice of s.
// Delete is O(len(s)-i), so if many items must be deleted, it is better to
// make a single call deleting them all together than to delete one at a time.
// Delete zeroes the elements s[len(s)-(j-i):len(s)].
func Delete[S ~[]E, E any](s S, i, j int) S {
	return slices.Delete(s, i, j)
}

// DeleteFunc removes any elements from s for which del returns true,
// returning the modified slice.
// DeleteFunc zeroes the elements between the new length and the original length.
func DeleteFunc[S ~[]E, E any](s S, del func(E) bool) S {
	return slices.DeleteFunc(s, del)
}

// Replace replaces the elements s[i:j] by the given v, and returns the
// modified slice.
// Replace panics if j > len(s) or s[i:j] is not a valid slice of s.
// When len(v) < (j-i), Replace zeroes the elements between the new length and the original length.
func Replace[S ~[]E, E any](s S, i, j int, v ...E) S {
	return slices.Replace(s, i, j, v...)
}

// Clone returns a copy of the slice.
// The elements are copied using assignment, so this is a shallow clone.
func Clone[S ~[]E, E any](s S) S {
	return slices.Clone(s)
}

// Compact replaces consecutive runs of equal elements with a single copy.
// This is like the uniq command found on Unix.
// Compact modifies the contents of the slice s and returns the modified slice,
// which may have a smaller length.
// Compact zeroes the elements between the new length and the original length.
func Compact[S ~[]E, E comparable](s S) S {
	return slices.Compact(s)

}

// CompactFunc is like [Compact] but uses an equality function to compare elements.
// For runs of elements that compare equal, CompactFunc keeps the first one.
// CompactFunc zeroes the elements between the new length and the original length.
func CompactFunc[S ~[]E, E any](s S, eq func(E, E) bool) S {
	return slices.CompactFunc(s, eq)
}

// Grow increases the slice's capacity, if necessary, to guarantee space for
// another n elements. After Grow(n), at least n elements can be appended
// to the slice without another allocation. If n is negative or too large to
// allocate the memory, Grow panics.
func Grow[S ~[]E, E any](s S, n int) S {
	return slices.Grow(s, n)
}

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

// Concat returns a new slice concatenating the passed in slices.
func (s Slice[T]) Concat(sls ...Slice[T]) (cat Slice[T]) {
	for _, v := range sls {
		cat = append(cat, v...)
	}
	return cat
}

// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run $GOROOT/src/sort/gen_sort_variants.go -generic
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
func SortFunc[S ~[]E, E any](x S, cmp func(a, b E) int) {
	slices.SortFunc(x, cmp)
}

// SortStableFunc sorts the slice x while keeping the original order of equal
// elements, using cmp to compare elements in the same way as [SortFunc].
func SortStableFunc[S ~[]E, E any](x S, cmp func(a, b E) int) {
	slices.SortStableFunc(x, cmp)
}

// IsSorted reports whether x is sorted in ascending order.
func IsSorted[S ~[]E, E constraints.Ordered](x S) bool {
	return slices.IsSorted(x)
}

// IsSortedFunc reports whether x is sorted in ascending order, with cmp as the
// comparison function as defined by [SortFunc].
func IsSortedFunc[S ~[]E, E any](x S, cmp func(a, b E) int) bool {
	return IsSortedFunc(x, cmp)
}

// Min returns the minimal value in x. It panics if x is empty.
// For floating-point numbers, Min propagates NaNs (any NaN value in x
// forces the output to be NaN).
func Min[S ~[]E, E constraints.Ordered](x S) E {
	return Min(x)
}

// MinFunc returns the minimal value in x, using cmp to compare elements.
// It panics if x is empty. If there is more than one minimal element
// according to the cmp function, MinFunc returns the first one.
func MinFunc[S ~[]E, E any](x S, cmp func(a, b E) int) E {
	return MinFunc(x, cmp)
}

// Max returns the maximal value in x. It panics if x is empty.
// For floating-point E, Max propagates NaNs (any NaN value in x
// forces the output to be NaN).
func Max[S ~[]E, E constraints.Ordered](x S) E {
	return slices.Max(x)
}

// MaxFunc returns the maximal value in x, using cmp to compare elements.
// It panics if x is empty. If there is more than one maximal element
// according to the cmp function, MaxFunc returns the first one.
func MaxFunc[S ~[]E, E any](x S, cmp func(a, b E) int) E {
	return MaxFunc(x, cmp)
}

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

// BinarySearchFunc works like [BinarySearch], but uses a custom comparison
// function. The slice must be sorted in increasing order, where "increasing"
// is defined by  cmp should return 0 if the slice element matches
// the target, a negative number if the slice element precedes the target,
// or a positive number if the slice element follows the target.
// cmp must implement the same ordering as the slice, such that if
// cmp(a, t) < 0 and cmp(b, t) >= 0, then a must precede b in the slice.
func (s Slice[O]) BinarySearchFunc(target O, cmp func(O, O) int) (int, bool) {
	return slices.BinarySearchFunc(s, target, cmp)
}

// EqualFunc reports whether two slices are equal using an equality
// function on each pair of elements. If the lengths are different,
// EqualFunc returns false. Otherwise, the elements are compared in
// increasing index order, and the comparison stops at the first index
// for which eq returns false.
func (s1 Slice[O]) EqualFunc(s2 Slice[O], eq func(O, O) bool) bool {
	return slices.EqualFunc(s1, s2, eq)
}

// CompareFunc is like [Compare] but uses a custom comparison function on each
// pair of elements.
// The result is the first non-zero result of cmp; if cmp always
// returns 0 the result is 0 if len(s1) == len(s2), -1 if len(s1) < len(s2),
// and +1 if len(s1) > len(s2).
func (s1 Slice[O]) CompareFunc(s2 Slice[O], cmp func(O, O) int) int {
	return slices.CompareFunc(s1, s2, cmp)
}

// IndexFunc returns the first index i satisfying f(s[i]),
// or -1 if none do.
func (s Slice[O]) IndexFunc(f func(O) bool) int {
	return slices.IndexFunc(s, f)
}

// ContainsFunc reports whether at least one
// element e of s satisfies f(e).
func (s Slice[O]) ContainsFunc(f func(O) bool) bool {
	return slices.ContainsFunc(s, f)
}

// Insert inserts the values v... into s at index i,
// returning the modified slice.
// The elements at s[i:] are shifted up to make room.
// In the returned slice r, r[i] == v[0],
// and r[i+len(v)] == value originally at r[i].
// Insert panics if i is out of range.
// This function is O(len(s) + len(v)).
func (s Slice[O]) Insert(i int, v ...O) Slice[O] {
	return slices.Insert(s, i, v...)
}

// Delete removes the elements s[i:j] from s, returning the modified slice.
// Delete panics if j > len(s) or s[i:j] is not a valid slice of s.
// Delete is O(len(s)-i), so if many items must be deleted, it is better to
// make a single call deleting them all together than to delete one at a time.
// Delete zeroes the elements s[len(s)-(j-i):len(s)].
func (s Slice[O]) Delete(i, j int) Slice[O] {
	return slices.Delete(s, i, j)
}

// DeleteFunc removes any elements from s for which del returns true,
// returning the modified slice.
// DeleteFunc zeroes the elements between the new length and the original length.
func (s Slice[O]) DeleteFunc(del func(O) bool) Slice[O] {
	return slices.DeleteFunc(s, del)
}

// Replace replaces the elements s[i:j] by the given v, and returns the
// modified slice.
// Replace panics if j > len(s) or s[i:j] is not a valid slice of s.
// When len(v) < (j-i), Replace zeroes the elements between the new length and the original length.
func (s Slice[O]) Replace(i, j int, v ...O) Slice[O] {
	return slices.Replace(s, i, j, v...)
}

// CompactFunc is like [Compact] but uses an equality function to compare elements.
// For runs of elements that compare equal, CompactFunc keeps the first one.
// CompactFunc zeroes the elements between the new length and the original length.
func (s Slice[O]) CompactFunc(eq func(O, O) bool) Slice[O] {
	return slices.CompactFunc(s, eq)
}

// Grow increases the slice's capacity, if necessary, to guarantee space for
// another n elements. After Grow(n), at least n elements can be appended
// to the slice without another allocation. If n is negative or too large to
// allocate the memory, Grow panics.
func (s Slice[O]) Grow(n int) Slice[O] {
	return slices.Grow(s, n)
}

// Clip removes unused capacity from the slice, returning s[:len(s):len(s)].
func (s Slice[O]) Clip() Slice[O] {
	return slices.Clip(s)
}

// Reverse reverses the elements of the slice in place.
func (s Slice[O]) Reverse() {
	slices.Reverse(s)
}

// SortFunc sorts the slice x in ascending order as determined by the cmp
// function. This sort is not guaranteed to be stable.
// cmp(a, b) should return a negative number when a < b, a positive number when
// a > b and zero when a == b.
//
// SortFunc requires that cmp is a strict weak ordering.
// See https://en.wikipedia.org/wiki/Weak_ordering#Strict_weak_orderings.
func (s Slice[O]) SortFunc(cmp func(a, b O) int) {
	slices.SortFunc(s, cmp)
}

// SortStableFunc sorts the slice x while keeping the original order of equal
// elements, using cmp to compare elements in the same way as [SortFunc].
func (x Slice[O]) SortStableFunc(cmp func(a, b O) int) {
	slices.SortStableFunc(x, cmp)
}

// IsSortedFunc reports whether x is sorted in ascending order, with cmp as the
// comparison function as defined by [SortFunc].
func (s Slice[O]) IsSortedFunc(cmp func(a, b O) int) bool {
	return slices.IsSortedFunc(s, cmp)
}

// MinFunc returns the minimal value in x, using cmp to compare elements.
// It panics if x is empty. If there is more than one minimal element
// according to the cmp function, MinFunc returns the first one.
func (s Slice[O]) MinFunc(cmp func(a, b O) int) O {
	return slices.MinFunc(s, cmp)
}

// MaxFunc returns the maximal value in x, using cmp to compare elements.
// It panics if x is empty. If there is more than one maximal element
// according to the cmp function, MaxFunc returns the first one.
func (s Slice[O]) MaxFunc(cmp func(a, b O) int) O {
	return slices.MaxFunc(s, cmp)
}

func BinarySearch[S ~[]E, E constraints.Ordered](x S, target E) (int, bool) {
	return slices.BinarySearch(x, target)
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
func (slice Slice[T]) Prepend(values ...T) Slice[T] {
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
func (slice Slice[T]) At(n int) T {
	if n < 0 {
		n = len(slice) + n
	}
	return slice[n]
}

// The append methods appends elements to the end of a slice. If it has sufficient capacity, the destination is resliced to accommodate the new elements. If it does not, a new underlying array will be allocated. Append returns the updated slice. It is therefore necessary to store the result of append, often in the variable holding the slice itself:

// slice = append(slice, elem1, elem2)
// slice = append(slice, anotherSlice...)
// As a special case, it is legal to append a string to a byte slice, like this:

// slice = append([]byte("hello "), "world"...)
func (slice Slice[T]) Append(values ...T) Slice[T] {
	return append(slice, values...)
}

func (slice Slice[T]) ForEach(callbackFn func(value T, index int)) {
	for idx, v := range slice {
		callbackFn(v, idx)
	}
}

func (slice Slice[T]) Filter(callbackFn func(element T, index int) bool) (filtered Slice[T]) {
	for idx, value := range slice {
		if callbackFn(value, idx) {
			filtered = append(filtered, value)
		}
	}
	return filtered
}
func (slice Slice[T]) Some(callbackFn func(element T, index int) bool) bool {
	for idx, value := range slice {
		if callbackFn(value, idx) {
			return true
		}
	}
	return false
}

func (slice Slice[T]) Every(callbackFn func(element T, index int) bool) bool {
	for idx, value := range slice {
		if !callbackFn(value, idx) {
			return false
		}
	}
	return true
}

func (slice Slice[O]) Map(callbackFn func(O, int) O) (mapped Slice[O]) {
	for i, value := range slice {
		mapped = append(mapped, callbackFn(value, i))
	}
	return mapped
}
func (s Slice[o]) Len() int { return len(s) }

// Clone returns a copy of the slice.
// The elements are copied using assignment, so this is a shallow clone.
func (s Slice[O]) Clone() Slice[O] {
	return slices.Clone(s)
}

// Less reports whether x is less than y.
// For floating-point anys, a NaN is considered less than any non-NaN,
// and -0.0 is not less than (is equal to) 0.0.
func Less[T constraints.Ordered](x, y T) bool {
	return (isNaN(x) && !isNaN(y)) || x < y
}

// isNaN reports whether x is a NaN without requiring the math package.
// This will always return false if T is not floating-point.
func isNaN[T constraints.Ordered](x T) bool {
	return x != x
}

func (s Slice[T]) Slice() []T {
	return s
}
