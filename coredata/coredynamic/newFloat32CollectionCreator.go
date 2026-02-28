package coredynamic

// newFloat32CollectionCreator provides factory methods for Float32Collection.
type newFloat32CollectionCreator struct{}

func (it newFloat32CollectionCreator) Empty() *Float32Collection {
	return EmptyCollection[float32]()
}

func (it newFloat32CollectionCreator) Cap(capacity int) *Float32Collection {
	return NewCollection[float32](capacity)
}

func (it newFloat32CollectionCreator) From(items []float32) *Float32Collection {
	return CollectionFrom[float32](items)
}

func (it newFloat32CollectionCreator) Clone(items []float32) *Float32Collection {
	return CollectionClone[float32](items)
}

func (it newFloat32CollectionCreator) Items(items ...float32) *Float32Collection {
	return CollectionFrom[float32](items)
}
