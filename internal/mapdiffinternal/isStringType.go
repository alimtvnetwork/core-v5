package mapdiffinternal

func isStringType(anyItem any) bool {
	_, isSuccess := anyItem.(string)

	return isSuccess
}
