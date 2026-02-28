package coregeneric

import (
	"fmt"
)

// SimpleSlice is a generic thin slice wrapper.
// It generalizes corestr.SimpleSlice from []string to []T.
type SimpleSlice[T any] []T

// EmptySimpleSlice creates an empty SimpleSlice[T].
func EmptySimpleSlice[T any]() *SimpleSlice[T] {
	s := SimpleSlice[T]([]T{})

	return &s
}

// NewSimpleSlice creates a SimpleSlice[T] with pre-allocated capacity.
func NewSimpleSlice[T any](capacity int) *SimpleSlice[T] {
	s := SimpleSlice[T](make([]T, 0, capacity))

	return &s
}

// SimpleSliceFrom wraps an existing slice (no copy).
func SimpleSliceFrom[T any](items []T) *SimpleSlice[T] {
	s := SimpleSlice[T](items)

	return &s
}

// SimpleSliceClone copies items into a new SimpleSlice[T].
func SimpleSliceClone[T any](items []T) *SimpleSlice[T] {
	s := NewSimpleSlice[T](len(items))

	return s.Adds(items...)
}

// Add appends a single item.
func (it *SimpleSlice[T]) Add(item T) *SimpleSlice[T] {
	*it = append(*it, item)

	return it
}

// AddIf appends the item only if the condition is true.
func (it *SimpleSlice[T]) AddIf(isAdd bool, item T) *SimpleSlice[T] {
	if !isAdd {
		return it
	}

	*it = append(*it, item)

	return it
}

// Adds appends multiple items (variadic).
func (it *SimpleSlice[T]) Adds(items ...T) *SimpleSlice[T] {
	if len(items) == 0 {
		return it
	}

	*it = append(*it, items...)

	return it
}

// AddSlice appends all items from a slice.
func (it *SimpleSlice[T]) AddSlice(items []T) *SimpleSlice[T] {
	if len(items) == 0 {
		return it
	}

	*it = append(*it, items...)

	return it
}

// AddsIf appends items only if the condition is true.
func (it *SimpleSlice[T]) AddsIf(isAdd bool, items ...T) *SimpleSlice[T] {
	if !isAdd {
		return it
	}

	return it.Adds(items...)
}

// AddFunc appends the result of calling the function.
func (it *SimpleSlice[T]) AddFunc(f func() T) *SimpleSlice[T] {
	*it = append(*it, f())

	return it
}

// First returns the first element. Panics on empty slice.
func (it *SimpleSlice[T]) First() T {
	return (*it)[0]
}

// Last returns the last element. Panics on empty slice.
func (it *SimpleSlice[T]) Last() T {
	return (*it)[it.LastIndex()]
}

// FirstOrDefault returns the first element or zero value if empty.
func (it *SimpleSlice[T]) FirstOrDefault() T {
	if it.IsEmpty() {
		var zero T

		return zero
	}

	return it.First()
}

// LastOrDefault returns the last element or zero value if empty.
func (it *SimpleSlice[T]) LastOrDefault() T {
	if it.IsEmpty() {
		var zero T

		return zero
	}

	return it.Last()
}

// Skip returns items after skipping the first n.
func (it *SimpleSlice[T]) Skip(count int) []T {
	if count >= it.Length() {
		return []T{}
	}

	return (*it)[count:]
}

// Take returns the first n items.
func (it *SimpleSlice[T]) Take(count int) []T {
	if count >= it.Length() {
		return *it
	}

	return (*it)[:count]
}

// Length returns the number of elements.
func (it *SimpleSlice[T]) Length() int {
	if it == nil {
		return 0
	}

	return len(*it)
}

// Count is an alias for Length.
func (it *SimpleSlice[T]) Count() int {
	return it.Length()
}

// IsEmpty returns true if the slice has no items.
func (it *SimpleSlice[T]) IsEmpty() bool {
	return it == nil || it.Length() == 0
}

// HasAnyItem returns true if the slice has at least one item.
func (it *SimpleSlice[T]) HasAnyItem() bool {
	return !it.IsEmpty()
}

// HasItems is an alias for HasAnyItem.
func (it *SimpleSlice[T]) HasItems() bool {
	return !it.IsEmpty()
}

// LastIndex returns the index of the last element.
func (it *SimpleSlice[T]) LastIndex() int {
	return it.Length() - 1
}

// HasIndex returns true if the given index is valid.
func (it *SimpleSlice[T]) HasIndex(index int) bool {
	return index >= 0 && it.LastIndex() >= index
}

// Items returns the underlying slice.
func (it *SimpleSlice[T]) Items() []T {
	return *it
}

// InsertAt inserts an item at the given index.
func (it *SimpleSlice[T]) InsertAt(index int, item T) *SimpleSlice[T] {
	s := *it
	s = append(s[:index+1], s[index:]...)
	s[index] = item
	*it = s

	return it
}

// ForEach calls fn for each item with its index.
func (it *SimpleSlice[T]) ForEach(fn func(index int, item T)) {
	for i, item := range *it {
		fn(i, item)
	}
}

// Filter returns a new SimpleSlice containing only items matching the predicate.
func (it *SimpleSlice[T]) Filter(predicate func(T) bool) *SimpleSlice[T] {
	result := EmptySimpleSlice[T]()

	for _, item := range *it {
		if predicate(item) {
			result.Add(item)
		}
	}

	return result
}

// CountFunc counts items matching the predicate.
func (it *SimpleSlice[T]) CountFunc(predicate func(index int, item T) bool) int {
	count := 0

	for i, item := range *it {
		if predicate(i, item) {
			count++
		}
	}

	return count
}

// Clone creates a copy of the SimpleSlice.
func (it *SimpleSlice[T]) Clone() *SimpleSlice[T] {
	if it.IsEmpty() {
		return EmptySimpleSlice[T]()
	}

	return SimpleSliceClone[T](*it)
}

// String returns a string representation.
func (it *SimpleSlice[T]) String() string {
	return fmt.Sprintf("%v", *it)
}
