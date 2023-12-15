package codegen

import (
	"sort"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coredata/stringslice"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/internal/convertinteranl"
)

type GoCode struct {
	codes                 *corestr.SimpleSlice
	imports               *corestr.Hashset
	testPkgName           string
	isDefaultPackageAdded bool
}

func (it *GoCode) Codes() *corestr.SimpleSlice {
	return it.codes
}

func (it *GoCode) Imports() *corestr.Hashset {
	return it.imports
}

func (it *GoCode) TestPkgName() string {
	return it.testPkgName
}

func (it *GoCode) AddImports(packages ...string) *GoCode {
	if len(packages) == 0 {
		return it
	}

	for _, pkg := range packages {
		pkg = corestr.StringUtils.WrapDoubleIfMissing(pkg)

		it.imports.Add(pkg)
	}

	return it
}

func (it *GoCode) AddCodes(codes ...string) *GoCode {
	if len(codes) == 0 {
		return it
	}

	for _, code := range codes {
		trimmedCode := strings.TrimSpace(code)

		if trimmedCode == "" {
			continue
		}

		it.codes.Add(trimmedCode)
	}

	return it
}

func (it *GoCode) AddCodesSlice(codes *corestr.SimpleSlice) *GoCode {
	if codes.IsEmpty() {
		return it
	}

	for _, code := range *codes {
		trimmedCode := strings.TrimSpace(code)

		if trimmedCode == "" {
			continue
		}

		it.codes.Add(trimmedCode)
	}

	return it
}

func (it *GoCode) IsCodeDefined() bool {
	if it == nil {
		return false
	}

	return it.codes.HasAnyItem()
}

func (it *GoCode) JoinCode() string {
	if it == nil {
		return ""
	}

	return it.codes.JoinLine()
}

func (it *GoCode) OptimizeImports(fullCode string) (organizedImports *corestr.Hashset) {
	if it == nil {
		return corestr.Empty.Hashset()
	}

	it.addDefaultPackages()

	it.imports = Utils.GetOptimizePackageImports(
		fullCode,
		it.imports,
	)

	return it.imports
}

func (it *GoCode) addDefaultPackages() *GoCode {
	if it.isDefaultPackageAdded {
		return it
	}

	it.AddImports(defaultPackages...)
	it.isDefaultPackageAdded = true

	return it
}

func (it *GoCode) CompileImports(fullCode string) string {
	if it == nil {
		return ""
	}

	it.OptimizeImports(fullCode)
	allImports := it.imports.List()
	sort.Strings(allImports)

	packagesTemplate := map[string]string{
		"$packageName": it.testPkgName,
		"$newPackages": strings.Join(allImports, "\n\t"),
	}

	packageHeader := Utils.ReplaceTemplate(
		testPkgHeaderTemplate,
		packagesTemplate,
	)

	return packageHeader
}

func (it *GoCode) Concat(goCodes ...*GoCode) *GoCode {
	if it == nil {
		return NewGoCode.Empty()
	}

	return NewGoCode.SameTestPackageMerge(
		it,
		goCodes...,
	)
}

func (it *GoCode) Append(goCodes ...*GoCode) *GoCode {
	if it == nil {
		return NewGoCode.Empty()
	}

	for _, goCode := range goCodes {
		if goCode == nil {
			continue
		}

		it.AddImports(goCode.imports.List()...)
		it.AddCodesSlice(goCode.codes)
	}

	return it
}

func (it *GoCode) Dispose() {
	if it == nil {
		return
	}

	it.codes.Dispose()
	it.imports.Dispose()
	it.testPkgName = ""
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
