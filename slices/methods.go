package slices

import (
	"golang.org/x/exp/slices"
)

func (s Slice[T]) Slice() []T {
	return s
}

// Concat returns a new slice concatenating the passed in slices.
func (s Slice[T]) Concat(sls ...Slice[T]) Slice[T] {
	for _, v := range sls {
		s = append(s, v...)
	}
	return s
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
func (s Slice[O]) Reverse() Slice[O] {
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
	for idx, v := range slice {
		callbackFn(v, idx)
	}
}

// Filter returns a new Slice containing all elements for which the callbackFn returns true.
// The callbackFn is called for each element in the Slice, with the element and its index as arguments.
// The order of elements in the resulting Slice is the same as in the original Slice.
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
	for i, v := range slice {
		mapped = append(mapped, callbackFn(v, i))
	}
	return mapped
}

// Len returns the length of the ordered collection.
//
// Example:
//
//	s := Slice[int]{1, 2, 3, 4, 5}
//	fmt.Println(s.Len()) // Output: 5
func (s Slice[o]) Len() int { return len(s) }

// Swap swaps the elements at indices i and j in the ordered collection.
// If i or j is negative, it counts from the end of the collection.
// The original collection is not modified, a new swapped collection is returned.
//
// Example:
//
//	o := NewOrdered[int]{1, 2, 3, 4, 5}
//	swapped := o.Swap(0, -1) // swap first and last elements
//	fmt.Println(swapped) // Output: [5 2 3 4 1]
func (o Slice[O]) Swap(i, j int) (s Slice[O]) {
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
func (s Slice[O]) Clone() Slice[O] {
	return slices.Clone(s)
}

func (s Slice[T]) CountFunc(target T, cmp func(T, T) bool) (count int) {
	for _, v := range s {
		if cmp(v, target) {
			count++
		}
	}
	return count
}

func (s Slice[T]) Fill(value T) (copy Slice[T]) {
	for range s {
		copy = append(copy, value)
	}
	return copy
}

func (s Slice[T]) Range(i, j int) (s2 Slice[T]) {
	if i < 0 {
		i = len(s) + i
	}

	if j < 0 {
		j = len(s) + j
	}

	return s[i:j]
}
