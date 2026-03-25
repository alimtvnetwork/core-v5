package reflectmodeltests

import (
	"reflect"

	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/coretests/results"
	"github.com/alimtvnetwork/core/reflectcore/reflectmodel"
)

// =============================================================================
// FieldProcessor nil receiver test cases (migrated from inline t.Error tests)
// =============================================================================

var fieldProcessorNilReceiverTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "IsFieldType on nil returns false",
		Func:  (*reflectmodel.FieldProcessor).IsFieldType,
		Args:  []any{reflect.TypeOf("")},
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "IsFieldKind on nil returns false",
		Func:  (*reflectmodel.FieldProcessor).IsFieldKind,
		Args:  []any{reflect.String},
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
}
