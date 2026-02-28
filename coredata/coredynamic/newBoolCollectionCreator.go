package coredynamic

// newBoolCollectionCreator provides factory methods for BoolCollection.
type newBoolCollectionCreator struct{}

func (it newBoolCollectionCreator) Empty() *BoolCollection {
	return EmptyCollection[bool]()
}

func (it newBoolCollectionCreator) Cap(capacity int) *BoolCollection {
	return NewCollection[bool](capacity)
}

func (it newBoolCollectionCreator) From(items []bool) *BoolCollection {
	return CollectionFrom[bool](items)
}

func (it newBoolCollectionCreator) Clone(items []bool) *BoolCollection {
	return CollectionClone[bool](items)
}

func (it newBoolCollectionCreator) Items(items ...bool) *BoolCollection {
	return CollectionFrom[bool](items)
}
