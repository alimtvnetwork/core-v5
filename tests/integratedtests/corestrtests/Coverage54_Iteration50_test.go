package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ═══════════════════════════════════════════════════════════════
// HashmapDiff
// ═══════════════════════════════════════════════════════════════

func Test_Cov54_HashmapDiff_Length(t *testing.T) {
	hd := corestr.HashmapDiff(map[string]string{"a": "1"})
	tc := coretestcases.CaseV1{Name: "HD Length", Expected: 1, Actual: hd.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_HashmapDiff_IsEmpty(t *testing.T) {
	hd := corestr.HashmapDiff(map[string]string{})
	tc := coretestcases.CaseV1{Name: "HD IsEmpty", Expected: true, Actual: hd.IsEmpty(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_HashmapDiff_HasAnyItem(t *testing.T) {
	hd := corestr.HashmapDiff(map[string]string{"a": "1"})
	tc := coretestcases.CaseV1{Name: "HD HasAnyItem", Expected: true, Actual: hd.HasAnyItem(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_HashmapDiff_LastIndex(t *testing.T) {
	hd := corestr.HashmapDiff(map[string]string{"a": "1", "b": "2"})
	tc := coretestcases.CaseV1{Name: "HD LastIndex", Expected: 1, Actual: hd.LastIndex(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_HashmapDiff_Raw(t *testing.T) {
	hd := corestr.HashmapDiff(map[string]string{"a": "1"})
	tc := coretestcases.CaseV1{Name: "HD Raw", Expected: "1", Actual: hd.Raw()["a"], Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_HashmapDiff_Raw_Nil(t *testing.T) {
	var hd *corestr.HashmapDiff
	tc := coretestcases.CaseV1{Name: "HD Raw nil", Expected: 0, Actual: len(hd.Raw()), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_HashmapDiff_AllKeysSorted(t *testing.T) {
	hd := corestr.HashmapDiff(map[string]string{"b": "2", "a": "1"})
	keys := hd.AllKeysSorted()
	tc := coretestcases.CaseV1{Name: "HD AllKeysSorted", Expected: "a", Actual: keys[0], Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_HashmapDiff_MapAnyItems(t *testing.T) {
	hd := corestr.HashmapDiff(map[string]string{"a": "1"})
	m := hd.MapAnyItems()
	tc := coretestcases.CaseV1{Name: "HD MapAnyItems", Expected: "1", Actual: m["a"], Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_HashmapDiff_MapAnyItems_Nil(t *testing.T) {
	var hd *corestr.HashmapDiff
	m := hd.MapAnyItems()
	tc := coretestcases.CaseV1{Name: "HD MapAnyItems nil", Expected: 0, Actual: len(m), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_HashmapDiff_IsRawEqual(t *testing.T) {
	hd := corestr.HashmapDiff(map[string]string{"a": "1"})
	tc := coretestcases.CaseV1{Name: "HD IsRawEqual", Expected: true, Actual: hd.IsRawEqual(map[string]string{"a": "1"}), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_HashmapDiff_IsRawEqual_Diff(t *testing.T) {
	hd := corestr.HashmapDiff(map[string]string{"a": "1"})
	tc := coretestcases.CaseV1{Name: "HD IsRawEqual diff", Expected: false, Actual: hd.IsRawEqual(map[string]string{"a": "2"}), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_HashmapDiff_HasAnyChanges(t *testing.T) {
	hd := corestr.HashmapDiff(map[string]string{"a": "1"})
	tc := coretestcases.CaseV1{Name: "HD HasAnyChanges", Expected: true, Actual: hd.HasAnyChanges(map[string]string{"a": "2"}), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_HashmapDiff_DiffRaw(t *testing.T) {
	hd := corestr.HashmapDiff(map[string]string{"a": "1"})
	diff := hd.DiffRaw(map[string]string{"a": "2"})
	tc := coretestcases.CaseV1{Name: "HD DiffRaw", Expected: true, Actual: len(diff) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_HashmapDiff_HashmapDiffUsingRaw(t *testing.T) {
	hd := corestr.HashmapDiff(map[string]string{"a": "1"})
	result := hd.HashmapDiffUsingRaw(map[string]string{"a": "2"})
	tc := coretestcases.CaseV1{Name: "HD HashmapDiffUsingRaw", Expected: true, Actual: result.HasAnyItem(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_HashmapDiff_HashmapDiffUsingRaw_Same(t *testing.T) {
	hd := corestr.HashmapDiff(map[string]string{"a": "1"})
	result := hd.HashmapDiffUsingRaw(map[string]string{"a": "1"})
	tc := coretestcases.CaseV1{Name: "HD HashmapDiffUsingRaw same", Expected: true, Actual: result.IsEmpty(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_HashmapDiff_DiffJsonMessage(t *testing.T) {
	hd := corestr.HashmapDiff(map[string]string{"a": "1"})
	msg := hd.DiffJsonMessage(map[string]string{"a": "2"})
	tc := coretestcases.CaseV1{Name: "HD DiffJsonMessage", Expected: true, Actual: len(msg) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_HashmapDiff_ToStringsSliceOfDiffMap(t *testing.T) {
	hd := corestr.HashmapDiff(map[string]string{"a": "1"})
	diff := hd.DiffRaw(map[string]string{"a": "2"})
	slice := hd.ToStringsSliceOfDiffMap(diff)
	tc := coretestcases.CaseV1{Name: "HD ToStringsSlice", Expected: true, Actual: len(slice) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_HashmapDiff_ShouldDiffMessage(t *testing.T) {
	hd := corestr.HashmapDiff(map[string]string{"a": "1"})
	msg := hd.ShouldDiffMessage("test", map[string]string{"a": "2"})
	tc := coretestcases.CaseV1{Name: "HD ShouldDiffMessage", Expected: true, Actual: len(msg) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_HashmapDiff_Serialize(t *testing.T) {
	hd := corestr.HashmapDiff(map[string]string{"a": "1"})
	data, err := hd.Serialize()
	tc := coretestcases.CaseV1{Name: "HD Serialize", Expected: true, Actual: err == nil && len(data) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_HashmapDiff_RawMapStringAnyDiff(t *testing.T) {
	hd := corestr.HashmapDiff(map[string]string{"a": "1"})
	m := hd.RawMapStringAnyDiff()
	tc := coretestcases.CaseV1{Name: "HD RawMapStringAnyDiff", Expected: true, Actual: len(m) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// KeyAnyValuePair
// ═══════════════════════════════════════════════════════════════

func Test_Cov54_KeyAnyValuePair_KeyName(t *testing.T) {
	kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
	tc := coretestcases.CaseV1{Name: "KAV KeyName", Expected: "k", Actual: kav.KeyName(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_KeyAnyValuePair_VariableName(t *testing.T) {
	kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
	tc := coretestcases.CaseV1{Name: "KAV VariableName", Expected: "k", Actual: kav.VariableName(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_KeyAnyValuePair_ValueAny(t *testing.T) {
	kav := &corestr.KeyAnyValuePair{Key: "k", Value: 42}
	tc := coretestcases.CaseV1{Name: "KAV ValueAny", Expected: 42, Actual: kav.ValueAny(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_KeyAnyValuePair_IsVariableNameEqual(t *testing.T) {
	kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
	tc := coretestcases.CaseV1{Name: "KAV IsVarNameEqual", Expected: true, Actual: kav.IsVariableNameEqual("k"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_KeyAnyValuePair_IsValueNull(t *testing.T) {
	kav := &corestr.KeyAnyValuePair{Key: "k", Value: nil}
	tc := coretestcases.CaseV1{Name: "KAV IsValueNull", Expected: true, Actual: kav.IsValueNull(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_KeyAnyValuePair_HasNonNull(t *testing.T) {
	kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
	tc := coretestcases.CaseV1{Name: "KAV HasNonNull", Expected: true, Actual: kav.HasNonNull(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_KeyAnyValuePair_HasValue(t *testing.T) {
	kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
	tc := coretestcases.CaseV1{Name: "KAV HasValue", Expected: true, Actual: kav.HasValue(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_KeyAnyValuePair_IsValueEmptyString(t *testing.T) {
	kav := &corestr.KeyAnyValuePair{Key: "k", Value: nil}
	tc := coretestcases.CaseV1{Name: "KAV IsValueEmptyString", Expected: true, Actual: kav.IsValueEmptyString(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_KeyAnyValuePair_ValueString(t *testing.T) {
	kav := &corestr.KeyAnyValuePair{Key: "k", Value: "hello"}
	tc := coretestcases.CaseV1{Name: "KAV ValueString", Expected: "hello", Actual: kav.ValueString(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_KeyAnyValuePair_String(t *testing.T) {
	kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
	tc := coretestcases.CaseV1{Name: "KAV String", Expected: true, Actual: len(kav.String()) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_KeyAnyValuePair_Compile(t *testing.T) {
	kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
	tc := coretestcases.CaseV1{Name: "KAV Compile", Expected: true, Actual: len(kav.Compile()) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_KeyAnyValuePair_Serialize(t *testing.T) {
	kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
	data, err := kav.Serialize()
	tc := coretestcases.CaseV1{Name: "KAV Serialize", Expected: true, Actual: err == nil && len(data) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_KeyAnyValuePair_SerializeMust(t *testing.T) {
	kav := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
	data := kav.SerializeMust()
	tc := coretestcases.CaseV1{Name: "KAV SerializeMust", Expected: true, Actual: len(data) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_KeyAnyValuePair_Json(t *testing.T) {
	kav := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
	j := kav.Json()
	tc := coretestcases.CaseV1{Name: "KAV Json", Expected: true, Actual: j.HasSafeItems(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_KeyAnyValuePair_AsJsoner(t *testing.T) {
	kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
	tc := coretestcases.CaseV1{Name: "KAV AsJsoner", Expected: true, Actual: kav.AsJsoner() != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_KeyAnyValuePair_AsJsonContractsBinder(t *testing.T) {
	kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
	tc := coretestcases.CaseV1{Name: "KAV AsJsonContractsBinder", Expected: true, Actual: kav.AsJsonContractsBinder() != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_KeyAnyValuePair_Clear(t *testing.T) {
	kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
	kav.Clear()
	tc := coretestcases.CaseV1{Name: "KAV Clear", Expected: "", Actual: kav.Key, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_KeyAnyValuePair_Dispose(t *testing.T) {
	kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
	kav.Dispose()
	tc := coretestcases.CaseV1{Name: "KAV Dispose", Expected: "", Actual: kav.Key, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// CollectionsOfCollection
// ═══════════════════════════════════════════════════════════════

func Test_Cov54_CollOfColl_IsEmpty(t *testing.T) {
	coc := corestr.New.CollectionsOfCollection.Cap(5)
	tc := coretestcases.CaseV1{Name: "CoC IsEmpty", Expected: true, Actual: coc.IsEmpty(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_CollOfColl_HasItems(t *testing.T) {
	coc := corestr.New.CollectionsOfCollection.Cap(5)
	col := corestr.New.Collection.Strings("a", "b")
	coc.Adds(*col)
	tc := coretestcases.CaseV1{Name: "CoC HasItems", Expected: true, Actual: coc.HasItems(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_CollOfColl_Length(t *testing.T) {
	coc := corestr.New.CollectionsOfCollection.Cap(5)
	col := corestr.New.Collection.Strings("a")
	coc.Adds(*col)
	tc := coretestcases.CaseV1{Name: "CoC Length", Expected: 1, Actual: coc.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_CollOfColl_AllIndividualItemsLength(t *testing.T) {
	coc := corestr.New.CollectionsOfCollection.Cap(5)
	col := corestr.New.Collection.Strings("a", "b")
	coc.Adds(*col)
	tc := coretestcases.CaseV1{Name: "CoC AllIndivLen", Expected: 2, Actual: coc.AllIndividualItemsLength(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_CollOfColl_Items(t *testing.T) {
	coc := corestr.New.CollectionsOfCollection.Cap(5)
	col := corestr.New.Collection.Strings("a")
	coc.Adds(*col)
	tc := coretestcases.CaseV1{Name: "CoC Items", Expected: 1, Actual: len(coc.Items()), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_CollOfColl_List(t *testing.T) {
	coc := corestr.New.CollectionsOfCollection.Cap(5)
	col := corestr.New.Collection.Strings("a", "b")
	coc.Adds(*col)
	tc := coretestcases.CaseV1{Name: "CoC List", Expected: 2, Actual: len(coc.List(0)), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_CollOfColl_ToCollection(t *testing.T) {
	coc := corestr.New.CollectionsOfCollection.Cap(5)
	col := corestr.New.Collection.Strings("a")
	coc.Adds(*col)
	result := coc.ToCollection()
	tc := coretestcases.CaseV1{Name: "CoC ToCollection", Expected: 1, Actual: result.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_CollOfColl_AddStrings(t *testing.T) {
	coc := corestr.New.CollectionsOfCollection.Cap(5)
	coc.AddStrings(false, []string{"x", "y"})
	tc := coretestcases.CaseV1{Name: "CoC AddStrings", Expected: 1, Actual: coc.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_CollOfColl_AddStrings_Empty(t *testing.T) {
	coc := corestr.New.CollectionsOfCollection.Cap(5)
	coc.AddStrings(false, []string{})
	tc := coretestcases.CaseV1{Name: "CoC AddStrings empty", Expected: 0, Actual: coc.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_CollOfColl_AsJsonContractsBinder(t *testing.T) {
	coc := corestr.New.CollectionsOfCollection.Cap(5)
	tc := coretestcases.CaseV1{Name: "CoC AsJsonContractsBinder", Expected: true, Actual: coc.AsJsonContractsBinder() != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// SimpleStringOnce (core methods)
// ═══════════════════════════════════════════════════════════════

func Test_Cov54_SSO_GetSetOnce(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Value("hello")
	tc := coretestcases.CaseV1{Name: "SSO Value", Expected: "hello", Actual: sso.Value(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_SSO_IsInitialized(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Value("x")
	tc := coretestcases.CaseV1{Name: "SSO IsInitialized", Expected: true, Actual: sso.IsInitialized(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_SSO_IsDefined(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Value("x")
	tc := coretestcases.CaseV1{Name: "SSO IsDefined", Expected: true, Actual: sso.IsDefined(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_SSO_IsUninitialized(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Default()
	tc := coretestcases.CaseV1{Name: "SSO IsUninitialized", Expected: true, Actual: sso.IsUninitialized(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_SSO_Invalidate(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Value("x")
	sso.Invalidate()
	tc := coretestcases.CaseV1{Name: "SSO Invalidate", Expected: true, Actual: sso.IsUninitialized(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_SSO_Reset(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Value("x")
	sso.Reset()
	tc := coretestcases.CaseV1{Name: "SSO Reset", Expected: true, Actual: sso.IsUninitialized(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_SSO_IsInvalid(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Default()
	tc := coretestcases.CaseV1{Name: "SSO IsInvalid", Expected: true, Actual: sso.IsInvalid(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_SSO_ValueBytes(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Value("ab")
	tc := coretestcases.CaseV1{Name: "SSO ValueBytes", Expected: 2, Actual: len(sso.ValueBytes()), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_SSO_SetOnUninitialized(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Default()
	err := sso.SetOnUninitialized("x")
	tc := coretestcases.CaseV1{Name: "SSO SetOnUninitialized", Expected: true, Actual: err == nil && sso.Value() == "x", Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_SSO_SetOnUninitialized_AlreadyInit(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Value("x")
	err := sso.SetOnUninitialized("y")
	tc := coretestcases.CaseV1{Name: "SSO SetOnUninitialized already", Expected: true, Actual: err != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_SSO_GetOnce(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Default()
	val := sso.GetOnce()
	tc := coretestcases.CaseV1{Name: "SSO GetOnce", Expected: "", Actual: val, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_SSO_GetOnceFunc(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Default()
	val := sso.GetOnceFunc(func() string { return "computed" })
	tc := coretestcases.CaseV1{Name: "SSO GetOnceFunc", Expected: "computed", Actual: val, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_SSO_SetOnceIfUninitialized(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Default()
	isSet := sso.SetOnceIfUninitialized("x")
	tc := coretestcases.CaseV1{Name: "SSO SetOnceIfUninitialized", Expected: true, Actual: isSet, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_SSO_IsEmpty(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Default()
	tc := coretestcases.CaseV1{Name: "SSO IsEmpty", Expected: true, Actual: sso.IsEmpty(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_SSO_Is(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Value("x")
	tc := coretestcases.CaseV1{Name: "SSO Is", Expected: true, Actual: sso.Is("x"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_SSO_IsContains(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Value("hello world")
	tc := coretestcases.CaseV1{Name: "SSO IsContains", Expected: true, Actual: sso.IsContains("world"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_SSO_Int(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Value("42")
	tc := coretestcases.CaseV1{Name: "SSO Int", Expected: 42, Actual: sso.Int(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_SSO_Byte(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Value("65")
	tc := coretestcases.CaseV1{Name: "SSO Byte", Expected: byte(65), Actual: sso.Byte(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_SSO_Boolean(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Value("yes")
	tc := coretestcases.CaseV1{Name: "SSO Boolean", Expected: true, Actual: sso.Boolean(false), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_SSO_BooleanDefault(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Value("true")
	tc := coretestcases.CaseV1{Name: "SSO BooleanDefault", Expected: true, Actual: sso.BooleanDefault(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_SSO_ConcatNew(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Value("hello")
	result := sso.ConcatNew(" world")
	tc := coretestcases.CaseV1{Name: "SSO ConcatNew", Expected: "hello world", Actual: result.Value(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_SSO_HasSafeNonEmpty(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Value("x")
	tc := coretestcases.CaseV1{Name: "SSO HasSafeNonEmpty", Expected: true, Actual: sso.HasSafeNonEmpty(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_SSO_SafeValue(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Value("x")
	tc := coretestcases.CaseV1{Name: "SSO SafeValue", Expected: "x", Actual: sso.SafeValue(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_SSO_SafeValue_Uninit(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Default()
	tc := coretestcases.CaseV1{Name: "SSO SafeValue uninit", Expected: "", Actual: sso.SafeValue(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_SSO_ValueInt(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Value("10")
	tc := coretestcases.CaseV1{Name: "SSO ValueInt", Expected: 10, Actual: sso.ValueInt(0), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_SSO_ValueFloat64(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Value("3.14")
	tc := coretestcases.CaseV1{Name: "SSO ValueFloat64", Expected: 3.14, Actual: sso.ValueFloat64(0), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_SSO_WithinRange_InRange(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Value("5")
	val, inRange := sso.WithinRange(true, 0, 10)
	tc := coretestcases.CaseV1{Name: "SSO WithinRange in", Expected: true, Actual: inRange && val == 5, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_SSO_WithinRange_OutOfRange(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Value("20")
	val, inRange := sso.WithinRange(true, 0, 10)
	tc := coretestcases.CaseV1{Name: "SSO WithinRange out", Expected: true, Actual: !inRange && val == 10, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_SSO_Int16(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Value("100")
	tc := coretestcases.CaseV1{Name: "SSO Int16", Expected: int16(100), Actual: sso.Int16(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov54_SSO_Int32(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Value("100")
	tc := coretestcases.CaseV1{Name: "SSO Int32", Expected: int32(100), Actual: sso.Int32(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}
