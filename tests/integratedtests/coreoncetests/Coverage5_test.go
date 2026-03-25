package coreoncetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coreonce"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── BoolOnce ──

func Test_Cov5_BoolOnce_Value(t *testing.T) {
	bo := coreonce.NewBoolOnce(func() bool { return true })
	actual := args.Map{
		"value":   bo.Value(),
		"execute": bo.Execute(),
		"string":  bo.String(),
	}
	expected := args.Map{
		"value":   true,
		"execute": true,
		"string":  "true",
	}
	expected.ShouldBeEqual(t, 0, "BoolOnce returns true -- true func", actual)
}

func Test_Cov5_BoolOnce_False(t *testing.T) {
	bo := coreonce.NewBoolOnce(func() bool { return false })
	actual := args.Map{
		"value":  bo.Value(),
		"string": bo.String(),
	}
	expected := args.Map{
		"value":  false,
		"string": "false",
	}
	expected.ShouldBeEqual(t, 0, "BoolOnce returns false -- false func", actual)
}

func Test_Cov5_BoolOnce_Serialize(t *testing.T) {
	bo := coreonce.NewBoolOnce(func() bool { return true })
	b, err := bo.Serialize()
	actual := args.Map{
		"noErr":    err == nil,
		"hasBytes": len(b) > 0,
	}
	expected := args.Map{
		"noErr":    true,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "BoolOnce.Serialize succeeds -- true", actual)
}

func Test_Cov5_BoolOnce_MarshalJSON(t *testing.T) {
	bo := coreonce.NewBoolOnce(func() bool { return true })
	b, err := bo.MarshalJSON()
	actual := args.Map{
		"noErr":    err == nil,
		"hasBytes": len(b) > 0,
	}
	expected := args.Map{
		"noErr":    true,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "BoolOnce.MarshalJSON succeeds -- true", actual)
}

func Test_Cov5_BoolOnce_UnmarshalJSON(t *testing.T) {
	bo := coreonce.NewBoolOncePtr(func() bool { return false })
	err := bo.UnmarshalJSON([]byte("true"))
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "BoolOnce.UnmarshalJSON succeeds -- true bytes", actual)
}

func Test_Cov5_BoolOnce_Ptr(t *testing.T) {
	bo := coreonce.NewBoolOncePtr(func() bool { return true })
	actual := args.Map{
		"notNil": bo != nil,
		"value":  bo.Value(),
	}
	expected := args.Map{
		"notNil": true,
		"value":  true,
	}
	expected.ShouldBeEqual(t, 0, "NewBoolOncePtr returns non-nil -- true", actual)
}
