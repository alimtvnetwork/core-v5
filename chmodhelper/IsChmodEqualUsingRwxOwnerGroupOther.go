package chmodhelper

import "gitlab.com/evatix-go/core/chmodhelper/chmodins"

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
