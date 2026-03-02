package coregeneric

// MapCollection transforms a Collection[T] into a Collection[U] using a mapping function.
// This is a package-level function because Go does not allow generic methods
// with additional type parameters on generic types.
func MapCollection[T any, U any](
	source *Collection[T],
	mapper func(T) U,
) *Collection[U] {
	if source == nil || source.IsEmpty() {
		return EmptyCollection[U]()
	}

	result := NewCollection[U](source.Length())

	for _, item := range source.items {
		result.Add(mapper(item))
	}

	return result
}

// FlatMapCollection transforms each element into a slice and flattens the results.
func FlatMapCollection[T any, U any](
	source *Collection[T],
	mapper func(T) []U,
) *Collection[U] {
	if source == nil || source.IsEmpty() {
		return EmptyCollection[U]()
	}

	result := NewCollection[U](source.Length())

	for _, item := range source.items {
		mapped := mapper(item)
		result.AddSlice(mapped)
	}

	return result
}

// ReduceCollection reduces a Collection[T] to a single value of type U.
func ReduceCollection[T any, U any](
	source *Collection[T],
	initial U,
	reducer func(accumulator U, item T) U,
) U {
	if source == nil {
		return initial
	}

	result := initial

	for _, item := range source.items {
		result = reducer(result, item)
	}

	return result
}

// GroupByCollection groups items by a key function, returning a map of key to Collection[T].
func GroupByCollection[T any, K comparable](
	source *Collection[T],
	keyFunc func(T) K,
) map[K]*Collection[T] {
	result := make(map[K]*Collection[T])

	if source == nil {
		return result
	}

	for _, item := range source.items {
		key := keyFunc(item)

		col, exists := result[key]
		if !exists {
			col = EmptyCollection[T]()
			result[key] = col
		}

		col.Add(item)
	}

	return result
}

// ContainsFunc checks if any item in the collection matches the predicate.
// Use this for Collection[T] where T is not comparable.
func ContainsFunc[T any](
	source *Collection[T],
	predicate func(T) bool,
) bool {
	if source == nil {
		return false
	}

	for _, item := range source.items {
		if predicate(item) {
			return true
		}
	}

	return false
}

// IndexOfFunc returns the index of the first item matching the predicate, or -1.
func IndexOfFunc[T any](
	source *Collection[T],
	predicate func(T) bool,
) int {
	if source == nil {
		return -1
	}

	for i, item := range source.items {
		if predicate(item) {
			return i
		}
	}

	return -1
}

// ContainsItem checks if a comparable item exists in the collection.
func ContainsItem[T comparable](
	source *Collection[T],
	item T,
) bool {
	if source == nil {
		return false
	}

	for _, existing := range source.items {
		if existing == item {
			return true
		}
	}

	return false
}

// IndexOfItem returns the index of the first occurrence of item, or -1.
func IndexOfItem[T comparable](
	source *Collection[T],
	item T,
) int {
	if source == nil {
		return -1
	}

	for i, existing := range source.items {
		if existing == item {
			return i
		}
	}

	return -1
}

// Distinct returns a new Collection with duplicate items removed.
// Requires T to be comparable.
func Distinct[T comparable](
	source *Collection[T],
) *Collection[T] {
	if source == nil {
		return EmptyCollection[T]()
	}

	seen := make(map[T]bool)
	result := EmptyCollection[T]()

	for _, item := range source.items {
		if !seen[item] {
			seen[item] = true
			result.Add(item)
		}
	}

	return result
}

// MapSimpleSlice transforms a SimpleSlice[T] into a SimpleSlice[U].
func MapSimpleSlice[T any, U any](
	source *SimpleSlice[T],
	mapper func(T) U,
) *SimpleSlice[U] {
	if source == nil || source.IsEmpty() {
		return EmptySimpleSlice[U]()
	}

	result := NewSimpleSlice[U](source.Length())

	for _, item := range *source {
		result.Add(mapper(item))
	}

	return result
}
