package reflectinternal

import "strings"

func GetFuncName(i interface{}) string {
	if IsNull(i) {
		return ""
	}

	funcFullName := GetFuncFullName(i)

	if len(funcFullName) == 0 {
		return ""
	}

	_, _, funcNameOnly := MethodNamePackageName(funcFullName)

	if strings.HasSuffix(funcNameOnly, "-fm") {
		return funcNameOnly[:len(funcNameOnly)-3]
	}

	return funcNameOnly
}
