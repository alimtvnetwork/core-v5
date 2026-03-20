package codestacktests

import (
	"testing"

	"github.com/alimtvnetwork/core/codestack"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── newTraceCollection ──

func Test_Cov9_NewTraceCollection_Default(t *testing.T) {
	// Arrange & Act
	tc := codestack.New.StackTrace.Default()

	// Assert
	actual := args.Map{"isNil": tc == nil}
	expected := args.Map{"isNil": false}
	expected.ShouldBeEqual(t, 0, "newTraceCollection Default", actual)
}

func Test_Cov9_NewTraceCollection_Empty(t *testing.T) {
	// Arrange & Act
	tc := codestack.New.StackTrace.Empty()

	// Assert
	actual := args.Map{"isEmpty": tc.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "newTraceCollection Empty", actual)
}

func Test_Cov9_NewTraceCollection_Using_Nil(t *testing.T) {
	// Arrange & Act
	tc := codestack.New.StackTrace.Using(false)

	// Assert
	actual := args.Map{"isEmpty": tc.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "newTraceCollection Using nil", actual)
}

func Test_Cov9_NewTraceCollection_Using_Clone(t *testing.T) {
	// Arrange
	trace := codestack.New.Create(codestack.Skip1)

	// Act
	tc := codestack.New.StackTrace.Using(true, trace)

	// Assert
	actual := args.Map{"length": tc.Length()}
	expected := args.Map{"length": 1}
	expected.ShouldBeEqual(t, 0, "newTraceCollection Using clone", actual)
}

// ── TraceCollection.IsEqualItems — nil paths ──

func Test_Cov9_IsEqualItems_BothNil(t *testing.T) {
	// Arrange
	var tc *codestack.TraceCollection

	// Act
	result := tc.IsEqualItems()

	// Assert
	actual := args.Map{"isEqual": result}
	expected := args.Map{"isEqual": true}
	expected.ShouldBeEqual(t, 0, "IsEqualItems both nil", actual)
}

func Test_Cov9_IsEqualItems_ReceiverNilItemsNot(t *testing.T) {
	// Arrange
	var tc *codestack.TraceCollection
	trace := codestack.New.Create(codestack.Skip1)

	// Act
	result := tc.IsEqualItems(trace)

	// Assert
	actual := args.Map{"isEqual": result}
	expected := args.Map{"isEqual": false}
	expected.ShouldBeEqual(t, 0, "IsEqualItems receiver nil items not", actual)
}

// ── TraceCollection.FilterWithLimit — isBreak branch ──

func Test_Cov9_FilterWithLimit_Break(t *testing.T) {
	// Arrange
	trace1 := codestack.New.Create(codestack.Skip1)
	trace2 := codestack.New.Create(codestack.Skip2)
	tc := codestack.New.StackTrace.Using(false, trace1, trace2)
	breakFilter := func(tr *codestack.Trace) (bool, bool) {
		return true, true // take first, then break
	}

	// Act
	result := tc.FilterWithLimit(10, breakFilter)

	// Assert
	actual := args.Map{"length": len(result)}
	expected := args.Map{"length": 1}
	expected.ShouldBeEqual(t, 0, "FilterWithLimit break on first", actual)
}

// ── TraceCollection.AddsUsingSkipUsingFilter — isBreak branch ──

func Test_Cov9_AddsUsingSkipUsingFilter_Break(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.Cap(10)
	breakFilter := func(tr *codestack.Trace) (bool, bool) {
		return true, true // take and break immediately
	}

	// Act
	result := tc.AddsUsingSkipUsingFilter(true, true, 0, 5, breakFilter)

	// Assert
	actual := args.Map{"hasItems": result.HasAnyItem()}
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "AddsUsingSkipUsingFilter with break", actual)
}

// ── ParseInjectUsingJson — error paths ──

func Test_Cov9_FileWithLine_ParseInjectUsingJsonMust_Panic(t *testing.T) {
	// Arrange
	fwl := &codestack.FileWithLine{}
	badResult := corejson.NewPtr("not-a-FileWithLine")
	badResult.Bytes = []byte("{invalid")

	// Act
	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		fwl.ParseInjectUsingJsonMust(badResult)
	}()

	// Assert
	actual := args.Map{"didPanic": didPanic}
	expected := args.Map{"didPanic": true}
	expected.ShouldBeEqual(t, 0, "FileWithLine ParseInjectUsingJsonMust panic", actual)
}

func Test_Cov9_Trace_ParseInjectUsingJson_Error(t *testing.T) {
	// Arrange
	tr := &codestack.Trace{}
	badResult := corejson.NewPtr("bad")
	badResult.Bytes = []byte("{invalid")

	// Act
	_, err := tr.ParseInjectUsingJson(badResult)

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "Trace ParseInjectUsingJson error", actual)
}

func Test_Cov9_Trace_ParseInjectUsingJsonMust_Panic(t *testing.T) {
	// Arrange
	tr := &codestack.Trace{}
	badResult := corejson.NewPtr("bad")
	badResult.Bytes = []byte("{invalid")

	// Act
	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		tr.ParseInjectUsingJsonMust(badResult)
	}()

	// Assert
	actual := args.Map{"didPanic": didPanic}
	expected := args.Map{"didPanic": true}
	expected.ShouldBeEqual(t, 0, "Trace ParseInjectUsingJsonMust panic", actual)
}

func Test_Cov9_TraceCollection_ParseInjectUsingJson_Error(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.Empty()
	badResult := corejson.NewPtr("bad")
	badResult.Bytes = []byte("{invalid")

	// Act
	result, err := tc.ParseInjectUsingJson(badResult)

	// Assert
	actual := args.Map{"hasError": err != nil, "isEmpty": result.IsEmpty()}
	expected := args.Map{"hasError": true, "isEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection ParseInjectUsingJson error", actual)
}

func Test_Cov9_TraceCollection_ParseInjectUsingJsonMust_Panic(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.Empty()
	badResult := corejson.NewPtr("bad")
	badResult.Bytes = []byte("{invalid")

	// Act
	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		tc.ParseInjectUsingJsonMust(badResult)
	}()

	// Assert
	actual := args.Map{"didPanic": didPanic}
	expected := args.Map{"didPanic": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection ParseInjectUsingJsonMust panic", actual)
}

// ── TraceCollection.PaginateAt — negative page panic ──

func Test_Cov9_TraceCollection_PaginateAt_NegativePanic(t *testing.T) {
	// Arrange
	trace := codestack.New.Create(codestack.Skip1)
	tc := codestack.New.StackTrace.Using(false, trace)

	// Act
	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		tc.PaginateAt(0, 5) // pageIndex=0 causes skipItems = 5*(0-1) = -5
	}()

	// Assert
	actual := args.Map{"didPanic": didPanic}
	expected := args.Map{"didPanic": true}
	expected.ShouldBeEqual(t, 0, "PaginateAt negative page panic", actual)
}
