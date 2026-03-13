package coreoncetests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coreonce"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================================================
// AnyOnce — uncovered paths
// ==========================================================================

func Test_Cov3_AnyOnce_ValueStringOnly(t *testing.T) {
	ao := coreonce.NewAnyOnce(func() any { return "hello" })
	actual := args.Map{
		"valueStringOnly": ao.ValueStringOnly(),
		"safeString":      ao.SafeString(),
		"valueStringMust": ao.ValueStringMust(),
		"valueOnly":       ao.ValueOnly(),
	}
	expected := args.Map{
		"valueStringOnly": ao.ValueString(),
		"safeString":      ao.ValueString(),
		"valueStringMust": ao.ValueString(),
		"valueOnly":       "hello",
	}
	expected.ShouldBeEqual(t, 0, "AnyOnce alias methods return expected -- hello", actual)
}

func Test_Cov3_AnyOnce_ValueString_Cached(t *testing.T) {
	ao := coreonce.NewAnyOnce(func() any { return 42 })
	// Call twice to test caching path
	first := ao.ValueString()
	second := ao.ValueString()
	actual := args.Map{"same": first == second}
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "AnyOnce ValueString cached returns same -- second call", actual)
}

func Test_Cov3_AnyOnce_CastFail(t *testing.T) {
	ao := coreonce.NewAnyOnce(func() any { return 42 })
	_, okStr := ao.CastValueString()
	_, okStrings := ao.CastValueStrings()
	_, okMap := ao.CastValueHashmapMap()
	_, okMapAny := ao.CastValueMapStringAnyMap()
	_, okBytes := ao.CastValueBytes()
	actual := args.Map{
		"okStr": okStr, "okStrings": okStrings,
		"okMap": okMap, "okMapAny": okMapAny, "okBytes": okBytes,
	}
	expected := args.Map{
		"okStr": false, "okStrings": false,
		"okMap": false, "okMapAny": false, "okBytes": false,
	}
	expected.ShouldBeEqual(t, 0, "AnyOnce Cast methods fail -- wrong type", actual)
}

func Test_Cov3_AnyOnce_ValueString_NilReturn(t *testing.T) {
	ao := coreonce.NewAnyOnce(func() any { return nil })
	result := ao.ValueString()
	actual := args.Map{"isAngelBracket": result != ""}
	expected := args.Map{"isAngelBracket": true}
	expected.ShouldBeEqual(t, 0, "AnyOnce ValueString nil returns angel bracket -- nil", actual)
}

// ==========================================================================
// ErrorOnce — with actual error
// ==========================================================================

func Test_Cov3_ErrorOnce_WithError(t *testing.T) {
	testErr := errors.New("test error")
	eo := coreonce.NewErrorOnce(func() error { return testErr })
	actual := args.Map{
		"hasError": eo.HasError(), "isEmptyError": eo.IsEmptyError(),
		"isEmpty": eo.IsEmpty(), "hasAnyItem": eo.HasAnyItem(),
		"isDefined": eo.IsDefined(), "isInvalid": eo.IsInvalid(),
		"isValid": eo.IsValid(), "isSuccess": eo.IsSuccess(),
		"isFailed": eo.IsFailed(), "isNull": eo.IsNull(),
		"isNullOrEmpty": eo.IsNullOrEmpty(),
		"message": eo.Message(),
		"isMessageEqual": eo.IsMessageEqual("test error"),
		"stringNotEmpty": eo.String() != "",
		"concatNew": eo.ConcatNew("extra") != nil,
		"concatNewStr": eo.ConcatNewString("extra") != "",
	}
	expected := args.Map{
		"hasError": true, "isEmptyError": false,
		"isEmpty": false, "hasAnyItem": true,
		"isDefined": true, "isInvalid": true,
		"isValid": false, "isSuccess": false,
		"isFailed": true, "isNull": false,
		"isNullOrEmpty": false,
		"message": "test error",
		"isMessageEqual": true,
		"stringNotEmpty": true,
		"concatNew": true, "concatNewStr": true,
	}
	expected.ShouldBeEqual(t, 0, "ErrorOnce with error returns expected -- test error", actual)
}

func Test_Cov3_ErrorOnce_Value(t *testing.T) {
	eo := coreonce.NewErrorOnce(func() error { return nil })
	actual := args.Map{
		"valueNil": eo.Value() == nil,
		"execute":  eo.Execute() == nil,
	}
	expected := args.Map{"valueNil": true, "execute": true}
	expected.ShouldBeEqual(t, 0, "ErrorOnce Value/Execute return nil -- no error", actual)
}

// ==========================================================================
// AnyErrorOnce — with error
// ==========================================================================

func Test_Cov3_AnyErrorOnce_WithError(t *testing.T) {
	testErr := errors.New("fail")
	aeo := coreonce.NewAnyErrorOnce(func() (any, error) { return nil, testErr })
	actual := args.Map{
		"hasError": aeo.HasError(), "isEmptyError": aeo.IsEmptyError(),
		"isFailed": aeo.IsFailed(), "isSuccess": aeo.IsSuccess(),
		"isNull": aeo.IsNull(), "isEmpty": aeo.IsEmpty(),
		"isValid": aeo.IsValid(), "isInvalid": aeo.IsInvalid(),
	}
	expected := args.Map{
		"hasError": true, "isEmptyError": false,
		"isFailed": true, "isSuccess": false,
		"isNull": true, "isEmpty": true,
		"isValid": false, "isInvalid": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyErrorOnce with error returns expected -- fail", actual)
}

func Test_Cov3_AnyErrorOnce_ValueString_Nil(t *testing.T) {
	aeo := coreonce.NewAnyErrorOnce(func() (any, error) { return nil, nil })
	val, err := aeo.ValueString()
	actual := args.Map{"val": val, "hasErr": err != nil}
	expected := args.Map{"val": "", "hasErr": false}
	expected.ShouldBeEqual(t, 0, "AnyErrorOnce ValueString nil value returns empty -- nil data", actual)
}

func Test_Cov3_AnyErrorOnce_CastFail(t *testing.T) {
	aeo := coreonce.NewAnyErrorOnce(func() (any, error) { return 42, nil })
	_, _, okStr := aeo.CastValueString()
	_, _, okStrings := aeo.CastValueStrings()
	_, _, okMap := aeo.CastValueHashmapMap()
	_, _, okMapAny := aeo.CastValueMapStringAnyMap()
	_, _, okBytes := aeo.CastValueBytes()
	actual := args.Map{
		"okStr": okStr, "okStrings": okStrings,
		"okMap": okMap, "okMapAny": okMapAny, "okBytes": okBytes,
	}
	expected := args.Map{
		"okStr": false, "okStrings": false,
		"okMap": false, "okMapAny": false, "okBytes": false,
	}
	expected.ShouldBeEqual(t, 0, "AnyErrorOnce Cast methods fail -- wrong type", actual)
}

// ==========================================================================
// BytesErrorOnce — with error path
// ==========================================================================

func Test_Cov3_BytesErrorOnce_WithError(t *testing.T) {
	testErr := errors.New("bytes error")
	beo := coreonce.NewBytesErrorOnce(func() ([]byte, error) { return nil, testErr })
	actual := args.Map{
		"hasError": beo.HasError(), "isEmpty": beo.IsEmpty(),
		"isValid": beo.IsValid(), "isInvalid": beo.IsInvalid(),
		"length": beo.Length(),
	}
	expected := args.Map{
		"hasError": true, "isEmpty": true,
		"isValid": false, "isInvalid": true,
		"length": 0,
	}
	expected.ShouldBeEqual(t, 0, "BytesErrorOnce with error returns expected -- error", actual)
}

func Test_Cov3_BytesErrorOnce_Valid(t *testing.T) {
	beo := coreonce.NewBytesErrorOnce(func() ([]byte, error) { return []byte("hi"), nil })
	actual := args.Map{
		"hasError": beo.HasError(), "isEmpty": beo.IsEmpty(),
		"string": beo.String(), "length": beo.Length(),
	}
	expected := args.Map{
		"hasError": false, "isEmpty": false,
		"string": "hi", "length": 2,
	}
	expected.ShouldBeEqual(t, 0, "BytesErrorOnce valid returns expected -- hi", actual)
}

// ==========================================================================
// StringOnce — additional edge cases
// ==========================================================================

func Test_Cov3_StringOnce_IsEqual_False(t *testing.T) {
	so := coreonce.NewStringOnce(func() string { return "hello" })
	actual := args.Map{
		"notEqual":     so.IsEqual("world"),
		"noPrefix":     so.HasPrefix("xyz"),
		"noSuffix":     so.HasSuffix("xyz"),
		"noContains":   so.IsContains("xyz"),
		"notStartWith": so.IsStartsWith("xyz"),
		"notEndWith":   so.IsEndsWith("xyz"),
	}
	expected := args.Map{
		"notEqual": false, "noPrefix": false, "noSuffix": false,
		"noContains": false, "notStartWith": false, "notEndWith": false,
	}
	expected.ShouldBeEqual(t, 0, "StringOnce negative checks return false -- mismatches", actual)
}

// ==========================================================================
// IntegerOnce — zero value
// ==========================================================================

func Test_Cov3_IntegerOnce_Zero(t *testing.T) {
	io := coreonce.NewIntegerOnce(func() int { return 0 })
	actual := args.Map{
		"isEmpty": io.IsEmpty(), "isZero": io.IsZero(),
		"isAboveZero": io.IsAboveZero(), "isAboveEqualZero": io.IsAboveEqualZero(),
		"isNegative": io.IsNegative(), "isPositive": io.IsPositive(),
	}
	expected := args.Map{
		"isEmpty": true, "isZero": true,
		"isAboveZero": false, "isAboveEqualZero": true,
		"isNegative": false, "isPositive": false,
	}
	expected.ShouldBeEqual(t, 0, "IntegerOnce zero value returns expected -- zero", actual)
}

// ==========================================================================
// ByteOnce — zero value
// ==========================================================================

func Test_Cov3_ByteOnce_Zero(t *testing.T) {
	bo := coreonce.NewByteOnce(func() byte { return 0 })
	actual := args.Map{
		"isEmpty": bo.IsEmpty(), "isZero": bo.IsZero(),
		"isNegative": bo.IsNegative(), "isPositive": bo.IsPositive(),
	}
	expected := args.Map{
		"isEmpty": true, "isZero": true,
		"isNegative": false, "isPositive": false,
	}
	expected.ShouldBeEqual(t, 0, "ByteOnce zero value returns expected -- zero", actual)
}

// ==========================================================================
// IntegersOnce — empty
// ==========================================================================

func Test_Cov3_IntegersOnce_Empty(t *testing.T) {
	io := coreonce.NewIntegersOnce(func() []int { return []int{} })
	actual := args.Map{"isEmpty": io.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "IntegersOnce empty returns true -- empty slice", actual)
}

// ==========================================================================
// StringsOnce — empty
// ==========================================================================

func Test_Cov3_StringsOnce_Empty(t *testing.T) {
	so := coreonce.NewStringsOnce(func() []string { return []string{} })
	actual := args.Map{
		"isEmpty": so.IsEmpty(), "hasAny": so.HasAnyItem(),
		"length": so.Length(),
	}
	expected := args.Map{"isEmpty": true, "hasAny": false, "length": 0}
	expected.ShouldBeEqual(t, 0, "StringsOnce empty returns expected -- empty slice", actual)
}

// ==========================================================================
// MapStringStringOnce — IsEqual edge: different values
// ==========================================================================

func Test_Cov3_MapStringStringOnce_Empty(t *testing.T) {
	mso := coreonce.NewMapStringStringOnce(func() map[string]string { return map[string]string{} })
	actual := args.Map{
		"isEmpty": mso.IsEmpty(), "length": mso.Length(),
		"isMissing": mso.IsMissing("any"),
	}
	expected := args.Map{"isEmpty": true, "length": 0, "isMissing": true}
	expected.ShouldBeEqual(t, 0, "MapStringStringOnce empty returns expected -- empty map", actual)
}

func Test_Cov3_MapStringStringOnce_UnmarshalJSON(t *testing.T) {
	mso := coreonce.NewMapStringStringOnce(func() map[string]string { return map[string]string{} })
	mb, _ := mso.MarshalJSON()
	err := mso.UnmarshalJSON(mb)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "MapStringStringOnce UnmarshalJSON returns no error -- valid", actual)
}
