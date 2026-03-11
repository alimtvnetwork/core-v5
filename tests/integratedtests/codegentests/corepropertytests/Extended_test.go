package corepropertytests

import (
	"testing"

	"github.com/alimtvnetwork/core/codegen/coreproperty"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_Write_Nil(t *testing.T) {
	// Arrange & Act
	result := coreproperty.Writer.Write(nil)

	// Assert
	if result != "nil" {
		t.Errorf("expected 'nil', got '%s'", result)
	}
}

func Test_Write_String(t *testing.T) {
	// Arrange & Act
	result := coreproperty.Writer.Write("hello")

	// Assert
	if result != `"hello"` {
		t.Errorf("expected '\"hello\"', got '%s'", result)
	}
}

func Test_Write_Bool(t *testing.T) {
	// Arrange & Act
	result := coreproperty.Writer.Write(true)

	// Assert
	if result != "true" {
		t.Errorf("expected 'true', got '%s'", result)
	}
}

func Test_Write_Int(t *testing.T) {
	// Arrange & Act
	result := coreproperty.Writer.Write(42)

	// Assert
	if result != "42" {
		t.Errorf("expected '42', got '%s'", result)
	}
}

func Test_Write_Float(t *testing.T) {
	// Arrange & Act
	result := coreproperty.Writer.Write(3.14)

	// Assert
	if result != "3.14" {
		t.Errorf("expected '3.14', got '%s'", result)
	}
}

func Test_Write_Byte(t *testing.T) {
	// Arrange & Act
	result := coreproperty.Writer.Write(byte(65))

	// Assert
	if result != "65" {
		t.Errorf("expected '65', got '%s'", result)
	}
}

func Test_Write_Int32(t *testing.T) {
	// Arrange & Act
	result := coreproperty.Writer.Write(int32(100))

	// Assert
	if result != "100" {
		t.Errorf("expected '100', got '%s'", result)
	}
}

func Test_Write_Int64(t *testing.T) {
	// Arrange & Act
	result := coreproperty.Writer.Write(int64(200))

	// Assert
	if result != "200" {
		t.Errorf("expected '200', got '%s'", result)
	}
}

func Test_Write_ArgsString(t *testing.T) {
	// Arrange & Act
	result := coreproperty.Writer.Write(args.String("test"))

	// Assert
	if result != "test" {
		t.Errorf("expected 'test', got '%s'", result)
	}
}

func Test_Write_Slice(t *testing.T) {
	// Arrange & Act
	result := coreproperty.Writer.Write([]string{"a", "b"})

	// Assert
	if result == "" || result == "nil" {
		t.Error("expected non-empty non-nil result for slice")
	}
}

func Test_Write_Map(t *testing.T) {
	// Arrange
	m := map[string]int{"a": 1}

	// Act
	result := coreproperty.Writer.Write(m)

	// Assert
	if result == "" || result == "nil" {
		t.Error("expected non-empty non-nil result for map")
	}
}

func Test_Write_Struct(t *testing.T) {
	// Arrange
	type sample struct {
		Name string
		Age  int
	}
	s := sample{Name: "test", Age: 30}

	// Act
	result := coreproperty.Writer.Write(s)

	// Assert
	if result == "" || result == "nil" {
		t.Error("expected non-empty non-nil result for struct")
	}
}

func Test_Write_Pointer(t *testing.T) {
	// Arrange
	val := "hello"
	ptr := &val

	// Act
	result := coreproperty.Writer.Write(ptr)

	// Assert
	if result == "" || result == "nil" {
		t.Error("expected non-empty non-nil result for pointer")
	}
}

func Test_Write_NilPointer(t *testing.T) {
	// Arrange
	var ptr *string

	// Act
	result := coreproperty.Writer.Write(ptr)

	// Assert
	if result != "nil" {
		t.Errorf("expected 'nil', got '%s'", result)
	}
}

func Test_WritePropertyOptions_SubRequest(t *testing.T) {
	// Arrange & Act
	result := coreproperty.Writer.WritePropertyOptions(true, []int{1, 2})

	// Assert
	if result == "" || result == "nil" {
		t.Error("expected non-empty result for sub-request slice")
	}
}

func Test_WriteStruct_Nil(t *testing.T) {
	// Arrange & Act
	result := coreproperty.Writer.WriteStruct(nil)

	// Assert
	if result != "nil" {
		t.Errorf("expected 'nil', got '%s'", result)
	}
}

func Test_Write_EmptySlice(t *testing.T) {
	// Arrange & Act
	result := coreproperty.Writer.Write([]string{})

	// Assert
	if result == "nil" {
		t.Error("expected non-nil result for empty slice")
	}
}

func Test_Write_EmptyMap(t *testing.T) {
	// Arrange & Act
	result := coreproperty.Writer.Write(map[string]int{})

	// Assert
	if result == "nil" {
		t.Error("expected non-nil result for empty map")
	}
}

func Test_WritePropertyOptions_Nil(t *testing.T) {
	// Arrange & Act
	result := coreproperty.Writer.WritePropertyOptions(false, nil)

	// Assert
	if result != "nil" {
		t.Errorf("expected 'nil', got '%s'", result)
	}
}

func Test_Write_Int8(t *testing.T) {
	// Arrange & Act
	result := coreproperty.Writer.Write(int8(10))

	// Assert
	if result != "10" {
		t.Errorf("expected '10', got '%s'", result)
	}
}

func Test_Write_Uint16(t *testing.T) {
	// Arrange & Act
	result := coreproperty.Writer.Write(uint16(500))

	// Assert
	if result != "500" {
		t.Errorf("expected '500', got '%s'", result)
	}
}

func Test_Write_Uint32(t *testing.T) {
	// Arrange & Act
	result := coreproperty.Writer.Write(uint32(1000))

	// Assert
	if result != "1000" {
		t.Errorf("expected '1000', got '%s'", result)
	}
}

func Test_Write_Uint64(t *testing.T) {
	// Arrange & Act
	result := coreproperty.Writer.Write(uint64(2000))

	// Assert
	if result != "2000" {
		t.Errorf("expected '2000', got '%s'", result)
	}
}

func Test_Write_Float32(t *testing.T) {
	// Arrange & Act
	result := coreproperty.Writer.Write(float32(1.5))

	// Assert
	if result != "1.5" {
		t.Errorf("expected '1.5', got '%s'", result)
	}
}
