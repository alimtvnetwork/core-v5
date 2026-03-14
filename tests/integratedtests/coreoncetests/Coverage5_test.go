package coreoncetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coreonce"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── ByteOnce additional ──

func Test_Cov5_ByteOnce_Aliases(t *testing.T) {
	bo := coreonce.NewByteOnce(func() byte { return 42 })
	actual := args.Map{
		"value":   int(bo.Value()),
		"execute": int(bo.Execute()),
		"string":  bo.String(),
	}
	expected := args.Map{"value": 42, "execute": 42, "string": "42"}
	expected.ShouldBeEqual(t, 0, "ByteOnce aliases -- 42", actual)
}

func Test_Cov5_ByteOnce_SerializeMarshal(t *testing.T) {
	bo := coreonce.NewByteOnce(func() byte { return 1 })
	bytes, err := bo.Serialize()
	mb, merr := bo.MarshalJSON()
	actual := args.Map{
		"hasBytes": len(bytes) > 0, "noErr": err == nil,
		"hasMb": len(mb) > 0, "noMerr": merr == nil,
	}
	expected := args.Map{
		"hasBytes": true, "noErr": true,
		"hasMb": true, "noMerr": true,
	}
	expected.ShouldBeEqual(t, 0, "ByteOnce serialize/marshal -- valid", actual)
}

// ── BoolOnce additional ──

func Test_Cov5_BoolOnce_Aliases(t *testing.T) {
	bo := coreonce.NewBoolOnce(func() bool { return true })
	actual := args.Map{
		"value":   bo.Value(),
		"execute": bo.Execute(),
		"isTrue":  bo.IsTrue(),
		"isFalse": bo.IsFalse(),
		"string":  bo.String(),
	}
	expected := args.Map{
		"value": true, "execute": true,
		"isTrue": true, "isFalse": false,
		"string": "true",
	}
	expected.ShouldBeEqual(t, 0, "BoolOnce aliases -- true", actual)
}

// ── StringOnce additional ──

func Test_Cov5_StringOnce_Aliases(t *testing.T) {
	so := coreonce.NewStringOnce(func() string { return "hello" })
	actual := args.Map{
		"value":   so.Value(),
		"execute": so.Execute(),
		"string":  so.String(),
		"isEmpty": so.IsEmpty(),
	}
	expected := args.Map{
		"value": "hello", "execute": "hello",
		"string": "hello", "isEmpty": false,
	}
	expected.ShouldBeEqual(t, 0, "StringOnce aliases -- hello", actual)
}

func Test_Cov5_StringOnce_Empty(t *testing.T) {
	so := coreonce.NewStringOnce(func() string { return "" })
	actual := args.Map{"isEmpty": so.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringOnce empty -- true", actual)
}

// ── ErrorOnce additional ──

func Test_Cov5_ErrorOnce_Nil(t *testing.T) {
	eo := coreonce.NewErrorOnce(func() error { return nil })
	actual := args.Map{
		"isNil":    eo.Value() == nil,
		"hasError": eo.HasError(),
		"string":   eo.String(),
	}
	expected := args.Map{
		"isNil": true, "hasError": false, "string": "",
	}
	expected.ShouldBeEqual(t, 0, "ErrorOnce nil -- no error", actual)
}

// ── BytesOnce additional ──

func Test_Cov5_BytesOnce_Aliases(t *testing.T) {
	bo := coreonce.NewBytesOnce(func() []byte { return []byte{1, 2} })
	actual := args.Map{
		"len":     len(bo.Value()),
		"execute": len(bo.Execute()),
		"string":  bo.String() != "",
	}
	expected := args.Map{
		"len": 2, "execute": 2, "string": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesOnce aliases -- 2 bytes", actual)
}
