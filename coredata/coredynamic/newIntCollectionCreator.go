package coredynamic

// newIntCollectionCreator provides factory methods for IntCollection.
type newIntCollectionCreator struct{}

func (it newIntCollectionCreator) Empty() *IntCollection {
	return EmptyCollection[int]()
}

func (it newIntCollectionCreator) Cap(capacity int) *IntCollection {
	return NewCollection[int](capacity)
}

func (it newIntCollectionCreator) From(items []int) *IntCollection {
	return CollectionFrom[int](items)
}

func (it newIntCollectionCreator) Clone(items []int) *IntCollection {
	return CollectionClone[int](items)
}

func (it newIntCollectionCreator) Items(items ...int) *IntCollection {
	return CollectionFrom[int](items)
}

func (it newIntCollectionCreator) LenCap(length, capacity int) *IntCollection {
	return &IntCollection{
		items: make([]int, length, capacity),
	}
}
