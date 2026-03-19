package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ── SimpleSlice ──

func Test_C32_SS_Add(t *testing.T) {
	ss := corestr.New.SimpleSlice.Empty()
	ss.Add("a")
	if ss.Length() != 1 { t.Fatal("expected 1") }
}

func Test_C32_SS_AddSplit(t *testing.T) {
	ss := corestr.New.SimpleSlice.Empty()
	ss.AddSplit("a,b", ",")
}

func Test_C32_SS_AddIf(t *testing.T) {
	ss := corestr.New.SimpleSlice.Empty()
	ss.AddIf(false, "skip")
	ss.AddIf(true, "add")
	if ss.Length() != 1 { t.Fatal("expected 1") }
}

func Test_C32_SS_Adds(t *testing.T) {
	ss := corestr.New.SimpleSlice.Empty()
	ss.Adds("a", "b")
	ss.Adds()
}

func Test_C32_SS_Append(t *testing.T) {
	ss := corestr.New.SimpleSlice.Empty()
	ss.Append("a")
	ss.Append()
}

func Test_C32_SS_AppendFmt(t *testing.T) {
	ss := corestr.New.SimpleSlice.Empty()
	ss.AppendFmt("hello %s", "world")
	ss.AppendFmt("")
}

func Test_C32_SS_AppendFmtIf(t *testing.T) {
	ss := corestr.New.SimpleSlice.Empty()
	ss.AppendFmtIf(true, "x%d", 1)
	ss.AppendFmtIf(false, "skip")
}

func Test_C32_SS_Length(t *testing.T) {
	ss := corestr.New.SimpleSlice.Lines("a")
	if ss.Length() != 1 { t.Fatal("expected 1") }
	var n *corestr.SimpleSlice
	if n.Length() != 0 { t.Fatal("expected 0") }
}

func Test_C32_SS_IsEmpty(t *testing.T) {
	ss := corestr.New.SimpleSlice.Empty()
	if !ss.IsEmpty() { t.Fatal("expected true") }
}

func Test_C32_SS_HasAnyItem(t *testing.T) {
	ss := corestr.New.SimpleSlice.Lines("a")
	if !ss.HasAnyItem() { t.Fatal("expected true") }
}

func Test_C32_SS_First(t *testing.T) {
	ss := corestr.New.SimpleSlice.Lines("a")
	if ss.First() != "a" { t.Fatal("expected a") }
}

func Test_C32_SS_Last(t *testing.T) {
	ss := corestr.New.SimpleSlice.Lines("a", "b")
	if ss.Last() != "b" { t.Fatal("expected b") }
}

func Test_C32_SS_FirstOrDefault(t *testing.T) {
	ss := corestr.New.SimpleSlice.Empty()
	_ = ss.FirstOrDefault()
	ss.Add("a")
	_ = ss.FirstOrDefault()
}

func Test_C32_SS_LastOrDefault(t *testing.T) {
	ss := corestr.New.SimpleSlice.Empty()
	_ = ss.LastOrDefault()
	ss.Add("a")
	_ = ss.LastOrDefault()
}

func Test_C32_SS_Strings(t *testing.T) {
	ss := corestr.New.SimpleSlice.Lines("a")
	_ = ss.Strings()
}

func Test_C32_SS_Join(t *testing.T) {
	ss := corestr.New.SimpleSlice.Lines("a", "b")
	_ = ss.Join(",")
}

func Test_C32_SS_JoinLine(t *testing.T) {
	ss := corestr.New.SimpleSlice.Lines("a", "b")
	_ = ss.JoinLine()
}

func Test_C32_SS_Take(t *testing.T) {
	ss := corestr.New.SimpleSlice.Lines("a", "b", "c")
	_ = ss.Take(2)
	_ = ss.Take(5)
}

func Test_C32_SS_Skip(t *testing.T) {
	ss := corestr.New.SimpleSlice.Lines("a", "b", "c")
	_ = ss.Skip(1)
}

func Test_C32_SS_Collection(t *testing.T) {
	ss := corestr.New.SimpleSlice.Lines("a")
	_ = ss.Collection(false)
}

func Test_C32_SS_Hashset(t *testing.T) {
	ss := corestr.New.SimpleSlice.Lines("a")
	_ = ss.Hashset()
}

func Test_C32_SS_Clear(t *testing.T) {
	ss := corestr.New.SimpleSlice.Lines("a")
	ss.Clear()
}

func Test_C32_SS_Dispose(t *testing.T) {
	ss := corestr.New.SimpleSlice.Lines("a")
	ss.Dispose()
}

func Test_C32_SS_Clone(t *testing.T) {
	ss := corestr.New.SimpleSlice.Lines("a")
	_ = ss.Clone(false)
}

func Test_C32_SS_DeepClone(t *testing.T) {
	ss := corestr.New.SimpleSlice.Lines("a")
	_ = ss.DeepClone()
}

func Test_C32_SS_AddAsTitleValue(t *testing.T) {
	ss := corestr.New.SimpleSlice.Empty()
	ss.AddAsTitleValue("key", "val")
}

func Test_C32_SS_AddAsTitleValueIf(t *testing.T) {
	ss := corestr.New.SimpleSlice.Empty()
	ss.AddAsTitleValueIf(true, "key", "val")
	ss.AddAsTitleValueIf(false, "key", "val")
}

func Test_C32_SS_AddAsCurlyTitleWrap(t *testing.T) {
	ss := corestr.New.SimpleSlice.Empty()
	ss.AddAsCurlyTitleWrap("title", "content")
}

func Test_C32_SS_AddAsCurlyTitleWrapIf(t *testing.T) {
	ss := corestr.New.SimpleSlice.Empty()
	ss.AddAsCurlyTitleWrapIf(true, "title", "content")
	ss.AddAsCurlyTitleWrapIf(false, "title", "content")
}

func Test_C32_SS_AddError(t *testing.T) {
	ss := corestr.New.SimpleSlice.Empty()
	ss.AddError(nil)
}

func Test_C32_SS_AddIf2(t *testing.T) {
	ss := corestr.New.SimpleSlice.Empty()
	ss.AddIf(true, "a")
	ss.AddIf(false, "c")
}

func Test_C32_SS_Sort(t *testing.T) {
	ss := corestr.New.SimpleSlice.Lines("b", "a")
	_ = ss.Sort()
}

func Test_C32_SS_Reverse(t *testing.T) {
	ss := corestr.New.SimpleSlice.Lines("a", "b")
	ss.Reverse()
}

func Test_C32_SS_IsContains(t *testing.T) {
	ss := corestr.New.SimpleSlice.Lines("a")
	if !ss.IsContains("a") { t.Fatal("expected true") }
}

func Test_C32_SS_IndexOf(t *testing.T) {
	ss := corestr.New.SimpleSlice.Lines("a", "b")
	_ = ss.IndexOf("a")
}

func Test_C32_SS_HasIndex(t *testing.T) {
	ss := corestr.New.SimpleSlice.Lines("a")
	_ = ss.HasIndex(0)
	_ = ss.HasIndex(99)
}

// ── newSimpleSliceCreator ──

func Test_C32_NSSC_Cap(t *testing.T)     { _ = corestr.New.SimpleSlice.Cap(5) }
func Test_C32_NSSC_Default(t *testing.T) { _ = corestr.New.SimpleSlice.Default() }
func Test_C32_NSSC_Deserialize(t *testing.T) {
	_, _ = corestr.New.SimpleSlice.Deserialize([]byte(`["a"]`))
	_, _ = corestr.New.SimpleSlice.Deserialize([]byte(`invalid`))
}
func Test_C32_NSSC_DeserializeJsoner(t *testing.T) {
	r := corejson.New([]string{"a"})
	_, _ = corestr.New.SimpleSlice.DeserializeJsoner(&r)
}
func Test_C32_NSSC_UsingLines(t *testing.T) {
	_ = corestr.New.SimpleSlice.UsingLines(true, "a")
	_ = corestr.New.SimpleSlice.UsingLines(false, "a")
	_ = corestr.New.SimpleSlice.UsingLines(false)
}
func Test_C32_NSSC_Lines(t *testing.T)         { _ = corestr.New.SimpleSlice.Lines("a") }
func Test_C32_NSSC_Split(t *testing.T)         { _ = corestr.New.SimpleSlice.Split("a,b", ",") }
func Test_C32_NSSC_SplitLines(t *testing.T)    { _ = corestr.New.SimpleSlice.SplitLines("a\nb") }
func Test_C32_NSSC_SpreadStrings(t *testing.T) { _ = corestr.New.SimpleSlice.SpreadStrings("a") }
func Test_C32_NSSC_Hashset(t *testing.T) {
	_ = corestr.New.SimpleSlice.Hashset(corestr.New.Hashset.StringsSpreadItems("a"))
	_ = corestr.New.SimpleSlice.Hashset(corestr.New.Hashset.Empty())
}
func Test_C32_NSSC_Map(t *testing.T) {
	_ = corestr.New.SimpleSlice.Map(map[string]string{"k": "v"})
	_ = corestr.New.SimpleSlice.Map(map[string]string{})
}
func Test_C32_NSSC_Create(t *testing.T)        { _ = corestr.New.SimpleSlice.Create([]string{"a"}) }
func Test_C32_NSSC_Strings(t *testing.T)       { _ = corestr.New.SimpleSlice.Strings([]string{"a"}) }
func Test_C32_NSSC_StringsPtr(t *testing.T)    { _ = corestr.New.SimpleSlice.StringsPtr(nil) }
func Test_C32_NSSC_StringsOptions(t *testing.T) {
	_ = corestr.New.SimpleSlice.StringsOptions(true, []string{"a"})
	_ = corestr.New.SimpleSlice.StringsOptions(false, []string{})
}
func Test_C32_NSSC_StringsClone(t *testing.T) {
	_ = corestr.New.SimpleSlice.StringsClone([]string{"a"})
	_ = corestr.New.SimpleSlice.StringsClone(nil)
}
func Test_C32_NSSC_Direct(t *testing.T) {
	_ = corestr.New.SimpleSlice.Direct(true, []string{"a"})
	_ = corestr.New.SimpleSlice.Direct(false, nil)
}
func Test_C32_NSSC_UsingSeparatorLine(t *testing.T) { _ = corestr.New.SimpleSlice.UsingSeparatorLine(",", "a,b") }
func Test_C32_NSSC_UsingLine(t *testing.T)          { _ = corestr.New.SimpleSlice.UsingLine("a|b") }
func Test_C32_NSSC_Empty(t *testing.T)              { _ = corestr.New.SimpleSlice.Empty() }
func Test_C32_NSSC_ByLen(t *testing.T)              { _ = corestr.New.SimpleSlice.ByLen([]string{"a"}) }

// ── SimpleStringOnce ──

func Test_C32_SSO_Value(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("hello")
	if sso.Value() != "hello" { t.Fatal("expected hello") }
}

func Test_C32_SSO_IsInitialized(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("x")
	if !sso.IsInitialized() { t.Fatal("expected true") }
}

func Test_C32_SSO_IsDefined(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("x")
	if !sso.IsDefined() { t.Fatal("expected true") }
}

func Test_C32_SSO_IsUninitialized(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Empty()
	if !sso.IsUninitialized() { t.Fatal("expected true") }
}

func Test_C32_SSO_Invalidate(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("x")
	sso.Invalidate()
}

func Test_C32_SSO_Reset(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("x")
	sso.Reset()
}

func Test_C32_SSO_IsInvalid(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Empty()
	if !sso.IsInvalid() { t.Fatal("expected true") }
}

func Test_C32_SSO_ValueBytes(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("x")
	_ = sso.ValueBytes()
	_ = sso.ValueBytesPtr()
}

func Test_C32_SSO_SetOnUninitialized(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Empty()
	_ = sso.SetOnUninitialized("v")
	err := sso.SetOnUninitialized("v2")
	if err == nil { t.Fatal("expected error") }
}

func Test_C32_SSO_GetSetOnce(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Empty()
	v := sso.GetSetOnce("hello")
	if v != "hello" { t.Fatal("expected hello") }
	v2 := sso.GetSetOnce("world")
	if v2 != "hello" { t.Fatal("expected hello still") }
}

func Test_C32_SSO_GetOnce(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("x")
	_ = sso.GetOnce()
}

func Test_C32_SSO_SetInitialize(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Empty()
	sso.SetInitialize()
}

func Test_C32_SSO_String(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("x")
	_ = sso.String()
}

func Test_C32_SSO_Dispose(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("x")
	sso.Dispose()
}

func Test_C32_SSO_Boolean(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("true")
	_ = sso.Boolean(true)
}

func Test_C32_SSO_Int(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("42")
	_ = sso.Int()
}

func Test_C32_SSO_ValueDefFloat64(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("3.14")
	_ = sso.ValueDefFloat64()
}

func Test_C32_SSO_JsonMethods(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("x")
	_ = sso.Json()
	_ = sso.JsonPtr()
	_ = sso.JsonModel()
	_ = sso.JsonModelAny()
	_, _ = sso.MarshalJSON()
	_ = sso.AsJsonContractsBinder()
	_ = sso.AsJsoner()
	_ = sso.AsJsonParseSelfInjector()
	_ = sso.AsJsonMarshaller()
	_, _ = sso.Serialize()
}

func Test_C32_SSO_UnmarshalJSON(t *testing.T) {
	sso := corestr.SimpleStringOnce{}
	_ = sso.UnmarshalJSON([]byte(`"hello"`))
}

func Test_C32_SSO_ParseInjectUsingJson(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Empty()
	r := corejson.New("hello")
	_, _ = sso.ParseInjectUsingJson(&r)
}

func Test_C32_SSO_ParseInjectUsingJsonMust_Panic(t *testing.T) {
	defer func() { recover() }()
	sso := corestr.New.SimpleStringOnce.Empty()
	bad := corejson.NewResult.UsingString(`invalid`)
	sso.ParseInjectUsingJsonMust(bad)
}

func Test_C32_SSO_JsonParseSelfInject(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Empty()
	r := corejson.New("hello")
	_ = sso.JsonParseSelfInject(&r)
}

func Test_C32_SSO_Deserialize(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("x")
	var s string
	_ = sso.Deserialize(&s)
}

func Test_C32_SSO_IsEmpty(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Empty()
	_ = sso.IsEmpty()
}

func Test_C32_SSO_HasNonEmpty(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("x")
	_ = sso.HasNonEmpty()
}

func Test_C32_SSO_IsWhitespace(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("  ")
	_ = sso.IsWhitespace()
}

func Test_C32_SSO_HasNonWhitespace(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("x")
	_ = sso.HasNonWhitespace()
}

func Test_C32_SSO_IsEqualText(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("x")
	_ = sso.IsEqualText("x")
}

func Test_C32_SSO_HasPrefix(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("hello")
	_ = sso.HasPrefix("hel")
}

func Test_C32_SSO_HasSuffix(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("hello")
	_ = sso.HasSuffix("llo")
}

func Test_C32_SSO_Contains(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("hello")
	_ = sso.Contains("ell")
}

func Test_C32_SSO_IsMatchRegex(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("hello123")
	_ = sso.IsMatchRegex("[0-9]+")
}

func Test_C32_SSO_Trim(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("  x  ")
	_ = sso.Trim()
}

func Test_C32_SSO_ToLower(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("HELLO")
	_ = sso.ToLower()
}

func Test_C32_SSO_ToUpper(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("hello")
	_ = sso.ToUpper()
}

func Test_C32_SSO_ValueLength(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("hello")
	_ = sso.ValueLength()
}

// ── newSimpleStringOnceCreator ──

func Test_C32_NSSOC_Any(t *testing.T)       { _ = corestr.New.SimpleStringOnce.Any(42) }
func Test_C32_NSSOC_Create(t *testing.T)     { _ = corestr.New.SimpleStringOnce.Create(true, "x") }
func Test_C32_NSSOC_CreatePtr(t *testing.T)  { _ = corestr.New.SimpleStringOnce.CreatePtr(true, "x") }
func Test_C32_NSSOC_Empty(t *testing.T)      { _ = corestr.New.SimpleStringOnce.Empty() }
func Test_C32_NSSOC_Init(t *testing.T)       { _ = corestr.New.SimpleStringOnce.Init("x") }
func Test_C32_NSSOC_InitPtr(t *testing.T)    { _ = corestr.New.SimpleStringOnce.InitPtr("x") }
func Test_C32_NSSOC_Uninitialized(t *testing.T) { _ = corestr.New.SimpleStringOnce.Uninitialized("x") }
