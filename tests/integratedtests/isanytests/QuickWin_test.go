package isanytests

import (
	"testing"

	"github.com/alimtvnetwork/core/isany"
)

func Test_QW_Conclusive_LeftNilRightDefined(t *testing.T) {
	// Cover the branch: left==nil || right==nil (with left nil, right defined)
	isEqual, isConclusive := isany.Conclusive(nil, "hello")
	if isEqual {
		t.Fatal("expected not equal")
	}
	if !isConclusive {
		t.Fatal("expected conclusive")
	}
}

func Test_QW_JsonEqual_BothMarshalError(t *testing.T) {
	// Cover the branch where both marshal errors exist and are different
	ch1 := make(chan int)
	ch2 := make(chan string)
	result := isany.JsonEqual(ch1, ch2)
	if result {
		t.Fatal("expected false for different marshal errors")
	}
}
