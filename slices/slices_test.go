package slices_test

import (
	"testing"

	. "github.com/cramanan/go-types/slices"
)

func eq[T comparable](a, b Slice[T]) bool {
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

func TestNew(t *testing.T) {
	type TestCase[T comparable] struct {
		desc string
		got  Slice[T]
		want Slice[T]
	}

	intTestCases := []TestCase[int]{
		{"New empty", New[int](), Slice[int]{}},
		{"New filled", New(1, 2, 3, 4), Slice[int]{1, 2, 3, 4}},
	}

	for _, tC := range intTestCases {
		t.Run(tC.desc, func(t *testing.T) {
			if !eq(tC.got, tC.want) {
				t.Errorf("got %v, wanted %v", tC.got, tC.want)

			}
		})
	}

	stringTestCases := []TestCase[string]{
		{"New empty", New[string](), Slice[string]{}},
		{"New filled", New("Hello", "World"), Slice[string]{"Hello", "World"}},
	}

	for _, tC := range stringTestCases {
		t.Run(tC.desc, func(t *testing.T) {
			if !eq(tC.got, tC.want) {
				t.Errorf("got %v, wanted %v", tC.got, tC.want)

			}
		})
	}
}

func TestFrom(t *testing.T) {
	type TestCase[T comparable] struct {
		desc string
		got  Slice[T]
		want Slice[T]
	}

	testCases := []TestCase[byte]{
		{"From", From([]byte("Hello")), Slice[byte]("Hello")},
		{"From Empty", From([]byte{}), Slice[byte]{}},
		{"From Slice", From(Slice[byte]{'H', 'e', 'l', 'l', 'o'}), Slice[byte]("Hello")},
		{"From nil", From[byte](nil), nil},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if !eq(tC.got, tC.want) {
				t.Errorf("got %v, wanted %v", tC.got, tC.want)

			}
		})
	}
}

func TestSlice_Prepend(t *testing.T) {
	type TestCase[T comparable] struct {
		desc string
		got  Slice[T]
		want Slice[T]
	}

	testCases := []TestCase[rune]{
		{"Prepend Empty to Empty", Slice[rune]{}.Prepend(), Slice[rune]{}},
		{"Prepend Empty to Slice", Slice[rune]("Nothing").Prepend(), Slice[rune]("Nothing")},
		{"Prepend Variadic to Empty", Slice[rune]{}.Prepend('H', 'E', 'L', 'L', 'O'), Slice[rune]("HELLO")},
		{"Prepend Variadic to Slice", Slice[rune]("World").Prepend('H', 'e', 'l', 'l', 'o', ' '), Slice[rune]("Hello World")},
		{"Prepend []rune... to Slice", Slice[rune]("xoxo").Prepend([]rune("xd lol ")...), []rune("xd lol xoxo")},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if !eq(tC.got, tC.want) {
				t.Errorf("got %v, wanted %v", tC.got, tC.want)
			}
		})
	}

}

func TestSlice_Append(t *testing.T) {
	type TestCase[T comparable] struct {
		desc string
		got  Slice[T]
		want Slice[T]
	}

	testCases := []TestCase[float64]{
		{"Append Empty to Empty", Slice[float64]{}.Append(), Slice[float64]{}},
		{"Append Empty to Slice", Slice[float64]{0.0}.Append(), Slice[float64]("Nothing")},
		{"Append Variadic to Empty", Slice[float64]{}.Append('W', 'o', 'r', 'L', 'O'), Slice[float64]("HELLO")},
		{"Append Variadic to Slice", Slice[float64]("Hello").Append('H', 'e', 'l', 'l', 'o', ' '), Slice[float64]("Hello World")},
		{"Append []rune... to Slice", Slice[float64]("xoxo").Append([]float64("xd lol ")...), []float64("xd lol xoxo")},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if !eq(tC.got, tC.want) {
				t.Errorf("got %v, wanted %v", tC.got, tC.want)
			}
		})
	}

}
