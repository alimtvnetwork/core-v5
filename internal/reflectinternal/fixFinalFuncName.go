package reflectinternal

import "strings"

func fixFinalFuncName(funcNameOnly string) string {
	if strings.HasSuffix(funcNameOnly, "-fm") {
		return funcNameOnly[:len(funcNameOnly)-3]
	}

	return funcNameOnly
}
