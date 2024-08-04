// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slices_test

import (
	"cmp"
	"math"
	"strings"
	"testing"

	. "github.com/cramanan/go-types/slices"
)

var equalIntTests = []struct {
	s1, s2 Slice[int]
	want   bool
}{
	{
		Slice[int]{1},
		nil,
		false,
	},
	{
		Slice[int]{},
		nil,
		true,
	},
	{
		Slice[int]{1, 2, 3},
		Slice[int]{1, 2, 3},
		true,
	},
	{
		Slice[int]{1, 2, 3},
		Slice[int]{1, 2, 3, 4},
		false,
	},
}

var equalFloatTests = []struct {
	s1, s2       Slice[float64]
	wantEqual    bool
	wantEqualNaN bool
}{
	{
		Slice[float64]{1, 2},
		Slice[float64]{1, 2},
		true,
		true,
	},
	{
		Slice[float64]{1, 2, math.NaN()},
		Slice[float64]{1, 2, math.NaN()},
		false,
		true,
	},
}

func TestEqual(t *testing.T) {
	for _, test := range equalIntTests {
		if got := Equal(test.s1, test.s2); got != test.want {
			t.Errorf("Equal(%v, %v) = %t, want %t", test.s1, test.s2, got, test.want)
		}
	}
	for _, test := range equalFloatTests {
		if got := Equal(test.s1, test.s2); got != test.wantEqual {
			t.Errorf("Equal(%v, %v) = %t, want %t", test.s1, test.s2, got, test.wantEqual)
		}
	}
}

// equal is simply ==.
func equal[T comparable](v1, v2 T) bool {
	return v1 == v2
}

// equalNaN is like == except that all NaNs are equal.
func equalNaN[T comparable](v1, v2 T) bool {
	isNaN := func(f T) bool { return f != f }
	return v1 == v2 || (isNaN(v1) && isNaN(v2))
}

// offByOne returns true if integers v1 and v2 differ by 1.
func offByOne(v1, v2 int) bool {
	return v1 == v2+1 || v1 == v2-1
}

func TestEqualFunc(t *testing.T) {
	for _, test := range equalIntTests {
		if got := EqualFunc(test.s1, test.s2, equal[int]); got != test.want {
			t.Errorf("EqualFunc(%v, %v, equal[int]) = %t, want %t", test.s1, test.s2, got, test.want)
		}
	}
	for _, test := range equalFloatTests {
		if got := EqualFunc(test.s1, test.s2, equal[float64]); got != test.wantEqual {
			t.Errorf("Equal(%v, %v, equal[float64]) = %t, want %t", test.s1, test.s2, got, test.wantEqual)
		}
		if got := EqualFunc(test.s1, test.s2, equalNaN[float64]); got != test.wantEqualNaN {
			t.Errorf("Equal(%v, %v, equalNaN[float64]) = %t, want %t", test.s1, test.s2, got, test.wantEqualNaN)
		}
	}

	s1 := Slice[int]{1, 2, 3}
	s2 := Slice[int]{2, 3, 4}
	if EqualFunc(s1, s1, offByOne) {
		t.Errorf("EqualFunc(%v, %v, offByOne) = true, want false", s1, s1)
	}
	if !EqualFunc(s1, s2, offByOne) {
		t.Errorf("EqualFunc(%v, %v, offByOne) = false, want true", s1, s2)
	}

	s3 := Slice[string]{"a", "b", "c"}
	s4 := Slice[string]{"A", "B", "C"}
	if !EqualFunc(s3, s4, strings.EqualFold) {
		t.Errorf("EqualFunc(%v, %v, strings.EqualFold) = false, want true", s3, s4)
	}

	cmpIntString := func(v1 int, v2 string) bool {
		return string(rune(v1)-1+'a') == v2
	}
	if !EqualFunc(s1, s3, cmpIntString) {
		t.Errorf("EqualFunc(%v, %v, cmpIntString) = false, want true", s1, s3)
	}
}

func BenchmarkEqualFunc_Large(b *testing.B) {
	type Large [4 * 1024]byte

	xs := make(Slice[Large], 1024)
	ys := make(Slice[Large], 1024)
	for i := 0; i < b.N; i++ {
		_ = EqualFunc(xs, ys, func(x, y Large) bool { return x == y })
	}
}

var compareIntTests = []struct {
	s1, s2 Slice[int]
	want   int
}{
	{
		Slice[int]{1},
		Slice[int]{1},
		0,
	},
	{
		Slice[int]{1},
		Slice[int]{},
		1,
	},
	{
		Slice[int]{},
		Slice[int]{1},
		-1,
	},
	{
		Slice[int]{},
		Slice[int]{},
		0,
	},
	{
		Slice[int]{1, 2, 3},
		Slice[int]{1, 2, 3},
		0,
	},
	{
		Slice[int]{1, 2, 3},
		Slice[int]{1, 2, 3, 4},
		-1,
	},
	{
		Slice[int]{1, 2, 3, 4},
		Slice[int]{1, 2, 3},
		+1,
	},
	{
		Slice[int]{1, 2, 3},
		Slice[int]{1, 4, 3},
		-1,
	},
	{
		Slice[int]{1, 4, 3},
		Slice[int]{1, 2, 3},
		+1,
	},
	{
		Slice[int]{1, 4, 3},
		Slice[int]{1, 2, 3, 8, 9},
		+1,
	},
}

var compareFloatTests = []struct {
	s1, s2 Slice[float64]
	want   int
}{
	{
		Slice[float64]{},
		Slice[float64]{},
		0,
	},
	{
		Slice[float64]{1},
		Slice[float64]{1},
		0,
	},
	{
		Slice[float64]{math.NaN()},
		Slice[float64]{math.NaN()},
		0,
	},
	{
		Slice[float64]{1, 2, math.NaN()},
		Slice[float64]{1, 2, math.NaN()},
		0,
	},
	{
		Slice[float64]{1, math.NaN(), 3},
		Slice[float64]{1, math.NaN(), 4},
		-1,
	},
	{
		Slice[float64]{1, math.NaN(), 3},
		Slice[float64]{1, 2, 4},
		-1,
	},
	{
		Slice[float64]{1, math.NaN(), 3},
		Slice[float64]{1, 2, math.NaN()},
		-1,
	},
	{
		Slice[float64]{1, 2, 3},
		Slice[float64]{1, 2, math.NaN()},
		+1,
	},
	{
		Slice[float64]{1, 2, 3},
		Slice[float64]{1, math.NaN(), 3},
		+1,
	},
	{
		Slice[float64]{1, math.NaN(), 3, 4},
		Slice[float64]{1, 2, math.NaN()},
		-1,
	},
}

func TestCompare(t *testing.T) {
	intWant := func(want bool) string {
		if want {
			return "0"
		}
		return "!= 0"
	}
	for _, test := range equalIntTests {
		if got := Compare(test.s1, test.s2); (got == 0) != test.want {
			t.Errorf("Compare(%v, %v) = %d, want %s", test.s1, test.s2, got, intWant(test.want))
		}
	}
	for _, test := range equalFloatTests {
		if got := Compare(test.s1, test.s2); (got == 0) != test.wantEqualNaN {
			t.Errorf("Compare(%v, %v) = %d, want %s", test.s1, test.s2, got, intWant(test.wantEqualNaN))
		}
	}

	for _, test := range compareIntTests {
		if got := Compare(test.s1, test.s2); got != test.want {
			t.Errorf("Compare(%v, %v) = %d, want %d", test.s1, test.s2, got, test.want)
		}
	}
	for _, test := range compareFloatTests {
		if got := Compare(test.s1, test.s2); got != test.want {
			t.Errorf("Compare(%v, %v) = %d, want %d", test.s1, test.s2, got, test.want)
		}
	}
}

func equalToCmp[T comparable](eq func(T, T) bool) func(T, T) int {
	return func(v1, v2 T) int {
		if eq(v1, v2) {
			return 0
		}
		return 1
	}
}

func TestCompareFunc(t *testing.T) {
	intWant := func(want bool) string {
		if want {
			return "0"
		}
		return "!= 0"
	}
	for _, test := range equalIntTests {
		if got := CompareFunc(test.s1, test.s2, equalToCmp(equal[int])); (got == 0) != test.want {
			t.Errorf("CompareFunc(%v, %v, equalToCmp(equal[int])) = %d, want %s", test.s1, test.s2, got, intWant(test.want))
		}
	}
	for _, test := range equalFloatTests {
		if got := CompareFunc(test.s1, test.s2, equalToCmp(equal[float64])); (got == 0) != test.wantEqual {
			t.Errorf("CompareFunc(%v, %v, equalToCmp(equal[float64])) = %d, want %s", test.s1, test.s2, got, intWant(test.wantEqual))
		}
	}

	for _, test := range compareIntTests {
		if got := CompareFunc(test.s1, test.s2, cmp.Compare[int]); got != test.want {
			t.Errorf("CompareFunc(%v, %v, cmp[int]) = %d, want %d", test.s1, test.s2, got, test.want)
		}
	}
	for _, test := range compareFloatTests {
		if got := CompareFunc(test.s1, test.s2, cmp.Compare[float64]); got != test.want {
			t.Errorf("CompareFunc(%v, %v, cmp[float64]) = %d, want %d", test.s1, test.s2, got, test.want)
		}
	}

	s1 := Slice[int]{1, 2, 3}
	s2 := Slice[int]{2, 3, 4}
	if got := CompareFunc(s1, s2, equalToCmp(offByOne)); got != 0 {
		t.Errorf("CompareFunc(%v, %v, offByOne) = %d, want 0", s1, s2, got)
	}

	s3 := Slice[string]{"a", "b", "c"}
	s4 := Slice[string]{"A", "B", "C"}
	if got := CompareFunc(s3, s4, strings.Compare); got != 1 {
		t.Errorf("CompareFunc(%v, %v, strings.Compare) = %d, want 1", s3, s4, got)
	}

	compareLower := func(v1, v2 string) int {
		return strings.Compare(strings.ToLower(v1), strings.ToLower(v2))
	}
	if got := CompareFunc(s3, s4, compareLower); got != 0 {
		t.Errorf("CompareFunc(%v, %v, compareLower) = %d, want 0", s3, s4, got)
	}

	cmpIntString := func(v1 int, v2 string) int {
		return strings.Compare(string(rune(v1)-1+'a'), v2)
	}
	if got := CompareFunc(s1, s3, cmpIntString); got != 0 {
		t.Errorf("CompareFunc(%v, %v, cmpIntString) = %d, want 0", s1, s3, got)
	}
}

var indexTests = []struct {
	s    Slice[int]
	v    int
	want int
}{
	{
		nil,
		0,
		-1,
	},
	{
		Slice[int]{},
		0,
		-1,
	},
	{
		Slice[int]{1, 2, 3},
		2,
		1,
	},
	{
		Slice[int]{1, 2, 2, 3},
		2,
		1,
	},
	{
		Slice[int]{1, 2, 3, 2},
		2,
		1,
	},
}

func TestIndex(t *testing.T) {
	for _, test := range indexTests {
		if got := Index(test.s, test.v); got != test.want {
			t.Errorf("Index(%v, %v) = %d, want %d", test.s, test.v, got, test.want)
		}
	}
}

func equalToIndex[T any](f func(T, T) bool, v1 T) func(T) bool {
	return func(v2 T) bool {
		return f(v1, v2)
	}
}

func BenchmarkIndex_Large(b *testing.B) {
	type Large [4 * 1024]byte

	ss := make(Slice[Large], 1024)
	for i := 0; i < b.N; i++ {
		_ = Index(ss, Large{1})
	}
}

func TestIndexFunc(t *testing.T) {
	for _, test := range indexTests {
		if got := IndexFunc(test.s, equalToIndex(equal[int], test.v)); got != test.want {
			t.Errorf("IndexFunc(%v, equalToIndex(equal[int], %v)) = %d, want %d", test.s, test.v, got, test.want)
		}
	}

	s1 := Slice[string]{"hi", "HI"}
	if got := IndexFunc(s1, equalToIndex(equal[string], "HI")); got != 1 {
		t.Errorf("IndexFunc(%v, equalToIndex(equal[string], %q)) = %d, want %d", s1, "HI", got, 1)
	}
	if got := IndexFunc(s1, equalToIndex(strings.EqualFold, "HI")); got != 0 {
		t.Errorf("IndexFunc(%v, equalToIndex(strings.EqualFold, %q)) = %d, want %d", s1, "HI", got, 0)
	}
}

func BenchmarkIndexFunc_Large(b *testing.B) {
	type Large [4 * 1024]byte

	ss := make(Slice[Large], 1024)
	for i := 0; i < b.N; i++ {
		_ = IndexFunc(ss, func(e Large) bool {
			return e == Large{1}
		})
	}
}

func TestContains(t *testing.T) {
	for _, test := range indexTests {
		if got := Contains(test.s, test.v); got != (test.want != -1) {
			t.Errorf("Contains(%v, %v) = %t, want %t", test.s, test.v, got, test.want != -1)
		}
	}
}

func TestContainsFunc(t *testing.T) {
	for _, test := range indexTests {
		if got := ContainsFunc(test.s, equalToIndex(equal[int], test.v)); got != (test.want != -1) {
			t.Errorf("ContainsFunc(%v, equalToIndex(equal[int], %v)) = %t, want %t", test.s, test.v, got, test.want != -1)
		}
	}

	s1 := Slice[string]{"hi", "HI"}
	if got := ContainsFunc(s1, equalToIndex(equal[string], "HI")); got != true {
		t.Errorf("ContainsFunc(%v, equalToContains(equal[string], %q)) = %t, want %t", s1, "HI", got, true)
	}
	if got := ContainsFunc(s1, equalToIndex(equal[string], "hI")); got != false {
		t.Errorf("ContainsFunc(%v, equalToContains(strings.EqualFold, %q)) = %t, want %t", s1, "hI", got, false)
	}
	if got := ContainsFunc(s1, equalToIndex(strings.EqualFold, "hI")); got != true {
		t.Errorf("ContainsFunc(%v, equalToContains(strings.EqualFold, %q)) = %t, want %t", s1, "hI", got, true)
	}
}

var insertTests = []struct {
	s    Slice[int]
	i    int
	add  Slice[int]
	want Slice[int]
}{
	{
		Slice[int]{1, 2, 3},
		0,
		Slice[int]{4},
		Slice[int]{4, 1, 2, 3},
	},
	{
		Slice[int]{1, 2, 3},
		1,
		Slice[int]{4},
		Slice[int]{1, 4, 2, 3},
	},
	{
		Slice[int]{1, 2, 3},
		3,
		Slice[int]{4},
		Slice[int]{1, 2, 3, 4},
	},
	{
		Slice[int]{1, 2, 3},
		2,
		Slice[int]{4, 5},
		Slice[int]{1, 2, 4, 5, 3},
	},
}

// func TestInsert(t *testing.T) {
// 	s := Slice[int]{1, 2, 3}
// 	if got := Insert(s, 0); !Equal(got, s) {
// 		t.Errorf("Insert(%v, 0) = %v, want %v", s, got, s)
// 	}
// 	for _, test := range insertTests {
// 		copy := Clone(test.s)
// 		if got := Insert(copy, test.i, test.add...); !Equal(got, test.want) {
// 			t.Errorf("Insert(%v, %d, %v...) = %v, want %v", test.s, test.i, test.add, got, test.want)
// 		}
// 	}

// 	if !testenv.OptimizationOff() && !race.Enabled {
// 		// Allocations should be amortized.
// 		const count = 50
// 		n := testing.AllocsPerRun(10, func() {
// 			s := Slice[int]{1, 2, 3}
// 			for i := 0; i < count; i++ {
// 				s = Insert(s, 0, 1)
// 			}
// 		})
// 		if n > count/2 {
// 			t.Errorf("too many allocations inserting %d elements: got %v, want less than %d", count, n, count/2)
// 		}
// 	}
// }

func TestInsertOverlap(t *testing.T) {
	const N = 10
	a := make(Slice[int], N)
	want := make(Slice[int], 2*N)
	for n := 0; n <= N; n++ { // length
		for i := 0; i <= n; i++ { // insertion point
			for x := 0; x <= N; x++ { // start of inserted data
				for y := x; y <= N; y++ { // end of inserted data
					for k := 0; k < N; k++ {
						a[k] = k
					}
					want = want[:0]
					want = append(want, a[:i]...)
					want = append(want, a[x:y]...)
					want = append(want, a[i:n]...)
					got := Insert(a[:n], i, a[x:y]...)
					if !Equal(got, want) {
						t.Errorf("Insert with overlap failed n=%d i=%d x=%d y=%d, got %v want %v", n, i, x, y, got, want)
					}
				}
			}
		}
	}
}

func TestInsertPanics(t *testing.T) {
	a := [3]int{}
	b := [1]int{}
	for _, test := range []struct {
		name string
		s    Slice[int]
		i    int
		v    Slice[int]
	}{
		// There are no values.
		{"with negative index", a[:1:1], -1, nil},
		{"with out-of-bounds index and > cap", a[:1:1], 2, nil},
		{"with out-of-bounds index and = cap", a[:1:2], 2, nil},
		{"with out-of-bounds index and < cap", a[:1:3], 2, nil},

		// There are values.
		{"with negative index", a[:1:1], -1, b[:]},
		{"with out-of-bounds index and > cap", a[:1:1], 2, b[:]},
		{"with out-of-bounds index and = cap", a[:1:2], 2, b[:]},
		{"with out-of-bounds index and < cap", a[:1:3], 2, b[:]},
	} {
		if !panics(func() { _ = Insert(test.s, test.i, test.v...) }) {
			t.Errorf("Insert %s: got no panic, want panic", test.name)
		}
	}
}

var deleteTests = []struct {
	s    Slice[int]
	i, j int
	want Slice[int]
}{
	{
		Slice[int]{1, 2, 3},
		0,
		0,
		Slice[int]{1, 2, 3},
	},
	{
		Slice[int]{1, 2, 3},
		0,
		1,
		Slice[int]{2, 3},
	},
	{
		Slice[int]{1, 2, 3},
		3,
		3,
		Slice[int]{1, 2, 3},
	},
	{
		Slice[int]{1, 2, 3},
		0,
		2,
		Slice[int]{3},
	},
	{
		Slice[int]{1, 2, 3},
		0,
		3,
		Slice[int]{},
	},
}

func TestDelete(t *testing.T) {
	for _, test := range deleteTests {
		copy := Clone(test.s)
		if got := Delete(copy, test.i, test.j); !Equal(got, test.want) {
			t.Errorf("Delete(%v, %d, %d) = %v, want %v", test.s, test.i, test.j, got, test.want)
		}
	}
}

var deleteFuncTests = []struct {
	s    Slice[int]
	fn   func(int) bool
	want Slice[int]
}{
	{
		nil,
		func(int) bool { return true },
		nil,
	},
	{
		Slice[int]{1, 2, 3},
		func(int) bool { return true },
		nil,
	},
	{
		Slice[int]{1, 2, 3},
		func(int) bool { return false },
		Slice[int]{1, 2, 3},
	},
	{
		Slice[int]{1, 2, 3},
		func(i int) bool { return i > 2 },
		Slice[int]{1, 2},
	},
	{
		Slice[int]{1, 2, 3},
		func(i int) bool { return i < 2 },
		Slice[int]{2, 3},
	},
	{
		Slice[int]{10, 2, 30},
		func(i int) bool { return i >= 10 },
		Slice[int]{2},
	},
}

func TestDeleteFunc(t *testing.T) {
	for i, test := range deleteFuncTests {
		copy := Clone(test.s)
		if got := DeleteFunc(copy, test.fn); !Equal(got, test.want) {
			t.Errorf("DeleteFunc case %d: got %v, want %v", i, got, test.want)
		}
	}
}

func panics(f func()) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			b = true
		}
	}()
	f()
	return false
}

func TestDeletePanics(t *testing.T) {
	s := Slice[int]{0, 1, 2, 3, 4}
	s = s[0:2]
	_ = s[0:4] // this is a valid slice of s

	for _, test := range []struct {
		name string
		s    Slice[int]
		i, j int
	}{
		{"with negative first index", Slice[int]{42}, -2, 1},
		{"with negative second index", Slice[int]{42}, 1, -1},
		{"with out-of-bounds first index", Slice[int]{42}, 2, 3},
		{"with out-of-bounds second index", Slice[int]{42}, 0, 2},
		{"with out-of-bounds both indexes", Slice[int]{42}, 2, 2},
		{"with invalid i>j", Slice[int]{42}, 1, 0},
		{"s[i:j] is valid and j > len(s)", s, 0, 4},
		{"s[i:j] is valid and i == j > len(s)", s, 3, 3},
	} {
		if !panics(func() { _ = Delete(test.s, test.i, test.j) }) {
			t.Errorf("Delete %s: got no panic, want panic", test.name)
		}
	}
}

func TestDeleteClearTail(t *testing.T) {
	mem := Slice[*int]{new(int), new(int), new(int), new(int), new(int), new(int)}
	s := mem[0:5] // there is 1 element beyond len(s), within cap(s)

	s = Delete(s, 2, 4)

	if mem[3] != nil || mem[4] != nil {
		// Check that potential memory leak is avoided
		t.Errorf("Delete: want nil discarded elements, got %v, %v", mem[3], mem[4])
	}
	if mem[5] == nil {
		t.Errorf("Delete: want unchanged elements beyond original len, got nil")
	}
}

func TestDeleteFuncClearTail(t *testing.T) {
	mem := Slice[*int]{new(int), new(int), new(int), new(int), new(int), new(int)}
	*mem[2], *mem[3] = 42, 42
	s := mem[0:5] // there is 1 element beyond len(s), within cap(s)

	s = DeleteFunc(s, func(i *int) bool {
		return i != nil && *i == 42
	})

	if mem[3] != nil || mem[4] != nil {
		// Check that potential memory leak is avoided
		t.Errorf("DeleteFunc: want nil discarded elements, got %v, %v", mem[3], mem[4])
	}
	if mem[5] == nil {
		t.Errorf("DeleteFunc: want unchanged elements beyond original len, got nil")
	}
}

func TestClone(t *testing.T) {
	s1 := Slice[int]{1, 2, 3}
	s2 := Clone(s1)
	if !Equal(s1, s2) {
		t.Errorf("Clone(%v) = %v, want %v", s1, s2, s1)
	}
	s1[0] = 4
	want := Slice[int]{1, 2, 3}
	if !Equal(s2, want) {
		t.Errorf("Clone(%v) changed unexpectedly to %v", want, s2)
	}
	if got := Clone(Slice[int](nil)); got != nil {
		t.Errorf("Clone(nil) = %#v, want nil", got)
	}
	if got := Clone(s1[:0]); got == nil || len(got) != 0 {
		t.Errorf("Clone(%v) = %#v, want %#v", s1[:0], got, s1[:0])
	}
}

var compactTests = []struct {
	name string
	s    Slice[int]
	want Slice[int]
}{
	{
		"nil",
		nil,
		nil,
	},
	{
		"one",
		Slice[int]{1},
		Slice[int]{1},
	},
	{
		"sorted",
		Slice[int]{1, 2, 3},
		Slice[int]{1, 2, 3},
	},
	{
		"1 item",
		Slice[int]{1, 1, 2},
		Slice[int]{1, 2},
	},
	{
		"unsorted",
		Slice[int]{1, 2, 1},
		Slice[int]{1, 2, 1},
	},
	{
		"many",
		Slice[int]{1, 2, 2, 3, 3, 4},
		Slice[int]{1, 2, 3, 4},
	},
}

func TestCompact(t *testing.T) {
	for _, test := range compactTests {
		copy := Clone(test.s)
		if got := Compact(copy); !Equal(got, test.want) {
			t.Errorf("Compact(%v) = %v, want %v", test.s, got, test.want)
		}
	}
}

func BenchmarkCompact(b *testing.B) {
	for _, c := range compactTests {
		b.Run(c.name, func(b *testing.B) {
			ss := make(Slice[int], 0, 64)
			for k := 0; k < b.N; k++ {
				ss = ss[:0]
				ss = append(ss, c.s...)
				_ = Compact(ss)
			}
		})
	}
}

func BenchmarkCompact_Large(b *testing.B) {
	type Large [4 * 1024]byte

	ss := make(Slice[Large], 1024)
	for i := 0; i < b.N; i++ {
		_ = Compact(ss)
	}
}

func TestCompactFunc(t *testing.T) {
	for _, test := range compactTests {
		copy := Clone(test.s)
		if got := CompactFunc(copy, equal[int]); !Equal(got, test.want) {
			t.Errorf("CompactFunc(%v, equal[int]) = %v, want %v", test.s, got, test.want)
		}
	}

	s1 := Slice[string]{"a", "a", "A", "B", "b"}
	copy := Clone(s1)
	want := Slice[string]{"a", "B"}
	if got := CompactFunc(copy, strings.EqualFold); !Equal(got, want) {
		t.Errorf("CompactFunc(%v, strings.EqualFold) = %v, want %v", s1, got, want)
	}
}

func TestCompactClearTail(t *testing.T) {
	one, two, three, four := 1, 2, 3, 4
	mem := Slice[*int]{&one, &one, &two, &two, &three, &four}
	s := mem[0:5] // there is 1 element beyond len(s), within cap(s)
	copy := Clone(s)

	s = Compact(s)
	want := Slice[*int]{&one, &two, &three}
	if !Equal(s, want) {
		t.Errorf("Compact(%v) = %v, want %v", copy, s, want)
	}

	if mem[3] != nil || mem[4] != nil {
		// Check that potential memory leak is avoided
		t.Errorf("Compact: want nil discarded elements, got %v, %v", mem[3], mem[4])
	}
	if mem[5] != &four {
		t.Errorf("Compact: want unchanged element beyond original len, got %v", mem[5])
	}
}

func TestCompactFuncClearTail(t *testing.T) {
	a, b, c, d, e, f := 1, 1, 2, 2, 3, 4
	mem := Slice[*int]{&a, &b, &c, &d, &e, &f}
	s := mem[0:5] // there is 1 element beyond len(s), within cap(s)
	copy := Clone(s)

	s = CompactFunc(s, func(x, y *int) bool {
		if x == nil || y == nil {
			return x == y
		}
		return *x == *y
	})
	want := Slice[*int]{&a, &c, &e}
	if !Equal(s, want) {
		t.Errorf("CompactFunc(%v) = %v, want %v", copy, s, want)
	}

	if mem[3] != nil || mem[4] != nil {
		// Check that potential memory leak is avoided
		t.Errorf("CompactFunc: want nil discarded elements, got %v, %v", mem[3], mem[4])
	}
	if mem[5] != &f {
		t.Errorf("CompactFunc: want unchanged elements beyond original len, got %v", mem[5])
	}
}

func BenchmarkCompactFunc_Large(b *testing.B) {
	type Large [4 * 1024]byte

	ss := make(Slice[Large], 1024)
	for i := 0; i < b.N; i++ {
		_ = CompactFunc(ss, func(a, b Large) bool { return a == b })
	}
}

func TestGrow(t *testing.T) {
	s1 := Slice[int]{1, 2, 3}

	copy := Clone(s1)
	s2 := Grow(copy, 1000)
	if !Equal(s1, s2) {
		t.Errorf("Grow(%v) = %v, want %v", s1, s2, s1)
	}
	if cap(s2) < 1000+len(s1) {
		t.Errorf("after Grow(%v) cap = %d, want >= %d", s1, cap(s2), 1000+len(s1))
	}

	// Test mutation of elements between length and capacity.
	copy = Clone(s1)
	s3 := Grow(copy[:1], 2)[:3]
	if !Equal(s1, s3) {
		t.Errorf("Grow should not mutate elements between length and capacity")
	}
	s3 = Grow(copy[:1], 1000)[:3]
	if !Equal(s1, s3) {
		t.Errorf("Grow should not mutate elements between length and capacity")
	}

	// // Test number of allocations.
	// if n := testing.AllocsPerRun(100, func() { _ = Grow(s2, cap(s2)-len(s2)) }); n != 0 {
	// 	t.Errorf("Grow should not allocate when given sufficient capacity; allocated %v times", n)
	// }
	// if n := testing.AllocsPerRun(100, func() { _ = Grow(s2, cap(s2)-len(s2)+1) }); n != 1 {
	// 	errorf := t.Errorf
	// 	if race.Enabled || testenv.OptimizationOff() {
	// 		errorf = t.Logf // this allocates multiple times in race detector mode
	// 	}
	// 	errorf("Grow should allocate once when given insufficient capacity; allocated %v times", n)
	// }

	// Test for negative growth sizes.
	var gotPanic bool
	func() {
		defer func() { gotPanic = recover() != nil }()
		_ = Grow(s1, -1)
	}()
	if !gotPanic {
		t.Errorf("Grow(-1) did not panic; expected a panic")
	}
}

func TestClip(t *testing.T) {
	s1 := Slice[int]{1, 2, 3, 4, 5, 6}[:3]
	orig := Clone(s1)
	if len(s1) != 3 {
		t.Errorf("len(%v) = %d, want 3", s1, len(s1))
	}
	if cap(s1) < 6 {
		t.Errorf("cap(%v[:3]) = %d, want >= 6", orig, cap(s1))
	}
	s2 := Clip(s1)
	if !Equal(s1, s2) {
		t.Errorf("Clip(%v) = %v, want %v", s1, s2, s1)
	}
	if cap(s2) != 3 {
		t.Errorf("cap(Clip(%v)) = %d, want 3", orig, cap(s2))
	}
}

func TestReverse(t *testing.T) {
	even := Slice[int]{3, 1, 4, 1, 5, 9} // len = 6
	Reverse(even)
	want := Slice[int]{9, 5, 1, 4, 1, 3}
	if !Equal(even, want) {
		t.Errorf("Reverse(even) = %v, want %v", even, want)
	}

	odd := Slice[int]{3, 1, 4, 1, 5, 9, 2} // len = 7
	Reverse(odd)
	want = Slice[int]{2, 9, 5, 1, 4, 1, 3}
	if !Equal(odd, want) {
		t.Errorf("Reverse(odd) = %v, want %v", odd, want)
	}

	words := Slice[string](strings.Fields("one two three"))
	Reverse(words)
	if want := strings.Fields("three two one"); !Equal(words, want) {
		t.Errorf("Reverse(words) = %v, want %v", words, want)
	}

	singleton := Slice[string]{"one"}
	Reverse(singleton)
	want2 := Slice[string]{"one"}
	if !Equal(singleton, want2) {
		t.Errorf("Reverse(singeleton) = %v, want %v", singleton, want)
	}

	Reverse[Slice[string]](nil)
}

// naiveReplace is a baseline implementation to the Replace function.
func naiveReplace[S Slice[E], E any](s S, i, j int, v ...E) S {
	s = Delete(s, i, j)
	s = Insert(s, i, v...)
	return s
}

func TestReplace(t *testing.T) {
	for _, test := range []struct {
		s, v Slice[int]
		i, j int
	}{
		{}, // all zero value
		{
			s: Slice[int]{1, 2, 3, 4},
			v: Slice[int]{5},
			i: 1,
			j: 2,
		},
		{
			s: Slice[int]{1, 2, 3, 4},
			v: Slice[int]{5, 6, 7, 8},
			i: 1,
			j: 2,
		},
		{
			s: func() Slice[int] {
				s := make(Slice[int], 3, 20)
				s[0] = 0
				s[1] = 1
				s[2] = 2
				return s
			}(),
			v: Slice[int]{3, 4, 5, 6, 7},
			i: 0,
			j: 1,
		},
	} {
		ss, vv := Clone(test.s), Clone(test.v)
		want := naiveReplace(ss, test.i, test.j, vv...)
		got := Replace(test.s, test.i, test.j, test.v...)
		if !Equal(got, want) {
			t.Errorf("Replace(%v, %v, %v, %v) = %v, want %v", test.s, test.i, test.j, test.v, got, want)
		}
	}
}

func TestReplacePanics(t *testing.T) {
	s := Slice[int]{0, 1, 2, 3, 4}
	s = s[0:2]
	_ = s[0:4] // this is a valid slice of s

	for _, test := range []struct {
		name string
		s, v Slice[int]
		i, j int
	}{
		{"indexes out of order", Slice[int]{1, 2}, Slice[int]{3}, 2, 1},
		{"large index", Slice[int]{1, 2}, Slice[int]{3}, 1, 10},
		{"negative index", Slice[int]{1, 2}, Slice[int]{3}, -1, 2},
		{"s[i:j] is valid and j > len(s)", s, nil, 0, 4},
	} {
		ss, vv := Clone(test.s), Clone(test.v)
		if !panics(func() { _ = Replace(ss, test.i, test.j, vv...) }) {
			t.Errorf("Replace %s: should have panicked", test.name)
		}
	}
}

func TestReplaceGrow(t *testing.T) {
	// When Replace needs to allocate a new slice, we want the original slice
	// to not be changed.
	a, b, c, d, e, f := 1, 2, 3, 4, 5, 6
	mem := Slice[*int]{&a, &b, &c, &d, &e, &f}
	memcopy := Clone(mem)
	s := mem[0:5] // there is 1 element beyond len(s), within cap(s)
	copy := Clone(s)
	original := s

	// The new elements don't fit within cap(s), so Replace will allocate.
	z := 99
	s = Replace(s, 1, 3, &z, &z, &z, &z)
	want := Slice[*int]{&a, &z, &z, &z, &z, &d, &e}

	if !Equal(s, want) {
		t.Errorf("Replace(%v, 1, 3, %v, %v, %v, %v) = %v, want %v", copy, &z, &z, &z, &z, s, want)
	}

	if !Equal(original, copy) {
		t.Errorf("original slice has changed, got %v, want %v", original, copy)
	}

	if !Equal(mem, memcopy) {
		// Changing the original tail s[len(s):cap(s)] is unwanted
		t.Errorf("original backing memory has changed, got %v, want %v", mem, memcopy)
	}
}

func TestReplaceClearTail(t *testing.T) {
	a, b, c, d, e, f := 1, 2, 3, 4, 5, 6
	mem := Slice[*int]{&a, &b, &c, &d, &e, &f}
	s := mem[0:5] // there is 1 element beyond len(s), within cap(s)
	copy := Clone(s)

	y, z := 8, 9
	s = Replace(s, 1, 4, &y, &z)
	want := Slice[*int]{&a, &y, &z, &e}
	if !Equal(s, want) {
		t.Errorf("Replace(%v) = %v, want %v", copy, s, want)
	}

	if mem[4] != nil {
		// Check that potential memory leak is avoided
		t.Errorf("Replace: want nil discarded element, got %v", mem[4])
	}
	if mem[5] != &f {
		t.Errorf("Replace: want unchanged elements beyond original len, got %v", mem[5])
	}
}

func TestReplaceOverlap(t *testing.T) {
	const N = 10
	a := make(Slice[int], N)
	want := make(Slice[int], 2*N)
	for n := 0; n <= N; n++ { // length
		for i := 0; i <= n; i++ { // insertion point 1
			for j := i; j <= n; j++ { // insertion point 2
				for x := 0; x <= N; x++ { // start of inserted data
					for y := x; y <= N; y++ { // end of inserted data
						for k := 0; k < N; k++ {
							a[k] = k
						}
						want = want[:0]
						want = append(want, a[:i]...)
						want = append(want, a[x:y]...)
						want = append(want, a[j:n]...)
						got := Replace(a[:n], i, j, a[x:y]...)
						if !Equal(got, want) {
							t.Errorf("Insert with overlap failed n=%d i=%d j=%d x=%d y=%d, got %v want %v", n, i, j, x, y, got, want)
						}
					}
				}
			}
		}
	}
}

func BenchmarkReplace(b *testing.B) {
	cases := []struct {
		name string
		s, v func() Slice[int]
		i, j int
	}{
		{
			name: "fast",
			s: func() Slice[int] {
				return make(Slice[int], 100)
			},
			v: func() Slice[int] {
				return make(Slice[int], 20)
			},
			i: 10,
			j: 40,
		},
		{
			name: "slow",
			s: func() Slice[int] {
				return make(Slice[int], 100)
			},
			v: func() Slice[int] {
				return make(Slice[int], 20)
			},
			i: 0,
			j: 2,
		},
	}

	for _, c := range cases {
		b.Run("naive-"+c.name, func(b *testing.B) {
			for k := 0; k < b.N; k++ {
				s := c.s()
				v := c.v()
				_ = naiveReplace(s, c.i, c.j, v...)
			}
		})
		b.Run("optimized-"+c.name, func(b *testing.B) {
			for k := 0; k < b.N; k++ {
				s := c.s()
				v := c.v()
				_ = Replace(s, c.i, c.j, v...)
			}
		})
	}

}

func TestInsertGrowthRate(t *testing.T) {
	b := make(Slice[byte], 1)
	maxCap := cap(b)
	nGrow := 0
	const N = 1e6
	for i := 0; i < N; i++ {
		b = Insert(b, len(b)-1, 0)
		if cap(b) > maxCap {
			maxCap = cap(b)
			nGrow++
		}
	}
	want := int(math.Log(N) / math.Log(1.25)) // 1.25 == growth rate for large slices
	if nGrow > want {
		t.Errorf("too many grows. got:%d want:%d", nGrow, want)
	}
}

func TestReplaceGrowthRate(t *testing.T) {
	b := make(Slice[byte], 2)
	maxCap := cap(b)
	nGrow := 0
	const N = 1e6
	for i := 0; i < N; i++ {
		b = Replace(b, len(b)-2, len(b)-1, 0, 0)
		if cap(b) > maxCap {
			maxCap = cap(b)
			nGrow++
		}
	}
	want := int(math.Log(N) / math.Log(1.25)) // 1.25 == growth rate for large slices
	if nGrow > want {
		t.Errorf("too many grows. got:%d want:%d", nGrow, want)
	}
}

// func apply[T any](v Slice[T], f func(T) Slice[T]) {
// 	f(v)
// }

// // Test type inference with a named slice type.
// func TestInference(t *testing.T) {
// 	s1 := []int{1, 2, 3}
// 	apply(s1, Reverse)
// 	want := []int{3, 2, 1}
// 	if !Equal(s1, want) {
// 		t.Errorf("Reverse(%v) = %v, want %v", []int{1, 2, 3}, s1, want)
// 	}

// 	type S []int
// 	s2 := S{4, 5, 6}
// 	apply(s2, Reverse)
// 	if want := (S{6, 5, 4}); !Equal(s2, want) {
// 		t.Errorf("Reverse(%v) = %v, want %v", S{4, 5, 6}, s2, want)
// 	}
// }

// func TestConcat(t *testing.T) {
// 	cases := []struct {
// 		s    Slice[Slice[int]]
// 		want Slice[int]
// 	}{
// 		{
// 			s:    Slice[Slice[int]]{nil},
// 			want: nil,
// 		},
// 		{
// 			s:    Slice[Slice[int]]{{1}},
// 			want: Slice[int]{1},
// 		},
// 		{
// 			s:    Slice[Slice[int]]{{1}, {2}},
// 			want: Slice[int]{1, 2},
// 		},
// 		{
// 			s:    Slice[Slice[int]]{{1}, nil, {2}},
// 			want: Slice[int]{1, 2},
// 		},
// 	}
// 	for _, tc := range cases {
// 		got := Concat(tc.s...)
// 		if !Equal(tc.want, got) {
// 			t.Errorf("Concat(%v) = %v, want %v", tc.s, got, tc.want)
// 		}
// 		var sink Slice[int]
// 		allocs := testing.AllocsPerRun(5, func() {
// 			sink = Concat(tc.s...)
// 		})
// 		_ = sink
// 		if allocs > 1 {
// 			errorf := t.Errorf
// 			if testenv.OptimizationOff() || race.Enabled {
// 				errorf = t.Logf
// 			}
// 			errorf("Concat(%v) allocated %v times; want 1", tc.s, allocs)
// 		}
// 	}
// }

// func TestConcat_too_large(t *testing.T) {
// 	// Use zero length element to minimize memory in testing
// 	type void struct{}
// 	cases := []struct {
// 		lengths     Slice[int]
// 		shouldPanic bool
// 	}{
// 		{
// 			lengths:     Slice[int]{0, 0},
// 			shouldPanic: false,
// 		},
// 		{
// 			lengths:     Slice[int]{math.MaxInt, 0},
// 			shouldPanic: false,
// 		},
// 		{
// 			lengths:     Slice[int]{0, math.MaxInt},
// 			shouldPanic: false,
// 		},
// 		{
// 			lengths:     Slice[int]{math.MaxInt - 1, 1},
// 			shouldPanic: false,
// 		},
// 		{
// 			lengths:     Slice[int]{math.MaxInt - 1, 1, 1},
// 			shouldPanic: true,
// 		},
// 		{
// 			lengths:     Slice[int]{math.MaxInt, 1},
// 			shouldPanic: true,
// 		},
// 		{
// 			lengths:     Slice[int]{math.MaxInt, math.MaxInt},
// 			shouldPanic: true,
// 		},
// 	}
// 	for _, tc := range cases {
// 		var r any
// 		ss := make([][]void, 0, len(tc.lengths))
// 		for _, l := range tc.lengths {
// 			s := make([]void, l)
// 			ss = append(ss, s)
// 		}
// 		func() {
// 			defer func() {
// 				r = recover()
// 			}()
// 			_ = Concat(ss...)
// 		}()
// 		if didPanic := r != nil; didPanic != tc.shouldPanic {
// 			t.Errorf("slices.Concat(lens(%v)) got panic == %v",
// 				tc.lengths, didPanic)
// 		}
// 	}
// }
