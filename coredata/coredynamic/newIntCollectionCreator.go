package coredynamic

// newIntCollectionCreator embeds the generic creator and adds LenCap.
type newIntCollectionCreator struct {
	newGenericCollectionCreator[int]
}

// LenCap creates a collection with specific length and capacity.
func (it newIntCollectionCreator) LenCap(length, capacity int) *IntCollection {
	return &IntCollection{
		items: make([]int, length, capacity),
	}
}
