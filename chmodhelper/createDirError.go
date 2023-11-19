package chmodhelper

import "errors"

func createDirError(dirPath string, err error) error {
	if IsPathExists(dirPath) && !IsDirectory(dirPath) {
		return errors.New(
			"dir : " + dirPath +
				", applyChmod :" + dirDefaultChmod.String() +
				", path exist but it is not a dir.",
		)
	}

	// has err
	return errors.New(
		"dir : " + dirPath +
			", applyChmod :" + dirDefaultChmod.String() +
			", " + err.Error(),
	)
}
