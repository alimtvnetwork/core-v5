package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ── HashmapDiff ──

func Test_C34_HD_Length(t *testing.T) {
	hd := corestr.HashmapDiff(map[string]string{"k": "v"})
	if hd.Length() != 1 { t.Fatal("expected 1") }
	var nilHd *corestr.HashmapDiff
	if nilHd.Length() != 0 { t.Fatal("expected 0") }
}

func Test_C34_HD_IsEmpty(t *testing.T)   { _ = corestr.HashmapDiff(map[string]string{}).IsEmpty() }
func Test_C34_HD_HasAnyItem(t *testing.T) { _ = corestr.HashmapDiff(map[string]string{"k": "v"}).HasAnyItem() }
func Test_C34_HD_LastIndex(t *testing.T)  { _ = corestr.HashmapDiff(map[string]string{"k": "v"}).LastIndex() }

func Test_C34_HD_AllKeysSorted(t *testing.T) {
	hd := corestr.HashmapDiff(map[string]string{"b": "2", "a": "1"})
	_ = hd.AllKeysSorted()
}

func Test_C34_HD_MapAnyItems(t *testing.T) {
	hd := corestr.HashmapDiff(map[string]string{"k": "v"})
	_ = hd.MapAnyItems()
	var nilHd *corestr.HashmapDiff
	_ = nilHd.MapAnyItems()
}

func Test_C34_HD_HasAnyChanges(t *testing.T) {
	hd := corestr.HashmapDiff(map[string]string{"k": "v"})
	_ = hd.HasAnyChanges(map[string]string{"k": "v2"})
}

func Test_C34_HD_IsRawEqual(t *testing.T) {
	hd := corestr.HashmapDiff(map[string]string{"k": "v"})
	_ = hd.IsRawEqual(map[string]string{"k": "v"})
}

func Test_C34_HD_HashmapDiffUsingRaw(t *testing.T) {
	hd := corestr.HashmapDiff(map[string]string{"k": "v"})
	_ = hd.HashmapDiffUsingRaw(map[string]string{"k": "v2"})
	_ = hd.HashmapDiffUsingRaw(map[string]string{"k": "v"})
}

func Test_C34_HD_DiffRaw(t *testing.T) {
	hd := corestr.HashmapDiff(map[string]string{"k": "v"})
	_ = hd.DiffRaw(map[string]string{"k": "v2"})
}

func Test_C34_HD_DiffJsonMessage(t *testing.T) {
	hd := corestr.HashmapDiff(map[string]string{"k": "v"})
	_ = hd.DiffJsonMessage(map[string]string{"k": "v2"})
}

func Test_C34_HD_ToStringsSliceOfDiffMap(t *testing.T) {
	hd := corestr.HashmapDiff(map[string]string{"k": "v"})
	_ = hd.ToStringsSliceOfDiffMap(map[string]string{"k": "v2"})
}

func Test_C34_HD_ShouldDiffMessage(t *testing.T) {
	hd := corestr.HashmapDiff(map[string]string{"k": "v"})
	_ = hd.ShouldDiffMessage("test", map[string]string{"k": "v2"})
}

func Test_C34_HD_LogShouldDiffMessage(t *testing.T) {
	hd := corestr.HashmapDiff(map[string]string{"k": "v"})
	_ = hd.LogShouldDiffMessage("test", map[string]string{"k": "v2"})
}

func Test_C34_HD_Raw(t *testing.T) {
	hd := corestr.HashmapDiff(map[string]string{"k": "v"})
	_ = hd.Raw()
	var nilHd *corestr.HashmapDiff
	_ = nilHd.Raw()
}

func Test_C34_HD_RawMapStringAnyDiff(t *testing.T) {
	hd := corestr.HashmapDiff(map[string]string{"k": "v"})
	_ = hd.RawMapStringAnyDiff()
}

func Test_C34_HD_Serialize(t *testing.T) {
	hd := corestr.HashmapDiff(map[string]string{"k": "v"})
	_, _ = hd.Serialize()
}

func Test_C34_HD_Deserialize(t *testing.T) {
	hd := corestr.HashmapDiff(map[string]string{"k": "v"})
	var target map[string]string
	_ = hd.Deserialize(&target)
}

// ── KeyValuePair ──

func Test_C34_KVP_Methods(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "v"}
	_ = kv.KeyName()
	_ = kv.VariableName()
	_ = kv.ValueString()
	_ = kv.IsVariableNameEqual("k")
	_ = kv.IsValueEqual("v")
	_ = kv.Json()
	_ = kv.JsonPtr()
	_, _ = kv.Serialize()
	_ = kv.SerializeMust()
	_ = kv.Compile()
	_ = kv.IsKeyEmpty()
	_ = kv.IsValueEmpty()
	_ = kv.HasKey()
	_ = kv.HasValue()
	_ = kv.IsKeyValueEmpty()
	_ = kv.TrimKey()
	_ = kv.TrimValue()
	_ = kv.String()
	_ = kv.FormatString("%s=%s")
	kv.Clear()
	kv2 := corestr.KeyValuePair{Key: "k", Value: "v"}
	kv2.Dispose()
	_ = kv2.IsKey("k")
	_ = kv2.Clone()
}

// ── KeyAnyValuePair ──

func Test_C34_KAVP_Methods(t *testing.T) {
	kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
	_ = kv.KeyName()
	_ = kv.VariableName()
	_ = kv.ValueAny()
	_ = kv.IsVariableNameEqual("k")
	_ = kv.SerializeMust()
	_ = kv.Compile()
	_ = kv.IsValueNull()
	_ = kv.HasNonNull()
	_ = kv.HasValue()
	_ = kv.IsValueEmptyString()
	_ = kv.IsValueWhitespace()
	_ = kv.ValueString()
	_, _ = kv.Serialize()
	_ = kv.Json()
	_ = kv.JsonPtr()
	_ = kv.String()
	_ = kv.AsJsonContractsBinder()
	_ = kv.AsJsoner()
	_ = kv.AsJsonParseSelfInjector()
	kv.Clear()
	kv2 := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
	kv2.Dispose()
}

func Test_C34_KAVP_ParseInjectUsingJson(t *testing.T) {
	kv := &corestr.KeyAnyValuePair{}
	r := corejson.New(corestr.KeyAnyValuePair{Key: "k", Value: "v"})
	_, _ = kv.ParseInjectUsingJson(&r)
}

func Test_C34_KAVP_ParseInjectUsingJsonMust_Panic(t *testing.T) {
	defer func() { recover() }()
	kv := &corestr.KeyAnyValuePair{}
	bad := corejson.NewResult.UsingString(`invalid`)
	kv.ParseInjectUsingJsonMust(bad)
}

func Test_C34_KAVP_JsonParseSelfInject(t *testing.T) {
	kv := &corestr.KeyAnyValuePair{}
	r := corejson.New(corestr.KeyAnyValuePair{Key: "k", Value: "v"})
	_ = kv.JsonParseSelfInject(&r)
}

// ── KeyValueCollection ──

func Test_C34_KVC_Basic(t *testing.T) {
	kvc := corestr.New.KeyValues.Empty()
	kvc.Add(corestr.KeyValuePair{Key: "k", Value: "v"})
	_ = kvc.Length()
	_ = kvc.Count()
	_ = kvc.HasAnyItem()
	_ = kvc.LastIndex()
	_ = kvc.HasIndex(0)
	_ = kvc.First()
	_ = kvc.FirstOrDefault()
	_ = kvc.Last()
	_ = kvc.LastOrDefault()
	_ = kvc.IsEmpty()
	_ = kvc.HasKey("k")
	_ = kvc.AllKeys()
	_ = kvc.AllKeysSorted()
	_ = kvc.AllValues()
	_ = kvc.Compile()
	_ = kvc.String()
	_ = kvc.SerializeMust()
}

func Test_C34_KVC_Find(t *testing.T) {
	kvc := corestr.New.KeyValues.Empty()
	kvc.Add(corestr.KeyValuePair{Key: "k", Value: "v"})
	_ = kvc.Find(func(i int, kv corestr.KeyValuePair) (corestr.KeyValuePair, bool, bool) {
		return kv, true, false
	})
}

func Test_C34_KVC_AddIf(t *testing.T) {
	kvc := corestr.New.KeyValues.Empty()
	kvc.AddIf(true, corestr.KeyValuePair{Key: "k", Value: "v"})
	kvc.AddIf(false, corestr.KeyValuePair{Key: "k2", Value: "v2"})
}

func Test_C34_KVC_AddMap(t *testing.T) {
	kvc := corestr.New.KeyValues.Empty()
	kvc.AddMap(map[string]string{"k": "v"})
}

func Test_C34_KVC_AddHashset(t *testing.T) {
	kvc := corestr.New.KeyValues.Empty()
	kvc.AddHashset(corestr.New.Hashset.StringsSpreadItems("a"))
}

func Test_C34_KVC_AddHashsetMap(t *testing.T) {
	kvc := corestr.New.KeyValues.Empty()
	kvc.AddHashsetMap(corestr.New.Hashmap.Empty())
}

func Test_C34_KVC_GetValueByKey(t *testing.T) {
	kvc := corestr.New.KeyValues.Empty()
	kvc.Add(corestr.KeyValuePair{Key: "k", Value: "v"})
	_ = kvc.GetValueByKey("k")
	_ = kvc.GetValueByKey("missing")
}

func Test_C34_KVC_Adds(t *testing.T) {
	kvc := corestr.New.KeyValues.Empty()
	kvc.Adds(corestr.KeyValuePair{Key: "k", Value: "v"})
}

func Test_C34_KVC_AddKeyValues(t *testing.T) {
	kvc := corestr.New.KeyValues.Empty()
	kvc.AddKeyValues("k", "v")
}

func Test_C34_KVC_Hashmap(t *testing.T) {
	kvc := corestr.New.KeyValues.Empty()
	kvc.Add(corestr.KeyValuePair{Key: "k", Value: "v"})
	_ = kvc.Hashmap()
	_ = corestr.New.KeyValues.Empty().Hashmap()
}

func Test_C34_KVC_HashmapOptions(t *testing.T) {
	kvc := corestr.New.KeyValues.Empty()
	kvc.Add(corestr.KeyValuePair{Key: "k", Value: "v"})
	_ = kvc.HashmapOptions(true)
}

func Test_C34_KVC_Clear(t *testing.T) {
	kvc := corestr.New.KeyValues.Empty()
	kvc.Add(corestr.KeyValuePair{Key: "k", Value: "v"})
	kvc.Clear()
}

func Test_C34_KVC_Dispose(t *testing.T) {
	kvc := corestr.New.KeyValues.Empty()
	kvc.Dispose()
}

func Test_C34_KVC_JsonMethods(t *testing.T) {
	kvc := corestr.New.KeyValues.Empty()
	kvc.Add(corestr.KeyValuePair{Key: "k", Value: "v"})
	_ = kvc.Json()
	_ = kvc.JsonPtr()
	_ = kvc.JsonModel()
	_ = kvc.JsonModelAny()
	_, _ = kvc.MarshalJSON()
	_ = kvc.AsJsonContractsBinder()
	_ = kvc.AsJsoner()
	_ = kvc.AsJsonMarshaller()
	_ = kvc.AsJsonParseSelfInjector()
	_, _ = kvc.Serialize()
}

// ── newKeyValuesCreator ──

func Test_C34_NKVC_Empty(t *testing.T) { _ = corestr.New.KeyValues.Empty() }
func Test_C34_NKVC_Cap(t *testing.T)   { _ = corestr.New.KeyValues.Cap(5) }
func Test_C34_NKVC_UsingKeyValuePairs(t *testing.T) {
	_ = corestr.New.KeyValues.UsingKeyValuePairs(corestr.KeyValuePair{Key: "k", Value: "v"})
}
func Test_C34_NKVC_UsingKeyValueStrings(t *testing.T) {
	_ = corestr.New.KeyValues.UsingKeyValueStrings("k", "v")
}
func Test_C34_NKVC_UsingMap(t *testing.T) {
	_ = corestr.New.KeyValues.UsingMap(map[string]string{"k": "v"})
}

// ── LeftRight ──

func Test_C34_LR_InvalidLeftRight(t *testing.T) {
	_ = corestr.InvalidLeftRight("msg")
	_ = corestr.InvalidLeftRightNoMessage()
}

func Test_C34_LR_Methods(t *testing.T) {
	lr := corestr.NewLeftRight("a", "b")
	_ = lr.LeftBytes()
	_ = lr.RightBytes()
	_ = lr.LeftTrim()
	_ = lr.RightTrim()
	_ = lr.IsLeftEmpty()
	_ = lr.IsRightEmpty()
	_ = lr.IsLeftWhitespace()
	_ = lr.IsRightWhitespace()
	_ = lr.HasValidNonEmptyLeft()
	_ = lr.HasValidNonEmptyRight()
	_ = lr.HasValidNonWhitespaceLeft()
	_ = lr.HasValidNonWhitespaceRight()
	_ = lr.HasSafeNonEmpty()
	_ = lr.Is("a", "b")
	_ = lr.Clone()
	_ = lr.String()
	lr.Clear()
	lr2 := corestr.NewLeftRight("a", "b")
	lr2.Dispose()
}

// ── LeftRightFromSplit ──

func Test_C34_LRFS_Methods(t *testing.T) {
	_ = corestr.LeftRightFromSplit("a=b", "=")
	_ = corestr.LeftRightFromSplitFull("a=b", "=")
	_ = corestr.LeftRightFromSplitTrimmed("a = b", "=")
	_ = corestr.LeftRightFromSplitFullTrimmed("a = b", "=")
}

// ── LeftMiddleRight ──

func Test_C34_LMR_Invalid(t *testing.T) {
	_ = corestr.InvalidLeftMiddleRight("msg")
	_ = corestr.InvalidLeftMiddleRightNoMessage()
}

func Test_C34_LMR_Methods(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("a", "b", "c")
	_ = lmr.LeftBytes()
	_ = lmr.RightBytes()
	_ = lmr.MiddleBytes()
	_ = lmr.LeftTrim()
	_ = lmr.RightTrim()
	_ = lmr.MiddleTrim()
	_ = lmr.IsLeftEmpty()
	_ = lmr.IsRightEmpty()
	_ = lmr.IsMiddleEmpty()
	_ = lmr.IsMiddleWhitespace()
	_ = lmr.IsLeftWhitespace()
	_ = lmr.IsRightWhitespace()
	_ = lmr.HasValidNonEmptyLeft()
	_ = lmr.HasValidNonEmptyRight()
	_ = lmr.HasValidNonEmptyMiddle()
	_ = lmr.HasValidNonWhitespaceLeft()
	_ = lmr.HasValidNonWhitespaceRight()
	_ = lmr.HasValidNonWhitespaceMiddle()
	_ = lmr.HasSafeNonEmpty()
	_ = lmr.IsAll("a", "b", "c")
	_ = lmr.Is("a", "b")
	_ = lmr.Clone()
	_ = lmr.ToLeftRight()
	lmr.Clear()
	lmr2 := corestr.NewLeftMiddleRight("a", "b", "c")
	lmr2.Dispose()
}

func Test_C34_LMRFS_Methods(t *testing.T) {
	_ = corestr.LeftMiddleRightFromSplit("a:b:c", ":")
	_ = corestr.LeftMiddleRightFromSplitTrimmed("a : b : c", ":")
	_ = corestr.LeftMiddleRightFromSplitN("a:b:c", ":", 3)
	_ = corestr.LeftMiddleRightFromSplitNTrimmed("a : b : c", ":", 3)
}

// ── ValidValue ──

func Test_C34_VV_Methods(t *testing.T) {
	vv := corestr.NewValidValue("hello")
	_ = vv.IsEmpty()
	_ = vv.ValueBytesOnce()
	_ = vv.ValueBytesOncePtr()
	_ = vv.Clone()
	_ = vv.String()
	vv.Clear()
	vv2 := corestr.NewValidValue("x")
	vv2.Dispose()
}

func Test_C34_VV_Creators(t *testing.T) {
	_ = corestr.NewValidValue("x")
	_ = corestr.NewValidValueEmpty()
	_ = corestr.InvalidValidValue("msg")
	_ = corestr.InvalidValidValueNoMessage()
	_ = corestr.NewValidValueUsingAny(false, true, "x")
	_ = corestr.NewValidValueUsingAnyAutoValid(false, "x")
}

func Test_C34_VV_JsonMethods(t *testing.T) {
	vv := corestr.NewValidValue("x")
	_ = vv.Json()
	_ = vv.JsonPtr()
	_ = vv.JsonModel()
	_ = vv.JsonModelAny()
	_, _ = vv.MarshalJSON()
	_, _ = vv.Serialize()
	_ = vv.AsJsonContractsBinder()
	_ = vv.AsJsoner()
	_ = vv.AsJsonMarshaller()
	_ = vv.AsJsonParseSelfInjector()
}

func Test_C34_VV_Boolean(t *testing.T) {
	_ = corestr.NewValidValue("true").Boolean()
}

func Test_C34_VV_Integer(t *testing.T) {
	_ = corestr.NewValidValue("42").Integer()
}

func Test_C34_VV_Float(t *testing.T) {
	_ = corestr.NewValidValue("3.14").Float()
}

func Test_C34_VV_Float64(t *testing.T) {
	_ = corestr.NewValidValue("3.14").Float64()
}

func Test_C34_VV_IsWhitespace(t *testing.T) {
	_ = corestr.NewValidValue("  ").IsWhitespace()
}

func Test_C34_VV_HasNonEmpty(t *testing.T) {
	_ = corestr.NewValidValue("x").HasNonEmpty()
}

func Test_C34_VV_Trim(t *testing.T) {
	_ = corestr.NewValidValue(" x ").Trim()
}

func Test_C34_VV_HasPrefix(t *testing.T) {
	_ = corestr.NewValidValue("hello").HasPrefix("hel")
}

func Test_C34_VV_HasSuffix(t *testing.T) {
	_ = corestr.NewValidValue("hello").HasSuffix("llo")
}

func Test_C34_VV_Contains(t *testing.T) {
	_ = corestr.NewValidValue("hello").Contains("ell")
}

func Test_C34_VV_IsMatchRegex(t *testing.T) {
	_ = corestr.NewValidValue("hello123").IsMatchRegex("[0-9]+")
}

func Test_C34_VV_ToLower(t *testing.T) {
	_ = corestr.NewValidValue("HELLO").ToLower()
}

func Test_C34_VV_ToUpper(t *testing.T) {
	_ = corestr.NewValidValue("hello").ToUpper()
}

func Test_C34_VV_ValueLength(t *testing.T) {
	_ = corestr.NewValidValue("hello").ValueLength()
}

// ── ValidValues ──

func Test_C34_VVS_Methods(t *testing.T) {
	vvs := corestr.NewValidValues(5)
	vvs.Add(corestr.NewValidValue("a"))
	_ = vvs.Length()
	_ = vvs.Count()
	_ = vvs.HasAnyItem()
	_ = vvs.LastIndex()
	_ = vvs.HasIndex(0)
	_ = vvs.IsEmpty()
	_ = vvs.SafeValueAt(0)
	_ = vvs.SafeValueAt(99)
}

func Test_C34_VVS_Creators(t *testing.T) {
	_ = corestr.EmptyValidValues()
	_ = corestr.NewValidValues(5)
	_ = corestr.NewValidValuesUsingValues(corestr.ValidValue{Value: "a", IsValid: true})
	_ = corestr.NewValidValuesUsingValues()
}

func Test_C34_VVS_AddFull(t *testing.T) {
	vvs := corestr.NewValidValues(5)
	vvs.AddFull("v", true, "")
}

func Test_C34_VVS_AddFullIf(t *testing.T) {
	vvs := corestr.NewValidValues(5)
	vvs.AddFullIf(true, "v", true, "")
	vvs.AddFullIf(false, "v2", true, "")
}

func Test_C34_VVS_AddItems(t *testing.T) {
	vvs := corestr.NewValidValues(5)
	vvs.AddItems(corestr.NewValidValue("a"), corestr.NewValidValue("b"))
}

func Test_C34_VVS_Find(t *testing.T) {
	vvs := corestr.NewValidValues(5)
	vvs.Add(corestr.NewValidValue("a"))
	_ = vvs.Find(func(i int, v *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
		return v, true, false
	})
}

func Test_C34_VVS_List(t *testing.T) {
	vvs := corestr.NewValidValues(5)
	vvs.Add(corestr.NewValidValue("a"))
	_ = vvs.List()
}

func Test_C34_VVS_Strings(t *testing.T) {
	vvs := corestr.NewValidValues(5)
	vvs.Add(corestr.NewValidValue("a"))
	_ = vvs.Strings()
}

func Test_C34_VVS_StringsUsingConditional(t *testing.T) {
	vvs := corestr.NewValidValues(5)
	vvs.Add(corestr.NewValidValue("a"))
	_ = vvs.StringsUsingConditional(func(v *corestr.ValidValue) bool { return true })
}

func Test_C34_VVS_Clear(t *testing.T) {
	vvs := corestr.NewValidValues(5)
	vvs.Clear()
}

func Test_C34_VVS_Dispose(t *testing.T) {
	vvs := corestr.NewValidValues(5)
	vvs.Dispose()
}

func Test_C34_VVS_JsonMethods(t *testing.T) {
	vvs := corestr.NewValidValues(5)
	vvs.Add(corestr.NewValidValue("a"))
	_ = vvs.Json()
	_ = vvs.JsonPtr()
	_ = vvs.JsonModel()
	_ = vvs.JsonModelAny()
}

// ── ValueStatus ──

func Test_C34_VS_Methods(t *testing.T) {
	_ = corestr.InvalidValueStatus("msg")
	_ = corestr.InvalidValueStatusNoMessage()
	vs := &corestr.ValueStatus{ValueValid: corestr.NewValidValue("x"), Index: 0}
	_ = vs.Clone()
}

// ── TextWithLineNumber ──

func Test_C34_TWLN_Methods(t *testing.T) {
	tw := &corestr.TextWithLineNumber{LineNumber: 1, Text: "hello"}
	_ = tw.HasLineNumber()
	_ = tw.IsInvalidLineNumber()
	_ = tw.Length()
	_ = tw.IsEmpty()
	_ = tw.IsEmptyText()
	_ = tw.IsEmptyTextLineBoth()
	var nilTw *corestr.TextWithLineNumber
	_ = nilTw.Length()
	_ = nilTw.IsEmpty()
	_ = nilTw.IsEmptyText()
	_ = nilTw.HasLineNumber()
	_ = nilTw.IsInvalidLineNumber()
}

// ── utils ──

func Test_C34_Utils_WrapDouble(t *testing.T)       { _ = corestr.StringUtils.WrapDouble("x") }
func Test_C34_Utils_WrapSingle(t *testing.T)       { _ = corestr.StringUtils.WrapSingle("x") }
func Test_C34_Utils_WrapTilda(t *testing.T)        { _ = corestr.StringUtils.WrapTilda("x") }
func Test_C34_Utils_WrapDoubleIfMissing(t *testing.T) {
	_ = corestr.StringUtils.WrapDoubleIfMissing("x")
	_ = corestr.StringUtils.WrapDoubleIfMissing(`"x"`)
	_ = corestr.StringUtils.WrapDoubleIfMissing("")
}
func Test_C34_Utils_WrapSingleIfMissing(t *testing.T) {
	_ = corestr.StringUtils.WrapSingleIfMissing("x")
	_ = corestr.StringUtils.WrapSingleIfMissing("'x'")
	_ = corestr.StringUtils.WrapSingleIfMissing("")
}

// ── CloneSlice / CloneSliceIf ──

func Test_C34_CloneSlice(t *testing.T) {
	_ = corestr.CloneSlice([]string{"a"})
	_ = corestr.CloneSlice(nil)
}

func Test_C34_CloneSliceIf(t *testing.T) {
	_ = corestr.CloneSliceIf(true, "a")
	_ = corestr.CloneSliceIf(false, "a")
	_ = corestr.CloneSliceIf(true)
}

// ── AnyToString ──

func Test_C34_AnyToString(t *testing.T) {
	_ = corestr.AnyToString(false, "hello")
	_ = corestr.AnyToString(true, "hello")
	_ = corestr.AnyToString(false, "")
}

// ── AllIndividualStringsOfStringsLength ──

func Test_C34_AllIndividualStringsOfStringsLength(t *testing.T) {
	s := [][]string{{"a", "b"}, {"c"}}
	_ = corestr.AllIndividualStringsOfStringsLength(&s)
	_ = corestr.AllIndividualStringsOfStringsLength(nil)
}

// ── AllIndividualsLengthOfSimpleSlices ──

func Test_C34_AllIndividualsLengthOfSimpleSlices(t *testing.T) {
	ss := corestr.New.SimpleSlice.Lines("a", "b")
	_ = corestr.AllIndividualsLengthOfSimpleSlices(ss)
	_ = corestr.AllIndividualsLengthOfSimpleSlices()
}

// ── NonChainedLinkedListNodes ──

func Test_C34_NCLLN_Methods(t *testing.T) {
	nc := corestr.NewNonChainedLinkedListNodes(5)
	if !nc.IsEmpty() { t.Fatal("expected empty") }
	ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
	nc.Adds(ll.Head())
	_ = nc.Length()
	_ = nc.HasItems()
	_ = nc.First()
	_ = nc.FirstOrDefault()
	_ = nc.Last()
	_ = nc.LastOrDefault()
	_ = nc.Items()
	_ = nc.IsChainingApplied()
	nc.ApplyChaining()
	_ = nc.ToChainedNodes()
}

func Test_C34_NCLLN_Empty(t *testing.T) {
	nc := corestr.NewNonChainedLinkedListNodes(0)
	_ = nc.FirstOrDefault()
	_ = nc.LastOrDefault()
	nc.ApplyChaining()
}

// ── NonChainedLinkedCollectionNodes ──

func Test_C34_NCLCN_Methods(t *testing.T) {
	nc := corestr.NewNonChainedLinkedCollectionNodes(5)
	if !nc.IsEmpty() { t.Fatal("expected empty") }
	lc := corestr.New.LinkedCollection.Strings([]string{"a"})
	nc.Adds(lc.Head())
	_ = nc.Length()
	_ = nc.HasItems()
	_ = nc.First()
	_ = nc.FirstOrDefault()
	_ = nc.Last()
	_ = nc.LastOrDefault()
	_ = nc.Items()
	_ = nc.IsChainingApplied()
	nc.ApplyChaining()
	_ = nc.ToChainedNodes()
}

func Test_C34_NCLCN_Empty(t *testing.T) {
	nc := corestr.NewNonChainedLinkedCollectionNodes(0)
	_ = nc.FirstOrDefault()
	_ = nc.LastOrDefault()
	nc.ApplyChaining()
}

// ── CollectionsOfCollection ──

func Test_C34_COC_Methods(t *testing.T) {
	coc := corestr.New.CollectionsOfCollection.Empty()
	_ = coc.IsEmpty()
	_ = coc.HasItems()
	_ = coc.Length()
	_ = coc.AllIndividualItemsLength()
	_ = coc.Items()
}

func Test_C34_COC_Add(t *testing.T) {
	coc := corestr.New.CollectionsOfCollection.Empty()
	coc.Add(corestr.New.Collection.Strings([]string{"a"}))
	coc.Add(nil)
}

func Test_C34_COC_AddStrings(t *testing.T) {
	coc := corestr.New.CollectionsOfCollection.Empty()
	coc.AddStrings(false, []string{"a"})
	coc.AddStrings(false, nil)
}

func Test_C34_COC_AddsStringsOfStrings(t *testing.T) {
	coc := corestr.New.CollectionsOfCollection.Empty()
	coc.AddsStringsOfStrings(false, []string{"a"}, []string{"b"})
	coc.AddsStringsOfStrings(false)
}

func Test_C34_COC_AddAsyncFuncItems(t *testing.T) {
	coc := corestr.New.CollectionsOfCollection.Empty()
	coc.AddAsyncFuncItems(func() *corestr.Collection { return corestr.New.Collection.Strings([]string{"a"}) })
	coc.AddAsyncFuncItems()
}

func Test_C34_COC_Adds(t *testing.T) {
	coc := corestr.New.CollectionsOfCollection.Empty()
	coc.Adds(*corestr.New.Collection.Strings([]string{"a"}))
	coc.Adds()
}

func Test_C34_COC_AddCollections(t *testing.T) {
	coc := corestr.New.CollectionsOfCollection.Empty()
	coc.AddCollections(corestr.New.Collection.Strings([]string{"a"}))
	coc.AddCollections()
}

func Test_C34_COC_List(t *testing.T) {
	coc := corestr.New.CollectionsOfCollection.Strings([]string{"a"})
	_ = coc.List(0)
}

func Test_C34_COC_ToCollection(t *testing.T) {
	coc := corestr.New.CollectionsOfCollection.Strings([]string{"a"})
	_ = coc.ToCollection()
}

func Test_C34_COC_String(t *testing.T) {
	coc := corestr.New.CollectionsOfCollection.Strings([]string{"a"})
	_ = coc.String()
	_ = corestr.New.CollectionsOfCollection.Empty().String()
}

func Test_C34_COC_JsonMethods(t *testing.T) {
	coc := corestr.New.CollectionsOfCollection.Strings([]string{"a"})
	_ = coc.Json()
	_ = coc.JsonPtr()
	_ = coc.JsonModel()
	_ = coc.JsonModelAny()
	_, _ = coc.MarshalJSON()
	_ = coc.AsJsonContractsBinder()
	_ = coc.AsJsoner()
	_ = coc.AsJsonMarshaller()
	_ = coc.AsJsonParseSelfInjector()
}

func Test_C34_COC_UnmarshalJSON(t *testing.T) {
	coc := &corestr.CollectionsOfCollection{}
	_ = coc.UnmarshalJSON([]byte(`[["a"]]`))
}

func Test_C34_COC_ParseInjectUsingJson(t *testing.T) {
	coc := corestr.New.CollectionsOfCollection.Empty()
	r := corejson.New([][]string{{"a"}})
	_, _ = coc.ParseInjectUsingJson(&r)
}

func Test_C34_COC_ParseInjectUsingJsonMust_Panic(t *testing.T) {
	defer func() { recover() }()
	coc := corestr.New.CollectionsOfCollection.Empty()
	bad := corejson.NewResult.UsingString(`invalid`)
	coc.ParseInjectUsingJsonMust(bad)
}

func Test_C34_COC_JsonParseSelfInject(t *testing.T) {
	coc := corestr.New.CollectionsOfCollection.Empty()
	r := corejson.New([][]string{{"a"}})
	_ = coc.JsonParseSelfInject(&r)
}

// ── newCollectionsOfCollectionCreator ──

func Test_C34_NCOCC_Empty(t *testing.T)    { _ = corestr.New.CollectionsOfCollection.Empty() }
func Test_C34_NCOCC_Cap(t *testing.T)      { _ = corestr.New.CollectionsOfCollection.Cap(5) }
func Test_C34_NCOCC_LenCap(t *testing.T)   { _ = corestr.New.CollectionsOfCollection.LenCap(0, 5) }
func Test_C34_NCOCC_Strings(t *testing.T)  { _ = corestr.New.CollectionsOfCollection.Strings([]string{"a"}) }
func Test_C34_NCOCC_SpreadStrings(t *testing.T) {
	_ = corestr.New.CollectionsOfCollection.SpreadStrings("a")
	_ = corestr.New.CollectionsOfCollection.SpreadStrings()
}
func Test_C34_NCOCC_CloneStrings(t *testing.T) {
	_ = corestr.New.CollectionsOfCollection.CloneStrings([]string{"a"})
}
func Test_C34_NCOCC_UsingCollections(t *testing.T) {
	_ = corestr.New.CollectionsOfCollection.UsingCollections(corestr.New.Collection.Strings([]string{"a"}))
}
func Test_C34_NCOCC_UsingStringsOfStrings(t *testing.T) {
	_ = corestr.New.CollectionsOfCollection.UsingStringsOfStrings(false, []string{"a"})
}

// ── HashsetsCollection ──

func Test_C34_HC_Methods(t *testing.T) {
	hc := corestr.New.HashsetsCollection.Empty()
	_ = hc.IsEmpty()
	_ = hc.HasItems()
	_ = hc.Length()
	hc.Add(corestr.New.Hashset.StringsSpreadItems("a"))
	_ = hc.IndexOf(0)
	_ = hc.ListPtr()
	_ = hc.List()
	_ = hc.StringsList()
	_ = hc.HasAll("a")
	_ = hc.ListDirectPtr()
	_ = hc.SortedListAsc()
	_ = hc.Json()
	_ = hc.JsonPtr()
	_ = hc.JsonModel()
	_ = hc.JsonModelAny()
	_, _ = hc.MarshalJSON()
	_ = hc.AsJsonContractsBinder()
	_ = hc.AsJsoner()
	_ = hc.AsJsonMarshaller()
	_ = hc.AsJsonParseSelfInjector()
	_ = hc.String()
	hc.Clear()
	hc2 := corestr.New.HashsetsCollection.Empty()
	hc2.Dispose()
}

func Test_C34_HC_AddNonEmpty(t *testing.T) {
	hc := corestr.New.HashsetsCollection.Empty()
	hc.AddNonEmpty(corestr.New.Hashset.Empty())
	hc.AddNonEmpty(corestr.New.Hashset.StringsSpreadItems("a"))
}

func Test_C34_HC_AddNonNil(t *testing.T) {
	hc := corestr.New.HashsetsCollection.Empty()
	hc.AddNonNil(nil)
	hc.AddNonNil(corestr.New.Hashset.StringsSpreadItems("a"))
}

func Test_C34_HC_Adds(t *testing.T) {
	hc := corestr.New.HashsetsCollection.Empty()
	hc.Adds(corestr.New.Hashset.StringsSpreadItems("a"))
}

func Test_C34_HC_AddHashsetsCollection(t *testing.T) {
	hc := corestr.New.HashsetsCollection.Empty()
	other := corestr.New.HashsetsCollection.Empty()
	other.Add(corestr.New.Hashset.StringsSpreadItems("a"))
	hc.AddHashsetsCollection(other)
}

func Test_C34_HC_UnmarshalJSON(t *testing.T) {
	hc := &corestr.HashsetsCollection{}
	_ = hc.UnmarshalJSON([]byte(`[["a"]]`))
}

// ── newHashsetsCollectionCreator ──

func Test_C34_NHCC_Empty(t *testing.T)  { _ = corestr.New.HashsetsCollection.Empty() }
func Test_C34_NHCC_Cap(t *testing.T)    { _ = corestr.New.HashsetsCollection.Cap(5) }
func Test_C34_NHCC_LenCap(t *testing.T) { _ = corestr.New.HashsetsCollection.LenCap(0, 5) }
func Test_C34_NHCC_UsingHashsets(t *testing.T) {
	_ = corestr.New.HashsetsCollection.UsingHashsets(corestr.New.Hashset.StringsSpreadItems("a"))
}
func Test_C34_NHCC_UsingHashsetsPointers(t *testing.T) {
	h := corestr.New.Hashset.StringsSpreadItems("a")
	_ = corestr.New.HashsetsCollection.UsingHashsetsPointers(h)
}

// ── emptyCreator ──

func Test_C34_EC_All(t *testing.T) {
	_ = corestr.Empty.Collection()
	_ = corestr.Empty.LinkedList()
	_ = corestr.Empty.SimpleSlice()
	_ = corestr.Empty.KeyAnyValuePair()
	_ = corestr.Empty.KeyValuePair()
	_ = corestr.Empty.KeyValueCollection()
	_ = corestr.Empty.LinkedCollections()
	_ = corestr.Empty.LeftRight()
	_ = corestr.Empty.SimpleStringOnce()
	_ = corestr.Empty.SimpleStringOncePtr()
	_ = corestr.Empty.Hashset()
	_ = corestr.Empty.HashsetsCollection()
	_ = corestr.Empty.Hashmap()
	_ = corestr.Empty.CharCollectionMap()
	_ = corestr.Empty.KeyValuesCollection()
	_ = corestr.Empty.CollectionsOfCollection()
	_ = corestr.Empty.CharHashsetMap()
}

// ── CharCollectionMap / CharHashsetMap ──

func Test_C34_CCM_Basic(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Empty()
	ccm.AddStrings("apple", "banana")
	_ = ccm.Length()
	_ = ccm.IsEmpty()
	_ = ccm.HasItems()
	_ = ccm.Has('a')
	_ = ccm.GetMap()
	_ = ccm.List()
	_ = ccm.SortedListAsc()
	_ = ccm.String()
	_ = ccm.SummaryString()
	_ = ccm.Json()
	_ = ccm.JsonPtr()
	_ = ccm.AllLengthsSum()
	ccm.Clear()
	ccm.Dispose()
}

func Test_C34_CHM_Basic(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(5)
	chm.AddStrings("apple", "banana")
	_ = chm.Length()
	_ = chm.IsEmpty()
	_ = chm.HasItems()
	_ = chm.Has('a')
	_ = chm.GetMap()
	_ = chm.List()
	_ = chm.SortedListAsc()
	_ = chm.String()
	_ = chm.SummaryString()
	_ = chm.Json()
	_ = chm.JsonPtr()
	_ = chm.AllLengthsSum()
	chm.Clear()
}

func Test_C34_NCCMC_CapSelfCap(t *testing.T) {
	_ = corestr.New.CharCollectionMap.CapSelfCap(5, 3)
}

func Test_C34_NCCMC_Items(t *testing.T) {
	_ = corestr.New.CharCollectionMap.Items(map[byte]*corestr.Collection{})
}

func Test_C34_NCCMC_ItemsPtrWithCap(t *testing.T) {
	_ = corestr.New.CharCollectionMap.ItemsPtrWithCap(5, map[byte]*corestr.Collection{})
}

func Test_C34_NCHMC_CapItems(t *testing.T) {
	_ = corestr.New.CharHashsetMap.CapItems(5, map[byte]*corestr.Hashset{})
}

func Test_C34_NCHMC_Strings(t *testing.T) {
	_ = corestr.New.CharHashsetMap.Strings("a", "b")
}

// ── DataModel conversions ──

func Test_C34_HashmapDataModel(t *testing.T) {
	hm := corestr.New.Hashmap.Empty()
	hm.AddOrUpdate("k", "v")
	dm := corestr.NewHashmapsDataModelUsing(hm)
	_ = corestr.NewHashmapUsingDataModel(dm)
}

func Test_C34_HashsetDataModel(t *testing.T) {
	hs := corestr.New.Hashset.StringsSpreadItems("a")
	dm := corestr.NewHashsetsDataModelUsing(hs)
	_ = corestr.NewHashsetUsingDataModel(dm)
}

func Test_C34_CharCollectionDataModel(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Empty()
	ccm.AddStrings("apple")
	dm := corestr.NewCharCollectionMapDataModelUsing(ccm)
	_ = corestr.NewCharCollectionMapUsingDataModel(dm)
}

func Test_C34_CharHashsetDataModel(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(5)
	chm.AddStrings("apple")
	dm := corestr.NewCharHashsetMapDataModelUsing(chm)
	_ = corestr.NewCharHashsetMapUsingDataModel(dm)
}

func Test_C34_HashsetsCollectionDataModel(t *testing.T) {
	hc := corestr.New.HashsetsCollection.Empty()
	hc.Add(corestr.New.Hashset.StringsSpreadItems("a"))
	dm := corestr.NewHashsetsCollectionDataModelUsing(hc)
	_ = corestr.NewHashsetsCollectionUsingDataModel(dm)
}
