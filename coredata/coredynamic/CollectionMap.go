package coredynamic

// Map transforms a Collection[T] into a Collection[U] using the given function.
// This is a package-level function because Go does not support generic methods
// on generic types with different type parameters.
//
// Usage:
//
//	names := coredynamic.Map(users, func(u User) string { return u.Name })
func Map[T any, U any](
	source *Collection[T],
	transform func(T) U,
) *Collection[U] {
	if source == nil || source.IsEmpty() {
		return EmptyCollection[U]()
	}

	result := NewCollection[U](source.Length())

	for _, item := range source.items {
		result.items = append(result.items, transform(item))
	}

	return result
}

// FlatMap transforms each item into a slice and flattens the results.
//
// Usage:
//
//	allTags := coredynamic.FlatMap(posts, func(p Post) []string { return p.Tags })
func FlatMap[T any, U any](
	source *Collection[T],
	transform func(T) []U,
) *Collection[U] {
	if source == nil || source.IsEmpty() {
		return EmptyCollection[U]()
	}

	result := NewCollection[U](source.Length())

	for _, item := range source.items {
		result.items = append(result.items, transform(item)...)
	}

	return result
}

// Reduce reduces a Collection[T] to a single value of type U.
//
// Usage:
//
//	total := coredynamic.Reduce(prices, 0, func(acc int, p Price) int { return acc + p.Amount })
func Reduce[T any, U any](
	source *Collection[T],
	initial U,
	reducer func(accumulator U, item T) U,
) U {
	if source == nil || source.IsEmpty() {
		return initial
	}

	result := initial

	for _, item := range source.items {
		result = reducer(result, item)
	}

	return result
}
