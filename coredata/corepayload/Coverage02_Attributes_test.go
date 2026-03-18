package corepayload

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
)

func TestAttributes_BasicOps(t *testing.T) {
	attr := New.Attributes.Empty()
	if attr == nil {
		t.Fatal("expected non-nil")
	}
	if attr.HasError() {
		t.Fatal("should not have error")
	}
	if !attr.IsValid() {
		t.Fatal("should be valid")
	}
	if attr.IsInvalid() {
		t.Fatal("should not be invalid")
	}

	var nilAttr *Attributes
	if !nilAttr.IsNull() {
		t.Fatal("nil should be null")
	}
	if nilAttr.Length() != 0 {
		t.Fatal("nil length should be 0")
	}
}

func TestAttributes_EmptyChecks(t *testing.T) {
	attr := New.Attributes.Empty()
	if !attr.IsEmpty() {
		t.Fatal("should be empty")
	}
	if attr.HasItems() {
		t.Fatal("should not have items")
	}
	if attr.HasDynamicPayloads() {
		t.Fatal("should not have dynamic payloads")
	}
	if attr.HasStringKeyValuePairs() {
		t.Fatal("should not have string kv")
	}
	if attr.HasAnyKeyValuePairs() {
		t.Fatal("should not have any kv")
	}
}

func TestAttributes_WithDynamicPayloads(t *testing.T) {
	payload := []byte(`{"key":"value"}`)
	attr := New.Attributes.UsingDynamicPayloadBytes(payload)
	if attr.IsEmpty() {
		t.Fatal("should not be empty")
	}
	if !attr.HasDynamicPayloads() {
		t.Fatal("should have payloads")
	}
	if attr.Length() == 0 {
		t.Fatal("should have length")
	}
	if attr.PayloadsString() == "" {
		t.Fatal("expected non-empty payloads string")
	}
	b := attr.Payloads()
	if len(b) == 0 {
		t.Fatal("expected bytes")
	}
}

func TestAttributes_WithKeyValues(t *testing.T) {
	hm := corestr.New.Hashmap.UsingMap(map[string]string{"k1": "v1"})
	attr := New.Attributes.UsingKeyValues(hm)
	if attr.IsKeyValuePairsEmpty() {
		t.Fatal("should have kv pairs")
	}
	if !attr.HasStringKey("k1") {
		t.Fatal("should have key")
	}
	v, found := attr.GetStringKeyValue("k1")
	if !found || v != "v1" {
		t.Fatal("unexpected value")
	}

	m := attr.Hashmap()
	if len(m) == 0 {
		t.Fatal("expected map")
	}
}

func TestAttributes_WithAnyKeyValues(t *testing.T) {
	anyMap := coredynamic.NewMapAnyItems(map[string]any{"k": "v"})
	attr := New.Attributes.UsingAnyKeyValues(anyMap)
	if attr.IsAnyKeyValuePairsEmpty() {
		t.Fatal("should have any kv")
	}
	if !attr.HasAnyKey("k") {
		t.Fatal("should have key")
	}
	m := attr.AnyKeyValMap()
	if len(m) == 0 {
		t.Fatal("expected map")
	}
}

func TestAttributes_PagingInfo(t *testing.T) {
	attr := New.Attributes.Empty()
	if attr.HasPagingInfo() {
		t.Fatal("should not have paging info")
	}
	if !attr.IsPagingInfoEmpty() {
		t.Fatal("paging info should be empty")
	}
}

func TestAttributes_AuthInfo(t *testing.T) {
	attr := New.Attributes.Empty()
	if attr.HasAuthInfo() {
		t.Fatal("should not have auth info")
	}
	if !attr.IsAuthInfoEmpty() {
		t.Fatal("auth info should be empty")
	}
	if attr.HasUserInfo() {
		t.Fatal("should not have user info")
	}
	if attr.HasSessionInfo() {
		t.Fatal("should not have session info")
	}
}

func TestAttributes_WithAuthInfo(t *testing.T) {
	ai := &AuthInfo{
		ActionType: "login",
		UserInfo:   &UserInfo{User: New.User.UsingName("alice")},
	}
	attr := New.Attributes.UsingAuthInfo(ai)
	if !attr.HasAuthInfo() {
		t.Fatal("should have auth info")
	}
	if !attr.HasUserInfo() {
		t.Fatal("should have user info")
	}
	if attr.VirtualUser() == nil {
		t.Fatal("should have virtual user")
	}
	if attr.AuthType() != "login" {
		t.Fatal("unexpected auth type")
	}
}

func TestAttributes_SetAuthInfo(t *testing.T) {
	attr := New.Attributes.Empty()
	attr.SetAuthInfo(&AuthInfo{ActionType: "test"})
	if attr.AuthType() != "test" {
		t.Fatal("unexpected")
	}

	var nilAttr *Attributes
	result := nilAttr.SetAuthInfo(&AuthInfo{ActionType: "test"})
	if result == nil {
		t.Fatal("expected non-nil")
	}
}

func TestAttributes_AddOrUpdateString(t *testing.T) {
	attr := New.Attributes.Empty()
	added := attr.AddOrUpdateString("k", "v")
	if !added {
		t.Fatal("expected added")
	}

	var nilAttr *Attributes
	if nilAttr.AddOrUpdateString("k", "v") {
		t.Fatal("nil should return false")
	}
}

func TestAttributes_Clear_Dispose(t *testing.T) {
	attr := New.Attributes.Empty()
	attr.Clear()
	attr.Dispose()

	var nilAttr *Attributes
	nilAttr.Clear()
	nilAttr.Dispose()
}

func TestAttributes_IsEqual(t *testing.T) {
	a1 := New.Attributes.UsingDynamicPayloadBytes([]byte(`"x"`))
	a2 := New.Attributes.UsingDynamicPayloadBytes([]byte(`"x"`))
	if !a1.IsEqual(a2) {
		t.Fatal("should be equal")
	}

	var nilA *Attributes
	if !nilA.IsEqual(nil) {
		t.Fatal("both nil should be equal")
	}
}

func TestAttributes_Clone(t *testing.T) {
	attr := New.Attributes.UsingDynamicPayloadBytes([]byte(`"test"`))
	c, err := attr.Clone(false)
	if err != nil {
		t.Fatal(err)
	}
	_ = c

	c2, err2 := attr.Clone(true)
	if err2 != nil {
		t.Fatal(err2)
	}
	_ = c2

	var nilAttr *Attributes
	cp, err3 := nilAttr.ClonePtr(false)
	if err3 != nil || cp != nil {
		t.Fatal("expected nil,nil for nil")
	}
}

func TestAttributes_Json(t *testing.T) {
	attr := New.Attributes.UsingDynamicPayloadBytes([]byte(`"x"`))
	s := attr.JsonString()
	if s == "" {
		t.Fatal("expected non-empty")
	}
	_ = attr.String()
	_ = attr.PrettyJsonString()
	_ = attr.Json()
	_ = attr.JsonPtr()
	_ = attr.JsonModel()
	_ = attr.JsonModelAny()
}

func TestAttributes_DeserializeDynamicPayloads(t *testing.T) {
	attr := New.Attributes.UsingDynamicPayloadBytes([]byte(`"hello"`))
	var out string
	err := attr.DeserializeDynamicPayloads(&out)
	if err != nil || out != "hello" {
		t.Fatal("unexpected")
	}
}

func TestAttributes_PayloadsJsonResult(t *testing.T) {
	attr := New.Attributes.UsingDynamicPayloadBytes([]byte(`"x"`))
	r := attr.PayloadsJsonResult()
	if r == nil {
		t.Fatal("expected non-nil")
	}

	empty := New.Attributes.Empty()
	r2 := empty.PayloadsJsonResult()
	_ = r2
}

func TestAttributes_PayloadsPrettyString(t *testing.T) {
	attr := New.Attributes.UsingDynamicPayloadBytes([]byte(`{"a":1}`))
	s := attr.PayloadsPrettyString()
	if s == "" {
		t.Fatal("expected non-empty")
	}

	empty := New.Attributes.Empty()
	s2 := empty.PayloadsPrettyString()
	if s2 != "" {
		t.Fatal("expected empty")
	}
}

func TestAttributes_NonPtr(t *testing.T) {
	attr := Attributes{}
	np := attr.NonPtr()
	_ = np
}

func TestAttributes_ParseInjectUsingJson(t *testing.T) {
	attr := New.Attributes.UsingDynamicPayloadBytes([]byte(`"x"`))
	jr := corejson.New(attr)
	target := &Attributes{}
	_, err := target.ParseInjectUsingJson(jr.Ptr())
	if err != nil {
		t.Fatal(err)
	}
}

func TestAttributes_HasSafeItems(t *testing.T) {
	attr := New.Attributes.UsingDynamicPayloadBytes([]byte(`"x"`))
	if !attr.HasSafeItems() {
		t.Fatal("should have safe items")
	}
}

func TestAttributes_Count_Capacity(t *testing.T) {
	attr := New.Attributes.UsingDynamicPayloadBytes([]byte(`"x"`))
	if attr.Count() == 0 {
		t.Fatal("expected non-zero")
	}
	if attr.Capacity() == 0 {
		t.Fatal("expected non-zero")
	}
}

func TestNewAttributesCreator(t *testing.T) {
	_ = New.Attributes.Empty()
	_ = New.Attributes.UsingDynamicPayloadBytes([]byte(`"x"`))
	_ = New.Attributes.UsingAuthInfo(&AuthInfo{})
	_ = New.Attributes.UsingKeyValues(corestr.Empty.Hashmap())
	_ = New.Attributes.UsingAnyKeyValues(coredynamic.EmptyMapAnyItems())
	_ = New.Attributes.UsingBasicError(nil)
	_ = New.Attributes.Create(nil, nil, []byte{})

	_, _ = New.Attributes.AllAny(nil, nil, nil, nil, "x")
	_, _ = New.Attributes.PageInfoAny(nil, "x")
	_, _ = New.Attributes.UsingDynamicPayloadAny(nil, "x")

	_ = New.Attributes.All(nil, nil, nil, nil, nil, nil, nil)
	_ = New.Attributes.UsingAuthInfoKeyValues(nil, nil)
	_ = New.Attributes.UsingKeyValuesPlusDynamic(nil, nil)
	_ = New.Attributes.UsingAuthInfoAnyKeyValues(nil, nil)
	_ = New.Attributes.UsingAnyKeyValuesPlusDynamic(nil, nil)
	_ = New.Attributes.UsingAuthInfoDynamicBytes(nil, nil)
	_ = New.Attributes.ErrFromTo(nil, nil, nil)
}

func TestNewAttributesCreator_Deserialize(t *testing.T) {
	attr := New.Attributes.UsingDynamicPayloadBytes([]byte(`"x"`))
	b, _ := corejson.Serialize.Raw(attr)
	a2, err := New.Attributes.Deserialize(b)
	if err != nil || a2 == nil {
		t.Fatal("unexpected")
	}

	_, err2 := New.Attributes.Deserialize([]byte("invalid"))
	if err2 == nil {
		t.Fatal("expected error")
	}
}

func TestNewAttributesCreator_CastOrDeserializeFrom(t *testing.T) {
	attr := New.Attributes.UsingDynamicPayloadBytes([]byte(`"x"`))
	_, _ = New.Attributes.CastOrDeserializeFrom(attr)
	_, err := New.Attributes.CastOrDeserializeFrom(nil)
	if err == nil {
		t.Fatal("expected error for nil")
	}
}

func TestEmptyCreator_Attributes(t *testing.T) {
	_ = Empty.Attributes()
	_ = Empty.AttributesDefaults()
	_ = Empty.PayloadWrapper()
	_ = Empty.PayloadsCollection()
}
