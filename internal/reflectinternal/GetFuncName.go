package reflectinternal

func GetFuncName(i interface{}) string {
	funcFullName := GetFuncFullName(i)

	if len(funcFullName) == 0 {
		return ""
	}

	_, _, funcNameOnly := MethodNamePackageName(funcFullName)

	return funcNameOnly
}
