package coregeneric

// typedHashmapCreator provides factory methods for Hashmap[K, V].
type typedHashmapCreator[K comparable, V any] struct{}

// Empty creates a zero-capacity Hashmap[K, V].
func (it typedHashmapCreator[K, V]) Empty() *Hashmap[K, V] {
	return EmptyHashmap[K, V]()
}

// Cap creates a Hashmap[K, V] with pre-allocated capacity.
func (it typedHashmapCreator[K, V]) Cap(capacity int) *Hashmap[K, V] {
	return NewHashmap[K, V](capacity)
}

// From wraps an existing map (no copy).
func (it typedHashmapCreator[K, V]) From(items map[K]V) *Hashmap[K, V] {
	return HashmapFrom[K, V](items)
}

// Clone copies entries from an existing map.
func (it typedHashmapCreator[K, V]) Clone(items map[K]V) *Hashmap[K, V] {
	return HashmapClone[K, V](items)
}
