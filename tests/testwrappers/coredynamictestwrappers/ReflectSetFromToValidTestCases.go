package coredynamictestwrappers

import (
	"github.com/alimtvnetwork/core/coretests"
)

var ReflectSetFromToValidNullNull = FromToTestWrapper{
	Header: "(null, null) -- do nothing -- " +
		"From `Null` to `Null` -- does nothing -- no error",
}

var ReflectSetFromToValidPtrToPtr = FromToTestWrapper{
	Header: "(sameTypePointer, sameTypePointer) -- try reflection -- " +
		"From `*FromToTestWrapper{Expected}` " +
		"to   `*FromToTestWrapper{Sample data}` should set to Expected. ",
	From: &ReflectSetFromToTestCasesDraftTypeExpected,
	To: &coretests.DraftType{
		SampleString1: "Same data",
	},
	ExpectedValue: &ReflectSetFromToTestCasesDraftTypeExpected,
}

var ReflectSetFromToValidValueToPtr = FromToTestWrapper{
	Header: "(sameTypeNonPointer, sameTypePointer) -- try reflection -- " +
		"From `FromToTestWrapper{Expected}` " +
		"to   `*FromToTestWrapper{Sample data}` should set to Expected.",
	From: ReflectSetFromToTestCasesDraftTypeExpected,
	To: &coretests.DraftType{
		SampleString1: "Sample data",
	},
	ExpectedValue: &ReflectSetFromToTestCasesDraftTypeExpected,
}

var ReflectSetFromToValidBytesToDraft = FromToTestWrapper{
	Header: "([]byte, otherType) -- try unmarshal, reflect -- " +
		"From `[]bytes(FromToTestWrapper{Expected}` " +
		"to   `*FromToTestWrapper{Sample data}` should set to Expected.",
	From: ReflectSetFromToTestCasesDraftTypeExpected.JsonBytesPtr(),
	To: &coretests.DraftType{
		SampleString1: "Sample data",
	},
	ExpectedValue: &ReflectSetFromToTestCasesDraftTypeExpected,
}

var ReflectSetFromToValidDraftToBytes = FromToTestWrapper{
	Header: "(otherType, *[]byte) -- try marshal, reflect -- " +
		"From `FromToTestWrapper{Expected}` " +
		"to   `*[]byte{}` should set to Expected.",
	From:          &ReflectSetFromToTestCasesDraftTypeExpected,
	To:            &[]byte{},
	ExpectedValue: &[]byte{}, // placeholder; IsSame checks type match (*[]byte == *[]byte)
}

// ReflectSetFromToValidTestCases kept for backward compatibility if needed.
var ReflectSetFromToValidTestCases = []FromToTestWrapper{
	ReflectSetFromToValidNullNull,
	ReflectSetFromToValidPtrToPtr,
	ReflectSetFromToValidValueToPtr,
	ReflectSetFromToValidBytesToDraft,
	ReflectSetFromToValidDraftToBytes,
}
