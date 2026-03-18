package corepayloadtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corepayload"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Attributes
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_Attributes_Empty(t *testing.T) {
	a := corepayload.New.Attributes.Empty()
	actual := args.Map{"empty": a.IsEmpty(), "len": a.Length()}
	expected := args.Map{"empty": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "Attributes empty", actual)
}

func Test_Cov14_Attributes_Nil(t *testing.T) {
	var a *corepayload.Attributes
	actual := args.Map{"empty": a.IsEmpty(), "len": a.Length()}
	expected := args.Map{"empty": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "Attributes nil", actual)
}

func Test_Cov14_Attributes_KeyValues(t *testing.T) {
	a := corepayload.New.Attributes.Empty()
	a.KeyValuePairs.AddOrUpdate("key", "val")
	v, ok := a.KeyValuePairs.Value("key")
	actual := args.Map{"val": v, "ok": ok, "len": a.StringKeyValuePairsLength()}
	expected := args.Map{"val": "val", "ok": true, "len": 1}
	expected.ShouldBeEqual(t, 0, "Attributes KeyValues", actual)
}

func Test_Cov14_Attributes_GetStringKeyValue(t *testing.T) {
	a := corepayload.New.Attributes.Empty()
	a.KeyValuePairs.AddOrUpdate("key", "val")
	v, found := a.GetStringKeyValue("key")
	_, notFound := a.GetStringKeyValue("miss")
	actual := args.Map{"val": v, "found": found, "notFound": notFound}
	expected := args.Map{"val": "val", "found": true, "notFound": false}
	expected.ShouldBeEqual(t, 0, "GetStringKeyValue", actual)
}

func Test_Cov14_Attributes_Clear(t *testing.T) {
	a := corepayload.New.Attributes.Empty()
	a.KeyValuePairs.AddOrUpdate("key", "val")
	a.Clear()
	actual := args.Map{"empty": a.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Attributes Clear", actual)
}

func Test_Cov14_Attributes_Json(t *testing.T) {
	a := corepayload.New.Attributes.Empty()
	a.KeyValuePairs.AddOrUpdate("k", "v")
	r := a.Json()
	actual := args.Map{"noErr": !r.HasError()}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Attributes Json", actual)
}

func Test_Cov14_Attributes_IsEqual_BothNil(t *testing.T) {
	var a1, a2 *corepayload.Attributes
	actual := args.Map{"eq": a1.IsEqual(a2)}
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "Attributes IsEqual both nil", actual)
}

func Test_Cov14_Attributes_IsEqual_OneNil(t *testing.T) {
	a1 := corepayload.New.Attributes.Empty()
	actual := args.Map{"eq": a1.IsEqual(nil)}
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "Attributes IsEqual one nil", actual)
}

func Test_Cov14_Attributes_IsEmptyError(t *testing.T) {
	a := corepayload.New.Attributes.Empty()
	actual := args.Map{"emptyErr": a.IsEmptyError()}
	expected := args.Map{"emptyErr": true}
	expected.ShouldBeEqual(t, 0, "Attributes IsEmptyError", actual)
}

func Test_Cov14_Attributes_DynamicBytesLength(t *testing.T) {
	a := corepayload.New.Attributes.Empty()
	actual := args.Map{"len": a.DynamicBytesLength()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Attributes DynamicBytesLength", actual)
}

func Test_Cov14_Attributes_AnyKeyValuePairsLength(t *testing.T) {
	a := corepayload.New.Attributes.Empty()
	actual := args.Map{"len": a.AnyKeyValuePairsLength()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Attributes AnyKeyValuePairsLength", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// AuthInfo / SessionInfo / UserInfo
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_AuthInfo_Empty(t *testing.T) {
	a := corepayload.AuthInfo{}
	actual := args.Map{"id": a.Identifier}
	expected := args.Map{"id": ""}
	expected.ShouldBeEqual(t, 0, "AuthInfo empty", actual)
}

func Test_Cov14_SessionInfo_Empty(t *testing.T) {
	s := corepayload.SessionInfo{}
	actual := args.Map{"id": s.Id}
	expected := args.Map{"id": ""}
	expected.ShouldBeEqual(t, 0, "SessionInfo empty", actual)
}

func Test_Cov14_UserInfo_Empty(t *testing.T) {
	u := corepayload.UserInfo{}
	actual := args.Map{"nilUser": u.User == nil}
	expected := args.Map{"nilUser": true}
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
// PayloadWrapper — core methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_PayloadWrapper_Nil(t *testing.T) {
	var pw *corepayload.PayloadWrapper
	actual := args.Map{
		"isNull":  pw.IsNull(),
		"isEmpty": pw.IsEmpty(),
		"len":     pw.Length(),
		"idStr":   pw.IdString(),
	}
	expected := args.Map{
		"isNull": true, "isEmpty": true, "len": 0, "idStr": "",
	}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper nil", actual)
}

func Test_Cov14_PayloadWrapper_UsingBytes(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.UsingBytes("n", "id1", "task", "cat", "entity", []byte(`{}`))
	actual := args.Map{
		"name":     pw.Name,
		"id":       pw.Identifier,
		"notEmpty": !pw.IsEmpty(),
		"hasItems": pw.HasItems(),
		"idStr":    pw.IdString(),
	}
	expected := args.Map{
		"name": "n", "id": "id1", "notEmpty": true, "hasItems": true, "idStr": "id1",
	}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper UsingBytes", actual)
}

func Test_Cov14_PayloadWrapper_PayloadName(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.UsingBytes("n", "id", "t", "c", "e", []byte(`{}`))
	actual := args.Map{"name": pw.PayloadName()}
	expected := args.Map{"name": "n"}
	expected.ShouldBeEqual(t, 0, "PayloadName", actual)
}

func Test_Cov14_PayloadWrapper_IdInteger(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.UsingBytes("n", "42", "t", "c", "e", []byte(`{}`))
	actual := args.Map{"v": pw.IdInteger()}
	expected := args.Map{"v": 42}
	expected.ShouldBeEqual(t, 0, "IdInteger", actual)
}

func Test_Cov14_PayloadWrapper_IdInteger_Invalid(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.UsingBytes("n", "abc", "t", "c", "e", []byte(`{}`))
	actual := args.Map{"v": pw.IdInteger()}
	expected := args.Map{"v": -1}
	expected.ShouldBeEqual(t, 0, "IdInteger invalid", actual)
}

func Test_Cov14_PayloadWrapper_HasSafeItems(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.UsingBytes("n", "id", "t", "c", "e", []byte(`{}`))
	actual := args.Map{"v": pw.HasSafeItems()}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "HasSafeItems", actual)
}

func Test_Cov14_PayloadWrapper_JsonString(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.UsingBytes("n", "id", "t", "c", "e", []byte(`{}`))
	actual := args.Map{"notEmpty": pw.JsonString() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "JsonString", actual)
}

func Test_Cov14_PayloadWrapper_PrettyJsonString(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.UsingBytes("n", "id", "t", "c", "e", []byte(`{}`))
	actual := args.Map{"notEmpty": pw.PrettyJsonString() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "PrettyJsonString", actual)
}

func Test_Cov14_PayloadWrapper_String(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.UsingBytes("n", "id", "t", "c", "e", []byte(`{}`))
	actual := args.Map{"notEmpty": pw.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "String", actual)
}

func Test_Cov14_PayloadWrapper_String_Nil(t *testing.T) {
	var pw *corepayload.PayloadWrapper
	actual := args.Map{"v": pw.String()}
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "String nil", actual)
}

func Test_Cov14_PayloadWrapper_MarshalJSON(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.UsingBytes("n", "id", "t", "c", "e", []byte(`{}`))
	b, err := pw.MarshalJSON()
	actual := args.Map{"hasBytes": len(b) > 0, "noErr": err == nil}
	expected := args.Map{"hasBytes": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "MarshalJSON", actual)
}

func Test_Cov14_PayloadWrapper_Serialize(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.UsingBytes("n", "id", "t", "c", "e", []byte(`{}`))
	b, err := pw.Serialize()
	actual := args.Map{"hasBytes": len(b) > 0, "noErr": err == nil}
	expected := args.Map{"hasBytes": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize", actual)
}

func Test_Cov14_PayloadWrapper_HasError_NoAttr(t *testing.T) {
	pw := &corepayload.PayloadWrapper{}
	actual := args.Map{"v": pw.HasError()}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "HasError no attr", actual)
}

func Test_Cov14_PayloadWrapper_HasError_Nil(t *testing.T) {
	var pw *corepayload.PayloadWrapper
	actual := args.Map{"v": pw.HasError()}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "HasError nil", actual)
}

func Test_Cov14_PayloadWrapper_InitializeAttributesOnNull(t *testing.T) {
	pw := &corepayload.PayloadWrapper{}
	a := pw.InitializeAttributesOnNull()
	actual := args.Map{"notNil": a != nil, "hasAttr": pw.HasAttributes()}
	expected := args.Map{"notNil": true, "hasAttr": true}
	expected.ShouldBeEqual(t, 0, "InitializeAttributesOnNull", actual)
}

func Test_Cov14_PayloadWrapper_PayloadsString(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.UsingBytes("n", "id", "t", "c", "e", []byte(`{"a":1}`))
	actual := args.Map{"v": pw.PayloadsString()}
	expected := args.Map{"v": `{"a":1}`}
	expected.ShouldBeEqual(t, 0, "PayloadsString", actual)
}

func Test_Cov14_PayloadWrapper_Json_Nil(t *testing.T) {
	var pw *corepayload.PayloadWrapper
	r := pw.Json()
	actual := args.Map{"empty": r.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Json nil", actual)
}

func Test_Cov14_PayloadWrapper_Json_Valid(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.UsingBytes("n", "id", "t", "c", "e", []byte(`{}`))
	r := pw.Json()
	actual := args.Map{"noErr": !r.HasError()}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Json valid", actual)
}

func Test_Cov14_PayloadWrapper_JsonPtr(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.UsingBytes("n", "id", "t", "c", "e", []byte(`{}`))
	r := pw.JsonPtr()
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "JsonPtr", actual)
}

func Test_Cov14_PayloadWrapper_Clone(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.UsingBytes("n", "id", "t", "c", "e", []byte(`{}`))
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

func Test_Cov14_PayloadWrapper_Clear(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.UsingBytes("n", "id", "t", "c", "e", []byte(`{}`))
	pw.Clear()
	actual := args.Map{"empty": pw.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Clear", actual)
}

func Test_Cov14_PayloadWrapper_Dispose(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.UsingBytes("n", "id", "t", "c", "e", []byte(`{}`))
	pw.Dispose()
	actual := args.Map{"null": pw.IsNull()}
	expected := args.Map{"null": true}
	expected.ShouldBeEqual(t, 0, "Dispose", actual)
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
	expected := args.Map{"isNull": true, "isEmpty": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection nil", actual)
}

func Test_Cov14_PayloadsCollection_Empty(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.Empty()
	actual := args.Map{"empty": pc.IsEmpty(), "len": pc.Length()}
	expected := args.Map{"empty": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection empty", actual)
}

func Test_Cov14_PayloadsCollection_Add(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.Empty()
	pw := corepayload.New.PayloadWrapper.UsingBytes("n", "id", "t", "c", "e", []byte(`{}`))
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
		Name: "n", Identifier: "id", EntityType: "e", Payloads: []byte(`{}`),
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
		Name: "n", Identifier: "id", EntityType: "e",
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
	u := corepayload.New.User.Empty()
	actual := args.Map{"empty": u.IsEmpty(), "id": u.Identifier}
	expected := args.Map{"empty": true, "id": ""}
	expected.ShouldBeEqual(t, 0, "User empty", actual)
}

func Test_Cov14_User_NonSysCreate(t *testing.T) {
	u := corepayload.New.User.NonSysCreate("alice", "admin")
	actual := args.Map{"name": u.Name, "type": u.Type, "sys": u.IsSystemUser}
	expected := args.Map{"name": "alice", "type": "admin", "sys": false}
	expected.ShouldBeEqual(t, 0, "User NonSysCreate", actual)
}

func Test_Cov14_User_Nil(t *testing.T) {
	var u *corepayload.User
	actual := args.Map{"isEmpty": u.IsEmpty(), "nameEmpty": u.IsNameEmpty()}
	expected := args.Map{"isEmpty": true, "nameEmpty": true}
	expected.ShouldBeEqual(t, 0, "User nil", actual)
}

func Test_Cov14_User_String(t *testing.T) {
	u := corepayload.New.User.NonSysCreate("alice", "admin")
	actual := args.Map{"notEmpty": u.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "User String", actual)
}

func Test_Cov14_User_Clone(t *testing.T) {
	u := corepayload.New.User.NonSysCreate("alice", "admin")
	c := u.Clone()
	actual := args.Map{"name": c.Name}
	expected := args.Map{"name": "alice"}
	expected.ShouldBeEqual(t, 0, "User Clone", actual)
}

func Test_Cov14_User_ClonePtr(t *testing.T) {
	u := corepayload.New.User.NonSysCreate("alice", "admin")
	c := u.ClonePtr()
	actual := args.Map{"notNil": c != nil, "name": c.Name}
	expected := args.Map{"notNil": true, "name": "alice"}
	expected.ShouldBeEqual(t, 0, "User ClonePtr", actual)
}

func Test_Cov14_User_ClonePtr_Nil(t *testing.T) {
	var u *corepayload.User
	c := u.ClonePtr()
	actual := args.Map{"nil": c == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "User ClonePtr nil", actual)
}

func Test_Cov14_User_NonSysCreateId(t *testing.T) {
	u := corepayload.New.User.NonSysCreateId("42", "alice", "admin")
	actual := args.Map{"id": u.Identifier, "v": u.IdentifierInteger()}
	expected := args.Map{"id": "42", "v": 42}
	expected.ShouldBeEqual(t, 0, "User NonSysCreateId", actual)
}

func Test_Cov14_User_Checkers(t *testing.T) {
	u := corepayload.New.User.NonSysCreate("alice", "admin")
	actual := args.Map{
		"valid":    u.IsValidUser(),
		"notSys":   u.IsNotSystemUser(),
		"hasType":  u.HasType(),
		"nameEq":   u.IsNameEqual("alice"),
		"nameNEq":  u.IsNameEqual("bob"),
		"noAuth":   u.IsAuthTokenEmpty(),
		"noPwd":    u.IsPasswordHashEmpty(),
	}
	expected := args.Map{
		"valid": true, "notSys": true, "hasType": true,
		"nameEq": true, "nameNEq": false,
		"noAuth": true, "noPwd": true,
	}
	expected.ShouldBeEqual(t, 0, "User checkers", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// TypedPayloadCollection
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_TypedPayloadCollection_Empty(t *testing.T) {
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	actual := args.Map{"empty": tc.IsEmpty(), "len": tc.Length()}
	expected := args.Map{"empty": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "TypedPayloadCollection empty", actual)
}

func Test_Cov14_TypedPayloadCollection_Add(t *testing.T) {
	tc := corepayload.NewTypedPayloadCollection[testUser](10)
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[testUser]("n", "id", "testUser", testUser{Name: "alice"})
	tc.Add(tw)
	actual := args.Map{"len": tc.Length(), "hasAny": tc.HasAnyItem()}
	expected := args.Map{"len": 1, "hasAny": true}
	expected.ShouldBeEqual(t, 0, "TypedPayloadCollection add", actual)
}

func cov14CreateTestCollection() *corepayload.TypedPayloadCollection[testUser] {
	tc := corepayload.NewTypedPayloadCollection[testUser](10)
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[testUser]("n", "id", "testUser", testUser{Name: "alice"})
	tc.Add(tw)
	return tc
}

func Test_Cov14_TypedPayloadCollection_Json(t *testing.T) {
	tc := cov14CreateTestCollection()
	r := tc.Json()
	actual := args.Map{"noErr": !r.HasError()}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TypedPayloadCollection Json", actual)
}

func Test_Cov14_TypedPayloadCollection_JsonPtr(t *testing.T) {
	tc := cov14CreateTestCollection()
	r := tc.JsonPtr()
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TypedPayloadCollection JsonPtr", actual)
}

func Test_Cov14_TypedPayloadCollection_String(t *testing.T) {
	tc := cov14CreateTestCollection()
	actual := args.Map{"notEmpty": tc.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TypedPayloadCollection String", actual)
}

func Test_Cov14_TypedPayloadCollection_PrettyJsonString(t *testing.T) {
	tc := cov14CreateTestCollection()
	actual := args.Map{"notEmpty": tc.PrettyJsonString() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TypedPayloadCollection PrettyJsonString", actual)
}

func Test_Cov14_TypedPayloadCollection_MarshalJSON(t *testing.T) {
	tc := cov14CreateTestCollection()
	b, err := tc.MarshalJSON()
	actual := args.Map{"hasBytes": len(b) > 0, "noErr": err == nil}
	expected := args.Map{"hasBytes": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "TypedPayloadCollection MarshalJSON", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// PayloadWrapper — NewPayloadWrapper from various creators
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_NewPayloadWrapper_Empty(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	actual := args.Map{"empty": pw.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "NewPayloadWrapper Empty", actual)
}

func Test_Cov14_NewPayloadWrapper_Create(t *testing.T) {
	pw, err := corepayload.New.PayloadWrapper.Create("n", "id", "task", "cat", map[string]string{"k": "v"})
	actual := args.Map{"notNil": pw != nil, "noErr": err == nil, "name": pw.Name}
	expected := args.Map{"notNil": true, "noErr": true, "name": "n"}
	expected.ShouldBeEqual(t, 0, "NewPayloadWrapper Create", actual)
}

func Test_Cov14_NewPayloadWrapper_Deserialize(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.UsingBytes("n", "id", "t", "c", "e", []byte(`{}`))
	b, _ := pw.Serialize()
	pw2, err := corepayload.New.PayloadWrapper.Deserialize(b)
	actual := args.Map{"notNil": pw2 != nil, "noErr": err == nil}
	expected := args.Map{"notNil": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "NewPayloadWrapper Deserialize", actual)
}

func Test_Cov14_NewPayloadWrapper_DeserializeUsingJsonResult(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.UsingBytes("n", "id", "t", "c", "e", []byte(`{}`))
	r := pw.JsonPtr()
	pw2, err := corepayload.New.PayloadWrapper.DeserializeUsingJsonResult(r)
	actual := args.Map{"notNil": pw2 != nil, "noErr": err == nil}
	expected := args.Map{"notNil": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "NewPayloadWrapper DeserializeUsingJsonResult", actual)
}

func Test_Cov14_NewPayloadWrapper_UsingBytesCreateInstruction(t *testing.T) {
	bci := &corepayload.BytesCreateInstruction{
		Name: "n", Identifier: "id", EntityType: "e", Payloads: []byte(`{}`),
	}
	pw := corepayload.New.PayloadWrapper.UsingBytesCreateInstruction(bci)
	actual := args.Map{"name": pw.Name}
	expected := args.Map{"name": "n"}
	expected.ShouldBeEqual(t, 0, "UsingBytesCreateInstruction", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// New.Attributes various creators
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_Attributes_UsingBasicError(t *testing.T) {
	a := corepayload.New.Attributes.UsingBasicError(nil)
	actual := args.Map{"notNil": a != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Attributes UsingBasicError", actual)
}

func Test_Cov14_Attributes_UsingAuthInfo(t *testing.T) {
	auth := &corepayload.AuthInfo{Identifier: "test"}
	a := corepayload.New.Attributes.UsingAuthInfo(auth)
	actual := args.Map{"notNil": a != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Attributes UsingAuthInfo", actual)
}

func Test_Cov14_Attributes_UsingDynamicPayloadBytes(t *testing.T) {
	a := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte(`{}`))
	actual := args.Map{"dynLen": a.DynamicBytesLength()}
	expected := args.Map{"dynLen": 2}
	expected.ShouldBeEqual(t, 0, "Attributes UsingDynamicPayloadBytes", actual)
}

func Test_Cov14_Attributes_Deserialize(t *testing.T) {
	a := corepayload.New.Attributes.Empty()
	b, _ := corejson.Serialize.Raw(a)
	a2, err := corepayload.New.Attributes.Deserialize(b)
	actual := args.Map{"notNil": a2 != nil, "noErr": err == nil}
	expected := args.Map{"notNil": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "Attributes Deserialize", actual)
}
