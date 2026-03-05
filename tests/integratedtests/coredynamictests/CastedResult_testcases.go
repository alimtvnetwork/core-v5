package coredynamictests

import (
	"errors"
	"reflect"

	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

type castedResultTestCase struct {
	Case coretestcases.CaseV1
	CR   *coredynamic.CastedResult
}

// ==========================================
// IsInvalid
// ==========================================

var castedResultIsInvalidTestCases = []castedResultTestCase{
	{
		Case: coretestcases.CaseV1{
			Title:         "IsInvalid true on nil receiver",
			ExpectedInput: "true",
		},
		CR: nil,
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "IsInvalid false when IsValid=true",
			ExpectedInput: "false",
		},
		CR: &coredynamic.CastedResult{IsValid: true},
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "IsInvalid true when IsValid=false",
			ExpectedInput: "true",
		},
		CR: &coredynamic.CastedResult{IsValid: false},
	},
}

// ==========================================
// IsNotNull
// ==========================================

var castedResultIsNotNullTestCases = []castedResultTestCase{
	{
		Case: coretestcases.CaseV1{
			Title:         "IsNotNull false on nil receiver",
			ExpectedInput: "false",
		},
		CR: nil,
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "IsNotNull true when IsNull=false",
			ExpectedInput: "true",
		},
		CR: &coredynamic.CastedResult{IsNull: false},
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "IsNotNull false when IsNull=true",
			ExpectedInput: "false",
		},
		CR: &coredynamic.CastedResult{IsNull: true},
	},
}

// ==========================================
// IsNotPointer
// ==========================================

var castedResultIsNotPointerTestCases = []castedResultTestCase{
	{
		Case: coretestcases.CaseV1{
			Title:         "IsNotPointer false on nil receiver",
			ExpectedInput: "false",
		},
		CR: nil,
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "IsNotPointer true when IsPointer=false",
			ExpectedInput: "true",
		},
		CR: &coredynamic.CastedResult{IsPointer: false},
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "IsNotPointer false when IsPointer=true",
			ExpectedInput: "false",
		},
		CR: &coredynamic.CastedResult{IsPointer: true},
	},
}

// ==========================================
// IsNotMatchingAcceptedType
// ==========================================

var castedResultIsNotMatchingAcceptedTypeTestCases = []castedResultTestCase{
	{
		Case: coretestcases.CaseV1{
			Title:         "IsNotMatchingAcceptedType false on nil receiver",
			ExpectedInput: "false",
		},
		CR: nil,
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "IsNotMatchingAcceptedType true when not matching",
			ExpectedInput: "true",
		},
		CR: &coredynamic.CastedResult{IsMatchingAcceptedType: false},
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "IsNotMatchingAcceptedType false when matching",
			ExpectedInput: "false",
		},
		CR: &coredynamic.CastedResult{IsMatchingAcceptedType: true},
	},
}

// ==========================================
// IsSourceKind
// ==========================================

type castedResultIsSourceKindTestCase struct {
	Case      coretestcases.CaseV1
	CR        *coredynamic.CastedResult
	CheckKind reflect.Kind
}

var castedResultIsSourceKindTestCases = []castedResultIsSourceKindTestCase{
	{
		Case: coretestcases.CaseV1{
			Title:         "IsSourceKind false on nil receiver",
			ExpectedInput: "false",
		},
		CR:        nil,
		CheckKind: reflect.String,
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "IsSourceKind true on kind match",
			ExpectedInput: "true",
		},
		CR:        &coredynamic.CastedResult{SourceKind: reflect.Int},
		CheckKind: reflect.Int,
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "IsSourceKind false on mismatch",
			ExpectedInput: "false",
		},
		CR:        &coredynamic.CastedResult{SourceKind: reflect.Int},
		CheckKind: reflect.String,
	},
}

// ==========================================
// HasError
// ==========================================

var castedResultHasErrorTestCases = []castedResultTestCase{
	{
		Case: coretestcases.CaseV1{
			Title:         "HasError false on nil receiver",
			ExpectedInput: "false",
		},
		CR: nil,
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "HasError true when error present",
			ExpectedInput: "true",
		},
		CR: &coredynamic.CastedResult{Error: errors.New("fail")},
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "HasError false when no error",
			ExpectedInput: "false",
		},
		CR: &coredynamic.CastedResult{},
	},
}

// ==========================================
// HasAnyIssues
// ==========================================

var castedResultHasAnyIssuesTestCases = []castedResultTestCase{
	{
		Case: coretestcases.CaseV1{
			Title:         "HasAnyIssues true on nil receiver",
			ExpectedInput: "true",
		},
		CR: nil,
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "HasAnyIssues true when invalid",
			ExpectedInput: "true",
		},
		CR: &coredynamic.CastedResult{IsValid: false, IsMatchingAcceptedType: true},
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "HasAnyIssues true when null",
			ExpectedInput: "true",
		},
		CR: &coredynamic.CastedResult{IsValid: true, IsNull: true, IsMatchingAcceptedType: true},
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "HasAnyIssues true when type not matching",
			ExpectedInput: "true",
		},
		CR: &coredynamic.CastedResult{IsValid: true, IsNull: false, IsMatchingAcceptedType: false},
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "HasAnyIssues false when all good",
			ExpectedInput: "false",
		},
		CR: &coredynamic.CastedResult{
			IsValid:                true,
			IsNull:                 false,
			IsMatchingAcceptedType: true,
		},
	},
}

// ==========================================
// SourceReflectType
// ==========================================

var castedResultSourceReflectTypeTestCases = []castedResultTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "Stores SourceReflectType correctly",
			ExpectedInput: args.Two[string, string]{
				First:  "string", // typeName
				Second: "true",   // isStringKind
			},
		},
		CR: &coredynamic.CastedResult{
			SourceReflectType: reflect.TypeOf(""),
			SourceKind:        reflect.String,
		},
	},
}

// ==========================================
// Casted
// ==========================================

var castedResultCastedValueTestCases = []castedResultTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "Casted stores value and HasAnyIssues false",
			ExpectedInput: args.Two[string, string]{
				First:  "42",    // castedValue
				Second: "false", // hasAnyIssues
			},
		},
		CR: &coredynamic.CastedResult{
			Casted:                 42,
			IsValid:                true,
			IsMatchingAcceptedType: true,
		},
	},
}

// ==========================================
// IsSourcePointer
// ==========================================

var castedResultIsSourcePointerTestCases = []castedResultTestCase{
	{
		Case: coretestcases.CaseV1{
			Title:         "IsSourcePointer field works",
			ExpectedInput: "true",
		},
		CR: &coredynamic.CastedResult{IsSourcePointer: true},
	},
}
