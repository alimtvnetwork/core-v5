package conditionaltests

import (
	"testing"

	"github.com/alimtvnetwork/core/conditional"
)

// ============================================================================
// typed_float32.go — all 11 functions
// ============================================================================

func Test_IfFloat32_Cov4(t *testing.T) {
	if r := conditional.IfFloat32(true, 1.5, 2.5); r != 1.5 {
		t.Errorf("expected 1.5, got %v", r)
	}
	if r := conditional.IfFloat32(false, 1.5, 2.5); r != 2.5 {
		t.Errorf("expected 2.5, got %v", r)
	}
}

func Test_IfFuncFloat32_Cov4(t *testing.T) {
	r := conditional.IfFuncFloat32(true, func() float32 { return 1.0 }, func() float32 { return 2.0 })
	if r != 1.0 {
		t.Error("expected 1.0")
	}
	r = conditional.IfFuncFloat32(false, func() float32 { return 1.0 }, func() float32 { return 2.0 })
	if r != 2.0 {
		t.Error("expected 2.0")
	}
}

func Test_IfTrueFuncFloat32_Cov4(t *testing.T) {
	if r := conditional.IfTrueFuncFloat32(true, func() float32 { return 3.14 }); r != 3.14 {
		t.Error("expected 3.14")
	}
	if r := conditional.IfTrueFuncFloat32(false, func() float32 { return 3.14 }); r != 0 {
		t.Error("expected 0")
	}
}

func Test_IfSliceFloat32_Cov4(t *testing.T) {
	r := conditional.IfSliceFloat32(true, []float32{1.0}, []float32{2.0, 3.0})
	if len(r) != 1 {
		t.Error("expected 1")
	}
	r = conditional.IfSliceFloat32(false, []float32{1.0}, []float32{2.0, 3.0})
	if len(r) != 2 {
		t.Error("expected 2")
	}
}

func Test_IfPtrFloat32_Cov4(t *testing.T) {
	a, b := float32(1.0), float32(2.0)
	if *conditional.IfPtrFloat32(true, &a, &b) != 1.0 {
		t.Error("expected 1.0")
	}
	if *conditional.IfPtrFloat32(false, &a, &b) != 2.0 {
		t.Error("expected 2.0")
	}
}

func Test_NilDefFloat32_Cov4(t *testing.T) {
	if r := conditional.NilDefFloat32(nil, 5.5); r != 5.5 {
		t.Error("expected 5.5")
	}
	v := float32(3.3)
	if r := conditional.NilDefFloat32(&v, 5.5); r != 3.3 {
		t.Error("expected 3.3")
	}
}

func Test_NilDefPtrFloat32_Cov4(t *testing.T) {
	r := conditional.NilDefPtrFloat32(nil, 9.9)
	if *r != 9.9 {
		t.Error("expected 9.9")
	}
	v := float32(1.1)
	r = conditional.NilDefPtrFloat32(&v, 9.9)
	if *r != 1.1 {
		t.Error("expected 1.1")
	}
}

func Test_ValueOrZeroFloat32_Cov4(t *testing.T) {
	if r := conditional.ValueOrZeroFloat32(nil); r != 0 {
		t.Error("expected 0")
	}
	v := float32(7.7)
	if r := conditional.ValueOrZeroFloat32(&v); r != 7.7 {
		t.Error("expected 7.7")
	}
}

func Test_PtrOrZeroFloat32_Cov4(t *testing.T) {
	r := conditional.PtrOrZeroFloat32(nil)
	if *r != 0 {
		t.Error("expected 0")
	}
	v := float32(4.4)
	r = conditional.PtrOrZeroFloat32(&v)
	if *r != 4.4 {
		t.Error("expected 4.4")
	}
}

func Test_NilValFloat32_Cov4(t *testing.T) {
	if r := conditional.NilValFloat32(nil, 1.0, 2.0); r != 1.0 {
		t.Error("expected 1.0")
	}
	v := float32(5.0)
	if r := conditional.NilValFloat32(&v, 1.0, 2.0); r != 2.0 {
		t.Error("expected 2.0")
	}
}

func Test_NilValPtrFloat32_Cov4(t *testing.T) {
	r := conditional.NilValPtrFloat32(nil, 1.0, 2.0)
	if *r != 1.0 {
		t.Error("expected 1.0")
	}
	v := float32(5.0)
	r = conditional.NilValPtrFloat32(&v, 1.0, 2.0)
	if *r != 2.0 {
		t.Error("expected 2.0")
	}
}

// ============================================================================
// typed_float64.go — NilDef, NilDefPtr, ValueOrZero, PtrOrZero, NilVal, NilValPtr
// (IfFloat64, IfFuncFloat64, IfTrueFuncFloat64 already covered by TypedWrappers)
// ============================================================================

func Test_NilDefFloat64_Cov4(t *testing.T) {
	if r := conditional.NilDefFloat64(nil, 5.5); r != 5.5 {
		t.Error("expected 5.5")
	}
	v := 3.3
	if r := conditional.NilDefFloat64(&v, 5.5); r != 3.3 {
		t.Error("expected 3.3")
	}
}

func Test_NilDefPtrFloat64_Cov4(t *testing.T) {
	r := conditional.NilDefPtrFloat64(nil, 9.9)
	if *r != 9.9 {
		t.Error("expected 9.9")
	}
}

func Test_ValueOrZeroFloat64_Cov4(t *testing.T) {
	if r := conditional.ValueOrZeroFloat64(nil); r != 0 {
		t.Error("expected 0")
	}
	v := 7.7
	if r := conditional.ValueOrZeroFloat64(&v); r != 7.7 {
		t.Error("expected 7.7")
	}
}

func Test_PtrOrZeroFloat64_Cov4(t *testing.T) {
	r := conditional.PtrOrZeroFloat64(nil)
	if *r != 0 {
		t.Error("expected 0")
	}
}

func Test_NilValFloat64_Cov4(t *testing.T) {
	if r := conditional.NilValFloat64(nil, 1.0, 2.0); r != 1.0 {
		t.Error("expected 1.0")
	}
	v := 5.0
	if r := conditional.NilValFloat64(&v, 1.0, 2.0); r != 2.0 {
		t.Error("expected 2.0")
	}
}

func Test_NilValPtrFloat64_Cov4(t *testing.T) {
	r := conditional.NilValPtrFloat64(nil, 1.0, 2.0)
	if *r != 1.0 {
		t.Error("expected 1.0")
	}
	v := 5.0
	r = conditional.NilValPtrFloat64(&v, 1.0, 2.0)
	if *r != 2.0 {
		t.Error("expected 2.0")
	}
}

func Test_IfSliceFloat64_Cov4(t *testing.T) {
	r := conditional.IfSliceFloat64(true, []float64{1.0}, []float64{2.0, 3.0})
	if len(r) != 1 {
		t.Error("expected 1")
	}
}

func Test_IfPtrFloat64_Cov4(t *testing.T) {
	a, b := 1.0, 2.0
	if *conditional.IfPtrFloat64(true, &a, &b) != 1.0 {
		t.Error("expected 1.0")
	}
}

func Test_IfFuncFloat64_Cov4(t *testing.T) {
	r := conditional.IfFuncFloat64(true, func() float64 { return 1.0 }, func() float64 { return 2.0 })
	if r != 1.0 {
		t.Error("expected 1.0")
	}
}

func Test_IfTrueFuncFloat64_Cov4(t *testing.T) {
	if r := conditional.IfTrueFuncFloat64(true, func() float64 { return 3.14 }); r != 3.14 {
		t.Error("expected 3.14")
	}
	if r := conditional.IfTrueFuncFloat64(false, func() float64 { return 3.14 }); r != 0 {
		t.Error("expected 0")
	}
}

// ============================================================================
// typed_int8.go — all 11 functions
// ============================================================================

func Test_IfInt8_Cov4(t *testing.T) {
	if r := conditional.IfInt8(true, 1, 2); r != 1 {
		t.Error("expected 1")
	}
	if r := conditional.IfInt8(false, 1, 2); r != 2 {
		t.Error("expected 2")
	}
}

func Test_IfFuncInt8_Cov4(t *testing.T) {
	r := conditional.IfFuncInt8(true, func() int8 { return 1 }, func() int8 { return 2 })
	if r != 1 {
		t.Error("expected 1")
	}
}

func Test_IfTrueFuncInt8_Cov4(t *testing.T) {
	if r := conditional.IfTrueFuncInt8(true, func() int8 { return 5 }); r != 5 {
		t.Error("expected 5")
	}
	if r := conditional.IfTrueFuncInt8(false, func() int8 { return 5 }); r != 0 {
		t.Error("expected 0")
	}
}

func Test_IfSliceInt8_Cov4(t *testing.T) {
	r := conditional.IfSliceInt8(true, []int8{1}, []int8{2, 3})
	if len(r) != 1 {
		t.Error("expected 1")
	}
}

func Test_IfPtrInt8_Cov4(t *testing.T) {
	a, b := int8(1), int8(2)
	if *conditional.IfPtrInt8(true, &a, &b) != 1 {
		t.Error("expected 1")
	}
}

func Test_NilDefInt8_Cov4(t *testing.T) {
	if r := conditional.NilDefInt8(nil, 5); r != 5 {
		t.Error("expected 5")
	}
	v := int8(3)
	if r := conditional.NilDefInt8(&v, 5); r != 3 {
		t.Error("expected 3")
	}
}

func Test_NilDefPtrInt8_Cov4(t *testing.T) {
	r := conditional.NilDefPtrInt8(nil, 9)
	if *r != 9 {
		t.Error("expected 9")
	}
}

func Test_ValueOrZeroInt8_Cov4(t *testing.T) {
	if r := conditional.ValueOrZeroInt8(nil); r != 0 {
		t.Error("expected 0")
	}
}

func Test_PtrOrZeroInt8_Cov4(t *testing.T) {
	r := conditional.PtrOrZeroInt8(nil)
	if *r != 0 {
		t.Error("expected 0")
	}
}

func Test_NilValInt8_Cov4(t *testing.T) {
	if r := conditional.NilValInt8(nil, 1, 2); r != 1 {
		t.Error("expected 1")
	}
}

func Test_NilValPtrInt8_Cov4(t *testing.T) {
	r := conditional.NilValPtrInt8(nil, 1, 2)
	if *r != 1 {
		t.Error("expected 1")
	}
}

// ============================================================================
// typed_int16.go — all 11 functions
// ============================================================================

func Test_IfInt16_Cov4(t *testing.T) {
	if r := conditional.IfInt16(true, 10, 20); r != 10 {
		t.Error("expected 10")
	}
}

func Test_IfFuncInt16_Cov4(t *testing.T) {
	r := conditional.IfFuncInt16(true, func() int16 { return 10 }, func() int16 { return 20 })
	if r != 10 {
		t.Error("expected 10")
	}
}

func Test_IfTrueFuncInt16_Cov4(t *testing.T) {
	if r := conditional.IfTrueFuncInt16(true, func() int16 { return 50 }); r != 50 {
		t.Error("expected 50")
	}
	if r := conditional.IfTrueFuncInt16(false, func() int16 { return 50 }); r != 0 {
		t.Error("expected 0")
	}
}

func Test_IfSliceInt16_Cov4(t *testing.T) {
	r := conditional.IfSliceInt16(true, []int16{1}, []int16{2, 3})
	if len(r) != 1 {
		t.Error("expected 1")
	}
}

func Test_IfPtrInt16_Cov4(t *testing.T) {
	a, b := int16(1), int16(2)
	if *conditional.IfPtrInt16(true, &a, &b) != 1 {
		t.Error("expected 1")
	}
}

func Test_NilDefInt16_Cov4(t *testing.T) {
	if r := conditional.NilDefInt16(nil, 5); r != 5 {
		t.Error("expected 5")
	}
}

func Test_NilDefPtrInt16_Cov4(t *testing.T) {
	r := conditional.NilDefPtrInt16(nil, 9)
	if *r != 9 {
		t.Error("expected 9")
	}
}

func Test_ValueOrZeroInt16_Cov4(t *testing.T) {
	if r := conditional.ValueOrZeroInt16(nil); r != 0 {
		t.Error("expected 0")
	}
}

func Test_PtrOrZeroInt16_Cov4(t *testing.T) {
	r := conditional.PtrOrZeroInt16(nil)
	if *r != 0 {
		t.Error("expected 0")
	}
}

func Test_NilValInt16_Cov4(t *testing.T) {
	if r := conditional.NilValInt16(nil, 1, 2); r != 1 {
		t.Error("expected 1")
	}
}

func Test_NilValPtrInt16_Cov4(t *testing.T) {
	r := conditional.NilValPtrInt16(nil, 1, 2)
	if *r != 1 {
		t.Error("expected 1")
	}
}

// ============================================================================
// typed_int32.go — all 11 functions
// ============================================================================

func Test_IfInt32_Cov4(t *testing.T) {
	if r := conditional.IfInt32(true, 10, 20); r != 10 {
		t.Error("expected 10")
	}
}

func Test_IfFuncInt32_Cov4(t *testing.T) {
	r := conditional.IfFuncInt32(true, func() int32 { return 10 }, func() int32 { return 20 })
	if r != 10 {
		t.Error("expected 10")
	}
}

func Test_IfTrueFuncInt32_Cov4(t *testing.T) {
	if r := conditional.IfTrueFuncInt32(true, func() int32 { return 50 }); r != 50 {
		t.Error("expected 50")
	}
	if r := conditional.IfTrueFuncInt32(false, func() int32 { return 50 }); r != 0 {
		t.Error("expected 0")
	}
}

func Test_IfSliceInt32_Cov4(t *testing.T) {
	r := conditional.IfSliceInt32(true, []int32{1}, []int32{2, 3})
	if len(r) != 1 {
		t.Error("expected 1")
	}
}

func Test_IfPtrInt32_Cov4(t *testing.T) {
	a, b := int32(1), int32(2)
	if *conditional.IfPtrInt32(true, &a, &b) != 1 {
		t.Error("expected 1")
	}
}

func Test_NilDefInt32_Cov4(t *testing.T) {
	if r := conditional.NilDefInt32(nil, 5); r != 5 {
		t.Error("expected 5")
	}
}

func Test_NilDefPtrInt32_Cov4(t *testing.T) {
	r := conditional.NilDefPtrInt32(nil, 9)
	if *r != 9 {
		t.Error("expected 9")
	}
}

func Test_ValueOrZeroInt32_Cov4(t *testing.T) {
	if r := conditional.ValueOrZeroInt32(nil); r != 0 {
		t.Error("expected 0")
	}
}

func Test_PtrOrZeroInt32_Cov4(t *testing.T) {
	r := conditional.PtrOrZeroInt32(nil)
	if *r != 0 {
		t.Error("expected 0")
	}
}

func Test_NilValInt32_Cov4(t *testing.T) {
	if r := conditional.NilValInt32(nil, 1, 2); r != 1 {
		t.Error("expected 1")
	}
}

func Test_NilValPtrInt32_Cov4(t *testing.T) {
	r := conditional.NilValPtrInt32(nil, 1, 2)
	if *r != 1 {
		t.Error("expected 1")
	}
}

// ============================================================================
// typed_int64.go — all 11 functions
// ============================================================================

func Test_IfInt64_Cov4(t *testing.T) {
	if r := conditional.IfInt64(true, 10, 20); r != 10 {
		t.Error("expected 10")
	}
}

func Test_IfFuncInt64_Cov4(t *testing.T) {
	r := conditional.IfFuncInt64(true, func() int64 { return 10 }, func() int64 { return 20 })
	if r != 10 {
		t.Error("expected 10")
	}
}

func Test_IfTrueFuncInt64_Cov4(t *testing.T) {
	if r := conditional.IfTrueFuncInt64(true, func() int64 { return 50 }); r != 50 {
		t.Error("expected 50")
	}
	if r := conditional.IfTrueFuncInt64(false, func() int64 { return 50 }); r != 0 {
		t.Error("expected 0")
	}
}

func Test_IfSliceInt64_Cov4(t *testing.T) {
	r := conditional.IfSliceInt64(true, []int64{1}, []int64{2, 3})
	if len(r) != 1 {
		t.Error("expected 1")
	}
}

func Test_IfPtrInt64_Cov4(t *testing.T) {
	a, b := int64(1), int64(2)
	if *conditional.IfPtrInt64(true, &a, &b) != 1 {
		t.Error("expected 1")
	}
}

func Test_NilDefInt64_Cov4(t *testing.T) {
	if r := conditional.NilDefInt64(nil, 5); r != 5 {
		t.Error("expected 5")
	}
}

func Test_NilDefPtrInt64_Cov4(t *testing.T) {
	r := conditional.NilDefPtrInt64(nil, 9)
	if *r != 9 {
		t.Error("expected 9")
	}
}

func Test_ValueOrZeroInt64_Cov4(t *testing.T) {
	if r := conditional.ValueOrZeroInt64(nil); r != 0 {
		t.Error("expected 0")
	}
}

func Test_PtrOrZeroInt64_Cov4(t *testing.T) {
	r := conditional.PtrOrZeroInt64(nil)
	if *r != 0 {
		t.Error("expected 0")
	}
}

func Test_NilValInt64_Cov4(t *testing.T) {
	if r := conditional.NilValInt64(nil, 1, 2); r != 1 {
		t.Error("expected 1")
	}
}

func Test_NilValPtrInt64_Cov4(t *testing.T) {
	r := conditional.NilValPtrInt64(nil, 1, 2)
	if *r != 1 {
		t.Error("expected 1")
	}
}

// ============================================================================
// typed_uint.go — all 11 functions
// ============================================================================

func Test_IfUint_Cov4(t *testing.T) {
	if r := conditional.IfUint(true, 10, 20); r != 10 {
		t.Error("expected 10")
	}
}

func Test_IfFuncUint_Cov4(t *testing.T) {
	r := conditional.IfFuncUint(true, func() uint { return 10 }, func() uint { return 20 })
	if r != 10 {
		t.Error("expected 10")
	}
}

func Test_IfTrueFuncUint_Cov4(t *testing.T) {
	if r := conditional.IfTrueFuncUint(true, func() uint { return 50 }); r != 50 {
		t.Error("expected 50")
	}
	if r := conditional.IfTrueFuncUint(false, func() uint { return 50 }); r != 0 {
		t.Error("expected 0")
	}
}

func Test_IfSliceUint_Cov4(t *testing.T) {
	r := conditional.IfSliceUint(true, []uint{1}, []uint{2, 3})
	if len(r) != 1 {
		t.Error("expected 1")
	}
}

func Test_IfPtrUint_Cov4(t *testing.T) {
	a, b := uint(1), uint(2)
	if *conditional.IfPtrUint(true, &a, &b) != 1 {
		t.Error("expected 1")
	}
}

func Test_NilDefUint_Cov4(t *testing.T) {
	if r := conditional.NilDefUint(nil, 5); r != 5 {
		t.Error("expected 5")
	}
}

func Test_NilDefPtrUint_Cov4(t *testing.T) {
	r := conditional.NilDefPtrUint(nil, 9)
	if *r != 9 {
		t.Error("expected 9")
	}
}

func Test_ValueOrZeroUint_Cov4(t *testing.T) {
	if r := conditional.ValueOrZeroUint(nil); r != 0 {
		t.Error("expected 0")
	}
}

func Test_PtrOrZeroUint_Cov4(t *testing.T) {
	r := conditional.PtrOrZeroUint(nil)
	if *r != 0 {
		t.Error("expected 0")
	}
}

func Test_NilValUint_Cov4(t *testing.T) {
	if r := conditional.NilValUint(nil, 1, 2); r != 1 {
		t.Error("expected 1")
	}
}

func Test_NilValPtrUint_Cov4(t *testing.T) {
	r := conditional.NilValPtrUint(nil, 1, 2)
	if *r != 1 {
		t.Error("expected 1")
	}
}

// ============================================================================
// typed_uint8.go — all 11 functions
// ============================================================================

func Test_IfUint8_Cov4(t *testing.T) {
	if r := conditional.IfUint8(true, 10, 20); r != 10 {
		t.Error("expected 10")
	}
}

func Test_IfFuncUint8_Cov4(t *testing.T) {
	r := conditional.IfFuncUint8(true, func() uint8 { return 10 }, func() uint8 { return 20 })
	if r != 10 {
		t.Error("expected 10")
	}
}

func Test_IfTrueFuncUint8_Cov4(t *testing.T) {
	if r := conditional.IfTrueFuncUint8(true, func() uint8 { return 50 }); r != 50 {
		t.Error("expected 50")
	}
	if r := conditional.IfTrueFuncUint8(false, func() uint8 { return 50 }); r != 0 {
		t.Error("expected 0")
	}
}

func Test_IfSliceUint8_Cov4(t *testing.T) {
	r := conditional.IfSliceUint8(true, []uint8{1}, []uint8{2, 3})
	if len(r) != 1 {
		t.Error("expected 1")
	}
}

func Test_IfPtrUint8_Cov4(t *testing.T) {
	a, b := uint8(1), uint8(2)
	if *conditional.IfPtrUint8(true, &a, &b) != 1 {
		t.Error("expected 1")
	}
}

func Test_NilDefUint8_Cov4(t *testing.T) {
	if r := conditional.NilDefUint8(nil, 5); r != 5 {
		t.Error("expected 5")
	}
}

func Test_NilDefPtrUint8_Cov4(t *testing.T) {
	r := conditional.NilDefPtrUint8(nil, 9)
	if *r != 9 {
		t.Error("expected 9")
	}
}

func Test_ValueOrZeroUint8_Cov4(t *testing.T) {
	if r := conditional.ValueOrZeroUint8(nil); r != 0 {
		t.Error("expected 0")
	}
}

func Test_PtrOrZeroUint8_Cov4(t *testing.T) {
	r := conditional.PtrOrZeroUint8(nil)
	if *r != 0 {
		t.Error("expected 0")
	}
}

func Test_NilValUint8_Cov4(t *testing.T) {
	if r := conditional.NilValUint8(nil, 1, 2); r != 1 {
		t.Error("expected 1")
	}
}

func Test_NilValPtrUint8_Cov4(t *testing.T) {
	r := conditional.NilValPtrUint8(nil, 1, 2)
	if *r != 1 {
		t.Error("expected 1")
	}
}

// ============================================================================
// typed_uint16.go — all 11 functions
// ============================================================================

func Test_IfUint16_Cov4(t *testing.T) {
	if r := conditional.IfUint16(true, 10, 20); r != 10 {
		t.Error("expected 10")
	}
}

func Test_IfFuncUint16_Cov4(t *testing.T) {
	r := conditional.IfFuncUint16(true, func() uint16 { return 10 }, func() uint16 { return 20 })
	if r != 10 {
		t.Error("expected 10")
	}
}

func Test_IfTrueFuncUint16_Cov4(t *testing.T) {
	if r := conditional.IfTrueFuncUint16(true, func() uint16 { return 50 }); r != 50 {
		t.Error("expected 50")
	}
}

func Test_IfSliceUint16_Cov4(t *testing.T) {
	r := conditional.IfSliceUint16(true, []uint16{1}, []uint16{2, 3})
	if len(r) != 1 {
		t.Error("expected 1")
	}
}

func Test_IfPtrUint16_Cov4(t *testing.T) {
	a, b := uint16(1), uint16(2)
	if *conditional.IfPtrUint16(true, &a, &b) != 1 {
		t.Error("expected 1")
	}
}

func Test_NilDefUint16_Cov4(t *testing.T) {
	if r := conditional.NilDefUint16(nil, 5); r != 5 {
		t.Error("expected 5")
	}
}

func Test_NilDefPtrUint16_Cov4(t *testing.T) {
	r := conditional.NilDefPtrUint16(nil, 9)
	if *r != 9 {
		t.Error("expected 9")
	}
}

func Test_ValueOrZeroUint16_Cov4(t *testing.T) {
	if r := conditional.ValueOrZeroUint16(nil); r != 0 {
		t.Error("expected 0")
	}
}

func Test_PtrOrZeroUint16_Cov4(t *testing.T) {
	r := conditional.PtrOrZeroUint16(nil)
	if *r != 0 {
		t.Error("expected 0")
	}
}

func Test_NilValUint16_Cov4(t *testing.T) {
	if r := conditional.NilValUint16(nil, 1, 2); r != 1 {
		t.Error("expected 1")
	}
}

func Test_NilValPtrUint16_Cov4(t *testing.T) {
	r := conditional.NilValPtrUint16(nil, 1, 2)
	if *r != 1 {
		t.Error("expected 1")
	}
}

// ============================================================================
// typed_uint32.go — all 11 functions
// ============================================================================

func Test_IfUint32_Cov4(t *testing.T) {
	if r := conditional.IfUint32(true, 10, 20); r != 10 {
		t.Error("expected 10")
	}
}

func Test_IfFuncUint32_Cov4(t *testing.T) {
	r := conditional.IfFuncUint32(true, func() uint32 { return 10 }, func() uint32 { return 20 })
	if r != 10 {
		t.Error("expected 10")
	}
}

func Test_IfTrueFuncUint32_Cov4(t *testing.T) {
	if r := conditional.IfTrueFuncUint32(true, func() uint32 { return 50 }); r != 50 {
		t.Error("expected 50")
	}
}

func Test_IfSliceUint32_Cov4(t *testing.T) {
	r := conditional.IfSliceUint32(true, []uint32{1}, []uint32{2, 3})
	if len(r) != 1 {
		t.Error("expected 1")
	}
}

func Test_IfPtrUint32_Cov4(t *testing.T) {
	a, b := uint32(1), uint32(2)
	if *conditional.IfPtrUint32(true, &a, &b) != 1 {
		t.Error("expected 1")
	}
}

func Test_NilDefUint32_Cov4(t *testing.T) {
	if r := conditional.NilDefUint32(nil, 5); r != 5 {
		t.Error("expected 5")
	}
}

func Test_NilDefPtrUint32_Cov4(t *testing.T) {
	r := conditional.NilDefPtrUint32(nil, 9)
	if *r != 9 {
		t.Error("expected 9")
	}
}

func Test_ValueOrZeroUint32_Cov4(t *testing.T) {
	if r := conditional.ValueOrZeroUint32(nil); r != 0 {
		t.Error("expected 0")
	}
}

func Test_PtrOrZeroUint32_Cov4(t *testing.T) {
	r := conditional.PtrOrZeroUint32(nil)
	if *r != 0 {
		t.Error("expected 0")
	}
}

func Test_NilValUint32_Cov4(t *testing.T) {
	if r := conditional.NilValUint32(nil, 1, 2); r != 1 {
		t.Error("expected 1")
	}
}

func Test_NilValPtrUint32_Cov4(t *testing.T) {
	r := conditional.NilValPtrUint32(nil, 1, 2)
	if *r != 1 {
		t.Error("expected 1")
	}
}

// ============================================================================
// typed_uint64.go — all 11 functions
// ============================================================================

func Test_IfUint64_Cov4(t *testing.T) {
	if r := conditional.IfUint64(true, 10, 20); r != 10 {
		t.Error("expected 10")
	}
}

func Test_IfFuncUint64_Cov4(t *testing.T) {
	r := conditional.IfFuncUint64(true, func() uint64 { return 10 }, func() uint64 { return 20 })
	if r != 10 {
		t.Error("expected 10")
	}
}

func Test_IfTrueFuncUint64_Cov4(t *testing.T) {
	if r := conditional.IfTrueFuncUint64(true, func() uint64 { return 50 }); r != 50 {
		t.Error("expected 50")
	}
}

func Test_IfSliceUint64_Cov4(t *testing.T) {
	r := conditional.IfSliceUint64(true, []uint64{1}, []uint64{2, 3})
	if len(r) != 1 {
		t.Error("expected 1")
	}
}

func Test_IfPtrUint64_Cov4(t *testing.T) {
	a, b := uint64(1), uint64(2)
	if *conditional.IfPtrUint64(true, &a, &b) != 1 {
		t.Error("expected 1")
	}
}

func Test_NilDefUint64_Cov4(t *testing.T) {
	if r := conditional.NilDefUint64(nil, 5); r != 5 {
		t.Error("expected 5")
	}
}

func Test_NilDefPtrUint64_Cov4(t *testing.T) {
	r := conditional.NilDefPtrUint64(nil, 9)
	if *r != 9 {
		t.Error("expected 9")
	}
}

func Test_ValueOrZeroUint64_Cov4(t *testing.T) {
	if r := conditional.ValueOrZeroUint64(nil); r != 0 {
		t.Error("expected 0")
	}
}

func Test_PtrOrZeroUint64_Cov4(t *testing.T) {
	r := conditional.PtrOrZeroUint64(nil)
	if *r != 0 {
		t.Error("expected 0")
	}
}

func Test_NilValUint64_Cov4(t *testing.T) {
	if r := conditional.NilValUint64(nil, 1, 2); r != 1 {
		t.Error("expected 1")
	}
}

func Test_NilValPtrUint64_Cov4(t *testing.T) {
	r := conditional.NilValPtrUint64(nil, 1, 2)
	if *r != 1 {
		t.Error("expected 1")
	}
}

// ============================================================================
// NilDefPtrBool, PtrOrZeroBool, NilValPtrBool (may be uncovered)
// ============================================================================

func Test_NilDefPtrBool_Cov4(t *testing.T) {
	r := conditional.NilDefPtrBool(nil, true)
	if *r != true {
		t.Error("expected true")
	}
	v := false
	r = conditional.NilDefPtrBool(&v, true)
	if *r != false {
		t.Error("expected false")
	}
}

func Test_PtrOrZeroBool_Cov4(t *testing.T) {
	r := conditional.PtrOrZeroBool(nil)
	if *r != false {
		t.Error("expected false")
	}
}

func Test_NilValPtrBool_Cov4(t *testing.T) {
	r := conditional.NilValPtrBool(nil, true, false)
	if *r != true {
		t.Error("expected true")
	}
}

// NilDefPtrByte, PtrOrZeroByte, NilValPtrByte
func Test_NilDefPtrByte_Cov4(t *testing.T) {
	r := conditional.NilDefPtrByte(nil, 9)
	if *r != 9 {
		t.Error("expected 9")
	}
}

func Test_PtrOrZeroByte_Cov4(t *testing.T) {
	r := conditional.PtrOrZeroByte(nil)
	if *r != 0 {
		t.Error("expected 0")
	}
}

func Test_NilValPtrByte_Cov4(t *testing.T) {
	r := conditional.NilValPtrByte(nil, 1, 2)
	if *r != 1 {
		t.Error("expected 1")
	}
}

// NilDefPtrInt, PtrOrZeroInt, NilValPtrInt
func Test_NilDefPtrInt_Cov4(t *testing.T) {
	r := conditional.NilDefPtrInt(nil, 42)
	if *r != 42 {
		t.Error("expected 42")
	}
}

func Test_PtrOrZeroInt_Cov4(t *testing.T) {
	r := conditional.PtrOrZeroInt(nil)
	if *r != 0 {
		t.Error("expected 0")
	}
}

func Test_NilValPtrInt_Cov4(t *testing.T) {
	r := conditional.NilValPtrInt(nil, -1, 1)
	if *r != -1 {
		t.Error("expected -1")
	}
}

// ============================================================================
// VoidFunctions with nil func in slice
// ============================================================================

func Test_VoidFunctions_WithNilFunc_Cov4(t *testing.T) {
	count := 0
	funcs := []func(){
		nil,
		func() { count++ },
	}
	conditional.VoidFunctions(true, funcs, []func(){})
	if count != 1 {
		t.Errorf("expected 1, got %d", count)
	}
}

// ============================================================================
// ErrorFunctionsExecuteResults with nil func in slice
// ============================================================================

func Test_ErrorFunctionsExecuteResults_WithNilFunc_Cov4(t *testing.T) {
	funcs := []func() error{
		nil,
		func() error { return nil },
	}
	err := conditional.ErrorFunctionsExecuteResults(true, funcs, []func() error{})
	if err != nil {
		t.Error("expected nil error")
	}
}

// ============================================================================
// FunctionsExecuteResults with nil and skip-take
// ============================================================================

func Test_FunctionsExecuteResults_SkipTake_Cov4(t *testing.T) {
	funcs := []func() (string, bool, bool){
		func() (string, bool, bool) { return "skip", false, false },
		func() (string, bool, bool) { return "take", true, false },
	}
	results := conditional.FunctionsExecuteResults[string](true, funcs, nil)
	if len(results) != 1 || results[0] != "take" {
		t.Error("expected only 'take'")
	}
}

func Test_FunctionsExecuteResults_NilFunc_Cov4(t *testing.T) {
	funcs := []func() (string, bool, bool){
		nil,
		func() (string, bool, bool) { return "a", true, false },
	}
	results := conditional.FunctionsExecuteResults[string](true, funcs, nil)
	if len(results) != 1 {
		t.Error("expected 1")
	}
}

func Test_FunctionsExecuteResults_Empty_Cov4(t *testing.T) {
	results := conditional.FunctionsExecuteResults[string](true, nil, nil)
	if results != nil {
		t.Error("expected nil")
	}
}
