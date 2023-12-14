package codegen

import (
	"strings"

	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/errcore"
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

func (it utils) GetOptimizePackageHeaders(
	code string,
	headerPackages *corestr.Hashset,
) *corestr.Hashset {
	headerLines := headerPackages.SimpleSlice()
	isImportStarted := false
	var removeIndexes []int

	for i, h := range headerLines.List() {
		h = strings.TrimSpace(h)
		if !isImportStarted && strings.HasPrefix(h, "import") {
			isImportStarted = true

			continue
		}

		if !isImportStarted {
			continue
		}

		if h == ")" || h == "" {
			continue
		}

		// after import
		_, pkgName := GetPkgName(h)
		pkgNameNext := pkgName + "."

		if !strings.Contains(code, pkgNameNext) {
			removeIndexes = append(removeIndexes, i)
		}
	}

	lines, err := headerLines.RemoveIndexes(removeIndexes...)
	errcore.HandleErr(err)

	return corestr.New.Hashset.Strings(lines.List())
}
