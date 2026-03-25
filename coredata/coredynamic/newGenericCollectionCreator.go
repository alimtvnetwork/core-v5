package coredynamic

// newGenericCollectionCreator provides factory methods for Collection[T].
// A single generic definition covers all primitive type variants.
//
// This is the coredynamic equivalent of coregeneric's typedCollectionCreator[T].
// Use the type-specific fields on newCollectionCreator (e.g., String, Int)
// which are instantiated from this generic creator, or use Any for Collection[any].
type newGenericCollectionCreator[T any] struct{}

// Empty creates a zero-capacity Collection[T].
func (it newGenericCollectionCreator[T]) Empty() *Collection[T] {
	return EmptyCollection[T]()
}

// Cap creates a Collection[T] with pre-allocated capacity.
func (it newGenericCollectionCreator[T]) Cap(capacity int) *Collection[T] {
	return NewCollection[T](capacity)
}

// From wraps an existing slice into a Collection[T] (no copy).
func (it newGenericCollectionCreator[T]) From(items []T) *Collection[T] {
	return CollectionFrom[T](items)
}

// Clone copies items into a new Collection[T].
func (it newGenericCollectionCreator[T]) Clone(items []T) *Collection[T] {
	return CollectionClone[T](items)
}

// Items creates a Collection[T] from variadic arguments.
func (it newGenericCollectionCreator[T]) Items(items ...T) *Collection[T] {
	return CollectionFrom[T](items)
}
