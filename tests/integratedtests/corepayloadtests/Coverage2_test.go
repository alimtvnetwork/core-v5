package corepayloadtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corepayload"
	"github.com/alimtvnetwork/core/coretests/args"
)

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
		"notNil":     pw != nil,
		"hasPayload": pw.HasPayload(),
	}
	expected := args.Map{
		"notNil":     true,
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
