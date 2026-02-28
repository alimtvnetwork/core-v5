package coredynamic

// newStringMapCollectionCreator provides factory methods for StringMapCollection.
type newStringMapCollectionCreator struct{}

func (it newStringMapCollectionCreator) Empty() *StringMapCollection {
	return EmptyCollection[map[string]string]()
}

func (it newStringMapCollectionCreator) Cap(capacity int) *StringMapCollection {
	return NewCollection[map[string]string](capacity)
}

func (it newStringMapCollectionCreator) From(items []map[string]string) *StringMapCollection {
	return CollectionFrom[map[string]string](items)
}

func (it newStringMapCollectionCreator) Clone(items []map[string]string) *StringMapCollection {
	return CollectionClone[map[string]string](items)
}

func (it newStringMapCollectionCreator) Items(items ...map[string]string) *StringMapCollection {
	return CollectionFrom[map[string]string](items)
}
