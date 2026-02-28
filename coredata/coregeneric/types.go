package coregeneric

// =============================================================================
// Generic Foundation Types
// =============================================================================
//
// The coregeneric package is built on five generic base types.
// ALL specialized type aliases below are derived FROM these generics.
//
// Base Types:
//
//	Collection[T any]          — thread-safe slice-backed list
//	Hashset[T comparable]      — thread-safe set (map[T]bool)
//	Hashmap[K comparable, V any] — thread-safe map wrapper
//	SimpleSlice[T any]         — lightweight typed slice wrapper
//	LinkedList[T any]          — singly-linked list with head/tail
//
// Constraint Hierarchy (most restrictive first):
//
//	cmp.Ordered  ⊂  comparable  ⊂  any
//	├── int, int8..64, uint, uint8..64, float32, float64, string
//	│
//	comparable
//	├── all Ordered types + bool + byte + structs (if all fields comparable)
//	│
//	any
//	├── all comparable types + slices, maps, funcs, interfaces
//
// Usage:
//
//	// Generic (custom types):
//	col := coregeneric.NewCollection[MyStruct](10)
//
//	// Typed shorthand (primitives):
//	col := coregeneric.New.Collection.String.Cap(10)
//
//	// Type alias:
//	var col coregeneric.StringCollection
//

// =============================================================================
// Collection[T any] — type aliases for common primitives
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
// Hashset[T comparable] — type aliases for common primitives
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
type BoolHashset = Hashset[bool]

// =============================================================================
// Hashmap[K comparable, V any] — type aliases for common key-value combinations
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
type IntBoolHashmap = Hashmap[int, bool]
type Int64StringHashmap = Hashmap[int64, string]
type Int64AnyHashmap = Hashmap[int64, any]

// =============================================================================
// SimpleSlice[T any] — type aliases for common primitives
// =============================================================================

type StringSimpleSlice = SimpleSlice[string]
type IntSimpleSlice = SimpleSlice[int]
type Int8SimpleSlice = SimpleSlice[int8]
type Int16SimpleSlice = SimpleSlice[int16]
type Int32SimpleSlice = SimpleSlice[int32]
type Int64SimpleSlice = SimpleSlice[int64]
type UintSimpleSlice = SimpleSlice[uint]
type Uint8SimpleSlice = SimpleSlice[uint8]
type Uint16SimpleSlice = SimpleSlice[uint16]
type Uint32SimpleSlice = SimpleSlice[uint32]
type Uint64SimpleSlice = SimpleSlice[uint64]
type Float32SimpleSlice = SimpleSlice[float32]
type Float64SimpleSlice = SimpleSlice[float64]
type ByteSimpleSlice = SimpleSlice[byte]
type BoolSimpleSlice = SimpleSlice[bool]
type AnySimpleSlice = SimpleSlice[any]

// =============================================================================
// LinkedList[T any] — type aliases for common primitives
// =============================================================================

type StringLinkedList = LinkedList[string]
type IntLinkedList = LinkedList[int]
type Int8LinkedList = LinkedList[int8]
type Int16LinkedList = LinkedList[int16]
type Int32LinkedList = LinkedList[int32]
type Int64LinkedList = LinkedList[int64]
type UintLinkedList = LinkedList[uint]
type Uint8LinkedList = LinkedList[uint8]
type Uint16LinkedList = LinkedList[uint16]
type Uint32LinkedList = LinkedList[uint32]
type Uint64LinkedList = LinkedList[uint64]
type Float32LinkedList = LinkedList[float32]
type Float64LinkedList = LinkedList[float64]
type ByteLinkedList = LinkedList[byte]
type BoolLinkedList = LinkedList[bool]
type AnyLinkedList = LinkedList[any]
