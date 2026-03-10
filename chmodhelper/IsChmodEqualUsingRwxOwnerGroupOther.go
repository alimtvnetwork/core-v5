package chmodhelper

import "github.com/alimtvnetwork/core/chmodhelper/chmodins"

func IsChmodEqualUsingRwxOwnerGroupOther(
	location string,
	rwx *chmodins.RwxOwnerGroupOther,
) bool {
	if rwx == nil {
		return false
	}

	return IsChmod(
		location,
		rwx.String())
}
