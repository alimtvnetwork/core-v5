package reflectmodeltests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/reflectcore/reflectmodel"
)

// ── FieldProcessor ──

func Test_Cov2_FieldProcessor_IsFieldType(t *testing.T) {
	fp := &reflectmodel.FieldProcessor{
		Name: "A", Index: 0,
		FieldType: reflect.TypeOf(42),
	}
	var nilFP *reflectmodel.FieldProcessor
	actual := args.Map{
		"isInt":    fp.IsFieldType(reflect.TypeOf(42)),
		"isString": fp.IsFieldType(reflect.TypeOf("")),
		"nil":      nilFP.IsFieldType(reflect.TypeOf(42)),
	}
	expected := args.Map{"isInt": true, "isString": false, "nil": false}
	expected.ShouldBeEqual(t, 0, "FieldProcessor returns correct value -- IsFieldType", actual)
}

func Test_Cov2_FieldProcessor_IsFieldKind(t *testing.T) {
	fp := &reflectmodel.FieldProcessor{
		FieldType: reflect.TypeOf(42),
	}
	var nilFP *reflectmodel.FieldProcessor
	actual := args.Map{
		"isInt":  fp.IsFieldKind(reflect.Int),
		"isStr":  fp.IsFieldKind(reflect.String),
		"nil":    nilFP.IsFieldKind(reflect.Int),
	}
	expected := args.Map{"isInt": true, "isStr": false, "nil": false}
	expected.ShouldBeEqual(t, 0, "FieldProcessor returns correct value -- IsFieldKind", actual)
}

// ── MethodProcessor extended ──

func Test_Cov2_MethodProcessor_Nil(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	// Only test methods that are safe on nil receivers
	actual := args.Map{
		"hasValid":  mp.HasValidFunc(),
		"isInvalid": mp.IsInvalid(),
		"funcNil":   mp.Func() == nil,
		"isPublic":  mp.IsPublicMethod(),
		"isPrivate": mp.IsPrivateMethod(),
	}
	expected := args.Map{
		"hasValid": false, "isInvalid": true,
		"funcNil": true, "isPublic": false, "isPrivate": false,
	}
	expected.ShouldBeEqual(t, 0, "MethodProcessor returns nil -- Nil", actual)
}

func Test_Cov2_MethodProcessor_IsEqual(t *testing.T) {
	var mp1, mp2 *reflectmodel.MethodProcessor
	actual := args.Map{"nilNil": mp1.IsEqual(mp2)}
	expected := args.Map{"nilNil": true}
	expected.ShouldBeEqual(t, 0, "MethodProcessor returns nil -- IsEqual nil", actual)
}

func Test_Cov2_MethodProcessor_IsNotEqual(t *testing.T) {
	var mp1, mp2 *reflectmodel.MethodProcessor
	actual := args.Map{"notEqual": mp1.IsNotEqual(mp2)}
	expected := args.Map{"notEqual": false}
	expected.ShouldBeEqual(t, 0, "MethodProcessor returns correct value -- IsNotEqual", actual)
}

func Test_Cov2_MethodProcessor_GetOutArgsTypes_Nil(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	result := mp.GetOutArgsTypes()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MethodProcessor returns nil -- GetOutArgsTypes nil", actual)
}

func Test_Cov2_MethodProcessor_GetInArgsTypes_Nil(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	result := mp.GetInArgsTypes()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MethodProcessor returns nil -- GetInArgsTypes nil", actual)
}

func Test_Cov2_MethodProcessor_GetInArgsTypesNames_Nil(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	result := mp.GetInArgsTypesNames()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MethodProcessor returns nil -- GetInArgsTypesNames nil", actual)
}

// ── rvUtils is unexported — cannot be tested from external package ──
