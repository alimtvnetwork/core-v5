package namevaluetests

import (
	"testing"

	"github.com/alimtvnetwork/core/namevalue"
)

func Test_QW_Collection_String_NilReceiver(t *testing.T) {
	var c *namevalue.Collection[string, string]
	s := c.String()
	if s != "" {
		t.Fatal("expected empty for nil")
	}
}

func Test_QW_Collection_JsonString_NilReceiver(t *testing.T) {
	var c *namevalue.Collection[string, string]
	s := c.JsonString()
	if s != "" {
		t.Fatal("expected empty for nil")
	}
}

func Test_QW_Instance_IsNull(t *testing.T) {
	var inst *namevalue.Instance[string, string]
	if !inst.IsNull() {
		t.Fatal("expected null")
	}
}

func Test_QW_Instance_String_Nil(t *testing.T) {
	var inst *namevalue.Instance[string, string]
	s := inst.String()
	if s != "" {
		t.Fatal("expected empty for nil")
	}
}

func Test_QW_Instance_JsonString_Nil(t *testing.T) {
	var inst *namevalue.Instance[string, string]
	s := inst.JsonString()
	if s != "" {
		t.Fatal("expected empty for nil")
	}
}
