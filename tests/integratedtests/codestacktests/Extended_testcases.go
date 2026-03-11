package codestacktests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// =============================================================================
// stacksTo: Reflection-based method discovery test cases
// =============================================================================

var extStacksToMethodDiscoveryTestCases = []coretestcases.CaseV1{
	{
		Title: "stacksTo methods discovered via reflection -- all exported methods found",
		ArrangeInput: args.Map{
			"when": "reflecting on StacksTo private struct var",
		},
		ExpectedInput: args.Map{
			"hasBytesMethod":         true,
			"hasBytesDefaultMethod":  true,
			"hasStringMethod":        true,
			"hasStringUsingFmtMethod": true,
			"hasJsonStringMethod":    true,
			"hasJsonStringDefaultMethod": true,
			"hasStringNoCountMethod": true,
			"hasStringDefaultMethod": true,
		},
	},
}

// =============================================================================
// stacksTo: Bytes
// =============================================================================

var extStacksToBytesTestCases = []coretestcases.CaseV1{
	{
		Title: "stacksTo Bytes returns non-empty -- skip index 0",
		ArrangeInput: args.Map{
			"when":      "calling Bytes via reflection",
			"skipIndex": 0,
		},
		ExpectedInput: args.Map{
			"notEmpty": true,
		},
	},
}

// =============================================================================
// stacksTo: BytesDefault
// =============================================================================

var extStacksToBytesDefaultTestCases = []coretestcases.CaseV1{
	{
		Title: "stacksTo BytesDefault returns non-empty -- default skip",
		ArrangeInput: args.Map{
			"when": "calling BytesDefault via reflection",
		},
		ExpectedInput: args.Map{
			"notEmpty": true,
		},
	},
}

// =============================================================================
// stacksTo: String
// =============================================================================

var extStacksToStringTestCases = []coretestcases.CaseV1{
	{
		Title: "stacksTo String returns non-empty -- skip 0 count 5",
		ArrangeInput: args.Map{
			"when":      "calling String via reflection",
			"skipIndex": 0,
			"count":     5,
		},
		ExpectedInput: args.Map{
			"notEmpty": true,
		},
	},
}

// =============================================================================
// stacksTo: StringUsingFmt
// =============================================================================

var extStacksToStringUsingFmtTestCases = []coretestcases.CaseV1{
	{
		Title: "stacksTo StringUsingFmt returns non-empty -- custom formatter",
		ArrangeInput: args.Map{
			"when": "calling StringUsingFmt via reflection",
		},
		ExpectedInput: args.Map{
			"notEmpty": true,
		},
	},
}

// =============================================================================
// stacksTo: JsonString
// =============================================================================

var extStacksToJsonStringTestCases = []coretestcases.CaseV1{
	{
		Title: "stacksTo JsonString returns non-empty -- skip 0",
		ArrangeInput: args.Map{
			"when":      "calling JsonString via reflection",
			"skipIndex": 0,
		},
		ExpectedInput: args.Map{
			"notEmpty": true,
		},
	},
}

// =============================================================================
// stacksTo: JsonStringDefault
// =============================================================================

var extStacksToJsonStringDefaultTestCases = []coretestcases.CaseV1{
	{
		Title: "stacksTo JsonStringDefault returns non-empty -- default skip",
		ArrangeInput: args.Map{
			"when": "calling JsonStringDefault via reflection",
		},
		ExpectedInput: args.Map{
			"notEmpty": true,
		},
	},
}

// =============================================================================
// stacksTo: StringNoCount
// =============================================================================

var extStacksToStringNoCountTestCases = []coretestcases.CaseV1{
	{
		Title: "stacksTo StringNoCount returns non-empty -- skip 0",
		ArrangeInput: args.Map{
			"when":      "calling StringNoCount via reflection",
			"skipIndex": 0,
		},
		ExpectedInput: args.Map{
			"notEmpty": true,
		},
	},
}

// =============================================================================
// stacksTo: StringDefault
// =============================================================================

var extStacksToStringDefaultTestCases = []coretestcases.CaseV1{
	{
		Title: "stacksTo StringDefault returns non-empty -- default skip",
		ArrangeInput: args.Map{
			"when": "calling StringDefault via reflection",
		},
		ExpectedInput: args.Map{
			"notEmpty": true,
		},
	},
}

// =============================================================================
// currentNameOf
// =============================================================================

var extNameOfAllTestCases = []coretestcases.CaseV1{
	{
		Title: "NameOf All returns parsed names -- full func name 'github.com/pkg.(*Struct).Method'",
		ArrangeInput: args.Map{
			"when":     "given full function name",
			"fullName": "github.com/mypackage.(*MyStruct).DoWork",
		},
		ExpectedInput: args.Map{
			"hasFullMethod":  true,
			"hasPackageName": true,
			"hasMethodName":  true,
		},
	},
	{
		Title: "NameOf All returns empty -- empty string input",
		ArrangeInput: args.Map{
			"when":     "given empty string",
			"fullName": "",
		},
		ExpectedInput: args.Map{
			"hasFullMethod":  false,
			"hasPackageName": false,
			"hasMethodName":  false,
		},
	},
}

var extNameOfMethodTestCases = []coretestcases.CaseV1{
	{
		Title: "NameOf Method returns current method name -- called from test",
		ArrangeInput: args.Map{
			"when": "called from test function",
		},
		ExpectedInput: args.Map{
			"hasMethodName": true,
		},
	},
}

var extNameOfPackageTestCases = []coretestcases.CaseV1{
	{
		Title: "NameOf Package returns current package name -- called from test",
		ArrangeInput: args.Map{
			"when": "called from test function",
		},
		ExpectedInput: args.Map{
			"hasPackageName": true,
		},
	},
}

// =============================================================================
// newTraceCollection
// =============================================================================

var extNewTraceCollectionCapTestCases = []coretestcases.CaseV1{
	{
		Title: "newTraceCollection Cap returns empty collection -- capacity 10",
		ArrangeInput: args.Map{
			"when":     "creating with capacity",
			"capacity": 10,
		},
		ExpectedInput: args.Map{
			"length":  0,
			"isEmpty": true,
		},
	},
}

var extNewTraceCollectionDefaultTestCases = []coretestcases.CaseV1{
	{
		Title: "newTraceCollection Default returns empty collection -- default capacity",
		ArrangeInput: args.Map{
			"when": "creating with default capacity",
		},
		ExpectedInput: args.Map{
			"length":  0,
			"isEmpty": true,
		},
	},
}

var extNewTraceCollectionUsingTestCases = []coretestcases.CaseV1{
	{
		Title: "newTraceCollection Using returns collection with items -- 2 traces no clone",
		ArrangeInput: args.Map{
			"when":    "creating using traces without clone",
			"isClone": false,
		},
		ExpectedInput: args.Map{
			"length": 2,
		},
	},
	{
		Title: "newTraceCollection Using returns cloned items -- 2 traces with clone",
		ArrangeInput: args.Map{
			"when":    "creating using traces with clone",
			"isClone": true,
		},
		ExpectedInput: args.Map{
			"length": 2,
		},
	},
}

var extNewTraceCollectionEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "newTraceCollection Empty returns empty collection -- empty",
		ArrangeInput: args.Map{
			"when": "creating empty",
		},
		ExpectedInput: args.Map{
			"length": 0,
		},
	},
}

// =============================================================================
// newStacksCreator
// =============================================================================

var extNewStacksCreatorDefaultTestCases = []coretestcases.CaseV1{
	{
		Title: "newStacksCreator Default returns collection -- skip 0 count 3",
		ArrangeInput: args.Map{
			"when":      "creating with default settings",
			"skipIndex": 0,
			"count":     3,
		},
		ExpectedInput: args.Map{
			"hasItems": true,
		},
	},
}

var extNewStacksCreatorSkipOneTestCases = []coretestcases.CaseV1{
	{
		Title: "newStacksCreator SkipOne returns collection -- skip 1",
		ArrangeInput: args.Map{
			"when": "creating with skip one",
		},
		ExpectedInput: args.Map{
			"hasItems": true,
		},
	},
}

var extNewStacksCreatorSkipNoneTestCases = []coretestcases.CaseV1{
	{
		Title: "newStacksCreator SkipNone returns collection -- skip none",
		ArrangeInput: args.Map{
			"when": "creating with skip none",
		},
		ExpectedInput: args.Map{
			"hasItems": true,
		},
	},
}

// =============================================================================
// dirGetter / fileGetter
// =============================================================================

var extDirGetterCurDirTestCases = []coretestcases.CaseV1{
	{
		Title: "Dir CurDir returns non-empty -- current directory",
		ArrangeInput: args.Map{
			"when": "calling CurDir",
		},
		ExpectedInput: args.Map{
			"notEmpty": true,
		},
	},
}

var extDirGetterRepoDirTestCases = []coretestcases.CaseV1{
	{
		Title: "Dir RepoDir returns non-empty -- repo directory",
		ArrangeInput: args.Map{
			"when": "calling RepoDir",
		},
		ExpectedInput: args.Map{
			"notEmpty": true,
		},
	},
}

var extFileGetterPathTestCases = []coretestcases.CaseV1{
	{
		Title: "File Path returns non-empty -- skip 0",
		ArrangeInput: args.Map{
			"when": "calling File.Path",
		},
		ExpectedInput: args.Map{
			"notEmpty": true,
		},
	},
}

var extFileGetterNameTestCases = []coretestcases.CaseV1{
	{
		Title: "File Name returns non-empty -- skip 0",
		ArrangeInput: args.Map{
			"when": "calling File.Name",
		},
		ExpectedInput: args.Map{
			"notEmpty": true,
		},
	},
}

var extFileGetterCurrentFilePathTestCases = []coretestcases.CaseV1{
	{
		Title: "File CurrentFilePath returns non-empty -- current file",
		ArrangeInput: args.Map{
			"when": "calling CurrentFilePath",
		},
		ExpectedInput: args.Map{
			"notEmpty": true,
		},
	},
}

// =============================================================================
// newCreator
// =============================================================================

var extNewCreatorDefaultTestCases = []coretestcases.CaseV1{
	{
		Title: "newCreator Default returns valid trace -- default skip",
		ArrangeInput: args.Map{
			"when": "calling New.Default()",
		},
		ExpectedInput: args.Map{
			"isOkay":       true,
			"hasFilePath":  true,
			"hasMethodName": true,
		},
	},
}

var extNewCreatorSkipOneTestCases = []coretestcases.CaseV1{
	{
		Title: "newCreator SkipOne returns valid trace -- skip one",
		ArrangeInput: args.Map{
			"when": "calling New.SkipOne()",
		},
		ExpectedInput: args.Map{
			"isOkay":      true,
			"hasFilePath": true,
		},
	},
}

var extNewCreatorPtrTestCases = []coretestcases.CaseV1{
	{
		Title: "newCreator Ptr returns non-nil trace ptr -- skip 0",
		ArrangeInput: args.Map{
			"when": "calling New.Ptr(0)",
		},
		ExpectedInput: args.Map{
			"isNil":  false,
			"isOkay": true,
		},
	},
}

// =============================================================================
// TraceCollection missing test case vars needed by TraceCollection_test.go
// =============================================================================

var traceCollectionSkipTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection Skip 1 -- returns remaining items",
		ArrangeInput: args.Map{
			"when": "skip first item",
			"skip": 1,
		},
		ExpectedInput: args.Map{
			"length":   2,
			"firstPkg": "pkg2",
		},
	},
}

var traceCollectionTakeTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection Take 2 -- returns first 2 items",
		ArrangeInput: args.Map{
			"when": "take first 2",
			"take": 2,
		},
		ExpectedInput: args.Map{
			"length":  2,
			"lastPkg": "pkg2",
		},
	},
}

var traceCollectionAddTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection Add -- two items added",
		ArrangeInput: args.Map{"when": "adding two items"},
		ExpectedInput: args.Map{
			"length":   2,
			"firstPkg": "add1",
			"lastPkg":  "add2",
		},
	},
}

var traceCollectionAddsTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection Adds -- three items added",
		ArrangeInput: args.Map{"when": "adding three items at once"},
		ExpectedInput: args.Map{
			"length": 3,
		},
	},
}

var traceCollectionAddsIfTrueTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection AddsIf true -- item added",
		ArrangeInput: args.Map{"when": "isAdd true"},
		ExpectedInput: args.Map{"length": 1},
	},
}

var traceCollectionAddsIfFalseTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection AddsIf false -- item not added",
		ArrangeInput: args.Map{"when": "isAdd false"},
		ExpectedInput: args.Map{"length": 0},
	},
}

var traceCollectionFirstOrDefaultEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection FirstOrDefault empty -- returns nil",
		ArrangeInput: args.Map{"when": "empty collection"},
		ExpectedInput: args.Map{"isNil": true},
	},
}

var traceCollectionFirstOrDefaultNonEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection FirstOrDefault non-empty -- returns first",
		ArrangeInput: args.Map{"when": "non-empty collection"},
		ExpectedInput: args.Map{
			"isNil":       false,
			"packageName": "pkg1",
		},
	},
}

var traceCollectionLastOrDefaultEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection LastOrDefault empty -- returns nil",
		ArrangeInput: args.Map{"when": "empty collection"},
		ExpectedInput: args.Map{"isNil": true},
	},
}

var traceCollectionLastOrDefaultNonEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection LastOrDefault non-empty -- returns last",
		ArrangeInput: args.Map{"when": "non-empty collection"},
		ExpectedInput: args.Map{
			"isNil":       false,
			"packageName": "pkg3",
		},
	},
}

var traceCollectionStringsTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection Strings -- returns 3 strings",
		ArrangeInput: args.Map{"when": "3 item collection"},
		ExpectedInput: args.Map{"length": 3},
	},
}

var traceCollectionFilterFuncTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection Filter line > 10 -- returns 2 items",
		ArrangeInput: args.Map{"when": "filter by line > 10"},
		ExpectedInput: args.Map{"length": 2},
	},
}

var traceCollectionSkipFilterPkgTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection SkipFilterPackageName pkg2 -- returns 2 items",
		ArrangeInput: args.Map{"when": "skip filter pkg2"},
		ExpectedInput: args.Map{"length": 2},
	},
}

var traceCollectionFilterMethodTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection FilterMethodName TestMethod -- returns 3 items",
		ArrangeInput: args.Map{"when": "filter by TestMethod"},
		ExpectedInput: args.Map{"length": 3},
	},
}

var traceCollectionCloneTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection Clone -- preserves items",
		ArrangeInput: args.Map{"when": "cloning 3 item collection"},
		ExpectedInput: args.Map{
			"length":   3,
			"firstPkg": "pkg1",
		},
	},
}

var traceCollectionClonePtrTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection ClonePtr -- returns non-nil clone",
		ArrangeInput: args.Map{"when": "cloning to ptr"},
		ExpectedInput: args.Map{
			"isNil":  false,
			"length": 3,
		},
	},
}

var traceCollectionClonePtrNilTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection ClonePtr nil -- returns nil",
		ArrangeInput: args.Map{"when": "nil collection"},
		ExpectedInput: args.Map{"isNil": true},
	},
}

var traceCollectionIsEqualTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection IsEqual -- same items returns true",
		ArrangeInput: args.Map{"when": "comparing equal collections"},
		ExpectedInput: args.Map{"isEqual": true},
	},
}

var traceCollectionIsEqualBothNilTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection IsEqual both nil -- returns true",
		ArrangeInput: args.Map{"when": "both nil"},
		ExpectedInput: args.Map{"isEqual": true},
	},
}

var traceCollectionIsEqualOneNilTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection IsEqual one nil -- returns false",
		ArrangeInput: args.Map{"when": "one nil"},
		ExpectedInput: args.Map{"isEqual": false},
	},
}

var traceCollectionLengthNilTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection Length nil -- returns 0",
		ArrangeInput: args.Map{"when": "nil collection"},
		ExpectedInput: args.Map{"length": 0},
	},
}

var traceCollectionHasIndexTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection HasIndex 2 -- returns true for 3 items",
		ArrangeInput: args.Map{
			"when":  "index within range",
			"index": 2,
		},
		ExpectedInput: args.Map{"hasIndex": true},
	},
	{
		Title: "TraceCollection HasIndex 5 -- returns false for 3 items",
		ArrangeInput: args.Map{
			"when":  "index out of range",
			"index": 5,
		},
		ExpectedInput: args.Map{"hasIndex": false},
	},
}

var traceCollectionReverseEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection Reverse empty -- returns empty",
		ArrangeInput: args.Map{"when": "empty collection"},
		ExpectedInput: args.Map{"length": 0},
	},
}

var traceCollectionReverseTwoTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection Reverse two -- swaps items",
		ArrangeInput: args.Map{"when": "two item collection"},
		ExpectedInput: args.Map{
			"firstPkg": "second",
			"lastPkg":  "first",
		},
	},
}

var traceCollectionConcatNewTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection ConcatNew -- creates new with extra item",
		ArrangeInput: args.Map{"when": "concat new trace"},
		ExpectedInput: args.Map{
			"newLength":      4,
			"originalLength": 3,
		},
	},
}

var traceCollectionDisposeTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection Dispose -- items nil",
		ArrangeInput: args.Map{"when": "dispose called"},
		ExpectedInput: args.Map{"itemsNil": true},
	},
}

var traceCollectionClearTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection Clear -- length 0",
		ArrangeInput: args.Map{"when": "clear called"},
		ExpectedInput: args.Map{"length": 0},
	},
}

var traceCollectionClearNilTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection Clear nil -- returns nil",
		ArrangeInput: args.Map{"when": "nil clear"},
		ExpectedInput: args.Map{"isNil": true},
	},
}

var traceCollectionCodeStacksStringTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection CodeStacksString -- returns non-empty",
		ArrangeInput: args.Map{"when": "non-empty collection"},
		ExpectedInput: args.Map{"notEmpty": true},
	},
}

var traceCollectionCodeStacksStringEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection CodeStacksString empty -- returns empty",
		ArrangeInput: args.Map{"when": "empty collection"},
		ExpectedInput: args.Map{"isEmpty": true},
	},
}

var traceCollectionJsonStringTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection JsonString -- returns non-empty",
		ArrangeInput: args.Map{"when": "non-empty collection"},
		ExpectedInput: args.Map{"notEmpty": true},
	},
}

var traceCollectionJsonStringEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection JsonString empty -- returns empty",
		ArrangeInput: args.Map{"when": "empty collection"},
		ExpectedInput: args.Map{"isEmpty": true},
	},
}

var traceCollectionSerializerTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection Serializer -- returns bytes no error",
		ArrangeInput: args.Map{"when": "non-empty collection"},
		ExpectedInput: args.Map{
			"hasError": false,
			"notEmpty": true,
		},
	},
}

var traceCollectionStackTracesBytesTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection StackTracesBytes -- returns non-empty bytes",
		ArrangeInput: args.Map{"when": "non-empty collection"},
		ExpectedInput: args.Map{"notEmpty": true},
	},
}

var traceCollectionStackTracesBytesEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection StackTracesBytes empty -- returns empty",
		ArrangeInput: args.Map{"when": "empty collection"},
		ExpectedInput: args.Map{"length": 0},
	},
}

var traceCollectionGetPagesSizeTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection GetPagesSize 2 -- returns 2 pages for 3 items",
		ArrangeInput: args.Map{
			"when":     "page size 2",
			"pageSize": 2,
		},
		ExpectedInput: args.Map{"pages": 2},
	},
	{
		Title: "TraceCollection GetPagesSize 0 -- returns 0 pages",
		ArrangeInput: args.Map{
			"when":     "page size 0",
			"pageSize": 0,
		},
		ExpectedInput: args.Map{"pages": 0},
	},
}

var traceCollectionLimitTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection LimitCollection 2 -- returns 2 items",
		ArrangeInput: args.Map{"when": "limit 2"},
		ExpectedInput: args.Map{"length": 2},
	},
}

var traceCollectionSafeLimitTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection SafeLimitCollection 10 -- returns 3 (capped)",
		ArrangeInput: args.Map{"when": "limit exceeds length"},
		ExpectedInput: args.Map{"length": 3},
	},
}

var traceCollectionFileWithLinesTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection FileWithLines -- returns 3 items",
		ArrangeInput: args.Map{"when": "3 item collection"},
		ExpectedInput: args.Map{"length": 3},
	},
}

var traceCollectionShortStringsTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection ShortStrings -- returns 3 strings",
		ArrangeInput: args.Map{"when": "3 item collection"},
		ExpectedInput: args.Map{"length": 3},
	},
}

var traceCollectionJoinUsingFmtTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection JoinUsingFmt -- returns joined string",
		ArrangeInput: args.Map{"when": "custom formatter with comma"},
		ExpectedInput: args.Map{"result": "pkg1,pkg2,pkg3"},
	},
}

var traceCollectionCsvStringsTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection CsvStrings -- returns 3 quoted strings",
		ArrangeInput: args.Map{"when": "3 item collection"},
		ExpectedInput: args.Map{"length": 3},
	},
}

var traceCollectionAddsPtrNilTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection AddsPtr nil trace -- skips nil",
		ArrangeInput: args.Map{"when": "nil trace pointer"},
		ExpectedInput: args.Map{"length": 0},
	},
}

var traceCollectionAddsPtrSkipTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection AddsPtr skip issues -- skips bad trace",
		ArrangeInput: args.Map{"when": "trace with issues"},
		ExpectedInput: args.Map{"length": 0},
	},
}

var traceCollectionStringTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection String -- returns non-empty",
		ArrangeInput: args.Map{"when": "non-empty collection"},
		ExpectedInput: args.Map{"notEmpty": true},
	},
}

var traceCollectionStringEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection String empty -- returns empty",
		ArrangeInput: args.Map{"when": "empty collection"},
		ExpectedInput: args.Map{"isEmpty": true},
	},
}
