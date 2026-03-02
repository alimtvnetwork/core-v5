package coregeneric

import (
	"cmp"
	"slices"
)

// =============================================================================
// Ordered constraint functions for Collection[T]
//
// These functions require T to satisfy cmp.Ordered (int, float, string, etc.)
// and provide type-safe sorting, min, max operations without custom comparators.
// =============================================================================

// SortCollection sorts a Collection[T] in ascending order (in-place).
// Requires T to be cmp.Ordered.
func SortCollection[T cmp.Ordered](source *Collection[T]) *Collection[T] {
	if source == nil {
		return nil
	}

	slices.Sort(source.items)

	return source
}

// SortCollectionDesc sorts a Collection[T] in descending order (in-place).
func SortCollectionDesc[T cmp.Ordered](source *Collection[T]) *Collection[T] {
	if source == nil {
		return nil
	}

	slices.SortFunc(source.items, func(a, b T) int {
		return cmp.Compare(b, a)
	})

	return source
}

// MinCollection returns the minimum element in a Collection[T].
// Panics on empty collection.
func MinCollection[T cmp.Ordered](source *Collection[T]) T {
	return slices.Min(source.items)
}

// MaxCollection returns the maximum element in a Collection[T].
// Panics on empty collection.
func MaxCollection[T cmp.Ordered](source *Collection[T]) T {
	return slices.Max(source.items)
}

// MinCollectionOrDefault returns the minimum element, or defVal if empty.
func MinCollectionOrDefault[T cmp.Ordered](source *Collection[T], defVal T) T {
	if source.IsEmpty() {
		return defVal
	}

	return slices.Min(source.items)
}

// MaxCollectionOrDefault returns the maximum element, or defVal if empty.
func MaxCollectionOrDefault[T cmp.Ordered](source *Collection[T], defVal T) T {
	if source.IsEmpty() {
		return defVal
	}

	return slices.Max(source.items)
}

// IsSortedCollection returns true if the collection is sorted in ascending order.
func IsSortedCollection[T cmp.Ordered](source *Collection[T]) bool {
	if source == nil {
		return true
	}

	return slices.IsSorted(source.items)
}

// SortSimpleSlice sorts a SimpleSlice[T] in ascending order (in-place).
func SortSimpleSlice[T cmp.Ordered](source *SimpleSlice[T]) *SimpleSlice[T] {
	slices.Sort([]T(*source))

	return source
}

// SortSimpleSliceDesc sorts a SimpleSlice[T] in descending order (in-place).
func SortSimpleSliceDesc[T cmp.Ordered](source *SimpleSlice[T]) *SimpleSlice[T] {
	slices.SortFunc([]T(*source), func(a, b T) int {
		return cmp.Compare(b, a)
	})

	return source
}

// MinSimpleSlice returns the minimum element in a SimpleSlice[T].
// Panics on empty slice.
func MinSimpleSlice[T cmp.Ordered](source *SimpleSlice[T]) T {
	return slices.Min([]T(*source))
}

// MaxSimpleSlice returns the maximum element in a SimpleSlice[T].
// Panics on empty slice.
func MaxSimpleSlice[T cmp.Ordered](source *SimpleSlice[T]) T {
	return slices.Max([]T(*source))
}

// SumCollection returns the sum of all elements in a Collection[T].
// Requires T to be a numeric ordered type.
func SumCollection[T cmp.Ordered](source *Collection[T]) T {
	var sum T

	if source == nil {
		return sum
	}

	for _, item := range source.items {
		sum += item
	}

	return sum
}

// SumSimpleSlice returns the sum of all elements in a SimpleSlice[T].
func SumSimpleSlice[T cmp.Ordered](source *SimpleSlice[T]) T {
	var sum T

	for _, item := range *source {
		sum += item
	}

	return sum
}

// ClampCollection clamps all values in the collection to [min, max].
func ClampCollection[T cmp.Ordered](source *Collection[T], minVal, maxVal T) *Collection[T] {
	if source == nil {
		return nil
	}

	for i, item := range source.items {
		source.items[i] = max(minVal, min(maxVal, item))
	}

	return source
}

// =============================================================================
// Ordered constraint functions for Hashset[T]
// =============================================================================

// SortedListHashset returns the hashset items as a sorted ascending slice.
func SortedListHashset[T cmp.Ordered](source *Hashset[T]) []T {
	list := source.List()
	slices.Sort(list)

	return list
}

// SortedListDescHashset returns the hashset items as a sorted descending slice.
func SortedListDescHashset[T cmp.Ordered](source *Hashset[T]) []T {
	list := source.List()
	slices.SortFunc(list, func(a, b T) int {
		return cmp.Compare(b, a)
	})

	return list
}

// MinHashset returns the minimum element in a Hashset[T].
// Panics on empty hashset.
func MinHashset[T cmp.Ordered](source *Hashset[T]) T {
	return slices.Min(source.List())
}

// MaxHashset returns the maximum element in a Hashset[T].
// Panics on empty hashset.
func MaxHashset[T cmp.Ordered](source *Hashset[T]) T {
	return slices.Max(source.List())
}

// MinHashsetOrDefault returns the minimum element, or defVal if empty.
func MinHashsetOrDefault[T cmp.Ordered](source *Hashset[T], defVal T) T {
	if source.IsEmpty() {
		return defVal
	}

	return slices.Min(source.List())
}

// MaxHashsetOrDefault returns the maximum element, or defVal if empty.
func MaxHashsetOrDefault[T cmp.Ordered](source *Hashset[T], defVal T) T {
	if source.IsEmpty() {
		return defVal
	}

	return slices.Max(source.List())
}

// SortedCollectionHashset returns the hashset as a sorted Collection[T].
func SortedCollectionHashset[T cmp.Ordered](source *Hashset[T]) *Collection[T] {
	sorted := SortedListHashset(source)

	return CollectionFrom[T](sorted)
}

// =============================================================================
// Ordered constraint functions for Hashmap[K, V]
// =============================================================================

// SortedKeysHashmap returns all keys sorted in ascending order.
// Requires K to be cmp.Ordered.
func SortedKeysHashmap[K cmp.Ordered, V any](source *Hashmap[K, V]) []K {
	keys := source.Keys()
	slices.Sort(keys)

	return keys
}

// SortedKeysDescHashmap returns all keys sorted in descending order.
func SortedKeysDescHashmap[K cmp.Ordered, V any](source *Hashmap[K, V]) []K {
	keys := source.Keys()
	slices.SortFunc(keys, func(a, b K) int {
		return cmp.Compare(b, a)
	})

	return keys
}

// MinKeyHashmap returns the minimum key.
// Panics on empty hashmap.
func MinKeyHashmap[K cmp.Ordered, V any](source *Hashmap[K, V]) K {
	return slices.Min(source.Keys())
}

// MaxKeyHashmap returns the maximum key.
// Panics on empty hashmap.
func MaxKeyHashmap[K cmp.Ordered, V any](source *Hashmap[K, V]) K {
	return slices.Max(source.Keys())
}

// MinKeyHashmapOrDefault returns the minimum key, or defVal if empty.
func MinKeyHashmapOrDefault[K cmp.Ordered, V any](source *Hashmap[K, V], defVal K) K {
	if source.IsEmpty() {
		return defVal
	}

	return slices.Min(source.Keys())
}

// MaxKeyHashmapOrDefault returns the maximum key, or defVal if empty.
func MaxKeyHashmapOrDefault[K cmp.Ordered, V any](source *Hashmap[K, V], defVal K) K {
	if source.IsEmpty() {
		return defVal
	}

	return slices.Max(source.Keys())
}

// SortedValuesHashmap returns all values sorted in ascending order.
// Requires V to be cmp.Ordered.
func SortedValuesHashmap[K comparable, V cmp.Ordered](source *Hashmap[K, V]) []V {
	values := source.Values()
	slices.Sort(values)

	return values
}

// MinValueHashmap returns the minimum value.
// Requires V to be cmp.Ordered. Panics on empty hashmap.
func MinValueHashmap[K comparable, V cmp.Ordered](source *Hashmap[K, V]) V {
	return slices.Min(source.Values())
}

// MaxValueHashmap returns the maximum value.
// Requires V to be cmp.Ordered. Panics on empty hashmap.
func MaxValueHashmap[K comparable, V cmp.Ordered](source *Hashmap[K, V]) V {
	return slices.Max(source.Values())
}

// MinValueHashmapOrDefault returns the minimum value, or defVal if empty.
func MinValueHashmapOrDefault[K comparable, V cmp.Ordered](source *Hashmap[K, V], defVal V) V {
	if source.IsEmpty() {
		return defVal
	}

	return slices.Min(source.Values())
}

// MaxValueHashmapOrDefault returns the maximum value, or defVal if empty.
func MaxValueHashmapOrDefault[K comparable, V cmp.Ordered](source *Hashmap[K, V], defVal V) V {
	if source.IsEmpty() {
		return defVal
	}

	return slices.Max(source.Values())
}
