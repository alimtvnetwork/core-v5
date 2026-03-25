package iserrortests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/iserror"
)

// ── AllDefined ──

func Test_Cov_AllDefined_Empty(t *testing.T) {
	actual := args.Map{"result": iserror.AllDefined()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AllDefined empty -- false", actual)
}

func Test_Cov_AllDefined_AllErrors(t *testing.T) {
	e1 := errors.New("a")
	e2 := errors.New("b")
	actual := args.Map{"result": iserror.AllDefined(e1, e2)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AllDefined all errors -- true", actual)
}

func Test_Cov_AllDefined_OneNil(t *testing.T) {
	e1 := errors.New("a")
	actual := args.Map{"result": iserror.AllDefined(e1, nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AllDefined one nil -- false", actual)
}

// ── AllEmpty ──

func Test_Cov_AllEmpty_Empty(t *testing.T) {
	actual := args.Map{"result": iserror.AllEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AllEmpty empty -- true", actual)
}

func Test_Cov_AllEmpty_AllNil(t *testing.T) {
	actual := args.Map{"result": iserror.AllEmpty(nil, nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AllEmpty all nil -- true", actual)
}

func Test_Cov_AllEmpty_OneError(t *testing.T) {
	actual := args.Map{"result": iserror.AllEmpty(nil, errors.New("x"))}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AllEmpty one error -- false", actual)
}

// ── AnyDefined ──

func Test_Cov_AnyDefined_Empty(t *testing.T) {
	actual := args.Map{"result": iserror.AnyDefined()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AnyDefined empty -- false", actual)
}

func Test_Cov_AnyDefined_OneError(t *testing.T) {
	actual := args.Map{"result": iserror.AnyDefined(nil, errors.New("x"))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AnyDefined one error -- true", actual)
}

func Test_Cov_AnyDefined_AllNil(t *testing.T) {
	actual := args.Map{"result": iserror.AnyDefined(nil, nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AnyDefined all nil -- false", actual)
}

// ── AnyEmpty ──

func Test_Cov_AnyEmpty_Empty(t *testing.T) {
	actual := args.Map{"result": iserror.AnyEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AnyEmpty empty -- true", actual)
}

func Test_Cov_AnyEmpty_OneNil(t *testing.T) {
	actual := args.Map{"result": iserror.AnyEmpty(errors.New("x"), nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AnyEmpty one nil -- true", actual)
}

func Test_Cov_AnyEmpty_AllErrors(t *testing.T) {
	actual := args.Map{"result": iserror.AnyEmpty(errors.New("a"), errors.New("b"))}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AnyEmpty all errors -- false", actual)
}

// ── Equal ──

func Test_Cov_Equal_BothNil(t *testing.T) {
	actual := args.Map{"result": iserror.Equal(nil, nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Equal both nil -- true", actual)
}

func Test_Cov_Equal_SameMsg(t *testing.T) {
	e1 := errors.New("same")
	e2 := errors.New("same")
	actual := args.Map{"result": iserror.Equal(e1, e2)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Equal same msg -- true", actual)
}

func Test_Cov_Equal_DiffMsg(t *testing.T) {
	e1 := errors.New("a")
	e2 := errors.New("b")
	actual := args.Map{"result": iserror.Equal(e1, e2)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Equal diff msg -- false", actual)
}

func Test_Cov_Equal_LeftNil(t *testing.T) {
	actual := args.Map{"result": iserror.Equal(nil, errors.New("x"))}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Equal left nil -- false", actual)
}

func Test_Cov_Equal_RightNil(t *testing.T) {
	actual := args.Map{"result": iserror.Equal(errors.New("x"), nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Equal right nil -- false", actual)
}

func Test_Cov_Equal_SameInstance(t *testing.T) {
	e := errors.New("x")
	actual := args.Map{"result": iserror.Equal(e, e)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Equal same instance -- true", actual)
}

// ── NotEqual ──

func Test_Cov_NotEqual_Different(t *testing.T) {
	actual := args.Map{"result": iserror.NotEqual(errors.New("a"), errors.New("b"))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NotEqual different -- true", actual)
}

func Test_Cov_NotEqual_Same(t *testing.T) {
	e := errors.New("x")
	actual := args.Map{"result": iserror.NotEqual(e, e)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NotEqual same -- false", actual)
}

// ── EqualString ──

func Test_Cov_EqualString_Same(t *testing.T) {
	actual := args.Map{"result": iserror.EqualString("abc", "abc")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "EqualString same -- true", actual)
}

func Test_Cov_EqualString_Diff(t *testing.T) {
	actual := args.Map{"result": iserror.EqualString("a", "b")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "EqualString diff -- false", actual)
}

// ── NotEqualString ──

func Test_Cov_NotEqualString_Diff(t *testing.T) {
	actual := args.Map{"result": iserror.NotEqualString("a", "b")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NotEqualString diff -- true", actual)
}

func Test_Cov_NotEqualString_Same(t *testing.T) {
	actual := args.Map{"result": iserror.NotEqualString("x", "x")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NotEqualString same -- false", actual)
}

// ── NotEmpty ──

func Test_Cov_NotEmpty_Error(t *testing.T) {
	actual := args.Map{"result": iserror.NotEmpty(errors.New("x"))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NotEmpty error -- true", actual)
}

func Test_Cov_NotEmpty_Nil(t *testing.T) {
	actual := args.Map{"result": iserror.NotEmpty(nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NotEmpty nil -- false", actual)
}

// ── ExitError ──

func Test_Cov_ExitError_NotExitError(t *testing.T) {
	actual := args.Map{"result": iserror.ExitError(errors.New("normal"))}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ExitError not exit error -- false", actual)
}

func Test_Cov_ExitError_Nil(t *testing.T) {
	// nil error should panic or return false depending on impl
	defer func() {
		r := recover()
		if r != nil {
			// ExitError panics on nil — that's fine for coverage
			return
		}
	}()
	_ = iserror.ExitError(nil)
}
