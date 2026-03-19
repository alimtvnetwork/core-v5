package ostypetests

import (
	"testing"

	"github.com/alimtvnetwork/core/ostype"
)

func Test_Cov_Group_IsWindows(t *testing.T) {
	if !ostype.WindowsGroup.IsWindows() {
		t.Error("expected windows")
	}
}

func Test_Cov_Group_IsUnix(t *testing.T) {
	if !ostype.UnixGroup.IsUnix() {
		t.Error("expected unix")
	}
}

func Test_Cov_Group_IsAndroid(t *testing.T) {
	if !ostype.AndroidGroup.IsAndroid() {
		t.Error("expected android")
	}
}

func Test_Cov_Group_IsInvalidGroup(t *testing.T) {
	if !ostype.InvalidGroup.IsInvalidGroup() {
		t.Error("expected invalid")
	}
}

func Test_Cov_Variation_Group_Android(t *testing.T) {
	g := ostype.Android.Group()
	if !g.IsAndroid() {
		t.Error("expected android group")
	}
}

func Test_Cov_Variation_Group_Unix(t *testing.T) {
	g := ostype.Linux.Group()
	if !g.IsUnix() {
		t.Error("expected unix group")
	}
}

func Test_Cov_Variation_IsActualGroupUnix(t *testing.T) {
	if !ostype.Linux.IsActualGroupUnix() {
		t.Error("expected actual group unix")
	}
}

func Test_Cov_Variation_IsPossibleUnixGroup(t *testing.T) {
	if !ostype.Linux.IsPossibleUnixGroup() {
		t.Error("expected possible unix")
	}
	if ostype.Windows.IsPossibleUnixGroup() {
		t.Error("windows should not be unix")
	}
}
