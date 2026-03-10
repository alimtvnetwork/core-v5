package issetter

import (
	"github.com/alimtvnetwork/core/errcore"
)

func NewMust(name string) Value {
	newType, err := New(name)
	errcore.HandleErr(err)

	return newType
}
