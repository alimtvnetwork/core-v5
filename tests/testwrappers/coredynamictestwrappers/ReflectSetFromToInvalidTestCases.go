package coredynamictestwrappers

import (
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/corevalidator"
)

var (
	ReflectSetFromToInvalidTestCases = []ReflectSetFromToTestWrapper{
		{
			Header: "(null, null) -- do nothing -- " +
				"From `Null` to `Null` -- does nothing -- no error",
			From:             nil,
			To:               nil,
			ExpectedValue:    &ReflectSetFromToTestCasesDraftTypeExpected,
			IsErrorExpected:  true,
			IsPanic:          true,
			PanicExpectation: "",
			Validator:        corevalidator.TextValidator{},
		},
		{
			Header: "(null, valid type - coretests.DraftType) -- should panic -- " +
				"From `Null` to `coretests.DraftType`",
			From: nil,
			To: &coretests.DraftType{
				SampleString1: "Same data",
			},
			ExpectedValue: &ReflectSetFromToTestCasesDraftTypeExpected,
		},
		{
			Header: "(sameTypePointer, sameTypePointer) -- try reflection -- " +
				"From `*ReflectSetFromToTestWrapper{Expected}` " +
				"to   `*ReflectSetFromToTestWrapper{Sample data}` should set to Expected. ",
			From: &ReflectSetFromToTestCasesDraftTypeExpected,
			To: &coretests.DraftType{
				SampleString1: "Same data",
			},
			ExpectedValue: &ReflectSetFromToTestCasesDraftTypeExpected,
		},
		{
			Header: "(sameTypeNonPointer, sameTypePointer) -- try reflection -- " +
				"From `ReflectSetFromToTestWrapper{Expected}` " +
				"to   `*ReflectSetFromToTestWrapper{Sample data}` should set to Expected.",
			From: ReflectSetFromToTestCasesDraftTypeExpected,
			To: &coretests.DraftType{
				SampleString1: "Sample data",
			},
			ExpectedValue: &ReflectSetFromToTestCasesDraftTypeExpected,
		},
		{
			Header: "(*[]byte, otherType) -- try unmarshal, reflect -- " +
				"From `*[]bytes(ReflectSetFromToTestWrapper{Expected}` " +
				"to   `*ReflectSetFromToTestWrapper{Sample data}` should set to Expected.",
			From: ReflectSetFromToTestCasesDraftTypeExpected.JsonBytesPtr(),
			To: &coretests.DraftType{
				SampleString1: "Sample data",
			},
			ExpectedValue: &ReflectSetFromToTestCasesDraftTypeExpected,
		},
		{
			Header: "(otherType, *[]byte) -- try marshal, reflect -- " +
				"From `ReflectSetFromToTestWrapper{Expected}` " +
				"to   `*[]byte{}` should set to Expected.",
			From:          ReflectSetFromToTestCasesDraftTypeExpected.JsonBytesPtr(),
			To:            &[]byte{},
			ExpectedValue: ReflectSetFromToTestCasesDraftTypeExpected.JsonBytesPtr(),
		},
	}
)
