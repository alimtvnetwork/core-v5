package stringutil

import "github.com/alimtvnetwork/core/constants"

func IsNotEmpty(str string) bool {
	return str != constants.EmptyString
}
