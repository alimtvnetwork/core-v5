package typesconvtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var boolPtrTestCases = []coretestcases.CaseV1{
	{
		Title:         "BoolPtr true",
		ArrangeInput:  args.Map{"value": true},
		ExpectedInput: args.Map{"notNil": true, "deref": true},
	},
	{
		Title:         "BoolPtr false",
		ArrangeInput:  args.Map{"value": false},
		ExpectedInput: args.Map{"notNil": true, "deref": false},
	},
}

var boolPtrToSimpleTestCases = []coretestcases.CaseV1{
	{
		Title:         "nil returns false",
		ArrangeInput:  args.Map{"isNil": true},
		ExpectedInput: args.Map{"result": false},
	},
	{
		Title:         "non-nil true returns true",
		ArrangeInput:  args.Map{"isNil": false, "value": true},
		ExpectedInput: args.Map{"result": true},
	},
}

var boolPtrToSimpleDefTestCases = []coretestcases.CaseV1{
	{
		Title:         "nil returns default true",
		ArrangeInput:  args.Map{"isNil": true, "defVal": true},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "non-nil false ignores default",
		ArrangeInput:  args.Map{"isNil": false, "value": false, "defVal": true},
		ExpectedInput: args.Map{"result": false},
	},
}

var boolPtrToDefPtrTestCases = []coretestcases.CaseV1{
	{
		Title:         "nil returns defVal ptr",
		ArrangeInput:  args.Map{"isNil": true, "defVal": true},
		ExpectedInput: args.Map{"deref": true},
	},
	{
		Title:         "non-nil returns original",
		ArrangeInput:  args.Map{"isNil": false, "value": false, "defVal": true},
		ExpectedInput: args.Map{"deref": false},
	},
}

var boolPtrDefValFuncTestCases = []coretestcases.CaseV1{
	{
		Title:         "nil calls func",
		ArrangeInput:  args.Map{"isNil": true},
		ExpectedInput: args.Map{"deref": true},
	},
	{
		Title:         "non-nil returns original",
		ArrangeInput:  args.Map{"isNil": false, "value": false},
		ExpectedInput: args.Map{"deref": false},
	},
}

var intPtrTestCases = []coretestcases.CaseV1{
	{
		Title:         "IntPtr creates pointer",
		ArrangeInput:  args.Map{"value": 42},
		ExpectedInput: args.Map{"notNil": true, "deref": 42},
	},
}

var intPtrToSimpleTestCases = []coretestcases.CaseV1{
	{
		Title:         "nil returns 0",
		ArrangeInput:  args.Map{"isNil": true},
		ExpectedInput: args.Map{"result": 0},
	},
	{
		Title:         "non-nil returns value",
		ArrangeInput:  args.Map{"isNil": false, "value": 7},
		ExpectedInput: args.Map{"result": 7},
	},
}

var intPtrToSimpleDefTestCases = []coretestcases.CaseV1{
	{
		Title:         "nil returns default",
		ArrangeInput:  args.Map{"isNil": true, "defVal": 99},
		ExpectedInput: args.Map{"result": 99},
	},
}

var intPtrToDefPtrTestCases = []coretestcases.CaseV1{
	{
		Title:         "nil returns defVal ptr",
		ArrangeInput:  args.Map{"isNil": true, "defVal": 10},
		ExpectedInput: args.Map{"deref": 10},
	},
}

var intPtrDefValFuncTestCases = []coretestcases.CaseV1{
	{
		Title:         "nil calls func",
		ArrangeInput:  args.Map{"isNil": true},
		ExpectedInput: args.Map{"deref": 55},
	},
}

var stringPtrTestCases = []coretestcases.CaseV1{
	{
		Title:         "StringPtr creates pointer",
		ArrangeInput:  args.Map{"value": "hello"},
		ExpectedInput: args.Map{"notNil": true, "deref": "hello"},
	},
}

var stringPtrToSimpleTestCases = []coretestcases.CaseV1{
	{
		Title:         "nil returns empty",
		ArrangeInput:  args.Map{"isNil": true},
		ExpectedInput: args.Map{"result": ""},
	},
	{
		Title:         "non-nil returns value",
		ArrangeInput:  args.Map{"isNil": false, "value": "world"},
		ExpectedInput: args.Map{"result": "world"},
	},
}

var stringPtrToSimpleDefTestCases = []coretestcases.CaseV1{
	{
		Title:         "nil returns default",
		ArrangeInput:  args.Map{"isNil": true, "defVal": "fallback"},
		ExpectedInput: args.Map{"result": "fallback"},
	},
}

var stringPtrToDefPtrTestCases = []coretestcases.CaseV1{
	{
		Title:         "nil returns defVal ptr",
		ArrangeInput:  args.Map{"isNil": true, "defVal": "default"},
		ExpectedInput: args.Map{"deref": "default"},
	},
}

var stringPtrDefValFuncTestCases = []coretestcases.CaseV1{
	{
		Title:         "nil calls func",
		ArrangeInput:  args.Map{"isNil": true},
		ExpectedInput: args.Map{"deref": "generated"},
	},
}

var stringToBoolTestCases = []coretestcases.CaseV1{
	{
		Title:         "empty string returns false",
		ArrangeInput:  args.Map{"value": ""},
		ExpectedInput: args.Map{"result": false},
	},
	{
		Title:         "yes returns true",
		ArrangeInput:  args.Map{"value": "yes"},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "Yes returns true",
		ArrangeInput:  args.Map{"value": "Yes"},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "no returns false",
		ArrangeInput:  args.Map{"value": "no"},
		ExpectedInput: args.Map{"result": false},
	},
	{
		Title:         "true returns true",
		ArrangeInput:  args.Map{"value": "true"},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "invalid returns false",
		ArrangeInput:  args.Map{"value": "xyz"},
		ExpectedInput: args.Map{"result": false},
	},
}

var stringPointerToBoolTestCases = []coretestcases.CaseV1{
	{
		Title:         "nil returns false",
		ArrangeInput:  args.Map{"isNil": true},
		ExpectedInput: args.Map{"result": false},
	},
	{
		Title:         "empty returns false",
		ArrangeInput:  args.Map{"isNil": false, "value": ""},
		ExpectedInput: args.Map{"result": false},
	},
	{
		Title:         "yes returns true",
		ArrangeInput:  args.Map{"isNil": false, "value": "yes"},
		ExpectedInput: args.Map{"result": true},
	},
}

var stringPointerToBoolPtrTestCases = []coretestcases.CaseV1{
	{
		Title:         "nil returns false ptr",
		ArrangeInput:  args.Map{"isNil": true},
		ExpectedInput: args.Map{"deref": false},
	},
	{
		Title:         "yes returns true ptr",
		ArrangeInput:  args.Map{"isNil": false, "value": "yes"},
		ExpectedInput: args.Map{"deref": true},
	},
}

var stringToBoolPtrTestCases = []coretestcases.CaseV1{
	{
		Title:         "empty returns false ptr",
		ArrangeInput:  args.Map{"value": ""},
		ExpectedInput: args.Map{"deref": false},
	},
	{
		Title:         "true returns true ptr",
		ArrangeInput:  args.Map{"value": "true"},
		ExpectedInput: args.Map{"deref": true},
	},
}

var bytePtrTestCases = []coretestcases.CaseV1{
	{
		Title:         "BytePtr creates pointer",
		ArrangeInput:  args.Map{"value": byte(5)},
		ExpectedInput: args.Map{"notNil": true},
	},
}

var bytePtrToSimpleTestCases = []coretestcases.CaseV1{
	{
		Title:         "nil returns 0",
		ArrangeInput:  args.Map{"isNil": true},
		ExpectedInput: args.Map{"result": byte(0)},
	},
}

var bytePtrToSimpleDefTestCases = []coretestcases.CaseV1{
	{
		Title:         "nil returns default",
		ArrangeInput:  args.Map{"isNil": true, "defVal": byte(9)},
		ExpectedInput: args.Map{"result": byte(9)},
	},
}

var bytePtrToDefPtrTestCases = []coretestcases.CaseV1{
	{
		Title:         "nil returns defVal ptr",
		ArrangeInput:  args.Map{"isNil": true, "defVal": byte(3)},
		ExpectedInput: args.Map{"notNil": true},
	},
}

var bytePtrDefValFuncTestCases = []coretestcases.CaseV1{
	{
		Title:         "nil calls func",
		ArrangeInput:  args.Map{"isNil": true},
		ExpectedInput: args.Map{"notNil": true},
	},
}

var floatPtrTestCases = []coretestcases.CaseV1{
	{
		Title:         "FloatPtr creates pointer",
		ArrangeInput:  args.Map{"value": float32(3.14)},
		ExpectedInput: args.Map{"notNil": true},
	},
}

var floatPtrToSimpleTestCases = []coretestcases.CaseV1{
	{
		Title:         "nil returns 0",
		ArrangeInput:  args.Map{"isNil": true},
		ExpectedInput: args.Map{"isZero": true},
	},
}

var floatPtrToSimpleDefTestCases = []coretestcases.CaseV1{
	{
		Title:         "nil returns default",
		ArrangeInput:  args.Map{"isNil": true, "defVal": float32(1.5)},
		ExpectedInput: args.Map{"result": float32(1.5)},
	},
}

var floatPtrToDefPtrTestCases = []coretestcases.CaseV1{
	{
		Title:         "nil returns defVal ptr",
		ArrangeInput:  args.Map{"isNil": true},
		ExpectedInput: args.Map{"notNil": true},
	},
}

var floatPtrDefValFuncTestCases = []coretestcases.CaseV1{
	{
		Title:         "nil calls func",
		ArrangeInput:  args.Map{"isNil": true},
		ExpectedInput: args.Map{"notNil": true},
	},
}
