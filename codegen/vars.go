package codegen

import (
	"github.com/alimtvnetwork/core/codegen/codegentype"
	"github.com/alimtvnetwork/core/coreutils/stringutil"
	"github.com/alimtvnetwork/core/internal/convertinternal"
)

var (
	New = newCreator{}

	defaultPackages = []string{
		"\"testing\"",
		"\"github.com/alimtvnetwork/core/coredata/corestr\"",
		"\"github.com/alimtvnetwork/core/coretests\"",
		"\"github.com/alimtvnetwork/core/coretests/args\"",
		"\"github.com/alimtvnetwork/core/converters\"",
	}

	vars = unitVariables{
		PackageName:        "$packageName",
		NewPackages:        "$newPackages",
		FuncName:           "$FuncName",
		ArrangeType:        "$ArrangeType",
		LinesPossible:      "$linesPossible",
		ActArgsSetup:       "$actArgsSetup",
		InArgs:             "$inArgs",
		OutArgs:            "$outArgs",
		FmtJoin:            "$fmtJoin",
		FmtOutputs:         "$fmtOutputs",
		Behaviour:          "$Behaviour",
		TestCaseName:       "$testCaseName",
		DirectFuncInvoke:   "$directFuncInvoke",
		Title:              "$title",
		ArrangeSetup:       "$arrangeSetup",
		ExpectedLines:      "$expectedLines",
		CaseItem:           "$caseItem",
		TestCases:          "$testCases",
		VerifyTypeOf:       "$VerifyTypeOf",
		VariablesSetup:     "$variablesSetup",
		ToCompiled:         "$toCompiled",
		TypeName:           "$typename",
		AllInArgsSpread:    "$allInArgsSpread",
		AllOutArgsSpread:   "$allOutArgsSpread",
		workFunc:           "WorkFunc",
		expect:             "Expect",
		inputExpected:      "input.Expect",
		inputPrefix:        "input",
		allInArgsCompiled:  "allInArgsCompiled",
		allOutArgsCompiled: "allOutArgsCompiled",
	}

	templateReplacerFunc = stringutil.
				ReplaceTemplate.
				DirectKeyUsingMapTrim

	Utils = utils{}

	functionTemplatesMap = [...]string{
		codegentype.Simple:           funcTemplate,
		codegentype.MultipleArranges: loopFuncTemplate,
	}

	pascalCaseFunc = convertinternal.Util.String.PascalCase
	camelCaseFunc  = convertinternal.Util.String.CamelCase
)
