package iserrortests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/iserror"
)

func Test_Cov_Equal_BothSameMessage(t *testing.T) {
	a := errors.New("err")
	b := errors.New("err")
	if !iserror.Equal(a, b) {
		t.Error("expected equal by message")
	}
}

func Test_Cov_Equal_DiffMessage(t *testing.T) {
	a := errors.New("a")
	b := errors.New("b")
	if iserror.Equal(a, b) {
		t.Error("expected not equal")
	}
}
