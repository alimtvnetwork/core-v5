package codegen

import "gitlab.com/auk-go/core/coredata/corestr"

type GeneratedCode struct {
	Codes       *corestr.SimpleSlice
	Packages    *corestr.Hashset
	TestPkgName string
}

func (it *GeneratedCode) AddPackages(packages ...string) *GeneratedCode {
	if len(packages) == 0 {
		return it
	}

	for _, pkg := range packages {
		pkg = corestr.StringUtils.WrapDoubleIfMissing(pkg)

		it.Packages.Add(pkg)
	}

	return it
}

func (it *GeneratedCode) AddCode(codes ...string) *GeneratedCode {
	if len(codes) == 0 {
		return it
	}

	for _, code := range codes {
		it.Codes.Add(code)
	}

	return it
}
