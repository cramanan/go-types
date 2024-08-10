//go:build go1.21

package ordered

import "slices"

func (s Ordered[T]) Concat(sls ...Ordered[T]) Ordered[T] {
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
func (s Ordered[O]) BinarySearchFunc(target O, cmp func(O, O) int) (int, bool) {
	return slices.BinarySearchFunc(s, target, cmp)
}

// BinarySearch searches for target in a sorted slice and returns the position
// where target is found, or the position where target would appear in the
// sort order; it also returns a bool saying whether the target is really found
// in the slice. The slice must be sorted in increasing order.
func (s Ordered[O]) BinarySearch(target O) (int, bool) {
	return slices.BinarySearch(s, target)
}

// EqualFunc reports whether two slices are equal using an equality
// function on each pair of elements. If the lengths are different,
// EqualFunc returns false. Otherwise, the elements are compared in
// increasing index order, and the comparison stops at the first index
// for which eq returns false.
func (s1 Ordered[O]) EqualFunc(s2 Ordered[O], eq func(O, O) bool) bool {
	return slices.EqualFunc(s1, s2, eq)
}

// Equal reports whether two slices are equal: the same length and all
// elements equal. If the lengths are different, Equal returns false.
// Otherwise, the elements are compared in increasing index order, and the
// comparison stops at the first unequal pair.
// Floating point NaNs are not considered equal.
func (s1 Ordered[O]) Equal(s2 Ordered[O]) bool {
	return slices.Equal(s1, s2)
}

// CompareFunc is like [Compare] but uses a custom comparison function on each
// pair of elements.
// The result is the first non-zero result of cmp; if cmp always
// returns 0 the result is 0 if len(s1) == len(s2), -1 if len(s1) < len(s2),
// and +1 if len(s1) > len(s2).
func (s1 Ordered[O]) CompareFunc(s2 Ordered[O], cmp func(O, O) int) int {
	return slices.CompareFunc(s1, s2, cmp)
}

// Index returns the index of the first occurrence of v in s,
// or -1 if not present.
func (s Ordered[O]) Index(v O) int {
	return slices.Index(s, v)
}

// IndexFunc returns the first index i satisfying f(s[i]),
// or -1 if none do.
func (s Ordered[O]) IndexFunc(f func(O) bool) int {
	return slices.IndexFunc(s, f)
}

// ContainsFunc reports whether at least one
// element e of s satisfies f(e).
func (s Ordered[O]) ContainsFunc(f func(O) bool) bool {
	return slices.ContainsFunc(s, f)
}

// Contains reports whether v is present in s.
func (s Ordered[O]) Contains(v O) bool {
	return slices.Contains(s, v)
}

// Insert inserts the values v... into s at index i,
// returning the modified slice.
// The elements at s[i:] are shifted up to make room.
// In the returned slice r, r[i] == v[0],
// and r[i+len(v)] == value originally at r[i].
// Insert panics if i is out of range.
// This function is O(len(s) + len(v)).
func (s Ordered[O]) Insert(i int, v ...O) Ordered[O] {
	return slices.Insert(s, i, v...)
}

// Delete removes the elements s[i:j] from s, returning the modified slice.
// Delete panics if j > len(s) or s[i:j] is not a valid slice of s.
// Delete is O(len(s)-i), so if many items must be deleted, it is better to
// make a single call deleting them all together than to delete one at a time.
// Delete zeroes the elements s[len(s)-(j-i):len(s)].
func (s Ordered[O]) Delete(i, j int) Ordered[O] {
	return slices.Delete(s, i, j)
}

// DeleteFunc removes any elements from s for which del returns true,
// returning the modified slice.
// DeleteFunc zeroes the elements between the new length and the original length.
func (s Ordered[O]) DeleteFunc(del func(O) bool) Ordered[O] {
	return slices.DeleteFunc(s, del)
}

// Replace replaces the elements s[i:j] by the given v, and returns the
// modified slice.
// Replace panics if j > len(s) or s[i:j] is not a valid slice of s.
// When len(v) < (j-i), Replace zeroes the elements between the new length and the original length.
func (s Ordered[O]) Replace(i, j int, v ...O) Ordered[O] {
	return slices.Replace(s, i, j, v...)
}

// CompactFunc is like [Compact] but uses an equality function to compare elements.
// For runs of elements that compare equal, CompactFunc keeps the first one.
// CompactFunc zeroes the elements between the new length and the original length.
func (s Ordered[O]) CompactFunc(eq func(O, O) bool) Ordered[O] {
	return slices.CompactFunc(s, eq)
}

func (s Ordered[O]) Compact() Ordered[O] {
	return slices.Compact(s)
}

// Grow increases the slice's capacity, if necessary, to guarantee space for
// another n elements. After Grow(n), at least n elements can be appended
// to the slice without another allocation. If n is negative or too large to
// allocate the memory, Grow panics.
func (s Ordered[O]) Grow(n int) Ordered[O] {
	return slices.Grow(s, n)
}

// Clip removes unused capacity from the slice, returning s[:len(s):len(s)].
func (s Ordered[O]) Clip() Ordered[O] {
	return slices.Clip(s)
}

// Reverse reverses the elements of the slice in place.
func (s Ordered[O]) Reverse() Ordered[O] {
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
func (s Ordered[O]) SortFunc(cmp func(a, b O) int) {
	slices.SortFunc(s, cmp)
}

// Sort returns a new sorted Ordered slice.
//
// The Sort method creates a clone of the original Ordered slice, sorts it using the slices.Sort function,
// and returns the sorted clone. The original slice remains unchanged.
func (s Ordered[O]) Sort() Ordered[O] {
	clone := s.Clone()
	slices.Sort(clone)
	return clone
}

// SortStableFunc sorts the slice x while keeping the original order of equal
// elements, using cmp to compare elements in the same way as [SortFunc].
func (x Ordered[O]) SortStableFunc(cmp func(a, b O) int) {
	slices.SortStableFunc(x, cmp)
}

// IsSortedFunc reports whether x is sorted in ascending order, with cmp as the
// comparison function as defined by [SortFunc].
func (s Ordered[O]) IsSortedFunc(cmp func(a, b O) int) bool {
	return slices.IsSortedFunc(s, cmp)
}

// IsSorted reports whether x is sorted in ascending order.
func (s Ordered[O]) IsSorted() bool {
	return slices.IsSorted(s)
}

// MinFunc returns the minimal value in x, using cmp to compare elements.
// It panics if x is empty. If there is more than one minimal element
// according to the cmp function, MinFunc returns the first one.
func (s Ordered[O]) MinFunc(cmp func(a, b O) int) O {
	return slices.MinFunc(s, cmp)
}

// Min returns the minimal value in x. It panics if x is empty. For floating-point numbers, Min propagates NaNs (any NaN value in x forces the output to be NaN).
func (s Ordered[O]) Min() O {
	return slices.Min(s)
}

// MaxFunc returns the maximal value in x, using cmp to compare elements.
// It panics if x is empty. If there is more than one maximal element
// according to the cmp function, MaxFunc returns the first one.
func (s Ordered[O]) MaxFunc(cmp func(a, b O) int) O {
	return slices.MaxFunc(s, cmp)
}

// Max returns the maximal value in x. It panics if x is empty. For floating-point E, Max propagates NaNs (any NaN value in x forces the output to be NaN).
func (s Ordered[O]) Max() O {
	return slices.Max(s)
}

// Prepend adds the given values to the beginning of the slice.
//
// Example:
//
//	s := Ordered[int]{1, 2, 3}
//	s = s.Prepend(4, 5)
//	fmt.Println(s) // Output: [4 5 1 2 3]
func (s Ordered[T]) Prepend(values ...T) Ordered[T] {
	return append(values, s...)
}

// At returns the element at the specified index in the slice.
// If the index is negative, it counts from the end of the slice.
//
// Example:
//
//	s := Ordered[int]{1, 2, 3, 4, 5}
//	fmt.Println(s.At(0))  // Output: 1
//	fmt.Println(s.At(-1)) // Output: 5
func (s Ordered[T]) At(n int) T {
	if n < 0 {
		n = len(s) + n
	}
	return s[n]
}

// Append adds one or more elements to the end of the slice.
//
// Example:
//
//	s := Ordered[int]{1, 2, 3}
//	s = s.Append(4, 5)
//	fmt.Println(s) // Output: [1 2 3 4 5]
func (slice Ordered[T]) Append(values ...T) Ordered[T] {
	return append(slice, values...)
}

// ForEach iterates over the slice and calls the provided callback function for each element.
// The callback function is called with the current element and its index as arguments.
func (slice Ordered[T]) ForEach(callbackFn func(value T, index int)) {
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
func (slice Ordered[O]) Map(callbackFn func(O, int) O) (mapped Ordered[O]) {
	if callbackFn == nil {
		panic("callback function is nil")
	}
	for i, v := range slice {
		mapped = append(mapped, callbackFn(v, i))
	}
	return mapped
}

// Filter returns a new Ordered containing all elements for which the callbackFn returns true.
// The callbackFn is called for each element in the Ordered, with the element and its index as arguments.
// The order of elements in the resulting Ordered is the same as in the original Ordered.
func (slice Ordered[T]) Filter(callbackFn func(element T, index int) bool) (filtered Ordered[T]) {
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

// Some returns true if at least one element in the Ordered slice satisfies the callback function.
//
// The Some method iterates over the elements of the Ordered slice, passing each element and its index to the callback function.
// If the callback function returns true for any element, Some immediately returns true. If the callback function never returns true, Some returns false.
func (slice Ordered[T]) Some(callbackFn func(element T, index int) bool) bool {
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

// Every returns true if every element in the Ordered slice satisfies the callback function.
//
// The Every method iterates over the elements of the Ordered slice, passing each element and its index to the callback function.
// If the callback function returns false for any element, Every immediately returns false. If the callback function returns true for all elements, Every returns true.
func (slice Ordered[T]) Every(callbackFn func(element T, index int) bool) bool {
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
//	s := Ordered[int]{1, 2, 3, 4, 5}
//	fmt.Println(s.Len()) // Output: 5
func (s Ordered[o]) Len() int { return len(s) }

// Swap swaps the elements at indices i and j in the ordered collection.
// If i or j is negative, it counts from the end of the collection.
// The original collection is not modified, a new swapped collection is returned.
//
// Example:
//
//	o := NewOrdered[int]{1, 2, 3, 4, 5}
//	swapped := o.Swap(0, -1) // swap first and last elements
//	fmt.Println(swapped) // Output: [5 2 3 4 1]
func (o Ordered[O]) Swap(i, j int) (s Ordered[O]) {
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

// Less reports whether the element at index i is less than the element at index j.
func (a Ordered[O]) Less(i, j int) bool { return a[i] < a[j] }

// Clone returns a copy of the slice.
// The elements are copied using assignment, so this is a shallow clone.
func (s Ordered[O]) Clone() Ordered[O] {
	return slices.Clone(s)
}

// CountFunc returns the count of elements in the Ordered slice that satisfy the comparison function with the target element.
//
// The CountFunc method iterates over the elements of the Ordered slice, comparing each element with the target element using the provided comparison function.
// If the comparison function returns true for an element, it is counted.
func (s Ordered[O]) CountFunc(target O, callbackFn func(O, O) bool) (count int) {
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

// Count returns the count of occurrences of the target element in the Ordered slice.
//
// The Count method iterates over the elements of the Ordered slice, comparing each element with the target element for equality.
// If an element is equal to the target, it is counted.
func (s Ordered[O]) Count(target O) (count int) {
	for _, v := range s {
		if v == target {
			count++
		}
	}
	return count
}

// Fill returns a new Ordered slice where all elements are replaced with the given value.
//
// The Fill method creates a new slice and appends the given value for each element in the original slice.
func (s Ordered[O]) Fill(value O) (copy Ordered[O]) {
	for range s {
		copy = append(copy, value)
	}
	return copy
}

// Range returns a new Ordered slice that includes elements from index i up to, but not including, index j.
//
// If i or j is negative, it is treated as an offset from the end of the slice.
func (s Ordered[O]) Range(i, j int) (s2 Ordered[O]) {
	if i < 0 {
		i = len(s) + i
	}

	if j < 0 {
		j = len(s) + j
	}

	return s[i:j]
}

func (s Ordered[O]) Slice() []O {
	return s
}

// Esxperimental
// Range generates a sequence of numbers from start to end with a given step.
// The step value must be non-zero.
// func Range[
// 	B constraints.Integer | constraints.Float,
// 	S constraints.Integer | constraints.Float,

// ](start, end B, step S) (slice []S) {
// 	if step == 0 {
// 		panic("step cannot be 0")
// 	}

// 	switch {
// 	case step > 0:
// 		for i := S(start); i < S(end); i += step {

// 		}

// 	case step < 0:
// 		for i := S(start); i < S(end); i += step {

// 		}
// 	}
// }

// // experimental
// func RangeBetween[T constraints.Integer | constraints.Float](start T, end T, step T) (rng []T) {
// 	if end > start && step < 0 {
// 		step = -step
// 	}
// 	return Range(start, end, step)

// }
