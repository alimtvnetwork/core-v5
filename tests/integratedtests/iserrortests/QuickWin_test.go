package iserrortests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/iserror"
)

func Test_QW_Equal_BothNonNilSameMessage(t *testing.T) {
	// Cover the Error() comparison branch
	e1 := errors.New("same")
	e2 := errors.New("same")
	if !iserror.Equal(e1, e2) {
		t.Fatal("expected true for same message")
	}
}

func Test_QW_Equal_BothNonNilDiffMessage(t *testing.T) {
	e1 := errors.New("a")
	e2 := errors.New("b")
	if iserror.Equal(e1, e2) {
		t.Fatal("expected false for different message")
	}
}

func Test_QW_Equal_LeftNilRightNot(t *testing.T) {
	if iserror.Equal(nil, errors.New("a")) {
		t.Fatal("expected false")
	}
}
