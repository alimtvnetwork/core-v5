package chmodhelper

func GetExistingChmodRwxWrapperMustPtr(
	filePath string,
) *RwxWrapper {
	wrapperPtr, err := GetExistingChmodRwxWrapperPtr(filePath)

	if err != nil {
		panic(err)
	}

	return wrapperPtr
}
