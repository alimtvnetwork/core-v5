package errcore

import (
	"errors"
	"testing"
)

func TestRawErrCollection_BasicOps(t *testing.T) {
	c := &RawErrCollection{}
	if !c.IsEmpty() {
		t.Fatal("should be empty")
	}
	if c.HasError() {
		t.Fatal("should not have error")
	}
	if c.HasAnyError() {
		t.Fatal("should not have any error")
	}
	if c.Length() != 0 {
		t.Fatal("expected 0")
	}
	if c.HasAnyIssues() {
		t.Fatal("should not have issues")
	}
	if !c.IsValid() {
		t.Fatal("should be valid")
	}
	if !c.IsSuccess() {
		t.Fatal("should be success")
	}
	if c.IsFailed() {
		t.Fatal("should not be failed")
	}
	if c.IsInvalid() {
		t.Fatal("should not be invalid")
	}
	if !c.IsDefined() == c.IsEmpty() {
		// ok
	}
	if !c.IsCollectionType() {
		t.Fatal("should be collection type")
	}
	if c.IsNull() {
		t.Fatal("items nil but struct not nil")
	}
}

func TestRawErrCollection_Add(t *testing.T) {
	c := &RawErrCollection{}
	c.Add(nil) // should be ignored
	c.Add(errors.New("e1"))
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestRawErrCollection_AddError(t *testing.T) {
	c := &RawErrCollection{}
	c.AddError(nil)
	c.AddError(errors.New("e"))
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestRawErrCollection_Adds(t *testing.T) {
	c := &RawErrCollection{}
	c.Adds()
	c.Adds(errors.New("a"), nil, errors.New("b"))
	if c.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func TestRawErrCollection_AddErrors(t *testing.T) {
	c := &RawErrCollection{}
	c.AddErrors(errors.New("a"))
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestRawErrCollection_AddString(t *testing.T) {
	c := &RawErrCollection{}
	c.AddString("")
	c.AddString("hello")
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestRawErrCollection_AddMsg(t *testing.T) {
	c := &RawErrCollection{}
	c.AddMsg("hello")
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestRawErrCollection_AddIf(t *testing.T) {
	c := &RawErrCollection{}
	c.AddIf(false, "skip")
	c.AddIf(true, "add")
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestRawErrCollection_AddFunc(t *testing.T) {
	c := &RawErrCollection{}
	c.AddFunc(nil)
	c.AddFunc(func() error { return nil })
	c.AddFunc(func() error { return errors.New("e") })
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestRawErrCollection_AddFuncIf(t *testing.T) {
	c := &RawErrCollection{}
	c.AddFuncIf(false, func() error { return errors.New("e") })
	c.AddFuncIf(true, nil)
	c.AddFuncIf(true, func() error { return errors.New("e") })
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestRawErrCollection_ConditionalAddError(t *testing.T) {
	c := &RawErrCollection{}
	c.ConditionalAddError(false, errors.New("e"))
	c.ConditionalAddError(true, errors.New("e"))
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestRawErrCollection_AddMsgStackTrace(t *testing.T) {
	c := &RawErrCollection{}
	c.AddMsgStackTrace("")
	c.AddMsgStackTrace("msg")
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestRawErrCollection_AddStackTrace(t *testing.T) {
	c := &RawErrCollection{}
	c.AddStackTrace(nil)
	c.AddStackTrace(errors.New("e"))
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestRawErrCollection_AddMsgErrStackTrace(t *testing.T) {
	c := &RawErrCollection{}
	c.AddMsgErrStackTrace("msg", nil)
	c.AddMsgErrStackTrace("msg", errors.New("e"))
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestRawErrCollection_AddMethodName(t *testing.T) {
	c := &RawErrCollection{}
	c.AddMethodName("")
	c.AddMethodName("msg")
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestRawErrCollection_AddMessages(t *testing.T) {
	c := &RawErrCollection{}
	c.AddMessages()
	c.AddMessages("a", "b")
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestRawErrCollection_AddErrorWithMessage(t *testing.T) {
	c := &RawErrCollection{}
	c.AddErrorWithMessage(nil, "msg")
	c.AddErrorWithMessage(errors.New("e"), "msg")
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestRawErrCollection_AddErrorWithMessageRef(t *testing.T) {
	c := &RawErrCollection{}
	c.AddErrorWithMessageRef(nil, "msg", "ref")
	c.AddErrorWithMessageRef(errors.New("e"), "msg", nil)
	c.AddErrorWithMessageRef(errors.New("e"), "msg", "ref")
	if c.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func TestRawErrCollection_AddFmt(t *testing.T) {
	c := &RawErrCollection{}
	c.AddFmt(nil, "fmt %s", "v")
	c.AddFmt(errors.New("e"), "fmt %s", "v")
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestRawErrCollection_Fmt(t *testing.T) {
	c := &RawErrCollection{}
	c.Fmt("", /* no args */ )
	c.Fmt("hello %s", "world")
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestRawErrCollection_FmtIf(t *testing.T) {
	c := &RawErrCollection{}
	c.FmtIf(false, "skip")
	c.FmtIf(true, "add %s", "v")
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestRawErrCollection_References(t *testing.T) {
	c := &RawErrCollection{}
	c.References("msg", "ref1")
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestRawErrCollection_AddWithRef(t *testing.T) {
	c := &RawErrCollection{}
	c.AddWithRef(nil, "ref")
	c.AddWithRef(errors.New("e"), "ref")
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestRawErrCollection_AddWithTraceRef(t *testing.T) {
	c := &RawErrCollection{}
	c.AddWithTraceRef(nil, []string{"t"}, "r")
	c.AddWithTraceRef(errors.New("e"), []string{"t"}, "r")
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestRawErrCollection_AddWithCompiledTraceRef(t *testing.T) {
	c := &RawErrCollection{}
	c.AddWithCompiledTraceRef(nil, "t", "r")
	c.AddWithCompiledTraceRef(errors.New("e"), "t", "r")
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestRawErrCollection_AddStringSliceAsErr(t *testing.T) {
	c := &RawErrCollection{}
	c.AddStringSliceAsErr()
	c.AddStringSliceAsErr("a", "", "b")
	if c.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func TestRawErrCollection_AddErrorGetters(t *testing.T) {
	c := &RawErrCollection{}
	c.AddErrorGetters()
	// can't easily test with actual errorGetter without a mock
}

func TestRawErrCollection_AddCompiledErrorGetters(t *testing.T) {
	c := &RawErrCollection{}
	c.AddCompiledErrorGetters()
}

func TestRawErrCollection_Strings(t *testing.T) {
	c := &RawErrCollection{}
	if len(c.Strings()) != 0 {
		t.Fatal("expected empty")
	}
	c.Add(errors.New("a"))
	if len(c.Strings()) != 1 {
		t.Fatal("expected 1")
	}
}

func TestRawErrCollection_String(t *testing.T) {
	c := &RawErrCollection{}
	if c.String() != "" {
		t.Fatal("expected empty")
	}
	c.Add(errors.New("a"))
	if c.String() == "" {
		t.Fatal("expected non-empty")
	}
}

func TestRawErrCollection_ErrorString(t *testing.T) {
	c := &RawErrCollection{}
	_ = c.ErrorString()
}

func TestRawErrCollection_Compile(t *testing.T) {
	c := &RawErrCollection{}
	_ = c.Compile()
}

func TestRawErrCollection_FullString(t *testing.T) {
	c := &RawErrCollection{}
	_ = c.FullString()
}

func TestRawErrCollection_FullStringWithTraces(t *testing.T) {
	c := &RawErrCollection{}
	_ = c.FullStringWithTraces()
}

func TestRawErrCollection_FullStringWithTracesIf(t *testing.T) {
	c := &RawErrCollection{}
	_ = c.FullStringWithTracesIf(true)
	_ = c.FullStringWithTracesIf(false)
}

func TestRawErrCollection_FullStringSplitByNewLine(t *testing.T) {
	c := &RawErrCollection{}
	_ = c.FullStringSplitByNewLine()
}

func TestRawErrCollection_FullStringWithoutReferences(t *testing.T) {
	c := &RawErrCollection{}
	_ = c.FullStringWithoutReferences()
}

func TestRawErrCollection_ReferencesCompiledString(t *testing.T) {
	c := &RawErrCollection{}
	_ = c.ReferencesCompiledString()
}

func TestRawErrCollection_StringUsingJoiner(t *testing.T) {
	c := &RawErrCollection{}
	if c.StringUsingJoiner(",") != "" {
		t.Fatal("expected empty")
	}
	c.Add(errors.New("a"))
	if c.StringUsingJoiner(",") == "" {
		t.Fatal("expected non-empty")
	}
}

func TestRawErrCollection_StringUsingJoinerAdditional(t *testing.T) {
	c := &RawErrCollection{}
	if c.StringUsingJoinerAdditional(",", "!") != "" {
		t.Fatal("expected empty")
	}
	c.Add(errors.New("a"))
	if c.StringUsingJoinerAdditional(",", "!") == "" {
		t.Fatal("expected non-empty")
	}
}

func TestRawErrCollection_StringWithAdditionalMessage(t *testing.T) {
	c := &RawErrCollection{}
	if c.StringWithAdditionalMessage("!") != "" {
		t.Fatal("expected empty")
	}
	c.Add(errors.New("a"))
	if c.StringWithAdditionalMessage("!") == "" {
		t.Fatal("expected non-empty")
	}
}

func TestRawErrCollection_CompiledError(t *testing.T) {
	c := RawErrCollection{}
	if c.CompiledError() != nil {
		t.Fatal("expected nil")
	}
	c.Items = append(c.Items, errors.New("a"))
	if c.CompiledError() == nil {
		t.Fatal("expected non-nil")
	}
}

func TestRawErrCollection_CompiledErrorUsingJoiner(t *testing.T) {
	c := &RawErrCollection{}
	if c.CompiledErrorUsingJoiner(",") != nil {
		t.Fatal("expected nil")
	}
	c.Add(errors.New("a"))
	if c.CompiledErrorUsingJoiner(",") == nil {
		t.Fatal("expected non-nil")
	}
}

func TestRawErrCollection_CompiledErrorUsingJoinerAdditionalMessage(t *testing.T) {
	c := &RawErrCollection{}
	if c.CompiledErrorUsingJoinerAdditionalMessage(",", "!") != nil {
		t.Fatal("expected nil")
	}
	c.Add(errors.New("a"))
	if c.CompiledErrorUsingJoinerAdditionalMessage(",", "!") == nil {
		t.Fatal("expected non-nil")
	}
}

func TestRawErrCollection_CompiledErrorWithStackTraces(t *testing.T) {
	c := &RawErrCollection{}
	if c.CompiledErrorWithStackTraces() != nil {
		t.Fatal("expected nil")
	}
	c.Add(errors.New("a"))
	if c.CompiledErrorWithStackTraces() == nil {
		t.Fatal("expected non-nil")
	}
}

func TestRawErrCollection_CompiledStackTracesString(t *testing.T) {
	c := &RawErrCollection{}
	if c.CompiledStackTracesString() != "" {
		t.Fatal("expected empty")
	}
	c.Add(errors.New("a"))
	if c.CompiledStackTracesString() == "" {
		t.Fatal("expected non-empty")
	}
}

func TestRawErrCollection_CompiledErrorUsingStackTraces(t *testing.T) {
	c := &RawErrCollection{}
	if c.CompiledErrorUsingStackTraces(",", []string{"t"}) != nil {
		t.Fatal("expected nil")
	}
	c.Add(errors.New("a"))
	if c.CompiledErrorUsingStackTraces(",", []string{"t"}) == nil {
		t.Fatal("expected non-nil")
	}
}

func TestRawErrCollection_CompiledJsonErrorWithStackTraces(t *testing.T) {
	c := &RawErrCollection{}
	c.Add(errors.New("e"))
	_ = c.CompiledJsonErrorWithStackTraces()
}

func TestRawErrCollection_CompiledJsonStringWithStackTraces(t *testing.T) {
	c := &RawErrCollection{}
	_ = c.CompiledJsonStringWithStackTraces()
}

func TestRawErrCollection_Serialize(t *testing.T) {
	c := &RawErrCollection{}
	b, err := c.Serialize()
	if b != nil || err != nil {
		t.Fatal("expected nil")
	}
}

func TestRawErrCollection_SerializeWithoutTraces(t *testing.T) {
	c := &RawErrCollection{}
	b, err := c.SerializeWithoutTraces()
	if b != nil || err != nil {
		t.Fatal("expected nil")
	}
}

func TestRawErrCollection_MarshalJSON(t *testing.T) {
	c := &RawErrCollection{}
	b, err := c.MarshalJSON()
	if b != nil || err != nil {
		t.Fatal("expected nil")
	}
}

func TestRawErrCollection_UnmarshalJSON(t *testing.T) {
	c := &RawErrCollection{}
	_ = c.UnmarshalJSON([]byte("[]"))
}

func TestRawErrCollection_Value(t *testing.T) {
	c := &RawErrCollection{}
	if c.Value() != nil {
		t.Fatal("expected nil")
	}
}

func TestRawErrCollection_Log(t *testing.T) {
	c := &RawErrCollection{}
	c.Log() // empty
	c.Add(errors.New("e"))
	c.Log() // with error
}

func TestRawErrCollection_LogWithTraces(t *testing.T) {
	c := &RawErrCollection{}
	c.LogWithTraces()
	c.Add(errors.New("e"))
	c.LogWithTraces()
}

func TestRawErrCollection_LogIf(t *testing.T) {
	c := &RawErrCollection{}
	c.LogIf(false)
}

func TestRawErrCollection_IsNull(t *testing.T) {
	c := &RawErrCollection{}
	_ = c.IsNull()
	_ = c.IsAnyNull()
}

func TestRawErrCollection_ClearDispose(t *testing.T) {
	c := &RawErrCollection{}
	c.Clear() // empty
	c.Add(errors.New("e"))
	c.Clear()
	if c.Length() != 0 {
		t.Fatal("expected 0")
	}

	c2 := &RawErrCollection{}
	c2.Dispose() // empty
	c2.Add(errors.New("e"))
	c2.Dispose()
}

func TestRawErrCollection_IsErrorsCollected(t *testing.T) {
	c := &RawErrCollection{}
	if c.IsErrorsCollected(nil) {
		t.Fatal("should not collect nil")
	}
	if !c.IsErrorsCollected(errors.New("e")) {
		t.Fatal("should collect")
	}
}

func TestRawErrCollection_CountStateChangeTracker(t *testing.T) {
	c := RawErrCollection{}
	tracker := c.CountStateChangeTracker()
	if !tracker.IsSameState() {
		t.Fatal("should be same")
	}
}

func TestRawErrCollection_ToRawErrCollection(t *testing.T) {
	c := RawErrCollection{}
	p := c.ToRawErrCollection()
	if p == nil {
		t.Fatal("expected non-nil")
	}
}

func TestRawErrCollection_ReflectSetTo(t *testing.T) {
	c := &RawErrCollection{}
	// value type
	err := c.ReflectSetTo(RawErrCollection{})
	if err == nil {
		t.Fatal("expected error")
	}
	// nil ptr
	var nilP *RawErrCollection
	err2 := c.ReflectSetTo(nilP)
	if err2 == nil {
		t.Fatal("expected error")
	}
	// valid ptr
	target := &RawErrCollection{}
	err3 := c.ReflectSetTo(target)
	if err3 != nil {
		t.Fatal("unexpected error")
	}
	// other type
	err4 := c.ReflectSetTo("other")
	if err4 == nil {
		t.Fatal("expected error")
	}
}
