package strings

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
func From[S ~string](value S) String {
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
Returns the first byte of strings.String

Same as:

	str[0]
*/
func (str String) First() byte {
	return str[0]
}

/*
Returns the last byte of strings.String

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
