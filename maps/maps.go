package maps

// Map is a generic type that wraps a map with keys of type K and values of type V.
type Map[K comparable, V any] map[K]V

// New allocates and initializes a map with keys of type K and values of type V.
func New[K comparable, V any]() Map[K, V] { return make(Map[K, V]) }

// From converts a map with keys of type K and values of type V into a Map type.
func From[M ~map[K]V, K comparable, V any](m M) Map[K, V] { return Map[K, V](m) }
