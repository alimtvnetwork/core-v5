package codegen

import (
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/internal/convertinteranl"
)

type utils struct{}

func (it utils) PascalCase(s string) string {
	return convertinteranl.Util.String.PascalCase(s)
}

func (it utils) CamelCase(s string) string {
	return convertinteranl.Util.String.CamelCase(s)
}

func (it utils) AllPackages(
	currentFuncPackage string,
	additionalPackages ...string,
) *corestr.Hashset {
	newPackages := corestr.
		New.
		Hashset.
		Strings(additionalPackages).
		Add(currentFuncPackage).
		WrapDoubleQuote()

	return newPackages
}
