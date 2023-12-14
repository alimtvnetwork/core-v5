package codegen

import (
	"gitlab.com/auk-go/core/coreutils/stringutil"
	"gitlab.com/auk-go/core/internal/convertinteranl"
)

var (
	NewCodeOutput = newCodeOutputCreator{}
	NewGoCode     = newGoCodeCreator{}

	defaultPackages = []string{
		"testing",
		"gitlab.com/auk-go/core/coredata/corestr",
		"gitlab.com/auk-go/core/coretests",
		"gitlab.com/auk-go/core/coretests/args",
	}

	vars = unitVariables{
		PackageName:      "$packageName",
		NewPackages:      "$newPackages",
		FuncName:         "$FuncName",
		ArrangeType:      "$ArrangeType",
		LinesPossible:    "$linesPossible",
		ActArgsSetup:     "$actArgsSetup",
		InArgs:           "$inArgs",
		OutArgs:          "$outArgs",
		FmtJoin:          "$fmtJoin",
		FmtOutputs:       "$fmtOutputs",
		Behaviour:        "$Behaviour",
		TestCaseName:     "$testCaseName",
		DirectFuncInvoke: "$directFuncInvoke",
		Title:            "$title",
		ArrangeSetup:     "$arrangeSetup",
		ExpectedLines:    "$expectedLines",
		CaseItem:         "$caseItem",
		TestCases:        "$testCases",
		VerifyTypeOf:     "$VerifyTypeOf",
		VariablesSetup:   "$variablesSetup",
		workFunc:         "WorkFunc",
		expect:           "Expect",
		inputExpectedVar: "input.Expect",
	}

	templateReplacerFunc = stringutil.
				ReplaceTemplate.
				DirectKeyUsingMapTrim

	Utils = utils{}

	pascalCaseFunc = convertinteranl.Util.String.PascalCase
	camelCaseFunc  = convertinteranl.Util.String.CamelCase
)
