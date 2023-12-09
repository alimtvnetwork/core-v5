package codegen

import (
	"gitlab.com/auk-go/core/coreindexes"
	"gitlab.com/auk-go/core/internal/convertinteranl"
)

var (
	NewCodeOutput  = newCodeOutputCreator{}
	indexByNameMap = map[int]string{
		coreindexes.First:   "First",
		coreindexes.Second:  "Second",
		coreindexes.Third:   "Third",
		coreindexes.Fourth:  "Fourth",
		coreindexes.Fifth:   "Fifth",
		coreindexes.Sixth:   "Sixth",
		coreindexes.Seventh: "Seventh",
	}

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
		inputExpectedVar: "input.Expect",
	}

	pascalCaseFunc = convertinteranl.Util.String.PascalCase
	camelCaseFunc  = convertinteranl.Util.String.CamelCase
)
