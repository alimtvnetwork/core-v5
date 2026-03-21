package corevalidatortests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/corevalidator"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
)

// ══════════════════════════════════════════════════════════════════════════════
// Condition
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov8_Condition_IsSplitByWhitespace_AllFalse(t *testing.T) {
	c := &corevalidator.Condition{}
	actual := args.Map{"split": c.IsSplitByWhitespace()}
	expected := args.Map{"split": false}
	expected.ShouldBeEqual(t, 0, "Condition.IsSplitByWhitespace returns non-empty -- all false", actual)
}

func Test_Cov8_Condition_IsSplitByWhitespace_UniqueWord(t *testing.T) {
	c := &corevalidator.Condition{IsUniqueWordOnly: true}
	actual := args.Map{"split": c.IsSplitByWhitespace()}
	expected := args.Map{"split": true}
	expected.ShouldBeEqual(t, 0, "Condition.IsSplitByWhitespace returns correct value -- unique word", actual)
}

func Test_Cov8_Condition_IsSplitByWhitespace_NonEmpty(t *testing.T) {
	c := &corevalidator.Condition{IsNonEmptyWhitespace: true}
	actual := args.Map{"split": c.IsSplitByWhitespace()}
	expected := args.Map{"split": true}
	expected.ShouldBeEqual(t, 0, "Condition.IsSplitByWhitespace returns empty -- non-empty", actual)
}

func Test_Cov8_Condition_IsSplitByWhitespace_Sort(t *testing.T) {
	c := &corevalidator.Condition{IsSortStringsBySpace: true}
	actual := args.Map{"split": c.IsSplitByWhitespace()}
	expected := args.Map{"split": true}
	expected.ShouldBeEqual(t, 0, "Condition.IsSplitByWhitespace returns correct value -- sort", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Parameter
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov8_Parameter_IsIgnoreCase_Sensitive(t *testing.T) {
	p := corevalidator.Parameter{IsCaseSensitive: true}
	actual := args.Map{"ignoreCase": p.IsIgnoreCase()}
	expected := args.Map{"ignoreCase": false}
	expected.ShouldBeEqual(t, 0, "Parameter.IsIgnoreCase returns correct value -- sensitive", actual)
}

func Test_Cov8_Parameter_IsIgnoreCase_Insensitive(t *testing.T) {
	p := corevalidator.Parameter{IsCaseSensitive: false}
	actual := args.Map{"ignoreCase": p.IsIgnoreCase()}
	expected := args.Map{"ignoreCase": true}
	expected.ShouldBeEqual(t, 0, "Parameter.IsIgnoreCase returns correct value -- insensitive", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// LineNumber
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov8_LineNumber_HasLineNumber_Valid(t *testing.T) {
	ln := &corevalidator.LineNumber{LineNumber: 5}
	actual := args.Map{"has": ln.HasLineNumber()}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "LineNumber.HasLineNumber returns non-empty -- valid", actual)
}

func Test_Cov8_LineNumber_HasLineNumber_Invalid(t *testing.T) {
	ln := &corevalidator.LineNumber{LineNumber: -1}
	actual := args.Map{"has": ln.HasLineNumber()}
	expected := args.Map{"has": false}
	expected.ShouldBeEqual(t, 0, "LineNumber.HasLineNumber returns error -- invalid", actual)
}

func Test_Cov8_LineNumber_IsMatch_BothInvalid(t *testing.T) {
	ln := &corevalidator.LineNumber{LineNumber: -1}
	actual := args.Map{"match": ln.IsMatch(-1)}
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "LineNumber.IsMatch returns error -- both invalid", actual)
}

func Test_Cov8_LineNumber_IsMatch_Same(t *testing.T) {
	ln := &corevalidator.LineNumber{LineNumber: 5}
	actual := args.Map{"match": ln.IsMatch(5)}
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "LineNumber.IsMatch returns correct value -- same", actual)
}

func Test_Cov8_LineNumber_IsMatch_Different(t *testing.T) {
	ln := &corevalidator.LineNumber{LineNumber: 5}
	actual := args.Map{"match": ln.IsMatch(3)}
	expected := args.Map{"match": false}
	expected.ShouldBeEqual(t, 0, "LineNumber.IsMatch returns correct value -- different", actual)
}

func Test_Cov8_LineNumber_IsMatch_InputInvalid(t *testing.T) {
	ln := &corevalidator.LineNumber{LineNumber: 5}
	actual := args.Map{"match": ln.IsMatch(-1)}
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "LineNumber.IsMatch returns error -- input invalid", actual)
}

func Test_Cov8_LineNumber_VerifyError_Match(t *testing.T) {
	ln := &corevalidator.LineNumber{LineNumber: 5}
	err := ln.VerifyError(5)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "LineNumber.VerifyError returns error -- match", actual)
}

func Test_Cov8_LineNumber_VerifyError_Mismatch(t *testing.T) {
	ln := &corevalidator.LineNumber{LineNumber: 5}
	err := ln.VerifyError(3)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "LineNumber.VerifyError returns error -- mismatch", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// TextValidator
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov8_TextValidator_ToString_SingleLine(t *testing.T) {
	tv := &corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal}
	result := tv.ToString(true)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.ToString returns non-empty -- single", actual)
}

func Test_Cov8_TextValidator_ToString_MultiLine(t *testing.T) {
	tv := &corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal}
	result := tv.ToString(false)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.ToString returns non-empty -- multi", actual)
}

func Test_Cov8_TextValidator_String(t *testing.T) {
	tv := &corevalidator.TextValidator{Search: "test", SearchAs: stringcompareas.Equal}
	result := tv.String()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.String returns non-empty -- with args", actual)
}

func Test_Cov8_TextValidator_SearchTextFinalized(t *testing.T) {
	tv := &corevalidator.TextValidator{Search: " hello ", SearchAs: stringcompareas.Equal, Condition: corevalidator.Condition{IsTrimCompare: true}}
	result := tv.SearchTextFinalized()
	actual := args.Map{"val": result}
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "TextValidator.SearchTextFinalized returns non-empty -- trimmed", actual)
}

func Test_Cov8_TextValidator_SearchTextFinalized_Cached(t *testing.T) {
	tv := &corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal}
	r1 := tv.SearchTextFinalizedPtr()
	r2 := tv.SearchTextFinalizedPtr()
	actual := args.Map{"same": r1 == r2}
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.SearchTextFinalized returns non-empty -- cached", actual)
}

func Test_Cov8_TextValidator_GetCompiledTerm_NoConditions(t *testing.T) {
	tv := &corevalidator.TextValidator{}
	result := tv.GetCompiledTermBasedOnConditions("hello world", false)
	actual := args.Map{"val": result}
	expected := args.Map{"val": "hello world"}
	expected.ShouldBeEqual(t, 0, "GetCompiledTerm returns empty -- no conditions", actual)
}

func Test_Cov8_TextValidator_GetCompiledTerm_Trim(t *testing.T) {
	tv := &corevalidator.TextValidator{Condition: corevalidator.Condition{IsTrimCompare: true}}
	result := tv.GetCompiledTermBasedOnConditions("  hello  ", false)
	actual := args.Map{"val": result}
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "GetCompiledTerm returns correct value -- trim", actual)
}

func Test_Cov8_TextValidator_GetCompiledTerm_SplitByWhitespace(t *testing.T) {
	tv := &corevalidator.TextValidator{Condition: corevalidator.Condition{IsNonEmptyWhitespace: true, IsSortStringsBySpace: true}}
	result := tv.GetCompiledTermBasedOnConditions("b a", false)
	actual := args.Map{"val": result}
	expected := args.Map{"val": "a b"}
	expected.ShouldBeEqual(t, 0, "GetCompiledTerm returns correct value -- split whitespace", actual)
}

func Test_Cov8_TextValidator_IsMatch_Equal(t *testing.T) {
	tv := &corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal}
	actual := args.Map{"match": tv.IsMatch("hello", true), "noMatch": tv.IsMatch("world", true)}
	expected := args.Map{"match": true, "noMatch": false}
	expected.ShouldBeEqual(t, 0, "TextValidator.IsMatch returns non-empty -- equal", actual)
}

func Test_Cov8_TextValidator_IsMatch_Contains(t *testing.T) {
	tv := &corevalidator.TextValidator{Search: "ell", SearchAs: stringcompareas.Contains}
	actual := args.Map{"match": tv.IsMatch("hello", true)}
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.IsMatch returns non-empty -- contains", actual)
}

func Test_Cov8_TextValidator_IsMatchMany_NilReceiver(t *testing.T) {
	var tv *corevalidator.TextValidator
	actual := args.Map{"match": tv.IsMatchMany(true, true, "a", "b")}
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.IsMatchMany returns nil -- nil", actual)
}

func Test_Cov8_TextValidator_IsMatchMany_EmptySkip(t *testing.T) {
	tv := &corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal}
	actual := args.Map{"match": tv.IsMatchMany(true, true)}
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.IsMatchMany returns empty -- empty skip", actual)
}

func Test_Cov8_TextValidator_IsMatchMany_EmptyNoSkip(t *testing.T) {
	tv := &corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal}
	actual := args.Map{"match": tv.IsMatchMany(false, true)}
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.IsMatchMany returns empty -- empty no skip", actual)
}

func Test_Cov8_TextValidator_IsMatchMany_AllMatch(t *testing.T) {
	tv := &corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal}
	actual := args.Map{"match": tv.IsMatchMany(false, true, "a", "a")}
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.IsMatchMany returns non-empty -- all match", actual)
}

func Test_Cov8_TextValidator_IsMatchMany_OneFails(t *testing.T) {
	tv := &corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal}
	actual := args.Map{"match": tv.IsMatchMany(false, true, "a", "b")}
	expected := args.Map{"match": false}
	expected.ShouldBeEqual(t, 0, "TextValidator.IsMatchMany returns non-empty -- one fails", actual)
}

func Test_Cov8_TextValidator_MethodName(t *testing.T) {
	tv := &corevalidator.TextValidator{SearchAs: stringcompareas.StartsWith}
	actual := args.Map{"name": tv.MethodName()}
	expected := args.Map{"name": "StartsWith"}
	expected.ShouldBeEqual(t, 0, "TextValidator.MethodName returns non-empty -- with args", actual)
}

func Test_Cov8_TextValidator_VerifyDetailError_Nil(t *testing.T) {
	var tv *corevalidator.TextValidator
	err := tv.VerifyDetailError(&corevalidator.Parameter{}, "hello")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.VerifyDetailError returns nil -- nil", actual)
}

func Test_Cov8_TextValidator_VerifyDetailError_Match(t *testing.T) {
	tv := &corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal}
	err := tv.VerifyDetailError(&corevalidator.Parameter{IsCaseSensitive: true}, "hello")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.VerifyDetailError returns error -- match", actual)
}

func Test_Cov8_TextValidator_VerifyDetailError_Mismatch(t *testing.T) {
	tv := &corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal}
	err := tv.VerifyDetailError(&corevalidator.Parameter{IsCaseSensitive: true}, "world")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.VerifyDetailError returns error -- mismatch", actual)
}

func Test_Cov8_TextValidator_VerifySimpleError_Nil(t *testing.T) {
	var tv *corevalidator.TextValidator
	err := tv.VerifySimpleError(0, &corevalidator.Parameter{}, "hello")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.VerifySimpleError returns nil -- nil", actual)
}

func Test_Cov8_TextValidator_VerifySimpleError_Match(t *testing.T) {
	tv := &corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal}
	err := tv.VerifySimpleError(0, &corevalidator.Parameter{IsCaseSensitive: true}, "hello")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.VerifySimpleError returns error -- match", actual)
}

func Test_Cov8_TextValidator_VerifySimpleError_Mismatch(t *testing.T) {
	tv := &corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal}
	err := tv.VerifySimpleError(0, &corevalidator.Parameter{IsCaseSensitive: true}, "world")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.VerifySimpleError returns error -- mismatch", actual)
}

func Test_Cov8_TextValidator_VerifyMany_ContinueOnError(t *testing.T) {
	tv := &corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal}
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	err := tv.VerifyMany(true, params, "a", "b")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.VerifyMany returns non-empty -- continue", actual)
}

func Test_Cov8_TextValidator_VerifyMany_StopOnFirst(t *testing.T) {
	tv := &corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal}
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	err := tv.VerifyMany(false, params, "a", "b")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.VerifyMany returns non-empty -- stop first", actual)
}

func Test_Cov8_TextValidator_VerifyFirstError_Nil(t *testing.T) {
	var tv *corevalidator.TextValidator
	err := tv.VerifyFirstError(&corevalidator.Parameter{}, "a")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.VerifyFirstError returns nil -- nil", actual)
}

func Test_Cov8_TextValidator_VerifyFirstError_EmptySkip(t *testing.T) {
	tv := &corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal}
	err := tv.VerifyFirstError(&corevalidator.Parameter{IsSkipCompareOnActualEmpty: true})
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.VerifyFirstError returns empty -- empty skip", actual)
}

func Test_Cov8_TextValidator_AllVerifyError_Nil(t *testing.T) {
	var tv *corevalidator.TextValidator
	err := tv.AllVerifyError(&corevalidator.Parameter{}, "a")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.AllVerifyError returns nil -- nil", actual)
}

func Test_Cov8_TextValidator_AllVerifyError_EmptySkip(t *testing.T) {
	tv := &corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal}
	err := tv.AllVerifyError(&corevalidator.Parameter{IsSkipCompareOnActualEmpty: true})
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.AllVerifyError returns empty -- empty skip", actual)
}

func Test_Cov8_TextValidator_AllVerifyError_WithErrors(t *testing.T) {
	tv := &corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal}
	err := tv.AllVerifyError(&corevalidator.Parameter{IsCaseSensitive: true}, "a", "b", "c")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.AllVerifyError returns error -- with errors", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// TextValidators
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov8_TextValidators_NilLength(t *testing.T) {
	var tvs *corevalidator.TextValidators
	actual := args.Map{"len": tvs.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "TextValidators returns nil -- nil length", actual)
}

func Test_Cov8_TextValidators_Count(t *testing.T) {
	tvs := corevalidator.NewTextValidators(5)
	tvs.Add(corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal})
	actual := args.Map{"count": tvs.Count()}
	expected := args.Map{"count": 0} // Count = LastIndex = Length - 1
	expected.ShouldBeEqual(t, 0, "TextValidators.Count returns non-empty -- with args", actual)
}

func Test_Cov8_TextValidators_Adds(t *testing.T) {
	tvs := corevalidator.NewTextValidators(5)
	tvs.Adds(
		corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal},
		corevalidator.TextValidator{Search: "b", SearchAs: stringcompareas.Equal},
	)
	actual := args.Map{"len": tvs.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "TextValidators.Adds returns non-empty -- with args", actual)
}

func Test_Cov8_TextValidators_Adds_Empty(t *testing.T) {
	tvs := corevalidator.NewTextValidators(5)
	tvs.Adds()
	actual := args.Map{"len": tvs.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "TextValidators.Adds returns empty -- empty", actual)
}

func Test_Cov8_TextValidators_AddSimple(t *testing.T) {
	tvs := corevalidator.NewTextValidators(5)
	tvs.AddSimple("hello", stringcompareas.Equal)
	actual := args.Map{"len": tvs.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TextValidators.AddSimple returns non-empty -- with args", actual)
}

func Test_Cov8_TextValidators_AddSimpleAllTrue(t *testing.T) {
	tvs := corevalidator.NewTextValidators(5)
	tvs.AddSimpleAllTrue("hello", stringcompareas.Equal)
	actual := args.Map{"len": tvs.Length(), "hasAny": tvs.HasAnyItem()}
	expected := args.Map{"len": 1, "hasAny": true}
	expected.ShouldBeEqual(t, 0, "TextValidators.AddSimpleAllTrue returns non-empty -- with args", actual)
}

func Test_Cov8_TextValidators_HasIndex(t *testing.T) {
	tvs := corevalidator.NewTextValidators(5)
	tvs.Add(corevalidator.TextValidator{Search: "a"})
	actual := args.Map{"has0": tvs.HasIndex(0), "has1": tvs.HasIndex(1)}
	expected := args.Map{"has0": true, "has1": false}
	expected.ShouldBeEqual(t, 0, "TextValidators.HasIndex returns non-empty -- with args", actual)
}

func Test_Cov8_TextValidators_String(t *testing.T) {
	tvs := corevalidator.NewTextValidators(5)
	tvs.Add(corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal})
	result := tvs.String()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TextValidators.String returns non-empty -- with args", actual)
}

func Test_Cov8_TextValidators_IsMatch_Empty(t *testing.T) {
	tvs := corevalidator.NewTextValidators(5)
	actual := args.Map{"match": tvs.IsMatch("hello", true)}
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "TextValidators.IsMatch returns empty -- empty", actual)
}

func Test_Cov8_TextValidators_IsMatch_AllPass(t *testing.T) {
	tvs := corevalidator.NewTextValidators(5)
	tvs.Add(corevalidator.TextValidator{Search: "hel", SearchAs: stringcompareas.StartsWith})
	tvs.Add(corevalidator.TextValidator{Search: "llo", SearchAs: stringcompareas.EndsWith})
	actual := args.Map{"match": tvs.IsMatch("hello", true)}
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "TextValidators.IsMatch returns non-empty -- all pass", actual)
}

func Test_Cov8_TextValidators_IsMatch_OneFails(t *testing.T) {
	tvs := corevalidator.NewTextValidators(5)
	tvs.Add(corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal})
	tvs.Add(corevalidator.TextValidator{Search: "world", SearchAs: stringcompareas.Equal})
	actual := args.Map{"match": tvs.IsMatch("hello", true)}
	expected := args.Map{"match": false}
	expected.ShouldBeEqual(t, 0, "TextValidators.IsMatch returns non-empty -- one fails", actual)
}

func Test_Cov8_TextValidators_IsMatchMany_Empty(t *testing.T) {
	tvs := corevalidator.NewTextValidators(5)
	actual := args.Map{"match": tvs.IsMatchMany(true, true, "a")}
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "TextValidators.IsMatchMany returns empty -- empty", actual)
}

func Test_Cov8_TextValidators_VerifyFirstError_Empty(t *testing.T) {
	tvs := corevalidator.NewTextValidators(5)
	err := tvs.VerifyFirstError(0, "hello", true)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidators.VerifyFirstError returns empty -- empty", actual)
}

func Test_Cov8_TextValidators_VerifyFirstError_Match(t *testing.T) {
	tvs := corevalidator.NewTextValidators(5)
	tvs.Add(corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal})
	err := tvs.VerifyFirstError(0, "hello", true)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidators.VerifyFirstError returns error -- match", actual)
}

func Test_Cov8_TextValidators_VerifyFirstError_Mismatch(t *testing.T) {
	tvs := corevalidator.NewTextValidators(5)
	tvs.Add(corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal})
	err := tvs.VerifyFirstError(0, "world", true)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidators.VerifyFirstError returns error -- mismatch", actual)
}

func Test_Cov8_TextValidators_VerifyErrorMany_Nil(t *testing.T) {
	var tvs *corevalidator.TextValidators
	err := tvs.VerifyErrorMany(true, &corevalidator.Parameter{}, "a")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidators.VerifyErrorMany returns nil -- nil", actual)
}

func Test_Cov8_TextValidators_VerifyErrorMany_Continue(t *testing.T) {
	tvs := corevalidator.NewTextValidators(5)
	tvs.Add(corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal})
	err := tvs.VerifyErrorMany(true, &corevalidator.Parameter{IsCaseSensitive: true}, "a")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidators.VerifyErrorMany returns error -- continue", actual)
}

func Test_Cov8_TextValidators_VerifyErrorMany_StopFirst(t *testing.T) {
	tvs := corevalidator.NewTextValidators(5)
	tvs.Add(corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal})
	err := tvs.VerifyErrorMany(false, &corevalidator.Parameter{IsCaseSensitive: true}, "a")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidators.VerifyErrorMany returns error -- stop first", actual)
}

func Test_Cov8_TextValidators_VerifyFirstErrorMany_Empty(t *testing.T) {
	tvs := corevalidator.NewTextValidators(5)
	err := tvs.VerifyFirstErrorMany(&corevalidator.Parameter{}, "a")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidators.VerifyFirstErrorMany returns empty -- empty", actual)
}

func Test_Cov8_TextValidators_AllVerifyErrorMany_Empty(t *testing.T) {
	tvs := corevalidator.NewTextValidators(5)
	err := tvs.AllVerifyErrorMany(&corevalidator.Parameter{}, "a")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidators.AllVerifyErrorMany returns empty -- empty", actual)
}

func Test_Cov8_TextValidators_AllVerifyError_Empty(t *testing.T) {
	tvs := corevalidator.NewTextValidators(5)
	err := tvs.AllVerifyError(0, "hello", true)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidators.AllVerifyError returns empty -- empty", actual)
}

func Test_Cov8_TextValidators_AllVerifyError_WithErrors(t *testing.T) {
	tvs := corevalidator.NewTextValidators(5)
	tvs.Add(corevalidator.TextValidator{Search: "x", SearchAs: stringcompareas.Equal})
	err := tvs.AllVerifyError(0, "y", true)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidators.AllVerifyError returns error -- with errors", actual)
}

func Test_Cov8_TextValidators_Dispose(t *testing.T) {
	tvs := corevalidator.NewTextValidators(5)
	tvs.Add(corevalidator.TextValidator{Search: "a"})
	tvs.Dispose()
	actual := args.Map{"nil": tvs.Items == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "TextValidators.Dispose returns non-empty -- with args", actual)
}

func Test_Cov8_TextValidators_Dispose_Nil(t *testing.T) {
	var tvs *corevalidator.TextValidators
	tvs.Dispose() // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "TextValidators.Dispose returns nil -- nil", actual)
}

func Test_Cov8_TextValidators_AsBasicSliceContractsBinder(t *testing.T) {
	tvs := corevalidator.NewTextValidators(5)
	binder := tvs.AsBasicSliceContractsBinder()
	actual := args.Map{"notNil": binder != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TextValidators.AsBasicSliceContractsBinder returns non-empty -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// LineValidator
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov8_LineValidator_IsMatch_LineAndText(t *testing.T) {
	lv := &corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: 5},
		TextValidator: corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal},
	}
	actual := args.Map{
		"matchBoth":  lv.IsMatch(5, "hello", true),
		"wrongLine":  lv.IsMatch(3, "hello", true),
		"wrongText":  lv.IsMatch(5, "world", true),
		"anyLine":    lv.IsMatch(-1, "hello", true),
	}
	expected := args.Map{
		"matchBoth": true, "wrongLine": false, "wrongText": false, "anyLine": true,
	}
	expected.ShouldBeEqual(t, 0, "LineValidator.IsMatch returns non-empty -- with args", actual)
}

func Test_Cov8_LineValidator_IsMatchMany_Nil(t *testing.T) {
	var lv *corevalidator.LineValidator
	actual := args.Map{"match": lv.IsMatchMany(true, true)}
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "LineValidator.IsMatchMany returns nil -- nil", actual)
}

func Test_Cov8_LineValidator_IsMatchMany_EmptySkip(t *testing.T) {
	lv := &corevalidator.LineValidator{
		TextValidator: corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal},
	}
	actual := args.Map{"match": lv.IsMatchMany(true, true)}
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "LineValidator.IsMatchMany returns empty -- empty skip", actual)
}

func Test_Cov8_LineValidator_IsMatchMany_AllMatch(t *testing.T) {
	lv := &corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal},
	}
	items := []corestr.TextWithLineNumber{
		{Text: "a", LineNumber: 0},
		{Text: "a", LineNumber: 1},
	}
	actual := args.Map{"match": lv.IsMatchMany(false, true, items...)}
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "LineValidator.IsMatchMany returns non-empty -- all match", actual)
}

func Test_Cov8_LineValidator_IsMatchMany_OneFails(t *testing.T) {
	lv := &corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal},
	}
	items := []corestr.TextWithLineNumber{
		{Text: "a", LineNumber: 0},
		{Text: "b", LineNumber: 1},
	}
	actual := args.Map{"match": lv.IsMatchMany(false, true, items...)}
	expected := args.Map{"match": false}
	expected.ShouldBeEqual(t, 0, "LineValidator.IsMatchMany returns non-empty -- one fails", actual)
}

func Test_Cov8_LineValidator_VerifyError_Match(t *testing.T) {
	lv := &corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	err := lv.VerifyError(params, 0, "hello")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "LineValidator.VerifyError returns error -- match", actual)
}

func Test_Cov8_LineValidator_VerifyError_LineMismatch(t *testing.T) {
	lv := &corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: 5},
		TextValidator: corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	err := lv.VerifyError(params, 3, "hello")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "LineValidator.VerifyError returns error -- line mismatch", actual)
}

func Test_Cov8_LineValidator_VerifyError_TextMismatch(t *testing.T) {
	lv := &corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	err := lv.VerifyError(params, 0, "world")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "LineValidator.VerifyError returns error -- text mismatch", actual)
}

func Test_Cov8_LineValidator_VerifyMany_Continue(t *testing.T) {
	lv := &corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	items := []corestr.TextWithLineNumber{{Text: "a", LineNumber: 0}}
	err := lv.VerifyMany(true, params, items...)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "LineValidator.VerifyMany returns non-empty -- continue", actual)
}

func Test_Cov8_LineValidator_VerifyMany_StopFirst(t *testing.T) {
	lv := &corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	items := []corestr.TextWithLineNumber{{Text: "a", LineNumber: 0}}
	err := lv.VerifyMany(false, params, items...)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "LineValidator.VerifyMany returns non-empty -- stop first", actual)
}

func Test_Cov8_LineValidator_VerifyFirstError_Nil(t *testing.T) {
	var lv *corevalidator.LineValidator
	err := lv.VerifyFirstError(&corevalidator.Parameter{})
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "LineValidator.VerifyFirstError returns nil -- nil", actual)
}

func Test_Cov8_LineValidator_VerifyFirstError_EmptySkip(t *testing.T) {
	lv := &corevalidator.LineValidator{
		TextValidator: corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal},
	}
	err := lv.VerifyFirstError(&corevalidator.Parameter{IsSkipCompareOnActualEmpty: true})
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "LineValidator.VerifyFirstError returns empty -- empty skip", actual)
}

func Test_Cov8_LineValidator_AllVerifyError_Nil(t *testing.T) {
	var lv *corevalidator.LineValidator
	err := lv.AllVerifyError(&corevalidator.Parameter{})
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "LineValidator.AllVerifyError returns nil -- nil", actual)
}

func Test_Cov8_LineValidator_AllVerifyError_EmptySkip(t *testing.T) {
	lv := &corevalidator.LineValidator{
		TextValidator: corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal},
	}
	err := lv.AllVerifyError(&corevalidator.Parameter{IsSkipCompareOnActualEmpty: true})
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "LineValidator.AllVerifyError returns empty -- empty skip", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// LinesValidators
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov8_LinesValidators_NilLength(t *testing.T) {
	var lv *corevalidator.LinesValidators
	actual := args.Map{"len": lv.Length(), "empty": lv.IsEmpty()}
	expected := args.Map{"len": 0, "empty": true}
	expected.ShouldBeEqual(t, 0, "LinesValidators returns nil -- nil", actual)
}

func Test_Cov8_LinesValidators_Basic(t *testing.T) {
	lv := corevalidator.NewLinesValidators(5)
	actual := args.Map{"len": lv.Length(), "empty": lv.IsEmpty(), "count": lv.Count(), "hasAny": lv.HasAnyItem()}
	expected := args.Map{"len": 0, "empty": true, "count": 0, "hasAny": false}
	expected.ShouldBeEqual(t, 0, "LinesValidators returns non-empty -- basic", actual)
}

func Test_Cov8_LinesValidators_Add(t *testing.T) {
	lv := corevalidator.NewLinesValidators(5)
	lv.Add(corevalidator.LineValidator{
		TextValidator: corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal},
	})
	actual := args.Map{"len": lv.Length(), "hasAny": lv.HasAnyItem(), "lastIndex": lv.LastIndex()}
	expected := args.Map{"len": 1, "hasAny": true, "lastIndex": 0}
	expected.ShouldBeEqual(t, 0, "LinesValidators.Add returns non-empty -- with args", actual)
}

func Test_Cov8_LinesValidators_AddPtr_Nil(t *testing.T) {
	lv := corevalidator.NewLinesValidators(5)
	lv.AddPtr(nil)
	actual := args.Map{"len": lv.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "LinesValidators.AddPtr returns nil -- nil", actual)
}

func Test_Cov8_LinesValidators_AddPtr_Valid(t *testing.T) {
	lv := corevalidator.NewLinesValidators(5)
	v := &corevalidator.LineValidator{
		TextValidator: corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal},
	}
	lv.AddPtr(v)
	actual := args.Map{"len": lv.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "LinesValidators.AddPtr returns non-empty -- valid", actual)
}

func Test_Cov8_LinesValidators_Adds(t *testing.T) {
	lv := corevalidator.NewLinesValidators(5)
	lv.Adds(
		corevalidator.LineValidator{TextValidator: corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal}},
		corevalidator.LineValidator{TextValidator: corevalidator.TextValidator{Search: "b", SearchAs: stringcompareas.Equal}},
	)
	actual := args.Map{"len": lv.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "LinesValidators.Adds returns non-empty -- with args", actual)
}

func Test_Cov8_LinesValidators_HasIndex(t *testing.T) {
	lv := corevalidator.NewLinesValidators(5)
	lv.Add(corevalidator.LineValidator{TextValidator: corevalidator.TextValidator{Search: "a"}})
	actual := args.Map{"has0": lv.HasIndex(0), "has1": lv.HasIndex(1)}
	expected := args.Map{"has0": true, "has1": false}
	expected.ShouldBeEqual(t, 0, "LinesValidators.HasIndex returns non-empty -- with args", actual)
}

func Test_Cov8_LinesValidators_String(t *testing.T) {
	lv := corevalidator.NewLinesValidators(5)
	lv.Add(corevalidator.LineValidator{TextValidator: corevalidator.TextValidator{Search: "a"}})
	actual := args.Map{"notEmpty": lv.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "LinesValidators.String returns non-empty -- with args", actual)
}

func Test_Cov8_LinesValidators_AsBasicSliceContractsBinder(t *testing.T) {
	lv := corevalidator.NewLinesValidators(5)
	binder := lv.AsBasicSliceContractsBinder()
	actual := args.Map{"notNil": binder != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "LinesValidators.AsBasicSliceContractsBinder returns non-empty -- with args", actual)
}

func Test_Cov8_LinesValidators_IsMatchText_Empty(t *testing.T) {
	lv := corevalidator.NewLinesValidators(5)
	actual := args.Map{"match": lv.IsMatchText("anything", true)}
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "LinesValidators.IsMatchText returns empty -- empty", actual)
}

func Test_Cov8_LinesValidators_IsMatchText_Match(t *testing.T) {
	lv := corevalidator.NewLinesValidators(5)
	lv.Add(corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal},
	})
	actual := args.Map{"match": lv.IsMatchText("hello", true)}
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "LinesValidators.IsMatchText returns non-empty -- match", actual)
}

func Test_Cov8_LinesValidators_IsMatchText_NoMatch(t *testing.T) {
	lv := corevalidator.NewLinesValidators(5)
	lv.Add(corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal},
	})
	actual := args.Map{"match": lv.IsMatchText("world", true)}
	expected := args.Map{"match": false}
	expected.ShouldBeEqual(t, 0, "LinesValidators.IsMatchText returns empty -- no match", actual)
}

func Test_Cov8_LinesValidators_IsMatch_Empty(t *testing.T) {
	lv := corevalidator.NewLinesValidators(5)
	actual := args.Map{"match": lv.IsMatch(true, true)}
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "LinesValidators.IsMatch returns empty -- empty", actual)
}

func Test_Cov8_LinesValidators_IsMatch_NoContentsSkip(t *testing.T) {
	lv := corevalidator.NewLinesValidators(5)
	lv.Add(corevalidator.LineValidator{TextValidator: corevalidator.TextValidator{Search: "a"}})
	actual := args.Map{"match": lv.IsMatch(true, true)}
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "LinesValidators.IsMatch returns empty -- no contents skip", actual)
}

func Test_Cov8_LinesValidators_IsMatch_NoContentsNoSkip(t *testing.T) {
	lv := corevalidator.NewLinesValidators(5)
	lv.Add(corevalidator.LineValidator{TextValidator: corevalidator.TextValidator{Search: "a"}})
	actual := args.Map{"match": lv.IsMatch(false, true)}
	expected := args.Map{"match": false}
	expected.ShouldBeEqual(t, 0, "LinesValidators.IsMatch returns empty -- no contents no skip", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// BaseLinesValidators
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov8_BaseLinesValidators_Nil(t *testing.T) {
	var blv *corevalidator.BaseLinesValidators
	actual := args.Map{"len": blv.LinesValidatorsLength(), "empty": blv.IsEmptyLinesValidators(), "has": blv.HasLinesValidators()}
	expected := args.Map{"len": 0, "empty": true, "has": false}
	expected.ShouldBeEqual(t, 0, "BaseLinesValidators returns nil -- nil", actual)
}

func Test_Cov8_BaseLinesValidators_Empty(t *testing.T) {
	blv := &corevalidator.BaseLinesValidators{}
	actual := args.Map{"len": blv.LinesValidatorsLength(), "empty": blv.IsEmptyLinesValidators()}
	expected := args.Map{"len": 0, "empty": true}
	expected.ShouldBeEqual(t, 0, "BaseLinesValidators returns empty -- empty", actual)
}

func Test_Cov8_BaseLinesValidators_WithItems(t *testing.T) {
	blv := &corevalidator.BaseLinesValidators{
		LinesValidators: []corevalidator.LineValidator{
			{TextValidator: corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal}},
		},
	}
	actual := args.Map{"len": blv.LinesValidatorsLength(), "has": blv.HasLinesValidators()}
	expected := args.Map{"len": 1, "has": true}
	expected.ShouldBeEqual(t, 0, "BaseLinesValidators returns non-empty -- with items", actual)
}

func Test_Cov8_BaseLinesValidators_ToLinesValidators_Empty(t *testing.T) {
	blv := &corevalidator.BaseLinesValidators{}
	lv := blv.ToLinesValidators()
	actual := args.Map{"notNil": lv != nil, "empty": lv.IsEmpty()}
	expected := args.Map{"notNil": true, "empty": true}
	expected.ShouldBeEqual(t, 0, "BaseLinesValidators.ToLinesValidators returns empty -- empty", actual)
}

func Test_Cov8_BaseLinesValidators_ToLinesValidators_WithItems(t *testing.T) {
	blv := &corevalidator.BaseLinesValidators{
		LinesValidators: []corevalidator.LineValidator{
			{TextValidator: corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal}},
		},
	}
	lv := blv.ToLinesValidators()
	actual := args.Map{"notNil": lv != nil, "len": lv.Length()}
	expected := args.Map{"notNil": true, "len": 1}
	expected.ShouldBeEqual(t, 0, "BaseLinesValidators.ToLinesValidators returns non-empty -- with items", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// BaseValidatorCoreCondition
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov8_BaseValidatorCoreCondition_WithCondition(t *testing.T) {
	cond := &corevalidator.Condition{IsTrimCompare: true}
	bvc := &corevalidator.BaseValidatorCoreCondition{ValidatorCoreCondition: cond}
	result := bvc.ValidatorCoreConditionDefault()
	actual := args.Map{"trim": result.IsTrimCompare}
	expected := args.Map{"trim": true}
	expected.ShouldBeEqual(t, 0, "BaseValidatorCoreCondition returns non-empty -- with condition", actual)
}

func Test_Cov8_BaseValidatorCoreCondition_NilCondition(t *testing.T) {
	bvc := &corevalidator.BaseValidatorCoreCondition{}
	result := bvc.ValidatorCoreConditionDefault()
	actual := args.Map{"trim": result.IsTrimCompare}
	expected := args.Map{"trim": false}
	expected.ShouldBeEqual(t, 0, "BaseValidatorCoreCondition returns nil -- nil condition", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SliceValidator — Messages
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov8_SliceValidator_ActualInputWithExpectingMessage(t *testing.T) {
	sv := &corevalidator.SliceValidator{
		ActualLines:   []string{"a"},
		ExpectedLines: []string{"b"},
		CompareAs:     stringcompareas.Equal,
	}
	result := sv.ActualInputWithExpectingMessage(0, "test")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SliceValidator.ActualInputWithExpectingMessage returns non-empty -- with args", actual)
}

func Test_Cov8_SliceValidator_ActualInputMessage(t *testing.T) {
	sv := &corevalidator.SliceValidator{ActualLines: []string{"hello"}}
	result := sv.ActualInputMessage(0, "test")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SliceValidator.ActualInputMessage returns non-empty -- with args", actual)
}

func Test_Cov8_SliceValidator_UserExpectingMessage(t *testing.T) {
	sv := &corevalidator.SliceValidator{ExpectedLines: []string{"hello"}}
	result := sv.UserExpectingMessage(0, "test")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SliceValidator.UserExpectingMessage returns non-empty -- with args", actual)
}

func Test_Cov8_SliceValidator_UserInputsMergeWithError_NoAttach(t *testing.T) {
	sv := &corevalidator.SliceValidator{ActualLines: []string{"a"}, ExpectedLines: []string{"b"}}
	params := &corevalidator.Parameter{IsAttachUserInputs: false}
	testErr := errors.New("test error")
	result := sv.UserInputsMergeWithError(params, testErr)
	actual := args.Map{"sameErr": result.Error() == "test error"}
	expected := args.Map{"sameErr": true}
	expected.ShouldBeEqual(t, 0, "UserInputsMergeWithError returns empty -- no attach", actual)
}

func Test_Cov8_SliceValidator_UserInputsMergeWithError_Attach_NilErr(t *testing.T) {
	sv := &corevalidator.SliceValidator{ActualLines: []string{"a"}, ExpectedLines: []string{"b"}}
	params := &corevalidator.Parameter{IsAttachUserInputs: true}
	result := sv.UserInputsMergeWithError(params, nil)
	actual := args.Map{"hasErr": result != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "UserInputsMergeWithError returns nil -- attach nil err", actual)
}

func Test_Cov8_SliceValidator_UserInputsMergeWithError_Attach_WithErr(t *testing.T) {
	sv := &corevalidator.SliceValidator{ActualLines: []string{"a"}, ExpectedLines: []string{"b"}}
	params := &corevalidator.Parameter{IsAttachUserInputs: true}
	testErr := errors.New("base")
	result := sv.UserInputsMergeWithError(params, testErr)
	actual := args.Map{"hasErr": result != nil, "containsBase": len(result.Error()) > 4}
	expected := args.Map{"hasErr": true, "containsBase": true}
	expected.ShouldBeEqual(t, 0, "UserInputsMergeWithError returns error -- attach with err", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SliceValidator — Constructors
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov8_NewSliceValidatorUsingErr(t *testing.T) {
	testErr := errors.New("line1\nline2")
	sv := corevalidator.NewSliceValidatorUsingErr(testErr, "line1\nline2", true, false, false, stringcompareas.Equal)
	actual := args.Map{"notNil": sv != nil, "actualLen": sv.ActualLinesLength(), "expectedLen": sv.ExpectingLinesLength()}
	expected := args.Map{"notNil": true, "actualLen": 2, "expectedLen": 2}
	expected.ShouldBeEqual(t, 0, "NewSliceValidatorUsingErr returns error -- with args", actual)
}

func Test_Cov8_NewSliceValidatorUsingAny(t *testing.T) {
	sv := corevalidator.NewSliceValidatorUsingAny("hello\nworld", "hello\nworld", false, false, false, stringcompareas.Equal)
	actual := args.Map{"notNil": sv != nil, "actualLen": sv.ActualLinesLength()}
	expected := args.Map{"notNil": true, "actualLen": 2}
	expected.ShouldBeEqual(t, 0, "NewSliceValidatorUsingAny returns non-empty -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SliceValidator — Verify
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov8_SliceValidator_VerifyFirstError_Nil(t *testing.T) {
	var sv *corevalidator.SliceValidator
	err := sv.VerifyFirstError(&corevalidator.Parameter{})
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SliceValidator.VerifyFirstError returns nil -- nil", actual)
}

func Test_Cov8_SliceValidator_AllVerifyError_Nil(t *testing.T) {
	var sv *corevalidator.SliceValidator
	err := sv.AllVerifyError(&corevalidator.Parameter{})
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SliceValidator.AllVerifyError returns nil -- nil", actual)
}

func Test_Cov8_SliceValidator_AllVerifyError_Match(t *testing.T) {
	sv := &corevalidator.SliceValidator{
		ActualLines:   []string{"hello", "world"},
		ExpectedLines: []string{"hello", "world"},
		CompareAs:     stringcompareas.Equal,
	}
	err := sv.AllVerifyError(&corevalidator.Parameter{IsCaseSensitive: true})
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SliceValidator.AllVerifyError returns error -- match", actual)
}

func Test_Cov8_SliceValidator_AllVerifyError_Mismatch(t *testing.T) {
	sv := &corevalidator.SliceValidator{
		ActualLines:   []string{"hello"},
		ExpectedLines: []string{"world"},
		CompareAs:     stringcompareas.Equal,
	}
	err := sv.AllVerifyError(&corevalidator.Parameter{IsCaseSensitive: true})
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SliceValidator.AllVerifyError returns error -- mismatch", actual)
}

func Test_Cov8_SliceValidator_AllVerifyErrorExceptLast_Nil(t *testing.T) {
	var sv *corevalidator.SliceValidator
	err := sv.AllVerifyErrorExceptLast(&corevalidator.Parameter{})
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SliceValidator.AllVerifyErrorExceptLast returns nil -- nil", actual)
}

func Test_Cov8_SliceValidator_AllVerifyErrorUptoLength_EmptyIgnore(t *testing.T) {
	sv := &corevalidator.SliceValidator{
		ActualLines:   []string{},
		ExpectedLines: []string{"a"},
		CompareAs:     stringcompareas.Equal,
	}
	err := sv.AllVerifyErrorUptoLength(false, &corevalidator.Parameter{IsSkipCompareOnActualEmpty: true}, 1)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyErrorUptoLength returns empty -- empty ignore", actual)
}

func Test_Cov8_SliceValidator_AllVerifyErrorUptoLength_BothNil(t *testing.T) {
	sv := &corevalidator.SliceValidator{CompareAs: stringcompareas.Equal}
	err := sv.AllVerifyErrorUptoLength(false, &corevalidator.Parameter{}, 0)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyErrorUptoLength returns nil -- both nil", actual)
}

func Test_Cov8_SliceValidator_AllVerifyErrorUptoLength_OneNil(t *testing.T) {
	sv := &corevalidator.SliceValidator{
		ActualLines: []string{"a"},
		CompareAs:   stringcompareas.Equal,
	}
	err := sv.AllVerifyErrorUptoLength(false, &corevalidator.Parameter{}, 1)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyErrorUptoLength returns nil -- one nil", actual)
}

func Test_Cov8_SliceValidator_AllVerifyErrorUptoLength_LengthMismatch(t *testing.T) {
	sv := &corevalidator.SliceValidator{
		ActualLines:   []string{"a"},
		ExpectedLines: []string{"a", "b"},
		CompareAs:     stringcompareas.Equal,
	}
	err := sv.AllVerifyErrorUptoLength(false, &corevalidator.Parameter{}, 2)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyErrorUptoLength returns error -- length mismatch", actual)
}

func Test_Cov8_SliceValidator_AllVerifyErrorTestCase_Nil(t *testing.T) {
	var sv *corevalidator.SliceValidator
	err := sv.AllVerifyErrorTestCase(0, "test", true)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyErrorTestCase returns nil -- nil", actual)
}

func Test_Cov8_SliceValidator_AllVerifyErrorQuick_Nil(t *testing.T) {
	var sv *corevalidator.SliceValidator
	err := sv.AllVerifyErrorQuick(0, "test", "a")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyErrorQuick returns nil -- nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SimpleSliceValidator
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov8_SimpleSliceValidator_SetActual(t *testing.T) {
	ssv := &corevalidator.SimpleSliceValidator{
		Expected:  corestr.New.SimpleSlice.Direct(false, []string{"a"}),
		CompareAs: stringcompareas.Equal,
	}
	result := ssv.SetActual([]string{"a"})
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SimpleSliceValidator.SetActual returns non-empty -- with args", actual)
}

func Test_Cov8_SimpleSliceValidator_SliceValidator(t *testing.T) {
	ssv := &corevalidator.SimpleSliceValidator{
		Expected:  corestr.New.SimpleSlice.Direct(false, []string{"a"}),
		CompareAs: stringcompareas.Equal,
	}
	ssv.SetActual([]string{"a"})
	sv := ssv.SliceValidator()
	actual := args.Map{"notNil": sv != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SimpleSliceValidator.SliceValidator returns non-empty -- with args", actual)
}

func Test_Cov8_SimpleSliceValidator_VerifyAll(t *testing.T) {
	ssv := &corevalidator.SimpleSliceValidator{
		Expected:  corestr.New.SimpleSlice.Direct(false, []string{"a"}),
		CompareAs: stringcompareas.Equal,
	}
	ssv.SetActual([]string{"a"})
	err := ssv.VerifyAll([]string{"a"}, &corevalidator.Parameter{IsCaseSensitive: true})
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleSliceValidator.VerifyAll returns non-empty -- with args", actual)
}

func Test_Cov8_SimpleSliceValidator_VerifyFirst(t *testing.T) {
	ssv := &corevalidator.SimpleSliceValidator{
		Expected:  corestr.New.SimpleSlice.Direct(false, []string{"a"}),
		CompareAs: stringcompareas.Equal,
	}
	ssv.SetActual([]string{"a"})
	err := ssv.VerifyFirst([]string{"a"}, &corevalidator.Parameter{IsCaseSensitive: true})
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleSliceValidator.VerifyFirst returns non-empty -- with args", actual)
}

func Test_Cov8_SimpleSliceValidator_VerifyUpto(t *testing.T) {
	ssv := &corevalidator.SimpleSliceValidator{
		Expected:  corestr.New.SimpleSlice.Direct(false, []string{"a", "b"}),
		CompareAs: stringcompareas.Equal,
	}
	ssv.SetActual([]string{"a", "b"})
	err := ssv.VerifyUpto([]string{"a", "b"}, &corevalidator.Parameter{IsCaseSensitive: true}, 1)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleSliceValidator.VerifyUpto returns non-empty -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// RangesSegment / RangeSegmentsValidator
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov8_RangesSegment_Fields(t *testing.T) {
	rs := corevalidator.RangesSegment{
		ExpectedLines: []string{"a"},
		CompareAs:     stringcompareas.Equal,
	}
	actual := args.Map{"len": len(rs.ExpectedLines)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "RangesSegment returns correct value -- fields", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Vars (predefined conditions)
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov8_DefaultDisabledCoreCondition(t *testing.T) {
	c := corevalidator.DefaultDisabledCoreCondition
	actual := args.Map{"trim": c.IsTrimCompare, "unique": c.IsUniqueWordOnly, "nonEmpty": c.IsNonEmptyWhitespace, "sort": c.IsSortStringsBySpace}
	expected := args.Map{"trim": false, "unique": false, "nonEmpty": false, "sort": false}
	expected.ShouldBeEqual(t, 0, "DefaultDisabledCoreCondition returns correct value -- with args", actual)
}

func Test_Cov8_DefaultTrimCoreCondition(t *testing.T) {
	c := corevalidator.DefaultTrimCoreCondition
	actual := args.Map{"trim": c.IsTrimCompare}
	expected := args.Map{"trim": true}
	expected.ShouldBeEqual(t, 0, "DefaultTrimCoreCondition returns correct value -- with args", actual)
}

func Test_Cov8_DefaultSortTrimCoreCondition(t *testing.T) {
	c := corevalidator.DefaultSortTrimCoreCondition
	actual := args.Map{"trim": c.IsTrimCompare, "nonEmpty": c.IsNonEmptyWhitespace, "sort": c.IsSortStringsBySpace}
	expected := args.Map{"trim": true, "nonEmpty": true, "sort": true}
	expected.ShouldBeEqual(t, 0, "DefaultSortTrimCoreCondition returns correct value -- with args", actual)
}

func Test_Cov8_DefaultUniqueWordsCoreCondition(t *testing.T) {
	c := corevalidator.DefaultUniqueWordsCoreCondition
	actual := args.Map{"trim": c.IsTrimCompare, "unique": c.IsUniqueWordOnly, "nonEmpty": c.IsNonEmptyWhitespace, "sort": c.IsSortStringsBySpace}
	expected := args.Map{"trim": true, "unique": true, "nonEmpty": true, "sort": true}
	expected.ShouldBeEqual(t, 0, "DefaultUniqueWordsCoreCondition returns correct value -- with args", actual)
}

func Test_Cov8_EmptyValidator(t *testing.T) {
	v := corevalidator.EmptyValidator
	actual := args.Map{"search": v.Search, "method": v.SearchAs.Name(), "trim": v.IsTrimCompare}
	expected := args.Map{"search": "", "method": "Equal", "trim": true}
	expected.ShouldBeEqual(t, 0, "EmptyValidator returns empty -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SliceValidators
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov8_SliceValidators_NilLength(t *testing.T) {
	var svs *corevalidator.SliceValidators
	actual := args.Map{"len": svs.Length(), "empty": svs.IsEmpty()}
	expected := args.Map{"len": 0, "empty": true}
	expected.ShouldBeEqual(t, 0, "SliceValidators returns nil -- nil", actual)
}

func Test_Cov8_SliceValidators_IsMatch_Empty(t *testing.T) {
	svs := &corevalidator.SliceValidators{}
	actual := args.Map{"match": svs.IsMatch(true)}
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "SliceValidators.IsMatch returns empty -- empty", actual)
}

func Test_Cov8_SliceValidators_IsMatch_AllPass(t *testing.T) {
	svs := &corevalidator.SliceValidators{
		Validators: []corevalidator.SliceValidator{
			{ActualLines: []string{"a"}, ExpectedLines: []string{"a"}, CompareAs: stringcompareas.Equal},
		},
	}
	actual := args.Map{"match": svs.IsMatch(true)}
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "SliceValidators.IsMatch returns non-empty -- all pass", actual)
}

func Test_Cov8_SliceValidators_IsMatch_Fail(t *testing.T) {
	svs := &corevalidator.SliceValidators{
		Validators: []corevalidator.SliceValidator{
			{ActualLines: []string{"a"}, ExpectedLines: []string{"b"}, CompareAs: stringcompareas.Equal},
		},
	}
	actual := args.Map{"match": svs.IsMatch(true)}
	expected := args.Map{"match": false}
	expected.ShouldBeEqual(t, 0, "SliceValidators.IsMatch returns non-empty -- fail", actual)
}

func Test_Cov8_SliceValidators_IsValid(t *testing.T) {
	svs := &corevalidator.SliceValidators{}
	actual := args.Map{"valid": svs.IsValid(true)}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "SliceValidators.IsValid returns non-empty -- with args", actual)
}

func Test_Cov8_SliceValidators_SetActualOnAll_Empty(t *testing.T) {
	svs := &corevalidator.SliceValidators{}
	svs.SetActualOnAll("a", "b") // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "SliceValidators.SetActualOnAll returns empty -- empty", actual)
}

func Test_Cov8_SliceValidators_VerifyAll_Empty(t *testing.T) {
	svs := &corevalidator.SliceValidators{}
	err := svs.VerifyAll("header", &corevalidator.Parameter{}, false)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SliceValidators.VerifyAll returns empty -- empty", actual)
}

func Test_Cov8_SliceValidators_VerifyAllError_Empty(t *testing.T) {
	svs := &corevalidator.SliceValidators{}
	err := svs.VerifyAllError(&corevalidator.Parameter{Header: "test"})
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SliceValidators.VerifyAllError returns empty -- empty", actual)
}

func Test_Cov8_SliceValidators_VerifyAllErrorUsingActual_Empty(t *testing.T) {
	svs := &corevalidator.SliceValidators{}
	err := svs.VerifyAllErrorUsingActual(&corevalidator.Parameter{Header: "test"}, "a")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SliceValidators.VerifyAllErrorUsingActual returns empty -- empty", actual)
}

func Test_Cov8_SliceValidators_VerifyFirst_Empty(t *testing.T) {
	svs := &corevalidator.SliceValidators{}
	err := svs.VerifyFirst(&corevalidator.Parameter{}, false)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SliceValidators.VerifyFirst returns empty -- empty", actual)
}

func Test_Cov8_SliceValidators_VerifyUpto_Empty(t *testing.T) {
	svs := &corevalidator.SliceValidators{}
	err := svs.VerifyUpto(false, false, 1, &corevalidator.Parameter{})
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SliceValidators.VerifyUpto returns empty -- empty", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// HeaderSliceValidators
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov8_HeaderSliceValidators_NilLength(t *testing.T) {
	var hsvs corevalidator.HeaderSliceValidators
	actual := args.Map{"len": hsvs.Length(), "empty": hsvs.IsEmpty()}
	expected := args.Map{"len": 0, "empty": true}
	expected.ShouldBeEqual(t, 0, "HeaderSliceValidators returns nil -- nil", actual)
}

func Test_Cov8_HeaderSliceValidators_IsMatch_Empty(t *testing.T) {
	hsvs := corevalidator.HeaderSliceValidators{}
	actual := args.Map{"match": hsvs.IsMatch(true)}
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "HeaderSliceValidators.IsMatch returns empty -- empty", actual)
}

func Test_Cov8_HeaderSliceValidators_IsValid(t *testing.T) {
	hsvs := corevalidator.HeaderSliceValidators{}
	actual := args.Map{"valid": hsvs.IsValid(true)}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "HeaderSliceValidators.IsValid returns non-empty -- with args", actual)
}

func Test_Cov8_HeaderSliceValidators_SetActualOnAll_Empty(t *testing.T) {
	hsvs := corevalidator.HeaderSliceValidators{}
	hsvs.SetActualOnAll("a") // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "HeaderSliceValidators.SetActualOnAll returns empty -- empty", actual)
}

func Test_Cov8_HeaderSliceValidators_VerifyAll_Empty(t *testing.T) {
	hsvs := corevalidator.HeaderSliceValidators{}
	err := hsvs.VerifyAll("header", &corevalidator.Parameter{}, false)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "HeaderSliceValidators.VerifyAll returns empty -- empty", actual)
}

func Test_Cov8_HeaderSliceValidators_VerifyAllError_Empty(t *testing.T) {
	hsvs := corevalidator.HeaderSliceValidators{}
	err := hsvs.VerifyAllError(&corevalidator.Parameter{Header: "test"})
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "HeaderSliceValidators.VerifyAllError returns empty -- empty", actual)
}

func Test_Cov8_HeaderSliceValidators_VerifyAllErrorUsingActual_Empty(t *testing.T) {
	hsvs := corevalidator.HeaderSliceValidators{}
	err := hsvs.VerifyAllErrorUsingActual(&corevalidator.Parameter{Header: "test"}, "a")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "HeaderSliceValidators.VerifyAllErrorUsingActual returns empty -- empty", actual)
}

func Test_Cov8_HeaderSliceValidators_VerifyFirst_Empty(t *testing.T) {
	hsvs := corevalidator.HeaderSliceValidators{}
	err := hsvs.VerifyFirst(&corevalidator.Parameter{}, false)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "HeaderSliceValidators.VerifyFirst returns empty -- empty", actual)
}

func Test_Cov8_HeaderSliceValidators_VerifyUpto_Empty(t *testing.T) {
	hsvs := corevalidator.HeaderSliceValidators{}
	err := hsvs.VerifyUpto(false, false, 1, &corevalidator.Parameter{})
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "HeaderSliceValidators.VerifyUpto returns empty -- empty", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// RangeSegmentsValidator
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov8_RangeSegmentsValidator_LengthOfVerifierSegments(t *testing.T) {
	rsv := &corevalidator.RangeSegmentsValidator{
		VerifierSegments: []corevalidator.RangesSegment{{}, {}},
	}
	actual := args.Map{"len": rsv.LengthOfVerifierSegments()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "RangeSegmentsValidator.LengthOfVerifierSegments returns non-empty -- with args", actual)
}

func Test_Cov8_RangeSegmentsValidator_SetActual(t *testing.T) {
	rsv := &corevalidator.RangeSegmentsValidator{Title: "test"}
	result := rsv.SetActual([]string{"a", "b", "c"})
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RangeSegmentsValidator.SetActual returns non-empty -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SliceValidator — isLengthOkay / isEmptyIgnoreCase
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov8_SliceValidator_Dispose_WithValidators(t *testing.T) {
	sv := &corevalidator.SliceValidator{
		ActualLines:   []string{"a"},
		ExpectedLines: []string{"a"},
		CompareAs:     stringcompareas.Equal,
	}
	_ = sv.ComparingValidators() // force lazy init
	sv.Dispose()
	actual := args.Map{"actualNil": sv.ActualLines == nil, "expectedNil": sv.ExpectedLines == nil}
	expected := args.Map{"actualNil": true, "expectedNil": true}
	expected.ShouldBeEqual(t, 0, "SliceValidator.Dispose returns non-empty -- with validators", actual)
}
