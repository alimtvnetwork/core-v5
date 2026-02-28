package coregeneric

// typedLinkedListCreator provides factory methods for LinkedList[T].
type typedLinkedListCreator[T any] struct{}

// Empty creates an empty LinkedList[T].
func (it typedLinkedListCreator[T]) Empty() *LinkedList[T] {
	return EmptyLinkedList[T]()
}

// From creates a LinkedList[T] from a slice of items.
func (it typedLinkedListCreator[T]) From(items []T) *LinkedList[T] {
	return LinkedListFrom[T](items)
}

// Items creates a LinkedList[T] from variadic arguments.
func (it typedLinkedListCreator[T]) Items(items ...T) *LinkedList[T] {
	return LinkedListFrom[T](items)
}
