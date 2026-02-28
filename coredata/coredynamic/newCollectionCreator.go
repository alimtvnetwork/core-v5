package coredynamic

// newCollectionCreator aggregates per-type collection creators.
//
// Usage:
//
//	coredynamic.New.Collection.String.Cap(10)
//	coredynamic.New.Collection.Int.Empty()
//	coredynamic.New.Collection.AnyMap.From(items)
type newCollectionCreator struct {
	String    newStringCollectionCreator
	Int       newIntCollectionCreator
	Int64     newInt64CollectionCreator
	Byte      newByteCollectionCreator
	ByteSlice newByteSliceCollectionCreator
	Bool      newBoolCollectionCreator
	Float32   newFloat32CollectionCreator
	Float64   newFloat64CollectionCreator
	AnyMap    newAnyMapCollectionCreator
	StringMap newStringMapCollectionCreator
	IntMap    newIntMapCollectionCreator
}
