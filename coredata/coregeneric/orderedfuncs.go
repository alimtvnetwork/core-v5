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
	slices.Sort(source.items)

	return source
}

// SortCollectionDesc sorts a Collection[T] in descending order (in-place).
func SortCollectionDesc[T cmp.Ordered](source *Collection[T]) *Collection[T] {
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
	for i, item := range source.items {
		source.items[i] = max(minVal, min(maxVal, item))
	}

	return source
}
