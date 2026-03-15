package corevalidatortests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/corevalidator"
)

// ── LineNumber ──

func Test_Cov3_LineNumber_String(t *testing.T) {
	ln := corevalidator.LineNumber{
		Number: 5,
		Line:   "hello world",
	}
	actual := args.Map{
		"number":   ln.Number,
		"line":     ln.Line,
		"notEmpty": ln.String() != "",
	}
	expected := args.Map{"number": 5, "line": "hello world", "notEmpty": true}
	expected.ShouldBeEqual(t, 0, "LineNumber String", actual)
}

// ── Condition ──

func Test_Cov3_Condition_Methods(t *testing.T) {
	c := corevalidator.Condition{
		IsValid: true,
		Message: "all good",
	}
	actual := args.Map{
		"isValid":  c.IsValid,
		"message":  c.Message,
		"hasError": c.HasError(),
	}
	expected := args.Map{
		"isValid": true, "message": "all good", "hasError": false,
	}
	expected.ShouldBeEqual(t, 0, "Condition methods", actual)
}

func Test_Cov3_Condition_HasError_True(t *testing.T) {
	c := corevalidator.Condition{IsValid: false, Message: "error"}
	actual := args.Map{"hasError": c.HasError()}
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "Condition HasError true", actual)
}

// ── Parameter ──

func Test_Cov3_Parameter_Methods(t *testing.T) {
	p := corevalidator.Parameter{
		Name:  "testParam",
		Value: "testValue",
	}
	actual := args.Map{
		"name":     p.Name,
		"value":    p.Value,
		"notEmpty": p.String() != "",
	}
	expected := args.Map{
		"name": "testParam", "value": "testValue", "notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Parameter methods", actual)
}

// ── RangesSegment ──

func Test_Cov3_RangesSegment(t *testing.T) {
	rs := corevalidator.RangesSegment{
		ActualStartIndex:    0,
		ActualEndIndex:      10,
		ExpectedStartIndex:  0,
		ExpectedEndIndex:    10,
	}
	actual := args.Map{
		"start":    rs.ActualStartIndex,
		"end":      rs.ActualEndIndex,
		"expStart": rs.ExpectedStartIndex,
		"expEnd":   rs.ExpectedEndIndex,
	}
	expected := args.Map{
		"start": 0, "end": 10, "expStart": 0, "expEnd": 10,
	}
	expected.ShouldBeEqual(t, 0, "RangesSegment struct", actual)
}

// ── BaseValidatorCoreCondition ──

func Test_Cov3_BaseValidatorCoreCondition_EmptyConditions(t *testing.T) {
	bv := corevalidator.BaseValidatorCoreCondition{
		Conditions: []corevalidator.Condition{},
	}
	actual := args.Map{
		"len":    len(bv.Conditions),
		"hasErr": bv.HasAnyError(),
	}
	expected := args.Map{"len": 0, "hasErr": false}
	expected.ShouldBeEqual(t, 0, "BaseValidatorCoreCondition empty", actual)
}

func Test_Cov3_BaseValidatorCoreCondition_WithError(t *testing.T) {
	bv := corevalidator.BaseValidatorCoreCondition{
		Conditions: []corevalidator.Condition{
			{IsValid: true, Message: "ok"},
			{IsValid: false, Message: "fail"},
		},
	}
	actual := args.Map{"hasErr": bv.HasAnyError()}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BaseValidatorCoreCondition with error", actual)
}
