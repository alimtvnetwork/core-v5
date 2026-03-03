package coreinstructiontests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================================================
// ClonePtr
// ==========================================================================

var fromToClonePtrCopiesTestCase = coretestcases.CaseV1{
	Title: "ClonePtr - copies From and To",
	ExpectedInput: []string{
		"true",
		"source",
		"destination",
	},
}

var fromToClonePtrNilTestCase = coretestcases.CaseV1{
	Title: "ClonePtr - nil receiver returns nil",
	ExpectedInput: []string{
		"true",
	},
}

// ==========================================================================
// Clone
// ==========================================================================

var fromToCloneCopiesTestCase = coretestcases.CaseV1{
	Title: "Clone - copies values",
	ExpectedInput: []string{
		"a",
		"b",
	},
}

// ==========================================================================
// IsNull
// ==========================================================================

var fromToIsNullNilTestCase = coretestcases.CaseV1{
	Title: "IsNull - nil returns true",
	ExpectedInput: []string{
		"true",
	},
}

var fromToIsNullNonNilTestCase = coretestcases.CaseV1{
	Title: "IsNull - non-nil returns false",
	ExpectedInput: []string{
		"false",
	},
}

// ==========================================================================
// IsFromEmpty
// ==========================================================================

var fromToIsFromEmptyEmptyTestCase = coretestcases.CaseV1{
	Title: "IsFromEmpty - empty From returns true",
	ExpectedInput: []string{
		"true",
	},
}

var fromToIsFromEmptyNilTestCase = coretestcases.CaseV1{
	Title: "IsFromEmpty - nil receiver returns true",
	ExpectedInput: []string{
		"true",
	},
}

// ==========================================================================
// IsToEmpty
// ==========================================================================

var fromToIsToEmptyEmptyTestCase = coretestcases.CaseV1{
	Title: "IsToEmpty - empty To returns true",
	ExpectedInput: []string{
		"true",
	},
}

var fromToIsToEmptyNonEmptyTestCase = coretestcases.CaseV1{
	Title: "IsToEmpty - non-empty returns false",
	ExpectedInput: []string{
		"false",
	},
}

// ==========================================================================
// String
// ==========================================================================

var fromToStringContainsTestCase = coretestcases.CaseV1{
	Title: "String - contains From and To",
	ExpectedInput: []string{
		"true",
		"true",
	},
}

// ==========================================================================
// FromName / ToName
// ==========================================================================

var fromToNamesTestCase = coretestcases.CaseV1{
	Title: "FromName/ToName return field values",
	ExpectedInput: []string{
		"src",
		"dst",
	},
}

// ==========================================================================
// SetFromName
// ==========================================================================

var fromToSetFromNameUpdatesTestCase = coretestcases.CaseV1{
	Title: "SetFromName - updates From",
	ExpectedInput: []string{
		"new",
	},
}

var fromToSetFromNameNilTestCase = coretestcases.CaseV1{
	Title: "SetFromName - nil receiver no panic",
	ExpectedInput: []string{
		"true",
	},
}

// ==========================================================================
// SetToName
// ==========================================================================

var fromToSetToNameUpdatesTestCase = coretestcases.CaseV1{
	Title: "SetToName - updates To",
	ExpectedInput: []string{
		"new",
	},
}

// ==========================================================================
// SourceDestination
// ==========================================================================

var fromToSourceDestMapsTestCase = coretestcases.CaseV1{
	Title: "SourceDestination - maps From->Source To->Destination",
	ExpectedInput: []string{
		"true",
		"src",
		"dst",
	},
}

var fromToSourceDestNilTestCase = coretestcases.CaseV1{
	Title: "SourceDestination - nil returns nil",
	ExpectedInput: []string{
		"true",
	},
}

// ==========================================================================
// Rename
// ==========================================================================

var fromToRenameMapsTestCase = coretestcases.CaseV1{
	Title: "Rename - maps From->Existing To->New",
	ExpectedInput: []string{
		"true",
		"old",
		"new",
	},
}

var fromToRenameNilTestCase = coretestcases.CaseV1{
	Title: "Rename - nil returns nil",
	ExpectedInput: []string{
		"true",
	},
}
