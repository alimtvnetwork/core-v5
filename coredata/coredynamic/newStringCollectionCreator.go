package coredynamic

// newStringCollectionCreator embeds the generic creator and adds
// string-specific convenience methods.
type newStringCollectionCreator struct {
	newGenericCollectionCreator[string]
}

// Create is an alias for From.
func (it newStringCollectionCreator) Create(items []string) *StringCollection {
	return CollectionFrom[string](items)
}

// LenCap creates a collection with specific length and capacity.
func (it newStringCollectionCreator) LenCap(length, capacity int) *StringCollection {
	return &StringCollection{
		items: make([]string, length, capacity),
	}
}
