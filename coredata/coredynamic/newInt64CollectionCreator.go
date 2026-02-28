package coredynamic

// newInt64CollectionCreator provides factory methods for Int64Collection.
type newInt64CollectionCreator struct{}

func (it newInt64CollectionCreator) Empty() *Int64Collection {
	return EmptyCollection[int64]()
}

func (it newInt64CollectionCreator) Cap(capacity int) *Int64Collection {
	return NewCollection[int64](capacity)
}

func (it newInt64CollectionCreator) From(items []int64) *Int64Collection {
	return CollectionFrom[int64](items)
}

func (it newInt64CollectionCreator) Clone(items []int64) *Int64Collection {
	return CollectionClone[int64](items)
}

func (it newInt64CollectionCreator) Items(items ...int64) *Int64Collection {
	return CollectionFrom[int64](items)
}

func (it newInt64CollectionCreator) LenCap(length, capacity int) *Int64Collection {
	return &Int64Collection{
		items: make([]int64, length, capacity),
	}
}
