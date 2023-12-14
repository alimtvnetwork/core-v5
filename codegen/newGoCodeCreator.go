package codegen

import "gitlab.com/auk-go/core/coredata/corestr"

type newGoCodeCreator struct{}

func (it newGoCodeCreator) Create(
	generator BaseGenerator, codes ...string,
) *GoCode {
	return &GoCode{
		Codes:       corestr.New.SimpleSlice.Strings(codes),
		Packages:    generator.AllPackages(),
		TestPkgName: generator.TestPkgName(),
	}
}

func (it newGoCodeCreator) All(
	testPkgName string,
	allImports *corestr.Hashset,
	codes ...string,
) *GoCode {
	return &GoCode{
		Codes:       corestr.New.SimpleSlice.Strings(codes),
		Packages:    allImports,
		TestPkgName: testPkgName,
	}
}
