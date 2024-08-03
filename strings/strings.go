package strings

import (
	"strings"
	"unicode"
)

type String string

/*
Returns a new strings.String

Same as:

	var str strings.String = ""
*/
func New() String {
	return String("")
}

/*
Returns a new strings.String from given value of type ~string (string and any aliases).
Specifying the type parameter should be redundent

Same as:

	str := strings.String("Hello World")
*/
func From[S ~string | ~[]byte | ~[]rune](value S) String {
	return String(value)
}

/*
Returns the native string type of the strings.String

Same as:

	str := string(strings.From("Hello World"))
*/
func (str String) String() string {
	return string(str)
}

/*
Returns the length of strings.String

Same as:

	len()
*/
func (str String) Length() int {
	return len(str)
}

/*
Returns the first byte of strings.String.
Will panic if String is empty.

Same as:

	str[0]
*/
func (str String) First() byte {
	return str[0]
}

/*
Returns the last byte of strings.String
Will panic if String is empty.

Same as:

	str[len(str) - 1]
*/
func (str String) Last() byte {
	return str[len(str)-1]
}

/*
Returns the nth byte of strings.String.

Negative indexing is supported (experimental)

Same as:

	str[n]
*/
func (str String) At(n int) byte {
	if n < 0 {
		n = len(str) + n
	}
	return str[n]
}

/*
Returns strings.String as a slice of bytes

Same as:

	[]byte(str)
*/
func (str String) Bytes() []byte {
	return []byte(str)
}

/*
Concatenates multiple strings.String together without modifying the base String.

Same as:

	strs1 + strs2 + ... + strsN
*/
func (str String) Concatenate(strs ...string) String {
	for _, value := range strs {
		str += String(value)
	}
	return str
}

// /*
// Decatenates str2 from str1 together without modifying the base String.

// Exact same as:

// 	strings.TrimSuffix
// */
// func (str1 String) Decatenate(str2 string) String {
// 	return String(strings.TrimSuffix(string(str1), str2))
// }

func (str String) Clone() String {
	return String(strings.Clone(string(str)))
}

func (str1 String) Compare(str2 String) int {
	return strings.Compare(string(str1), string(str2))
}

func (str String) Contains(substr String) bool {
	return strings.Contains(string(str), string(substr))
}

func (str String) ContainsAny(substr String) bool {
	return strings.ContainsAny(string(str), string(substr))
}

func (str String) ContainsRune(substr rune) bool {
	return strings.ContainsRune(string(str), substr)
}

func (str String) Count(substr String) int {
	return strings.Count(string(str), string(substr))
}

func (str String) Cut(sep String) (before String, after String, found bool) {
	bf, af, found := strings.Cut(string(str), string(sep))
	return String(bf), String(af), found
}

func (str String) CutPrefix(prefix String) (after String, found bool) {
	af, found := strings.CutPrefix(string(str), string(prefix))
	return String(af), found
}

func (str String) CutSuffix(prefix String) (after String, found bool) {
	af, found := strings.CutSuffix(string(str), string(prefix))
	return String(af), found
}

func (str String) EqualFold(t String) bool {
	return strings.EqualFold(string(str), string(t))
}

func (str String) Fields() (fields []String) {
	native := strings.Fields(string(str))
	for _, value := range native {
		fields = append(fields, String(value))
	}
	return fields
}

func (str String) FieldsFunc(f func(rune) bool) (fields []String) {
	native := strings.FieldsFunc(string(str), f)
	fields = make([]String, len(native))
	for i, value := range native {
		fields[i] = String(value)
	}
	return fields
}

func (str String) HasPrefix(prefix String) bool {
	return strings.HasPrefix(string(str), string(prefix))
}

func (str String) HasSuffix(prefix String) bool {
	return strings.HasSuffix(string(str), string(prefix))
}

func (str String) Index(substr String) int {
	return strings.Index(string(str), string(substr))
}

func (str String) IndexAny(substr String) int {
	return strings.IndexAny(string(str), string(substr))
}

func (str String) IndexByte(c byte) int {
	return strings.IndexByte(string(str), c)
}
func (str String) IndexFunc(f func(rune) bool) int {
	return strings.IndexFunc(string(str), f)
}

func (str String) IndexRune(substr rune) int {
	return strings.IndexRune(string(str), substr)
}

func Join(elems []String, sep String) String {
	strElems := make([]string, len(elems))
	for i, elem := range elems {
		strElems[i] = string(elem)
	}
	return String(strings.Join(strElems, string(sep)))
}

func (str String) LastIndex(substr String) int {
	return strings.LastIndex(string(str), string(substr))
}

func (str String) LastIndexAny(chars String) int {
	return strings.LastIndexAny(string(str), string(chars))
}

func (str String) LastIndexByte(c byte) int {
	return strings.LastIndexByte(string(str), c)
}

func (str String) LastIndexFunc(f func(rune) bool) int {
	return strings.LastIndexFunc(string(str), f)
}

func (str String) Map(mapping func(rune) rune) String {
	return String(strings.Map(mapping, string(str)))
}

func (str String) Repeat(count int) String {
	return String(strings.Repeat(string(str), count))
}

func (str String) Replace(old String, new String, n int) String {
	return String(strings.Replace(string(str), string(old), string(new), n))
}

func (str String) ReplaceAll(old, new String) String {
	return str.Replace(old, new, -1)
}

func (str String) Split(sep String) []String {
	native := strings.Split(string(str), string(sep))
	strs := make([]String, len(native))
	for i, value := range native {
		strs[i] = String(value)
	}
	return strs
}

func (str String) SplitAfter(sep String) []String {
	native := strings.SplitAfter(string(str), string(sep))
	strs := make([]String, len(native))
	for i, value := range native {
		strs[i] = String(value)
	}
	return strs
}

func (str String) SplitAfterN(sep String, n int) []String {
	native := strings.SplitAfterN(string(str), string(sep), n)
	strs := make([]String, len(native))
	for i, value := range native {
		strs[i] = String(value)
	}
	return strs
}

func (str String) SplitN(sep String, n int) []String {
	native := strings.SplitN(string(str), string(sep), n)
	strs := make([]String, len(native))
	for i, value := range native {
		strs[i] = String(value)
	}
	return strs
}

func (str String) ToLower() String {
	return String(strings.ToLower(string(str)))
}

func (str String) ToLowerSpecial(c unicode.SpecialCase) String {
	return String(strings.ToLowerSpecial(c, string(str)))
}

func (str String) ToTitle() String {
	return String(strings.ToTitle(string(str)))
}

func (str String) ToTitleSpecial(c unicode.SpecialCase) String {
	return String(strings.ToTitleSpecial(c, string(str)))
}

func (str String) ToUpper() String {
	return String(strings.ToUpper(string(str)))
}

func (str String) ToUpperSpecial(c unicode.SpecialCase) String {
	return String(strings.ToUpperSpecial(c, string(str)))
}

func (str String) ToValidUTF8(replacement String) String {
	return String(strings.ToValidUTF8(string(str), string(replacement)))
}

func (str String) Trim(cutset String) String {
	return String(strings.Trim(string(str), string(cutset)))
}

func (str String) TrimFunc(s String, f func(rune) bool) String {
	return String(strings.TrimFunc(string(str), f))
}

func (str String) TrimLeft(cutset String) String {
	return String(strings.TrimLeft(string(str), string(cutset)))
}

func (str String) TrimLeftFunc(s String, f func(rune) bool) String {
	return String(strings.TrimLeftFunc(string(str), f))
}

func (str String) TrimPrefix(prefix String) String {
	return String(strings.TrimPrefix(string(str), string(prefix)))
}

func (str String) TrimRight(cutset String) String {
	return String(strings.TrimRight(string(str), string(cutset)))
}

func (str String) TrimRightFunc(s String, f func(rune) bool) String {
	return String(strings.TrimRightFunc(string(str), f))
}

func (str String) TrimSpace() String {
	return String(strings.TrimSpace(string(str)))
}

func (str String) TrimSuffix(suffix String) String {
	return String(strings.TrimSuffix(string(str), string(suffix)))
}
