package chmodhelpertests

import (
	"os"
	"path"

	"gitlab.com/evatix-go/core/internal/fsinternal"
	"gitlab.com/evatix-go/core/msgtype"
	"gitlab.com/evatix-go/core/tests/testwrappers/chmodhelpertestwrappers"
)

func createDefaultPaths(pathCreators *[]*chmodhelpertestwrappers.CreatePathsInstruction) {
	for _, pathCreate := range *pathCreators {
		dir := pathCreate.Dir
		if fsinternal.IsPathExists(dir) {
			err1 := os.RemoveAll(dir)

			msgtype.SimpleHandleErr(err1, dir)
		}

		fileMode := parseRwxLinuxToFileMode(&pathCreate.ApplyRwx)

		mkDirErr := os.MkdirAll(dir, fileMode)

		msgtype.SimpleHandleErr(mkDirErr, dir)

		for _, filePath := range pathCreate.Files {
			compiledPath := path.Join(dir, filePath)
			_, err := os.Create(compiledPath)

			msgtype.SimpleHandleErr(err, compiledPath)

			chmodErr := os.Chmod(compiledPath, fileMode)

			msgtype.SimpleHandleErr(chmodErr, compiledPath)
		}
	}
}
