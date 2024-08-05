package ordered

import (
	"github.com/cramanan/go-types/slices"
	"golang.org/x/exp/constraints"
)

type Ordered[O constraints.Ordered] slices.Slice[O]

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
func New[T constraints.Ordered](values ...T) Ordered[T] {
	return append(*new(Ordered[T]), values...)
}

func From[O constraints.Ordered](s Ordered[O]) Ordered[O] {
	return Ordered[O](s)
}

// Equal reports whether two slices are equal: the same length and all
// elements equal. If the lengths are different, Equal returns false.
// Otherwise, the elements are compared in increasing index order, and the
// comparison stops at the first unequal pair.
// Floating point NaNs are not considered equal.
func (s1 Ordered[O]) Equal(s2 Ordered[O]) bool {
	return slices.Equal(s1, s2)
}

// Compare compares the elements of s1 and s2, using [constraints.Compare] on each pair
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
