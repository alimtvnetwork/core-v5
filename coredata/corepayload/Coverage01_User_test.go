package corepayload

import (
	"testing"
)

func TestUser_BasicOps(t *testing.T) {
	u := New.User.Create(false, "alice", "admin")
	if u.IsEmpty() {
		t.Fatal("should not be empty")
	}
	if !u.IsValidUser() {
		t.Fatal("should be valid")
	}
	if u.IsNameEmpty() {
		t.Fatal("name should not be empty")
	}
	if !u.IsNameEqual("alice") {
		t.Fatal("name should match")
	}
	if !u.IsNotSystemUser() {
		t.Fatal("should not be system user")
	}
	if !u.IsVirtualUser() {
		t.Fatal("should be virtual user")
	}
	if !u.HasType() {
		t.Fatal("should have type")
	}
	if u.IsTypeEmpty() {
		t.Fatal("type should not be empty")
	}
	if u.IsAuthTokenEmpty() != true {
		t.Fatal("auth token should be empty")
	}
	if u.HasAuthToken() {
		t.Fatal("should not have auth token")
	}
	if u.IsPasswordHashEmpty() != true {
		t.Fatal("password hash should be empty")
	}
	if u.HasPasswordHash() {
		t.Fatal("should not have password hash")
	}
}

func TestUser_NilChecks(t *testing.T) {
	var nilU *User
	if !nilU.IsEmpty() {
		t.Fatal("nil should be empty")
	}
	if nilU.IsValidUser() {
		t.Fatal("nil should not be valid")
	}
	if !nilU.IsNameEmpty() {
		t.Fatal("nil name should be empty")
	}
	if nilU.IsNameEqual("x") {
		t.Fatal("nil should not match")
	}
}

func TestUser_IdentifierInteger(t *testing.T) {
	u := New.User.All(false, "123", "alice", "admin", "", "")
	if u.IdentifierInteger() != 123 {
		t.Fatal("expected 123")
	}
	u2 := New.User.Empty()
	if u2.IdentifierInteger() >= 0 {
		t.Fatal("expected invalid")
	}
}

func TestUser_IdentifierUnsignedInteger(t *testing.T) {
	u := New.User.All(false, "5", "alice", "admin", "", "")
	if u.IdentifierUnsignedInteger() != 5 {
		t.Fatal("expected 5")
	}
	u2 := New.User.Empty()
	if u2.IdentifierUnsignedInteger() != 0 {
		t.Fatal("expected 0")
	}
}

func TestUser_Clone(t *testing.T) {
	u := New.User.Create(true, "bob", "sys")
	c := u.Clone()
	if c.Name != "bob" {
		t.Fatal("clone mismatch")
	}

	cp := u.ClonePtr()
	if cp == nil || cp.Name != "bob" {
		t.Fatal("clone ptr mismatch")
	}

	var nilU *User
	if nilU.ClonePtr() != nil {
		t.Fatal("expected nil")
	}
}

func TestUser_Json(t *testing.T) {
	u := New.User.Create(false, "alice", "admin")
	j := u.Json()
	if j.HasError() {
		t.Fatal(j.Error)
	}
	jp := u.JsonPtr()
	if jp == nil {
		t.Fatal("expected non-nil")
	}
	s := u.String()
	if s == "" {
		t.Fatal("expected non-empty")
	}
	_ = u.PrettyJsonString()
}

func TestUser_Serialize_Deserialize(t *testing.T) {
	u := New.User.Create(false, "alice", "admin")
	b, err := u.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	u2 := New.User.Empty()
	err2 := u2.Deserialize(b)
	if err2 != nil {
		t.Fatal(err2)
	}
	if u2.Name != "alice" {
		t.Fatal("deserialize mismatch")
	}
}

func TestUser_Ptr(t *testing.T) {
	u := User{Name: "x"}
	p := u.Ptr()
	if p == nil {
		t.Fatal("expected non-nil")
	}
}

func TestNewUserCreator(t *testing.T) {
	_ = New.User.Empty()
	_ = New.User.Create(false, "a", "t")
	_ = New.User.NonSysCreate("a", "t")
	_ = New.User.NonSysCreateId("1", "a", "t")
	_ = New.User.System("a", "t")
	_ = New.User.SystemId("1", "a", "t")
	_ = New.User.UsingName("a")
	_ = New.User.All(true, "1", "a", "t", "tok", "hash")
}

func TestNewUserCreator_Deserialize(t *testing.T) {
	u := New.User.Create(false, "alice", "admin")
	b, _ := u.Serialize()
	u2, err := New.User.Deserialize(b)
	if err != nil || u2.Name != "alice" {
		t.Fatal("unexpected")
	}

	_, err2 := New.User.Deserialize([]byte("invalid"))
	if err2 == nil {
		t.Fatal("expected error")
	}
}

func TestNewUserCreator_CastOrDeserializeFrom(t *testing.T) {
	u := New.User.Create(false, "alice", "admin")
	u2, err := New.User.CastOrDeserializeFrom(u)
	_ = u2
	_ = err

	_, err2 := New.User.CastOrDeserializeFrom(nil)
	if err2 == nil {
		t.Fatal("expected error for nil")
	}
}

// ================== UserInfo ==================

func TestUserInfo_BasicOps(t *testing.T) {
	ui := &UserInfo{
		User:       New.User.Create(false, "alice", "admin"),
		SystemUser: New.User.System("root", "sys"),
	}
	if ui.IsEmpty() {
		t.Fatal("should not be empty")
	}
	if !ui.HasUser() {
		t.Fatal("should have user")
	}
	if !ui.HasSystemUser() {
		t.Fatal("should have system user")
	}

	var nilUI *UserInfo
	if !nilUI.IsEmpty() {
		t.Fatal("nil should be empty")
	}
}

func TestUserInfo_SetMethods(t *testing.T) {
	var nilUI *UserInfo
	ui := nilUI.SetUser(New.User.UsingName("alice"))
	if ui == nil || ui.User == nil {
		t.Fatal("expected user set")
	}

	var nilUI2 *UserInfo
	ui2 := nilUI2.SetSystemUser(New.User.System("root", "sys"))
	if ui2 == nil || ui2.SystemUser == nil {
		t.Fatal("expected system user set")
	}

	var nilUI3 *UserInfo
	ui3 := nilUI3.SetUserSystemUser(
		New.User.UsingName("alice"),
		New.User.System("root", "sys"),
	)
	if ui3 == nil {
		t.Fatal("expected non-nil")
	}
}

func TestUserInfo_Clone(t *testing.T) {
	ui := &UserInfo{User: New.User.UsingName("alice")}
	c := ui.Clone()
	_ = c

	cp := ui.ClonePtr()
	if cp == nil {
		t.Fatal("expected non-nil")
	}

	var nilUI *UserInfo
	if nilUI.ClonePtr() != nil {
		t.Fatal("expected nil")
	}
}

func TestUserInfo_ToNonPtr(t *testing.T) {
	ui := &UserInfo{User: New.User.UsingName("alice")}
	np := ui.ToNonPtr()
	_ = np

	var nilUI *UserInfo
	np2 := nilUI.ToNonPtr()
	_ = np2
}

// ================== SessionInfo ==================

func TestSessionInfo_BasicOps(t *testing.T) {
	si := &SessionInfo{Id: "123", User: New.User.UsingName("alice")}
	if si.IsEmpty() {
		t.Fatal("should not be empty")
	}
	if !si.IsValid() {
		t.Fatal("should be valid")
	}
	if !si.HasUser() {
		t.Fatal("should have user")
	}
	if si.IsUserEmpty() {
		t.Fatal("user should not be empty")
	}
	if si.IsUserNameEmpty() {
		t.Fatal("username should not be empty")
	}
	if !si.IsUsernameEqual("alice") {
		t.Fatal("should match")
	}

	var nilSI *SessionInfo
	if !nilSI.IsEmpty() {
		t.Fatal("nil should be empty")
	}
}

func TestSessionInfo_IdentifierInteger(t *testing.T) {
	si := SessionInfo{Id: "42"}
	if si.IdentifierInteger() != 42 {
		t.Fatal("expected 42")
	}
	si2 := SessionInfo{}
	if si2.IdentifierInteger() >= 0 {
		t.Fatal("expected invalid")
	}
}

func TestSessionInfo_Clone(t *testing.T) {
	si := &SessionInfo{Id: "1", User: New.User.UsingName("x")}
	c := si.Clone()
	_ = c
	cp := si.ClonePtr()
	if cp == nil {
		t.Fatal("expected non-nil")
	}

	var nilSI *SessionInfo
	if nilSI.ClonePtr() != nil {
		t.Fatal("expected nil")
	}
}

// ================== PagingInfo ==================

func TestPagingInfo_BasicOps(t *testing.T) {
	pi := &PagingInfo{TotalPages: 5, CurrentPageIndex: 1, PerPageItems: 10, TotalItems: 50}
	if pi.IsEmpty() {
		t.Fatal("should not be empty")
	}
	if !pi.HasTotalPages() {
		t.Fatal("should have total pages")
	}
	if !pi.HasCurrentPageIndex() {
		t.Fatal("should have current page")
	}
	if !pi.HasPerPageItems() {
		t.Fatal("should have per page")
	}
	if !pi.HasTotalItems() {
		t.Fatal("should have total items")
	}

	var nilPI *PagingInfo
	if !nilPI.IsEmpty() {
		t.Fatal("nil should be empty")
	}
}

func TestPagingInfo_IsEqual(t *testing.T) {
	pi1 := &PagingInfo{TotalPages: 5, CurrentPageIndex: 1, PerPageItems: 10, TotalItems: 50}
	pi2 := &PagingInfo{TotalPages: 5, CurrentPageIndex: 1, PerPageItems: 10, TotalItems: 50}
	if !pi1.IsEqual(pi2) {
		t.Fatal("should be equal")
	}

	var nilPI *PagingInfo
	if !nilPI.IsEqual(nil) {
		t.Fatal("both nil should be equal")
	}

	pi3 := &PagingInfo{TotalPages: 3}
	if pi1.IsEqual(pi3) {
		t.Fatal("should not be equal")
	}
}

func TestPagingInfo_Clone(t *testing.T) {
	pi := &PagingInfo{TotalPages: 5}
	c := pi.Clone()
	if c.TotalPages != 5 {
		t.Fatal("clone mismatch")
	}
	cp := pi.ClonePtr()
	if cp == nil {
		t.Fatal("expected non-nil")
	}

	var nilPI *PagingInfo
	if nilPI.ClonePtr() != nil {
		t.Fatal("expected nil")
	}
}

// ================== AuthInfo ==================

func TestAuthInfo_BasicOps(t *testing.T) {
	ai := &AuthInfo{ActionType: "login", ResourceName: "/api"}
	if ai.IsEmpty() {
		t.Fatal("should not be empty")
	}
	if !ai.HasActionType() {
		t.Fatal("should have action type")
	}
	if !ai.HasResourceName() {
		t.Fatal("should have resource name")
	}

	var nilAI *AuthInfo
	if !nilAI.IsEmpty() {
		t.Fatal("nil should be empty")
	}
}

func TestAuthInfo_SetMethods(t *testing.T) {
	var nilAI *AuthInfo
	ai := nilAI.SetActionType("login")
	if ai == nil || ai.ActionType != "login" {
		t.Fatal("unexpected")
	}

	var nilAI2 *AuthInfo
	ai2 := nilAI2.SetResourceName("/api")
	if ai2 == nil {
		t.Fatal("unexpected")
	}

	var nilAI3 *AuthInfo
	ai3 := nilAI3.SetIdentifier("123")
	if ai3 == nil {
		t.Fatal("unexpected")
	}

	var nilAI4 *AuthInfo
	ai4 := nilAI4.SetSessionInfo(&SessionInfo{Id: "1"})
	if ai4 == nil {
		t.Fatal("unexpected")
	}

	var nilAI5 *AuthInfo
	ai5 := nilAI5.SetUserInfo(&UserInfo{User: New.User.UsingName("x")})
	if ai5 == nil {
		t.Fatal("unexpected")
	}

	var nilAI6 *AuthInfo
	ai6 := nilAI6.SetUser(New.User.UsingName("alice"))
	if ai6 == nil {
		t.Fatal("unexpected")
	}

	var nilAI7 *AuthInfo
	ai7 := nilAI7.SetSystemUser(New.User.System("root", "sys"))
	if ai7 == nil {
		t.Fatal("unexpected")
	}

	var nilAI8 *AuthInfo
	ai8 := nilAI8.SetUserSystemUser(
		New.User.UsingName("alice"),
		New.User.System("root", "sys"),
	)
	if ai8 == nil {
		t.Fatal("unexpected")
	}
}

func TestAuthInfo_Clone(t *testing.T) {
	ai := &AuthInfo{ActionType: "login"}
	c := ai.Clone()
	if c.ActionType != "login" {
		t.Fatal("clone mismatch")
	}
	cp := ai.ClonePtr()
	if cp == nil {
		t.Fatal("expected non-nil")
	}

	var nilAI *AuthInfo
	if nilAI.ClonePtr() != nil {
		t.Fatal("expected nil")
	}
}

func TestAuthInfo_Json(t *testing.T) {
	ai := AuthInfo{ActionType: "login"}
	_ = ai.String()
	_ = ai.PrettyJsonString()
	_ = ai.Json()
	_ = ai.JsonPtr()
}

func TestAuthInfo_IdentifierInteger(t *testing.T) {
	ai := AuthInfo{Identifier: "42"}
	if ai.IdentifierInteger() != 42 {
		t.Fatal("expected 42")
	}
	ai2 := AuthInfo{}
	if ai2.IdentifierInteger() >= 0 {
		t.Fatal("expected invalid")
	}
}
