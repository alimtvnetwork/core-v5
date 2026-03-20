package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// SimpleStringOnce — Init/Set/Get
// ══════════════════════════════════════════════════════════════════════════════

func Test_I28_SSO_Value_Empty(t *testing.T) {
	var sso corestr.SimpleStringOnce
	actual := args.Map{"val": sso.Value(), "init": sso.IsInitialized(), "defined": sso.IsDefined(), "uninit": sso.IsUninitialized()}
	expected := args.Map{"val": "", "init": false, "defined": false, "uninit": true}
	expected.ShouldBeEqual(t, 0, "SSO empty", actual)
}

func Test_I28_SSO_SetOnUninitialized(t *testing.T) {
	var sso corestr.SimpleStringOnce
	err := sso.SetOnUninitialized("hello")
	actual := args.Map{"noErr": err == nil, "val": sso.Value(), "init": sso.IsInitialized()}
	expected := args.Map{"noErr": true, "val": "hello", "init": true}
	expected.ShouldBeEqual(t, 0, "SSO SetOnUninitialized", actual)
}

func Test_I28_SSO_SetOnUninitialized_AlreadyInit(t *testing.T) {
	var sso corestr.SimpleStringOnce
	_ = sso.SetOnUninitialized("first")
	err := sso.SetOnUninitialized("second")
	actual := args.Map{"hasErr": err != nil, "val": sso.Value()}
	expected := args.Map{"hasErr": true, "val": "first"}
	expected.ShouldBeEqual(t, 0, "SSO SetOnUninitialized already init", actual)
}

func Test_I28_SSO_GetSetOnce(t *testing.T) {
	var sso corestr.SimpleStringOnce
	v1 := sso.GetSetOnce("first")
	v2 := sso.GetSetOnce("second")
	actual := args.Map{"v1": v1, "v2": v2}
	expected := args.Map{"v1": "first", "v2": "first"}
	expected.ShouldBeEqual(t, 0, "SSO GetSetOnce", actual)
}

func Test_I28_SSO_GetOnce(t *testing.T) {
	var sso corestr.SimpleStringOnce
	v := sso.GetOnce()
	actual := args.Map{"val": v, "init": sso.IsInitialized()}
	expected := args.Map{"val": "", "init": true}
	expected.ShouldBeEqual(t, 0, "SSO GetOnce", actual)
}

func Test_I28_SSO_GetOnce_AlreadyInit(t *testing.T) {
	var sso corestr.SimpleStringOnce
	sso.GetSetOnce("hello")
	v := sso.GetOnce()
	actual := args.Map{"val": v}
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "SSO GetOnce already init", actual)
}

func Test_I28_SSO_GetOnceFunc(t *testing.T) {
	var sso corestr.SimpleStringOnce
	v := sso.GetOnceFunc(func() string { return "computed" })
	v2 := sso.GetOnceFunc(func() string { return "other" })
	actual := args.Map{"v": v, "v2": v2}
	expected := args.Map{"v": "computed", "v2": "computed"}
	expected.ShouldBeEqual(t, 0, "SSO GetOnceFunc", actual)
}

func Test_I28_SSO_SetOnceIfUninitialized(t *testing.T) {
	var sso corestr.SimpleStringOnce
	ok1 := sso.SetOnceIfUninitialized("hello")
	ok2 := sso.SetOnceIfUninitialized("world")
	actual := args.Map{"ok1": ok1, "ok2": ok2, "val": sso.Value()}
	expected := args.Map{"ok1": true, "ok2": false, "val": "hello"}
	expected.ShouldBeEqual(t, 0, "SSO SetOnceIfUninitialized", actual)
}

func Test_I28_SSO_Invalidate_Reset(t *testing.T) {
	var sso corestr.SimpleStringOnce
	sso.GetSetOnce("hello")
	sso.Invalidate()
	actual := args.Map{"init": sso.IsInitialized(), "val": sso.Value()}
	expected := args.Map{"init": false, "val": ""}
	expected.ShouldBeEqual(t, 0, "SSO Invalidate", actual)

	sso.GetSetOnce("world")
	sso.Reset()
	actual2 := args.Map{"init": sso.IsInitialized()}
	expected2 := args.Map{"init": false}
	expected2.ShouldBeEqual(t, 0, "SSO Reset", actual2)
}

func Test_I28_SSO_IsInvalid(t *testing.T) {
	var sso corestr.SimpleStringOnce
	actual := args.Map{"invalid": sso.IsInvalid()}
	expected := args.Map{"invalid": true}
	expected.ShouldBeEqual(t, 0, "SSO IsInvalid uninit", actual)

	sso.GetSetOnce("hello")
	actual2 := args.Map{"invalid": sso.IsInvalid()}
	expected2 := args.Map{"invalid": false}
	expected2.ShouldBeEqual(t, 0, "SSO IsInvalid init", actual2)
}

func Test_I28_SSO_IsInvalid_Nil(t *testing.T) {
	var sso *corestr.SimpleStringOnce
	actual := args.Map{"invalid": sso.IsInvalid()}
	expected := args.Map{"invalid": true}
	expected.ShouldBeEqual(t, 0, "SSO IsInvalid nil", actual)
}

func Test_I28_SSO_SetInitialize_SetUnInit(t *testing.T) {
	var sso corestr.SimpleStringOnce
	sso.SetInitialize()
	actual := args.Map{"init": sso.IsInitialized()}
	expected := args.Map{"init": true}
	expected.ShouldBeEqual(t, 0, "SSO SetInitialize", actual)

	sso.SetUnInit()
	actual2 := args.Map{"init": sso.IsInitialized()}
	expected2 := args.Map{"init": false}
	expected2.ShouldBeEqual(t, 0, "SSO SetUnInit", actual2)
}

// ══════════════════════════════════════════════════════════════════════════════
// SimpleStringOnce — Bytes, Checks
// ══════════════════════════════════════════════════════════════════════════════

func Test_I28_SSO_ValueBytes(t *testing.T) {
	var sso corestr.SimpleStringOnce
	sso.GetSetOnce("abc")
	actual := args.Map{"len": len(sso.ValueBytes()), "lenPtr": len(sso.ValueBytesPtr())}
	expected := args.Map{"len": 3, "lenPtr": 3}
	expected.ShouldBeEqual(t, 0, "SSO ValueBytes", actual)
}

func Test_I28_SSO_IsEmpty_IsWhitespace_Trim(t *testing.T) {
	var sso corestr.SimpleStringOnce
	sso.GetSetOnce("  hi  ")
	actual := args.Map{"empty": sso.IsEmpty(), "ws": sso.IsWhitespace(), "trim": sso.Trim()}
	expected := args.Map{"empty": false, "ws": false, "trim": "hi"}
	expected.ShouldBeEqual(t, 0, "SSO checks", actual)
}

func Test_I28_SSO_HasValidNonEmpty_HasValidNonWhitespace(t *testing.T) {
	var sso corestr.SimpleStringOnce
	sso.GetSetOnce("x")
	actual := args.Map{"hv": sso.HasValidNonEmpty(), "hvw": sso.HasValidNonWhitespace(), "safe": sso.HasSafeNonEmpty()}
	expected := args.Map{"hv": true, "hvw": true, "safe": true}
	expected.ShouldBeEqual(t, 0, "SSO HasValid", actual)
}

func Test_I28_SSO_SafeValue(t *testing.T) {
	var sso corestr.SimpleStringOnce
	actual := args.Map{"uninit": sso.SafeValue()}
	expected := args.Map{"uninit": ""}
	expected.ShouldBeEqual(t, 0, "SSO SafeValue uninit", actual)

	sso.GetSetOnce("hello")
	actual2 := args.Map{"init": sso.SafeValue()}
	expected2 := args.Map{"init": "hello"}
	expected2.ShouldBeEqual(t, 0, "SSO SafeValue init", actual2)
}

// ══════════════════════════════════════════════════════════════════════════════
// SimpleStringOnce — Numeric conversions
// ══════════════════════════════════════════════════════════════════════════════

func Test_I28_SSO_Int(t *testing.T) {
	var sso corestr.SimpleStringOnce
	sso.GetSetOnce("42")
	actual := args.Map{"val": sso.Int()}
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "SSO Int", actual)
}

func Test_I28_SSO_Int_Err(t *testing.T) {
	var sso corestr.SimpleStringOnce
	sso.GetSetOnce("abc")
	actual := args.Map{"val": sso.Int()}
	expected := args.Map{"val": 0}
	expected.ShouldBeEqual(t, 0, "SSO Int err", actual)
}

func Test_I28_SSO_Byte(t *testing.T) {
	var sso corestr.SimpleStringOnce
	sso.GetSetOnce("100")
	actual := args.Map{"val": sso.Byte()}
	expected := args.Map{"val": byte(100)}
	expected.ShouldBeEqual(t, 0, "SSO Byte", actual)
}

func Test_I28_SSO_Byte_OutOfRange(t *testing.T) {
	var sso corestr.SimpleStringOnce
	sso.GetSetOnce("300")
	actual := args.Map{"val": sso.Byte()}
	expected := args.Map{"val": byte(0)}
	expected.ShouldBeEqual(t, 0, "SSO Byte out of range", actual)
}

func Test_I28_SSO_Byte_Err(t *testing.T) {
	var sso corestr.SimpleStringOnce
	sso.GetSetOnce("abc")
	actual := args.Map{"val": sso.Byte()}
	expected := args.Map{"val": byte(0)}
	expected.ShouldBeEqual(t, 0, "SSO Byte err", actual)
}

func Test_I28_SSO_Int16(t *testing.T) {
	var sso corestr.SimpleStringOnce
	sso.GetSetOnce("100")
	actual := args.Map{"val": sso.Int16()}
	expected := args.Map{"val": int16(100)}
	expected.ShouldBeEqual(t, 0, "SSO Int16", actual)
}

func Test_I28_SSO_Int16_OutOfRange(t *testing.T) {
	var sso corestr.SimpleStringOnce
	sso.GetSetOnce("99999")
	actual := args.Map{"val": sso.Int16()}
	expected := args.Map{"val": int16(0)}
	expected.ShouldBeEqual(t, 0, "SSO Int16 out of range", actual)
}

func Test_I28_SSO_Int16_Err(t *testing.T) {
	var sso corestr.SimpleStringOnce
	sso.GetSetOnce("abc")
	actual := args.Map{"val": sso.Int16()}
	expected := args.Map{"val": int16(0)}
	expected.ShouldBeEqual(t, 0, "SSO Int16 err", actual)
}

func Test_I28_SSO_Int32(t *testing.T) {
	var sso corestr.SimpleStringOnce
	sso.GetSetOnce("1000")
	actual := args.Map{"val": sso.Int32()}
	expected := args.Map{"val": int32(1000)}
	expected.ShouldBeEqual(t, 0, "SSO Int32", actual)
}

func Test_I28_SSO_Int32_Err(t *testing.T) {
	var sso corestr.SimpleStringOnce
	sso.GetSetOnce("abc")
	actual := args.Map{"val": sso.Int32()}
	expected := args.Map{"val": int32(0)}
	expected.ShouldBeEqual(t, 0, "SSO Int32 err", actual)
}

func Test_I28_SSO_Uint16(t *testing.T) {
	var sso corestr.SimpleStringOnce
	sso.GetSetOnce("100")
	val, inRange := sso.Uint16()
	actual := args.Map{"val": val, "inRange": inRange}
	expected := args.Map{"val": uint16(100), "inRange": true}
	expected.ShouldBeEqual(t, 0, "SSO Uint16", actual)
}

func Test_I28_SSO_Uint32(t *testing.T) {
	var sso corestr.SimpleStringOnce
	sso.GetSetOnce("1000")
	val, inRange := sso.Uint32()
	actual := args.Map{"val": val, "inRange": inRange}
	expected := args.Map{"val": uint32(1000), "inRange": true}
	expected.ShouldBeEqual(t, 0, "SSO Uint32", actual)
}

func Test_I28_SSO_WithinRange_InRange(t *testing.T) {
	var sso corestr.SimpleStringOnce
	sso.GetSetOnce("50")
	val, inRange := sso.WithinRange(true, 0, 100)
	actual := args.Map{"val": val, "inRange": inRange}
	expected := args.Map{"val": 50, "inRange": true}
	expected.ShouldBeEqual(t, 0, "SSO WithinRange in range", actual)
}

func Test_I28_SSO_WithinRange_Below(t *testing.T) {
	var sso corestr.SimpleStringOnce
	sso.GetSetOnce("-5")
	val, inRange := sso.WithinRange(true, 0, 100)
	actual := args.Map{"val": val, "inRange": inRange}
	expected := args.Map{"val": 0, "inRange": false}
	expected.ShouldBeEqual(t, 0, "SSO WithinRange below", actual)
}

func Test_I28_SSO_WithinRange_Above(t *testing.T) {
	var sso corestr.SimpleStringOnce
	sso.GetSetOnce("200")
	val, inRange := sso.WithinRange(true, 0, 100)
	actual := args.Map{"val": val, "inRange": inRange}
	expected := args.Map{"val": 100, "inRange": false}
	expected.ShouldBeEqual(t, 0, "SSO WithinRange above", actual)
}

func Test_I28_SSO_WithinRange_NoBoundary(t *testing.T) {
	var sso corestr.SimpleStringOnce
	sso.GetSetOnce("200")
	val, inRange := sso.WithinRange(false, 0, 100)
	actual := args.Map{"val": val, "inRange": inRange}
	expected := args.Map{"val": 200, "inRange": false}
	expected.ShouldBeEqual(t, 0, "SSO WithinRange no boundary", actual)
}

func Test_I28_SSO_WithinRange_Err(t *testing.T) {
	var sso corestr.SimpleStringOnce
	sso.GetSetOnce("abc")
	val, inRange := sso.WithinRange(true, 0, 100)
	actual := args.Map{"val": val, "inRange": inRange}
	expected := args.Map{"val": 0, "inRange": false}
	expected.ShouldBeEqual(t, 0, "SSO WithinRange err", actual)
}

func Test_I28_SSO_WithinRangeDefault(t *testing.T) {
	var sso corestr.SimpleStringOnce
	sso.GetSetOnce("50")
	val, inRange := sso.WithinRangeDefault(0, 100)
	actual := args.Map{"val": val, "inRange": inRange}
	expected := args.Map{"val": 50, "inRange": true}
	expected.ShouldBeEqual(t, 0, "SSO WithinRangeDefault", actual)
}

func Test_I28_SSO_Boolean(t *testing.T) {
	var sso corestr.SimpleStringOnce
	sso.GetSetOnce("yes")
	actual := args.Map{"val": sso.Boolean(false)}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "SSO Boolean yes", actual)
}

func Test_I28_SSO_Boolean_True(t *testing.T) {
	var sso corestr.SimpleStringOnce
	sso.GetSetOnce("true")
	actual := args.Map{"val": sso.Boolean(false)}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "SSO Boolean true", actual)
}

func Test_I28_SSO_Boolean_ConsiderInit_Uninit(t *testing.T) {
	var sso corestr.SimpleStringOnce
	actual := args.Map{"val": sso.Boolean(true)}
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "SSO Boolean consider init uninit", actual)
}

func Test_I28_SSO_Boolean_ParseErr(t *testing.T) {
	var sso corestr.SimpleStringOnce
	sso.GetSetOnce("abc")
	actual := args.Map{"val": sso.Boolean(false)}
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "SSO Boolean parse err", actual)
}

func Test_I28_SSO_BooleanDefault(t *testing.T) {
	var sso corestr.SimpleStringOnce
	sso.GetSetOnce("y")
	actual := args.Map{"val": sso.BooleanDefault()}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "SSO BooleanDefault", actual)
}

func Test_I28_SSO_IsValueBool(t *testing.T) {
	var sso corestr.SimpleStringOnce
	sso.GetSetOnce("1")
	actual := args.Map{"val": sso.IsValueBool()}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "SSO IsValueBool", actual)
}

func Test_I28_SSO_IsSetter(t *testing.T) {
	var sso corestr.SimpleStringOnce
	sso.GetSetOnce("yes")
	is := sso.IsSetter(false)
	actual := args.Map{"true": is.IsTrue()}
	expected := args.Map{"true": true}
	expected.ShouldBeEqual(t, 0, "SSO IsSetter yes", actual)
}

func Test_I28_SSO_IsSetter_ConsiderInit_Uninit(t *testing.T) {
	var sso corestr.SimpleStringOnce
	is := sso.IsSetter(true)
	actual := args.Map{"false": is.IsFalse()}
	expected := args.Map{"false": true}
	expected.ShouldBeEqual(t, 0, "SSO IsSetter uninit", actual)
}

func Test_I28_SSO_IsSetter_ParseErr(t *testing.T) {
	var sso corestr.SimpleStringOnce
	sso.GetSetOnce("abc")
	is := sso.IsSetter(false)
	actual := args.Map{"uninit": is.IsUninitialized()}
	expected := args.Map{"uninit": true}
	expected.ShouldBeEqual(t, 0, "SSO IsSetter parse err", actual)
}

func Test_I28_SSO_ValueInt(t *testing.T) {
	var sso corestr.SimpleStringOnce
	sso.GetSetOnce("42")
	actual := args.Map{"val": sso.ValueInt(0), "defInt": sso.ValueDefInt()}
	expected := args.Map{"val": 42, "defInt": 42}
	expected.ShouldBeEqual(t, 0, "SSO ValueInt", actual)
}

func Test_I28_SSO_ValueInt_Err(t *testing.T) {
	var sso corestr.SimpleStringOnce
	sso.GetSetOnce("abc")
	actual := args.Map{"val": sso.ValueInt(99), "defInt": sso.ValueDefInt()}
	expected := args.Map{"val": 99, "defInt": 0}
	expected.ShouldBeEqual(t, 0, "SSO ValueInt err", actual)
}

func Test_I28_SSO_ValueByte(t *testing.T) {
	var sso corestr.SimpleStringOnce
	sso.GetSetOnce("100")
	actual := args.Map{"val": sso.ValueByte(0), "def": sso.ValueDefByte()}
	expected := args.Map{"val": byte(100), "def": byte(100)}
	expected.ShouldBeEqual(t, 0, "SSO ValueByte", actual)
}

func Test_I28_SSO_ValueFloat64(t *testing.T) {
	var sso corestr.SimpleStringOnce
	sso.GetSetOnce("3.14")
	actual := args.Map{"close": sso.ValueFloat64(0) > 3.1, "def": sso.ValueDefFloat64() > 3.1}
	expected := args.Map{"close": true, "def": true}
	expected.ShouldBeEqual(t, 0, "SSO ValueFloat64", actual)
}

func Test_I28_SSO_NonPtr_Ptr(t *testing.T) {
	var sso corestr.SimpleStringOnce
	sso.GetSetOnce("hello")
	np := sso.NonPtr()
	p := sso.Ptr()
	actual := args.Map{"npVal": np.Value(), "pSame": p == &sso}
	expected := args.Map{"npVal": "hello", "pSame": true}
	expected.ShouldBeEqual(t, 0, "SSO NonPtr/Ptr", actual)
}

func Test_I28_SSO_ConcatNew(t *testing.T) {
	var sso corestr.SimpleStringOnce
	sso.GetSetOnce("hello")
	newSSO := sso.ConcatNew(" world")
	actual := args.Map{"val": newSSO.Value()}
	expected := args.Map{"val": "hello world"}
	expected.ShouldBeEqual(t, 0, "SSO ConcatNew", actual)
}

func Test_I28_SSO_ConcatNewUsingStrings(t *testing.T) {
	var sso corestr.SimpleStringOnce
	sso.GetSetOnce("a")
	newSSO := sso.ConcatNewUsingStrings("-", "b", "c")
	actual := args.Map{"val": newSSO.Value()}
	expected := args.Map{"val": "a-b-c"}
	expected.ShouldBeEqual(t, 0, "SSO ConcatNewUsingStrings", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// HashmapDiff
// ══════════════════════════════════════════════════════════════════════════════

func Test_I28_HashmapDiff_Length(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1", "b": "2"}
	actual := args.Map{"len": hd.Length(), "empty": hd.IsEmpty(), "hasAny": hd.HasAnyItem(), "lastIdx": hd.LastIndex()}
	expected := args.Map{"len": 2, "empty": false, "hasAny": true, "lastIdx": 1}
	expected.ShouldBeEqual(t, 0, "HashmapDiff basics", actual)
}

func Test_I28_HashmapDiff_Nil(t *testing.T) {
	var hd *corestr.HashmapDiff
	actual := args.Map{"len": hd.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "HashmapDiff nil length", actual)
}

func Test_I28_HashmapDiff_AllKeysSorted(t *testing.T) {
	hd := corestr.HashmapDiff{"b": "2", "a": "1"}
	keys := hd.AllKeysSorted()
	actual := args.Map{"first": keys[0], "second": keys[1]}
	expected := args.Map{"first": "a", "second": "b"}
	expected.ShouldBeEqual(t, 0, "HashmapDiff AllKeysSorted", actual)
}

func Test_I28_HashmapDiff_MapAnyItems(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1"}
	mai := hd.MapAnyItems()
	actual := args.Map{"len": len(mai)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "HashmapDiff MapAnyItems", actual)
}

func Test_I28_HashmapDiff_MapAnyItems_Nil(t *testing.T) {
	var hd *corestr.HashmapDiff
	mai := hd.MapAnyItems()
	actual := args.Map{"len": len(mai)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "HashmapDiff MapAnyItems nil", actual)
}

func Test_I28_HashmapDiff_Raw(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1"}
	raw := hd.Raw()
	actual := args.Map{"len": len(raw)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "HashmapDiff Raw", actual)
}

func Test_I28_HashmapDiff_Raw_Nil(t *testing.T) {
	var hd *corestr.HashmapDiff
	raw := hd.Raw()
	actual := args.Map{"len": len(raw)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "HashmapDiff Raw nil", actual)
}

func Test_I28_HashmapDiff_IsRawEqual(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1"}
	actual := args.Map{"eq": hd.IsRawEqual(map[string]string{"a": "1"}), "neq": hd.IsRawEqual(map[string]string{"a": "2"})}
	expected := args.Map{"eq": true, "neq": false}
	expected.ShouldBeEqual(t, 0, "HashmapDiff IsRawEqual", actual)
}

func Test_I28_HashmapDiff_HasAnyChanges(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1"}
	actual := args.Map{"changes": hd.HasAnyChanges(map[string]string{"a": "2"}), "noChanges": hd.HasAnyChanges(map[string]string{"a": "1"})}
	expected := args.Map{"changes": true, "noChanges": false}
	expected.ShouldBeEqual(t, 0, "HashmapDiff HasAnyChanges", actual)
}

func Test_I28_HashmapDiff_DiffRaw(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1", "b": "2"}
	diff := hd.DiffRaw(map[string]string{"a": "1", "b": "99"})
	actual := args.Map{"hasDiff": len(diff) > 0}
	expected := args.Map{"hasDiff": true}
	expected.ShouldBeEqual(t, 0, "HashmapDiff DiffRaw", actual)
}

func Test_I28_HashmapDiff_HashmapDiffUsingRaw_NoDiff(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1"}
	diff := hd.HashmapDiffUsingRaw(map[string]string{"a": "1"})
	actual := args.Map{"empty": diff.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "HashmapDiff HashmapDiffUsingRaw no diff", actual)
}

func Test_I28_HashmapDiff_HashmapDiffUsingRaw_HasDiff(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1"}
	diff := hd.HashmapDiffUsingRaw(map[string]string{"a": "2"})
	actual := args.Map{"hasDiff": diff.HasAnyItem()}
	expected := args.Map{"hasDiff": true}
	expected.ShouldBeEqual(t, 0, "HashmapDiff HashmapDiffUsingRaw has diff", actual)
}

func Test_I28_HashmapDiff_DiffJsonMessage(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1"}
	msg := hd.DiffJsonMessage(map[string]string{"a": "2"})
	actual := args.Map{"notEmpty": msg != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "HashmapDiff DiffJsonMessage", actual)
}

func Test_I28_HashmapDiff_ShouldDiffMessage(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1"}
	msg := hd.ShouldDiffMessage("test", map[string]string{"a": "2"})
	actual := args.Map{"notEmpty": msg != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "HashmapDiff ShouldDiffMessage", actual)
}

func Test_I28_HashmapDiff_LogShouldDiffMessage(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1"}
	msg := hd.LogShouldDiffMessage("test", map[string]string{"a": "2"})
	actual := args.Map{"notEmpty": msg != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "HashmapDiff LogShouldDiffMessage", actual)
}

func Test_I28_HashmapDiff_ToStringsSliceOfDiffMap(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1"}
	diff := hd.DiffRaw(map[string]string{"a": "2"})
	strs := hd.ToStringsSliceOfDiffMap(diff)
	actual := args.Map{"hasItems": len(strs) > 0}
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "HashmapDiff ToStringsSliceOfDiffMap", actual)
}

func Test_I28_HashmapDiff_RawMapStringAnyDiff(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1"}
	d := hd.RawMapStringAnyDiff()
	actual := args.Map{"notNil": d != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "HashmapDiff RawMapStringAnyDiff", actual)
}

func Test_I28_HashmapDiff_Serialize(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1"}
	b, err := hd.Serialize()
	actual := args.Map{"noErr": err == nil, "hasBytes": len(b) > 0}
	expected := args.Map{"noErr": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "HashmapDiff Serialize", actual)
}

func Test_I28_HashmapDiff_Deserialize(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1"}
	target := map[string]string{}
	err := hd.Deserialize(&target)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "HashmapDiff Deserialize", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// HashmapDataModel
// ══════════════════════════════════════════════════════════════════════════════

func Test_I28_HashmapDataModel_NewUsing(t *testing.T) {
	dm := &corestr.HashmapDataModel{Items: map[string]string{"a": "1"}}
	hm := corestr.NewHashmapUsingDataModel(dm)
	actual := args.Map{"notNil": hm != nil, "has": hm.Has("a")}
	expected := args.Map{"notNil": true, "has": true}
	expected.ShouldBeEqual(t, 0, "HashmapDataModel NewUsing", actual)
}

func Test_I28_HashmapDataModel_NewFromCollection(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("k", "v")
	dm := corestr.NewHashmapsDataModelUsing(hm)
	actual := args.Map{"notNil": dm != nil, "len": len(dm.Items)}
	expected := args.Map{"notNil": true, "len": 1}
	expected.ShouldBeEqual(t, 0, "HashmapDataModel NewFromCollection", actual)
}
