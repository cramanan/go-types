package slices

type Slice[T any] []T

func New[T any](values ...T) Slice[T] {
	return append(Slice[T]{}, values...)
}

func From[T any](slice []T) Slice[T] {
	return slice
}

func (slice Slice[T]) Prepend(values ...T) Slice[T] {
	return append(values, slice...)
}

func (slice Slice[T]) Append(values ...T) Slice[T] {
	return append(slice, values...)
}

func (slice1 Slice[T]) Merge(slice2 Slice[T]) Slice[T] {
	return append(slice1, slice2...)
}

func (slice Slice[T]) At(n int) T {
	if n < 0 {
		n = len(slice) + n
	}
	return slice[n]
}

func (slice *Slice[T]) Pop() (last T) {
	last = (*slice)[len(*slice)-1]
	*slice = (*slice)[:len(*slice)-1]
	return last
}

func (slice Slice[T]) Length() int {
	return len(slice)
}

func (slice Slice[T]) ForEach(callbackFn func(value T, index int)) {
	for idx, v := range slice {
		callbackFn(v, idx)
	}
}

func (slice Slice[T]) Filter(callbackFn func(element T, index int) bool) (filtered Slice[T]) {
	for idx, value := range slice {
		if callbackFn(value, idx) {
			filtered = append(filtered, value)
		}
	}
	return filtered
}

func (slice Slice[T]) Some(callbackFn func(element T, index int) bool) bool {
	for idx, value := range slice {
		if callbackFn(value, idx) {
			return true
		}
	}
	return false
}

func (slice Slice[T]) Every(callbackFn func(element T, index int) bool) bool {
	for idx, value := range slice {
		if !callbackFn(value, idx) {
			return false
		}
	}
	return true
}

func (slice Slice[T]) Reverse() Slice[T] {
	reversed := make(Slice[T], len(slice))
	for i := len(slice) - 1; i > 0; i-- {
		reversed[i] = slice[i]
	}
	return reversed
}
