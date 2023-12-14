package codegen

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coredata/stringslice"
)

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

func (it *GeneratedCode) JoinCode() string {
	if it == nil {
		return ""
	}

	return it.Codes.JoinLine()
}

func (it *GeneratedCode) OptimizeImports(fullCode string) (organizedImports *corestr.Hashset) {
	if it == nil {
		return corestr.Empty.Hashset()
	}

	it.addDefaultPackages()

	it.Packages = Utils.GetOptimizePackageImports(
		fullCode,
		it.Packages,
	)

	return it.Packages
}

func (it *GeneratedCode) addDefaultPackages() *GeneratedCode {
	return it.AddPackages(defaultPackages...)
}

func (it *GeneratedCode) CompileImports(fullCode string) string {
	if it == nil {
		return ""
	}

	it.OptimizeImports(fullCode)

	packagesTemplate := map[string]string{
		"$packageName": it.TestPkgName,
		"$newPackages": it.Packages.JoinLine(),
	}

	packageHeader := Utils.ReplaceTemplate(
		testPkgHeaderTemplate,
		packagesTemplate,
	)

	return packageHeader
}

func (it *GeneratedCode) CompileFullCode() string {
	if it == nil {
		return ""
	}

	fullCode := it.JoinCode()
	compiledImports := it.CompileImports(fullCode)

	return stringslice.Joins(
		constants.NewLineUnix,
		compiledImports,
		"",
		fullCode,
		"",
	)
}
