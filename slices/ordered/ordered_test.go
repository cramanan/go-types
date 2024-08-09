package ordered_test

import (
	"testing"

	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"

	. "github.com/cramanan/go-types/slices/ordered"
)

func eq[T ~[]A, A comparable](s1, s2 T) bool {
	return slices.Equal(s1, s2)
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
		{"Spread", New([]int{1, 2, 3, 4, 5, 6}...), Ordered[int]{1, 2, 3, 4, 5, 6}},
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

	type Alias []int

	testCases := []TestCase[int]{
		{"Empty", From[[]int](nil), nil},
		{"Empty", From[Ordered[int]](nil), nil},
		{"Populated", From(New(1, 22, 333, 4444, 5555)), []int{1, 22, 333, 4444, 5555}},
		{"Populated", From(Alias([]int{1, 22, 333, 4444, 5555})), Ordered[int]{1, 22, 333, 4444, 5555}},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if !eq(tC.got, tC.want) {
				t.Errorf("%s = %v, want %v", tC.desc, tC.got, tC.want)
			}
		})
	}
}

func TestFor_Each(t *testing.T) {
	o := New(1, 2, 3, 4, 5, 6)
	evens := *new(Ordered[int])
	even := func(value int, idx int) {
		if value%2 == 0 {
			evens = append(evens, value)
		}
	}
	o.ForEach(even)
	if want := []int{2, 4, 6}; !eq(evens, want) {
		t.Errorf("Error: %v != %v", evens, want)
	}
	odds := *new(Ordered[int])
	odd := func(value int, idx int) {
		if value%2 != 0 {
			odds = append(odds, value)
		}
	}
	o.ForEach(odd)
	if want := []int{1, 3, 5}; !eq(odds, want) {
		t.Errorf("Error: %v != %v", odds, want)
	}
}

func TestFilter(t *testing.T) {
	callback := func(value string, idx int) bool { return value != "Bob" }
	callbackFn := func(value string, idx int) bool { return value == "Bob" }

	testCases := []struct {
		desc      string
		got, want []string
	}{
		{"Remove Bobs", New("Alice", "Bob", "John", "Bob", "Bob", "Jane").Filter(callback), []string{"Alice", "John", "Jane"}},
		{"Only Bobs", New("Alice", "Bob", "John", "Bob", "Bob", "Jane").Filter(callbackFn), []string{"Bob", "Bob", "Bob"}},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if !eq(tC.got, tC.want) {
				t.Errorf("Error: %v != %v", tC.got, tC.want)
			}
		})
	}

	defer func() {
		if reason := recover(); reason != "callback function is nil" {
			t.Error("Should have panicked but didn't")
		}
	}()
	New(870987, 7697869876, 658678675).Every(nil)
}

func TestSome(t *testing.T) {
	callbackFn := func(value string, idx int) bool { return value == "Bob" }

	testCases := []struct {
		desc      string
		got, want bool
	}{
		{"No Empty", New[string]().Some(callbackFn), false},
		{"Have Bobs", New("Alice", "Bob", "John", "Bob", "Bob", "Jane").Some(callbackFn), true},
		{"No Bobs", New("Alice", "Alice", "John", "Alice", "Alice", "Alice").Some(callbackFn), false},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.got != tC.want {
				t.Errorf("Error: %v != %v", tC.got, tC.want)
			}
		})
	}

	defer func() {
		if reason := recover(); reason != "callback function is nil" {
			t.Error("Should have panicked but didn't")
		}
	}()
	New(870987, 7697869876, 658678675).Every(nil)
}

func TestEvery(t *testing.T) {
	callbackFn := func(value string, idx int) bool { return value == "Bob" }

	testCases := []struct {
		desc      string
		got, want bool
	}{
		{"Empty", New[string]().Every(callbackFn), true},
		{"Have Bobs", New("Alice", "Bob", "John", "Bob", "Bob", "Jane").Every(callbackFn), false},
		{"Only Bobs", New("Bob", "Bob", "Bob", "Bob", "Bob", "Bob").Every(callbackFn), true},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.got != tC.want {
				t.Errorf("Error: %v != %v", tC.got, tC.want)
			}
		})
	}

	defer func() {
		if reason := recover(); reason != "callback function is nil" {
			t.Error("Should have panicked but didn't")
		}
	}()
	New(870987, 7697869876, 658678675).Every(nil)
}

func doubled(value int, _ int) int { return 2 * value }

func TestMap(t *testing.T) {
	testCases := []struct {
		desc      string
		got, want []int
	}{
		{"Empty", New[int]().Map(doubled), nil},
		{"Map", New(100, 200, 300, 400).Map(doubled), New(100*2, 200*2, 300*2, 400*2)},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if !eq(tC.got, tC.want) {
				t.Errorf("Error: %v != %v", tC.got, tC.want)
			}
		})
	}

	defer func() {
		if reason := recover(); reason != "callback function is nil" {
			t.Error("Should have panicked but didn't")
		}
	}()
	New(870987, 7697869876, 658678675).Map(nil)
}
