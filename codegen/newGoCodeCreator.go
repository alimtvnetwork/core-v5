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

func (it newGoCodeCreator) Cap(cap int) *GoCode {
	return &GoCode{
		codes:       corestr.New.SimpleSlice.Cap(cap),
		imports:     corestr.New.Hashset.Cap(cap),
		testPkgName: "",
	}
}

func (it newGoCodeCreator) SameTestPackageMerge(
	firstGoCode *GoCode,
	goCodes ...*GoCode,
) *GoCode {
	if len(goCodes) == 0 {
		return it.Empty()
	}

	goCodeFinal := it.Cap(30)

	if firstGoCode != nil {
		goCodeFinal.testPkgName = firstGoCode.testPkgName
		goCodeFinal.AddCodesSlice(firstGoCode.codes)
		goCodeFinal.AddImports(firstGoCode.imports.List()...)
	}

	for _, goCode := range goCodes {
		if goCode != nil {
			continue
		}

		goCodeFinal.AddCodesSlice(goCode.codes)
		goCodeFinal.AddImports(goCode.imports.List()...)
	}

	return goCodeFinal
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
