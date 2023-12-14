package codegen

import "gitlab.com/auk-go/core/coredata/corestr"

type utils struct{}

func (it utils) PascalCase() string {

}

func (it utils) CamelCase() string {

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
