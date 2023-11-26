package reflectinternal

import (
	"reflect"
	"runtime"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/stringslice"
)

type functionGetter struct{}

func (it functionGetter) GetFunc(i interface{}) *runtime.Func {
	if IsNull(i) {
		return nil
	}

	rv := reflect.ValueOf(i)

	if rv.Kind() != reflect.Func {
		return nil
	}

	return runtime.FuncForPC(rv.Pointer())
}

// GetFuncFullName
//
// Get the function name, passing non function may result panic
func (it functionGetter) GetFuncFullName(i interface{}) string {
	f := GetFunc(i)

	if f == nil {
		return ""
	}

	return f.Name()
}

func (it functionGetter) GetFuncFullNameWithName(i interface{}) (fullName, name string) {
	fullName = GetFuncFullName(i)

	if len(fullName) == 0 {
		return "", ""
	}

	_, _, funcNameOnly := MethodNamePackageName(fullName)

	return fullName, fixFinalFuncName(funcNameOnly)
}

func (it functionGetter) GetFuncName(i interface{}) string {
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

func fixFinalFuncName(funcNameOnly string) string {
	if strings.HasSuffix(funcNameOnly, "-fm") {
		return funcNameOnly[:len(funcNameOnly)-3]
	}

	return funcNameOnly
}

func MethodNamePackageName(fullFuncName string) (fullMethodName, packageName, methodName string) {
	if fullFuncName == "" {
		return "", "", ""
	}

	hasComplexName :=
		strings.HasPrefix(
			fullFuncName,
			gitlabDotCom,
		) ||
			strings.HasPrefix(
				fullFuncName,
				gitHubDotCom,
			) ||
			strings.LastIndexByte(
				fullFuncName,
				constants.ForwardChar,
			) > -1

	if hasComplexName {
		forwardSlashFound := strings.LastIndexByte(
			fullFuncName,
			constants.ForwardChar,
		)

		return MethodNamePackageName(fullFuncName[forwardSlashFound+1:])
	}

	splitsByDot := strings.Split(fullFuncName, constants.Dot)
	packageName, methodName = stringslice.FirstLastDefault(splitsByDot)

	return fixFinalFuncName(fullFuncName), packageName, fixFinalFuncName(methodName)
}
