package strings

type String string

func New() String {
	return String("")
}

func From[T ~string](s T) String {
	return String(s)
}
