package ordered_test

import (
	"reflect"
	"testing"

	"golang.org/x/exp/constraints"

	. "github.com/cramanan/go-types/slices/ordered"
)

func eq[T any](s1, s2 []T) bool {
	return reflect.DeepEqual(s1, s2)
}

func TestNew(t *testing.T) {
	type TestCase[T constraints.Ordered] struct {
		desc      string
		got, want Ordered[T]
	}
	testCases := []TestCase[int]{
		{"Empty", New[int](), nil},
		{"Empty", make(Ordered[int], 0), Ordered[int]{}},
		{"Populated", make(Ordered[int], 10), Ordered[int]{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{"Populated", New(1, 2, 3, 4, 5, 6), Ordered[int]{1, 2, 3, 4, 5, 6}},

		{"Empty Spread", New([]int{}...), nil},
		{"Spreaded", New([]int{1, 2, 3, 4, 5, 6}...), Ordered[int]{1, 2, 3, 4, 5, 6}},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if !eq(tC.got, tC.want) {
				t.Errorf("%s = %v, want %v", tC.desc, tC.got, tC.want)
			}
		})
	}
}

func TestFrom(t *testing.T) {
	type TestCase[T constraints.Ordered] struct {
		desc      string
		got, want Ordered[T]
	}
	testCases := []TestCase[float64]{
		{"Empty", From[float64](nil), nil},
		{"Empty", From([]float64{}), Ordered[float64]{}},

		{"Populated", From([]float64{1.0, 2.0, 3.0, 4.0}), Ordered[float64]{1, 2, 3, 4}},
		{"Populated", From(New[float64](1., 2, 3, 4, 5, 6)), Ordered[float64]{1, 2, 3, 4, 5, 6}},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if !eq(tC.got, tC.want) {
				t.Errorf("%s = %v, want %v", tC.desc, tC.got, tC.want)
			}
		})
	}
}
