package args

import (
	"reflect"

	"gitlab.com/auk-go/core/internal/reflectinternal"
)

type FuncWrap struct {
	Name                 string         `json:",omitempty"`
	FullName             string         `json:",omitempty"`
	Func                 any            `json:"-"`
	isInvalid            bool           `json:"IsInvalid,omitempty"`
	rvType               reflect.Type   `json:"-"`
	rv                   reflect.Value  `json:"-"`
	inArgsTypesNames     []string       `json:"-"`
	inArgsTypes          []reflect.Type `json:"-"`
	outArgsTypes         []reflect.Type `json:"-"`
	pkgNameOnly          string
	funcDirectInvokeName string
	pkgPath              string
	inArgsMap            Map
	outArgsMap           Map
	inArgsNames          []string
	outArgsTypesNames    []string
	outArgsNames         []string
}

func (it *FuncWrap) GetFuncName() string {
	if it == nil {
		return ""
	}

	return it.Name
}

func (it *FuncWrap) GetPascalCaseFuncName() string {
	if it == nil {
		return ""
	}

	return pascalCaseFunc(it.Name)
}

func (it *FuncWrap) HasValidFunc() bool {
	return it != nil &&
		!it.isInvalid &&
		it.rv.IsValid() &&
		reflectinternal.Is.Func(it.Func)
}

func (it *FuncWrap) IsInvalid() bool {
	return it == nil ||
		it.isInvalid ||
		!it.rv.IsValid() ||
		!it.HasValidFunc()
}

func (it *FuncWrap) IsValid() bool {
	return !it.IsInvalid()
}

func (it *FuncWrap) PkgPath() string {
	if it.IsInvalid() {
		return ""
	}

	if len(it.pkgPath) > 0 {
		return it.pkgPath
	}

	it.pkgPath = reflectinternal.GetFunc.GetPkgPathFullName(it.FullName)

	return it.pkgPath
}

func (it *FuncWrap) PkgNameOnly() string {
	if it.IsInvalid() {
		return ""
	}

	if len(it.pkgNameOnly) > 0 {
		return it.pkgNameOnly
	}

	it.pkgNameOnly = reflectinternal.Utils.PkgNameOnly(it.Func)

	return it.pkgNameOnly
}

func (it *FuncWrap) FuncDirectInvokeName() string {
	if it.IsInvalid() {
		return ""
	}

	if len(it.funcDirectInvokeName) > 0 {
		return it.funcDirectInvokeName
	}

	it.funcDirectInvokeName = reflectinternal.GetFunc.FuncDirectInvokeNameUsingFullName(it.FullName)

	return it.funcDirectInvokeName
}

func (it *FuncWrap) GetType() reflect.Type {
	if it.IsInvalid() {
		return nil
	}

	return it.rvType
}

func (it *FuncWrap) IsPublicMethod() bool {
	return it != nil && it.rvType.PkgPath() == ""
}

func (it *FuncWrap) IsPrivateMethod() bool {
	return it != nil && it.rvType.PkgPath() != ""
}

func (it *FuncWrap) IsNotEqual(another *FuncWrap) bool {
	return !it.IsEqual(another)
}

func (it *FuncWrap) IsEqualValue(another FuncWrap) bool {
	return it.IsEqual(&another)
}

func (it *FuncWrap) IsEqual(another *FuncWrap) bool {
	if it == nil && another == nil {
		return true
	}

	if it == nil || another == nil {
		return false
	}

	if it == another {
		return true
	}

	if it.IsInvalid() != another.IsInvalid() {
		return false
	}

	if it.Name != it.Name {
		return false
	}

	if it.IsPublicMethod() != it.IsPublicMethod() {
		return false
	}

	if it.ArgsCount() != it.ArgsCount() {
		return false
	}

	if it.ReturnLength() != it.ReturnLength() {
		return false
	}

	isInArgsOkay, _ := it.InArgsVerifyRv(another.GetInArgsTypes())

	if !isInArgsOkay {
		return false
	}

	isOutArgsOkay, _ := it.OutArgsVerifyRv(another.GetOutArgsTypes())

	if !isOutArgsOkay {
		return false
	}

	return true
}

func (it FuncWrap) AsFuncWrapper() FuncWrapper {
	return &it
}
