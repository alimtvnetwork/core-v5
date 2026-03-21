package codestacktests

import (
	"testing"

	"github.com/alimtvnetwork/core/codestack"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── newTraceCollection.Default ──
// Covers newTraceCollection.go L13-15

func Test_Cov10_NewTraceCollection_Default(t *testing.T) {
	// Arrange & Act
	tc := codestack.New.StackTrace.Default()

	// Assert
	actual := args.Map{"isNil": tc == nil}
	expected := args.Map{"isNil": false}
	expected.ShouldBeEqual(t, 0, "Default returns non-nil TraceCollection", actual)
}

// ── newTraceCollection.Using with nil traces ──
// Covers newTraceCollection.go L21-23

func Test_Cov10_NewTraceCollection_Using_NilTraces(t *testing.T) {
	// Arrange & Act
	tc := codestack.New.StackTrace.Using(false, nil...)

	// Assert
	actual := args.Map{"isNil": tc == nil, "length": tc.Length()}
	expected := args.Map{"isNil": false, "length": 0}
	expected.ShouldBeEqual(t, 0, "Using nil traces returns empty collection", actual)
}

// ── newTraceCollection.Using with isClone=true ──
// Covers newTraceCollection.go L33-35

func Test_Cov10_NewTraceCollection_Using_Clone(t *testing.T) {
	// Arrange
	trace := codestack.New.Create(1)

	// Act
	tc := codestack.New.StackTrace.Using(true, trace)

	// Assert
	actual := args.Map{"isNil": tc == nil, "hasItems": tc.Length() > 0}
	expected := args.Map{"isNil": false, "hasItems": true}
	expected.ShouldBeEqual(t, 0, "Using isClone=true clones traces", actual)
}

// ── AddsUsingSkip with isSkipInvalid=true, isBreakOnceInvalid=false ──
// Covers TraceCollection.go L75 (continue path)

func Test_Cov10_AddsUsingSkip_SkipContinue(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.Cap(10)

	// Act — large skip index produces invalid traces; isBreakOnceInvalid=false means continue
	tc.AddsUsingSkip(true, false, 900, 5)

	// Assert — all traces should be skipped (invalid at high skip index)
	actual := args.Map{"length": tc.Length()}
	expected := args.Map{"length": 0}
	expected.ShouldBeEqual(t, 0, "AddsUsingSkip continues past invalid -- no break", actual)
}

// ── AddsUsingSkipUsingFilter: no-break path ──
// Covers TraceCollection.go L119-120 (continue) and L141 (end-of-loop return)

func Test_Cov10_AddsUsingSkipUsingFilter_SkipContinueAndEndReturn(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.Cap(10)
	takeAll := func(trace *codestack.Trace) (bool, bool) {
		return true, false
	}

	// Act — large skip index means invalid traces, isBreakOnceInvalid=false → continue
	tc.AddsUsingSkipUsingFilter(true, false, 900, 5, takeAll)

	// Assert
	actual := args.Map{"length": tc.Length()}
	expected := args.Map{"length": 0}
	expected.ShouldBeEqual(t, 0, "AddsUsingSkipUsingFilter continues and returns -- no items", actual)
}

// ── FilterWithLimit: natural loop exhaustion ──
// Covers TraceCollection.go L529 (end-of-loop return)

func Test_Cov10_FilterWithLimit_NaturalExhaustion(t *testing.T) {
	// Arrange — create collection from current stack
	tc := codestack.New.StackTrace.Collect(1, 3)
	takeAll := func(trace *codestack.Trace) (bool, bool) {
		return true, false // take all, never break
	}

	// Act — limit larger than collection, so loop finishes naturally
	result := tc.FilterWithLimit(100, takeAll)

	// Assert
	actual := args.Map{"hasItems": len(result) > 0}
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "FilterWithLimit returns all items -- natural loop end", actual)
}

// ── GetSinglePageCollection: negative page panic ──
// Covers TraceCollection.go L419-426

func Test_Cov10_GetSinglePageCollection_NegativePagePanic(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.Collect(1, 20)

	// Act
	didPanic := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		tc.GetSinglePageCollection(5, -1)
	}()

	// Assert
	actual := args.Map{"didPanic": didPanic}
	expected := args.Map{"didPanic": true}
	expected.ShouldBeEqual(t, 0, "GetSinglePageCollection negative page panics", actual)
}
