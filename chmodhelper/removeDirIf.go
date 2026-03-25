package chmodhelper

import (
	"github.com/alimtvnetwork/core/internal/pathinternal"
)

func removeDirIf(
	isRemoveAllDirBeforeCreate bool,
	dir string,
	funcName string,
) error {
	return pathinternal.RemoveDirIf(
		isRemoveAllDirBeforeCreate,
		dir,
		funcName,
	)
}
