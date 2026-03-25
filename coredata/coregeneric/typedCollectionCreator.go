package coregeneric

// typedCollectionCreator provides factory methods for Collection[T].
// A single generic definition covers all primitive type variants.
type typedCollectionCreator[T any] struct{}

// Empty creates a zero-capacity Collection[T].
func (it typedCollectionCreator[T]) Empty() *Collection[T] {
	return EmptyCollection[T]()
}

// Cap creates a Collection[T] with pre-allocated capacity.
func (it typedCollectionCreator[T]) Cap(capacity int) *Collection[T] {
	return NewCollection[T](capacity)
}

// From wraps an existing slice into a Collection[T] (no copy).
func (it typedCollectionCreator[T]) From(items []T) *Collection[T] {
	return CollectionFrom[T](items)
}

// Clone copies items into a new Collection[T].
func (it typedCollectionCreator[T]) Clone(items []T) *Collection[T] {
	return CollectionClone[T](items)
}

// Items creates a Collection[T] from variadic arguments.
func (it typedCollectionCreator[T]) Items(items ...T) *Collection[T] {
	return CollectionFrom[T](items)
}

// LenCap creates a Collection[T] with specific length and capacity.
func (it typedCollectionCreator[T]) LenCap(length, capacity int) *Collection[T] {
	return CollectionLenCap[T](length, capacity)
}
