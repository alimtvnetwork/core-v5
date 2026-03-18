package reflectinternaltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/internal/reflectinternal"
	"github.com/alimtvnetwork/core/reflectcore/reflectmodel"
)

// ── reflectConverter — uncovered branches ──

func Test_Cov5_Converter_ReflectValuesToInterfaces_Empty(t *testing.T) {
	result := reflectinternal.Converter.ReflectValuesToInterfaces(nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ReflectValuesToInterfaces empty", actual)
}

func Test_Cov5_Converter_ReflectValuesToInterfaces_Valid(t *testing.T) {
	rvs := []reflect.Value{reflect.ValueOf(42), reflect.ValueOf("hello")}
	result := reflectinternal.Converter.ReflectValuesToInterfaces(rvs)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "ReflectValuesToInterfaces valid", actual)
}

func Test_Cov5_Converter_ReflectValueToAnyValue_Ptr(t *testing.T) {
	x := 42
	rv := reflect.ValueOf(&x)
	result := reflectinternal.Converter.ReflectValueToAnyValue(rv)
	actual := args.Map{"val": result}
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "ReflectValueToAnyValue ptr", actual)
}

func Test_Cov5_Converter_ReflectValueToAnyValue_String(t *testing.T) {
	rv := reflect.ValueOf("hello")
	result := reflectinternal.Converter.ReflectValueToAnyValue(rv)
	actual := args.Map{"val": result}
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "ReflectValueToAnyValue string", actual)
}

func Test_Cov5_Converter_ReflectValueToAnyValue_Int(t *testing.T) {
	rv := reflect.ValueOf(42)
	result := reflectinternal.Converter.ReflectValueToAnyValue(rv)
	actual := args.Map{"val": result}
	expected := args.Map{"val": int64(42)}
	expected.ShouldBeEqual(t, 0, "ReflectValueToAnyValue int", actual)
}

func Test_Cov5_Converter_ReflectValueToAnyValue_Default(t *testing.T) {
	rv := reflect.ValueOf(3.14)
	result := reflectinternal.Converter.ReflectValueToAnyValue(rv)
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ReflectValueToAnyValue default", actual)
}

func Test_Cov5_Converter_InterfacesToTypes(t *testing.T) {
	result := reflectinternal.Converter.InterfacesToTypes([]any{42, "hello"})
	emptyResult := reflectinternal.Converter.InterfacesToTypes(nil)
	actual := args.Map{"len": len(result), "emptyLen": len(emptyResult)}
	expected := args.Map{"len": 2, "emptyLen": 0}
	expected.ShouldBeEqual(t, 0, "InterfacesToTypes", actual)
}

func Test_Cov5_Converter_InterfacesToTypesNames(t *testing.T) {
	result := reflectinternal.Converter.InterfacesToTypesNames([]any{42, "hello"})
	emptyResult := reflectinternal.Converter.InterfacesToTypesNames(nil)
	actual := args.Map{"len": len(result), "emptyLen": len(emptyResult)}
	expected := args.Map{"len": 2, "emptyLen": 0}
	expected.ShouldBeEqual(t, 0, "InterfacesToTypesNames", actual)
}

func Test_Cov5_Converter_InterfacesToTypesNamesWithValues(t *testing.T) {
	result := reflectinternal.Converter.InterfacesToTypesNamesWithValues([]any{42})
	emptyResult := reflectinternal.Converter.InterfacesToTypesNamesWithValues(nil)
	actual := args.Map{"len": len(result), "emptyLen": len(emptyResult)}
	expected := args.Map{"len": 1, "emptyLen": 0}
	expected.ShouldBeEqual(t, 0, "InterfacesToTypesNamesWithValues", actual)
}

func Test_Cov5_Converter_ReflectValueToPointerReflectValue(t *testing.T) {
	rv := reflect.ValueOf(42)
	result := reflectinternal.Converter.ReflectValueToPointerReflectValue(rv)
	actual := args.Map{"isPtr": result.Kind() == reflect.Ptr}
	expected := args.Map{"isPtr": true}
	expected.ShouldBeEqual(t, 0, "ReflectValueToPointerReflectValue", actual)
}

func Test_Cov5_Converter_ToPtrRvIfNotAlready_NonPtr(t *testing.T) {
	rv := reflect.ValueOf(42)
	result := reflectinternal.Converter.ToPtrRvIfNotAlready(rv)
	actual := args.Map{"isPtr": result.Kind() == reflect.Ptr}
	expected := args.Map{"isPtr": true}
	expected.ShouldBeEqual(t, 0, "ToPtrRvIfNotAlready non-ptr", actual)
}

func Test_Cov5_Converter_ToPtrRvIfNotAlready_AlreadyPtr(t *testing.T) {
	x := 42
	rv := reflect.ValueOf(&x)
	result := reflectinternal.Converter.ToPtrRvIfNotAlready(rv)
	actual := args.Map{"isPtr": result.Kind() == reflect.Ptr}
	expected := args.Map{"isPtr": true}
	expected.ShouldBeEqual(t, 0, "ToPtrRvIfNotAlready already-ptr", actual)
}

func Test_Cov5_Converter_ReducePointer(t *testing.T) {
	x := 42
	result := reflectinternal.Converter.ReducePointer(&x, 3)
	actual := args.Map{"valid": result.IsValid, "kind": result.Kind == reflect.Int}
	expected := args.Map{"valid": true, "kind": true}
	expected.ShouldBeEqual(t, 0, "ReducePointer", actual)
}

func Test_Cov5_Converter_ReducePointerDefault(t *testing.T) {
	x := 42
	result := reflectinternal.Converter.ReducePointerDefault(&x)
	actual := args.Map{"valid": result.IsValid}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "ReducePointerDefault", actual)
}

func Test_Cov5_Converter_ReducePointerRvDefault(t *testing.T) {
	x := 42
	rv := reflect.ValueOf(&x)
	result := reflectinternal.Converter.ReducePointerRvDefault(rv)
	actual := args.Map{"valid": result.IsValid}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "ReducePointerRvDefault", actual)
}

func Test_Cov5_Converter_ReducePointerDefaultToType(t *testing.T) {
	x := 42
	result := reflectinternal.Converter.ReducePointerDefaultToType(&x)
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ReducePointerDefaultToType", actual)
}

func Test_Cov5_Converter_ReducePointerRvDefaultToType_Nil(t *testing.T) {
	rv := reflect.Value{}
	result := reflectinternal.Converter.ReducePointerRvDefaultToType(rv)
	actual := args.Map{"nil": result == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "ReducePointerRvDefaultToType nil", actual)
}

func Test_Cov5_Converter_ReflectValToInterfaces_Slice(t *testing.T) {
	rv := reflect.ValueOf([]int{1, 2, 3})
	result := reflectinternal.Converter.ReflectValToInterfaces(false, rv)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "ReflectValToInterfaces slice", actual)
}

func Test_Cov5_Converter_ReflectValToInterfaces_PtrSlice(t *testing.T) {
	s := []int{1, 2}
	rv := reflect.ValueOf(&s)
	result := reflectinternal.Converter.ReflectValToInterfaces(false, rv)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "ReflectValToInterfaces ptr-slice", actual)
}

func Test_Cov5_Converter_ReflectValToInterfaces_NotSlice(t *testing.T) {
	rv := reflect.ValueOf(42)
	result := reflectinternal.Converter.ReflectValToInterfaces(false, rv)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ReflectValToInterfaces not-slice", actual)
}

func Test_Cov5_Converter_ReflectValToInterfaces_EmptySlice(t *testing.T) {
	rv := reflect.ValueOf([]int{})
	result := reflectinternal.Converter.ReflectValToInterfaces(false, rv)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ReflectValToInterfaces empty-slice", actual)
}

func Test_Cov5_Converter_ReflectValToInterfacesAsync(t *testing.T) {
	rv := reflect.ValueOf([]int{1, 2, 3})
	result := reflectinternal.Converter.ReflectValToInterfacesAsync(rv)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "ReflectValToInterfacesAsync", actual)
}

func Test_Cov5_Converter_ReflectValToInterfacesAsync_Ptr(t *testing.T) {
	s := []int{1, 2}
	rv := reflect.ValueOf(&s)
	result := reflectinternal.Converter.ReflectValToInterfacesAsync(rv)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "ReflectValToInterfacesAsync ptr", actual)
}

func Test_Cov5_Converter_ReflectValToInterfacesAsync_NotSlice(t *testing.T) {
	rv := reflect.ValueOf(42)
	result := reflectinternal.Converter.ReflectValToInterfacesAsync(rv)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ReflectValToInterfacesAsync not-slice", actual)
}

func Test_Cov5_Converter_ReflectValToInterfacesUsingProcessor(t *testing.T) {
	rv := reflect.ValueOf([]int{1, 2, 3})
	result := reflectinternal.Converter.ReflectValToInterfacesUsingProcessor(false,
		func(item any) (any, bool, bool) { return item, true, false }, rv)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "ReflectValToInterfacesUsingProcessor", actual)
}

func Test_Cov5_Converter_ReflectValToInterfacesUsingProcessor_Break(t *testing.T) {
	rv := reflect.ValueOf([]int{1, 2, 3})
	result := reflectinternal.Converter.ReflectValToInterfacesUsingProcessor(false,
		func(item any) (any, bool, bool) { return item, true, true }, rv)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "ReflectValToInterfacesUsingProcessor break", actual)
}

func Test_Cov5_Converter_ReflectInterfaceVal(t *testing.T) {
	result := reflectinternal.Converter.ReflectInterfaceVal(42)
	actual := args.Map{"val": result}
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "ReflectInterfaceVal", actual)
}

func Test_Cov5_Converter_ReflectInterfaceVal_Ptr(t *testing.T) {
	x := 42
	result := reflectinternal.Converter.ReflectInterfaceVal(&x)
	actual := args.Map{"val": result}
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "ReflectInterfaceVal ptr", actual)
}

func Test_Cov5_Converter_ToPointerRv(t *testing.T) {
	result := reflectinternal.Converter.ToPointerRv(42)
	nilResult := reflectinternal.Converter.ToPointerRv(nil)
	actual := args.Map{"notNil": result != nil, "nilResult": nilResult == nil}
	expected := args.Map{"notNil": true, "nilResult": true}
	expected.ShouldBeEqual(t, 0, "ToPointerRv", actual)
}

func Test_Cov5_Converter_ToPointer(t *testing.T) {
	result := reflectinternal.Converter.ToPointer(42)
	nilResult := reflectinternal.Converter.ToPointer(nil)
	actual := args.Map{"notNil": result != nil, "nilResult": nilResult == nil}
	expected := args.Map{"notNil": true, "nilResult": true}
	expected.ShouldBeEqual(t, 0, "ToPointer", actual)
}

// ── reflectGetUsingReflectValue — uncovered branches ──

func Test_Cov5_RvGetter_PublicValuesMapStruct_NonStruct(t *testing.T) {
	rv := reflect.ValueOf(42)
	_, err := reflectinternal.ReflectGetterUsingReflectValue.PublicValuesMapStruct(rv)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "PublicValuesMapStruct non-struct", actual)
}

func Test_Cov5_RvGetter_PublicValuesMapStruct_WithUnexported(t *testing.T) {
	type S struct {
		Pub  int
		priv int //nolint:unused
	}
	rv := reflect.ValueOf(S{Pub: 1})
	m, err := reflectinternal.ReflectGetterUsingReflectValue.PublicValuesMapStruct(rv)
	actual := args.Map{"len": len(m), "noErr": err == nil}
	expected := args.Map{"len": 1, "noErr": true}
	expected.ShouldBeEqual(t, 0, "PublicValuesMapStruct unexported", actual)
}

func Test_Cov5_RvGetter_FieldNameWithTypeMap_Struct(t *testing.T) {
	type S struct{ A int; B string }
	rv := reflect.ValueOf(S{})
	m := reflectinternal.ReflectGetterUsingReflectValue.FieldNameWithTypeMap(rv)
	actual := args.Map{"len": len(m)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "FieldNameWithTypeMap struct", actual)
}

func Test_Cov5_RvGetter_FieldNameWithTypeMap_Ptr(t *testing.T) {
	type S struct{ A int }
	rv := reflect.ValueOf(&S{})
	m := reflectinternal.ReflectGetterUsingReflectValue.FieldNameWithTypeMap(rv)
	actual := args.Map{"len": len(m)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "FieldNameWithTypeMap ptr", actual)
}

func Test_Cov5_RvGetter_FieldNameWithTypeMap_NonStruct(t *testing.T) {
	rv := reflect.ValueOf(42)
	m := reflectinternal.ReflectGetterUsingReflectValue.FieldNameWithTypeMap(rv)
	actual := args.Map{"nil": m == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "FieldNameWithTypeMap non-struct", actual)
}

func Test_Cov5_RvGetter_FieldNameWithValuesMap(t *testing.T) {
	type S struct{ A int }
	rv := reflect.ValueOf(S{A: 42})
	m, err := reflectinternal.ReflectGetterUsingReflectValue.FieldNameWithValuesMap(rv)
	actual := args.Map{"len": len(m), "noErr": err == nil}
	expected := args.Map{"len": 1, "noErr": true}
	expected.ShouldBeEqual(t, 0, "FieldNameWithValuesMap", actual)
}

func Test_Cov5_RvGetter_FieldNamesMap_Ptr(t *testing.T) {
	type S struct{ A int }
	rv := reflect.ValueOf(&S{})
	m, err := reflectinternal.ReflectGetterUsingReflectValue.FieldNamesMap(rv)
	actual := args.Map{"len": len(m), "noErr": err == nil}
	expected := args.Map{"len": 1, "noErr": true}
	expected.ShouldBeEqual(t, 0, "FieldNamesMap ptr", actual)
}

func Test_Cov5_RvGetter_FieldNamesMap_NonStruct(t *testing.T) {
	rv := reflect.ValueOf(42)
	_, err := reflectinternal.ReflectGetterUsingReflectValue.FieldNamesMap(rv)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "FieldNamesMap non-struct", actual)
}

func Test_Cov5_RvGetter_StructFieldsMap_Ptr(t *testing.T) {
	type S struct{ A int }
	rv := reflect.ValueOf(&S{})
	m := reflectinternal.ReflectGetterUsingReflectValue.StructFieldsMap(rv)
	actual := args.Map{"len": len(m)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "StructFieldsMap ptr", actual)
}

func Test_Cov5_RvGetter_StructFieldsMap_NonStruct(t *testing.T) {
	rv := reflect.ValueOf(42)
	m := reflectinternal.ReflectGetterUsingReflectValue.StructFieldsMap(rv)
	actual := args.Map{"nil": m == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "StructFieldsMap non-struct", actual)
}

func Test_Cov5_RvGetter_NullFieldsMap(t *testing.T) {
	type S struct{ A *int }
	rv := reflect.ValueOf(S{})
	m := reflectinternal.ReflectGetterUsingReflectValue.NullFieldsMap(3, rv)
	actual := args.Map{"gt0": len(m) > 0}
	expected := args.Map{"gt0": true}
	expected.ShouldBeEqual(t, 0, "NullFieldsMap", actual)
}

func Test_Cov5_RvGetter_NullOrZeroFieldsMap(t *testing.T) {
	type S struct{ A int; B *int }
	rv := reflect.ValueOf(S{})
	m := reflectinternal.ReflectGetterUsingReflectValue.NullOrZeroFieldsMap(3, rv)
	actual := args.Map{"gt0": len(m) > 0}
	expected := args.Map{"gt0": true}
	expected.ShouldBeEqual(t, 0, "NullOrZeroFieldsMap", actual)
}

// ── reflectTypeConverter — uncovered branches ──

func Test_Cov5_ReflectType_SafeName(t *testing.T) {
	result := reflectinternal.ReflectType.SafeName(42)
	nilResult := reflectinternal.ReflectType.SafeName(nil)
	actual := args.Map{"val": result, "nil": nilResult}
	expected := args.Map{"val": "int", "nil": ""}
	expected.ShouldBeEqual(t, 0, "SafeName", actual)
}

func Test_Cov5_ReflectType_SafeTypeNameOfSliceOrSingle(t *testing.T) {
	single := reflectinternal.ReflectType.SafeTypeNameOfSliceOrSingle(true, 42)
	slice := reflectinternal.ReflectType.SafeTypeNameOfSliceOrSingle(false, []int{1})
	actual := args.Map{"single": single, "slice": slice}
	expected := args.Map{"single": "int", "slice": "int"}
	expected.ShouldBeEqual(t, 0, "SafeTypeNameOfSliceOrSingle", actual)
}

func Test_Cov5_ReflectType_SliceFirstItemTypeName(t *testing.T) {
	result := reflectinternal.ReflectType.SliceFirstItemTypeName([]int{1})
	nilResult := reflectinternal.ReflectType.SliceFirstItemTypeName(nil)
	actual := args.Map{"val": result, "nil": nilResult}
	expected := args.Map{"val": "int", "nil": ""}
	expected.ShouldBeEqual(t, 0, "SliceFirstItemTypeName", actual)
}

func Test_Cov5_ReflectType_NamesStringUsingReflectType(t *testing.T) {
	full := reflectinternal.ReflectType.NamesStringUsingReflectType(true, reflect.TypeOf(42))
	empty := reflectinternal.ReflectType.NamesStringUsingReflectType(true)
	actual := args.Map{"notEmpty": full != "", "empty": empty}
	expected := args.Map{"notEmpty": true, "empty": ""}
	expected.ShouldBeEqual(t, 0, "NamesStringUsingReflectType", actual)
}

func Test_Cov5_ReflectType_TypeNamesString(t *testing.T) {
	full := reflectinternal.ReflectType.TypeNamesString(true, 42)
	empty := reflectinternal.ReflectType.TypeNamesString(true)
	actual := args.Map{"notEmpty": full != "", "empty": empty}
	expected := args.Map{"notEmpty": true, "empty": ""}
	expected.ShouldBeEqual(t, 0, "TypeNamesString", actual)
}

func Test_Cov5_ReflectType_NamesUsingReflectType_Short(t *testing.T) {
	result := reflectinternal.ReflectType.NamesUsingReflectType(false, reflect.TypeOf(42))
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "NamesUsingReflectType short", actual)
}

func Test_Cov5_ReflectType_NamesReferenceString(t *testing.T) {
	result := reflectinternal.ReflectType.NamesReferenceString(true, 42)
	empty := reflectinternal.ReflectType.NamesReferenceString(true)
	actual := args.Map{"notEmpty": result != "", "empty": empty}
	expected := args.Map{"notEmpty": true, "empty": ""}
	expected.ShouldBeEqual(t, 0, "NamesReferenceString", actual)
}

func Test_Cov5_ReflectType_Names(t *testing.T) {
	full := reflectinternal.ReflectType.Names(true, 42)
	short := reflectinternal.ReflectType.Names(false, 42)
	empty := reflectinternal.ReflectType.Names(true)
	actual := args.Map{"fullLen": len(full), "shortLen": len(short), "emptyLen": len(empty)}
	expected := args.Map{"fullLen": 1, "shortLen": 1, "emptyLen": 0}
	expected.ShouldBeEqual(t, 0, "Names", actual)
}

func Test_Cov5_ReflectType_Name(t *testing.T) {
	result := reflectinternal.ReflectType.Name(42)
	nilResult := reflectinternal.ReflectType.Name(nil)
	actual := args.Map{"val": result, "nil": nilResult}
	expected := args.Map{"val": "int", "nil": ""}
	expected.ShouldBeEqual(t, 0, "Name", actual)
}

func Test_Cov5_ReflectType_NameUsingFmt(t *testing.T) {
	result := reflectinternal.ReflectType.NameUsingFmt(42)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NameUsingFmt", actual)
}

// ── sliceConverter — uncovered branches ──

func Test_Cov5_SliceConverter_Length(t *testing.T) {
	actual := args.Map{
		"slice": reflectinternal.SliceConverter.Length([]int{1, 2}),
		"map":   reflectinternal.SliceConverter.Length(map[string]int{"a": 1}),
		"nil":   reflectinternal.SliceConverter.Length(nil),
		"int":   reflectinternal.SliceConverter.Length(42),
	}
	expected := args.Map{"slice": 2, "map": 1, "nil": 0, "int": 0}
	expected.ShouldBeEqual(t, 0, "SliceConverter Length", actual)
}

func Test_Cov5_SliceConverter_ToStringsRv_Ptr(t *testing.T) {
	s := []int{1, 2}
	rv := reflect.ValueOf(&s)
	result, err := reflectinternal.SliceConverter.ToStringsRv(rv)
	actual := args.Map{"len": len(result), "noErr": err == nil}
	expected := args.Map{"len": 2, "noErr": true}
	expected.ShouldBeEqual(t, 0, "ToStringsRv ptr", actual)
}

func Test_Cov5_SliceConverter_ToStringsRv_NotSlice(t *testing.T) {
	rv := reflect.ValueOf(42)
	_, err := reflectinternal.SliceConverter.ToStringsRv(rv)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ToStringsRv not-slice", actual)
}

func Test_Cov5_SliceConverter_ToStringsRv_Empty(t *testing.T) {
	rv := reflect.ValueOf([]int{})
	result, err := reflectinternal.SliceConverter.ToStringsRv(rv)
	actual := args.Map{"len": len(result), "noErr": err == nil}
	expected := args.Map{"len": 0, "noErr": true}
	expected.ShouldBeEqual(t, 0, "ToStringsRv empty", actual)
}

func Test_Cov5_SliceConverter_ToStrings(t *testing.T) {
	result, err := reflectinternal.SliceConverter.ToStrings([]int{1, 2})
	actual := args.Map{"len": len(result), "noErr": err == nil}
	expected := args.Map{"len": 2, "noErr": true}
	expected.ShouldBeEqual(t, 0, "ToStrings", actual)
}

func Test_Cov5_SliceConverter_ToStringsMust(t *testing.T) {
	result := reflectinternal.SliceConverter.ToStringsMust([]int{1, 2})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "ToStringsMust", actual)
}

func Test_Cov5_SliceConverter_ToStringsRvUsingProcessor(t *testing.T) {
	rv := reflect.ValueOf([]int{1, 2, 3})
	result, err := reflectinternal.SliceConverter.ToStringsRvUsingProcessor(rv,
		func(i int, item any) (string, bool, bool) { return "x", true, i == 1 })
	actual := args.Map{"len": len(result), "noErr": err == nil}
	expected := args.Map{"len": 2, "noErr": true}
	expected.ShouldBeEqual(t, 0, "ToStringsRvUsingProcessor", actual)
}

func Test_Cov5_SliceConverter_ToStringsRvUsingSimpleProcessor(t *testing.T) {
	rv := reflect.ValueOf([]int{1, 2})
	result, err := reflectinternal.SliceConverter.ToStringsRvUsingSimpleProcessor(rv, true,
		func(i int, item any) string {
			if i == 0 {
				return ""
			}
			return "x"
		})
	actual := args.Map{"len": len(result), "noErr": err == nil}
	expected := args.Map{"len": 1, "noErr": true}
	expected.ShouldBeEqual(t, 0, "ToStringsRvUsingSimpleProcessor", actual)
}

func Test_Cov5_SliceConverter_ToAnyItemsAsync(t *testing.T) {
	result := reflectinternal.SliceConverter.ToAnyItemsAsync([]int{1, 2})
	nilResult := reflectinternal.SliceConverter.ToAnyItemsAsync(nil)
	actual := args.Map{"len": len(result), "nilLen": len(nilResult)}
	expected := args.Map{"len": 2, "nilLen": 0}
	expected.ShouldBeEqual(t, 0, "ToAnyItemsAsync", actual)
}

// ── mapConverter — uncovered branches ──

func Test_Cov5_MapConverter_Length(t *testing.T) {
	actual := args.Map{"val": reflectinternal.MapConverter.Length(map[string]int{"a": 1})}
	expected := args.Map{"val": 1}
	expected.ShouldBeEqual(t, 0, "MapConverter Length", actual)
}

func Test_Cov5_MapConverter_ToStringsRv_Ptr(t *testing.T) {
	m := map[string]int{"a": 1}
	rv := reflect.ValueOf(&m)
	result, err := reflectinternal.MapConverter.ToStringsRv(rv)
	actual := args.Map{"len": len(result), "noErr": err == nil}
	expected := args.Map{"len": 1, "noErr": true}
	expected.ShouldBeEqual(t, 0, "MapConverter ToStringsRv ptr", actual)
}

func Test_Cov5_MapConverter_ToStringsRv_NotMap(t *testing.T) {
	rv := reflect.ValueOf(42)
	_, err := reflectinternal.MapConverter.ToStringsRv(rv)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MapConverter ToStringsRv not-map", actual)
}

func Test_Cov5_MapConverter_ToStringsRv_IntKey(t *testing.T) {
	m := map[int]string{1: "a"}
	rv := reflect.ValueOf(m)
	_, err := reflectinternal.MapConverter.ToStringsRv(rv)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MapConverter ToStringsRv int-key", actual)
}

func Test_Cov5_MapConverter_ToKeysStrings(t *testing.T) {
	result, err := reflectinternal.MapConverter.ToKeysStrings(map[string]int{"a": 1})
	actual := args.Map{"len": len(result), "noErr": err == nil}
	expected := args.Map{"len": 1, "noErr": true}
	expected.ShouldBeEqual(t, 0, "ToKeysStrings", actual)
}

func Test_Cov5_MapConverter_ToValuesAny(t *testing.T) {
	result, err := reflectinternal.MapConverter.ToValuesAny(map[string]int{"a": 1})
	nilResult, _ := reflectinternal.MapConverter.ToValuesAny(nil)
	actual := args.Map{"len": len(result), "noErr": err == nil, "nilLen": len(nilResult)}
	expected := args.Map{"len": 1, "noErr": true, "nilLen": 0}
	expected.ShouldBeEqual(t, 0, "ToValuesAny", actual)
}

func Test_Cov5_MapConverter_ToKeysAny(t *testing.T) {
	result, err := reflectinternal.MapConverter.ToKeysAny(map[string]int{"a": 1})
	nilResult, _ := reflectinternal.MapConverter.ToKeysAny(nil)
	actual := args.Map{"len": len(result), "noErr": err == nil, "nilLen": len(nilResult)}
	expected := args.Map{"len": 1, "noErr": true, "nilLen": 0}
	expected.ShouldBeEqual(t, 0, "ToKeysAny", actual)
}

func Test_Cov5_MapConverter_ToKeysValuesAny(t *testing.T) {
	keys, vals, err := reflectinternal.MapConverter.ToKeysValuesAny(map[string]int{"a": 1})
	nilKeys, _, _ := reflectinternal.MapConverter.ToKeysValuesAny(nil)
	actual := args.Map{"keysLen": len(keys), "valsLen": len(vals), "noErr": err == nil, "nilKeysLen": len(nilKeys)}
	expected := args.Map{"keysLen": 1, "valsLen": 1, "noErr": true, "nilKeysLen": 0}
	expected.ShouldBeEqual(t, 0, "ToKeysValuesAny", actual)
}

func Test_Cov5_MapConverter_ToStrings_Nil(t *testing.T) {
	result, err := reflectinternal.MapConverter.ToStrings(nil)
	actual := args.Map{"len": len(result), "noErr": err == nil}
	expected := args.Map{"len": 0, "noErr": true}
	expected.ShouldBeEqual(t, 0, "ToStrings nil", actual)
}

func Test_Cov5_MapConverter_ToStringsMust_ReflectValue(t *testing.T) {
	m := map[string]int{"a": 1}
	rv := reflect.ValueOf(m)
	result := reflectinternal.MapConverter.ToStringsMust(rv)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "ToStringsMust reflect.Value", actual)
}

func Test_Cov5_MapConverter_ToSortedStrings(t *testing.T) {
	result, err := reflectinternal.MapConverter.ToSortedStrings(map[string]int{"b": 2, "a": 1})
	nilResult, _ := reflectinternal.MapConverter.ToSortedStrings(nil)
	actual := args.Map{"len": len(result), "noErr": err == nil, "nilLen": len(nilResult)}
	expected := args.Map{"len": 2, "noErr": true, "nilLen": 0}
	expected.ShouldBeEqual(t, 0, "ToSortedStrings", actual)
}

func Test_Cov5_MapConverter_ToSortedStringsMust(t *testing.T) {
	result := reflectinternal.MapConverter.ToSortedStringsMust(map[string]int{"b": 2, "a": 1})
	nilResult := reflectinternal.MapConverter.ToSortedStringsMust(nil)
	actual := args.Map{"len": len(result), "nilLen": len(nilResult)}
	expected := args.Map{"len": 2, "nilLen": 0}
	expected.ShouldBeEqual(t, 0, "ToSortedStringsMust", actual)
}

func Test_Cov5_MapConverter_ToMapStringAnyRv(t *testing.T) {
	m := map[string]int{"a": 1}
	rv := reflect.ValueOf(m)
	result, err := reflectinternal.MapConverter.ToMapStringAnyRv(rv)
	actual := args.Map{"len": len(result), "noErr": err == nil}
	expected := args.Map{"len": 1, "noErr": true}
	expected.ShouldBeEqual(t, 0, "ToMapStringAnyRv", actual)
}

func Test_Cov5_MapConverter_ToMapStringAnyRv_IntKey(t *testing.T) {
	m := map[int]string{1: "a"}
	rv := reflect.ValueOf(m)
	result, err := reflectinternal.MapConverter.ToMapStringAnyRv(rv)
	actual := args.Map{"len": len(result), "noErr": err == nil}
	expected := args.Map{"len": 1, "noErr": true}
	expected.ShouldBeEqual(t, 0, "ToMapStringAnyRv int-key", actual)
}

func Test_Cov5_MapConverter_ToMapStringAny(t *testing.T) {
	result, err := reflectinternal.MapConverter.ToMapStringAny(map[string]int{"a": 1})
	nilResult, _ := reflectinternal.MapConverter.ToMapStringAny(nil)
	actual := args.Map{"len": len(result), "noErr": err == nil, "nilLen": len(nilResult)}
	expected := args.Map{"len": 1, "noErr": true, "nilLen": 0}
	expected.ShouldBeEqual(t, 0, "ToMapStringAny", actual)
}

// ── isChecker — uncovered branches ──

func Test_Cov5_Is_Conclusive_Slices(t *testing.T) {
	eq, conc := reflectinternal.Is.Conclusive([]int{1}, []int{1})
	actual := args.Map{"eq": eq, "conc": conc}
	expected := args.Map{"eq": false, "conc": false}
	expected.ShouldBeEqual(t, 0, "Conclusive slices", actual)
}

func Test_Cov5_Is_Conclusive_BothNilPtrs(t *testing.T) {
	var a, b *int
	eq, conc := reflectinternal.Is.Conclusive(a, b)
	actual := args.Map{"eq": eq, "conc": conc}
	expected := args.Map{"eq": true, "conc": true}
	expected.ShouldBeEqual(t, 0, "Conclusive both nil ptrs", actual)
}

func Test_Cov5_Is_Conclusive_OneNilPtr(t *testing.T) {
	var a *int
	x := 42
	eq, conc := reflectinternal.Is.Conclusive(a, &x)
	actual := args.Map{"eq": eq, "conc": conc}
	expected := args.Map{"eq": false, "conc": true}
	expected.ShouldBeEqual(t, 0, "Conclusive one nil ptr", actual)
}

func Test_Cov5_Is_AnyEqual_Slices(t *testing.T) {
	actual := args.Map{"val": reflectinternal.Is.AnyEqual([]int{1}, []int{1})}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "AnyEqual slices", actual)
}

func Test_Cov5_Is_NotFunc(t *testing.T) {
	actual := args.Map{
		"int":  reflectinternal.Is.NotFunc(42),
		"nil":  reflectinternal.Is.NotFunc(nil),
	}
	expected := args.Map{"int": true, "nil": true}
	expected.ShouldBeEqual(t, 0, "NotFunc", actual)
}

func Test_Cov5_Is_SliceOrArrayOf(t *testing.T) {
	actual := args.Map{
		"slice": reflectinternal.Is.SliceOrArrayOf(reflect.TypeOf([]int{})),
		"int":   reflectinternal.Is.SliceOrArrayOf(reflect.TypeOf(42)),
	}
	expected := args.Map{"slice": true, "int": false}
	expected.ShouldBeEqual(t, 0, "SliceOrArrayOf", actual)
}

func Test_Cov5_Is_NotNull(t *testing.T) {
	actual := args.Map{
		"val": reflectinternal.Is.NotNull(42),
		"nil": reflectinternal.Is.NotNull(nil),
	}
	expected := args.Map{"val": true, "nil": false}
	expected.ShouldBeEqual(t, 0, "NotNull", actual)
}

func Test_Cov5_Is_Defined(t *testing.T) {
	actual := args.Map{"val": reflectinternal.Is.Defined(42)}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "Defined", actual)
}

func Test_Cov5_Is_ZeroRv_Array(t *testing.T) {
	rv := reflect.ValueOf([2]int{0, 0})
	actual := args.Map{"zero": reflectinternal.Is.ZeroRv(rv)}
	expected := args.Map{"zero": true}
	expected.ShouldBeEqual(t, 0, "ZeroRv array", actual)
}

func Test_Cov5_Is_ZeroRv_Struct(t *testing.T) {
	type S struct{ A int }
	rv := reflect.ValueOf(S{})
	actual := args.Map{"zero": reflectinternal.Is.ZeroRv(rv)}
	expected := args.Map{"zero": true}
	expected.ShouldBeEqual(t, 0, "ZeroRv struct", actual)
}

func Test_Cov5_Is_InterfaceRv(t *testing.T) {
	rv := reflect.ValueOf(42)
	actual := args.Map{"val": reflectinternal.Is.InterfaceRv(rv)}
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "InterfaceRv", actual)
}

func Test_Cov5_Is_Interface(t *testing.T) {
	actual := args.Map{"val": reflectinternal.Is.Interface(42)}
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "Interface", actual)
}

func Test_Cov5_Is_StructRv(t *testing.T) {
	type S struct{}
	rv := reflect.ValueOf(S{})
	actual := args.Map{"val": reflectinternal.Is.StructRv(rv)}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "StructRv", actual)
}

// ── looper — uncovered branches ──

func Test_Cov5_Looper_FieldsFor(t *testing.T) {
	type S struct{ A int; B string }
	count := 0
	err := reflectinternal.Looper.FieldsFor(S{}, func(fp *reflectmodel.FieldProcessor) error {
		return nil
	})
	_ = count
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "FieldsFor", actual)
}

func Test_Cov5_Looper_FieldNames(t *testing.T) {
	type S struct{ A int; B string }
	names, err := reflectinternal.Looper.FieldNames(S{})
	actual := args.Map{"len": len(names), "noErr": err == nil}
	expected := args.Map{"len": 2, "noErr": true}
	expected.ShouldBeEqual(t, 0, "FieldNames", actual)
}

func Test_Cov5_Looper_FieldsMap(t *testing.T) {
	type S struct{ A int }
	m, err := reflectinternal.Looper.FieldsMap(S{})
	actual := args.Map{"len": len(m), "noErr": err == nil}
	expected := args.Map{"len": 1, "noErr": true}
	expected.ShouldBeEqual(t, 0, "FieldsMap", actual)
}

func Test_Cov5_Looper_MethodsMap(t *testing.T) {
	type S struct{}
	m, err := reflectinternal.Looper.MethodsMap(S{})
	actual := args.Map{"notNil": m != nil, "noErr": err == nil}
	expected := args.Map{"notNil": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "MethodsMap", actual)
}

func Test_Cov5_Looper_ReducePointer(t *testing.T) {
	x := 42
	result := reflectinternal.Looper.ReducePointer(&x, 3)
	actual := args.Map{"valid": result.IsValid}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "ReducePointer", actual)
}

func Test_Cov5_Looper_ReducePointerDefault(t *testing.T) {
	result := reflectinternal.Looper.ReducePointerDefault(42)
	actual := args.Map{"valid": result.IsValid}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "ReducePointerDefault", actual)
}

func Test_Cov5_Looper_Slice(t *testing.T) {
	err := reflectinternal.Looper.Slice([]int{1, 2}, func(total, index int, item any) error { return nil })
	nilErr := reflectinternal.Looper.Slice(nil, func(total, index int, item any) error { return nil })
	actual := args.Map{"noErr": err == nil, "nilNoErr": nilErr == nil}
	expected := args.Map{"noErr": true, "nilNoErr": true}
	expected.ShouldBeEqual(t, 0, "Slice", actual)
}

func Test_Cov5_Looper_Map(t *testing.T) {
	err := reflectinternal.Looper.Map(map[string]int{"a": 1}, func(total, index int, key, value any) error { return nil })
	nilErr := reflectinternal.Looper.Map(nil, func(total, index int, key, value any) error { return nil })
	actual := args.Map{"noErr": err == nil, "nilNoErr": nilErr == nil}
	expected := args.Map{"noErr": true, "nilNoErr": true}
	expected.ShouldBeEqual(t, 0, "Map", actual)
}

func Test_Cov5_Looper_ToPointerReflectValue(t *testing.T) {
	type S struct{}
	_, err := reflectinternal.Looper.ToPointerReflectValue(S{})
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ToPointerReflectValue struct", actual)
}

func Test_Cov5_Looper_ToPointerReflectValue_Invalid(t *testing.T) {
	_, err := reflectinternal.Looper.ToPointerReflectValue(42)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ToPointerReflectValue invalid", actual)
}

// ── codeStack — uncovered ──

func Test_Cov5_CodeStack_NewDefault(t *testing.T) {
	st := reflectinternal.CodeStack.NewDefault()
	actual := args.Map{"isOkay": st.IsOkay}
	expected := args.Map{"isOkay": true}
	expected.ShouldBeEqual(t, 0, "NewDefault", actual)
}

func Test_Cov5_CodeStack_LastFileWithLine(t *testing.T) {
	result := reflectinternal.CodeStack.LastFileWithLine(0, 2)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "LastFileWithLine", actual)
}

func Test_Cov5_CodeStack_NewStacks(t *testing.T) {
	result := reflectinternal.CodeStack.NewStacks(0, 2)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NewStacks", actual)
}

func Test_Cov5_CodeStack_StacksStrings(t *testing.T) {
	result := reflectinternal.CodeStack.StacksStrings(0)
	actual := args.Map{"gt0": len(result) > 0}
	expected := args.Map{"gt0": true}
	expected.ShouldBeEqual(t, 0, "StacksStrings", actual)
}

func Test_Cov5_CodeStack_StacksStringsCount(t *testing.T) {
	result := reflectinternal.CodeStack.StacksStringsCount(0, 2)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "StacksStringsCount", actual)
}

func Test_Cov5_CodeStack_StacksStringsFiltered(t *testing.T) {
	result := reflectinternal.CodeStack.StacksStringsFiltered(0, 4)
	actual := args.Map{"gt0": len(result) > 0}
	expected := args.Map{"gt0": true}
	expected.ShouldBeEqual(t, 0, "StacksStringsFiltered", actual)
}

func Test_Cov5_CodeStack_StacksString(t *testing.T) {
	result := reflectinternal.CodeStack.StacksString(0)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StacksString", actual)
}

func Test_Cov5_CodeStack_StacksStringDefault(t *testing.T) {
	result := reflectinternal.CodeStack.StacksStringDefault(0)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StacksStringDefault", actual)
}

func Test_Cov5_CodeStack_StacksStringCount(t *testing.T) {
	result := reflectinternal.CodeStack.StacksStringCount(0, 2)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StacksStringCount", actual)
}

func Test_Cov5_CodeStack_SingleStack(t *testing.T) {
	result := reflectinternal.CodeStack.SingleStack(0)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SingleStack", actual)
}

// ── reflectPath ──

func Test_Cov5_Path_CurDir(t *testing.T) {
	result := reflectinternal.Path.CurDir()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Path CurDir", actual)
}

func Test_Cov5_Path_RepoDir(t *testing.T) {
	result := reflectinternal.Path.RepoDir()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Path RepoDir", actual)
}

// ── getFunc — uncovered ──

func Test_Cov5_GetFunc_NameOnlyByStack(t *testing.T) {
	result := reflectinternal.GetFunc.NameOnlyByStack(0)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NameOnlyByStack", actual)
}

func Test_Cov5_GetFunc_GetMethodsRv(t *testing.T) {
	type S struct{}
	rv := reflect.ValueOf(S{})
	result := reflectinternal.GetFunc.GetMethodsRv(rv)
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "GetMethodsRv", actual)
}

func Test_Cov5_GetFunc_GetMethodsMapRv(t *testing.T) {
	type S struct{}
	rv := reflect.ValueOf(S{})
	result := reflectinternal.GetFunc.GetMethodsMapRv(rv)
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "GetMethodsMapRv", actual)
}

func Test_Cov5_GetFunc_GetMethodProcessorsMap(t *testing.T) {
	type S struct{}
	rv := reflect.ValueOf(S{})
	result := reflectinternal.GetFunc.GetMethodProcessorsMap(rv)
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "GetMethodProcessorsMap", actual)
}

func Test_Cov5_GetFunc_GetPkgPath(t *testing.T) {
	fn := func() {}
	result := reflectinternal.GetFunc.GetPkgPath(fn)
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "GetPkgPath", actual)
}

// ── Utils uncovered ──

func Test_Cov5_Utils_VerifyReflectTypes_Mismatch(t *testing.T) {
	ok, err := reflectinternal.Utils.VerifyReflectTypes(
		"TestRoot",
		[]reflect.Type{reflect.TypeOf(0), reflect.TypeOf("")},
		[]reflect.Type{reflect.TypeOf(0), reflect.TypeOf(0)},
	)
	actual := args.Map{"ok": ok, "hasErr": err != nil}
	expected := args.Map{"ok": false, "hasErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyReflectTypes mismatch", actual)
}

func Test_Cov5_Utils_FullNameToPkgName(t *testing.T) {
	result := reflectinternal.Utils.FullNameToPkgName("mypackage.MyFunc")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "FullNameToPkgName", actual)
}

// ── TypeNameToValidVariableName — more branches ──

func Test_Cov5_TypeNameToValidVariableName_SliceBrackets(t *testing.T) {
	result := reflectinternal.TypeNameToValidVariableName("[]MyType")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TypeNameToValidVariableName slice", actual)
}

func Test_Cov5_TypeNameToValidVariableName_PtrStar(t *testing.T) {
	result := reflectinternal.TypeNameToValidVariableName("*MyType")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TypeNameToValidVariableName ptr", actual)
}

func Test_Cov5_TypeNameToValidVariableName_DotWithSlice(t *testing.T) {
	result := reflectinternal.TypeNameToValidVariableName("[]pkg.MyType")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TypeNameToValidVariableName dot-slice", actual)
}

func Test_Cov5_TypeNameToValidVariableName_DotWithPtr(t *testing.T) {
	result := reflectinternal.TypeNameToValidVariableName("*pkg.MyType")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TypeNameToValidVariableName dot-ptr", actual)
}
