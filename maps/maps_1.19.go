//go:build go1.19 && !go1.21

// Package maps defines various functions useful with maps of any type.
package maps

import "golang.org/x/exp/maps"

// Keys returns the keys of the map m.
// The keys will be in an indeterminate order.
func Keys[M ~map[K]V, K comparable, V any](m M) []K { return maps.Keys(m) }

// Values returns the values of the map m.
// The values will be in an indeterminate order.
func Values[M ~map[K]V, K comparable, V any](m M) []V { return maps.Values(m) }

// Equal reports whether two maps contain the same key/value pairs.
// Values are compared using ==.
func Equal[M1, M2 ~map[K]V, K, V comparable](m1 M1, m2 M2) bool { return maps.Equal(m1, m2) }

// EqualFunc is like Equal, but compares values using eq.
// Keys are still compared with ==.
func EqualFunc[M1 ~map[K]V1, M2 ~map[K]V2, K comparable, V1, V2 any](m1 M1, m2 M2, eq func(V1, V2) bool) bool {
	return maps.EqualFunc(m1, m2, eq)
}

// Clear removes all entries from m, leaving it empty.
func Clear[M ~map[K]V, K comparable, V any](m M) { maps.Clear(m) }

// Clone returns a copy of m.  This is a shallow clone:
// the new keys and values are set using ordinary assignment.
func Clone[M ~map[K]V, K comparable, V any](m M) M { return maps.Clone(m) }

// Copy copies all key/value pairs in src adding them to dst.
// When a key in src is already present in dst,
// the value in dst will be overwritten by the value associated
// with the key in src.
func Copy[M1 ~map[K]V, M2 ~map[K]V, K comparable, V any](dst M1, src M2) { maps.Copy(dst, src) }

// DeleteFunc deletes any key/value pairs from m for which del returns true.
func DeleteFunc[M ~map[K]V, K comparable, V any](m M, del func(K, V) bool) { maps.DeleteFunc(m, del) }

// Executes a provided function once per each key/value pair in the Map.
func ForEach[M ~map[K]V, K comparable, V any](m M, callbackFn func(K, V)) {
	if callbackFn == nil {
		panic("callback function is nil")
	}
	for key, value := range m {
		callbackFn(key, value)
	}
}

// Filter returns a new map containing only the key-value pairs for which the callback function returns true.
func Filter[M ~map[K]V, K comparable, V any](m M, callbackFn func(K, V) bool) (filtered M) {
	if callbackFn == nil {
		panic("callback function is nil")
	}
	filtered = make(M)
	for key, value := range m {
		if callbackFn(key, value) {
			filtered[key] = value
		}
	}
	return filtered
}

// Some returns true if at least one key-value pair in the map causes the callback function to return true.
func Some[M ~map[K]V, K comparable, V any](m M, callbackFn func(K, V) bool) bool {
	if callbackFn == nil {
		panic("callback function is nil")
	}
	for key, value := range m {
		if callbackFn(key, value) {
			return true
		}
	}
	return false
}

// Every returns true if all key-value pairs in the map cause the callback function to return true.
func Every[M ~map[K]V, K comparable, V any](m M, callbackFn func(K, V) bool) bool {
	if callbackFn == nil {
		panic("callback function is nil")
	}
	for key, value := range m {
		if !callbackFn(key, value) {
			return false
		}
	}
	return true
}

// MapFunc applies a transformation function to each value in the map and returns a new map with the transformed values.
func MapFunc[M ~map[K]V, K comparable, V any, W any](m M, callbackFn func(K, V) W) (mapped map[K]W) {
	mapped = make(map[K]W)
	for key, value := range m {
		mapped[key] = callbackFn(key, value)
	}
	return mapped
}

// Reduce applies a reduction function to each value in the map and returns a single value.
func Reduce[M ~map[K]V, K comparable, V any, I any](m M, callbackFn func(I, K, V) I, initialValue I) I {
	result := initialValue
	for key, value := range m {
		result = callbackFn(result, key, value)
	}
	return result
}

// Size returns the number of key-value pairs in the map.
func Size[M ~map[K]V, K comparable, V any](m M) int { return len(m) }

// IsEmpty returns true if the map is empty
func IsEmpty[M ~map[K]V, K comparable, V any](m M) bool { return Size(m) == 0 }

// Keys returns the keys of the map m.
// The keys will be in an indeterminate order.
func (m Map[K, V]) Keys() []K { return maps.Keys(m) }

// Values returns the values of the map m.
// The values will be in an indeterminate order.
func (m Map[K, V]) Values() []V { return maps.Values(m) }

// Executes a provided function once per each key/value pair in the Map.
func (m Map[K, V]) ForEach(callbackFn func(K, V)) { ForEach(m, callbackFn) }

// Filter returns a new map containing only the key-value pairs for which the callback function returns true.
func (m Map[K, V]) Filter(callbackFn func(K, V) bool) Map[K, V] { return Filter(m, callbackFn) }

// Some returns true if at least one key-value pair in the map causes the callback function to return true.
func (m Map[K, V]) Some(callbackFn func(K, V) bool) bool { return Some(m, callbackFn) }

// Every returns true if all key-value pairs in the map cause the callback function to return true.
func (m Map[K, V]) Every(callbackFn func(K, V) bool) bool { return Every(m, callbackFn) }

// Len returns the size of the map.
func (m Map[K, V]) Size() int { return Size(m) }

// Len returns the size of the map.
func (m Map[K, V]) IsEmpty() bool { return IsEmpty(m) }
