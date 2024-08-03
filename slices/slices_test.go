package slices_test

import (
	"testing"
	"unicode"

	. "github.com/cramanan/types/slices"
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
		{"Append Empty to Slice", Slice[float64]{0.0}.Append(), Slice[float64]{0.0}},
		{"Append Variadic to Empty", Slice[float64]{}.Append(), []float64{}},
		{"Append Variadic to Slice", Slice[float64]{0.0, 1.1}.Append(), Slice[float64]{0.0, 1.1}},
		{"Append []float64... to slice", Slice[float64]{69}.Append([]float64{420.0}...), []float64{69, 420}},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if !eq(tC.got, tC.want) {
				t.Errorf("got %v, wanted %v", tC.got, tC.want)
			}
		})
	}
}

var alphabet = Slice[rune]("abcdefghijklmnopqrstuvwxyz")

func TestSlice_At(t *testing.T) {
	type TestCase[T comparable] struct {
		desc string
		got  T
		want T
	}

	testCases := []TestCase[rune]{
		{"At First", alphabet.At(0), 'a'},
		{"At Last", alphabet.At(-1), 'z'},
		{"At Middle", alphabet.At(len(alphabet) / 2), 'n'},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.got != tC.want {
				t.Errorf("got %q, want %q", tC.got, tC.want)
			}
		})
	}
}

func TestSlice_AtOutOfRangePanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("alphabet.At(100) should have panicked but didn't.")
		}
	}()

	alphabet.At(100)
}

func TestSlice_Length(t *testing.T) {
	got := alphabet.Length()
	want := len(alphabet)
	if got != want && got != 26 {
		t.Errorf("alphabet.Length = %d, want %d", got, want)
	}
}

func TestSlice_ForEach(t *testing.T) {
	recipient := New[rune]()
	toUppercase := func(value rune, _ int) {
		recipient = append(recipient, unicode.ToUpper(value))
	}
	alphabet.ForEach(toUppercase)
	want := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	if string(recipient) != want {
		t.Errorf("recipient = %v, want %s", recipient, want)
	}
}

func TestSlice_Filter(t *testing.T) {
	slice := New(10, 0, 0, 20, 0, 30, 0, 40)
	removeZeros := func(element int, _ int) bool { return element != 0 }
	got := slice.Filter(removeZeros)
	want := New(10, 20, 30, 40)
	if !eq(got, want) {
		t.Errorf("slice.Filter(removeZeros) = %v, want %v", got, want)
	}
}

func TestSlice_Some(t *testing.T) {
	slice := New(false, false, true, false, false)
	some := func(element bool, _ int) bool { return element }
	got := slice.Some(some)
	if !got {
		t.Errorf("slice.Some(some) = %t, want %t", got, true)
	}

}

func TestSlice_Every(t *testing.T) {
	slice := New(0, 1, 2, 3, 4, 5, 6)
	under10 := func(element int, _ int) bool { return element < 10 }
	got := slice.Every(under10)
	if !got {
		t.Errorf("slice.Every(under10) = %t, want %t", got, true)
	}
}
