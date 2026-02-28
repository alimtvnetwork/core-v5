package chmodhelper

func CreateDirFilesWithRwxPermissions(
	isRemoveAllDirBeforeCreate bool,
	dirFilesWithRwxPermissions []DirFilesWithRwxPermission,
) error {
	if dirFilesWithRwxPermissions == nil || len(dirFilesWithRwxPermissions) == 0 {
		return nil
	}

	for _, pathCreate := range dirFilesWithRwxPermissions {
		createErr := CreateDirFilesWithRwxPermission(
			isRemoveAllDirBeforeCreate,
			&pathCreate)

		if createErr != nil {
			return createErr
		}
	}

	return nil
}
