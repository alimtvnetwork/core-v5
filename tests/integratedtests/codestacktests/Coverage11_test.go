package codestacktests

import (
	"testing"

	"github.com/alimtvnetwork/core/codestack"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── newTraceCollection.Default ──
// Covers newTraceCollection.go L13-15

func Test_Cov11_TraceCollection_Default(t *testing.T) {
	// Arrange & Act
	tc := codestack.New.TraceCollection.Default()

	// Assert
	actual := args.Map{"notNil": tc != nil, "length": tc.Length()}
	expected := args.Map{"notNil": true, "length": 0}
	expected.ShouldBeEqual(t, 0, "TraceCollection Default creates empty collection", actual)
}

// ── newTraceCollection.Using with nil traces ──
// Covers newTraceCollection.go L21-23

func Test_Cov11_TraceCollection_Using_NilTraces(t *testing.T) {
	// Arrange & Act
	tc := codestack.New.TraceCollection.Using(false, nil...)

	// Assert
	actual := args.Map{"notNil": tc != nil, "length": tc.Length()}
	expected := args.Map{"notNil": true, "length": 0}
	expected.ShouldBeEqual(t, 0, "TraceCollection Using nil traces returns empty", actual)
}

// ── newTraceCollection.Using with clone=true ──
// Covers newTraceCollection.go L33-35

func Test_Cov11_TraceCollection_Using_Clone(t *testing.T) {
	// Arrange — create some traces
	traces := []codestack.Trace{
		{},
		{},
	}

	// Act
	tc := codestack.New.TraceCollection.Using(true, traces...)

	// Assert
	actual := args.Map{"notNil": tc != nil, "length": tc.Length()}
	expected := args.Map{"notNil": true, "length": 2}
	expected.ShouldBeEqual(t, 0, "TraceCollection Using clone copies traces", actual)
}

// ── GetSinglePageCollection: pageIndex=0 triggers negative skipItems panic ──
// Covers TraceCollection.go L419-426

func Test_Cov11_GetSinglePageCollection_ZeroPagePanic(t *testing.T) {
	// Arrange — create collection with items
	tc := codestack.New.TraceCollection.Default()
	traces := []codestack.Trace{{}, {}, {}}
	tc.Adds(traces...)

	// Act
	didPanic := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		tc.GetSinglePageCollection(2, 0) // pageIndex=0 → skipItems = 2*(0-1) = -2 → panic
	}()

	// Assert
	actual := args.Map{"didPanic": didPanic}
	expected := args.Map{"didPanic": true}
	expected.ShouldBeEqual(t, 0, "GetSinglePageCollection panics on zero pageIndex", actual)
}
