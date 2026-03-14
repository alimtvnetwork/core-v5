package coremathtests

import (
	"math"
	"testing"

	"github.com/alimtvnetwork/core/coremath"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── MaxByte ──

func Test_Cov_MaxByte_LeftGreater(t *testing.T) {
	actual := args.Map{"result": coremath.MaxByte(10, 5)}
	expected := args.Map{"result": byte(10)}
	expected.ShouldBeEqual(t, 0, "MaxByte left greater", actual)
}

func Test_Cov_MaxByte_RightGreater(t *testing.T) {
	actual := args.Map{"result": coremath.MaxByte(3, 8)}
	expected := args.Map{"result": byte(8)}
	expected.ShouldBeEqual(t, 0, "MaxByte right greater", actual)
}

// ── MinByte ──

func Test_Cov_MinByte_LeftSmaller(t *testing.T) {
	actual := args.Map{"result": coremath.MinByte(2, 9)}
	expected := args.Map{"result": byte(2)}
	expected.ShouldBeEqual(t, 0, "MinByte left smaller", actual)
}

func Test_Cov_MinByte_RightSmaller(t *testing.T) {
	actual := args.Map{"result": coremath.MinByte(9, 3)}
	expected := args.Map{"result": byte(3)}
	expected.ShouldBeEqual(t, 0, "MinByte right smaller", actual)
}

// ── MaxFloat32 ──

func Test_Cov_MaxFloat32_LeftGreater(t *testing.T) {
	actual := args.Map{"result": coremath.MaxFloat32(3.5, 1.2)}
	expected := args.Map{"result": float32(3.5)}
	expected.ShouldBeEqual(t, 0, "MaxFloat32 left greater", actual)
}

func Test_Cov_MaxFloat32_RightGreater(t *testing.T) {
	actual := args.Map{"result": coremath.MaxFloat32(1.0, 2.5)}
	expected := args.Map{"result": float32(2.5)}
	expected.ShouldBeEqual(t, 0, "MaxFloat32 right greater", actual)
}

// ── MinFloat32 ──

func Test_Cov_MinFloat32_LeftSmaller(t *testing.T) {
	actual := args.Map{"result": coremath.MinFloat32(1.0, 5.0)}
	expected := args.Map{"result": float32(1.0)}
	expected.ShouldBeEqual(t, 0, "MinFloat32 left smaller", actual)
}

func Test_Cov_MinFloat32_RightSmaller(t *testing.T) {
	actual := args.Map{"result": coremath.MinFloat32(5.0, 2.0)}
	expected := args.Map{"result": float32(2.0)}
	expected.ShouldBeEqual(t, 0, "MinFloat32 right smaller", actual)
}

// ── MaxInt ──

func Test_Cov_MaxInt_LeftGreater(t *testing.T) {
	actual := args.Map{"result": coremath.MaxInt(10, 5)}
	expected := args.Map{"result": 10}
	expected.ShouldBeEqual(t, 0, "MaxInt left greater", actual)
}

func Test_Cov_MaxInt_RightGreater(t *testing.T) {
	actual := args.Map{"result": coremath.MaxInt(3, 8)}
	expected := args.Map{"result": 8}
	expected.ShouldBeEqual(t, 0, "MaxInt right greater", actual)
}

// ── MinInt ──

func Test_Cov_MinInt_LeftSmaller(t *testing.T) {
	actual := args.Map{"result": coremath.MinInt(2, 9)}
	expected := args.Map{"result": 2}
	expected.ShouldBeEqual(t, 0, "MinInt left smaller", actual)
}

func Test_Cov_MinInt_RightSmaller(t *testing.T) {
	actual := args.Map{"result": coremath.MinInt(9, 3)}
	expected := args.Map{"result": 3}
	expected.ShouldBeEqual(t, 0, "MinInt right smaller", actual)
}

// ── IsRangeWithin.Integer ──

func Test_Cov_IntegerWithin_ToByte(t *testing.T) {
	actual := args.Map{
		"inRange":    coremath.IsRangeWithin.Integer.ToByte(100),
		"outOfRange": coremath.IsRangeWithin.Integer.ToByte(300),
		"negative":   coremath.IsRangeWithin.Integer.ToByte(-1),
	}
	expected := args.Map{"inRange": true, "outOfRange": false, "negative": false}
	expected.ShouldBeEqual(t, 0, "IntegerWithin ToByte", actual)
}

func Test_Cov_IntegerWithin_ToUnsignedInt16(t *testing.T) {
	actual := args.Map{
		"inRange":  coremath.IsRangeWithin.Integer.ToUnsignedInt16(100),
		"outRange": coremath.IsRangeWithin.Integer.ToUnsignedInt16(70000),
	}
	expected := args.Map{"inRange": true, "outRange": false}
	expected.ShouldBeEqual(t, 0, "IntegerWithin ToUnsignedInt16", actual)
}

func Test_Cov_IntegerWithin_ToUnsignedInt32(t *testing.T) {
	actual := args.Map{
		"inRange":  coremath.IsRangeWithin.Integer.ToUnsignedInt32(100),
		"negative": coremath.IsRangeWithin.Integer.ToUnsignedInt32(-1),
	}
	expected := args.Map{"inRange": true, "negative": false}
	expected.ShouldBeEqual(t, 0, "IntegerWithin ToUnsignedInt32", actual)
}

func Test_Cov_IntegerWithin_ToUnsignedInt64(t *testing.T) {
	actual := args.Map{
		"positive": coremath.IsRangeWithin.Integer.ToUnsignedInt64(100),
		"negative": coremath.IsRangeWithin.Integer.ToUnsignedInt64(-1),
	}
	expected := args.Map{"positive": true, "negative": false}
	expected.ShouldBeEqual(t, 0, "IntegerWithin ToUnsignedInt64", actual)
}

func Test_Cov_IntegerWithin_ToInt8(t *testing.T) {
	actual := args.Map{
		"inRange":  coremath.IsRangeWithin.Integer.ToInt8(50),
		"outRange": coremath.IsRangeWithin.Integer.ToInt8(200),
	}
	expected := args.Map{"inRange": true, "outRange": false}
	expected.ShouldBeEqual(t, 0, "IntegerWithin ToInt8", actual)
}

func Test_Cov_IntegerWithin_ToInt16(t *testing.T) {
	actual := args.Map{
		"inRange":  coremath.IsRangeWithin.Integer.ToInt16(1000),
		"outRange": coremath.IsRangeWithin.Integer.ToInt16(40000),
	}
	expected := args.Map{"inRange": true, "outRange": false}
	expected.ShouldBeEqual(t, 0, "IntegerWithin ToInt16", actual)
}

func Test_Cov_IntegerWithin_ToInt32(t *testing.T) {
	actual := args.Map{
		"inRange":  coremath.IsRangeWithin.Integer.ToInt32(1000),
		"outRange": coremath.IsRangeWithin.Integer.ToInt32(math.MaxInt32 + 1),
	}
	expected := args.Map{"inRange": true, "outRange": false}
	expected.ShouldBeEqual(t, 0, "IntegerWithin ToInt32", actual)
}

// ── IsRangeWithin.Integer16 ──

func Test_Cov_Integer16Within_ToByte(t *testing.T) {
	actual := args.Map{
		"inRange":  coremath.IsRangeWithin.Integer16.ToByte(100),
		"outRange": coremath.IsRangeWithin.Integer16.ToByte(300),
		"negative": coremath.IsRangeWithin.Integer16.ToByte(-1),
	}
	expected := args.Map{"inRange": true, "outRange": false, "negative": false}
	expected.ShouldBeEqual(t, 0, "Integer16Within ToByte", actual)
}

func Test_Cov_Integer16Within_ToUnsignedInt16(t *testing.T) {
	actual := args.Map{
		"positive": coremath.IsRangeWithin.Integer16.ToUnsignedInt16(100),
		"negative": coremath.IsRangeWithin.Integer16.ToUnsignedInt16(-1),
	}
	expected := args.Map{"positive": true, "negative": false}
	expected.ShouldBeEqual(t, 0, "Integer16Within ToUnsignedInt16", actual)
}

func Test_Cov_Integer16Within_ToUnsignedInt32(t *testing.T) {
	actual := args.Map{
		"positive": coremath.IsRangeWithin.Integer16.ToUnsignedInt32(100),
		"negative": coremath.IsRangeWithin.Integer16.ToUnsignedInt32(-1),
	}
	expected := args.Map{"positive": true, "negative": false}
	expected.ShouldBeEqual(t, 0, "Integer16Within ToUnsignedInt32", actual)
}

func Test_Cov_Integer16Within_ToUnsignedInt64(t *testing.T) {
	actual := args.Map{
		"positive": coremath.IsRangeWithin.Integer16.ToUnsignedInt64(100),
		"negative": coremath.IsRangeWithin.Integer16.ToUnsignedInt64(-1),
	}
	expected := args.Map{"positive": true, "negative": false}
	expected.ShouldBeEqual(t, 0, "Integer16Within ToUnsignedInt64", actual)
}

func Test_Cov_Integer16Within_ToInt8(t *testing.T) {
	actual := args.Map{
		"inRange":  coremath.IsRangeWithin.Integer16.ToInt8(50),
		"outRange": coremath.IsRangeWithin.Integer16.ToInt8(200),
	}
	expected := args.Map{"inRange": true, "outRange": false}
	expected.ShouldBeEqual(t, 0, "Integer16Within ToInt8", actual)
}

// ── IsRangeWithin.Integer32 ──

func Test_Cov_Integer32Within_ToByte(t *testing.T) {
	actual := args.Map{
		"inRange":  coremath.IsRangeWithin.Integer32.ToByte(200),
		"outRange": coremath.IsRangeWithin.Integer32.ToByte(300),
	}
	expected := args.Map{"inRange": true, "outRange": false}
	expected.ShouldBeEqual(t, 0, "Integer32Within ToByte", actual)
}

func Test_Cov_Integer32Within_ToUnsignedInt16(t *testing.T) {
	actual := args.Map{
		"inRange":  coremath.IsRangeWithin.Integer32.ToUnsignedInt16(1000),
		"outRange": coremath.IsRangeWithin.Integer32.ToUnsignedInt16(70000),
	}
	expected := args.Map{"inRange": true, "outRange": false}
	expected.ShouldBeEqual(t, 0, "Integer32Within ToUnsignedInt16", actual)
}

func Test_Cov_Integer32Within_ToUnsignedInt32(t *testing.T) {
	actual := args.Map{
		"positive": coremath.IsRangeWithin.Integer32.ToUnsignedInt32(100),
		"negative": coremath.IsRangeWithin.Integer32.ToUnsignedInt32(-1),
	}
	expected := args.Map{"positive": true, "negative": false}
	expected.ShouldBeEqual(t, 0, "Integer32Within ToUnsignedInt32", actual)
}

func Test_Cov_Integer32Within_ToUnsignedInt64(t *testing.T) {
	actual := args.Map{
		"positive": coremath.IsRangeWithin.Integer32.ToUnsignedInt64(100),
		"negative": coremath.IsRangeWithin.Integer32.ToUnsignedInt64(-1),
	}
	expected := args.Map{"positive": true, "negative": false}
	expected.ShouldBeEqual(t, 0, "Integer32Within ToUnsignedInt64", actual)
}

func Test_Cov_Integer32Within_ToInt8(t *testing.T) {
	actual := args.Map{
		"inRange":  coremath.IsRangeWithin.Integer32.ToInt8(50),
		"outRange": coremath.IsRangeWithin.Integer32.ToInt8(200),
	}
	expected := args.Map{"inRange": true, "outRange": false}
	expected.ShouldBeEqual(t, 0, "Integer32Within ToInt8", actual)
}

func Test_Cov_Integer32Within_ToInt16(t *testing.T) {
	actual := args.Map{
		"inRange":  coremath.IsRangeWithin.Integer32.ToInt16(1000),
		"outRange": coremath.IsRangeWithin.Integer32.ToInt16(40000),
	}
	expected := args.Map{"inRange": true, "outRange": false}
	expected.ShouldBeEqual(t, 0, "Integer32Within ToInt16", actual)
}

func Test_Cov_Integer32Within_ToInt(t *testing.T) {
	actual := args.Map{
		"inRange": coremath.IsRangeWithin.Integer32.ToInt(1000),
	}
	expected := args.Map{"inRange": true}
	expected.ShouldBeEqual(t, 0, "Integer32Within ToInt", actual)
}

// ── IsRangeWithin.Integer64 ──

func Test_Cov_Integer64Within_ToByte(t *testing.T) {
	actual := args.Map{
		"inRange":  coremath.IsRangeWithin.Integer64.ToByte(200),
		"outRange": coremath.IsRangeWithin.Integer64.ToByte(300),
	}
	expected := args.Map{"inRange": true, "outRange": false}
	expected.ShouldBeEqual(t, 0, "Integer64Within ToByte", actual)
}

func Test_Cov_Integer64Within_ToUnsignedInt16(t *testing.T) {
	actual := args.Map{
		"inRange":  coremath.IsRangeWithin.Integer64.ToUnsignedInt16(100),
		"outRange": coremath.IsRangeWithin.Integer64.ToUnsignedInt16(70000),
	}
	expected := args.Map{"inRange": true, "outRange": false}
	expected.ShouldBeEqual(t, 0, "Integer64Within ToUnsignedInt16", actual)
}

func Test_Cov_Integer64Within_ToUnsignedInt32(t *testing.T) {
	actual := args.Map{
		"inRange":  coremath.IsRangeWithin.Integer64.ToUnsignedInt32(100),
		"outRange": coremath.IsRangeWithin.Integer64.ToUnsignedInt32(-1),
	}
	expected := args.Map{"inRange": true, "outRange": false}
	expected.ShouldBeEqual(t, 0, "Integer64Within ToUnsignedInt32", actual)
}

func Test_Cov_Integer64Within_ToUnsignedInt64(t *testing.T) {
	actual := args.Map{
		"positive": coremath.IsRangeWithin.Integer64.ToUnsignedInt64(100),
		"negative": coremath.IsRangeWithin.Integer64.ToUnsignedInt64(-1),
	}
	expected := args.Map{"positive": true, "negative": false}
	expected.ShouldBeEqual(t, 0, "Integer64Within ToUnsignedInt64", actual)
}

func Test_Cov_Integer64Within_ToInt8(t *testing.T) {
	actual := args.Map{
		"inRange":  coremath.IsRangeWithin.Integer64.ToInt8(50),
		"outRange": coremath.IsRangeWithin.Integer64.ToInt8(200),
	}
	expected := args.Map{"inRange": true, "outRange": false}
	expected.ShouldBeEqual(t, 0, "Integer64Within ToInt8", actual)
}

func Test_Cov_Integer64Within_ToInt16(t *testing.T) {
	actual := args.Map{
		"inRange":  coremath.IsRangeWithin.Integer64.ToInt16(1000),
		"outRange": coremath.IsRangeWithin.Integer64.ToInt16(40000),
	}
	expected := args.Map{"inRange": true, "outRange": false}
	expected.ShouldBeEqual(t, 0, "Integer64Within ToInt16", actual)
}

func Test_Cov_Integer64Within_ToInt32(t *testing.T) {
	actual := args.Map{
		"inRange":  coremath.IsRangeWithin.Integer64.ToInt32(1000),
		"outRange": coremath.IsRangeWithin.Integer64.ToInt32(int64(math.MaxInt32) + 1),
	}
	expected := args.Map{"inRange": true, "outRange": false}
	expected.ShouldBeEqual(t, 0, "Integer64Within ToInt32", actual)
}

func Test_Cov_Integer64Within_ToInt(t *testing.T) {
	actual := args.Map{
		"inRange": coremath.IsRangeWithin.Integer64.ToInt(1000),
	}
	expected := args.Map{"inRange": true}
	expected.ShouldBeEqual(t, 0, "Integer64Within ToInt", actual)
}

// ── IsRangeWithin.UnsignedInteger16 ──

func Test_Cov_UnsignedInt16Within_ToByte(t *testing.T) {
	actual := args.Map{
		"inRange":  coremath.IsRangeWithin.UnsignedInteger16.ToByte(200),
		"outRange": coremath.IsRangeWithin.UnsignedInteger16.ToByte(300),
	}
	expected := args.Map{"inRange": true, "outRange": false}
	expected.ShouldBeEqual(t, 0, "UnsignedInt16Within ToByte", actual)
}

func Test_Cov_UnsignedInt16Within_ToInt8(t *testing.T) {
	actual := args.Map{
		"inRange":  coremath.IsRangeWithin.UnsignedInteger16.ToInt8(50),
		"outRange": coremath.IsRangeWithin.UnsignedInteger16.ToInt8(200),
	}
	expected := args.Map{"inRange": true, "outRange": false}
	expected.ShouldBeEqual(t, 0, "UnsignedInt16Within ToInt8", actual)
}

// ── IsOutOfRange.Integer ──

func Test_Cov_IntegerOutOfRange_ToByte(t *testing.T) {
	actual := args.Map{
		"inRange":  coremath.IsOutOfRange.Integer.ToByte(100),
		"outRange": coremath.IsOutOfRange.Integer.ToByte(300),
	}
	expected := args.Map{"inRange": false, "outRange": true}
	expected.ShouldBeEqual(t, 0, "IntegerOutOfRange ToByte", actual)
}

func Test_Cov_IntegerOutOfRange_ToUnsignedInt16(t *testing.T) {
	actual := args.Map{
		"inRange":  coremath.IsOutOfRange.Integer.ToUnsignedInt16(100),
		"outRange": coremath.IsOutOfRange.Integer.ToUnsignedInt16(70000),
	}
	expected := args.Map{"inRange": false, "outRange": true}
	expected.ShouldBeEqual(t, 0, "IntegerOutOfRange ToUnsignedInt16", actual)
}

func Test_Cov_IntegerOutOfRange_ToUnsignedInt32(t *testing.T) {
	actual := args.Map{
		"inRange": coremath.IsOutOfRange.Integer.ToUnsignedInt32(100),
	}
	expected := args.Map{"inRange": false}
	expected.ShouldBeEqual(t, 0, "IntegerOutOfRange ToUnsignedInt32", actual)
}

func Test_Cov_IntegerOutOfRange_ToUnsignedInt64(t *testing.T) {
	actual := args.Map{
		"positive": coremath.IsOutOfRange.Integer.ToUnsignedInt64(100),
		"negative": coremath.IsOutOfRange.Integer.ToUnsignedInt64(-1),
	}
	expected := args.Map{"positive": false, "negative": true}
	expected.ShouldBeEqual(t, 0, "IntegerOutOfRange ToUnsignedInt64", actual)
}

func Test_Cov_IntegerOutOfRange_ToInt8(t *testing.T) {
	actual := args.Map{
		"inRange":  coremath.IsOutOfRange.Integer.ToInt8(50),
		"outRange": coremath.IsOutOfRange.Integer.ToInt8(200),
	}
	expected := args.Map{"inRange": false, "outRange": true}
	expected.ShouldBeEqual(t, 0, "IntegerOutOfRange ToInt8", actual)
}

func Test_Cov_IntegerOutOfRange_ToInt16(t *testing.T) {
	actual := args.Map{
		"inRange":  coremath.IsOutOfRange.Integer.ToInt16(1000),
		"outRange": coremath.IsOutOfRange.Integer.ToInt16(40000),
	}
	expected := args.Map{"inRange": false, "outRange": true}
	expected.ShouldBeEqual(t, 0, "IntegerOutOfRange ToInt16", actual)
}

func Test_Cov_IntegerOutOfRange_ToInt32(t *testing.T) {
	actual := args.Map{
		"inRange": coremath.IsOutOfRange.Integer.ToInt32(1000),
	}
	expected := args.Map{"inRange": false}
	expected.ShouldBeEqual(t, 0, "IntegerOutOfRange ToInt32", actual)
}

func Test_Cov_IntegerOutOfRange_ToInt(t *testing.T) {
	actual := args.Map{
		"inRange": coremath.IsOutOfRange.Integer.ToInt(1000),
	}
	expected := args.Map{"inRange": false}
	expected.ShouldBeEqual(t, 0, "IntegerOutOfRange ToInt", actual)
}

// ── IsOutOfRange.Integer64 ──

func Test_Cov_Integer64OutOfRange_Byte(t *testing.T) {
	actual := args.Map{
		"inRange":  coremath.IsOutOfRange.Integer64.Byte(100),
		"outRange": coremath.IsOutOfRange.Integer64.Byte(300),
	}
	expected := args.Map{"inRange": false, "outRange": true}
	expected.ShouldBeEqual(t, 0, "Integer64OutOfRange Byte", actual)
}

func Test_Cov_Integer64OutOfRange_UnsignedInt16(t *testing.T) {
	actual := args.Map{
		"inRange":  coremath.IsOutOfRange.Integer64.UnsignedInt16(100),
		"outRange": coremath.IsOutOfRange.Integer64.UnsignedInt16(70000),
	}
	expected := args.Map{"inRange": false, "outRange": true}
	expected.ShouldBeEqual(t, 0, "Integer64OutOfRange UnsignedInt16", actual)
}

func Test_Cov_Integer64OutOfRange_UnsignedInt32(t *testing.T) {
	actual := args.Map{
		"inRange":  coremath.IsOutOfRange.Integer64.UnsignedInt32(100),
		"outRange": coremath.IsOutOfRange.Integer64.UnsignedInt32(-1),
	}
	expected := args.Map{"inRange": false, "outRange": true}
	expected.ShouldBeEqual(t, 0, "Integer64OutOfRange UnsignedInt32", actual)
}

func Test_Cov_Integer64OutOfRange_UnsignedInt64(t *testing.T) {
	actual := args.Map{
		"positive": coremath.IsOutOfRange.Integer64.UnsignedInt64(100),
		"negative": coremath.IsOutOfRange.Integer64.UnsignedInt64(-1),
	}
	expected := args.Map{"positive": false, "negative": true}
	expected.ShouldBeEqual(t, 0, "Integer64OutOfRange UnsignedInt64", actual)
}

func Test_Cov_Integer64OutOfRange_Int8(t *testing.T) {
	actual := args.Map{
		"inRange":  coremath.IsOutOfRange.Integer64.Int8(50),
		"outRange": coremath.IsOutOfRange.Integer64.Int8(200),
	}
	expected := args.Map{"inRange": false, "outRange": true}
	expected.ShouldBeEqual(t, 0, "Integer64OutOfRange Int8", actual)
}

func Test_Cov_Integer64OutOfRange_Int16(t *testing.T) {
	actual := args.Map{
		"inRange":  coremath.IsOutOfRange.Integer64.Int16(1000),
		"outRange": coremath.IsOutOfRange.Integer64.Int16(40000),
	}
	expected := args.Map{"inRange": false, "outRange": true}
	expected.ShouldBeEqual(t, 0, "Integer64OutOfRange Int16", actual)
}

func Test_Cov_Integer64OutOfRange_Int32(t *testing.T) {
	actual := args.Map{
		"inRange":  coremath.IsOutOfRange.Integer64.Int32(1000),
		"outRange": coremath.IsOutOfRange.Integer64.Int32(int64(math.MaxInt32) + 1),
	}
	expected := args.Map{"inRange": false, "outRange": true}
	expected.ShouldBeEqual(t, 0, "Integer64OutOfRange Int32", actual)
}

func Test_Cov_Integer64OutOfRange_Int(t *testing.T) {
	actual := args.Map{
		"inRange": coremath.IsOutOfRange.Integer64.Int(1000),
	}
	expected := args.Map{"inRange": false}
	expected.ShouldBeEqual(t, 0, "Integer64OutOfRange Int", actual)
}
