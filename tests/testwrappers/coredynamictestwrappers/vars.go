package coredynamictestwrappers

import "github.com/alimtvnetwork/core/coretests"

var (
	ReflectSetFromToTestCasesDraftTypeInput = coretests.DraftType{
		SampleString1: "Same data",
		SampleString2: "",
		SampleInteger: 0,
	}
	ReflectSetFromToTestCasesDraftTypeExpected = coretests.DraftType{
		SampleString1: "Expected",
		SampleString2: "",
		SampleInteger: 0,
	}

	nilBytes []byte = nil
)
