package coregeneric

import (
	"cmp"
	"sort"
)

// LessFunc is a comparator function for pointer elements.
// It receives two non-nil dereferenced values and returns true if a should come before b.
type LessFunc[T any] func(a, b T) bool

// PointerSliceSorter is a generic, nil-safe sorter for []*T slices.
// It implements sort.Interface and allows swapping the Less function at any time.
//
// By default, ascending order is used for cmp.Ordered types.
// Nil pointers are consistently pushed to the end (ascending) or beginning (descending).
type PointerSliceSorter[T cmp.Ordered] struct {
	items   []*T
	lessFunc LessFunc[T]
	nilFirst bool
}

// NewPointerSliceSorterAsc creates an ascending sorter for []*T.
// Nil pointers are sorted to the end.
func NewPointerSliceSorterAsc[T cmp.Ordered](items []*T) *PointerSliceSorter[T] {
	return &PointerSliceSorter[T]{
		items: items,
		lessFunc: func(a, b T) bool {
			return a < b
		},
		nilFirst: false,
	}
}

// NewPointerSliceSorterDesc creates a descending sorter for []*T.
// Nil pointers are sorted to the end.
func NewPointerSliceSorterDesc[T cmp.Ordered](items []*T) *PointerSliceSorter[T] {
	return &PointerSliceSorter[T]{
		items: items,
		lessFunc: func(a, b T) bool {
			return a > b
		},
		nilFirst: false,
	}
}

// NewPointerSliceSorterFunc creates a sorter with a custom Less function.
// nilFirst controls where nil pointers end up: true = beginning, false = end.
func NewPointerSliceSorterFunc[T cmp.Ordered](items []*T, lessFunc LessFunc[T], nilFirst bool) *PointerSliceSorter[T] {
	return &PointerSliceSorter[T]{
		items:    items,
		lessFunc: lessFunc,
		nilFirst: nilFirst,
	}
}

// SetLessFunc replaces the comparator function. Returns the sorter for chaining.
func (s *PointerSliceSorter[T]) SetLessFunc(lessFunc LessFunc[T]) *PointerSliceSorter[T] {
	s.lessFunc = lessFunc

	return s
}

// SetNilFirst controls nil placement. true = nils first, false = nils last.
func (s *PointerSliceSorter[T]) SetNilFirst(nilFirst bool) *PointerSliceSorter[T] {
	s.nilFirst = nilFirst

	return s
}

// SetAsc switches to ascending order.
func (s *PointerSliceSorter[T]) SetAsc() *PointerSliceSorter[T] {
	s.lessFunc = func(a, b T) bool {
		return a < b
	}

	return s
}

// SetDesc switches to descending order.
func (s *PointerSliceSorter[T]) SetDesc() *PointerSliceSorter[T] {
	s.lessFunc = func(a, b T) bool {
		return a > b
	}

	return s
}

// SetItems replaces the underlying slice. Returns the sorter for chaining.
func (s *PointerSliceSorter[T]) SetItems(items []*T) *PointerSliceSorter[T] {
	s.items = items

	return s
}

// Items returns the underlying slice.
func (s *PointerSliceSorter[T]) Items() []*T {
	return s.items
}

// Sort sorts the slice in place and returns the sorter for chaining.
func (s *PointerSliceSorter[T]) Sort() *PointerSliceSorter[T] {
	sort.Sort(s)

	return s
}

// IsSorted returns true if the slice is already sorted according to the current Less function.
func (s *PointerSliceSorter[T]) IsSorted() bool {
	return sort.IsSorted(s)
}

// --- sort.Interface implementation ---

// Len returns the number of elements.
func (s *PointerSliceSorter[T]) Len() int {
	if s.items == nil {
		return 0
	}

	return len(s.items)
}

// Less compares two elements, handling nil pointers safely.
func (s *PointerSliceSorter[T]) Less(i, j int) bool {
	pi, pj := s.items[i], s.items[j]

	if pi == nil && pj == nil {
		return false
	}

	if pi == nil {
		return s.nilFirst
	}

	if pj == nil {
		return !s.nilFirst
	}

	return s.lessFunc(*pi, *pj)
}

// Swap swaps two elements.
func (s *PointerSliceSorter[T]) Swap(i, j int) {
	s.items[i], s.items[j] = s.items[j], s.items[i]
}
