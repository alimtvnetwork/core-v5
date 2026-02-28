package coredynamic

// newAnyCollectionCreator provides factory methods for Collection[any].
type newAnyCollectionCreator struct{}

// Empty creates a zero-capacity Collection[any].
func (it newAnyCollectionCreator) Empty(zero ...any) *Collection[any] {
	return EmptyCollection[any]()
}

// Cap creates a Collection[any] with pre-allocated capacity.
func (it newAnyCollectionCreator) Cap(capacity int) *Collection[any] {
	return NewCollection[any](capacity)
}

// From wraps an existing slice into a Collection[any] (no copy).
func (it newAnyCollectionCreator) From(items []any) *Collection[any] {
	return CollectionFrom[any](items)
}

// Clone copies items into a new Collection[any].
func (it newAnyCollectionCreator) Clone(items []any) *Collection[any] {
	return CollectionClone[any](items)
}

// Items creates a Collection[any] from variadic arguments.
func (it newAnyCollectionCreator) Items(items ...any) *Collection[any] {
	return CollectionFrom[any](items)
}
