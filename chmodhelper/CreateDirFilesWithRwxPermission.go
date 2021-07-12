package chmodhelper

func CreateDirFilesWithRwxPermission(
	isRemoveAllDirBeforeCreate bool,
	dirFilesWithRwxPermission *DirFilesWithRwxPermission,
) error {
	fileMode, fileModeErr := ParseRwxOwnerGroupOtherToFileMode(
		&dirFilesWithRwxPermission.ApplyRwx)

	if fileModeErr != nil {
		return fileModeErr
	}

	err := CreateDirWithFiles(
		isRemoveAllDirBeforeCreate,
		fileMode,
		&dirFilesWithRwxPermission.DirWithFiles)

	if err != nil {
		return err
	}

	return nil
}
