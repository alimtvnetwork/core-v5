package coredynamic

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

// --- Conditional Add ---

// AddIf appends an item only when isAdd is true.
func (it *Collection[T]) AddIf(isAdd bool, item T) *Collection[T] {
	if !isAdd {
		return it
	}
	it.items = append(it.items, item)
	return it
}

// AddManyIf appends multiple items only when isAdd is true.
func (it *Collection[T]) AddManyIf(isAdd bool, items ...T) *Collection[T] {
	if !isAdd || len(items) == 0 {
		return it
	}
	it.items = append(it.items, items...)
	return it
}

// --- Collection Merging ---

// AddCollection appends all items from another collection. Skips nil/empty.
func (it *Collection[T]) AddCollection(other *Collection[T]) *Collection[T] {
	if other == nil || other.IsEmpty() {
		return it
	}
	it.items = append(it.items, other.items...)
	return it
}

// AddCollections appends items from multiple collections. Skips nil/empty.
func (it *Collection[T]) AddCollections(others ...*Collection[T]) *Collection[T] {
	for _, other := range others {
		if other == nil || other.IsEmpty() {
			continue
		}
		it.items = append(it.items, other.items...)
	}
	return it
}

// ConcatNew creates a new collection containing this collection's items
// plus the additional items. Does not mutate the original.
func (it *Collection[T]) ConcatNew(items ...T) *Collection[T] {
	newLen := it.Length() + len(items)
	result := NewCollection[T](newLen)
	result.items = append(result.items, it.items...)
	result.items = append(result.items, items...)
	return result
}

// --- Clone ---

// Clone creates a deep copy of the collection.
func (it *Collection[T]) Clone() *Collection[T] {
	if it == nil {
		return EmptyCollection[T]()
	}
	return CollectionClone[T](it.items)
}

// --- Capacity ---

// Capacity returns the current capacity of the underlying slice.
func (it *Collection[T]) Capacity() int {
	if it == nil || it.items == nil {
		return 0
	}
	return cap(it.items)
}

// AddCapacity grows the collection's capacity by the specified amount.
func (it *Collection[T]) AddCapacity(additional int) *Collection[T] {
	if additional <= 0 {
		return it
	}
	return it.Resize(it.Capacity() + additional)
}

// Resize grows the collection to the new capacity if it's larger than the current.
func (it *Collection[T]) Resize(newCapacity int) *Collection[T] {
	if cap(it.items) >= newCapacity {
		return it
	}
	newItems := make([]T, len(it.items), newCapacity)
	copy(newItems, it.items)
	it.items = newItems
	return it
}

// --- Reorder ---

// Reverse reverses the items in place.
func (it *Collection[T]) Reverse() *Collection[T] {
	length := it.Length()
	if length <= 1 {
		return it
	}

	lastIndex := length - 1
	mid := length / 2

	for i := 0; i < mid; i++ {
		it.items[i], it.items[lastIndex-i] =
			it.items[lastIndex-i], it.items[i]
	}
	return it
}

// InsertAt inserts items at the given index.
func (it *Collection[T]) InsertAt(index int, items ...T) *Collection[T] {
	if len(items) == 0 {
		return it
	}

	// Grow slice
	it.items = append(it.items, items...)

	// Shift existing items right
	copy(it.items[index+len(items):], it.items[index:])

	// Place new items
	copy(it.items[index:], items)

	return it
}

// --- Search ---

// IndexOfFunc returns the index of the first item matching the predicate, or -1.
func (it *Collection[T]) IndexOfFunc(
	predicate func(T) bool,
) int {
	for i, item := range it.items {
		if predicate(item) {
			return i
		}
	}
	return constants.InvalidIndex
}

// ContainsFunc returns true if any item matches the predicate.
func (it *Collection[T]) ContainsFunc(
	predicate func(T) bool,
) bool {
	return it.IndexOfFunc(predicate) >= 0
}

// Note: Map, FlatMap, and Reduce are defined in CollectionMap.go

// --- String Helpers ---

// SafeAt returns the item at index, or the zero value if index is out of range.
func (it *Collection[T]) SafeAt(index int) T {
	if !it.HasIndex(index) {
		var zero T
		return zero
	}
	return it.items[index]
}

// SprintItems returns a string slice where each item is formatted via fmt.Sprintf.
func (it *Collection[T]) SprintItems(format string) []string {
	slice := make([]string, it.Length())

	for i, item := range it.items {
		slice[i] = fmt.Sprintf(format, item)
	}
	return slice
}
