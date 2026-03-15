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
	expected.ShouldBeEqual(t, 0, "Writer nil", actual)
}

func Test_Cov_Writer_String(t *testing.T) {
	// Act
	result := coreproperty.Writer.Write("hello")

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Writer string", actual)
}

func Test_Cov_Writer_Bool(t *testing.T) {
	// Act
	actual := args.Map{"result": coreproperty.Writer.Write(true)}

	// Assert
	expected := args.Map{"result": "true"}
	expected.ShouldBeEqual(t, 0, "Writer bool", actual)
}

func Test_Cov_Writer_Int(t *testing.T) {
	// Act
	actual := args.Map{"result": coreproperty.Writer.Write(42)}

	// Assert
	expected := args.Map{"result": "42"}
	expected.ShouldBeEqual(t, 0, "Writer int", actual)
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
	expected.ShouldBeEqual(t, 0, "Writer struct", actual)
}

func Test_Cov_Writer_Slice(t *testing.T) {
	// Act
	result := coreproperty.Writer.Write([]string{"a", "b"})

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Writer slice", actual)
}

func Test_Cov_Writer_Map(t *testing.T) {
	// Act
	result := coreproperty.Writer.Write(map[string]int{"a": 1})

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Writer map", actual)
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
	expected.ShouldBeEqual(t, 0, "Writer pointer", actual)
}

func Test_Cov_Writer_ArgsString(t *testing.T) {
	// Act
	result := coreproperty.Writer.Write(args.String("test"))

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": "test"}
	expected.ShouldBeEqual(t, 0, "Writer args.String", actual)
}

func Test_Cov_Writer_WritePropertyOptions_SubRequest(t *testing.T) {
	// Act
	result := coreproperty.Writer.WritePropertyOptions(true, []int{1, 2})

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "WritePropertyOptions sub-request", actual)
}

func Test_Cov_Writer_Byte(t *testing.T) {
	// Act
	result := coreproperty.Writer.Write(byte(5))

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": "5"}
	expected.ShouldBeEqual(t, 0, "Writer byte", actual)
}

func Test_Cov_Writer_Float64(t *testing.T) {
	// Act
	result := coreproperty.Writer.Write(float64(3.14))

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Writer float64", actual)
}
