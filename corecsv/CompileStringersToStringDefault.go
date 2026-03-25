package corecsv

import "github.com/alimtvnetwork/core/constants"

func CompileStringersToStringDefault(
	compileStringerFunctions ...func() string,
) string {
	return CompileStringersToString(
		constants.CommaSpace,
		true,
		false,
		compileStringerFunctions...)
}
