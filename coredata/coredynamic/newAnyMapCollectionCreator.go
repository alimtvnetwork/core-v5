package coredynamic

// newAnyMapCollectionCreator provides factory methods for AnyMapCollection.
type newAnyMapCollectionCreator struct{}

func (it newAnyMapCollectionCreator) Empty() *AnyMapCollection {
	return EmptyCollection[map[string]any]()
}

func (it newAnyMapCollectionCreator) Cap(capacity int) *AnyMapCollection {
	return NewCollection[map[string]any](capacity)
}

func (it newAnyMapCollectionCreator) From(items []map[string]any) *AnyMapCollection {
	return CollectionFrom[map[string]any](items)
}

func (it newAnyMapCollectionCreator) Clone(items []map[string]any) *AnyMapCollection {
	return CollectionClone[map[string]any](items)
}

func (it newAnyMapCollectionCreator) Items(items ...map[string]any) *AnyMapCollection {
	return CollectionFrom[map[string]any](items)
}
