package chmodhelpertestwrappers

import "github.com/alimtvnetwork/core/chmodhelper/chmodins"

type RwxCompileValueTestWrapper struct {
	Existing, Input, Expected chmodins.RwxOwnerGroupOther
}
