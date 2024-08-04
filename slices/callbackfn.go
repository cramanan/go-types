package slices

// The basic '==' operator as a function
func Equal[T ordered](e1, e2 T) bool { return e1 == e2 }

// The basic '!=' operator as a function
//
// Warning: NotEqual is ineffective when comparing 2 empty Slices
func NotEqual[T ordered](e1, e2 T) bool { return e1 != e2 }

// The basic '<' operator as a function
func Less[T ordered](e1 T, e2 T) bool { return e1 < e2 }

// The basic '<=' operator as a function
func LessOrEqual[T ordered](e1 T, e2 T) bool { return e1 <= e2 }

// The basic '>' operator as a function
func Greater[T ordered](e1 T, e2 T) bool { return e1 > e2 }

// The basic '>=' operator as a function
func GreaterOrEqual[T ordered](e1 T, e2 T) bool { return e1 >= e2 }

// CompFunc is the Ordered function for operators
func UseOrdered[T ordered](f func(T, T) bool, value T) func(T) bool {
	return func(comp T) bool {
		return f(comp, value)
	}
}
