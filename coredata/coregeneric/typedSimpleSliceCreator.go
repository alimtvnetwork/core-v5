package coregeneric

// typedSimpleSliceCreator provides factory methods for SimpleSlice[T].
type typedSimpleSliceCreator[T any] struct{}

// Empty creates an empty SimpleSlice[T].
func (it typedSimpleSliceCreator[T]) Empty() *SimpleSlice[T] {
	return EmptySimpleSlice[T]()
}

// Cap creates a SimpleSlice[T] with pre-allocated capacity.
func (it typedSimpleSliceCreator[T]) Cap(capacity int) *SimpleSlice[T] {
	return NewSimpleSlice[T](capacity)
}

// From wraps an existing slice (no copy).
func (it typedSimpleSliceCreator[T]) From(items []T) *SimpleSlice[T] {
	return SimpleSliceFrom[T](items)
}

// Clone copies items into a new SimpleSlice[T].
func (it typedSimpleSliceCreator[T]) Clone(items []T) *SimpleSlice[T] {
	return SimpleSliceClone[T](items)
}

// Items creates a SimpleSlice[T] from variadic arguments.
func (it typedSimpleSliceCreator[T]) Items(items ...T) *SimpleSlice[T] {
	return SimpleSliceFrom[T](items)
}
