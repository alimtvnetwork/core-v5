package corepayloadtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corepayload"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Attributes — full coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_Attributes_NewEmpty(t *testing.T) {
	a := corepayload.NewAttributes.Empty()
	actual := args.Map{"empty": a.IsEmpty(), "len": a.Length()}
	expected := args.Map{"empty": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "Attributes empty", actual)
}

func Test_Cov14_Attributes_AddGet(t *testing.T) {
	a := corepayload.NewAttributes.Empty()
	a.Set("key", "val")
	v, ok := a.Get("key")
	actual := args.Map{"val": v, "ok": ok}
	expected := args.Map{"val": "val", "ok": true}
	expected.ShouldBeEqual(t, 0, "Attributes add/get", actual)
}

func Test_Cov14_Attributes_Get_Missing(t *testing.T) {
	a := corepayload.NewAttributes.Empty()
	_, ok := a.Get("missing")
	actual := args.Map{"ok": ok}
	expected := args.Map{"ok": false}
	expected.ShouldBeEqual(t, 0, "Attributes get missing", actual)
}

func Test_Cov14_Attributes_Nil(t *testing.T) {
	var a *corepayload.Attributes
	actual := args.Map{"nil": a.IsNull(), "empty": a.IsEmpty(), "len": a.Length()}
	expected := args.Map{"nil": true, "empty": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "Attributes nil", actual)
}

func Test_Cov14_Attributes_Has(t *testing.T) {
	a := corepayload.NewAttributes.Empty()
	a.Set("key", "val")
	actual := args.Map{"has": a.Has("key"), "notHas": a.Has("miss")}
	expected := args.Map{"has": true, "notHas": false}
	expected.ShouldBeEqual(t, 0, "Attributes Has", actual)
}

func Test_Cov14_Attributes_Remove(t *testing.T) {
	a := corepayload.NewAttributes.Empty()
	a.Set("key", "val")
	a.Remove("key")
	actual := args.Map{"has": a.Has("key")}
	expected := args.Map{"has": false}
	expected.ShouldBeEqual(t, 0, "Attributes Remove", actual)
}

func Test_Cov14_Attributes_GetString(t *testing.T) {
	a := corepayload.NewAttributes.Empty()
	a.Set("key", "val")
	actual := args.Map{"v": a.GetString("key"), "missing": a.GetString("miss")}
	expected := args.Map{"v": "val", "missing": ""}
	expected.ShouldBeEqual(t, 0, "Attributes GetString", actual)
}

func Test_Cov14_Attributes_Clear(t *testing.T) {
	a := corepayload.NewAttributes.Empty()
	a.Set("key", "val")
	a.Clear()
	actual := args.Map{"empty": a.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Attributes Clear", actual)
}

func Test_Cov14_Attributes_Json(t *testing.T) {
	a := corepayload.NewAttributes.Empty()
	a.Set("k", "v")
	r := a.Json()
	actual := args.Map{"noErr": !r.HasError()}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Attributes Json", actual)
}

func Test_Cov14_Attributes_String(t *testing.T) {
	a := corepayload.NewAttributes.Empty()
	a.Set("k", "v")
	actual := args.Map{"notEmpty": a.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Attributes String", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// AuthInfo / SessionInfo / UserInfo
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_AuthInfo_Empty(t *testing.T) {
	a := corepayload.AuthInfo{}
	actual := args.Map{"token": a.Token}
	expected := args.Map{"token": ""}
	expected.ShouldBeEqual(t, 0, "AuthInfo empty", actual)
}

func Test_Cov14_SessionInfo_Empty(t *testing.T) {
	s := corepayload.SessionInfo{}
	actual := args.Map{"sid": s.SessionID}
	expected := args.Map{"sid": ""}
	expected.ShouldBeEqual(t, 0, "SessionInfo empty", actual)
}

func Test_Cov14_UserInfo_Empty(t *testing.T) {
	u := corepayload.UserInfo{}
	actual := args.Map{"id": u.UserID}
	expected := args.Map{"id": ""}
	expected.ShouldBeEqual(t, 0, "UserInfo empty", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// PagingInfo
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_PagingInfo_Empty(t *testing.T) {
	p := corepayload.PagingInfo{}
	actual := args.Map{"page": p.PageIndex, "size": p.PageSize}
	expected := args.Map{"page": 0, "size": 0}
	expected.ShouldBeEqual(t, 0, "PagingInfo empty", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// PayloadWrapper — comprehensive
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_PayloadWrapper_Nil(t *testing.T) {
	var pw *corepayload.PayloadWrapper
	actual := args.Map{
		"isNull":   pw.IsNull(),
		"isEmpty":  pw.IsEmpty(),
		"hasItems": pw.HasItems(),
		"len":      pw.Length(),
		"name":     pw.NameValue(),
		"id":       pw.IdString(),
	}
	expected := args.Map{
		"isNull": true, "isEmpty": true, "hasItems": false,
		"len": 0, "name": "", "id": "",
	}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper nil", actual)
}

func Test_Cov14_PayloadWrapper_Valid(t *testing.T) {
	pw := corepayload.NewPayloadWrapper.From("name", "id1", "entity", []byte(`{"a":1}`))
	actual := args.Map{
		"name":     pw.Name,
		"id":       pw.Identifier,
		"notEmpty": !pw.IsEmpty(),
		"hasItems": pw.HasItems(),
	}
	expected := args.Map{
		"name": "name", "id": "id1", "notEmpty": true, "hasItems": true,
	}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper valid", actual)
}

func Test_Cov14_PayloadWrapper_Json_Nil(t *testing.T) {
	var pw *corepayload.PayloadWrapper
	r := pw.Json()
	actual := args.Map{"empty": r.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper Json nil", actual)
}

func Test_Cov14_PayloadWrapper_Json_Valid(t *testing.T) {
	pw := corepayload.NewPayloadWrapper.From("n", "id", "e", []byte(`{"a":1}`))
	r := pw.Json()
	actual := args.Map{"noErr": !r.HasError()}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper Json valid", actual)
}

func Test_Cov14_PayloadWrapper_String_Nil(t *testing.T) {
	var pw *corepayload.PayloadWrapper
	actual := args.Map{"v": pw.String()}
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper String nil", actual)
}

func Test_Cov14_PayloadWrapper_Getters(t *testing.T) {
	pw := corepayload.NewPayloadWrapper.From("n", "id", "e", []byte(`{}`))
	actual := args.Map{
		"name":       pw.NameValue(),
		"id":         pw.IdString(),
		"entity":     pw.EntityType,
		"payloads":   pw.PayloadsString(),
		"hasPayload": pw.HasPayloads(),
	}
	expected := args.Map{
		"name": "n", "id": "id", "entity": "e",
		"payloads": "{}", "hasPayload": true,
	}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper getters", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// PayloadsCollection
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_PayloadsCollection_Nil(t *testing.T) {
	var pc *corepayload.PayloadsCollection
	actual := args.Map{
		"isNull":  pc.IsNull(),
		"isEmpty": pc.IsEmpty(),
		"len":     pc.Length(),
	}
	expected := args.Map{
		"isNull": true, "isEmpty": true, "len": 0,
	}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection nil", actual)
}

func Test_Cov14_PayloadsCollection_Empty(t *testing.T) {
	pc := corepayload.NewPayloadsCollection.Empty()
	actual := args.Map{"empty": pc.IsEmpty(), "len": pc.Length()}
	expected := args.Map{"empty": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection empty", actual)
}

func Test_Cov14_PayloadsCollection_Add(t *testing.T) {
	pc := corepayload.NewPayloadsCollection.Empty()
	pw := corepayload.NewPayloadWrapper.From("n", "id", "e", []byte(`{}`))
	pc.Add(pw)
	actual := args.Map{"len": pc.Length(), "hasAny": pc.HasAnyItem()}
	expected := args.Map{"len": 1, "hasAny": true}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection add", actual)
}

func Test_Cov14_PayloadsCollection_Json_Nil(t *testing.T) {
	var pc *corepayload.PayloadsCollection
	r := pc.Json()
	actual := args.Map{"empty": r.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection Json nil", actual)
}

func Test_Cov14_PayloadsCollection_First_Nil(t *testing.T) {
	var pc *corepayload.PayloadsCollection
	f := pc.First()
	actual := args.Map{"nil": f == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection First nil", actual)
}

func Test_Cov14_PayloadsCollection_Last_Nil(t *testing.T) {
	var pc *corepayload.PayloadsCollection
	l := pc.Last()
	actual := args.Map{"nil": l == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection Last nil", actual)
}

func Test_Cov14_PayloadsCollection_String_Nil(t *testing.T) {
	var pc *corepayload.PayloadsCollection
	actual := args.Map{"v": pc.String()}
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection String nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// BytesCreateInstruction / PayloadCreateInstruction
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_BytesCreateInstruction(t *testing.T) {
	bci := corepayload.BytesCreateInstruction{
		Name:       "n",
		Identifier: "id",
		EntityType: "e",
		Payloads:   []byte(`{}`),
	}
	actual := args.Map{"name": bci.Name, "id": bci.Identifier}
	expected := args.Map{"name": "n", "id": "id"}
	expected.ShouldBeEqual(t, 0, "BytesCreateInstruction", actual)
}

func Test_Cov14_BytesCreateInstruction_String(t *testing.T) {
	bci := corepayload.BytesCreateInstruction{
		Name: "n", Identifier: "id", EntityType: "e", Payloads: []byte(`{}`),
	}
	actual := args.Map{"notEmpty": bci.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "BytesCreateInstruction String", actual)
}

func Test_Cov14_PayloadCreateInstruction(t *testing.T) {
	pci := corepayload.PayloadCreateInstruction{
		Name:       "n",
		Identifier: "id",
		EntityType: "e",
	}
	actual := args.Map{"name": pci.Name}
	expected := args.Map{"name": "n"}
	expected.ShouldBeEqual(t, 0, "PayloadCreateInstruction", actual)
}

func Test_Cov14_PayloadCreateInstruction_String(t *testing.T) {
	pci := corepayload.PayloadCreateInstruction{
		Name: "n", Identifier: "id", EntityType: "e",
	}
	actual := args.Map{"notEmpty": pci.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "PayloadCreateInstruction String", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// User
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_User_Empty(t *testing.T) {
	u := corepayload.NewUser.Empty()
	actual := args.Map{"empty": u.IsEmpty(), "id": u.UserID}
	expected := args.Map{"empty": true, "id": ""}
	expected.ShouldBeEqual(t, 0, "User empty", actual)
}

func Test_Cov14_User_Set(t *testing.T) {
	u := corepayload.NewUser.From("uid", "uname")
	actual := args.Map{"id": u.UserID, "name": u.Username}
	expected := args.Map{"id": "uid", "name": "uname"}
	expected.ShouldBeEqual(t, 0, "User set", actual)
}

func Test_Cov14_User_Nil(t *testing.T) {
	var u *corepayload.User
	actual := args.Map{"isNull": u.IsNull(), "isEmpty": u.IsEmpty()}
	expected := args.Map{"isNull": true, "isEmpty": true}
	expected.ShouldBeEqual(t, 0, "User nil", actual)
}

func Test_Cov14_User_Json(t *testing.T) {
	u := corepayload.NewUser.From("uid", "uname")
	r := u.Json()
	actual := args.Map{"noErr": !r.HasError()}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "User Json", actual)
}

func Test_Cov14_User_String(t *testing.T) {
	u := corepayload.NewUser.From("uid", "uname")
	actual := args.Map{"notEmpty": u.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "User String", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// PayloadTypeExpander
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_PayloadTypeExpander(t *testing.T) {
	pw := corepayload.NewPayloadWrapper.From("n", "id", "e", []byte(`{}`))
	expanded := corepayload.PayloadTypeExpander(pw, "newType")
	actual := args.Map{"type": expanded.EntityType}
	expected := args.Map{"type": "newType"}
	expected.ShouldBeEqual(t, 0, "PayloadTypeExpander", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// PayloadWrapper — additional methods for coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_PayloadWrapper_IdInteger(t *testing.T) {
	pw := corepayload.NewPayloadWrapper.From("n", "42", "e", []byte(`{}`))
	actual := args.Map{"v": pw.IdInteger()}
	expected := args.Map{"v": 42}
	expected.ShouldBeEqual(t, 0, "IdInteger", actual)
}

func Test_Cov14_PayloadWrapper_IdInteger_Invalid(t *testing.T) {
	pw := corepayload.NewPayloadWrapper.From("n", "abc", "e", []byte(`{}`))
	actual := args.Map{"v": pw.IdInteger()}
	expected := args.Map{"v": -1}
	expected.ShouldBeEqual(t, 0, "IdInteger invalid", actual)
}

func Test_Cov14_PayloadWrapper_Clear(t *testing.T) {
	pw := corepayload.NewPayloadWrapper.From("n", "id", "e", []byte(`{}`))
	pw.Clear()
	actual := args.Map{"empty": pw.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper Clear", actual)
}

func Test_Cov14_PayloadWrapper_Dispose(t *testing.T) {
	pw := corepayload.NewPayloadWrapper.From("n", "id", "e", []byte(`{}`))
	pw.Dispose()
	actual := args.Map{"null": pw.IsNull()}
	expected := args.Map{"null": true}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper Dispose", actual)
}

func Test_Cov14_PayloadWrapper_HasError_Nil(t *testing.T) {
	var pw *corepayload.PayloadWrapper
	actual := args.Map{"v": pw.HasError()}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper HasError nil", actual)
}

func Test_Cov14_PayloadWrapper_HasSafeItems(t *testing.T) {
	pw := corepayload.NewPayloadWrapper.From("n", "id", "e", []byte(`{}`))
	actual := args.Map{"v": pw.HasSafeItems()}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "HasSafeItems", actual)
}

func Test_Cov14_PayloadWrapper_JsonPtr(t *testing.T) {
	pw := corepayload.NewPayloadWrapper.From("n", "id", "e", []byte(`{}`))
	r := pw.JsonPtr()
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "JsonPtr", actual)
}

func Test_Cov14_PayloadWrapper_PrettyJsonString(t *testing.T) {
	pw := corepayload.NewPayloadWrapper.From("n", "id", "e", []byte(`{}`))
	actual := args.Map{"notEmpty": pw.PrettyJsonString() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "PrettyJsonString", actual)
}

func Test_Cov14_PayloadWrapper_JsonString(t *testing.T) {
	pw := corepayload.NewPayloadWrapper.From("n", "id", "e", []byte(`{}`))
	actual := args.Map{"notEmpty": pw.JsonString() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "JsonString", actual)
}

func Test_Cov14_PayloadWrapper_MarshalJSON(t *testing.T) {
	pw := corepayload.NewPayloadWrapper.From("n", "id", "e", []byte(`{}`))
	b, err := pw.MarshalJSON()
	actual := args.Map{"hasBytes": len(b) > 0, "noErr": err == nil}
	expected := args.Map{"hasBytes": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "MarshalJSON", actual)
}

func Test_Cov14_PayloadWrapper_Serialize(t *testing.T) {
	pw := corepayload.NewPayloadWrapper.From("n", "id", "e", []byte(`{}`))
	b, err := pw.Serialize()
	actual := args.Map{"hasBytes": len(b) > 0, "noErr": err == nil}
	expected := args.Map{"hasBytes": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize", actual)
}

func Test_Cov14_PayloadWrapper_Clone(t *testing.T) {
	pw := corepayload.NewPayloadWrapper.From("n", "id", "e", []byte(`{}`))
	c := pw.Clone(true)
	actual := args.Map{"notNil": c != nil, "name": c.Name}
	expected := args.Map{"notNil": true, "name": "n"}
	expected.ShouldBeEqual(t, 0, "Clone", actual)
}

func Test_Cov14_PayloadWrapper_Clone_Nil(t *testing.T) {
	var pw *corepayload.PayloadWrapper
	c := pw.Clone(true)
	actual := args.Map{"nil": c == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "Clone nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// PayloadWrapper — Attributes
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_PayloadWrapper_Attributes(t *testing.T) {
	pw := corepayload.NewPayloadWrapper.From("n", "id", "e", []byte(`{}`))
	a := pw.InitializeAttributesOnNull()
	a.Set("k", "v")
	actual := args.Map{"notNil": pw.Attributes() != nil, "has": pw.Attributes().Has("k")}
	expected := args.Map{"notNil": true, "has": true}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper Attributes", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Generic helpers — typed_collection_funcs / typed_collection_paging
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_TypedPayloadCollection_Json(t *testing.T) {
	tc := createTestCollection()
	r := tc.Json()
	actual := args.Map{"noErr": !r.HasError()}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TypedPayloadCollection Json", actual)
}

func Test_Cov14_TypedPayloadCollection_JsonPtr(t *testing.T) {
	tc := createTestCollection()
	r := tc.JsonPtr()
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TypedPayloadCollection JsonPtr", actual)
}

func Test_Cov14_TypedPayloadCollection_String(t *testing.T) {
	tc := createTestCollection()
	actual := args.Map{"notEmpty": tc.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TypedPayloadCollection String", actual)
}

func Test_Cov14_TypedPayloadCollection_PrettyJsonString(t *testing.T) {
	tc := createTestCollection()
	actual := args.Map{"notEmpty": tc.PrettyJsonString() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TypedPayloadCollection PrettyJsonString", actual)
}

func Test_Cov14_TypedPayloadCollection_MarshalJSON(t *testing.T) {
	tc := createTestCollection()
	b, err := tc.MarshalJSON()
	actual := args.Map{"hasBytes": len(b) > 0, "noErr": err == nil}
	expected := args.Map{"hasBytes": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "TypedPayloadCollection MarshalJSON", actual)
}

func createTestCollection() *corepayload.TypedPayloadCollection[testUser] {
	tc := corepayload.NewTypedPayloadCollection[testUser]()
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[testUser]("n", "id", "testUser", testUser{Name: "alice"})
	tc.Add(tw)
	return tc
}

// ══════════════════════════════════════════════════════════════════════════════
// PayloadWrapper — NewPayloadWrapperFromJsonResult
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_NewPayloadWrapper_FromInstruction(t *testing.T) {
	pw := corepayload.NewPayloadWrapper.From("n", "id", "e", []byte(`{}`))
	actual := args.Map{"name": pw.Name}
	expected := args.Map{"name": "n"}
	expected.ShouldBeEqual(t, 0, "NewPayloadWrapper From", actual)
}

func Test_Cov14_NewPayloadWrapper_FromJsonResult(t *testing.T) {
	r := corejson.New(map[string]string{"k": "v"})
	pw := corepayload.NewPayloadWrapper.FromJsonResult("n", "id", "e", &r)
	actual := args.Map{"name": pw.Name, "hasPayloads": pw.HasPayloads()}
	expected := args.Map{"name": "n", "hasPayloads": true}
	expected.ShouldBeEqual(t, 0, "NewPayloadWrapper FromJsonResult", actual)
}
