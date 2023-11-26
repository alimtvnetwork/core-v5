package reflectinternal

import (
	"reflect"
	"runtime"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/stringslice"
)

type getFunc struct{}

func (it getFunc) RunTime(i interface{}) *runtime.Func {
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
func (it getFunc) FullName(i interface{}) string {
	f := it.RunTime(i)

	if f == nil {
		return ""
	}

	return f.Name()
}

func (it getFunc) FullNameWithName(i interface{}) (fullName, name string) {
	fullName = it.FullName(i)

	if len(fullName) == 0 {
		return "", ""
	}

	_, _, funcNameOnly := it.All(fullName)

	return fullName, it.fixFinalFuncName(funcNameOnly)
}

func (it getFunc) Name(i interface{}) string {
	if IsNull(i) {
		return ""
	}

	funcFullName := it.FullName(i)

	if len(funcFullName) == 0 {
		return ""
	}

	_, _, funcNameOnly := it.All(funcFullName)

	return it.fixFinalFuncName(funcNameOnly)
}

func (it getFunc) fixFinalFuncName(funcNameOnly string) string {
	if strings.HasSuffix(funcNameOnly, "-fm") {
		return funcNameOnly[:len(funcNameOnly)-3]
	}

	return funcNameOnly
}

func (it getFunc) All(fullFuncName string) (fullMethodName, packageName, methodName string) {
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

		return it.All(fullFuncName[forwardSlashFound+1:])
	}

	splitsByDot := strings.Split(fullFuncName, constants.Dot)
	packageName, methodName = stringslice.FirstLastDefault(splitsByDot)

	return it.fixFinalFuncName(fullFuncName), packageName, it.fixFinalFuncName(methodName)
}
