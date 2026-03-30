package reflectmodeltests

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/alimtvnetwork/core/reflectcore/reflectmodel"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage13 — reflectcore/reflectmodel gaps (Iteration 16)
//
// Targets:
//   - MethodProcessor.IsEqual branches (IsInvalid, Name, IsPublic, ArgsCount, ReturnLength)
//   - MethodProcessor.GetInArgsTypes/GetInArgsTypesNames zero-args path
//   - MethodProcessor.validationError (invalid receiver)
//   - isNull (all branches via rvUtils.IsNull)
//   - rvUtils: ArgsToReflectValues empty, ReflectValueToAnyValue nil,
//     InterfacesToTypes empty, IsReflectTypeMatchAny,
//     PrependWithSpaces with prependingLinesSpaceCount > 0,
//     WithSpaces empty
// ══════════════════════════════════════════════════════════════════════════════

// ---------- MethodProcessor.IsEqual branches ----------

func Test_Cov13_IsEqual_DifferentIsInvalid(t *testing.T) {
	// Arrange
	a := newMethodProcessor("PublicMethod")
	var b *reflectmodel.MethodProcessor // nil = invalid

	// Act
	result := a.IsEqual(b)

	// Assert
	if result {
		t.Fatal("expected false when IsInvalid differs")
	}
}

func Test_Cov13_IsEqual_DifferentName(t *testing.T) {
	// Arrange
	a := newMethodProcessor("PublicMethod")
	b := newMethodProcessor("NoArgsMethod")

	// Act
	result := a.IsEqual(b)

	// Assert
	if result {
		t.Fatal("expected false when names differ")
	}
}

func Test_Cov13_IsEqual_DifferentArgsCount(t *testing.T) {
	// Arrange
	a := newMethodProcessor("PublicMethod")  // 3 in-args (receiver + string + int)
	b := newMethodProcessor("NoArgsMethod")  // 1 in-arg (receiver only)
	// Force same name so we get past name check
	b.Name = a.Name

	// Act
	result := a.IsEqual(b)

	// Assert
	if result {
		t.Fatal("expected false when ArgsCount differs")
	}
}

func Test_Cov13_IsEqual_DifferentReturnLength(t *testing.T) {
	// Arrange
	a := newMethodProcessor("PublicMethod") // returns (string, error) = 2
	b := newMethodProcessor("MultiReturn")  // returns (int, string, error) = 3
	// Force same name and args count
	b.Name = a.Name

	// Act
	result := a.IsEqual(b)

	// Assert
	if result {
		t.Fatal("expected false when ReturnLength differs")
	}
}

func Test_Cov13_IsEqual_BothEqual(t *testing.T) {
	// Arrange
	a := newMethodProcessor("PublicMethod")
	b := newMethodProcessor("PublicMethod")

	// Act
	result := a.IsEqual(b)

	// Assert
	if !result {
		t.Fatal("expected true for identical methods")
	}
}

// ---------- GetInArgsTypes / GetInArgsTypesNames zero-args path ----------

func Test_Cov13_GetInArgsTypes_ZeroArgs(t *testing.T) {
	// Arrange — NoArgsMethod has 1 in-arg (receiver), not 0.
	// We need a method with 0 NumIn... but Go methods always have receiver.
	// The coverage line 229-231 is argsCount==0 returning empty.
	// This path is hit when ArgsCount() == 0, which means NumIn() == 0.
	// With real reflect.Method that's impossible (receiver), so this is dead code.
	// Documenting: unreachable for real methods.
	t.Skip("ArgsCount==0 is unreachable for real reflect.Method (always has receiver)")
}

func Test_Cov13_GetInArgsTypesNames_ZeroArgs(t *testing.T) {
	t.Skip("ArgsCount==0 is unreachable for real reflect.Method (always has receiver)")
}

// ---------- validationError via Invoke on invalid ----------

func Test_Cov13_Invoke_InvalidProcessor(t *testing.T) {
	// Arrange
	var mp *reflectmodel.MethodProcessor // nil

	// Act
	_, err := mp.Invoke()

	// Assert
	if err == nil {
		t.Fatal("expected error from nil MethodProcessor.Invoke")
	}
}

func Test_Cov13_Invoke_InvalidNonNilProcessor(t *testing.T) {
	// Arrange — create a MethodProcessor with zero-value ReflectMethod (IsInvalid returns false
	// since pointer is non-nil, but the method type is invalid)
	// This should trigger the validationError IsInvalid branch (line 276-284).
	// Actually IsInvalid checks `it == nil`, so non-nil is valid.
	// The validationError IsInvalid branch needs it.IsInvalid() == true,
	// which only happens when it == nil, already covered above.
	t.Skip("validationError IsInvalid branch requires nil receiver, covered by nil Invoke test")
}

// ---------- rvUtils.IsNull (covers isNull.go + utils.go IsNull) ----------

func Test_Cov13_IsNull_Nil(t *testing.T) {
	// Arrange
	var utils reflectmodel.RvUtils

	// Act
	result := utils.IsNull(nil)

	// Assert
	if !result {
		t.Fatal("expected true for nil")
	}
}

func Test_Cov13_IsNull_NilMap(t *testing.T) {
	// Arrange
	var utils reflectmodel.RvUtils
	var m map[string]string

	// Act
	result := utils.IsNull(m)

	// Assert
	if !result {
		t.Fatal("expected true for nil map")
	}
}

func Test_Cov13_IsNull_NilSlice(t *testing.T) {
	// Arrange
	var utils reflectmodel.RvUtils
	var s []int

	// Act
	result := utils.IsNull(s)

	// Assert
	if !result {
		t.Fatal("expected true for nil slice")
	}
}

func Test_Cov13_IsNull_NilPtr(t *testing.T) {
	// Arrange
	var utils reflectmodel.RvUtils
	var p *int

	// Act
	result := utils.IsNull(p)

	// Assert
	if !result {
		t.Fatal("expected true for nil pointer")
	}
}

func Test_Cov13_IsNull_NilChan(t *testing.T) {
	// Arrange
	var utils reflectmodel.RvUtils
	var ch chan int

	// Act
	result := utils.IsNull(ch)

	// Assert
	if !result {
		t.Fatal("expected true for nil chan")
	}
}

func Test_Cov13_IsNull_NilFunc(t *testing.T) {
	// Arrange
	var utils reflectmodel.RvUtils
	var fn func()

	// Act
	result := utils.IsNull(fn)

	// Assert
	if !result {
		t.Fatal("expected true for nil func")
	}
}

func Test_Cov13_IsNull_NilUnsafePointer(t *testing.T) {
	// Arrange
	var utils reflectmodel.RvUtils
	var up unsafe.Pointer

	// Act
	result := utils.IsNull(up)

	// Assert
	if !result {
		t.Fatal("expected true for nil unsafe.Pointer")
	}
}

func Test_Cov13_IsNull_NonNilValue(t *testing.T) {
	// Arrange
	var utils reflectmodel.RvUtils

	// Act
	result := utils.IsNull(42)

	// Assert
	if result {
		t.Fatal("expected false for non-nil int")
	}
}

// ---------- rvUtils.ArgsToReflectValues empty ----------

func Test_Cov13_ArgsToReflectValues_Empty(t *testing.T) {
	// Arrange
	var utils reflectmodel.RvUtils

	// Act
	result := utils.ArgsToReflectValues([]any{})

	// Assert
	if len(result) != 0 {
		t.Fatal("expected empty slice")
	}
}

// ---------- rvUtils.ReflectValueToAnyValue nil ----------

func Test_Cov13_ReflectValueToAnyValue_NilValue(t *testing.T) {
	// Arrange
	var utils reflectmodel.RvUtils

	// Act — pass nil which IsNull returns true
	result := utils.ReflectValueToAnyValue(reflect.Value{})

	// Assert
	if result != nil {
		t.Fatal("expected nil for zero reflect.Value")
	}
}

// ---------- rvUtils.InterfacesToTypes empty ----------

func Test_Cov13_InterfacesToTypes_Empty(t *testing.T) {
	// Arrange
	var utils reflectmodel.RvUtils

	// Act
	result := utils.InterfacesToTypes([]any{})

	// Assert
	if len(result) != 0 {
		t.Fatal("expected empty slice")
	}
}

// ---------- rvUtils.IsReflectTypeMatchAny ----------

func Test_Cov13_IsReflectTypeMatchAny_Match(t *testing.T) {
	// Arrange
	var utils reflectmodel.RvUtils

	// Act
	ok, err := utils.IsReflectTypeMatchAny("hello", "world")

	// Assert
	if !ok {
		t.Fatal("expected match for same types")
	}
	if err != nil {
		t.Fatal("expected no error")
	}
}

func Test_Cov13_IsReflectTypeMatchAny_Mismatch(t *testing.T) {
	// Arrange
	var utils reflectmodel.RvUtils

	// Act
	ok, err := utils.IsReflectTypeMatchAny("hello", 42)

	// Assert
	if ok {
		t.Fatal("expected no match for different types")
	}
	if err == nil {
		t.Fatal("expected error for type mismatch")
	}
}

// ---------- rvUtils.PrependWithSpaces with prependingLinesSpaceCount > 0 ----------

func Test_Cov13_PrependWithSpaces_WithPrependingSpaces(t *testing.T) {
	// Arrange
	var utils reflectmodel.RvUtils
	existing := []string{"line1", "line2"}

	// Act
	result := utils.PrependWithSpaces(
		2,
		existing,
		4,
		"header",
	)

	// Assert
	if len(result) != 3 {
		t.Fatalf("expected 3 lines, got %d", len(result))
	}
	if result[0] != "    header" {
		t.Fatalf("expected prepending line with 4 spaces, got %q", result[0])
	}
}

// ---------- rvUtils.WithSpaces empty ----------

func Test_Cov13_WithSpaces_Empty(t *testing.T) {
	// Arrange
	var utils reflectmodel.RvUtils

	// Act
	result := utils.WithSpaces(4)

	// Assert
	if len(result) != 0 {
		t.Fatal("expected empty slice")
	}
}
