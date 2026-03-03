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
	Title:         "ClonePtr - nil receiver returns nil",
	ExpectedInput: "true",
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
	Title:         "IsNull - nil returns true",
	ExpectedInput: "true",
}

var fromToIsNullNonNilTestCase = coretestcases.CaseV1{
	Title:         "IsNull - non-nil returns false",
	ExpectedInput: "false",
}

// ==========================================================================
// IsFromEmpty
// ==========================================================================

var fromToIsFromEmptyEmptyTestCase = coretestcases.CaseV1{
	Title:         "IsFromEmpty - empty From returns true",
	ExpectedInput: "true",
}

var fromToIsFromEmptyNilTestCase = coretestcases.CaseV1{
	Title:         "IsFromEmpty - nil receiver returns true",
	ExpectedInput: "true",
}

// ==========================================================================
// IsToEmpty
// ==========================================================================

var fromToIsToEmptyEmptyTestCase = coretestcases.CaseV1{
	Title:         "IsToEmpty - empty To returns true",
	ExpectedInput: "true",
}

var fromToIsToEmptyNonEmptyTestCase = coretestcases.CaseV1{
	Title:         "IsToEmpty - non-empty returns false",
	ExpectedInput: "false",
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
	Title:         "SetFromName - updates From",
	ExpectedInput: "new",
}

var fromToSetFromNameNilTestCase = coretestcases.CaseV1{
	Title:         "SetFromName - nil receiver no panic",
	ExpectedInput: "true",
}

// ==========================================================================
// SetToName
// ==========================================================================

var fromToSetToNameUpdatesTestCase = coretestcases.CaseV1{
	Title:         "SetToName - updates To",
	ExpectedInput: "new",
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
	Title:         "SourceDestination - nil returns nil",
	ExpectedInput: "true",
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
	Title:         "Rename - nil returns nil",
	ExpectedInput: "true",
}
