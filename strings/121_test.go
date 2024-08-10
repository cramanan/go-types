//go:build go1.20

package strings_test

import (
	"strings"
	"testing"
	"unicode"
	"unicode/utf8"
	"unsafe"

	. "github.com/cramanan/go-types/strings"
)

func TestCompareStrings_121(t *testing.T) {
	// unsafeString converts a []byte to a string with no allocation.
	// The caller must not modify b while the result string is in use.
	unsafeString := func(b []byte) string {
		return unsafe.String(unsafe.SliceData(b), len(b))
	}

	lengths := make([]int, 0) // lengths to test in ascending order
	for i := 0; i <= 128; i++ {
		lengths = append(lengths, i)
	}
	lengths = append(lengths, 256, 512, 1024, 1333, 4095, 4096, 4097)

	// if !testing.Short() || testenv.Builder() != "" {
	// 	lengths = append(lengths, 65535, 65536, 65537, 99999)
	// }

	n := lengths[len(lengths)-1]
	a := make([]byte, n+1)
	b := make([]byte, n+1)
	lastLen := 0
	for _, len := range lengths {
		// randomish but deterministic data. No 0 or 255.
		for i := 0; i < len; i++ {
			a[i] = byte(1 + 31*i%254)
			b[i] = byte(1 + 31*i%254)
		}
		// data past the end is different
		for i := len; i <= n; i++ {
			a[i] = 8
			b[i] = 9
		}

		sa, sb := unsafeString(a), unsafeString(b)
		cmp := Compare(sa[:len], sb[:len])
		if cmp != 0 {
			t.Errorf(`CompareIdentical(%d) = %d`, len, cmp)
		}
		if len > 0 {
			cmp = Compare(sa[:len-1], sb[:len])
			if cmp != -1 {
				t.Errorf(`CompareAshorter(%d) = %d`, len, cmp)
			}
			cmp = Compare(sa[:len], sb[:len-1])
			if cmp != 1 {
				t.Errorf(`CompareBshorter(%d) = %d`, len, cmp)
			}
		}
		for k := lastLen; k < len; k++ {
			b[k] = a[k] - 1
			cmp = Compare(unsafeString(a[:len]), unsafeString(b[:len]))
			if cmp != 1 {
				t.Errorf(`CompareAbigger(%d,%d) = %d`, len, k, cmp)
			}
			b[k] = a[k] + 1
			cmp = Compare(unsafeString(a[:len]), unsafeString(b[:len]))
			if cmp != -1 {
				t.Errorf(`CompareBbigger(%d,%d) = %d`, len, k, cmp)
			}
			b[k] = a[k]
		}
		lastLen = len
	}
}

func TestClone_121(t *testing.T) {
	var cloneTests = []string{
		"",
		strings.Clone(""),
		strings.Repeat("a", 42)[:0],
		"short",
		strings.Repeat("a", 42),
	}
	for _, input := range cloneTests {
		clone := strings.Clone(input)
		if clone != input {
			t.Errorf("Clone(%q) = %q; want %q", input, clone, input)
		}

		if len(input) != 0 && unsafe.StringData(clone) == unsafe.StringData(input) {
			t.Errorf("Clone(%q) return value should not reference inputs backing memory.", input)
		}

		if len(input) == 0 && unsafe.StringData(clone) != unsafe.StringData(emptyString) {
			t.Errorf("Clone(%#v) return value should be equal to empty string.", unsafe.StringData(input))
		}
	}
}

func TestMap_121(t *testing.T) {
	// Run a couple of awful growth/shrinkage tests
	a := tenRunes('a')
	// 1.  Grow. This triggers two reallocations in Map.
	maxRune := func(rune) rune { return unicode.MaxRune }
	m := Map(maxRune, a)
	expect := tenRunes(unicode.MaxRune)
	if m != expect {
		t.Errorf("growing: expected %q got %q", expect, m)
	}

	// 2. Shrink
	minRune := func(rune) rune { return 'a' }
	m = Map(minRune, tenRunes(unicode.MaxRune))
	expect = a
	if m != expect {
		t.Errorf("shrinking: expected %q got %q", expect, m)
	}

	// 3. Rot13
	m = Map(rot13, "a to zed")
	expect = "n gb mrq"
	if m != expect {
		t.Errorf("rot13: expected %q got %q", expect, m)
	}

	// 4. Rot13^2
	m = Map(rot13, Map(rot13, "a to zed"))
	expect = "a to zed"
	if m != expect {
		t.Errorf("rot13: expected %q got %q", expect, m)
	}

	// 5. Drop
	dropNotLatin := func(r rune) rune {
		if unicode.Is(unicode.Latin, r) {
			return r
		}
		return -1
	}
	m = Map(dropNotLatin, "Hello, 세계")
	expect = "Hello"
	if m != expect {
		t.Errorf("drop: expected %q got %q", expect, m)
	}

	// 6. Identity
	identity := func(r rune) rune {
		return r
	}
	orig := "Input string that we expect not to be copied."
	m = Map(identity, orig)
	if unsafe.StringData(orig) != unsafe.StringData(m) {
		t.Error("unexpected copy during identity map")
	}

	// 7. Handle invalid UTF-8 sequence
	replaceNotLatin := func(r rune) rune {
		if unicode.Is(unicode.Latin, r) {
			return r
		}
		return utf8.RuneError
	}
	m = Map(replaceNotLatin, "Hello\255World")
	expect = "Hello\uFFFDWorld"
	if m != expect {
		t.Errorf("replace invalid sequence: expected %q got %q", expect, m)
	}

	// 8. Check utf8.RuneSelf and utf8.MaxRune encoding
	encode := func(r rune) rune {
		switch r {
		case utf8.RuneSelf:
			return unicode.MaxRune
		case unicode.MaxRune:
			return utf8.RuneSelf
		}
		return r
	}
	s := string(rune(utf8.RuneSelf)) + string(utf8.MaxRune)
	r := string(utf8.MaxRune) + string(rune(utf8.RuneSelf)) // reverse of s
	m = Map(encode, s)
	if m != r {
		t.Errorf("encoding not handled correctly: expected %q got %q", r, m)
	}
	m = Map(encode, r)
	if m != s {
		t.Errorf("encoding not handled correctly: expected %q got %q", s, m)
	}

	// 9. Check mapping occurs in the front, middle and back
	trimSpaces := func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}
	m = Map(trimSpaces, "   abc    123   ")
	expect = "abc123"
	if m != expect {
		t.Errorf("trimSpaces: expected %q got %q", expect, m)
	}
}
