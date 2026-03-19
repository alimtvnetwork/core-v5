package ostypetests

import (
	"testing"

	"github.com/alimtvnetwork/core/ostype"
)

func Test_QW_Group_IsAnyEnumsEqual_NoMatch(t *testing.T) {
	g := ostype.Linux
	if g.IsAnyEnumsEqual() {
		t.Fatal("expected false for empty enums")
	}
}

func Test_QW_Group_MinByte(t *testing.T) {
	_ = ostype.Linux.MinByte()
}

func Test_QW_Variation_IsAnyEnumsEqual_NoMatch(t *testing.T) {
	v := ostype.Linux
	if v.IsAnyEnumsEqual() {
		t.Fatal("expected false for empty enums")
	}
}

func Test_QW_Variation_MinByte(t *testing.T) {
	_ = ostype.Linux.MinByte()
}
