package strings

import (
	"unicode"
)

// Returns the native string type of the String //
func (s String) String() string { return string(s) }

// Returns the length of String
func (s String) Len() int { return len(s) }

// Returns the nth String of String.
//
// Negative indexing is supported (experimental)
func (s String) At(n int) String {
	if n < 0 {
		n = len(s) + n
	}
	return String(s[n])
}

//
// Returns the nth byte of String.

// Negative indexing is supported (experimental)
func (s String) ByteAt(n int) byte {
	if n < 0 {
		n = len(s) + n
	}
	return s[n]
}

//
// Returns the nth rune of String.

// Negative indexing is supported (experimental)
func (s String) RuneAt(n int) rune {
	if n < 0 {
		n = len(s) + n
	}
	return rune(s[n])
}

// Returns String as a slice of bytes //
func (s String) Bytes() []byte {
	return []byte(s)
}

// Concatenates multiple String together without modifying the base String. //
func (s String) Concatenate(strs ...String) String {
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
func (s String) Clone() String {
	return String(Clone(s))
}

// Compare returns an integer comparing two strings lexicographically.
// The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
//
// Compare is included only for symmetry with package bytes.
// It is usually clearer and always faster to use the built-in
// string comparison operators ==, <, >, and so on.
func (str1 String) Compare(str2 String) int {
	return Compare((str1), (str2))
}

// Contains reports whether substr is within s.
func (s String) Contains(substr String) bool {
	return Contains((s), (substr))
}

// ContainsAny reports whether any Unicode code points in chars are within s.
func (s String) ContainsAny(chars String) bool {
	return ContainsAny((s), (chars))
}

// ContainsRune reports whether the Unicode code point r is within s.
func (s String) ContainsRune(r rune) bool {
	return ContainsRune((s), r)
}

// Count counts the number of non-overlapping instances of substr in s.
// If substr is an empty string, Count returns 1 + the number of Unicode code points in s.
func (s String) Count(substr String) int {
	return Count((s), (substr))
}

// Cut slices s around the first instance of sep,
// returning the text before and after sep.
// The found result reports whether sep appears in s.
// If sep does not appear in s, cut returns s, "", false
func (s String) Cut(sep String) (before String, after String, found bool) {
	bf, af, found := Cut((s), (sep))
	return String(bf), String(af), found
}

// CutPrefix returns s without the provided leading prefix string
// and reports whether it found the prefix.
// If s doesn't start with prefix, CutPrefix returns s, false.
// If prefix is the empty string, CutPrefix returns s, true.
func (s String) CutPrefix(prefix String) (after String, found bool) {
	af, found := CutPrefix((s), (prefix))
	return String(af), found
}

// CutSuffix returns s without the provided ending suffix string
// and reports whether it found the suffix.
// If s doesn't end with suffix, CutSuffix returns s, false.
// If suffix is the empty string, CutSuffix returns s, true.
func (s String) CutSuffix(prefix String) (after String, found bool) {
	af, found := CutSuffix((s), (prefix))
	return String(af), found
}

// EqualFold reports whether s and t, interpreted as UTF-8 strings,
// are equal under simple Unicode case-folding, which is a more general
// form of case-insensitivity.
func (s String) EqualFold(t String) bool {
	return EqualFold((s), (t))
}

// Fields splits the string s around each instance of one or more consecutive white space
// characters, as defined by unicode.IsSpace, returning a slice of substrings of s or an
// empty slice if s contains only white space.
func (s String) Fields() (fields []String) {
	native := Fields((s))
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
func (s String) FieldsFunc(f func(rune) bool) (fields []String) {
	native := FieldsFunc((s), f)
	fields = make([]String, len(native))
	for i, value := range native {
		fields[i] = String(value)
	}
	return fields
}

// HasPrefix reports whether the string s begins with prefix.
func (s String) HasPrefix(prefix String) bool {
	return HasPrefix((s), (prefix))
}

// HasSuffix reports whether the string s ends with suffix.
func (s String) HasSuffix(prefix String) bool {
	return HasSuffix((s), (prefix))
}

// Index returns the index of the first instance of substr in s, or -1 if substr is not present in s.
func (s String) Index(substr String) int {
	return Index((s), (substr))
}

// IndexAny returns the index of the first instance of any Unicode code point
// from chars in s, or -1 if no Unicode code point from chars is present in s.
func (s String) IndexAny(substr String) int {
	return IndexAny((s), (substr))
}

// IndexByte returns the index of the first instance of c in s, or -1 if c is not present in s.
func (s String) IndexByte(c byte) int {
	return IndexByte((s), c)
}

// IndexFunc returns the index into s of the first Unicode
// code point satisfying f(c), or -1 if none do.
func (s String) IndexFunc(f func(rune) bool) int {
	return IndexFunc((s), f)
}

// IndexRune returns the index of the first instance of the Unicode code point
// r, or -1 if rune is not present in s.
// If r is utf8.RuneError, it returns the first instance of any
// invalid UTF-8 byte sequence.
func (s String) IndexRune(substr rune) int {
	return IndexRune((s), substr)
}

// LastIndex returns the index of the last instance of substr in s, or -1 if substr is not present in s.
func (s String) LastIndex(substr String) int {
	return LastIndex((s), (substr))
}

// LastIndexAny returns the index of the last instance of any Unicode code
// point from chars in s, or -1 if no Unicode code point from chars is
// present in s.
func (s String) LastIndexAny(chars String) int {
	return LastIndexAny((s), (chars))
}

// LastIndexByte returns the index of the last instance of c in s, or -1 if c is not present in s.
func (s String) LastIndexByte(c byte) int {
	return LastIndexByte((s), c)
}

// LastIndexFunc returns the index into s of the last
// Unicode code point satisfying f(c), or -1 if none do.
func (s String) LastIndexFunc(f func(rune) bool) int {
	return LastIndexFunc((s), f)
}

// Map returns a copy of the string s with all its characters modified
// according to the mapping function. If mapping returns a negative value, the character is
// dropped from the string with no replacement.
func (s String) Map(mapping func(rune) rune) String {
	return String(Map(mapping, (s)))
}

// Repeat returns a new string consisting of count copies of the string s.
//
// It panics if count is negative or if the result of (len(s) * count)
// overflows.
func (s String) Repeat(count int) String {
	return String(Repeat((s), count))
}

// Replace returns a copy of the string s with the first n
// non-overlapping instances of old replaced by new.
// If old is empty, it matches at the beginning of the string
// and after each UTF-8 sequence, yielding up to k+1 replacements
// for a k-rune string.
// If n < 0, there is no limit on the number of replacements.
func (s String) Replace(old String, new String, n int) String {
	return String(Replace((s), (old), (new), n))
}

// ReplaceAll returns a copy of the string s with all
// non-overlapping instances of old replaced by new.
// If old is empty, it matches at the beginning of the string
// and after each UTF-8 sequence, yielding up to k+1 replacements
// for a k-rune string.
func (s String) ReplaceAll(old, new String) String {
	return String(ReplaceAll((s), (old), (new)))
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
func (s String) Split(sep String) []String {
	native := Split((s), (sep))
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
func (s String) SplitAfter(sep String) []String {
	native := SplitAfter((s), (sep))
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
func (s String) SplitAfterN(sep String, n int) []String {
	native := SplitAfterN((s), (sep), n)
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
func (s String) SplitN(sep String, n int) []String {
	native := SplitN((s), (sep), n)
	strs := make([]String, len(native))
	for i, value := range native {
		strs[i] = String(value)
	}
	return strs
}

// ToLower returns s with all Unicode letters mapped to their lower case.
func (s String) ToLower() String {
	return String(ToLower((s)))
}

// ToLowerSpecial returns a copy of the string s with all Unicode letters mapped to their
// lower case using the case mapping specified by c.
func (s String) ToLowerSpecial(c unicode.SpecialCase) String {
	return ToLowerSpecial(c, s)
}

// ToTitle returns a copy of the string s with all Unicode letters mapped to
// their Unicode title case.
func (s String) ToTitle() String {
	return String(ToTitle((s)))
}

// ToTitleSpecial returns a copy of the string s with all Unicode letters mapped to their
// Unicode title case, giving priority to the special casing rules.
func (s String) ToTitleSpecial(c unicode.SpecialCase) String {
	return String(ToTitleSpecial(c, (s)))
}

// ToUpper returns s with all Unicode letters mapped to their upper case.
func (s String) ToUpper() String {
	return String(ToUpper((s)))
}

// ToUpperSpecial returns a copy of the string s with all Unicode letters mapped to their
// upper case using the case mapping specified by c.
func (s String) ToUpperSpecial(c unicode.SpecialCase) String {
	return String(ToUpperSpecial(c, (s)))
}

// ToValidUTF8 returns a copy of the string s with each run of invalid UTF-8 byte sequences
// replaced by the replacement string, which may be empty.
func (s String) ToValidUTF8(replacement String) String {
	return String(ToValidUTF8((s), (replacement)))
}

// Trim returns a slice of the string s with all leading and
// trailing Unicode code points contained in cutset removed.
func (s String) Trim(cutset String) String {
	return String(Trim((s), (cutset)))
}

// TrimFunc returns a slice of the string s with all leading
// and trailing Unicode code points c satisfying f(c) removed.
func (s String) TrimFunc(f func(rune) bool) String {
	return String(TrimFunc((s), f))
}

// TrimLeft returns a slice of the string s with all leading
// Unicode code points contained in cutset removed.
//
// To remove a prefix, use [TrimPrefix] instead.
func (s String) TrimLeft(cutset String) String {
	return String(TrimLeft((s), (cutset)))
}

// TrimLeftFunc returns a slice of the string s with all leading
// Unicode code points c satisfying f(c) removed.
func (s String) TrimLeftFunc(f func(rune) bool) String {
	return String(TrimLeftFunc((s), f))
}

// TrimPrefix returns s without the provided leading prefix string.
// If s doesn't start with prefix, s is returned unchanged.
func (s String) TrimPrefix(prefix String) String {
	return String(TrimPrefix((s), (prefix)))
}

// TrimRight returns a slice of the string s, with all trailing
// Unicode code points contained in cutset removed.
//
// To remove a suffix, use [TrimSuffix] instead.
func (s String) TrimRight(cutset String) String {
	return String(TrimRight((s), (cutset)))
}

// TrimRightFunc returns a slice of the string s with all trailing
// Unicode code points c satisfying f(c) removed.
func (s String) TrimRightFunc(f func(rune) bool) String {
	return String(TrimRightFunc((s), f))
}

// TrimSpace returns a slice of the string s, with all leading
// and trailing white space removed, as defined by Unicode.
func (s String) TrimSpace() String {
	return String(TrimSpace((s)))
}

// TrimSuffix returns s without the provided trailing suffix string.
// If s doesn't end with suffix, s is returned unchanged.
func (s String) TrimSuffix(suffix String) String {
	return String(TrimSuffix((s), (suffix)))
}
