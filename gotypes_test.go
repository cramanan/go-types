package gotypes

import (
	"testing"

	"github.com/cramanan/go-types/functions"
	"github.com/cramanan/go-types/slices"
)

func eq[T comparable](a, b slices.Slice[T]) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestSlicesWithFunctions(t *testing.T) {

	type TestCase[T any] struct {
		desc      string
		got, want slices.Slice[T]
	}

	testCases := []TestCase[int]{
		{"Nothing", slices.New[int]().SortFunc(functions.Ascending), slices.New[int]()},
		{"Nothing to sort", slices.New(1).SortFunc(functions.Ascending), slices.New(1)},
		{"Sort Ascending",
			slices.New(9999999999999999, 10, 10000000, 0).SortFunc(functions.Ascending),
			slices.New(0, 10, 10000000, 9999999999999999)},
		{"Sort Descending",
			slices.New(9999999999999999, 10, 10000000, 0).SortFunc(functions.Descending),
			slices.New(0, 10, 10000000, 9999999999999999).Reverse()},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if !eq(tC.got, tC.want) {
				t.Errorf("%s got %v, want %v", tC.desc, tC.got, tC.want)
			}
		})
	}
}
