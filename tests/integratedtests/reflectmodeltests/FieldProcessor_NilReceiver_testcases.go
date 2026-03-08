package reflectmodeltests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/reflectcore/reflectmodel"
	"reflect"
)

// =============================================================================
// FieldProcessor nil receiver test cases (migrated from inline t.Error tests)
// =============================================================================

var fieldProcessorNilReceiverTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "IsFieldType on nil returns false",
		Func:  (*reflectmodel.FieldProcessor).IsFieldType,
		Args:  []any{reflect.TypeOf("")},
		Expected: args.Map{
			"value":    "false",
			"panicked": false,
		},
	},
	{
		Title: "IsFieldKind on nil returns false",
		Func:  (*reflectmodel.FieldProcessor).IsFieldKind,
		Args:  []any{reflect.String},
		Expected: args.Map{
			"value":    "false",
			"panicked": false,
		},
	},
}
