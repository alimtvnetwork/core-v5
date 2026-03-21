package corepropertytests

import (
	"testing"

	"github.com/alimtvnetwork/core/codegen/coreproperty"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_Cov_Writer_Nil(t *testing.T) {
	// Act
	actual := args.Map{"result": coreproperty.Writer.Write(nil)}

	// Assert
	expected := args.Map{"result": "nil"}
	expected.ShouldBeEqual(t, 0, "Writer returns nil -- nil", actual)
}

func Test_Cov_Writer_String(t *testing.T) {
	// Act
	result := coreproperty.Writer.Write("hello")

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Writer returns correct value -- string", actual)
}

func Test_Cov_Writer_Bool(t *testing.T) {
	// Act
	actual := args.Map{"result": coreproperty.Writer.Write(true)}

	// Assert
	expected := args.Map{"result": "true"}
	expected.ShouldBeEqual(t, 0, "Writer returns correct value -- bool", actual)
}

func Test_Cov_Writer_Int(t *testing.T) {
	// Act
	actual := args.Map{"result": coreproperty.Writer.Write(42)}

	// Assert
	expected := args.Map{"result": "42"}
	expected.ShouldBeEqual(t, 0, "Writer returns correct value -- int", actual)
}

func Test_Cov_Writer_Struct(t *testing.T) {
	// Arrange
	type TestStruct struct {
		Name string
		Age  int
	}

	// Act
	result := coreproperty.Writer.Write(TestStruct{Name: "Alice", Age: 30})

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Writer returns correct value -- struct", actual)
}

func Test_Cov_Writer_Slice(t *testing.T) {
	// Act
	result := coreproperty.Writer.Write([]string{"a", "b"})

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Writer returns correct value -- slice", actual)
}

func Test_Cov_Writer_Map(t *testing.T) {
	// Act
	result := coreproperty.Writer.Write(map[string]int{"a": 1})

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Writer returns correct value -- map", actual)
}

func Test_Cov_Writer_Pointer(t *testing.T) {
	// Arrange
	val := "hello"
	var nilPtr *string

	// Act
	result := coreproperty.Writer.Write(&val)
	nilResult := coreproperty.Writer.Write(nilPtr)

	// Assert
	actual := args.Map{"notEmpty": result != "", "nilResult": nilResult}
	expected := args.Map{"notEmpty": true, "nilResult": "nil"}
	expected.ShouldBeEqual(t, 0, "Writer returns correct value -- pointer", actual)
}

func Test_Cov_Writer_ArgsString(t *testing.T) {
	// Act
	result := coreproperty.Writer.Write(args.String("test"))

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": "test"}
	expected.ShouldBeEqual(t, 0, "Writer returns correct value -- args.String", actual)
}

func Test_Cov_Writer_WritePropertyOptions_SubRequest(t *testing.T) {
	// Act
	result := coreproperty.Writer.WritePropertyOptions(true, []int{1, 2})

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "WritePropertyOptions returns correct value -- sub-request", actual)
}

func Test_Cov_Writer_Byte(t *testing.T) {
	// Act
	result := coreproperty.Writer.Write(byte(5))

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": "5"}
	expected.ShouldBeEqual(t, 0, "Writer returns correct value -- byte", actual)
}

func Test_Cov_Writer_Float64(t *testing.T) {
	// Act
	result := coreproperty.Writer.Write(float64(3.14))

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Writer returns correct value -- float64", actual)
}
