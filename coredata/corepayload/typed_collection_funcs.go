package corepayload

// MapTypedPayloads transforms each TypedPayloadWrapper[T] into a value of type U.
//
// This is a package-level function because Go does not allow generic methods
// with additional type parameters on generic types.
//
// Usage:
//
//	names := corepayload.MapTypedPayloads[User, string](collection,
//	    func(item *TypedPayloadWrapper[User]) string {
//	        return item.Data().Name
//	    },
//	)
func MapTypedPayloads[T any, U any](
	source *TypedPayloadCollection[T],
	mapper func(*TypedPayloadWrapper[T]) U,
) []U {
	if source.IsEmpty() {
		return []U{}
	}

	result := make([]U, source.Length())

	for i, item := range source.items {
		result[i] = mapper(item)
	}

	return result
}

// MapTypedPayloadData transforms each deserialized data T into a value of type U.
//
// Usage:
//
//	emails := corepayload.MapTypedPayloadData[User, string](collection,
//	    func(user User) string {
//	        return user.Email
//	    },
//	)
func MapTypedPayloadData[T any, U any](
	source *TypedPayloadCollection[T],
	mapper func(T) U,
) []U {
	if source.IsEmpty() {
		return []U{}
	}

	result := make([]U, source.Length())

	for i, item := range source.items {
		result[i] = mapper(item.Data())
	}

	return result
}

// FlatMapTypedPayloads transforms each wrapper into a slice and flattens the results.
//
// Usage:
//
//	allTags := corepayload.FlatMapTypedPayloads[User, string](collection,
//	    func(item *TypedPayloadWrapper[User]) []string {
//	        return item.Data().Tags
//	    },
//	)
func FlatMapTypedPayloads[T any, U any](
	source *TypedPayloadCollection[T],
	mapper func(*TypedPayloadWrapper[T]) []U,
) []U {
	if source.IsEmpty() {
		return []U{}
	}

	var result []U

	for _, item := range source.items {
		result = append(result, mapper(item)...)
	}

	return result
}

// FlatMapTypedPayloadData transforms each data item into a slice and flattens the results.
func FlatMapTypedPayloadData[T any, U any](
	source *TypedPayloadCollection[T],
	mapper func(T) []U,
) []U {
	if source.IsEmpty() {
		return []U{}
	}

	var result []U

	for _, item := range source.items {
		result = append(result, mapper(item.Data())...)
	}

	return result
}

// ReduceTypedPayloads reduces a TypedPayloadCollection[T] to a single value of type U.
//
// Usage:
//
//	totalAge := corepayload.ReduceTypedPayloads[User, int](collection, 0,
//	    func(acc int, item *TypedPayloadWrapper[User]) int {
//	        return acc + item.Data().Age
//	    },
//	)
func ReduceTypedPayloads[T any, U any](
	source *TypedPayloadCollection[T],
	initial U,
	reducer func(accumulator U, item *TypedPayloadWrapper[T]) U,
) U {
	result := initial

	if source.IsEmpty() {
		return result
	}

	for _, item := range source.items {
		result = reducer(result, item)
	}

	return result
}

// ReduceTypedPayloadData reduces the data items to a single value.
func ReduceTypedPayloadData[T any, U any](
	source *TypedPayloadCollection[T],
	initial U,
	reducer func(accumulator U, data T) U,
) U {
	result := initial

	if source.IsEmpty() {
		return result
	}

	for _, item := range source.items {
		result = reducer(result, item.Data())
	}

	return result
}

// GroupTypedPayloads groups items by a key function.
//
// Usage:
//
//	byCategory := corepayload.GroupTypedPayloads[User, string](collection,
//	    func(item *TypedPayloadWrapper[User]) string {
//	        return item.CategoryName()
//	    },
//	)
func GroupTypedPayloads[T any, K comparable](
	source *TypedPayloadCollection[T],
	keyFunc func(*TypedPayloadWrapper[T]) K,
) map[K]*TypedPayloadCollection[T] {
	result := make(map[K]*TypedPayloadCollection[T])

	if source.IsEmpty() {
		return result
	}

	for _, item := range source.items {
		key := keyFunc(item)
		group, exists := result[key]
		isNewGroup := !exists

		if isNewGroup {
			group = EmptyTypedPayloadCollection[T]()
			result[key] = group
		}

		group.Add(item)
	}

	return result
}

// GroupTypedPayloadData groups items by a key derived from the data.
func GroupTypedPayloadData[T any, K comparable](
	source *TypedPayloadCollection[T],
	keyFunc func(T) K,
) map[K]*TypedPayloadCollection[T] {
	return GroupTypedPayloads[T, K](source, func(item *TypedPayloadWrapper[T]) K {
		return keyFunc(item.Data())
	})
}

// PartitionTypedPayloads splits the collection into two: items matching the predicate
// and items that don't.
//
// Usage:
//
//	active, inactive := corepayload.PartitionTypedPayloads[User](collection,
//	    func(item *TypedPayloadWrapper[User]) bool {
//	        return item.Data().IsActive
//	    },
//	)
func PartitionTypedPayloads[T any](
	source *TypedPayloadCollection[T],
	predicate func(*TypedPayloadWrapper[T]) bool,
) (matching, notMatching *TypedPayloadCollection[T]) {
	matching = EmptyTypedPayloadCollection[T]()
	notMatching = EmptyTypedPayloadCollection[T]()

	if source.IsEmpty() {
		return matching, notMatching
	}

	for _, item := range source.items {
		if predicate(item) {
			matching.Add(item)
		} else {
			notMatching.Add(item)
		}
	}

	return matching, notMatching
}

// AnyTypedPayload returns true if any item matches the predicate.
func AnyTypedPayload[T any](
	source *TypedPayloadCollection[T],
	predicate func(*TypedPayloadWrapper[T]) bool,
) bool {
	if source.IsEmpty() {
		return false
	}

	for _, item := range source.items {
		if predicate(item) {
			return true
		}
	}

	return false
}

// AllTypedPayloads returns true if all items match the predicate.
// Returns true for empty collections (vacuous truth).
func AllTypedPayloads[T any](
	source *TypedPayloadCollection[T],
	predicate func(*TypedPayloadWrapper[T]) bool,
) bool {
	if source.IsEmpty() {
		return true
	}

	for _, item := range source.items {
		isFailed := !predicate(item)

		if isFailed {
			return false
		}
	}

	return true
}

// ConvertTypedPayloads transforms a TypedPayloadCollection[T] into
// a TypedPayloadCollection[U] by re-deserializing each wrapper's bytes.
//
// Usage:
//
//	adminCol := corepayload.ConvertTypedPayloads[User, Admin](userCol)
func ConvertTypedPayloads[T any, U any](
	source *TypedPayloadCollection[T],
) (*TypedPayloadCollection[U], error) {
	if source.IsEmpty() {
		return EmptyTypedPayloadCollection[U](), nil
	}

	result := NewTypedPayloadCollection[U](source.Length())

	for _, item := range source.items {
		converted, err := NewTypedPayloadWrapper[U](item.Wrapper)

		if err != nil {
			return nil, err
		}

		result.items = append(result.items, converted)
	}

	return result, nil
}
