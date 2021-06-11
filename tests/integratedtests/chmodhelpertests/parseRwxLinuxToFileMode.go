package chmodhelpertests

import (
	"os"

	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
)

func parseRwxLinuxToFileMode(rwxOwnerGroupOther *chmodins.RwxOwnerGroupOther) os.FileMode {
	varWrapper, err := chmodhelper.ParseRwxOwnerGroupOtherToRwxVariableWrapper(
		rwxOwnerGroupOther)

	if err != nil {
		panic(err)
	}

	return varWrapper.ToCompileFixedPtr().ToFileMode()
}
