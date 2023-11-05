package coredynamictestwrappers

import (
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/corevalidator"
	"gitlab.com/auk-go/core/enums/stringcompareas"
)

var (
	ReflectSetFromToInvalidTestCases = []ReflectSetFromToTestWrapper{
		// {
		// 	Header: "(null, null) -- do nothing -- " +
		// 		"From `Null` to `Null` -- does nothing -- no error",
		// 	From:             nil,
		// 	To:               nil,
		// 	ExpectedValue:    "",
		// 	IsExpectingError: false,
		// 	HasPanic:         false,
		// 	Validator: corevalidator.TextValidator{
		// 		Search:                 "",
		// 		SearchAs:               stringcompareas.Equal,
		// 		ValidatorCoreCondition: corevalidator.DefaultTrimCoreCondition,
		// 	},
		// },
		{
			Header: "(null, valid type - coretests.DraftType) -- should panic -- " +
				"From `Null` to `coretests.DraftType`",
			From: nil,
			To: &coretests.DraftType{
				SampleString1: "Same data",
			},
			// ExpectedValue:    &ReflectSetFromToTestCasesDraftTypeExpected,
			ExpectedValue:    "Invalid : value cannot process it. `from` is nil, cannot set null or nil to destination.\"! Supported Types: https://t.ly/1Lpt,  Ref(s) { \"(FromType, ToType) = (<nil>, *coretests.DraftType)\" }",
			IsExpectingError: false,
			HasPanic:         true,
			Validator: corevalidator.TextValidator{
				Search:                 "Invalid : value cannot process it. `from` is nil, cannot set null or nil to destination.\"! Supported Types: https://t.ly/1Lpt,  Ref(s) { \"(FromType, ToType) = (<nil>, *coretests.DraftType)\" }",
				SearchAs:               stringcompareas.Equal,
				ValidatorCoreCondition: corevalidator.DefaultTrimCoreCondition,
			},
		},
		// {
		// 	Header: "(sameTypePointer, sameTypePointer) -- try reflection -- " +
		// 		"From `*ReflectSetFromToTestWrapper{Expected}` " +
		// 		"to   `*ReflectSetFromToTestWrapper{Sample data}` should set to Expected. ",
		// 	From: &ReflectSetFromToTestCasesDraftTypeExpected,
		// 	To: &coretests.DraftType{
		// 		SampleString1: "Same data",
		// 	},
		// 	ExpectedValue: &ReflectSetFromToTestCasesDraftTypeExpected,
		// },
		// {
		// 	Header: "(sameTypeNonPointer, sameTypePointer) -- try reflection -- " +
		// 		"From `ReflectSetFromToTestWrapper{Expected}` " +
		// 		"to   `*ReflectSetFromToTestWrapper{Sample data}` should set to Expected.",
		// 	From: ReflectSetFromToTestCasesDraftTypeExpected,
		// 	To: &coretests.DraftType{
		// 		SampleString1: "Sample data",
		// 	},
		// 	ExpectedValue: &ReflectSetFromToTestCasesDraftTypeExpected,
		// },
		// {
		// 	Header: "(*[]byte, otherType) -- try unmarshal, reflect -- " +
		// 		"From `*[]bytes(ReflectSetFromToTestWrapper{Expected}` " +
		// 		"to   `*ReflectSetFromToTestWrapper{Sample data}` should set to Expected.",
		// 	From: ReflectSetFromToTestCasesDraftTypeExpected.JsonBytesPtr(),
		// 	To: &coretests.DraftType{
		// 		SampleString1: "Sample data",
		// 	},
		// 	ExpectedValue: &ReflectSetFromToTestCasesDraftTypeExpected,
		// },
		// {
		// 	Header: "(otherType, *[]byte) -- try marshal, reflect -- " +
		// 		"From `ReflectSetFromToTestWrapper{Expected}` " +
		// 		"to   `*[]byte{}` should set to Expected.",
		// 	From:          ReflectSetFromToTestCasesDraftTypeExpected.JsonBytesPtr(),
		// 	To:            &[]byte{},
		// 	ExpectedValue: ReflectSetFromToTestCasesDraftTypeExpected.JsonBytesPtr(),
		// },
	}
)
