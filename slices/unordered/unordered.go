package unordered

import "github.com/cramanan/go-types/slices"

type Unordered[T any] slices.Slice[T]

func New[T any](values ...T) Unordered[T] {
	return append(*new(Unordered[T]), values...)
}

func From[O any](s Unordered[O]) Unordered[O] {
	return s
}
