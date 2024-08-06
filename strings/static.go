package strings

import (
	"strings"
	"unicode"
)

// Returns the length of strings.String
func Len(s String) int { return len(s) }

// Returns the nth String of String.
//
// Negative indexing is supported (experimental)
func At(s String, n int) String {
	if n < 0 {
		n = len(s) + n
	}
	return String(s[n])
}

// Returns the nth byte of String.
//
// Negative indexing is supported (experimental)
func ByteAt(s String, n int) byte {
	if n < 0 {
		n = len(s) + n
	}
	return s[n]
}

// Returns the nth rune of strings.String.
//
// Negative indexing is supported (experimental)
func RuneAt(s String, n int) rune {
	if n < 0 {
		n = len(s) + n
	}
	return rune(s[n])
}

// Concatenates multiple strings.String together without modifying the base String. //
func Concatenate(s String, strs ...string) String {
	for _, value := range strs {
		s += String(value)
	}
	return s
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
func Clone(s String) String {
	return String(strings.Clone(string(s)))
}

// Compare returns an integer comparing two strings lexicographically.
// The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
//
// Compare is included only for symmetry with package bytes.
// It is usually clearer and always faster to use the built-in
// string comparison operators ==, <, >, and so on.
func Compare(str1 String, str2 String) int {
	return strings.Compare(string(str1), string(str2))
}

// Contains reports whether substr is within s.
func Contains(s String, substr String) bool {
	return strings.Contains(string(s), string(substr))
}

// ContainsAny reports whether any Unicode code points in chars are within s.
func ContainsAny(s String, chars String) bool {
	return strings.ContainsAny(string(s), string(chars))
}

// ContainsRune reports whether the Unicode code point r is within s.
func ContainsRune(s String, r rune) bool {
	return strings.ContainsRune(string(s), r)
}

// Count counts the number of non-overlapping instances of substr in s.
// If substr is an empty string, Count returns 1 + the number of Unicode code points in s.
func Count(s String, substr String) int {
	return strings.Count(string(s), string(substr))
}

// Cut slices s around the first instance of sep,
// returning the text before and after sep.
// The found result reports whether sep appears in s.
// If sep does not appear in s, cut returns s, "", false
func Cut(s String, sep String) (before String, after String, found bool) {
	bf, af, found := strings.Cut(string(s), string(sep))
	return String(bf), String(af), found
}

// CutPrefix returns s without the provided leading prefix string
// and reports whether it found the prefix.
// If s doesn't start with prefix, CutPrefix returns s, false.
// If prefix is the empty string, CutPrefix returns s, true.
func CutPrefix(s, prefix String) (after String, found bool) {
	if !HasPrefix(s, prefix) {
		return s, false
	}
	return s[len(prefix):], true
}

// CutSuffix returns s without the provided ending suffix string
// and reports whether it found the suffix.
// If s doesn't end with suffix, CutSuffix returns s, false.
// If suffix is the empty string, CutSuffix returns s, true.
func CutSuffix(s, suffix String) (before String, found bool) {
	if !HasSuffix(s, suffix) {
		return s, false
	}
	return s[:len(s)-len(suffix)], true
}

// EqualFold reports whether s and t, interpreted as UTF-8 strings,
// are equal under simple Unicode case-folding, which is a more general
// form of case-insensitivity.
func EqualFold(s String, t String) bool {
	return strings.EqualFold(string(s), string(t))
}

// Fields splits the string s around each instance of one or more consecutive white space
// characters, as defined by unicode.IsSpace, returning a slice of substrings of s or an
// empty slice if s contains only white space.
func Fields(s String) (fields []String) {
	native := strings.Fields(string(s))
	for _, value := range native {
		fields = append(fields, String(value))
	}
	return fields
}

// FieldsFunc splits the string s at each run of Unicode code points c satisfying f(c)
// and returns an array of slices of s. If all code points in s satisfy f(c) or the
// string is empty, an empty slice is returned.
//
// FieldsFunc makes no guarantees about the order in which it calls f(c)
// and assumes that f always returns the same value for a given c.
func FieldsFunc(s String, f func(rune) bool) (fields []String) {
	native := strings.FieldsFunc(string(s), f)
	fields = make([]String, len(native))
	for i, value := range native {
		fields[i] = String(value)
	}
	return fields
}

// HasPrefix reports whether the string s begins with prefix.
func HasPrefix(s String, prefix String) bool {
	return strings.HasPrefix(string(s), string(prefix))
}

// HasSuffix reports whether the string s ends with suffix.
func HasSuffix(s String, prefix String) bool {
	return strings.HasSuffix(string(s), string(prefix))
}

// Index returns the index of the first instance of substr in s, or -1 if substr is not present in s.
func Index(s String, substr String) int {
	return strings.Index(string(s), string(substr))
}

// IndexAny returns the index of the first instance of any Unicode code point
// from chars in s, or -1 if no Unicode code point from chars is present in s.
func IndexAny(s String, substr String) int {
	return strings.IndexAny(string(s), string(substr))
}

// IndexByte returns the index of the first instance of c in s, or -1 if c is not present in s.
func IndexByte(s String, c byte) int {
	return strings.IndexByte(string(s), c)
}

// IndexFunc returns the index into s of the first Unicode
// code point satisfying f(c), or -1 if none do.
func IndexFunc(s String, f func(rune) bool) int {
	return strings.IndexFunc(string(s), f)
}

// IndexRune returns the index of the first instance of the Unicode code point
// r, or -1 if rune is not present in s.
// If r is utf8.RuneError, it returns the first instance of any
// invalid UTF-8 byte sequence.
func IndexRune(s String, substr rune) int {
	return strings.IndexRune(string(s), substr)
}

// Join concatenates the elements of its first argument to create a single string. The separator
// string sep is placed between elements in the resulting string.
func Join(elems []String, sep String) String {
	strElems := make([]string, len(elems))
	for i, elem := range elems {
		strElems[i] = string(elem)
	}
	return String(strings.Join(strElems, string(sep)))
}

// LastIndex returns the index of the last instance of substr in s, or -1 if substr is not present in s.
func LastIndex(s String, substr String) int {
	return strings.LastIndex(string(s), string(substr))
}

// LastIndexAny returns the index of the last instance of any Unicode code
// point from chars in s, or -1 if no Unicode code point from chars is
// present in s.
func LastIndexAny(s String, chars String) int {
	return strings.LastIndexAny(string(s), string(chars))
}

// LastIndexByte returns the index of the last instance of c in s, or -1 if c is not present in s.
func LastIndexByte(s String, c byte) int {
	return strings.LastIndexByte(string(s), c)
}

// LastIndexFunc returns the index into s of the last
// Unicode code point satisfying f(c), or -1 if none do.
func LastIndexFunc(s String, f func(rune) bool) int {
	return strings.LastIndexFunc(string(s), f)
}

// Map returns a copy of the string s with all its characters modified
// according to the mapping function. If mapping returns a negative value, the character is
// dropped from the string with no replacement.
func Map(mapping func(rune) rune, s String) String {
	return String(strings.Map(mapping, string(s)))
}

// Repeat returns a new string consisting of count copies of the string s.
//
// It panics if count is negative or if the result of (len(s) * count)
// overflows.
func Repeat(s String, count int) String {
	return String(strings.Repeat(string(s), count))
}

// Replace returns a copy of the string s with the first n
// non-overlapping instances of old replaced by new.
// If old is empty, it matches at the beginning of the string
// and after each UTF-8 sequence, yielding up to k+1 replacements
// for a k-rune string.
// If n < 0, there is no limit on the number of replacements.
func Replace(s String, old String, new String, n int) String {
	return String(strings.Replace(string(s), string(old), string(new), n))
}

// ReplaceAll returns a copy of the string s with all
// non-overlapping instances of old replaced by new.
// If old is empty, it matches at the beginning of the string
// and after each UTF-8 sequence, yielding up to k+1 replacements
// for a k-rune string.
func ReplaceAll(s, old, new String) String {
	return String(strings.ReplaceAll(string(s), string(old), string(new)))
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
func Split(s String, sep String) []String {
	native := strings.Split(string(s), string(sep))
	strs := make([]String, len(native))
	for i, value := range native {
		strs[i] = String(value)
	}
	return strs
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
func SplitAfter(s String, sep String) []String {
	native := strings.SplitAfter(string(s), string(sep))
	strs := make([]String, len(native))
	for i, value := range native {
		strs[i] = String(value)
	}
	return strs
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
func SplitAfterN(s String, sep String, n int) []String {
	native := strings.SplitAfterN(string(s), string(sep), n)
	strs := make([]String, len(native))
	for i, value := range native {
		strs[i] = String(value)
	}
	return strs
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
func SplitN(s String, sep String, n int) []String {
	native := strings.SplitN(string(s), string(sep), n)
	strs := make([]String, len(native))
	for i, value := range native {
		strs[i] = String(value)
	}
	return strs
}

// ToLower returns s with all Unicode letters mapped to their lower case.
func ToLower(s String) String {
	return String(strings.ToLower(string(s)))
}

// ToLowerSpecial returns a copy of the string s with all Unicode letters mapped to their
// lower case using the case mapping specified by c.
func ToLowerSpecial(c unicode.SpecialCase, s String) String {
	return String(strings.ToLowerSpecial(c, string(s)))
}

// ToTitle returns a copy of the string s with all Unicode letters mapped to
// their Unicode title case.
func ToTitle(s String) String {
	return String(strings.ToTitle(string(s)))
}

// ToTitleSpecial returns a copy of the string s with all Unicode letters mapped to their
// Unicode title case, giving priority to the special casing rules.
func ToTitleSpecial(c unicode.SpecialCase, s String) String {
	return String(strings.ToTitleSpecial(c, string(s)))
}

// ToUpper returns s with all Unicode letters mapped to their upper case.
func ToUpper(s String) String {
	return String(strings.ToUpper(string(s)))
}

// ToUpperSpecial returns a copy of the string s with all Unicode letters mapped to their
// upper case using the case mapping specified by c.
func ToUpperSpecial(c unicode.SpecialCase, s String) String {
	return String(strings.ToUpperSpecial(c, string(s)))
}

// ToValidUTF8 returns a copy of the string s with each run of invalid UTF-8 byte sequences
// replaced by the replacement string, which may be empty.
func ToValidUTF8(s String, replacement String) String {
	return String(strings.ToValidUTF8(string(s), string(replacement)))
}

// Trim returns a slice of the string s with all leading and
// trailing Unicode code points contained in cutset removed.
func Trim(s String, cutset String) String {
	return String(strings.Trim(string(s), string(cutset)))
}

// TrimFunc returns a slice of the string s with all leading
// and trailing Unicode code points c satisfying f(c) removed.
func TrimFunc(s String, f func(rune) bool) String {
	return String(strings.TrimFunc(string(s), f))
}

// TrimLeft returns a slice of the string s with all leading
// Unicode code points contained in cutset removed.
//
// To remove a prefix, use [TrimPrefix] instead.
func TrimLeft(s String, cutset String) String {
	return String(strings.TrimLeft(string(s), string(cutset)))
}

// TrimLeftFunc returns a slice of the string s with all leading
// Unicode code points c satisfying f(c) removed.
func TrimLeftFunc(s String, f func(rune) bool) String {
	return String(strings.TrimLeftFunc(string(s), f))
}

// TrimPrefix returns s without the provided leading prefix string.
// If s doesn't start with prefix, s is returned unchanged.
func TrimPrefix(s String, prefix String) String {
	return String(strings.TrimPrefix(string(s), string(prefix)))
}

// TrimRight returns a slice of the string s, with all trailing
// Unicode code points contained in cutset removed.
//
// To remove a suffix, use [TrimSuffix] instead.
func TrimRight(s String, cutset String) String {
	return String(strings.TrimRight(string(s), string(cutset)))
}

// TrimRightFunc returns a slice of the string s with all trailing
// Unicode code points c satisfying f(c) removed.
func TrimRightFunc(s String, f func(rune) bool) String {
	return String(strings.TrimRightFunc(string(s), f))
}

// TrimSpace returns a slice of the string s, with all leading
// and trailing white space removed, as defined by Unicode.
func TrimSpace(s String) String {
	return String(strings.TrimSpace(string(s)))
}

// TrimSuffix returns s without the provided trailing suffix string.
// If s doesn't end with suffix, s is returned unchanged.
func TrimSuffix(s String, suffix String) String {
	return String(strings.TrimSuffix(string(s), string(suffix)))
}
