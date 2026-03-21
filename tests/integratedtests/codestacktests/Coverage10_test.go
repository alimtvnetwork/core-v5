package codestacktests

import (
	"testing"

	"github.com/alimtvnetwork/core/codestack"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── newStacksCreator.All with isBreakOnceInvalid=false ──
// Covers TraceCollection.go L75 (continue path on invalid trace)

func Test_Cov10_AddsUsingSkip_SkipContinue(t *testing.T) {
	// Arrange & Act — large skip index means invalid traces
	// isSkipInvalid=true, isBreakOnceInvalid=false → continue past invalid
	tc := codestack.New.StackTrace.All(true, false, 900, 5)

	// Assert — all traces should be skipped
	actual := args.Map{"length": tc.Length()}
	expected := args.Map{"length": 0}
	expected.ShouldBeEqual(t, 0, "All continues past invalid -- no break", actual)
}

// ── newStacksCreator.Default ──
// Covers newTraceCollection.go L13-15 (Default) indirectly

func Test_Cov10_StackTrace_Default(t *testing.T) {
	// Arrange & Act
	tc := codestack.New.StackTrace.Default(1, 3)

	// Assert
	actual := args.Map{"hasItems": tc.HasAnyItem()}
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "StackTrace Default returns items", actual)
}

// ── FilterWithLimit: natural loop exhaustion ──
// Covers TraceCollection.go L529 (end-of-loop return)

func Test_Cov10_FilterWithLimit_NaturalExhaustion(t *testing.T) {
	// Arrange
	tcVal := codestack.New.StackTrace.Default(1, 3)
	tc := &tcVal
	takeAll := func(trace *codestack.Trace) (bool, bool) {
		return true, false
	}

	// Act — limit larger than collection
	result := tc.FilterWithLimit(100, takeAll)

	// Assert
	actual := args.Map{"hasItems": len(result) > 0}
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "FilterWithLimit returns all -- natural loop end", actual)
}

// ── GetSinglePageCollection: negative page panic ──
// Covers TraceCollection.go L419-426

func Test_Cov10_GetSinglePageCollection_NegativePagePanic(t *testing.T) {
	// Arrange — need length >= eachPageSize for the method to not short-circuit
	tcVal := codestack.New.StackTrace.Default(1, 30)
	tc := &tcVal

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
