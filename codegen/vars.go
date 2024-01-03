package codegen

import (
	"gitlab.com/auk-go/core/codegen/codegentype"
	"gitlab.com/auk-go/core/coreutils/stringutil"
	"gitlab.com/auk-go/core/internal/convertinteranl"
)

var (
	New = newCreator{}

	defaultPackages = []string{
		"\"testing\"",
		"\"gitlab.com/auk-go/core/coredata/corestr\"",
		"\"gitlab.com/auk-go/core/coretests\"",
		"\"gitlab.com/auk-go/core/coretests/args\"",
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
		ToCompiled:       "$toCompiled",
		TypeName:         "$typename",
		workFunc:         "WorkFunc",
		expect:           "Expect",
		inputExpected:    "input.Expect",
		inputPrefix:      "input",
	}

	templateReplacerFunc = stringutil.
				ReplaceTemplate.
				DirectKeyUsingMapTrim

	Utils = utils{}

	functionTemplatesMap = [...]string{
		codegentype.Simple:           funcTemplate,
		codegentype.MultipleArranges: loopFuncTemplate,
	}

	pascalCaseFunc = convertinteranl.Util.String.PascalCase
	camelCaseFunc  = convertinteranl.Util.String.CamelCase
)
