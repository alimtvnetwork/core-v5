package codegen

import "gitlab.com/auk-go/core/coredata/corestr"

type newGoCodeCreator struct{}

func (it newGoCodeCreator) Create(
	generator BaseGenerator,
	codes ...string,
) *GoCode {
	return &GoCode{
		codes:       corestr.New.SimpleSlice.Strings(codes),
		imports:     generator.AllPackages(),
		testPkgName: generator.TestPkgName(),
	}
}

func (it newGoCodeCreator) Empty() *GoCode {
	return &GoCode{
		codes:       corestr.New.SimpleSlice.Empty(),
		imports:     corestr.New.Hashset.Empty(),
		testPkgName: "",
	}
}

func (it newGoCodeCreator) All(
	testPkgName string,
	allImports *corestr.Hashset,
	codes ...string,
) *GoCode {
	return &GoCode{
		codes:       corestr.New.SimpleSlice.Strings(codes),
		imports:     allImports,
		testPkgName: testPkgName,
	}
}
