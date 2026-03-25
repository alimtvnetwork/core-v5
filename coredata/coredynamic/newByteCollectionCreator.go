package coredynamic

// newByteCollectionCreator embeds the generic creator and adds LenCap.
type newByteCollectionCreator struct {
	newGenericCollectionCreator[byte]
}

// LenCap creates a collection with specific length and capacity.
func (it newByteCollectionCreator) LenCap(length, capacity int) *ByteCollection {
	return &ByteCollection{
		items: make([]byte, length, capacity),
	}
}
