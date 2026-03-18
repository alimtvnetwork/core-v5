package corepayloadtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corepayload"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// AttributesGetters — All methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov11_Attributes_IsNull(t *testing.T) {
	var a *corepayload.Attributes
	a2 := &corepayload.Attributes{}
	actual := args.Map{"nil": a.IsNull(), "notNil": a2.IsNull()}
	expected := args.Map{"nil": true, "notNil": false}
	expected.ShouldBeEqual(t, 0, "IsNull", actual)
}

func Test_Cov11_Attributes_HasSafeItems(t *testing.T) {
	a := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("data"))
	actual := args.Map{"val": a.HasSafeItems()}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "HasSafeItems", actual)
}

func Test_Cov11_Attributes_HasStringKey(t *testing.T) {
	a := corepayload.New.Attributes.Empty()
	a.KeyValuePairs.AddOrUpdate("k", "v")
	actual := args.Map{"has": a.HasStringKey("k"), "no": a.HasStringKey("x")}
	expected := args.Map{"has": true, "no": false}
	expected.ShouldBeEqual(t, 0, "HasStringKey", actual)
}

func Test_Cov11_Attributes_HasStringKey_NoKV(t *testing.T) {
	a := &corepayload.Attributes{}
	actual := args.Map{"val": a.HasStringKey("k")}
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "HasStringKey no kv", actual)
}

func Test_Cov11_Attributes_HasAnyKey(t *testing.T) {
	a := corepayload.New.Attributes.Empty()
	a.AnyKeyValuePairs.Add("k", 42)
	actual := args.Map{"has": a.HasAnyKey("k"), "no": a.HasAnyKey("x")}
	expected := args.Map{"has": true, "no": false}
	expected.ShouldBeEqual(t, 0, "HasAnyKey", actual)
}

func Test_Cov11_Attributes_HasAnyKey_NoAKV(t *testing.T) {
	a := &corepayload.Attributes{}
	actual := args.Map{"val": a.HasAnyKey("k")}
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "HasAnyKey no akv", actual)
}

func Test_Cov11_Attributes_Payloads_Empty(t *testing.T) {
	var a *corepayload.Attributes
	actual := args.Map{"len": len(a.Payloads())}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Payloads empty", actual)
}

func Test_Cov11_Attributes_Payloads_Valid(t *testing.T) {
	a := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("data"))
	actual := args.Map{"len": len(a.Payloads())}
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "Payloads valid", actual)
}

func Test_Cov11_Attributes_PayloadsString_Empty(t *testing.T) {
	var a *corepayload.Attributes
	actual := args.Map{"val": a.PayloadsString()}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "PayloadsString empty", actual)
}

func Test_Cov11_Attributes_PayloadsString_Valid(t *testing.T) {
	a := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("hello"))
	actual := args.Map{"val": a.PayloadsString()}
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "PayloadsString valid", actual)
}

func Test_Cov11_Attributes_AnyKeyValMap_Empty(t *testing.T) {
	var a *corepayload.Attributes
	m := a.AnyKeyValMap()
	actual := args.Map{"len": len(m)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyKeyValMap empty", actual)
}

func Test_Cov11_Attributes_Hashmap_Empty(t *testing.T) {
	var a *corepayload.Attributes
	m := a.Hashmap()
	actual := args.Map{"len": len(m)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Hashmap empty", actual)
}

func Test_Cov11_Attributes_CompiledError(t *testing.T) {
	a := corepayload.New.Attributes.Empty()
	actual := args.Map{"noErr": a.CompiledError() == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "CompiledError", actual)
}

func Test_Cov11_Attributes_HasIssuesOrEmpty(t *testing.T) {
	a := corepayload.New.Attributes.Empty()
	actual := args.Map{"val": a.HasIssuesOrEmpty()}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "HasIssuesOrEmpty empty", actual)
}

func Test_Cov11_Attributes_IsSafeValid(t *testing.T) {
	a := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("data"))
	actual := args.Map{"val": a.IsSafeValid()}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "IsSafeValid", actual)
}

func Test_Cov11_Attributes_HasAnyItem(t *testing.T) {
	a := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("data"))
	actual := args.Map{"val": a.HasAnyItem()}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "HasAnyItem", actual)
}

func Test_Cov11_Attributes_Count(t *testing.T) {
	a := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("data"))
	actual := args.Map{"val": a.Count()}
	expected := args.Map{"val": 4}
	expected.ShouldBeEqual(t, 0, "Count", actual)
}

func Test_Cov11_Attributes_Capacity(t *testing.T) {
	a := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("data"))
	actual := args.Map{"val": a.Capacity()}
	expected := args.Map{"val": 4}
	expected.ShouldBeEqual(t, 0, "Capacity", actual)
}

func Test_Cov11_Attributes_Length_Nil(t *testing.T) {
	var a *corepayload.Attributes
	actual := args.Map{"val": a.Length()}
	expected := args.Map{"val": 0}
	expected.ShouldBeEqual(t, 0, "Length nil", actual)
}

func Test_Cov11_Attributes_HasPagingInfo(t *testing.T) {
	a := &corepayload.Attributes{PagingInfo: &corepayload.PagingInfo{TotalPages: 5}}
	var a2 *corepayload.Attributes
	actual := args.Map{"has": a.HasPagingInfo(), "nil": a2.HasPagingInfo()}
	expected := args.Map{"has": true, "nil": false}
	expected.ShouldBeEqual(t, 0, "HasPagingInfo", actual)
}

func Test_Cov11_Attributes_HasKeyValuePairs(t *testing.T) {
	a := corepayload.New.Attributes.Empty()
	a.KeyValuePairs.AddOrUpdate("k", "v")
	var a2 *corepayload.Attributes
	actual := args.Map{"has": a.HasKeyValuePairs(), "nil": a2.HasKeyValuePairs()}
	expected := args.Map{"has": true, "nil": false}
	expected.ShouldBeEqual(t, 0, "HasKeyValuePairs", actual)
}

func Test_Cov11_Attributes_HasFromTo(t *testing.T) {
	a := &corepayload.Attributes{}
	var a2 *corepayload.Attributes
	actual := args.Map{"no": a.HasFromTo(), "nil": a2.HasFromTo()}
	expected := args.Map{"no": false, "nil": false}
	expected.ShouldBeEqual(t, 0, "HasFromTo", actual)
}

func Test_Cov11_Attributes_IsValid(t *testing.T) {
	a := corepayload.New.Attributes.Empty()
	var a2 *corepayload.Attributes
	actual := args.Map{"valid": a.IsValid(), "nil": a2.IsValid()}
	expected := args.Map{"valid": true, "nil": false}
	expected.ShouldBeEqual(t, 0, "IsValid", actual)
}

func Test_Cov11_Attributes_IsInvalid(t *testing.T) {
	var a *corepayload.Attributes
	actual := args.Map{"val": a.IsInvalid()}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "IsInvalid nil", actual)
}

func Test_Cov11_Attributes_HasError(t *testing.T) {
	a := corepayload.New.Attributes.Empty()
	var a2 *corepayload.Attributes
	actual := args.Map{"no": a.HasError(), "nil": a2.HasError()}
	expected := args.Map{"no": false, "nil": false}
	expected.ShouldBeEqual(t, 0, "HasError", actual)
}

func Test_Cov11_Attributes_Error_Empty(t *testing.T) {
	a := corepayload.New.Attributes.Empty()
	actual := args.Map{"noErr": a.Error() == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Error empty", actual)
}

func Test_Cov11_Attributes_IsEmptyError_Nil(t *testing.T) {
	var a *corepayload.Attributes
	actual := args.Map{"val": a.IsEmptyError()}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "IsEmptyError nil", actual)
}

func Test_Cov11_Attributes_DynamicBytesLength_Nil(t *testing.T) {
	var a *corepayload.Attributes
	actual := args.Map{"val": a.DynamicBytesLength()}
	expected := args.Map{"val": 0}
	expected.ShouldBeEqual(t, 0, "DynamicBytesLength nil", actual)
}

func Test_Cov11_Attributes_StringKeyValuePairsLength_Nil(t *testing.T) {
	var a *corepayload.Attributes
	actual := args.Map{"val": a.StringKeyValuePairsLength()}
	expected := args.Map{"val": 0}
	expected.ShouldBeEqual(t, 0, "StringKeyValuePairsLength nil", actual)
}

func Test_Cov11_Attributes_AnyKeyValuePairsLength_Nil(t *testing.T) {
	var a *corepayload.Attributes
	actual := args.Map{"val": a.AnyKeyValuePairsLength()}
	expected := args.Map{"val": 0}
	expected.ShouldBeEqual(t, 0, "AnyKeyValuePairsLength nil", actual)
}

func Test_Cov11_Attributes_IsEmpty_Nil(t *testing.T) {
	var a *corepayload.Attributes
	actual := args.Map{"val": a.IsEmpty()}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "IsEmpty nil", actual)
}

func Test_Cov11_Attributes_HasItems(t *testing.T) {
	a := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("data"))
	actual := args.Map{"val": a.HasItems()}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "HasItems", actual)
}

func Test_Cov11_Attributes_IsPagingInfoEmpty(t *testing.T) {
	a := &corepayload.Attributes{}
	var a2 *corepayload.Attributes
	actual := args.Map{"empty": a.IsPagingInfoEmpty(), "nil": a2.IsPagingInfoEmpty()}
	expected := args.Map{"empty": true, "nil": true}
	expected.ShouldBeEqual(t, 0, "IsPagingInfoEmpty", actual)
}

func Test_Cov11_Attributes_IsKeyValuePairsEmpty(t *testing.T) {
	a := &corepayload.Attributes{}
	var a2 *corepayload.Attributes
	actual := args.Map{"empty": a.IsKeyValuePairsEmpty(), "nil": a2.IsKeyValuePairsEmpty()}
	expected := args.Map{"empty": true, "nil": true}
	expected.ShouldBeEqual(t, 0, "IsKeyValuePairsEmpty", actual)
}

func Test_Cov11_Attributes_IsAnyKeyValuePairsEmpty(t *testing.T) {
	a := &corepayload.Attributes{}
	var a2 *corepayload.Attributes
	actual := args.Map{"empty": a.IsAnyKeyValuePairsEmpty(), "nil": a2.IsAnyKeyValuePairsEmpty()}
	expected := args.Map{"empty": true, "nil": true}
	expected.ShouldBeEqual(t, 0, "IsAnyKeyValuePairsEmpty", actual)
}

func Test_Cov11_Attributes_IsUserInfoEmpty(t *testing.T) {
	a := &corepayload.Attributes{}
	var a2 *corepayload.Attributes
	actual := args.Map{"empty": a.IsUserInfoEmpty(), "nil": a2.IsUserInfoEmpty()}
	expected := args.Map{"empty": true, "nil": true}
	expected.ShouldBeEqual(t, 0, "IsUserInfoEmpty", actual)
}

func Test_Cov11_Attributes_VirtualUser_Empty(t *testing.T) {
	a := &corepayload.Attributes{}
	actual := args.Map{"nil": a.VirtualUser() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "VirtualUser empty", actual)
}

func Test_Cov11_Attributes_VirtualUser_Valid(t *testing.T) {
	a := &corepayload.Attributes{AuthInfo: &corepayload.AuthInfo{UserInfo: &corepayload.UserInfo{User: &corepayload.User{Name: "u"}}}}
	actual := args.Map{"name": a.VirtualUser().Name}
	expected := args.Map{"name": "u"}
	expected.ShouldBeEqual(t, 0, "VirtualUser valid", actual)
}

func Test_Cov11_Attributes_SystemUser_Empty(t *testing.T) {
	a := &corepayload.Attributes{}
	actual := args.Map{"nil": a.SystemUser() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "SystemUser empty", actual)
}

func Test_Cov11_Attributes_SystemUser_Valid(t *testing.T) {
	a := &corepayload.Attributes{AuthInfo: &corepayload.AuthInfo{UserInfo: &corepayload.UserInfo{SystemUser: &corepayload.User{Name: "su"}}}}
	actual := args.Map{"name": a.SystemUser().Name}
	expected := args.Map{"name": "su"}
	expected.ShouldBeEqual(t, 0, "SystemUser valid", actual)
}

func Test_Cov11_Attributes_SessionUser_Empty(t *testing.T) {
	a := &corepayload.Attributes{}
	actual := args.Map{"nil": a.SessionUser() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "SessionUser empty", actual)
}

func Test_Cov11_Attributes_SessionUser_Valid(t *testing.T) {
	a := &corepayload.Attributes{AuthInfo: &corepayload.AuthInfo{SessionInfo: &corepayload.SessionInfo{Id: "1", User: &corepayload.User{Name: "su"}}}}
	actual := args.Map{"name": a.SessionUser().Name}
	expected := args.Map{"name": "su"}
	expected.ShouldBeEqual(t, 0, "SessionUser valid", actual)
}

func Test_Cov11_Attributes_IsAuthInfoEmpty(t *testing.T) {
	a := &corepayload.Attributes{}
	var a2 *corepayload.Attributes
	actual := args.Map{"empty": a.IsAuthInfoEmpty(), "nil": a2.IsAuthInfoEmpty()}
	expected := args.Map{"empty": true, "nil": true}
	expected.ShouldBeEqual(t, 0, "IsAuthInfoEmpty", actual)
}

func Test_Cov11_Attributes_IsSessionInfoEmpty(t *testing.T) {
	a := &corepayload.Attributes{}
	var a2 *corepayload.Attributes
	actual := args.Map{"empty": a.IsSessionInfoEmpty(), "nil": a2.IsSessionInfoEmpty()}
	expected := args.Map{"empty": true, "nil": true}
	expected.ShouldBeEqual(t, 0, "IsSessionInfoEmpty", actual)
}

func Test_Cov11_Attributes_HasUserInfo(t *testing.T) {
	a := &corepayload.Attributes{AuthInfo: &corepayload.AuthInfo{UserInfo: &corepayload.UserInfo{User: &corepayload.User{Name: "u"}}}}
	actual := args.Map{"val": a.HasUserInfo()}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "HasUserInfo", actual)
}

func Test_Cov11_Attributes_HasAuthInfo(t *testing.T) {
	a := &corepayload.Attributes{AuthInfo: &corepayload.AuthInfo{ActionType: "login"}}
	actual := args.Map{"val": a.HasAuthInfo()}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "HasAuthInfo", actual)
}

func Test_Cov11_Attributes_HasSessionInfo(t *testing.T) {
	a := &corepayload.Attributes{AuthInfo: &corepayload.AuthInfo{SessionInfo: &corepayload.SessionInfo{Id: "1"}}}
	actual := args.Map{"val": a.HasSessionInfo()}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "HasSessionInfo", actual)
}

func Test_Cov11_Attributes_SessionInfo_Empty(t *testing.T) {
	a := &corepayload.Attributes{}
	actual := args.Map{"nil": a.SessionInfo() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "SessionInfo empty", actual)
}

func Test_Cov11_Attributes_SessionInfo_Valid(t *testing.T) {
	a := &corepayload.Attributes{AuthInfo: &corepayload.AuthInfo{SessionInfo: &corepayload.SessionInfo{Id: "1"}}}
	actual := args.Map{"id": a.SessionInfo().Id}
	expected := args.Map{"id": "1"}
	expected.ShouldBeEqual(t, 0, "SessionInfo valid", actual)
}

func Test_Cov11_Attributes_AuthType_Empty(t *testing.T) {
	a := &corepayload.Attributes{}
	actual := args.Map{"val": a.AuthType()}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "AuthType empty", actual)
}

func Test_Cov11_Attributes_AuthType_Valid(t *testing.T) {
	a := &corepayload.Attributes{AuthInfo: &corepayload.AuthInfo{ActionType: "login"}}
	actual := args.Map{"val": a.AuthType()}
	expected := args.Map{"val": "login"}
	expected.ShouldBeEqual(t, 0, "AuthType valid", actual)
}

func Test_Cov11_Attributes_ResourceName_Empty(t *testing.T) {
	a := &corepayload.Attributes{}
	actual := args.Map{"val": a.ResourceName()}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "ResourceName empty", actual)
}

func Test_Cov11_Attributes_ResourceName_Valid(t *testing.T) {
	a := &corepayload.Attributes{AuthInfo: &corepayload.AuthInfo{ResourceName: "res"}}
	actual := args.Map{"val": a.ResourceName()}
	expected := args.Map{"val": "res"}
	expected.ShouldBeEqual(t, 0, "ResourceName valid", actual)
}

func Test_Cov11_Attributes_HasStringKeyValuePairs(t *testing.T) {
	a := corepayload.New.Attributes.Empty()
	a.KeyValuePairs.AddOrUpdate("k", "v")
	actual := args.Map{"val": a.HasStringKeyValuePairs()}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "HasStringKeyValuePairs", actual)
}

func Test_Cov11_Attributes_HasAnyKeyValuePairs(t *testing.T) {
	a := corepayload.New.Attributes.Empty()
	a.AnyKeyValuePairs.Add("k", 42)
	actual := args.Map{"val": a.HasAnyKeyValuePairs()}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "HasAnyKeyValuePairs", actual)
}

func Test_Cov11_Attributes_HasDynamicPayloads(t *testing.T) {
	a := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("data"))
	actual := args.Map{"val": a.HasDynamicPayloads()}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "HasDynamicPayloads", actual)
}

func Test_Cov11_Attributes_GetStringKeyValue_Nil(t *testing.T) {
	var a *corepayload.Attributes
	v, found := a.GetStringKeyValue("k")
	actual := args.Map{"val": v, "found": found}
	expected := args.Map{"val": "", "found": false}
	expected.ShouldBeEqual(t, 0, "GetStringKeyValue nil", actual)
}

func Test_Cov11_Attributes_GetStringKeyValue_Valid(t *testing.T) {
	a := corepayload.New.Attributes.Empty()
	a.KeyValuePairs.AddOrUpdate("k", "v")
	v, found := a.GetStringKeyValue("k")
	actual := args.Map{"val": v, "found": found}
	expected := args.Map{"val": "v", "found": true}
	expected.ShouldBeEqual(t, 0, "GetStringKeyValue valid", actual)
}

func Test_Cov11_Attributes_GetAnyKeyValue_Nil(t *testing.T) {
	var a *corepayload.Attributes
	_, found := a.GetAnyKeyValue("k")
	actual := args.Map{"found": found}
	expected := args.Map{"found": false}
	expected.ShouldBeEqual(t, 0, "GetAnyKeyValue nil", actual)
}

func Test_Cov11_Attributes_IsErrorDifferent(t *testing.T) {
	a := corepayload.New.Attributes.Empty()
	actual := args.Map{"val": a.IsErrorDifferent(nil)}
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "IsErrorDifferent", actual)
}

func Test_Cov11_Attributes_IsErrorEqual(t *testing.T) {
	a := corepayload.New.Attributes.Empty()
	actual := args.Map{"val": a.IsErrorEqual(nil)}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "IsErrorEqual nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// AttributesSetters — All methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov11_Attributes_SetAuthInfo_Nil(t *testing.T) {
	var a *corepayload.Attributes
	result := a.SetAuthInfo(&corepayload.AuthInfo{ActionType: "x"})
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SetAuthInfo nil", actual)
}

func Test_Cov11_Attributes_SetAuthInfo_NonNil(t *testing.T) {
	a := corepayload.New.Attributes.Empty()
	result := a.SetAuthInfo(&corepayload.AuthInfo{ActionType: "x"})
	actual := args.Map{"same": result == a}
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "SetAuthInfo non-nil", actual)
}

func Test_Cov11_Attributes_SetUserInfo_Nil(t *testing.T) {
	var a *corepayload.Attributes
	result := a.SetUserInfo(&corepayload.UserInfo{})
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SetUserInfo nil", actual)
}

func Test_Cov11_Attributes_SetUserInfo_NonNil(t *testing.T) {
	a := &corepayload.Attributes{AuthInfo: &corepayload.AuthInfo{}}
	result := a.SetUserInfo(&corepayload.UserInfo{})
	actual := args.Map{"same": result == a}
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "SetUserInfo non-nil", actual)
}

func Test_Cov11_Attributes_AddNewStringKeyValueOnly_Nil(t *testing.T) {
	var a *corepayload.Attributes
	actual := args.Map{"val": a.AddNewStringKeyValueOnly("k", "v")}
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "AddNewStringKeyValueOnly nil", actual)
}

func Test_Cov11_Attributes_AddNewStringKeyValueOnly_Valid(t *testing.T) {
	a := corepayload.New.Attributes.Empty()
	actual := args.Map{"val": a.AddNewStringKeyValueOnly("k", "v")}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "AddNewStringKeyValueOnly valid", actual)
}

func Test_Cov11_Attributes_AddNewAnyKeyValueOnly_Nil(t *testing.T) {
	var a *corepayload.Attributes
	actual := args.Map{"val": a.AddNewAnyKeyValueOnly("k", 42)}
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "AddNewAnyKeyValueOnly nil", actual)
}

func Test_Cov11_Attributes_AddNewAnyKeyValueOnly_Valid(t *testing.T) {
	a := corepayload.New.Attributes.Empty()
	actual := args.Map{"val": a.AddNewAnyKeyValueOnly("k", 42)}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "AddNewAnyKeyValueOnly valid", actual)
}

func Test_Cov11_Attributes_AddOrUpdateString_Nil(t *testing.T) {
	var a *corepayload.Attributes
	actual := args.Map{"val": a.AddOrUpdateString("k", "v")}
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "AddOrUpdateString nil", actual)
}

func Test_Cov11_Attributes_AddOrUpdateAnyItem_Nil(t *testing.T) {
	var a *corepayload.Attributes
	actual := args.Map{"val": a.AddOrUpdateAnyItem("k", 42)}
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "AddOrUpdateAnyItem nil", actual)
}

func Test_Cov11_Attributes_SetBasicErr_Nil(t *testing.T) {
	var a *corepayload.Attributes
	result := a.SetBasicErr(nil)
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SetBasicErr nil", actual)
}

func Test_Cov11_Attributes_Clear_Nil(t *testing.T) {
	var a *corepayload.Attributes
	a.Clear() // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Clear nil", actual)
}

func Test_Cov11_Attributes_Clear_Valid(t *testing.T) {
	a := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("data"))
	a.Clear()
	actual := args.Map{"len": len(a.DynamicPayloads)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Clear valid", actual)
}

func Test_Cov11_Attributes_Dispose_Nil(t *testing.T) {
	var a *corepayload.Attributes
	a.Dispose() // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Dispose nil", actual)
}

func Test_Cov11_Attributes_HandleErr_NoError(t *testing.T) {
	a := corepayload.New.Attributes.Empty()
	a.HandleErr() // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "HandleErr no error", actual)
}

func Test_Cov11_Attributes_HandleError_NoError(t *testing.T) {
	a := corepayload.New.Attributes.Empty()
	a.HandleError() // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "HandleError no error", actual)
}

func Test_Cov11_Attributes_MustBeEmptyError_Empty(t *testing.T) {
	a := corepayload.New.Attributes.Empty()
	a.MustBeEmptyError() // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "MustBeEmptyError", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Attributes — NewAttributesCreator
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov11_NewAttributes_Empty(t *testing.T) {
	a := corepayload.New.Attributes.Empty()
	actual := args.Map{"notNil": a != nil, "valid": a.IsValid()}
	expected := args.Map{"notNil": true, "valid": true}
	expected.ShouldBeEqual(t, 0, "NewAttributes.Empty", actual)
}

func Test_Cov11_NewAttributes_UsingDynamicPayloadBytes(t *testing.T) {
	a := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("data"))
	actual := args.Map{"len": len(a.DynamicPayloads)}
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "NewAttributes.UsingDynamicPayloadBytes", actual)
}

func Test_Cov11_NewAttributes_UsingAuthInfo(t *testing.T) {
	ai := &corepayload.AuthInfo{ActionType: "login"}
	a := corepayload.New.Attributes.UsingAuthInfo(ai)
	actual := args.Map{"action": a.AuthInfo.ActionType}
	expected := args.Map{"action": "login"}
	expected.ShouldBeEqual(t, 0, "NewAttributes.UsingAuthInfo", actual)
}

func Test_Cov11_NewAttributes_UsingAuthInfoDynamicBytes(t *testing.T) {
	a := corepayload.New.Attributes.UsingAuthInfoDynamicBytes(&corepayload.AuthInfo{ActionType: "x"}, []byte("data"))
	actual := args.Map{"action": a.AuthInfo.ActionType, "len": len(a.DynamicPayloads)}
	expected := args.Map{"action": "x", "len": 4}
	expected.ShouldBeEqual(t, 0, "NewAttributes.UsingAuthInfoDynamicBytes", actual)
}

func Test_Cov11_NewAttributes_Deserialize_Valid(t *testing.T) {
	a := corepayload.New.Attributes.Empty()
	b := a.JsonPtr().Bytes
	a2, err := corepayload.New.Attributes.Deserialize(b)
	actual := args.Map{"notNil": a2 != nil, "noErr": err == nil}
	expected := args.Map{"notNil": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "NewAttributes.Deserialize", actual)
}

func Test_Cov11_NewAttributes_Deserialize_Bad(t *testing.T) {
	_, err := corepayload.New.Attributes.Deserialize([]byte("{bad"))
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NewAttributes.Deserialize bad", actual)
}

func Test_Cov11_NewAttributes_DeserializeMany_Bad(t *testing.T) {
	_, err := corepayload.New.Attributes.DeserializeMany([]byte("{bad"))
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NewAttributes.DeserializeMany bad", actual)
}

func Test_Cov11_NewAttributes_CastOrDeserializeFrom_Nil(t *testing.T) {
	_, err := corepayload.New.Attributes.CastOrDeserializeFrom(nil)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NewAttributes.CastOrDeserializeFrom nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// EmptyCreator
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov11_Empty_Attributes(t *testing.T) {
	a := corepayload.Empty.Attributes()
	actual := args.Map{"notNil": a != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Empty.Attributes", actual)
}

func Test_Cov11_Empty_AttributesDefaults(t *testing.T) {
	a := corepayload.Empty.AttributesDefaults()
	actual := args.Map{"notNil": a != nil, "hasKV": a.KeyValuePairs != nil, "hasAKV": a.AnyKeyValuePairs != nil}
	expected := args.Map{"notNil": true, "hasKV": true, "hasAKV": true}
	expected.ShouldBeEqual(t, 0, "Empty.AttributesDefaults", actual)
}

func Test_Cov11_Empty_PayloadWrapper(t *testing.T) {
	pw := corepayload.Empty.PayloadWrapper()
	actual := args.Map{"notNil": pw != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Empty.PayloadWrapper", actual)
}

func Test_Cov11_Empty_PayloadsCollection(t *testing.T) {
	pc := corepayload.Empty.PayloadsCollection()
	actual := args.Map{"notNil": pc != nil, "empty": pc.IsEmpty()}
	expected := args.Map{"notNil": true, "empty": true}
	expected.ShouldBeEqual(t, 0, "Empty.PayloadsCollection", actual)
}
