package reflectinternal

func GetFuncOnly(i interface{}) string {
	funcFullName := GetFuncFullName(i)

	if len(funcFullName) == 0 {
		return ""
	}

	_, _, funcNameOnly := MethodNamePackageName(funcFullName)

	return funcNameOnly
}
