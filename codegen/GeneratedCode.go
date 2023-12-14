package codegen

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coredata/stringslice"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/internal/convertinteranl"
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
		trimmedCode := strings.TrimSpace(code)

		if trimmedCode == "" {
			continue
		}

		it.Codes.Add(trimmedCode)
	}

	return it
}

func (it *GeneratedCode) IsCodeDefined() bool {
	if it == nil {
		return false
	}

	return it.Codes.HasAnyItem()
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

func (it *GeneratedCode) CompileFullCode() (string, error) {
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

func (it *GeneratedCode) FormatCode(code string) (string, error) {
	s, err := convertinteranl.CodeFormatter.Golang(code)

	return s, errcore.StackEnhance.Error(err)
}
