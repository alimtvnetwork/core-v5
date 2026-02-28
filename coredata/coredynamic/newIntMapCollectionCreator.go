package coredynamic

// newIntMapCollectionCreator provides factory methods for IntMapCollection.
type newIntMapCollectionCreator struct{}

func (it newIntMapCollectionCreator) Empty() *IntMapCollection {
	return EmptyCollection[map[string]int]()
}

func (it newIntMapCollectionCreator) Cap(capacity int) *IntMapCollection {
	return NewCollection[map[string]int](capacity)
}

func (it newIntMapCollectionCreator) From(items []map[string]int) *IntMapCollection {
	return CollectionFrom[map[string]int](items)
}

func (it newIntMapCollectionCreator) Clone(items []map[string]int) *IntMapCollection {
	return CollectionClone[map[string]int](items)
}

func (it newIntMapCollectionCreator) Items(items ...map[string]int) *IntMapCollection {
	return CollectionFrom[map[string]int](items)
}
