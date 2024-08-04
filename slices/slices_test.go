package slices_test

import (
	"fmt"
	"reflect"
	"testing"
	"unicode"

	. "github.com/cramanan/go-types/slices"
)

func eq[T comparable](a, b Slice[T]) bool {
	return reflect.DeepEqual(a, b)
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
		{"Prepend Empty to Empty", Slice[rune]{}.Prepend(), nil},
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
	want := true
	if got != want {
		t.Errorf("slice.Some(some) = %t, want %t", got, want)
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

func TestSlice_IndexFunc(t *testing.T) {
	type users struct {
		id   int
		name string
	}
	user1 := users{0, "John"}
	user2 := users{1, "Jane"}
	cmpFunc := func(u users) bool {
		return u.id == 0
	}

	slice := New(user1, user2)

	got := slice.IndexFunc(cmpFunc)
	want := 0
	if got != want {
		t.Errorf("slice.IndexFunc(cmpFunc) = %d, want %d", got, want)
	}
}

func TestSlice_ContainsFunc(t *testing.T) {
	type users struct {
		id   int
		name string
	}
	user1 := users{0, "John"}
	user2 := users{1, "Jane"}
	hasJohn := func(u users) bool { return u.name == "John" }

	slice := New(user1, user2)

	got := slice.ContainsFunc(hasJohn)
	want := true
	if got != want {
		t.Errorf("slice.IndexFunc(hasJohn) = %t, want %t", got, want)
	}
}

func TestMap(t *testing.T) {
	type sameType[T any] struct {
		desc string
		got  Slice[T]
		want Slice[T]
	}

	slice := New(1, 2, 3, 4)
	doubled := func(i int) int { return i * 2 }
	squared := func(i int) int { return i * i }
	nothing := func(i int) int { return i }

	testCases := []sameType[int]{
		{"Doubled", Map(slice, doubled), New(2, 4, 6, 8)},
		{"Squared", Map(slice, squared), New(1, 4, 9, 16)},
		{"nothing", Map(slice, nothing), slice},
		{"nothing", Map(slice, nothing), slice},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if !eq(tC.got, tC.want) {
				t.Errorf("got %v, want %v", tC.got, tC.want)
			}
		})
	}

	slice2 := New(1.1, 2.2, 3.3, 4.4)
	sprint := func(i float64) string { return fmt.Sprint(i) }
	fill0 := func(i float64) string { return "0" }

	testCases2 := []sameType[string]{
		{"Sprint", Map(slice2, sprint), New("1.1", "2.2", "3.3", "4.4")},
		{"Fill 0", Map(slice2, fill0), New("0", "0", "0", "0")},
	}

	for _, tC := range testCases2 {
		t.Run(tC.desc, func(t *testing.T) {
			if !eq(tC.got, tC.want) {
				t.Errorf("got %v, want %v", tC.got, tC.want)
			}
		})
	}
}

func TestReduceSum(t *testing.T) {
	numbers := Slice[int]{1, 2, 3, 4, 5}
	sum := func(acc int, current int) int { return acc + current }
	got := Reduce(numbers, sum, 0)
	if got != 15 {
		t.Errorf("expected sum to be 15, got %d", got)
	}
}

func TestReduceConcat(t *testing.T) {
	strings := Slice[string]{"hello", "world", "go"}
	concatenated := Reduce(strings, func(acc string, current string) string {
		return acc + " " + current
	}, "")
	if concatenated != " hello world go" {
		t.Errorf("expected concatenated string to be ' hello world go', got '%s'", concatenated)
	}
}

func TestReduceEmptySlice(t *testing.T) {
	var numbers Slice[int]
	sum := Reduce(numbers, func(acc int, current int) int {
		return acc + current
	}, 0)
	if sum != 0 {
		t.Errorf("expected sum to be 0 for empty slice, got %d", sum)
	}
}

func TestReduceSingleElementSlice(t *testing.T) {
	numbers := Slice[int]{5}
	sum := Reduce(numbers, func(acc int, current int) int {
		return acc + current
	}, 0)
	if sum != 5 {
		t.Errorf("expected sum to be 5 for single-element slice, got %d", sum)
	}
}

func TestReduceInitialValue(t *testing.T) {
	numbers := Slice[int]{1, 2, 3}
	sum := Reduce(numbers, func(acc int, current int) int {
		return acc + current
	}, 10)
	if sum != 16 {
		t.Errorf("expected sum to be 16 with initial value 10, got %d", sum)
	}
}

func TestReducePanicsOnNilCallback(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected Reduce to panic on nil callback")
		}
	}()
	numbers := Slice[int]{1, 2, 3}
	Reduce(numbers, nil, 0)
}
