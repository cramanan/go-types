package strings

import (
	"fmt"
	"strings"
	"unicode"
)

// NewReader returns a new [Reader] reading from s. It is similar to bytes.NewBufferString but more efficient and non-writable.
func NewReader[S IString](s S) *strings.Reader { return strings.NewReader(string(s)) }

// NewReplacer returns a new [Replacer] from a list of old, new string
// pairs. Replacements are performed in the order they appear in the
// target string, without overlapping matches. The old string
// comparisons are done in argument order.
//
// NewReplacer panics if given an odd number of arguments.
func NewReplacer[S IString](s ...S) *strings.Replacer {
	strs := make([]string, len(s))
	for i, v := range s {
		strs[i] = string(v)
	}
	return strings.NewReplacer(strs...)
}

// Len returns the length of the string | []byte | []rune s.
func Len[S IString](s S) int { return len(s) }

// At returns the character at index n, the type parameter defines the return type
func At[C IChar | ~string, S IString](s S, n int) C {
	idx := n
	if n < 0 {
		n = len(s) + n
	}

	if n < 0 || n > len(s)-1 {
		panic(fmt.Sprintf("index out of range: strings.At(%d) for %T of length %d", idx, s, len(s)))
	}

	return C(string(s)[n])
}

// Concatenate concatenates a variable number of strings into a single string.
// The type of the resulting string is determined by the type parameter T.
func Concatenate[T IString](first T, strs ...T) T {
	toString := append([]string{}, string(first))
	for _, s := range strs {
		toString = append(toString, string(s))
	}

	res := ""
	for _, s := range toString {
		res += s
	}

	return T(res)
}

// Clone returns a fresh copy of s.
// It guarantees to make a copy of s into a new allocation,
// which can be important when retaining only a small substring
// of a much larger string. Using Clone can help such programs
// use less memory. Of course, since using Clone makes a copy,
// overuse of Clone can make programs use more memory.
// Clone should typically be used only rarely, and only when
// profiling indicates that it is needed.
// For strings of length zero the string "" will be returned
// and no allocation is made.
func Clone[S IString](s S) S {
	return S(strings.Clone(string(s)))
}

// Compare returns an integer comparing two strings lexicographically.
// The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
//
// Compare is included only for symmetry with package bytes.
// It is usually clearer and always faster to use the built-in
// string comparison operators ==, <, >, and so on.
func Compare[S1, S2 IString](str1 S1, str2 S2) int {
	return strings.Compare(string(str1), string(str2))
}

// Contains reports whether substr is within s.
func Contains[S1, S2 IString | IChar](s S1, substr S2) bool {
	return strings.Contains(string(s), string(substr))
}

// ContainsAny reports whether any Unicode code points in chars are within s.
func ContainsAny[S1, S2 IString | IChar](s S1, chars S2) bool {
	return strings.ContainsAny(string(s), string(chars))
}

func ContainsFunc[S IString](s S, callbackFn func(rune) bool) bool {
	if callbackFn == nil {
		panic("callback function is nil")
	}
	return IndexFunc(string(s), callbackFn) >= 0
}

// Count counts the number of non-overlapping instances of substr in s.
// If substr is an empty string, Count returns 1 + the number of Unicode code points in s.
func Count[S1, S2 IString | IChar](s S1, substr S2) int {
	return strings.Count(string(s), string(substr))
}

// Cut slices s around the first instance of sep,
// returning the text before and after sep.
// The found result reports whether sep appears in s.
// If sep does not appear in s, cut returns s, "", false
func Cut[S1 IString, S2 IString | IChar](s S1, sep S2) (before S1, after S1, found bool) {
	bf, af, found := strings.Cut(string(s), string(sep))
	return S1(bf), S1(af), found
}

// HasPrefix reports whether the string s begins with prefix.
func HasPrefix[S1, S2 IString | IChar](s S1, prefix S2) bool {
	return strings.HasPrefix(string(s), string(prefix))
}

// HasSuffix reports whether the string s ends with suffix.
func HasSuffix[S1, S2 IString | IChar](s S1, suffix S2) bool {
	return strings.HasSuffix(string(s), string(suffix))
}

// CutPrefix returns s without the provided leading prefix string
// and reports whether it found the prefix.
// If s doesn't start with prefix, CutPrefix returns s, false.
// If prefix is the empty string, CutPrefix returns s, true.
func CutPrefix[S1 IString, S2 IString | IChar](s S1, prefix S2) (after S1, found bool) {
	if !HasPrefix(s, prefix) {
		return (s), false
	}
	return S1(string(s)[len(string(prefix)):]), true
}

// CutSuffix returns s without the provided ending suffix string
// and reports whether it found the suffix.
// If s doesn't end with suffix, CutSuffix returns s, false.
// If suffix is the empty string, CutSuffix returns s, true.
func CutSuffix[S1 IString, S2 IString | IChar](s S1, suffix S2) (after S1, found bool) {
	if !HasSuffix(s, suffix) {
		return s, false
	}
	return S1(string(s)[:len(s)-len(string(suffix))]), true
}

// EqualFold reports whether s and t, interpreted as UTF-8 strings,
// are equal under simple Unicode case-folding, which is a more general
// form of case-insensitivity.
func EqualFold[S1, S2 IString | IChar](s S1, t S2) bool {
	return strings.EqualFold(string(s), string(t))
}

// Fields splits the string s around each instance of one or more consecutive white space
// characters, as defined by unicode.IsSpace, returning a slice of substrings of s or an
// empty slice if s contains only white space.
func Fields[S IString](s S) (fields []S) {
	native := strings.Fields(string(s))
	for _, value := range native {
		fields = append(fields, S(value))
	}
	return fields
}

// FieldsFunc splits the string s at each run of Unicode code points c satisfying f(c)
// and returns an array of slices of s. If all code points in s satisfy f(c) or the
// string is empty, an empty slice is returned.
//
// FieldsFunc makes no guarantees about the order in which it calls f(c)
// and assumes that f always returns the same value for a given c.
func FieldsFunc[S IString, C IChar](s S, f func(C) bool) (fields []S) {
	fn := func(c rune) bool { return f(C(c)) }
	for _, value := range strings.FieldsFunc(string(s), fn) {
		fields = append(fields, S(value))
	}
	return fields
}

// Index returns the index of the first instance of substr in s, or -1 if substr is not present in s.
func Index[S1, S2 IString | IChar](str S1, substr S2) int {
	switch sub := any(substr).(type) {
	case rune:
		return strings.IndexRune(string(str), sub)
	case byte:
		return strings.IndexByte(string(str), sub)
	case string:
		return strings.Index(string(str), sub)
	default:
		return -1
	}
}

// IndexAny returns the index of the first instance of any Unicode code point
// from chars in s, or -1 if no Unicode code point from chars is present in s.
func IndexAny[S1, S2 IString | IChar](s S1, substr S2) int {
	return strings.IndexAny(string(s), string(substr))
}

// IndexFunc returns the index into s of the first Unicode
// code point satisfying f(c), or -1 if none do.
func IndexFunc[S IString, C IChar](s S, f func(C) bool) int {
	fn := func(c rune) bool { return f(C(c)) }
	return strings.IndexFunc(string(s), fn)
}

// Join concatenates the elements of its first argument to create a single string. The separator
// string sep is placed between elements in the resulting string.
func Join[S1 IString | IChar, S2 IString](elems []S1, sep S2) S2 {
	strElems := make([]string, len(elems))
	for i, elem := range elems {
		strElems[i] = string(elem)
	}
	return S2(strings.Join(strElems, string(sep)))
}

// LastIndex returns the index of the last instance of substr in s, or -1 if substr is not present in s.
func LastIndex[S1, S2 IString | IChar](s S1, substr S2) int {
	return strings.LastIndex(string(s), string(substr))
}

// LastIndex returns the index of the last instance of substr in s, or -1 if substr is not present in s.
func LastIndexFunc[S IString, C IChar](s S, f func(C) bool) int {
	fn := func(c rune) bool { return f(C(c)) }
	return strings.LastIndexFunc(string(s), fn)
}

// LastIndexAny returns the index of the last instance of any Unicode code
// point from chars in s, or -1 if no Unicode code point from chars is
// present in s.
func LastIndexAny[S1, S2 IString | IChar](s S1, chars S2) int {
	return strings.LastIndexAny(string(s), string(chars))
}

// Map returns a copy of the string s with all its characters modified
// according to the mapping function. If mapping returns a negative value, the character is
// dropped from the string with no replacement.
func Map[S IString](callbackFn func(rune) rune, s S) S {
	if callbackFn == nil {
		panic("callback function is nil")
	}
	return S(strings.Map(callbackFn, string(s)))
}

// Repeat returns a new string consisting of count copies of the string s.
//
// It panics if count is negative or if the result of (len(s) * count)
// overflows.
func Repeat[S IString](s S, count int) S { return S(strings.Repeat(string(s), count)) }

// Replace returns a copy of the string s with the first n
// non-overlapping instances of old replaced by new.
// If old is empty, it matches at the beginning of the string
// and after each UTF-8 sequence, yielding up to k+1 replacements
// for a k-rune string.
// If n < 0, there is no limit on the number of replacements.
func Replace[S1 IString, S2 IString | IChar](s S1, old S2, new S2, n int) S1 {
	return S1(strings.Replace(string(s), string(old), string(new), n))
}

// ReplaceAll returns a copy of the string s with all
// non-overlapping instances of old replaced by new.
// If old is empty, it matches at the beginning of the string
// and after each UTF-8 sequence, yielding up to k+1 replacements
// for a k-rune string.
func ReplaceAll[S1 IString, S2 IString | IChar](s S1, old, new S2) S1 {
	return S1(strings.ReplaceAll(string(s), string(old), string(new)))
}

// Split slices s into all substrings separated by sep and returns a slice of
// the substrings between those separators.
//
// If s does not contain sep and sep is not empty, Split returns a
// slice of length 1 whose only element is s.
//
// If sep is empty, Split splits after each UTF-8 sequence. If both s
// and sep are empty, Split returns an empty slice.
//
// It is equivalent to [SplitN] with a count of -1.
//
// To split around the first instance of a separator, see Cut.
func Split[S1 IString, S2 IString | IChar](s S1, sep S2) (res []S1) {
	for _, value := range strings.Split(string(s), string(sep)) {
		res = append(res, S1(value))
	}
	return res
}

// SplitAfter slices s into all substrings after each instance of sep and
// returns a slice of those substrings.
//
// If s does not contain sep and sep is not empty, SplitAfter returns
// a slice of length 1 whose only element is s.
//
// If sep is empty, SplitAfter splits after each UTF-8 sequence. If
// both s and sep are empty, SplitAfter returns an empty slice.
//
// It is equivalent to [SplitAfterN] with a count of -1.
func SplitAfter[S1 IString, S2 IString | IChar](s S1, sep S2) (res []S1) {
	for _, value := range strings.SplitAfter(string(s), string(sep)) {
		res = append(res, S1(value))
	}
	return res
}

// SplitAfterN slices s into substrings after each instance of sep and
// returns a slice of those substrings.
//
// The count determines the number of substrings to return:
//
//	n > 0: at most n substrings; the last substring will be the unsplit remainder.
//	n == 0: the result is nil (zero substrings)
//	n < 0: all substrings
//
// Edge cases for s and sep (for example, empty strings) are handled
// as described in the documentation for SplitAfter.
func SplitAfterN[S1 IString, S2 IString | IChar](s S1, sep S2, n int) (res []S1) {
	for _, value := range strings.SplitAfterN(string(s), string(sep), n) {
		res = append(res, S1(value))
	}
	return res
}

// SplitN slices s into substrings separated by sep and returns a slice of
// the substrings between those separators.
//
// The count determines the number of substrings to return:
//
//	n > 0: at most n substrings; the last substring will be the unsplit remainder.
//	n == 0: the result is nil (zero substrings)
//	n < 0: all substrings
//
// Edge cases for s and sep (for example, empty strings) are handled
// as described in the documentation for [Split].
//
// To split around the first instance of a separator, see Cut.
func SplitN[S1 IString, S2 IString | IChar](s S1, sep S2, n int) (res []S1) {
	for _, value := range strings.SplitN(string(s), string(sep), n) {
		res = append(res, S1(value))
	}
	return res
}

// ToLower returns s with all Unicode letters mapped to their lower case.
func ToLower[S IString](s S) S { return S(strings.ToLower(string(s))) }

// ToLowerSpecial returns a copy of the string s with all Unicode letters mapped to their
// lower case using the case mapping specified by c.
func ToLowerSpecial[S IString](c unicode.SpecialCase, s S) S {
	return S(strings.ToLowerSpecial(c, string(s)))
}

// ToTitle returns a copy of the string s with all Unicode letters mapped to
// their Unicode title case.
func ToTitle[S IString](s S) S { return S(strings.ToTitle(string(s))) }

// ToTitleSpecial returns a copy of the string s with all Unicode letters mapped to their
// Unicode title case, giving priority to the special casing rules.
func ToTitleSpecial[S IString](c unicode.SpecialCase, s S) S {
	return S(strings.ToTitleSpecial(c, string(s)))
}

// ToUpper returns s with all Unicode letters mapped to their upper case.
func ToUpper[S IString](s S) S { return S(strings.ToUpper(string(s))) }

// ToUpperSpecial returns a copy of the string s with all Unicode letters mapped to their
// upper case using the case mapping specified by c.
func ToUpperSpecial[S IString](c unicode.SpecialCase, s S) S {
	return S(strings.ToUpperSpecial(c, string(s)))
}

// ToValidUTF8 returns a copy of the string s with each run of invalid UTF-8 byte sequences
// replaced by the replacement string, which may be empty.
func ToValidUTF8[S1 IString, S2 IString | IChar](s S1, replacement S2) S1 {
	return S1(strings.ToValidUTF8(string(s), string(replacement)))
}

// Trim returns a slice of the string s with all leading and
// trailing Unicode code points contained in cutset removed.
func Trim[S1 IString, S2 IString | IChar](s S1, cutset S2) S1 {
	return S1(strings.Trim(string(s), string(cutset)))
}

// TrimFunc returns a slice of the string s with all leading
// and trailing Unicode code points c satisfying f(c) removed.
func TrimFunc[S IString](s S, callbackFn func(rune) bool) S {
	if callbackFn == nil {
		panic("callback function is nil")
	}
	return S(strings.TrimFunc(string(s), callbackFn))
}

// TrimLeft returns a slice of the string s with all leading
// Unicode code points contained in cutset removed.
//
// To remove a prefix, use [TrimPrefix] instead.
func TrimLeft[C IString, S IString](s S, cutset C) S {
	return S(strings.TrimLeft(string(s), string(cutset)))
}

// TrimLeftFunc returns a slice of the string s with all leading
// Unicode code points c satisfying f(c) removed.
func TrimLeftFunc[S IString](s S, callbackFn func(rune) bool) S {
	if callbackFn == nil {
		panic("callback function is nil")
	}
	return S(strings.TrimLeftFunc(string(s), callbackFn))
}

// TrimPrefix returns s without the provided leading prefix string.
// If s doesn't start with prefix, s is returned unchanged.
func TrimPrefix[S IString, C IString | IChar](s S, prefix C) S {
	return S(strings.TrimPrefix(string(s), string(prefix)))
}

// TrimRight returns a slice of the string s, with all trailing
// Unicode code points contained in cutset removed.
//
// To remove a suffix, use [TrimSuffix] instead.
func TrimRight[S1 IString, S2 IString | IChar](s S1, cutset S2) S1 {
	return S1(strings.TrimRight(string(s), string(cutset)))
}

// TrimRightFunc returns a slice of the string s with all trailing
// Unicode code points c satisfying f(c) removed.
func TrimRightFunc[S IString](s S, callbackFn func(rune) bool) S {
	if callbackFn == nil {
		panic("callback function is nil")
	}
	return S(strings.TrimRightFunc(string(s), callbackFn))
}

// TrimSpace returns a slice of the string s, with all leading
// and trailing white space removed, as defined by Unicode.
func TrimSpace[S IString](s S) S {
	return S(strings.TrimSpace(string(s)))
}

// TrimSuffix returns s without the provided trailing suffix string.
// If s doesn't end with suffix, s is returned unchanged.
func TrimSuffix[S IString, C IString | IChar](s S, suffix C) S {
	return S(strings.TrimSuffix(string(s), string(suffix)))
}
