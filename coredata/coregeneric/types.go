package coregeneric

// =============================================================================
// Collection type aliases for common primitive types
// =============================================================================

type StringCollection = Collection[string]
type IntCollection = Collection[int]
type Int8Collection = Collection[int8]
type Int16Collection = Collection[int16]
type Int32Collection = Collection[int32]
type Int64Collection = Collection[int64]
type UintCollection = Collection[uint]
type Uint8Collection = Collection[uint8]
type Uint16Collection = Collection[uint16]
type Uint32Collection = Collection[uint32]
type Uint64Collection = Collection[uint64]
type Float32Collection = Collection[float32]
type Float64Collection = Collection[float64]
type ByteCollection = Collection[byte]
type BoolCollection = Collection[bool]
type AnyCollection = Collection[any]

// =============================================================================
// Hashset type aliases for common primitive types
// =============================================================================

type StringHashset = Hashset[string]
type IntHashset = Hashset[int]
type Int8Hashset = Hashset[int8]
type Int16Hashset = Hashset[int16]
type Int32Hashset = Hashset[int32]
type Int64Hashset = Hashset[int64]
type UintHashset = Hashset[uint]
type Uint8Hashset = Hashset[uint8]
type Uint16Hashset = Hashset[uint16]
type Uint32Hashset = Hashset[uint32]
type Uint64Hashset = Hashset[uint64]
type Float32Hashset = Hashset[float32]
type Float64Hashset = Hashset[float64]
type ByteHashset = Hashset[byte]

// =============================================================================
// Hashmap type aliases for common key-value combinations
// =============================================================================

type StringStringHashmap = Hashmap[string, string]
type StringIntHashmap = Hashmap[string, int]
type StringInt64Hashmap = Hashmap[string, int64]
type StringFloat64Hashmap = Hashmap[string, float64]
type StringBoolHashmap = Hashmap[string, bool]
type StringAnyHashmap = Hashmap[string, any]
type IntStringHashmap = Hashmap[int, string]
type IntIntHashmap = Hashmap[int, int]
type IntAnyHashmap = Hashmap[int, any]

// =============================================================================
// SimpleSlice type aliases for common primitive types
// =============================================================================

type StringSimpleSlice = SimpleSlice[string]
type IntSimpleSlice = SimpleSlice[int]
type Int8SimpleSlice = SimpleSlice[int8]
type Int16SimpleSlice = SimpleSlice[int16]
type Int32SimpleSlice = SimpleSlice[int32]
type Int64SimpleSlice = SimpleSlice[int64]
type UintSimpleSlice = SimpleSlice[uint]
type Float32SimpleSlice = SimpleSlice[float32]
type Float64SimpleSlice = SimpleSlice[float64]
type ByteSimpleSlice = SimpleSlice[byte]
type BoolSimpleSlice = SimpleSlice[bool]
type AnySimpleSlice = SimpleSlice[any]

// =============================================================================
// LinkedList type aliases for common primitive types
// =============================================================================

type StringLinkedList = LinkedList[string]
type IntLinkedList = LinkedList[int]
type Int64LinkedList = LinkedList[int64]
type Float64LinkedList = LinkedList[float64]
type AnyLinkedList = LinkedList[any]
