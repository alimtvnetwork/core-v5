package coregeneric

// newHashsetCreator aggregates per-type hashset creators.
//
// Usage:
//
//	coregeneric.New.Hashset.String.Cap(10)
//	coregeneric.New.Hashset.Int.From([]int{1, 2, 3})
//	coregeneric.New.Hashset.Float64.Items(1.0, 2.5)
type newHashsetCreator struct {
	String  typedHashsetCreator[string]
	Int     typedHashsetCreator[int]
	Int8    typedHashsetCreator[int8]
	Int16   typedHashsetCreator[int16]
	Int32   typedHashsetCreator[int32]
	Int64   typedHashsetCreator[int64]
	Uint    typedHashsetCreator[uint]
	Uint8   typedHashsetCreator[uint8]
	Uint16  typedHashsetCreator[uint16]
	Uint32  typedHashsetCreator[uint32]
	Uint64  typedHashsetCreator[uint64]
	Float32 typedHashsetCreator[float32]
	Float64 typedHashsetCreator[float64]
	Byte    typedHashsetCreator[byte]
}
