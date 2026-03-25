package coredynamic

// newCollectionCreator aggregates per-type collection creators.
//
// Simple types use newGenericCollectionCreator[T] directly.
// Types with extra methods (LenCap, Create) embed the generic creator
// and extend it.
//
// Usage:
//
//	coredynamic.New.Collection.String.Cap(10)
//	coredynamic.New.Collection.Int.Empty()
//	coredynamic.New.Collection.Any.From(items)
type newCollectionCreator struct {
	Any       newAnyCollectionCreator
	String    newStringCollectionCreator
	Int       newIntCollectionCreator
	Int64     newInt64CollectionCreator
	Byte      newByteCollectionCreator
	ByteSlice newGenericCollectionCreator[[]byte]
	Bool      newGenericCollectionCreator[bool]
	Float32   newGenericCollectionCreator[float32]
	Float64   newGenericCollectionCreator[float64]
	AnyMap    newGenericCollectionCreator[map[string]any]
	StringMap newGenericCollectionCreator[map[string]string]
	IntMap    newGenericCollectionCreator[map[string]int]
}
