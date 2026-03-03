package coredynamictests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================
// Constructors
//
// Each case tests a different factory function for creating Dynamic values.
// The "constructor" key holds a string tag used by the test runner's switch.
// Constructor function references are validated at init via Dynamic_method_refs.go.
// ==========================================

var dynamicConstructorTestCases = []coretestcases.CaseV1{
	{
		Title: "NewDynamicValid creates valid Dynamic",
		ArrangeInput: args.Map{
			"constructorRef": refNewDynamicValid, // build error if renamed
			"inputData":      "hello",            // string data to wrap
		},
		// [0] IsValid=true, [1] Value()="hello"
		ExpectedInput: []string{"true", "hello"},
	},
	{
		Title: "NewDynamic with isValid=false creates invalid Dynamic",
		ArrangeInput: args.Map{
			"constructorRef": refNewDynamic, // build error if renamed
			"isValid":        false,         // explicitly invalid
		},
		// [0] IsValid=false, [1] IsInvalid=true
		ExpectedInput: []string{"false", "true"},
	},
	{
		Title: "InvalidDynamic creates invalid nil Dynamic",
		ArrangeInput: args.Map{
			"constructorRef": refInvalidDynamic, // build error if renamed
		},
		// [0] IsValid=false, [1] IsNull=true
		ExpectedInput: []string{"false", "true"},
	},
	{
		Title: "InvalidDynamicPtr creates invalid nil Dynamic pointer",
		ArrangeInput: args.Map{
			"constructorRef": refInvalidDynamicPtr, // build error if renamed
		},
		// [0] ptr!=nil=true, [1] IsValid=false, [2] IsNull=true
		ExpectedInput: []string{"true", "false", "true"},
	},
	{
		Title: "NewDynamicPtr creates pointer Dynamic",
		ArrangeInput: args.Map{
			"constructorRef": refNewDynamicPtr, // build error if renamed
			"inputData":      42,               // integer data to wrap
			"isValid":        true,
		},
		// [0] ptr!=nil=true, [1] IsValid=true, [2] Value()=42
		ExpectedInput: []string{"true", "true", "42"},
	},
}

// ==========================================
// Clone
//
// Tests Clone(), ClonePtr(), and NonPtr() methods.
// Method name strings are kept minimal; constructorRef ensures
// the actual Dynamic methods are compile-time validated.
// ==========================================

var dynamicCloneTestCases = []coretestcases.CaseV1{
	{
		Title: "Clone creates independent copy",
		ArrangeInput: args.Map{
			"scenario":  "clone",  // which clone method to test
			"inputData": "data",   // string data to wrap
		},
		// [0] cloned.Value()="data", [1] cloned.IsValid()=true
		ExpectedInput: []string{"data", "true"},
	},
	{
		Title: "ClonePtr creates independent pointer copy",
		ArrangeInput: args.Map{
			"scenario":  "clonePtr",
			"inputData": "data",
		},
		// [0] ptr!=nil=true, [1] cloned.Value()="data"
		ExpectedInput: []string{"true", "data"},
	},
	{
		Title: "ClonePtr returns nil on nil receiver",
		ArrangeInput: args.Map{
			"scenario":    "clonePtr",
			"nilReceiver": true, // test nil receiver safety
		},
		// [0] result==nil → true
		ExpectedInput: []string{"true"},
	},
	{
		Title: "NonPtr returns value copy",
		ArrangeInput: args.Map{
			"scenario":  "nonPtr",
			"inputData": "x",
		},
		// [0] nonPtr.Value()="x"
		ExpectedInput: []string{"x"},
	},
}

// ==========================================
// Type Checks — Bool method refs
//
// Each case stores a DynamicBoolMethodRef in "checkRef".
// This provides compile-time safety: renaming the method
// on *Dynamic causes a build error at the ref declaration.
//
// "inputData" holds the ACTUAL Go value to wrap in Dynamic.
// No more magic strings like "slice_int" — the real data is here.
// ==========================================

var dynamicTypeCheckTestCases = []coretestcases.CaseV1{
	// --- DataValue (special: tests Data() == Value()) ---
	{
		Title: "Data and Value return same inner data",
		ArrangeInput: args.Map{
			"scenario":  "dataValue", // special multi-method check
			"inputData": 99,          // integer wrapped in Dynamic
		},
		// [0] Data()="99", [1] Data()==Value() → true
		ExpectedInput: []string{"99", "true"},
	},

	// --- IsNull ---
	{
		Title: "IsNull true for nil data",
		ArrangeInput: args.Map{
			"checkRef":  refIsNull, // build error if IsNull renamed
			"inputData": nil,       // nil data → IsNull should be true
			"isValid":   true,      // valid Dynamic but nil data
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsNull false for non-nil data",
		ArrangeInput: args.Map{
			"checkRef":  refIsNull,
			"inputData": "x", // non-nil string
		},
		ExpectedInput: []string{"false"},
	},

	// --- String (non-empty check) ---
	{
		Title: "String returns non-empty for valid",
		ArrangeInput: args.Map{
			"scenario":  "stringNonEmpty", // special: checks len > 0
			"inputData": "hello",
		},
		// [0] String()!="" → true
		ExpectedInput: []string{"true"},
	},

	// --- IsStringType ---
	{
		Title: "IsStringType true for string",
		ArrangeInput: args.Map{
			"checkRef":  refIsStringType,
			"inputData": "text", // string type → true
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsStringType false for int",
		ArrangeInput: args.Map{
			"checkRef":  refIsStringType,
			"inputData": 42, // int type → false
		},
		ExpectedInput: []string{"false"},
	},

	// --- IsNumber ---
	{
		Title: "IsNumber true for int",
		ArrangeInput: args.Map{
			"checkRef":  refIsNumber,
			"inputData": 42, // int is a number
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsNumber false for string",
		ArrangeInput: args.Map{
			"checkRef":  refIsNumber,
			"inputData": "x", // string is not a number
		},
		ExpectedInput: []string{"false"},
	},

	// --- IsPrimitive ---
	{
		Title: "IsPrimitive true for int",
		ArrangeInput: args.Map{
			"checkRef":  refIsPrimitive,
			"inputData": 10, // int is primitive
		},
		ExpectedInput: []string{"true"},
	},

	// --- IsFunc ---
	{
		Title: "IsFunc true for function",
		ArrangeInput: args.Map{
			"checkRef":  refIsFunc,
			"inputData": func() {}, // anonymous function
		},
		ExpectedInput: []string{"true"},
	},

	// --- IsSliceOrArray ---
	{
		Title: "IsSliceOrArray true for slice",
		ArrangeInput: args.Map{
			"checkRef":  refIsSliceOrArray,
			"inputData": []int{1, 2, 3}, // actual int slice
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsSliceOrArray false for string",
		ArrangeInput: args.Map{
			"checkRef":  refIsSliceOrArray,
			"inputData": "x", // string is not a slice
		},
		ExpectedInput: []string{"false"},
	},

	// --- IsSliceOrArrayOrMap ---
	{
		Title: "IsSliceOrArrayOrMap true for map",
		ArrangeInput: args.Map{
			"checkRef":  refIsSliceOrArrayOrMap,
			"inputData": map[string]int{"a": 1}, // actual map
		},
		ExpectedInput: []string{"true"},
	},

	// --- IsMap ---
	{
		Title: "IsMap true for map",
		ArrangeInput: args.Map{
			"checkRef":  refIsMap,
			"inputData": map[string]int{"x": 1}, // actual map
		},
		ExpectedInput: []string{"true"},
	},

	// --- IsPointer ---
	{
		Title: "IsPointer true for pointer data",
		ArrangeInput: args.Map{
			"checkRef":  refIsPointer,
			"scenario":  "pointer", // special: wraps &val
		},
		ExpectedInput: []string{"true"},
	},

	// --- IsValueType ---
	{
		Title: "IsValueType true for non-pointer",
		ArrangeInput: args.Map{
			"checkRef":  refIsValueType,
			"inputData": 42, // plain int, not a pointer
		},
		ExpectedInput: []string{"true"},
	},
}

// ==========================================
// IsStruct
// ==========================================

var dynamicIsStructTestCases = []coretestcases.CaseV1{
	{
		Title: "IsStruct true for struct",
		ArrangeInput: args.Map{
			"checkRef":  refIsStruct,
			"scenario":  "struct", // special: creates a struct value
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsStruct false for int",
		ArrangeInput: args.Map{
			"checkRef":  refIsStruct,
			"inputData": 5, // int is not a struct
		},
		ExpectedInput: []string{"false"},
	},
}

// ==========================================
// Length
//
// Tests Dynamic.Length() for slices, nil, and maps.
// "inputData" holds the actual Go value; no magic strings.
// ==========================================

var dynamicLengthTestCases = []coretestcases.CaseV1{
	{
		Title: "Length returns slice length",
		ArrangeInput: args.Map{
			"inputData": []int{1, 2, 3}, // 3-element slice
		},
		// Length() = 3
		ExpectedInput: []string{"3"},
	},
	{
		Title: "Length returns 0 for nil data",
		ArrangeInput: args.Map{
			"inputData": nil,    // nil data
			"isValid":   false,  // invalid Dynamic
		},
		ExpectedInput: []string{"0"},
	},
	{
		Title: "Length returns map length",
		ArrangeInput: args.Map{
			"inputData": map[string]int{"a": 1, "b": 2}, // 2-entry map
		},
		ExpectedInput: []string{"2"},
	},
}

// ==========================================
// Value Extraction
//
// Tests ValueInt, ValueBool, ValueString, etc.
// "inputData" is the actual Go value passed to NewDynamicValid.
// Expected lines document the return value.
// ==========================================

var dynamicValueIntTestCases = []coretestcases.CaseV1{
	{
		Title: "ValueInt returns int value",
		ArrangeInput: args.Map{
			"inputData": 42, // int → ValueInt returns 42
		},
		ExpectedInput: []string{"42"},
	},
	{
		Title: "ValueInt returns -1 for non-int",
		ArrangeInput: args.Map{
			"inputData": "not-int", // string → ValueInt returns -1 (default)
		},
		ExpectedInput: []string{"-1"},
	},
}

var dynamicValueBoolTestCases = []coretestcases.CaseV1{
	{
		Title: "ValueBool returns true",
		ArrangeInput: args.Map{
			"inputData": true, // bool true
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "ValueBool returns false for non-bool",
		ArrangeInput: args.Map{
			"inputData": "x", // string → ValueBool returns false
		},
		ExpectedInput: []string{"false"},
	},
}

var dynamicValueStringTestCases = []coretestcases.CaseV1{
	{
		Title: "ValueString returns string directly",
		ArrangeInput: args.Map{
			"inputData": "hello", // string → returned as-is
		},
		// [0] ValueString() = "hello"
		ExpectedInput: []string{"hello"},
	},
	{
		Title: "ValueString formats non-string as non-empty",
		ArrangeInput: args.Map{
			"inputData": 42, // int → formatted to non-empty string
		},
		// [0] ValueString()!="" → true
		ExpectedInput: []string{"true"},
	},
	{
		Title: "ValueString returns empty for nil data",
		ArrangeInput: args.Map{
			"inputData": nil, // nil → empty string
			"isValid":   true,
		},
		// [0] ValueString()=="" → true
		ExpectedInput: []string{"true"},
	},
}

var dynamicValueStringsTestCases = []coretestcases.CaseV1{
	{
		Title: "ValueStrings returns []string",
		ArrangeInput: args.Map{
			"inputData": []string{"a", "b"}, // actual []string
		},
		// Returns the slice elements directly
		ExpectedInput: []string{"a", "b"},
	},
	{
		Title: "ValueStrings returns nil for non-[]string",
		ArrangeInput: args.Map{
			"inputData": 42, // int → ValueStrings returns nil
		},
		// [0] ValueStrings()==nil → true
		ExpectedInput: []string{"true"},
	},
}

var dynamicValueUIntTestCases = []coretestcases.CaseV1{
	{
		Title: "ValueUInt returns uint value",
		ArrangeInput: args.Map{
			"inputData": uint(10), // typed uint
		},
		ExpectedInput: []string{"10"},
	},
}

var dynamicValueInt64TestCases = []coretestcases.CaseV1{
	{
		Title: "ValueInt64 returns int64 value",
		ArrangeInput: args.Map{
			"inputData": int64(999), // typed int64
		},
		ExpectedInput: []string{"999"},
	},
}

var dynamicBytesTestCases = []coretestcases.CaseV1{
	{
		Title: "Bytes returns []byte",
		ArrangeInput: args.Map{
			"inputData": []byte("raw"), // actual byte slice
		},
		// [0] ok=true, [1] string(bytes)="raw"
		ExpectedInput: []string{"true", "raw"},
	},
	{
		Title: "Bytes returns false for non-bytes",
		ArrangeInput: args.Map{
			"inputData": "str", // plain string, not []byte
		},
		// [0] ok=false
		ExpectedInput: []string{"false"},
	},
	{
		Title: "Bytes returns nil,false on nil receiver",
		ArrangeInput: args.Map{
			"nilReceiver": true, // test nil *Dynamic safety
		},
		// [0] bytes==nil → true, [1] ok=false
		ExpectedInput: []string{"true", "false"},
	},
}

var dynamicIntDefaultTestCases = []coretestcases.CaseV1{
	{
		Title: "IntDefault parses int value",
		ArrangeInput: args.Map{
			"inputData":    42, // int data
			"defaultValue": 0,  // fallback (not used here)
		},
		// [0] isSuccess=true, [1] val=42
		ExpectedInput: []string{"true", "42"},
	},
	{
		Title: "IntDefault returns default on nil data",
		ArrangeInput: args.Map{
			"inputData":    nil, // nil data
			"isValid":      true,
			"defaultValue": 99, // fallback value
		},
		// [0] isSuccess=false, [1] val=99 (the default)
		ExpectedInput: []string{"false", "99"},
	},
}

var dynamicValueNullErrTestCases = []coretestcases.CaseV1{
	{
		Title: "ValueNullErr returns error on nil receiver",
		ArrangeInput: args.Map{
			"nilReceiver": true, // nil *Dynamic receiver
		},
		// [0] err!=nil → true
		ExpectedInput: []string{"true"},
	},
	{
		Title: "ValueNullErr returns error on null data",
		ArrangeInput: args.Map{
			"inputData": nil, // nil inner data
			"isValid":   true,
		},
		// [0] err!=nil → true
		ExpectedInput: []string{"true"},
	},
	{
		Title: "ValueNullErr returns nil for valid data",
		ArrangeInput: args.Map{
			"inputData": "ok", // non-nil data
		},
		// [0] err!=nil → false (no error)
		ExpectedInput: []string{"false"},
	},
}

// ==========================================
// Reflect
//
// Tests ReflectKind, IsReflectKind, ReflectTypeName,
// ReflectType, IsReflectTypeOf.
// ==========================================

var dynamicReflectTestCases = []coretestcases.CaseV1{
	{
		Title: "ReflectKind returns String for string",
		ArrangeInput: args.Map{
			"scenario":  "reflectKind",
			"inputData": "text", // string → Kind=String
		},
		ExpectedInput: []string{"string"},
	},
	{
		Title: "ReflectKind returns Int for int",
		ArrangeInput: args.Map{
			"scenario":  "reflectKind",
			"inputData": 42, // int → Kind=Int
		},
		ExpectedInput: []string{"int"},
	},
	{
		Title: "IsReflectKind matches correctly",
		ArrangeInput: args.Map{
			"scenario":  "isReflectKindMatch",
			"inputData": "x", // string → check against reflect.String
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsReflectKind returns false on mismatch",
		ArrangeInput: args.Map{
			"scenario":  "isReflectKindMismatch",
			"inputData": "x", // string → check against reflect.Int → false
		},
		ExpectedInput: []string{"false"},
	},
	{
		Title: "ReflectTypeName returns non-empty",
		ArrangeInput: args.Map{
			"scenario":    "reflectTypeName",
			"checkRef":    refReflectTypeName, // build error if renamed
			"inputData":   "text",
		},
		// [0] ReflectTypeName()!="" → true
		ExpectedInput: []string{"true"},
	},
	{
		Title: "ReflectType returns correct type",
		ArrangeInput: args.Map{
			"scenario":  "reflectType",
			"inputData": 42, // check ReflectType() == reflect.TypeOf(42)
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsReflectTypeOf matches type",
		ArrangeInput: args.Map{
			"scenario":  "isReflectTypeOf",
			"inputData": "hello", // check against reflect.TypeOf("")
		},
		ExpectedInput: []string{"true"},
	},
}

// ==========================================
// Loop
// ==========================================

var dynamicLoopTestCases = []coretestcases.CaseV1{
	{
		Title: "Loop iterates slice items",
		ArrangeInput: args.Map{
			"scenario": "iterate",
		},
		// [0] isCalled=true, [1..3] collected items
		ExpectedInput: []string{"true", "a", "b", "c"},
	},
	{
		Title: "Loop returns false for invalid",
		ArrangeInput: args.Map{
			"scenario": "invalid",
		},
		// [0] isCalled=false
		ExpectedInput: []string{"false"},
	},
	{
		Title: "Loop respects break",
		ArrangeInput: args.Map{
			"scenario": "break",
		},
		// [0] iterations before break = 2
		ExpectedInput: []string{"2"},
	},
}

// ==========================================
// ItemAccess
// ==========================================

var dynamicItemAccessTestCases = []coretestcases.CaseV1{
	{
		Title: "ItemUsingIndex returns correct element",
		ArrangeInput: args.Map{
			"scenario": "itemUsingIndex",
		},
		// [0] index 0 = "a", [1] index 1 = "b"
		ExpectedInput: []string{"a", "b"},
	},
	{
		Title: "ItemUsingKey returns map value",
		ArrangeInput: args.Map{
			"scenario": "itemUsingKey",
		},
		// [0] map["k"] = 42
		ExpectedInput: []string{"42"},
	},
}

// ==========================================
// IsStructStringNullOrEmpty
// ==========================================

var dynamicStructStringNullOrEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "IsStructStringNullOrEmpty true on nil data",
		ArrangeInput: args.Map{
			"checkRef":  refIsStructStringNullOrEmpty,
			"inputData": nil,  // nil data → true
			"isValid":   true,
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsStructStringNullOrEmpty false for non-empty",
		ArrangeInput: args.Map{
			"checkRef":  refIsStructStringNullOrEmpty,
			"inputData": "text", // non-empty → false
		},
		ExpectedInput: []string{"false"},
	},
}
