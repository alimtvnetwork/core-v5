package reqtypetests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/reqtype"
)

// requestIdentityTestCases
// Expected: name, isValid, isInvalid
var requestIdentityTestCases = []coretestcases.CaseV1{
	{
		Title:        "Request.Name returns 'Invalid' and isValid false -- Invalid type",
		ArrangeInput: reqtype.Invalid,
		ExpectedInput: args.Map{
			"name":      "Invalid",
			"isValid":   "false",
			"isInvalid": "true",
		},
	},
	{
		Title:        "Request.Name returns 'CreateUsingAliasMap' and isValid true -- Create type",
		ArrangeInput: reqtype.Create,
		ExpectedInput: args.Map{
			"name":      "CreateUsingAliasMap",
			"isValid":   "true",
			"isInvalid": "false",
		},
	},
	{
		Title:        "Request.Name returns 'Read' and isValid true -- Read type",
		ArrangeInput: reqtype.Read,
		ExpectedInput: args.Map{
			"name":      "Read",
			"isValid":   "true",
			"isInvalid": "false",
		},
	},
	{
		Title:        "Request.Name returns 'Update' and isValid true -- Update type",
		ArrangeInput: reqtype.Update,
		ExpectedInput: args.Map{
			"name":      "Update",
			"isValid":   "true",
			"isInvalid": "false",
		},
	},
	{
		Title:        "Request.Name returns 'Delete' and isValid true -- Delete type",
		ArrangeInput: reqtype.Delete,
		ExpectedInput: args.Map{
			"name":      "Delete",
			"isValid":   "true",
			"isInvalid": "false",
		},
	},
	{
		Title:        "Request.Name returns 'Drop' and isValid true -- Drop type",
		ArrangeInput: reqtype.Drop,
		ExpectedInput: args.Map{
			"name":      "Drop",
			"isValid":   "true",
			"isInvalid": "false",
		},
	},
}

// requestLogicalGroupTestCases
// Expected: isCreateLogically, isDropLogically, isCrudOnly, isReadOrEdit
var requestLogicalGroupTestCases = []coretestcases.CaseV1{
	{
		Title:        "Request returns isCreateLogically true and isCrudOnly true -- Create type",
		ArrangeInput: reqtype.Create,
		ExpectedInput: args.Map{
			"isCreateLogically": "true",
			"isDropLogically":   "false",
			"isCrudOnly":        "true",
			"isReadOrEdit":      "false",
		},
	},
	{
		Title:        "Request returns isReadOrEdit true and isCrudOnly true -- Read type",
		ArrangeInput: reqtype.Read,
		ExpectedInput: args.Map{
			"isCreateLogically": "false",
			"isDropLogically":   "false",
			"isCrudOnly":        "true",
			"isReadOrEdit":      "true",
		},
	},
	{
		Title:        "Request returns isReadOrEdit true and isCrudOnly true -- Update type",
		ArrangeInput: reqtype.Update,
		ExpectedInput: args.Map{
			"isCreateLogically": "false",
			"isDropLogically":   "false",
			"isCrudOnly":        "true",
			"isReadOrEdit":      "true",
		},
	},
	{
		Title:        "Request returns isCrudOnly true -- Delete type",
		ArrangeInput: reqtype.Delete,
		ExpectedInput: args.Map{
			"isCreateLogically": "false",
			"isDropLogically":   "false",
			"isCrudOnly":        "true",
			"isReadOrEdit":      "false",
		},
	},
	{
		Title:        "Request returns isDropLogically true and isCrudOnly true -- Drop type",
		ArrangeInput: reqtype.Drop,
		ExpectedInput: args.Map{
			"isCreateLogically": "false",
			"isDropLogically":   "true",
			"isCrudOnly":        "true",
			"isReadOrEdit":      "false",
		},
	},
	{
		Title:        "Request returns isCreateLogically true and isReadOrEdit true -- CreateOrUpdate type",
		ArrangeInput: reqtype.CreateOrUpdate,
		ExpectedInput: args.Map{
			"isCreateLogically": "true",
			"isDropLogically":   "false",
			"isCrudOnly":        "true",
			"isReadOrEdit":      "true",
		},
	},
	{
		Title:        "Request returns isCrudOnly false -- Append type",
		ArrangeInput: reqtype.Append,
		ExpectedInput: args.Map{
			"isCreateLogically": "false",
			"isDropLogically":   "false",
			"isCrudOnly":        "false",
			"isReadOrEdit":      "false",
		},
	},
}

// requestHttpTestCases
// Expected: isGet, isPost, isPut, isDelete, isPatch
var requestHttpTestCases = []coretestcases.CaseV1{
	{
		Title:        "Request.IsGetHttp returns true -- GetHttp type",
		ArrangeInput: reqtype.GetHttp,
		ExpectedInput: args.Map{
			"isGet":    "true",
			"isPost":   "false",
			"isPut":    "false",
			"isDelete": "false",
			"isPatch":  "false",
		},
	},
	{
		Title:        "Request.IsPostHttp returns true -- PostHttp type",
		ArrangeInput: reqtype.PostHttp,
		ExpectedInput: args.Map{
			"isGet":    "false",
			"isPost":   "true",
			"isPut":    "false",
			"isDelete": "false",
			"isPatch":  "false",
		},
	},
	{
		Title:        "Request.IsPutHttp returns true -- PutHttp type",
		ArrangeInput: reqtype.PutHttp,
		ExpectedInput: args.Map{
			"isGet":    "false",
			"isPost":   "false",
			"isPut":    "true",
			"isDelete": "false",
			"isPatch":  "false",
		},
	},
	{
		Title:        "Request.IsDeleteHttp returns true -- DeleteHttp type",
		ArrangeInput: reqtype.DeleteHttp,
		ExpectedInput: args.Map{
			"isGet":    "false",
			"isPost":   "false",
			"isPut":    "false",
			"isDelete": "true",
			"isPatch":  "false",
		},
	},
	{
		Title:        "Request.IsPatchHttp returns true -- PatchHttp type",
		ArrangeInput: reqtype.PatchHttp,
		ExpectedInput: args.Map{
			"isGet":    "false",
			"isPost":   "false",
			"isPut":    "false",
			"isDelete": "false",
			"isPatch":  "true",
		},
	},
	{
		Title:        "Request returns all HTTP false -- Create (non-HTTP) type",
		ArrangeInput: reqtype.Create,
		ExpectedInput: args.Map{
			"isGet":    "false",
			"isPost":   "false",
			"isPut":    "false",
			"isDelete": "false",
			"isPatch":  "false",
		},
	},
}
