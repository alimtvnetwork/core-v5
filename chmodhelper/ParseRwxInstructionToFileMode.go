package chmodhelper

import (
	"os"

	"github.com/alimtvnetwork/core/chmodhelper/chmodins"
)

func ParseRwxOwnerGroupOtherToFileModeMust(
	rwxOwnerGroupOther *chmodins.RwxOwnerGroupOther,
) os.FileMode {
	fileMode, err := ParseRwxOwnerGroupOtherToFileMode(
		rwxOwnerGroupOther)

	if err != nil {
		panic(err)
	}

	return fileMode
}
