package corecmptests

import (
	"testing"

	"github.com/alimtvnetwork/core/corecmp"
	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ============================================================================
// Integer16
// ============================================================================

func Test_Cov3_Integer16_Equal(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer16(5, 5) == corecomparator.Equal}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer16 returns Equal -- same values", actual)
}

func Test_Cov3_Integer16_LeftLess(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer16(3, 5) == corecomparator.LeftLess}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer16 returns LeftLess -- left < right", actual)
}

func Test_Cov3_Integer16_LeftGreater(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer16(7, 5) == corecomparator.LeftGreater}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer16 returns LeftGreater -- left > right", actual)
}

// ============================================================================
// Integer64
// ============================================================================

func Test_Cov3_Integer64_Equal(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer64(100, 100) == corecomparator.Equal}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer64 returns Equal -- same values", actual)
}

func Test_Cov3_Integer64_LeftLess(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer64(50, 100) == corecomparator.LeftLess}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer64 returns LeftLess -- left < right", actual)
}

func Test_Cov3_Integer64_LeftGreater(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer64(200, 100) == corecomparator.LeftGreater}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer64 returns LeftGreater -- left > right", actual)
}

// ============================================================================
// IntegerPtr
// ============================================================================

func Test_Cov3_IntegerPtr_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.IntegerPtr(nil, nil) == corecomparator.Equal}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IntegerPtr returns Equal -- both nil", actual)
}

func Test_Cov3_IntegerPtr_LeftNil(t *testing.T) {
	v := 5
	actual := args.Map{"result": corecmp.IntegerPtr(nil, &v) == corecomparator.NotEqual}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IntegerPtr returns NotEqual -- left nil", actual)
}

func Test_Cov3_IntegerPtr_RightNil(t *testing.T) {
	v := 5
	actual := args.Map{"result": corecmp.IntegerPtr(&v, nil) == corecomparator.NotEqual}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IntegerPtr returns NotEqual -- right nil", actual)
}

func Test_Cov3_IntegerPtr_BothEqual(t *testing.T) {
	a, b := 5, 5
	actual := args.Map{"result": corecmp.IntegerPtr(&a, &b) == corecomparator.Equal}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IntegerPtr returns Equal -- both 5", actual)
}

// ============================================================================
// BytePtr
// ============================================================================

func Test_Cov3_BytePtr_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.BytePtr(nil, nil) == corecomparator.Equal}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "BytePtr returns Equal -- both nil", actual)
}

func Test_Cov3_BytePtr_LeftNil(t *testing.T) {
	v := byte(5)
	actual := args.Map{"result": corecmp.BytePtr(nil, &v) == corecomparator.NotEqual}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "BytePtr returns NotEqual -- left nil", actual)
}

func Test_Cov3_BytePtr_BothEqual(t *testing.T) {
	a, b := byte(5), byte(5)
	actual := args.Map{"result": corecmp.BytePtr(&a, &b) == corecomparator.Equal}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "BytePtr returns Equal -- both 5", actual)
}

// ============================================================================
// IsIntegersEqual
// ============================================================================

func Test_Cov3_IsIntegersEqual_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.IsIntegersEqual(nil, nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual returns true -- both nil", actual)
}

func Test_Cov3_IsIntegersEqual_LeftNil(t *testing.T) {
	actual := args.Map{"result": corecmp.IsIntegersEqual(nil, []int{1})}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual returns false -- left nil", actual)
}

func Test_Cov3_IsIntegersEqual_Same(t *testing.T) {
	actual := args.Map{"result": corecmp.IsIntegersEqual([]int{1, 2}, []int{1, 2})}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual returns true -- same values", actual)
}

func Test_Cov3_IsIntegersEqual_Different(t *testing.T) {
	actual := args.Map{"result": corecmp.IsIntegersEqual([]int{1, 2}, []int{1, 3})}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual returns false -- different values", actual)
}

// ============================================================================
// IsStringsEqualPtr
// ============================================================================

func Test_Cov3_IsStringsEqualPtr_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.IsStringsEqualPtr(nil, nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr returns true -- both nil", actual)
}

func Test_Cov3_IsStringsEqualPtr_LeftNil(t *testing.T) {
	actual := args.Map{"result": corecmp.IsStringsEqualPtr(nil, []string{"a"})}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr returns false -- left nil", actual)
}

func Test_Cov3_IsStringsEqualPtr_DiffLen(t *testing.T) {
	actual := args.Map{"result": corecmp.IsStringsEqualPtr([]string{"a"}, []string{"a", "b"})}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr returns false -- different lengths", actual)
}

func Test_Cov3_IsStringsEqualPtr_Same(t *testing.T) {
	actual := args.Map{"result": corecmp.IsStringsEqualPtr([]string{"a", "b"}, []string{"a", "b"})}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr returns true -- same values", actual)
}
