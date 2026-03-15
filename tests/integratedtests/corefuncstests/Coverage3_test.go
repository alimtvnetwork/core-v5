package corefuncstests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/corefuncs"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── ActionReturnsErrorFuncWrapper — AsActionFunc ──

func Test_Cov3_ActionErr_AsActionFunc(t *testing.T) {
	// Arrange
	w := corefuncs.New.ActionErr("test", func() error { return nil })

	// Act — should not panic
	w.AsActionFunc()()

	// Assert
	actual := args.Map{"called": true}
	expected := args.Map{"called": true}
	expected.ShouldBeEqual(t, 0, "ActionErr AsActionFunc -- no panic", actual)
}

func Test_Cov3_ActionErr_AsActionReturnsErrorFunc_Success(t *testing.T) {
	// Arrange
	w := corefuncs.New.ActionErr("test", func() error { return nil })

	// Act
	err := w.AsActionReturnsErrorFunc()()

	// Assert
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ActionErr success -- nil", actual)
}

// ── InOutErrFuncWrapper — AsActionFunc / AsActionReturnsErrorFunc ──

func Test_Cov3_LegacyInOutErr_AsActionFunc(t *testing.T) {
	// Arrange
	w := corefuncs.New.LegacyInOutErr("test", func(input any) (any, error) {
		return nil, nil
	})

	// Act
	w.AsActionFunc("input")()

	// Assert
	actual := args.Map{"called": true}
	expected := args.Map{"called": true}
	expected.ShouldBeEqual(t, 0, "LegacyInOutErr AsActionFunc", actual)
}

func Test_Cov3_LegacyInOutErr_AsActionReturnsErrorFunc_Fail(t *testing.T) {
	// Arrange
	w := corefuncs.New.LegacyInOutErr("test", func(input any) (any, error) {
		return nil, errors.New("fail")
	})

	// Act
	err := w.AsActionReturnsErrorFunc("input")()

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "LegacyInOutErr failure", actual)
}

func Test_Cov3_LegacyInOutErr_AsActionReturnsErrorFunc_Success(t *testing.T) {
	// Arrange
	w := corefuncs.New.LegacyInOutErr("test", func(input any) (any, error) {
		return "ok", nil
	})

	// Act
	err := w.AsActionReturnsErrorFunc("input")()

	// Assert
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "LegacyInOutErr success", actual)
}

// ── ResultDelegatingFuncWrapper — AsActionFunc / AsActionReturnsErrorFunc ──

func Test_Cov3_LegacyResultDelegating_AsActionFunc(t *testing.T) {
	// Arrange
	w := corefuncs.New.LegacyResultDelegating("test", func(target any) error {
		return nil
	})

	// Act
	w.AsActionFunc("target")()

	// Assert
	actual := args.Map{"called": true}
	expected := args.Map{"called": true}
	expected.ShouldBeEqual(t, 0, "LegacyResultDelegating AsActionFunc", actual)
}

func Test_Cov3_LegacyResultDelegating_AsActionReturnsErrorFunc_Fail(t *testing.T) {
	// Arrange
	w := corefuncs.New.LegacyResultDelegating("test", func(target any) error {
		return errors.New("fail")
	})

	// Act
	err := w.AsActionReturnsErrorFunc("target")()

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "LegacyResultDelegating failure", actual)
}

func Test_Cov3_LegacyResultDelegating_AsActionReturnsErrorFunc_Success(t *testing.T) {
	// Arrange
	w := corefuncs.New.LegacyResultDelegating("test", func(target any) error {
		return nil
	})

	// Act
	err := w.AsActionReturnsErrorFunc("target")()

	// Assert
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "LegacyResultDelegating success", actual)
}

// ── Generic wrappers — AsActionFunc / AsActionReturnsErrorFunc ──

func Test_Cov3_InOutErrWrapperOf_AsActionFunc(t *testing.T) {
	// Arrange
	w := corefuncs.NewInOutErrWrapper[string, int]("test", func(s string) (int, error) {
		return 0, nil
	})

	// Act
	w.AsActionFunc("x")()

	// Assert
	actual := args.Map{"called": true}
	expected := args.Map{"called": true}
	expected.ShouldBeEqual(t, 0, "InOutErrWrapperOf AsActionFunc", actual)
}

func Test_Cov3_InOutErrWrapperOf_AsActionReturnsErrorFunc_Fail(t *testing.T) {
	// Arrange
	w := corefuncs.NewInOutErrWrapper[string, int]("test", func(s string) (int, error) {
		return 0, errors.New("fail")
	})

	// Act
	err := w.AsActionReturnsErrorFunc("x")()

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "InOutErrWrapperOf failure", actual)
}

func Test_Cov3_InOutFuncWrapperOf_AsActionFunc(t *testing.T) {
	// Arrange
	w := corefuncs.NewInOutWrapper[string, int]("test", func(s string) int { return 0 })

	// Act
	w.AsActionFunc("x")()

	// Assert
	actual := args.Map{"called": true}
	expected := args.Map{"called": true}
	expected.ShouldBeEqual(t, 0, "InOutFuncWrapperOf AsActionFunc", actual)
}

func Test_Cov3_InOutFuncWrapperOf_AsActionReturnsErrorFunc(t *testing.T) {
	// Arrange
	w := corefuncs.NewInOutWrapper[string, int]("test", func(s string) int { return 0 })

	// Act
	err := w.AsActionReturnsErrorFunc("x")()

	// Assert
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "InOutFuncWrapperOf returns nil", actual)
}

func Test_Cov3_InActionErrWrapperOf_AsActionFunc(t *testing.T) {
	// Arrange
	w := corefuncs.NewInActionErrWrapper[string]("test", func(s string) error { return nil })

	// Act
	w.AsActionFunc("x")()

	// Assert
	actual := args.Map{"called": true}
	expected := args.Map{"called": true}
	expected.ShouldBeEqual(t, 0, "InActionErrWrapperOf AsActionFunc", actual)
}

func Test_Cov3_InActionErrWrapperOf_AsActionReturnsErrorFunc_Fail(t *testing.T) {
	// Arrange
	w := corefuncs.NewInActionErrWrapper[string]("test", func(s string) error {
		return errors.New("fail")
	})

	// Act
	err := w.AsActionReturnsErrorFunc("x")()

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "InActionErrWrapperOf failure", actual)
}

func Test_Cov3_InActionErrWrapperOf_AsActionReturnsErrorFunc_Success(t *testing.T) {
	// Arrange
	w := corefuncs.NewInActionErrWrapper[string]("test", func(s string) error { return nil })

	// Act
	err := w.AsActionReturnsErrorFunc("x")()

	// Assert
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "InActionErrWrapperOf success", actual)
}

func Test_Cov3_ResultDelegatingWrapperOf_AsActionFunc(t *testing.T) {
	// Arrange
	w := corefuncs.NewResultDelegatingWrapper[*string]("test", func(t *string) error { return nil })

	// Act
	var s string
	w.AsActionFunc(&s)()

	// Assert
	actual := args.Map{"called": true}
	expected := args.Map{"called": true}
	expected.ShouldBeEqual(t, 0, "ResultDelegatingWrapperOf AsActionFunc", actual)
}

func Test_Cov3_ResultDelegatingWrapperOf_AsActionReturnsErrorFunc_Fail(t *testing.T) {
	// Arrange
	w := corefuncs.NewResultDelegatingWrapper[*string]("test", func(t *string) error {
		return errors.New("fail")
	})

	// Act
	var s string
	err := w.AsActionReturnsErrorFunc(&s)()

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "ResultDelegatingWrapperOf failure", actual)
}

func Test_Cov3_SerializeWrapperOf_AsActionReturnsErrorFunc_Fail(t *testing.T) {
	// Arrange
	w := corefuncs.NewSerializeWrapper[string]("test", func(s string) ([]byte, error) {
		return nil, errors.New("fail")
	})

	// Act
	err := w.AsActionReturnsErrorFunc("x")()

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "SerializeWrapperOf failure", actual)
}

func Test_Cov3_SerializeWrapperOf_AsActionReturnsErrorFunc_Success(t *testing.T) {
	// Arrange
	w := corefuncs.NewSerializeWrapper[string]("test", func(s string) ([]byte, error) {
		return []byte(s), nil
	})

	// Act
	err := w.AsActionReturnsErrorFunc("x")()

	// Assert
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "SerializeWrapperOf success", actual)
}

// ── NamedActionFuncWrapper — AsActionFunc ──

func Test_Cov3_NamedAction_AsActionFunc(t *testing.T) {
	// Arrange
	var called bool
	w := corefuncs.New.NamedAction("test", func(name string) { called = true })

	// Act
	w.AsActionFunc()()

	// Assert
	actual := args.Map{"called": called}
	expected := args.Map{"called": true}
	expected.ShouldBeEqual(t, 0, "NamedAction AsActionFunc", actual)
}
