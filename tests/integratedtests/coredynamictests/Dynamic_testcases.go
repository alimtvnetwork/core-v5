package coredynamictests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================
// Constructors
// ==========================================

var dynamicConstructorTestCases = []coretestcases.CaseV1{
	{
		Title: "NewDynamicValid creates valid Dynamic",
		ArrangeInput: args.Map{
			"constructor": "NewDynamicValid",
			"value":       "hello",
		},
		ExpectedInput: []string{"true", "hello"},
	},
	{
		Title: "NewDynamic with isValid=false creates invalid Dynamic",
		ArrangeInput: args.Map{
			"constructor": "NewDynamic",
			"isValid":     false,
		},
		ExpectedInput: []string{"false", "true"},
	},
	{
		Title: "InvalidDynamic creates invalid nil Dynamic",
		ArrangeInput: args.Map{
			"constructor": "InvalidDynamic",
		},
		ExpectedInput: []string{"false", "true"},
	},
	{
		Title: "InvalidDynamicPtr creates invalid nil Dynamic pointer",
		ArrangeInput: args.Map{
			"constructor": "InvalidDynamicPtr",
		},
		ExpectedInput: []string{"true", "false", "true"},
	},
	{
		Title: "NewDynamicPtr creates pointer Dynamic",
		ArrangeInput: args.Map{
			"constructor": "NewDynamicPtr",
			"value":       42,
			"isValid":     true,
		},
		ExpectedInput: []string{"true", "true", "42"},
	},
}

// ==========================================
// Clone
// ==========================================

var dynamicCloneTestCases = []coretestcases.CaseV1{
	{
		Title: "Clone creates independent copy",
		ArrangeInput: args.Map{
			"method": "Clone",
			"value":  "data",
		},
		ExpectedInput: []string{"data", "true"},
	},
	{
		Title: "ClonePtr creates independent pointer copy",
		ArrangeInput: args.Map{
			"method": "ClonePtr",
			"value":  "data",
		},
		ExpectedInput: []string{"true", "data"},
	},
	{
		Title: "ClonePtr returns nil on nil receiver",
		ArrangeInput: args.Map{
			"method":   "ClonePtr",
			"receiver": "nil",
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "NonPtr returns value copy",
		ArrangeInput: args.Map{
			"method": "NonPtr",
			"value":  "x",
		},
		ExpectedInput: []string{"x"},
	},
}

// ==========================================
// Type Checks
// ==========================================

var dynamicTypeCheckTestCases = []coretestcases.CaseV1{
	{
		Title: "Data and Value return same inner data",
		ArrangeInput: args.Map{
			"check": "DataValue",
			"value": 99,
		},
		ExpectedInput: []string{"99", "true"},
	},
	{
		Title: "IsNull true for nil data",
		ArrangeInput: args.Map{
			"check":   "IsNull",
			"isValid": true,
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsNull false for non-nil data",
		ArrangeInput: args.Map{
			"check": "IsNull",
			"value": "x",
		},
		ExpectedInput: []string{"false"},
	},
	{
		Title: "String returns non-empty for valid",
		ArrangeInput: args.Map{
			"check": "String",
			"value": "hello",
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsStringType true for string",
		ArrangeInput: args.Map{
			"check": "IsStringType",
			"value": "text",
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsStringType false for int",
		ArrangeInput: args.Map{
			"check": "IsStringType",
			"value": 42,
		},
		ExpectedInput: []string{"false"},
	},
	{
		Title: "IsNumber true for int",
		ArrangeInput: args.Map{
			"check": "IsNumber",
			"value": 42,
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsNumber false for string",
		ArrangeInput: args.Map{
			"check": "IsNumber",
			"value": "x",
		},
		ExpectedInput: []string{"false"},
	},
	{
		Title: "IsPrimitive true for int",
		ArrangeInput: args.Map{
			"check": "IsPrimitive",
			"value": 10,
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsFunc true for function",
		ArrangeInput: args.Map{
			"check": "IsFunc",
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsSliceOrArray true for slice",
		ArrangeInput: args.Map{
			"check": "IsSliceOrArray",
			"value": "slice_int",
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsSliceOrArray false for string",
		ArrangeInput: args.Map{
			"check": "IsSliceOrArray",
			"value": "x",
		},
		ExpectedInput: []string{"false"},
	},
	{
		Title: "IsSliceOrArrayOrMap true for map",
		ArrangeInput: args.Map{
			"check": "IsSliceOrArrayOrMap",
			"value": "map_string_int",
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsMap true for map",
		ArrangeInput: args.Map{
			"check": "IsMap",
			"value": "map_string_int",
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsPointer true for pointer data",
		ArrangeInput: args.Map{
			"check": "IsPointer",
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsValueType true for non-pointer",
		ArrangeInput: args.Map{
			"check": "IsValueType",
			"value": 42,
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
			"value": "struct",
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsStruct false for int",
		ArrangeInput: args.Map{
			"value": 5,
		},
		ExpectedInput: []string{"false"},
	},
}

// ==========================================
// Length
// ==========================================

var dynamicLengthTestCases = []coretestcases.CaseV1{
	{
		Title: "Length returns slice length",
		ArrangeInput: args.Map{
			"value": "slice_int",
		},
		ExpectedInput: []string{"3"},
	},
	{
		Title: "Length returns 0 for nil data",
		ArrangeInput: args.Map{
			"value": "nil",
		},
		ExpectedInput: []string{"0"},
	},
	{
		Title: "Length returns map length",
		ArrangeInput: args.Map{
			"value": "map_2",
		},
		ExpectedInput: []string{"2"},
	},
}

// ==========================================
// Value Extraction
// ==========================================

var dynamicValueIntTestCases = []coretestcases.CaseV1{
	{
		Title: "ValueInt returns int value",
		ArrangeInput: args.Map{
			"value": 42,
		},
		ExpectedInput: []string{"42"},
	},
	{
		Title: "ValueInt returns -1 for non-int",
		ArrangeInput: args.Map{
			"value": "not-int",
		},
		ExpectedInput: []string{"-1"},
	},
}

var dynamicValueBoolTestCases = []coretestcases.CaseV1{
	{
		Title: "ValueBool returns true",
		ArrangeInput: args.Map{
			"value": true,
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "ValueBool returns false for non-bool",
		ArrangeInput: args.Map{
			"value": "x",
		},
		ExpectedInput: []string{"false"},
	},
}

var dynamicValueStringTestCases = []coretestcases.CaseV1{
	{
		Title: "ValueString returns string directly",
		ArrangeInput: args.Map{
			"value": "hello",
		},
		ExpectedInput: []string{"hello"},
	},
	{
		Title: "ValueString formats non-string as non-empty",
		ArrangeInput: args.Map{
			"value": 42,
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "ValueString returns empty for nil data",
		ArrangeInput: args.Map{
			"value": "nil",
		},
		ExpectedInput: []string{"true"},
	},
}

var dynamicValueStringsTestCases = []coretestcases.CaseV1{
	{
		Title: "ValueStrings returns []string",
		ArrangeInput: args.Map{
			"value": "strings_ab",
		},
		ExpectedInput: []string{"a", "b"},
	},
	{
		Title: "ValueStrings returns nil for non-[]string",
		ArrangeInput: args.Map{
			"value": 42,
		},
		ExpectedInput: []string{"true"},
	},
}

var dynamicValueUIntTestCases = []coretestcases.CaseV1{
	{
		Title: "ValueUInt returns uint value",
		ArrangeInput: args.Map{
			"value": uint(10),
		},
		ExpectedInput: []string{"10"},
	},
}

var dynamicValueInt64TestCases = []coretestcases.CaseV1{
	{
		Title: "ValueInt64 returns int64 value",
		ArrangeInput: args.Map{
			"value": int64(999),
		},
		ExpectedInput: []string{"999"},
	},
}

var dynamicBytesTestCases = []coretestcases.CaseV1{
	{
		Title: "Bytes returns []byte",
		ArrangeInput: args.Map{
			"value": "bytes_raw",
		},
		ExpectedInput: []string{"true", "raw"},
	},
	{
		Title: "Bytes returns false for non-bytes",
		ArrangeInput: args.Map{
			"value": "str",
		},
		ExpectedInput: []string{"false"},
	},
	{
		Title: "Bytes returns nil,false on nil receiver",
		ArrangeInput: args.Map{
			"receiver": "nil",
		},
		ExpectedInput: []string{"true", "false"},
	},
}

var dynamicIntDefaultTestCases = []coretestcases.CaseV1{
	{
		Title: "IntDefault parses int value",
		ArrangeInput: args.Map{
			"value":        42,
			"defaultValue": 0,
		},
		ExpectedInput: []string{"true", "42"},
	},
	{
		Title: "IntDefault returns default on nil data",
		ArrangeInput: args.Map{
			"value":        "nil",
			"defaultValue": 99,
		},
		ExpectedInput: []string{"false", "99"},
	},
}

var dynamicValueNullErrTestCases = []coretestcases.CaseV1{
	{
		Title: "ValueNullErr returns error on nil receiver",
		ArrangeInput: args.Map{
			"receiver": "nil",
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "ValueNullErr returns error on null data",
		ArrangeInput: args.Map{
			"value": "nil",
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "ValueNullErr returns nil for valid data",
		ArrangeInput: args.Map{
			"value": "ok",
		},
		ExpectedInput: []string{"false"},
	},
}

// ==========================================
// Reflect
// ==========================================

var dynamicReflectTestCases = []coretestcases.CaseV1{
	{
		Title: "ReflectKind returns String for string",
		ArrangeInput: args.Map{
			"check": "ReflectKind",
			"value": "text",
		},
		ExpectedInput: []string{"string"},
	},
	{
		Title: "ReflectKind returns Int for int",
		ArrangeInput: args.Map{
			"check": "ReflectKind",
			"value": 42,
		},
		ExpectedInput: []string{"int"},
	},
	{
		Title: "IsReflectKind matches correctly",
		ArrangeInput: args.Map{
			"check": "IsReflectKind",
			"value": "x",
			"match": true,
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsReflectKind returns false on mismatch",
		ArrangeInput: args.Map{
			"check": "IsReflectKind",
			"value": "x",
			"match": false,
		},
		ExpectedInput: []string{"false"},
	},
	{
		Title: "ReflectTypeName returns non-empty",
		ArrangeInput: args.Map{
			"check": "ReflectTypeName",
			"value": "text",
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "ReflectType returns correct type",
		ArrangeInput: args.Map{
			"check": "ReflectType",
			"value": 42,
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsReflectTypeOf matches type",
		ArrangeInput: args.Map{
			"check": "IsReflectTypeOf",
			"value": "hello",
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
		ExpectedInput: []string{"true", "a", "b", "c"},
	},
	{
		Title: "Loop returns false for invalid",
		ArrangeInput: args.Map{
			"scenario": "invalid",
		},
		ExpectedInput: []string{"false"},
	},
	{
		Title: "Loop respects break",
		ArrangeInput: args.Map{
			"scenario": "break",
		},
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
			"method": "ItemUsingIndex",
		},
		ExpectedInput: []string{"a", "b"},
	},
	{
		Title: "ItemUsingKey returns map value",
		ArrangeInput: args.Map{
			"method": "ItemUsingKey",
		},
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
			"value": "nil",
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsStructStringNullOrEmpty false for non-empty",
		ArrangeInput: args.Map{
			"value": "text",
		},
		ExpectedInput: []string{"false"},
	},
}
