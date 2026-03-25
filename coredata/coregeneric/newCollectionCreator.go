package coregeneric

// newCollectionCreator aggregates per-type collection creators.
//
// Usage:
//
//	coregeneric.New.Collection.String.Cap(10)
//	coregeneric.New.Collection.Int.Empty()
//	coregeneric.New.Collection.Float64.Items(1.0, 2.5)
//	coregeneric.New.Collection.Any.From(existingSlice)
type newCollectionCreator struct {
	String  typedCollectionCreator[string]
	Int     typedCollectionCreator[int]
	Int8    typedCollectionCreator[int8]
	Int16   typedCollectionCreator[int16]
	Int32   typedCollectionCreator[int32]
	Int64   typedCollectionCreator[int64]
	Uint    typedCollectionCreator[uint]
	Uint8   typedCollectionCreator[uint8]
	Uint16  typedCollectionCreator[uint16]
	Uint32  typedCollectionCreator[uint32]
	Uint64  typedCollectionCreator[uint64]
	Float32 typedCollectionCreator[float32]
	Float64 typedCollectionCreator[float64]
	Byte    typedCollectionCreator[byte]
	Bool    typedCollectionCreator[bool]
	Any     typedCollectionCreator[any]
}
