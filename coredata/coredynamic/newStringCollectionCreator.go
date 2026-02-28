package coredynamic

// newStringCollectionCreator provides factory methods for StringCollection.
type newStringCollectionCreator struct{}

// Empty creates a zero-capacity string collection.
func (it newStringCollectionCreator) Empty() *StringCollection {
	return EmptyCollection[string]()
}

// Cap creates a string collection with the given capacity.
func (it newStringCollectionCreator) Cap(capacity int) *StringCollection {
	return NewCollection[string](capacity)
}

// From wraps an existing slice (no copy).
func (it newStringCollectionCreator) From(items []string) *StringCollection {
	return CollectionFrom[string](items)
}

// Clone creates a collection by copying the given slice.
func (it newStringCollectionCreator) Clone(items []string) *StringCollection {
	return CollectionClone[string](items)
}

// Create is an alias for From.
func (it newStringCollectionCreator) Create(items []string) *StringCollection {
	return CollectionFrom[string](items)
}

// Items creates a collection from variadic items.
func (it newStringCollectionCreator) Items(items ...string) *StringCollection {
	return CollectionFrom[string](items)
}

// LenCap creates a collection with specific length and capacity.
func (it newStringCollectionCreator) LenCap(length, capacity int) *StringCollection {
	return &StringCollection{
		items: make([]string, length, capacity),
	}
}
