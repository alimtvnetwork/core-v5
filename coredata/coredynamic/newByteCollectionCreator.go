package coredynamic

// newByteCollectionCreator provides factory methods for ByteCollection.
type newByteCollectionCreator struct{}

func (it newByteCollectionCreator) Empty() *ByteCollection {
	return EmptyCollection[byte]()
}

func (it newByteCollectionCreator) Cap(capacity int) *ByteCollection {
	return NewCollection[byte](capacity)
}

func (it newByteCollectionCreator) From(items []byte) *ByteCollection {
	return CollectionFrom[byte](items)
}

func (it newByteCollectionCreator) Clone(items []byte) *ByteCollection {
	return CollectionClone[byte](items)
}

func (it newByteCollectionCreator) LenCap(length, capacity int) *ByteCollection {
	return &ByteCollection{
		items: make([]byte, length, capacity),
	}
}
