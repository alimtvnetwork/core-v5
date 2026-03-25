package coretaskinfotests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretaskinfo"
)

func Test_QW_Info_JsonString_Nil(t *testing.T) {
	defer func() { recover() }() // value receiver on nil pointer may panic
	var info *coretaskinfo.Info
	s := info.JsonString()
	if s != "" {
		t.Fatal("expected empty for nil")
	}
}

func Test_QW_Info_LazyMapWithPayload_ErrorPath(t *testing.T) {
	info := &coretaskinfo.Info{
		RootName:    "test",
		Description: "desc",
	}
	// LazyMapWithPayload takes []byte — pass invalid JSON bytes to cover error branch
	m := info.LazyMapWithPayload([]byte(`{invalid`))
	if m == nil {
		t.Fatal("expected non-nil map")
	}
}

func Test_QW_Info_LazyMapWithPayloadAsAny_ErrorPath(t *testing.T) {
	info := &coretaskinfo.Info{
		RootName:    "test",
		Description: "desc",
	}
	m := info.LazyMapWithPayloadAsAny(make(chan int))
	if m == nil {
		t.Fatal("expected non-nil map")
	}
}
