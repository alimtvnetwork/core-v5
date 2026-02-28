package coredynamic

// newByteSliceCollectionCreator provides factory methods for ByteSliceCollection.
type newByteSliceCollectionCreator struct{}

func (it newByteSliceCollectionCreator) Empty() *ByteSliceCollection {
	return EmptyCollection[[]byte]()
}

func (it newByteSliceCollectionCreator) Cap(capacity int) *ByteSliceCollection {
	return NewCollection[[]byte](capacity)
}

func (it newByteSliceCollectionCreator) From(items [][]byte) *ByteSliceCollection {
	return CollectionFrom[[]byte](items)
}

func (it newByteSliceCollectionCreator) Clone(items [][]byte) *ByteSliceCollection {
	return CollectionClone[[]byte](items)
}
