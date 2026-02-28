package coredynamic

// newInt64CollectionCreator embeds the generic creator and adds LenCap.
type newInt64CollectionCreator struct {
	newGenericCollectionCreator[int64]
}

// LenCap creates a collection with specific length and capacity.
func (it newInt64CollectionCreator) LenCap(length, capacity int) *Int64Collection {
	return &Int64Collection{
		items: make([]int64, length, capacity),
	}
}
