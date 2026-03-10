package coredynamictests

import (
	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/coretests/results"
)

// =============================================================================
// MapAnyItems nil receiver test cases
// (migrated from standalone CaseV1 variables in MapAnyItemsEdge_testcases.go)
//
// Note: Some MapAnyItems nil tests require ArrangeInput for the right-side
// argument (e.g., IsEqualRaw), so they remain in CaseV1. Only pure
// nil-receiver zero-arg methods are migrated here.
// =============================================================================

var mapAnyItemsNilSafeTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "IsEqualRaw on nil with non-nil map returns false",
		Func: func(m *coredynamic.MapAnyItems) bool {
			return m.IsEqualRaw(map[string]any{"k": "v"})
		},
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "IsEqualRaw on nil with nil map returns true",
		Func: func(m *coredynamic.MapAnyItems) bool {
			return m.IsEqualRaw(nil)
		},
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "Length on nil returns 0",
		Func:  (*coredynamic.MapAnyItems).Length,
		Expected: results.ResultAny{
			Value:    "0",
			Panicked: false,
		},
	},
	{
		Title: "IsEmpty on nil returns true",
		Func:  (*coredynamic.MapAnyItems).IsEmpty,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "HasAnyItem on nil returns false",
		Func:  (*coredynamic.MapAnyItems).HasAnyItem,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "HasKey on nil returns false",
		Func: func(m *coredynamic.MapAnyItems) bool {
			return m.HasKey("anything")
		},
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "Clear on nil does not panic",
		Func:  (*coredynamic.MapAnyItems).Clear,
		Expected: results.ResultAny{
			Panicked: false,
		},
		CompareFields: []string{"panicked", "returnCount"},
	},
	{
		Title: "DeepClear on nil does not panic",
		Func:  (*coredynamic.MapAnyItems).DeepClear,
		Expected: results.ResultAny{
			Panicked: false,
		},
		CompareFields: []string{"panicked", "returnCount"},
	},
	{
		Title: "Dispose on nil does not panic",
		Func:  (*coredynamic.MapAnyItems).Dispose,
		Expected: results.ResultAny{
			Panicked: false,
		},
		CompareFields: []string{"panicked", "returnCount"},
	},
	{
		Title: "ClonePtr on nil returns nil and error",
		Func: func(m *coredynamic.MapAnyItems) bool {
			ptr, err := m.ClonePtr()
			return ptr == nil && err != nil
		},
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
}
