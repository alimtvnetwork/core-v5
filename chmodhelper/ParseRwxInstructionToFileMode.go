package chmodhelper

import (
	"os"

	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
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
