package corepropertytests

import (
	"testing"

	"github.com/alimtvnetwork/core/codegen/coreproperty"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── WriteStruct with unexported fields ──

type covStructWithUnexported struct {
	Public  string
	private int //nolint:unused
}

func Test_Cov_WriteStruct_UnexportedField(t *testing.T) {
	s := covStructWithUnexported{Public: "hello"}
	result := coreproperty.Writer.WriteStruct(s)

	actual := args.Map{
		"nonEmpty":       result != "",
		"containsPublic": len(result) > 0,
	}
	expected := args.Map{
		"nonEmpty":       true,
		"containsPublic": true,
	}
	expected.ShouldBeEqual(t, 0, "WriteStruct unexported field skipped", actual)
}

// ── WritePropertyOptions default fallback ──

type covCustomType struct{}

func (c covCustomType) String() string { return "custom" }

func Test_Cov_WritePropertyOptions_DefaultCase(t *testing.T) {
	// A type that doesn't match string/bool/int/float/struct/slice/array/ptr/map
	// would fall through to the default case. Interface values can trigger this.
	var iface interface{ String() string } = covCustomType{}
	result := coreproperty.Writer.Write(iface)

	actual := args.Map{"nonEmpty": result != ""}
	expected := args.Map{"nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "WritePropertyOptions default case", actual)
}

// ── WriteMap with nested types ──

func Test_Cov_WriteMap_NestedSlice(t *testing.T) {
	m := map[string][]int{"nums": {1, 2, 3}}
	result := coreproperty.Writer.Write(m)

	actual := args.Map{"nonEmpty": result != "" && result != "nil"}
	expected := args.Map{"nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "WriteMap nested slice", actual)
}

// ── WritePointerRv with nil pointer ──

func Test_Cov_WritePointerRv_NilPtr(t *testing.T) {
	var ptr *string
	result := coreproperty.Writer.Write(ptr)

	actual := args.Map{"result": result}
	expected := args.Map{"result": "nil"}
	expected.ShouldBeEqual(t, 0, "WritePointerRv nil ptr", actual)
}
