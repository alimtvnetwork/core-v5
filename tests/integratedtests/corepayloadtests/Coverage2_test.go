package corepayloadtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corepayload"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Attributes ──

func Test_Cov2_Attributes_NewAndGetters(t *testing.T) {
	attrs := corepayload.New.Attributes.Create(map[string]string{"key": "val"})
	emptyAttrs := corepayload.New.Attributes.Empty()

	actual := args.Map{
		"len":         attrs.Length(),
		"isEmpty":     attrs.IsEmpty(),
		"hasAny":      attrs.HasAnyItem(),
		"emptyLen":    emptyAttrs.Length(),
		"emptyIsEmpty": emptyAttrs.IsEmpty(),
	}
	expected := args.Map{
		"len":         1,
		"isEmpty":     false,
		"hasAny":      true,
		"emptyLen":    0,
		"emptyIsEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Attributes new and getters", actual)
}

func Test_Cov2_Attributes_GetSet(t *testing.T) {
	attrs := corepayload.New.Attributes.Empty()
	attrs.Set("k1", "v1")
	val := attrs.Get("k1")
	missing := attrs.Get("missing")

	actual := args.Map{
		"val":     val,
		"missing": missing,
		"has":     attrs.Has("k1"),
		"hasNot":  attrs.Has("missing"),
	}
	expected := args.Map{
		"val":     "v1",
		"missing": "",
		"has":     true,
		"hasNot":  false,
	}
	expected.ShouldBeEqual(t, 0, "Attributes get/set", actual)
}

func Test_Cov2_Attributes_Json(t *testing.T) {
	attrs := corepayload.New.Attributes.Create(map[string]string{"k": "v"})
	bytes, err := attrs.MarshalJSON()

	actual := args.Map{
		"hasBytes": len(bytes) > 0,
		"noErr":    err == nil,
	}
	expected := args.Map{
		"hasBytes": true,
		"noErr":    true,
	}
	expected.ShouldBeEqual(t, 0, "Attributes JSON", actual)
}

func Test_Cov2_Attributes_Clone(t *testing.T) {
	attrs := corepayload.New.Attributes.Create(map[string]string{"k": "v"})
	clone := attrs.Clone()

	actual := args.Map{
		"notNil": clone != nil,
		"len":    clone.Length(),
	}
	expected := args.Map{
		"notNil": true,
		"len":    1,
	}
	expected.ShouldBeEqual(t, 0, "Attributes clone", actual)
}

// ── PagingInfo ──

func Test_Cov2_PagingInfo(t *testing.T) {
	p := corepayload.PagingInfo{
		PageIndex: 0,
		PageSize:  10,
		Total:     25,
	}

	actual := args.Map{
		"totalPages": p.TotalPages(),
		"hasNext":    p.HasNextPage(),
		"hasPrev":    p.HasPreviousPage(),
	}
	expected := args.Map{
		"totalPages": 3,
		"hasNext":    true,
		"hasPrev":    false,
	}
	expected.ShouldBeEqual(t, 0, "PagingInfo", actual)
}

func Test_Cov2_PagingInfo_LastPage(t *testing.T) {
	p := corepayload.PagingInfo{
		PageIndex: 2,
		PageSize:  10,
		Total:     25,
	}

	actual := args.Map{
		"hasNext": p.HasNextPage(),
		"hasPrev": p.HasPreviousPage(),
	}
	expected := args.Map{
		"hasNext": false,
		"hasPrev": true,
	}
	expected.ShouldBeEqual(t, 0, "PagingInfo last page", actual)
}

// ── SessionInfo ──

func Test_Cov2_SessionInfo(t *testing.T) {
	s := corepayload.SessionInfo{SessionId: "abc123"}

	actual := args.Map{
		"hasSession": s.HasSession(),
		"sessionId":  s.SessionId,
	}
	expected := args.Map{
		"hasSession": true,
		"sessionId":  "abc123",
	}
	expected.ShouldBeEqual(t, 0, "SessionInfo", actual)
}

func Test_Cov2_SessionInfo_Empty(t *testing.T) {
	s := corepayload.SessionInfo{}

	actual := args.Map{"hasSession": s.HasSession()}
	expected := args.Map{"hasSession": false}
	expected.ShouldBeEqual(t, 0, "SessionInfo empty", actual)
}

// ── AuthInfo ──

func Test_Cov2_AuthInfo(t *testing.T) {
	a := corepayload.AuthInfo{Token: "tok", UserId: "uid"}

	actual := args.Map{
		"hasToken":  a.HasToken(),
		"hasUserId": a.HasUserId(),
		"token":     a.Token,
	}
	expected := args.Map{
		"hasToken":  true,
		"hasUserId": true,
		"token":     "tok",
	}
	expected.ShouldBeEqual(t, 0, "AuthInfo", actual)
}

// ── PayloadWrapper ──

func Test_Cov2_PayloadWrapper_Basic(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.UsingPayload(map[string]string{"k": "v"})

	actual := args.Map{
		"notNil":   pw != nil,
		"hasPayload": pw.HasPayload(),
	}
	expected := args.Map{
		"notNil":   true,
		"hasPayload": true,
	}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper basic", actual)
}

func Test_Cov2_PayloadWrapper_Empty(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()

	actual := args.Map{
		"notNil":     pw != nil,
		"hasPayload": pw.HasPayload(),
	}
	expected := args.Map{
		"notNil":     true,
		"hasPayload": false,
	}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper empty", actual)
}
