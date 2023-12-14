package codegen

type OverridingNames struct {
	TestPkgName string
	FuncCall    string
}

func (it OverridingNames) HasTestPkgName() bool {
	return len(it.TestPkgName) > 0
}

func (it OverridingNames) HasFuncCall() bool {
	return len(it.FuncCall) > 0
}
