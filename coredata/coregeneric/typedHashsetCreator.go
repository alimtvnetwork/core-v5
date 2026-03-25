package coregeneric

// typedHashsetCreator provides factory methods for Hashset[T].
type typedHashsetCreator[T comparable] struct{}

// Empty creates a zero-capacity Hashset[T].
func (it typedHashsetCreator[T]) Empty() *Hashset[T] {
	return EmptyHashset[T]()
}

// Cap creates a Hashset[T] with pre-allocated capacity.
func (it typedHashsetCreator[T]) Cap(capacity int) *Hashset[T] {
	return NewHashset[T](capacity)
}

// From creates a Hashset[T] from a slice of items.
func (it typedHashsetCreator[T]) From(items []T) *Hashset[T] {
	return HashsetFrom[T](items)
}

// Items creates a Hashset[T] from variadic arguments.
func (it typedHashsetCreator[T]) Items(items ...T) *Hashset[T] {
	return HashsetFrom[T](items)
}

// UsingMap creates a Hashset[T] from an existing map.
func (it typedHashsetCreator[T]) UsingMap(itemsMap map[T]bool) *Hashset[T] {
	return HashsetFromMap[T](itemsMap)
}
