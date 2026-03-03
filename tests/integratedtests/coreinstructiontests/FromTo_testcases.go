package coreinstructiontests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var fromToClonePtrTestCases = []coretestcases.CaseV1{
	{
		Title: "ClonePtr - copies From and To",
		ExpectedInput: []string{
			"true",
			"source",
			"destination",
		},
	},
	{
		Title: "ClonePtr - nil receiver returns nil",
		ExpectedInput: []string{
			"true",
		},
	},
}

var fromToCloneTestCases = []coretestcases.CaseV1{
	{
		Title: "Clone - copies values",
		ExpectedInput: []string{
			"a",
			"b",
		},
	},
}

var fromToIsNullTestCases = []coretestcases.CaseV1{
	{
		Title: "IsNull - nil returns true",
		ExpectedInput: []string{
			"true",
		},
	},
	{
		Title: "IsNull - non-nil returns false",
		ExpectedInput: []string{
			"false",
		},
	},
}

var fromToIsFromEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "IsFromEmpty - empty From returns true",
		ExpectedInput: []string{
			"true",
		},
	},
	{
		Title: "IsFromEmpty - nil receiver returns true",
		ExpectedInput: []string{
			"true",
		},
	},
}

var fromToIsToEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "IsToEmpty - empty To returns true",
		ExpectedInput: []string{
			"true",
		},
	},
	{
		Title: "IsToEmpty - non-empty returns false",
		ExpectedInput: []string{
			"false",
		},
	},
}

var fromToStringTestCases = []coretestcases.CaseV1{
	{
		Title: "String - contains From and To",
		ExpectedInput: []string{
			"true",
			"true",
		},
	},
}

var fromToNamesTestCases = []coretestcases.CaseV1{
	{
		Title: "FromName/ToName return field values",
		ExpectedInput: []string{
			"src",
			"dst",
		},
	},
}

var fromToSetFromNameTestCases = []coretestcases.CaseV1{
	{
		Title: "SetFromName - updates From",
		ExpectedInput: []string{
			"new",
		},
	},
	{
		Title: "SetFromName - nil receiver no panic",
		ExpectedInput: []string{
			"true",
		},
	},
}

var fromToSetToNameTestCases = []coretestcases.CaseV1{
	{
		Title: "SetToName - updates To",
		ExpectedInput: []string{
			"new",
		},
	},
}

var fromToSourceDestTestCases = []coretestcases.CaseV1{
	{
		Title: "SourceDestination - maps From->Source To->Destination",
		ExpectedInput: []string{
			"true",
			"src",
			"dst",
		},
	},
	{
		Title: "SourceDestination - nil returns nil",
		ExpectedInput: []string{
			"true",
		},
	},
}

var fromToRenameTestCases = []coretestcases.CaseV1{
	{
		Title: "Rename - maps From->Existing To->New",
		ExpectedInput: []string{
			"true",
			"old",
			"new",
		},
	},
	{
		Title: "Rename - nil returns nil",
		ExpectedInput: []string{
			"true",
		},
	},
}
