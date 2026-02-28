package coredynamic

// newGenericCollectionCreator provides factory methods for Collection[T] with any type.
type newGenericCollectionCreator struct{}

// Empty creates a zero-capacity Collection[T].
func (it newGenericCollectionCreator) Empty(zero ...any) *Collection[any] {
	return EmptyCollection[any]()
}

// Cap creates a Collection[T] with pre-allocated capacity.
func (it newGenericCollectionCreator) Cap(capacity int) *Collection[any] {
	return NewCollection[any](capacity)
}

// From wraps an existing slice into a Collection[T] (no copy).
func (it newGenericCollectionCreator) From(items []any) *Collection[any] {
	return CollectionFrom[any](items)
}

// Clone copies items into a new Collection[T].
func (it newGenericCollectionCreator) Clone(items []any) *Collection[any] {
	return CollectionClone[any](items)
}

// Items creates a Collection[T] from variadic arguments.
func (it newGenericCollectionCreator) Items(items ...any) *Collection[any] {
	return CollectionFrom[any](items)
}
