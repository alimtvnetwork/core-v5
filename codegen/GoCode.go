package codegen

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coredata/stringslice"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/internal/convertinteranl"
)

type GoCode struct {
	Codes       *corestr.SimpleSlice
	Packages    *corestr.Hashset
	TestPkgName string
}

func (it *GoCode) AddPackages(packages ...string) *GoCode {
	if len(packages) == 0 {
		return it
	}

	for _, pkg := range packages {
		pkg = corestr.StringUtils.WrapDoubleIfMissing(pkg)

		it.Packages.Add(pkg)
	}

	return it
}

func (it *GoCode) AddCode(codes ...string) *GoCode {
	if len(codes) == 0 {
		return it
	}

	for _, code := range codes {
		trimmedCode := strings.TrimSpace(code)

		if trimmedCode == "" {
			continue
		}

		it.Codes.Add(trimmedCode)
	}

	return it
}

func (it *GoCode) IsCodeDefined() bool {
	if it == nil {
		return false
	}

	return it.Codes.HasAnyItem()
}

func (it *GoCode) JoinCode() string {
	if it == nil {
		return ""
	}

	return it.Codes.JoinLine()
}

func (it *GoCode) OptimizeImports(fullCode string) (organizedImports *corestr.Hashset) {
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

func (it *GoCode) addDefaultPackages() *GoCode {
	return it.AddPackages(defaultPackages...)
}

func (it *GoCode) CompileImports(fullCode string) string {
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

func (it *GoCode) CompileFullCode() (string, error) {
	if it == nil {
		return "", nil
	}

	fullCode := it.JoinCode()
	compiledImports := it.CompileImports(fullCode)

	finalCode := stringslice.Joins(
		constants.NewLineUnix,
		compiledImports,
		"",
		fullCode,
		"",
	)

	return it.FormatCode(finalCode)
}

func (it *GoCode) FormatCode(code string) (string, error) {
	s, err := convertinteranl.CodeFormatter.Golang(code)

	return s, errcore.StackEnhance.Error(err)
}
