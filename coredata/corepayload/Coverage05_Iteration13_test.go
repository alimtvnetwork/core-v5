package corepayload

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ══════════════════════════════════════════════════════════════════════════════
// Attributes — uncovered branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I13_Attributes_IsEqual_SamePtrBranch(t *testing.T) {
	a := &Attributes{}
	if !a.IsEqual(a) {
		t.Fatal("same ptr should be equal")
	}
}

func Test_I13_Attributes_IsEqual_NilVsNonNil(t *testing.T) {
	var a *Attributes
	b := &Attributes{}
	if a.IsEqual(b) {
		t.Fatal("nil vs non-nil should not be equal")
	}
}

func Test_I13_Attributes_Clone_Shallow(t *testing.T) {
	a := &Attributes{
		KeyValuePairs:    corestr.Empty.Hashmap(),
		AnyKeyValuePairs: coredynamic.EmptyMapAnyItems(),
		DynamicPayloads:  []byte("test"),
	}
	cloned, err := a.Clone(false)
	if err != nil {
		t.Fatal("expected no error:", err)
	}
	_ = cloned
}

func Test_I13_Attributes_Clone_Deep(t *testing.T) {
	a := &Attributes{
		KeyValuePairs:    corestr.Empty.Hashmap(),
		AnyKeyValuePairs: coredynamic.EmptyMapAnyItems(),
		DynamicPayloads:  []byte("test"),
	}
	cloned, err := a.ClonePtr(true)
	if err != nil {
		t.Fatal("expected no error:", err)
	}
	if cloned == nil {
		t.Fatal("expected non-nil clone")
	}
}

func Test_I13_Attributes_ClonePtr_Nil(t *testing.T) {
	var a *Attributes
	cloned, err := a.ClonePtr(false)
	if err != nil || cloned != nil {
		t.Fatal("expected nil clone for nil input")
	}
}

func Test_I13_Attributes_IsEqual_DiffPaging(t *testing.T) {
	a := &Attributes{PagingInfo: &PagingInfo{PageIndex: 1}}
	b := &Attributes{PagingInfo: &PagingInfo{PageIndex: 2}}
	if a.IsEqual(b) {
		t.Fatal("expected not equal with different paging")
	}
}

func Test_I13_Attributes_IsEqual_DiffKeyValues(t *testing.T) {
	h1 := corestr.Empty.Hashmap()
	h1.Add("k", "v1")
	h2 := corestr.Empty.Hashmap()
	h2.Add("k", "v2")
	a := &Attributes{KeyValuePairs: h1}
	b := &Attributes{KeyValuePairs: h2}
	if a.IsEqual(b) {
		t.Fatal("expected not equal with different key values")
	}
}

func Test_I13_Attributes_IsEqual_DiffDynamicPayloads(t *testing.T) {
	a := &Attributes{DynamicPayloads: []byte("aaa")}
	b := &Attributes{DynamicPayloads: []byte("bbb")}
	if a.IsEqual(b) {
		t.Fatal("expected not equal with different dynamic payloads")
	}
}

func Test_I13_Attributes_IsEqual_DiffAnyKeyValues(t *testing.T) {
	m1 := coredynamic.EmptyMapAnyItems()
	m2 := coredynamic.EmptyMapAnyItems()
	m2.AddOrUpdate("key", "val")
	a := &Attributes{AnyKeyValuePairs: m1}
	b := &Attributes{AnyKeyValuePairs: m2}
	if a.IsEqual(b) {
		t.Fatal("expected not equal with different any key values")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// Attributes — JSON
// ══════════════════════════════════════════════════════════════════════════════

func Test_I13_Attributes_Json(t *testing.T) {
	a := &Attributes{
		KeyValuePairs:   corestr.Empty.Hashmap(),
		DynamicPayloads: []byte("test"),
	}
	j := a.Json()
	if j.HasError() {
		t.Fatal("expected no json error")
	}
	// Deserialize back
	a2 := &Attributes{}
	err := a2.ParseInjectUsingJson(j)
	if err != nil {
		t.Fatal("expected no error:", err)
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// PayloadWrapper — uncovered branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I13_PayloadWrapper_NilMarshal(t *testing.T) {
	var pw *PayloadWrapper
	_, err := pw.MarshalJSON()
	if err == nil {
		t.Fatal("expected error for nil marshal")
	}
}

func Test_I13_PayloadWrapper_NilUnmarshal(t *testing.T) {
	var pw *PayloadWrapper
	err := pw.UnmarshalJSON([]byte("{}"))
	if err == nil {
		t.Fatal("expected error for nil unmarshal")
	}
}

func Test_I13_PayloadWrapper_RoundTrip(t *testing.T) {
	pw := &PayloadWrapper{
		Name:       "test",
		Identifier: "id1",
		Payloads:   []byte(`{"key":"val"}`),
	}
	jsonBytes, err := pw.MarshalJSON()
	if err != nil {
		t.Fatal("expected no marshal error:", err)
	}
	pw2 := &PayloadWrapper{}
	err = pw2.UnmarshalJSON(jsonBytes)
	if err != nil {
		t.Fatal("expected no unmarshal error:", err)
	}
	if pw2.Name != "test" {
		t.Fatal("expected name test")
	}
}

func Test_I13_PayloadWrapper_HasSafeItems(t *testing.T) {
	pw := &PayloadWrapper{
		Payloads: []byte(`{"key":"val"}`),
	}
	_ = pw.HasSafeItems()
}

func Test_I13_PayloadWrapper_IsEmpty(t *testing.T) {
	pw := &PayloadWrapper{}
	if !pw.IsEmpty() {
		t.Fatal("expected empty")
	}
}

func Test_I13_PayloadWrapper_ReCreate(t *testing.T) {
	pw := &PayloadWrapper{
		Name:     "test",
		Payloads: []byte(`{"x":1}`),
	}
	j, _ := pw.MarshalJSON()
	pw2, err := pw.ReCreateUsingJsonBytes(j)
	if err != nil {
		t.Fatal("expected no error:", err)
	}
	if pw2.Name != "test" {
		t.Fatal("expected name test")
	}
}

func Test_I13_PayloadWrapper_ReCreateUsingJsonResult(t *testing.T) {
	pw := &PayloadWrapper{
		Name:     "test",
		Payloads: []byte(`{"x":1}`),
	}
	j, _ := pw.MarshalJSON()
	jr := corejson.NewResult.UsingBytes(j)
	pw2, err := pw.ReCreateUsingJsonResult(jr)
	if err != nil {
		t.Fatal("expected no error:", err)
	}
	if pw2.Name != "test" {
		t.Fatal("expected name test")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// PayloadsCollection — uncovered branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I13_PayloadsCollection_AddAndFilter(t *testing.T) {
	pc := New.PayloadsCollection.Empty()
	pc.Add(*New.PayloadWrapper.UsingBytes("a", "id1", "task", "cat", "ent", []byte(`"x"`)))
	pc.Add(*New.PayloadWrapper.UsingBytes("b", "id2", "task", "cat", "ent", []byte(`"y"`)))
	if pc.Length() != 2 {
		t.Fatal("expected 2 items")
	}
	filtered := pc.Filter(func(w *PayloadWrapper) bool {
		return w.Name == "a"
	})
	if len(filtered) != 1 {
		t.Fatal("expected 1 filtered item")
	}
}

func Test_I13_PayloadsCollection_NilReceiver(t *testing.T) {
	var pc *PayloadsCollection
	if pc.HasAnyItem() {
		t.Fatal("nil should have no items")
	}
}

func Test_I13_PayloadsCollection_Json(t *testing.T) {
	pc := Empty.PayloadsCollection()
	pw := &PayloadWrapper{Name: "a", Payloads: []byte(`{"k":"v"}`)}
	pc.Add(pw)
	j := pc.Json()
	if j.HasError() {
		t.Fatal("expected no json error")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// Empty and New creators
// ══════════════════════════════════════════════════════════════════════════════

func Test_I13_EmptyCreators(t *testing.T) {
	_ = Empty.Attributes()
	_ = Empty.AttributesDefaults()
	_ = Empty.PayloadWrapper()
	_ = Empty.PayloadsCollection()
}

func Test_I13_NewAttributes_All(t *testing.T) {
	a := New.Attributes.All(
		nil,
		corestr.Empty.Hashmap(),
		coredynamic.EmptyMapAnyItems(),
		nil,
		[]byte("dp"),
		nil,
		nil,
	)
	if a == nil {
		t.Fatal("expected non-nil attributes")
	}
}

func Test_I13_PayloadWrapper_Getters(t *testing.T) {
	pw := &PayloadWrapper{
		Name:           "test",
		Identifier:     "id1",
		TaskTypeName:   "task",
		EntityType:     "entity",
		CategoryName:   "cat",
		HasManyRecords: true,
		Payloads:       []byte(`{"key":"val"}`),
		Attributes:     Empty.AttributesDefaults(),
	}
	_ = pw.PayloadsString()
	_ = pw.HasPayloads()
	_ = pw.IsPayloadsEmpty()
	_ = pw.HasAttributes()
	if pw.IsEmpty() {
		t.Fatal("expected not empty")
	}
}
