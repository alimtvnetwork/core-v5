package codegen

import (
	"gitlab.com/auk-go/core/internal/convertinteranl"
)

var (
	NewCodeOutput = newCodeOutputCreator{}

	unitTestVars = unitVariables{
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
		inputExpectedVar: "input.Expect",
	}

	pascalCaseFunc = convertinteranl.Util.String.PascalCase
	camelCaseFunc  = convertinteranl.Util.String.CamelCase
)
