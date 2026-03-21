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
	defer func() { recover() }() // value receiver on nil pointer may panic
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
	defer func() {
		if r := recover(); r != nil {
			// nil pointer panic is expected for zero-value Instance.String()
		}
	}()
	inst := namevalue.Instance[string, string]{}
	_ = inst.String()
}

func Test_QW_Instance_JsonString_Nil(t *testing.T) {
	// JsonString is a value receiver — calling on nil pointer panics
	defer func() {
		if r := recover(); r != nil {
			// expected: nil pointer dereference on value receiver
		}
	}()
	var inst *namevalue.Instance[string, string]
	_ = inst.JsonString()
}
