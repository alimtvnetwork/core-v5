package coredynamic

// GroupBy groups items by a key function, returning a map of key to Collection[T].
// The key type K must be comparable for use as a map key.
//
// Usage:
//
//	groups := coredynamic.GroupBy(users, func(u User) string { return u.Department })
func GroupBy[T any, K comparable](
	source *Collection[T],
	keyFunc func(T) K,
) map[K]*Collection[T] {
	if source == nil || source.IsEmpty() {
		return map[K]*Collection[T]{}
	}

	result := make(map[K]*Collection[T])

	for _, item := range source.items {
		key := keyFunc(item)
		col, exists := result[key]
		if !exists {
			col = NewCollection[T](0)
			result[key] = col
		}
		col.items = append(col.items, item)
	}
	return result
}

// GroupByLock is the mutex-protected variant of GroupBy.
func GroupByLock[T any, K comparable](
	source *Collection[T],
	keyFunc func(T) K,
) map[K]*Collection[T] {
	source.Lock()
	defer source.Unlock()
	return GroupBy(source, keyFunc)
}

// GroupByCount returns the count of items per group key.
//
// Usage:
//
//	counts := coredynamic.GroupByCount(words, func(w string) string { return w })
func GroupByCount[T any, K comparable](
	source *Collection[T],
	keyFunc func(T) K,
) map[K]int {
	if source == nil || source.IsEmpty() {
		return map[K]int{}
	}

	result := make(map[K]int)

	for _, item := range source.items {
		result[keyFunc(item)]++
	}
	return result
}
