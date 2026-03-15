package iserrortests

import (
	"errors"
	"os/exec"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/iserror"
)

// ── Defined / Empty additional coverage ──

func Test_Cov2_Defined_Error(t *testing.T) {
	actual := args.Map{"result": iserror.Defined(errors.New("x"))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Defined error -- true", actual)
}

func Test_Cov2_Defined_Nil(t *testing.T) {
	actual := args.Map{"result": iserror.Defined(nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Defined nil -- false", actual)
}

func Test_Cov2_Empty_Error(t *testing.T) {
	actual := args.Map{"result": iserror.Empty(errors.New("x"))}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Empty error -- false", actual)
}

func Test_Cov2_Empty_Nil(t *testing.T) {
	actual := args.Map{"result": iserror.Empty(nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Empty nil -- true", actual)
}

// ── ExitError with actual ExitError ──

func Test_Cov2_ExitError_RealExitError(t *testing.T) {
	exitErr := &exec.ExitError{}
	actual := args.Map{"result": iserror.ExitError(exitErr)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "ExitError real ExitError -- true", actual)
}

// ── Equal with same error message different instances ──

func Test_Cov2_Equal_SameMessage(t *testing.T) {
	e1 := errors.New("same msg")
	e2 := errors.New("same msg")
	actual := args.Map{"result": iserror.Equal(e1, e2)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Equal same message different instances -- true", actual)
}

// ── NotEqual both nil ──

func Test_Cov2_NotEqual_BothNil(t *testing.T) {
	actual := args.Map{"result": iserror.NotEqual(nil, nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NotEqual both nil -- false", actual)
}

// ── NotEqualString same ──

func Test_Cov2_NotEqualString_EmptyStrings(t *testing.T) {
	actual := args.Map{"result": iserror.NotEqualString("", "")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NotEqualString empty -- false", actual)
}
