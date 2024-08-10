package maps_test

import (
	"math"
	"strconv"
	"testing"

	. "github.com/cramanan/go-types/maps"
)

var m1 = Map[int, int]{1: 2, 2: 4, 4: 8, 8: 16}
var m2 = Map[int, string]{1: "2", 2: "4", 4: "8", 8: "16"}

func TestEqual(t *testing.T) {
	if !Equal(m1, m1) {
		t.Errorf("Equal(%v, %v) = false, want true", m1, m1)
	}
	if Equal(m1, (map[int]int)(nil)) {
		t.Errorf("Equal(%v, nil) = true, want false", m1)
	}
	if Equal((map[int]int)(nil), m1) {
		t.Errorf("Equal(nil, %v) = true, want false", m1)
	}
	if !Equal[map[int]int, map[int]int](nil, nil) {
		t.Error("Equal(nil, nil) = false, want true")
	}
	if ms := map[int]int{1: 2}; Equal(m1, ms) {
		t.Errorf("Equal(%v, %v) = true, want false", m1, ms)
	}

	// Comparing NaN for equality is expected to fail.
	mf := map[int]float64{1: 0, 2: math.NaN()}
	if Equal(mf, mf) {
		t.Errorf("Equal(%v, %v) = true, want false", mf, mf)
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

// equalStr compares ints and strings.
func equalIntStr(v1 int, v2 string) bool {
	return strconv.Itoa(v1) == v2
}

func TestEqualFunc(t *testing.T) {
	if !EqualFunc(m1, m1, equal[int]) {
		t.Errorf("EqualFunc(%v, %v, equal) = false, want true", m1, m1)
	}
	if EqualFunc(m1, (map[int]int)(nil), equal[int]) {
		t.Errorf("EqualFunc(%v, nil, equal) = true, want false", m1)
	}
	if EqualFunc((map[int]int)(nil), m1, equal[int]) {
		t.Errorf("EqualFunc(nil, %v, equal) = true, want false", m1)
	}
	if !EqualFunc[map[int]int, map[int]int](nil, nil, equal[int]) {
		t.Error("EqualFunc(nil, nil, equal) = false, want true")
	}
	if ms := map[int]int{1: 2}; EqualFunc(m1, ms, equal[int]) {
		t.Errorf("EqualFunc(%v, %v, equal) = true, want false", m1, ms)
	}

	// Comparing NaN for equality is expected to fail.
	mf := map[int]float64{1: 0, 2: math.NaN()}
	if EqualFunc(mf, mf, equal[float64]) {
		t.Errorf("EqualFunc(%v, %v, equal) = true, want false", mf, mf)
	}
	// But it should succeed using equalNaN.
	if !EqualFunc(mf, mf, equalNaN[float64]) {
		t.Errorf("EqualFunc(%v, %v, equalNaN) = false, want true", mf, mf)
	}

	if !EqualFunc(m1, m2, equalIntStr) {
		t.Errorf("EqualFunc(%v, %v, equalIntStr) = false, want true", m1, m2)
	}
}

func TestClone(t *testing.T) {
	mc := Clone(m1)
	if !Equal(mc, m1) {
		t.Errorf("Clone(%v) = %v, want %v", m1, mc, m1)
	}
	mc[16] = 32
	if Equal(mc, m1) {
		t.Errorf("Equal(%v, %v) = true, want false", mc, m1)
	}
}

func TestCloneNil(t *testing.T) {
	var m1 map[string]int
	mc := Clone(m1)
	if mc != nil {
		t.Errorf("Clone(%v) = %v, want %v", m1, mc, m1)
	}
}

func TestCopy(t *testing.T) {
	mc := Clone(m1)
	Copy(mc, mc)
	if !Equal(mc, m1) {
		t.Errorf("Copy(%v, %v) = %v, want %v", m1, m1, mc, m1)
	}
	Copy(mc, map[int]int{16: 32})
	want := map[int]int{1: 2, 2: 4, 4: 8, 8: 16, 16: 32}
	if !Equal(mc, want) {
		t.Errorf("Copy result = %v, want %v", mc, want)
	}

	type M1 map[int]bool
	type M2 map[int]bool
	Copy(make(M1), make(M2))
}

func TestDeleteFunc(t *testing.T) {
	mc := Clone(m1)
	DeleteFunc(mc, func(int, int) bool { return false })
	if !Equal(mc, m1) {
		t.Errorf("DeleteFunc(%v, true) = %v, want %v", m1, mc, m1)
	}
	DeleteFunc(mc, func(k, v int) bool { return k > 3 })
	want := map[int]int{1: 2, 2: 4}
	if !Equal(mc, want) {
		t.Errorf("DeleteFunc result = %v, want %v", mc, want)
	}
}

var n map[int]int

func BenchmarkMapClone(b *testing.B) {
	var m = make(map[int]int)
	for i := 0; i < 1000000; i++ {
		m[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n = Clone(m)
	}
}

func TestCloneWithDelete(t *testing.T) {
	var m = make(map[int]int)
	for i := 0; i < 32; i++ {
		m[i] = i
	}
	for i := 8; i < 32; i++ {
		delete(m, i)
	}
	m2 := Clone(m)
	if len(m2) != 8 {
		t.Errorf("len2(m2) = %d, want %d", len(m2), 8)
	}
	for i := 0; i < 8; i++ {
		if m2[i] != m[i] {
			t.Errorf("m2[%d] = %d, want %d", i, m2[i], m[i])
		}
	}
}

func TestCloneWithMapAssign(t *testing.T) {
	var m = make(map[int]int)
	const N = 25
	for i := 0; i < N; i++ {
		m[i] = i
	}
	m2 := Clone(m)
	if len(m2) != N {
		t.Errorf("len2(m2) = %d, want %d", len(m2), N)
	}
	for i := 0; i < N; i++ {
		if m2[i] != m[i] {
			t.Errorf("m2[%d] = %d, want %d", i, m2[i], m[i])
		}
	}
}

func TestCloneLarge(t *testing.T) {
	// See issue 64474.
	type K [17]float64 // > 128 bytes
	type V [17]float64

	var zero float64
	negZero := -zero

	for tst := 0; tst < 3; tst++ {
		// Initialize m with a key and value.
		m := map[K]V{}
		var k1 K
		var v1 V
		m[k1] = v1

		switch tst {
		case 0: // nothing, just a 1-entry map
		case 1:
			// Add more entries to make it 2 buckets
			// 1 entry already
			// 7 more fill up 1 bucket
			// 1 more to grow to 2 buckets
			for i := 0; i < 7+1; i++ {
				m[K{float64(i) + 1}] = V{}
			}
		case 2:
			// Capture the map mid-grow
			// 1 entry already
			// 7 more fill up 1 bucket
			// 5 more (13 total) fill up 2 buckets
			// 13 more (26 total) fill up 4 buckets
			// 1 more to start the 4->8 bucket grow
			for i := 0; i < 7+5+13+1; i++ {
				m[K{float64(i) + 1}] = V{}
			}
		}

		// Clone m, which should freeze the map's contents.
		c := Clone(m)

		// Update m with new key and value.
		k2, v2 := k1, v1
		k2[0] = negZero
		v2[0] = 1.0
		m[k2] = v2

		// Make sure c still has its old key and value.
		for k, v := range c {
			if math.Signbit(k[0]) {
				t.Errorf("tst%d: sign bit of key changed; got %v want %v", tst, k, k1)
			}
			if v != v1 {
				t.Errorf("tst%d: value changed; got %v want %v", tst, v, v1)
			}
		}
	}
}

func TestForEach(t *testing.T) {
	m := New[string, int]()
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3

	called := 0
	callbackFn := func(k string, v int) {
		called++
	}
	m.ForEach(callbackFn)

	if called != 3 {
		t.Errorf("Foreach got %d, want %d", called, 3)
	}
}

func TestFilter(t *testing.T) {
	m := New[rune, int]()
	m['a'] = 1
	m['b'] = 2
	m['c'] = 3
	m['d'] = 4
	m['e'] = 5
	m['f'] = 6
	m2 := m.Clone()

	callbackFn := func(k rune, v int) bool { return v%2 == 0 }
	got := m.Filter(callbackFn)
	size := got.Size()

	if size != 3 {
		t.Errorf("Size() got %d, want %d", size, 6)
	}

	want := Map[rune, int]{'b': 2, 'd': 4, 'f': 6}

	if !Equal(got, want) {
		t.Errorf("Filter got %v, want %v", got, want)
	}

	if !Equal(m, m2) {
		t.Errorf("Clone got %v, want %v", m2, m)
	}

}

func TestSome(t *testing.T) {
	m := New[byte, int]()
	m['a'] = 1
	m['b'] = 2
	m['c'] = 3

	callbackFn := func(k byte, v int) bool { return v%2 == 0 }

	got := m.Some(callbackFn)
	want := true
	if got != want {
		t.Errorf("Some(callbackFn) got %t, want %t", got, want)
	}

	callbackFn = func(k byte, v int) bool { return v > 3 }
	got = m.Some(callbackFn)
	want = false
	if got != want {
		t.Errorf("Some(callbackFn) got %t, want %t", got, want)
	}
}

func TestEvery(t *testing.T) {
	m := New[float32, int]()
	m[0.0] = 2
	m[1.1] = 4
	m[2.2] = 6

	callbackFn := func(k float32, v int) bool { return v%2 == 0 }
	got := m.Every(callbackFn)
	want := true
	if got != want {
		t.Errorf("Every(callbackFn) got %t, want %t", got, want)
	}

	callbackFn = func(k float32, v int) bool { return v > 3 }
	got = m.Every(callbackFn)
	want = false
	if got != want {
		t.Errorf("Every(callbackFn) got %t, want %t", got, want)
	}
}

func TestSize(t *testing.T) {
	m := New[string, int]()
	got := m.Size()
	if got != 0 {
		t.Errorf("Size() got %d, want %d", got, 0)
	}

	m["a"] = 1
	got = m.Size()
	if got != 1 {
		t.Errorf("Size() got %d, want %d", got, 1)
	}

	m["b"] = 2
	got = m.Size()
	if got != 2 {
		t.Errorf("Size() got %d, want %d", got, 2)
	}
}

func TestIsEmpty(t *testing.T) {
	m := New[string, int]()
	got := m.IsEmpty()
	if got != true {
		t.Errorf("IsEmpty() got %t, want %t", got, true)
	}

	m["a"] = 1
	got = m.IsEmpty()
	if got != false {
		t.Errorf("IsEmpty() got %t, want %t", got, false)
	}

	m.Clear()
	got = m.IsEmpty()
	if got != true {
		t.Errorf("IsEmpty() got %t, want %t", got, true)
	}
}

func TestMap_GetSet(t *testing.T) {
	m3 := New[string, int]()
	got := m3.Set("first", 123)
	want := m3["first"]
	if got != want {
		t.Errorf("Set(\"first\") got %d, want %d", got, want)
	}

	key := "second"
	if got, ok := m3.Get(key); ok {
		t.Errorf("Get(%q) got a value when it shouldn't: %v", key, got)
	}
}
