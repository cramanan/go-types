package slices

//
//This package defines Ordered operators as functions for Func methods
//

// The basic '==' operator as a function
func Equal[T ordered](e1, e2 T) bool { return e1 == e2 }

// The basic '!=' operator as a function
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
