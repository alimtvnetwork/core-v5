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
		Title:        "Invalid request identity",
		ArrangeInput: reqtype.Invalid,
		ExpectedInput: args.Map{
			"name":      "Invalid",
			"isValid":   "false",
			"isInvalid": "true",
		},
	},
	{
		Title:        "Create request identity",
		ArrangeInput: reqtype.Create,
		ExpectedInput: args.Map{
			"name":      "CreateUsingAliasMap",
			"isValid":   "true",
			"isInvalid": "false",
		},
	},
	{
		Title:        "Read request identity",
		ArrangeInput: reqtype.Read,
		ExpectedInput: args.Map{
			"name":      "Read",
			"isValid":   "true",
			"isInvalid": "false",
		},
	},
	{
		Title:        "Update request identity",
		ArrangeInput: reqtype.Update,
		ExpectedInput: args.Map{
			"name":      "Update",
			"isValid":   "true",
			"isInvalid": "false",
		},
	},
	{
		Title:        "Delete request identity",
		ArrangeInput: reqtype.Delete,
		ExpectedInput: args.Map{
			"name":      "Delete",
			"isValid":   "true",
			"isInvalid": "false",
		},
	},
	{
		Title:        "Drop request identity",
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
		Title:        "Create is logically a create operation",
		ArrangeInput: reqtype.Create,
		ExpectedInput: args.Map{
			"isCreateLogically": "true",
			"isDropLogically":   "false",
			"isCrudOnly":        "true",
			"isReadOrEdit":      "false",
		},
	},
	{
		Title:        "Read is logically a read/edit operation",
		ArrangeInput: reqtype.Read,
		ExpectedInput: args.Map{
			"isCreateLogically": "false",
			"isDropLogically":   "false",
			"isCrudOnly":        "true",
			"isReadOrEdit":      "true",
		},
	},
	{
		Title:        "Update is logically a read/edit and CRUD operation",
		ArrangeInput: reqtype.Update,
		ExpectedInput: args.Map{
			"isCreateLogically": "false",
			"isDropLogically":   "false",
			"isCrudOnly":        "true",
			"isReadOrEdit":      "true",
		},
	},
	{
		Title:        "Delete is logically a CRUD operation",
		ArrangeInput: reqtype.Delete,
		ExpectedInput: args.Map{
			"isCreateLogically": "false",
			"isDropLogically":   "false",
			"isCrudOnly":        "true",
			"isReadOrEdit":      "false",
		},
	},
	{
		Title:        "Drop is logically a drop operation",
		ArrangeInput: reqtype.Drop,
		ExpectedInput: args.Map{
			"isCreateLogically": "false",
			"isDropLogically":   "true",
			"isCrudOnly":        "true",
			"isReadOrEdit":      "false",
		},
	},
	{
		Title:        "CreateOrUpdate is logically a create operation",
		ArrangeInput: reqtype.CreateOrUpdate,
		ExpectedInput: args.Map{
			"isCreateLogically": "true",
			"isDropLogically":   "false",
			"isCrudOnly":        "false",
			"isReadOrEdit":      "false",
		},
	},
	{
		Title:        "Append is not a CRUD operation",
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
		Title:        "GetHttp matches only GET",
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
		Title:        "PostHttp matches only POST",
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
		Title:        "PutHttp matches only PUT",
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
		Title:        "DeleteHttp matches only DELETE",
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
		Title:        "PatchHttp matches only PATCH",
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
		Title:        "Create is not an HTTP method",
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
