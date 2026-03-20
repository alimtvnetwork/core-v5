package coregeneric

import "iter"

// All returns an iter.Seq2[int, T] over all nodes with index.
//
//	for i, value := range linkedList.All() { ... }
func (it *LinkedList[T]) All() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		i := 0
		current := it.head

		for current != nil {
			if !yield(i, current.Element) {
				return
			}

			current = current.next
			i++
		}
	}
}

// Values returns an iter.Seq[T] over all node values.
//
//	for value := range linkedList.Values() { ... }
func (it *LinkedList[T]) Values() iter.Seq[T] {
	return func(yield func(T) bool) {
		current := it.head

		for current != nil {
			if !yield(current.Element) {
				return
			}

			current = current.next
		}
	}
}
