package corejsontests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── New / NewPtr ──

func Test_New_Simple_Cov(t *testing.T) {
	r := corejson.New("hello")
	actual := args.Map{
		"hasError":  r.Error != nil,
		"isEmpty":   r.IsEmpty(),
		"hasBytes":  len(r.Bytes) > 0,
		"typeName":  r.TypeName != "",
	}
	expected := args.Map{
		"hasError":  false,
		"isEmpty":   false,
		"hasBytes":  true,
		"typeName":  true,
	}
	expected.ShouldBeEqual(t, 0, "New_Simple", actual)
}

func Test_NewPtr_Simple_Cov(t *testing.T) {
	r := corejson.NewPtr("hello")
	actual := args.Map{
		"isNil":    r == nil,
		"hasError": r.Error != nil,
	}
	expected := args.Map{
		"isNil":    false,
		"hasError": false,
	}
	expected.ShouldBeEqual(t, 0, "NewPtr_Simple", actual)
}

func Test_New_Struct_Cov(t *testing.T) {
	type testS struct{ A int }
	r := corejson.New(testS{A: 42})
	actual := args.Map{
		"hasError":         r.Error != nil,
		"hasBytes":         len(r.Bytes) > 0,
		"stringNotEmpty":   r.String() != "",
		"jsonStringNotNil": r.JsonString() != "",
	}
	expected := args.Map{
		"hasError":         false,
		"hasBytes":         true,
		"stringNotEmpty":   true,
		"jsonStringNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "New_Struct", actual)
}

func Test_New_Nil_Cov(t *testing.T) {
	r := corejson.New(nil)
	actual := args.Map{
		"hasError": r.Error != nil,
	}
	expected := args.Map{
		"hasError": false,
	}
	expected.ShouldBeEqual(t, 0, "New_Nil", actual)
}

// ── Result methods ──

func Test_Result_IsEmpty_Cov(t *testing.T) {
	empty := corejson.Result{}
	actual := args.Map{
		"isEmpty":       empty.IsEmpty(),
		"isEmptyError":  empty.IsEmptyError() != nil,
		"hasError":      empty.HasError(),
		"hasNoError":    !empty.HasError(),
		"isValid":       empty.IsValid(),
	}
	expected := args.Map{
		"isEmpty":       true,
		"isEmptyError":  true,
		"hasError":      false,
		"hasNoError":    true,
		"isValid":       false,
	}
	expected.ShouldBeEqual(t, 0, "Result_IsEmpty", actual)
}

func Test_Result_String_Cov(t *testing.T) {
	r := corejson.New(42)
	s := r.String()
	if s == "" {
		t.Error("should not be empty")
	}
}

func Test_Result_JsonString_Cov(t *testing.T) {
	r := corejson.New(42)
	s := r.JsonString()
	if s != "42" {
		t.Errorf("expected 42, got %s", s)
	}
}

func Test_Result_SafeBytes_Cov(t *testing.T) {
	r := corejson.New(42)
	b := r.SafeBytes()
	if len(b) == 0 {
		t.Error("should not be empty")
	}
}

func Test_Result_SafeBytes_Empty_Cov(t *testing.T) {
	r := corejson.Result{}
	b := r.SafeBytes()
	if b == nil {
		t.Error("should not be nil")
	}
}

func Test_Result_MustBytes_Cov(t *testing.T) {
	r := corejson.New(42)
	b := r.MustBytes()
	if len(b) == 0 {
		t.Error("should not be empty")
	}
}

func Test_Result_PrettyJsonString_Cov(t *testing.T) {
	r := corejson.New(map[string]int{"a": 1})
	s := r.PrettyJsonString()
	if s == "" {
		t.Error("should not be empty")
	}
}

func Test_Result_PrettyJsonBytes_Cov(t *testing.T) {
	r := corejson.New(map[string]int{"a": 1})
	b := r.PrettyJsonBytes()
	if len(b) == 0 {
		t.Error("should not be empty")
	}
}

func Test_Result_Clone_Cov(t *testing.T) {
	r := corejson.New(42)
	c := r.Clone()
	actual := args.Map{
		"hasError": c.Error != nil,
		"hasBytes": len(c.Bytes) > 0,
	}
	expected := args.Map{
		"hasError": false,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "Result_Clone", actual)
}

func Test_Result_ClonePtr_Cov(t *testing.T) {
	r := corejson.New(42)
	c := r.ClonePtr()
	if c == nil {
		t.Error("should not be nil")
	}
}

func Test_Result_IsEqual_SameContent_Cov(t *testing.T) {
	r1 := corejson.New(42)
	r2 := corejson.New(42)
	if !r1.IsEqual(&r2) {
		t.Error("same content should be equal")
	}
}

func Test_Result_IsEqual_DiffContent_Cov(t *testing.T) {
	r1 := corejson.New(42)
	r2 := corejson.New(43)
	if r1.IsEqual(&r2) {
		t.Error("different content should not be equal")
	}
}

func Test_Result_IsEqual_Nil_Cov(t *testing.T) {
	r := corejson.New(42)
	if r.IsEqual(nil) {
		t.Error("should not be equal to nil")
	}
}

func Test_Result_Unmarshal_Cov(t *testing.T) {
	type testS struct{ A int }
	r := corejson.New(testS{A: 42})
	var out testS
	err := r.Unmarshal(&out)
	if err != nil || out.A != 42 {
		t.Error("should unmarshal correctly")
	}
}

// ── Empty creator ──

func Test_EmptyResult_Cov(t *testing.T) {
	r := corejson.Empty.Result()
	if !r.IsEmpty() {
		t.Error("should be empty")
	}
}

func Test_EmptyResultPtr_Cov(t *testing.T) {
	r := corejson.Empty.ResultPtr()
	if r == nil || !r.IsEmpty() {
		t.Error("should be empty ptr")
	}
}

// ── Serialize / Deserialize ──

func Test_Serialize_Default_Cov(t *testing.T) {
	r := corejson.Serialize.Default(42)
	if r.HasError() {
		t.Error("should not error")
	}
}

func Test_Serialize_DefaultPtr_Cov(t *testing.T) {
	r := corejson.Serialize.DefaultPtr(42)
	if r == nil || r.HasError() {
		t.Error("should not error")
	}
}

func Test_Deserialize_FromResult_Cov(t *testing.T) {
	type testS struct{ A int }
	r := corejson.New(testS{A: 42})
	var out testS
	err := corejson.Deserialize.FromResult(r, &out)
	if err != nil || out.A != 42 {
		t.Error("should deserialize")
	}
}

func Test_Deserialize_FromBytes_Cov(t *testing.T) {
	type testS struct{ A int }
	var out testS
	err := corejson.Deserialize.FromBytes([]byte(`{"A":42}`), &out)
	if err != nil || out.A != 42 {
		t.Error("should deserialize")
	}
}

// ── CastAny ──

func Test_CastAny_ToString_Cov(t *testing.T) {
	r := corejson.CastAny.ToString(42)
	if r == "" {
		t.Error("should not be empty")
	}
}

// ── AnyTo ──

func Test_AnyTo_Result_Cov(t *testing.T) {
	r := corejson.AnyTo.Result(42)
	if r.HasError() {
		t.Error("should not error")
	}
}

func Test_AnyTo_ResultPtr_Cov(t *testing.T) {
	r := corejson.AnyTo.ResultPtr(42)
	if r == nil {
		t.Error("should not be nil")
	}
}

// ── NewResult ──

func Test_NewResult_UsingBytes_Cov(t *testing.T) {
	r := corejson.NewResult.UsingBytes([]byte(`"hello"`))
	if r.IsEmpty() {
		t.Error("should not be empty")
	}
}

func Test_NewResult_UsingBytesPtr_Cov(t *testing.T) {
	r := corejson.NewResult.UsingBytesPtr([]byte(`"hello"`))
	if r == nil {
		t.Error("should not be nil")
	}
}

// ── BytesCollection ──

func Test_NewBytesCollection_Cap_Cov(t *testing.T) {
	c := corejson.NewBytesCollection.Cap(5)
	if c == nil {
		t.Error("should not be nil")
	}
}

// ── NewResultsCollection ──

func Test_NewResultsCollection_Cap_Cov(t *testing.T) {
	c := corejson.NewResultsCollection.Cap(5)
	if c == nil {
		t.Error("should not be nil")
	}
}

// ── NewResultsPtrCollection ──

func Test_NewResultsPtrCollection_Cap_Cov(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Cap(5)
	if c == nil {
		t.Error("should not be nil")
	}
}

// ── Pretty ──

func Test_Pretty_FromBytes_Cov(t *testing.T) {
	result := corejson.Pretty.FromBytes([]byte(`{"a":1}`))
	if len(result) == 0 {
		t.Error("should not be empty")
	}
}

func Test_Pretty_FromString_Cov(t *testing.T) {
	result := corejson.Pretty.FromString(`{"a":1}`)
	if result == "" {
		t.Error("should not be empty")
	}
}

// ── NewMapResults ──

func Test_NewMapResults_Empty_Cov(t *testing.T) {
	m := corejson.NewMapResults.Empty()
	if m == nil {
		t.Error("should not be nil")
	}
}
