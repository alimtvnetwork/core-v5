package coredynamic

import "gitlab.com/auk-go/core/constants"

// --- Package-level search functions for comparable types ---

// Contains returns true if the collection contains the given item.
func Contains[T comparable](col *Collection[T], item T) bool {
	return IndexOf(col, item) >= 0
}

// IndexOf returns the index of the first occurrence of item, or -1.
func IndexOf[T comparable](col *Collection[T], item T) int {
	for i, v := range col.items {
		if v == item {
			return i
		}
	}
	return constants.InvalidIndex
}

// Has is an alias for Contains.
func Has[T comparable](col *Collection[T], item T) bool {
	return Contains(col, item)
}

// HasAll returns true if the collection contains all of the given items.
func HasAll[T comparable](col *Collection[T], items ...T) bool {
	if col.IsEmpty() {
		return false
	}

	for _, item := range items {
		if !Contains(col, item) {
			return false
		}
	}
	return true
}

// LastIndexOf returns the index of the last occurrence of item, or -1.
func LastIndexOf[T comparable](col *Collection[T], item T) int {
	for i := col.Length() - 1; i >= 0; i-- {
		if col.items[i] == item {
			return i
		}
	}
	return constants.InvalidIndex
}

// Count returns the number of occurrences of item in the collection.
func Count[T comparable](col *Collection[T], item T) int {
	n := 0

	for _, v := range col.items {
		if v == item {
			n++
		}
	}
	return n
}

// --- Thread-safe variants ---

// ContainsLock returns true if the collection contains the item, with mutex protection.
func ContainsLock[T comparable](col *Collection[T], item T) bool {
	col.Lock()
	defer col.Unlock()
	return Contains(col, item)
}

// IndexOfLock returns the index of the first occurrence with mutex protection.
func IndexOfLock[T comparable](col *Collection[T], item T) int {
	col.Lock()
	defer col.Unlock()
	return IndexOf(col, item)
}
