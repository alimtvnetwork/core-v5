package coredynamic

// newFloat64CollectionCreator provides factory methods for Float64Collection.
type newFloat64CollectionCreator struct{}

func (it newFloat64CollectionCreator) Empty() *Float64Collection {
	return EmptyCollection[float64]()
}

func (it newFloat64CollectionCreator) Cap(capacity int) *Float64Collection {
	return NewCollection[float64](capacity)
}

func (it newFloat64CollectionCreator) From(items []float64) *Float64Collection {
	return CollectionFrom[float64](items)
}

func (it newFloat64CollectionCreator) Clone(items []float64) *Float64Collection {
	return CollectionClone[float64](items)
}

func (it newFloat64CollectionCreator) Items(items ...float64) *Float64Collection {
	return CollectionFrom[float64](items)
}
