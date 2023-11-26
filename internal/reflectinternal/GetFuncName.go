package reflectinternal

func GetFuncName(i interface{}) string {
	if IsNull(i) {
		return ""
	}

	funcFullName := GetFuncFullName(i)

	if len(funcFullName) == 0 {
		return ""
	}

	_, _, funcNameOnly := MethodNamePackageName(funcFullName)

	return fixFinalFuncName(funcNameOnly)
}
