package reqtypetests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/reqtype"
)

// requestIdentityTestCases
// Expected order: name, isValid, isInvalid
var requestIdentityTestCases = []coretestcases.CaseV1{
	{
		Title:         "Invalid request identity",
		ArrangeInput:  reqtype.Invalid,
		ExpectedInput: []string{"Invalid", "false", "true"},
	},
	{
		Title:         "Create request identity",
		ArrangeInput:  reqtype.Create,
		ExpectedInput: []string{"Create", "true", "false"},
	},
	{
		Title:         "Read request identity",
		ArrangeInput:  reqtype.Read,
		ExpectedInput: []string{"Read", "true", "false"},
	},
	{
		Title:         "Update request identity",
		ArrangeInput:  reqtype.Update,
		ExpectedInput: []string{"Update", "true", "false"},
	},
	{
		Title:         "Delete request identity",
		ArrangeInput:  reqtype.Delete,
		ExpectedInput: []string{"Delete", "true", "false"},
	},
	{
		Title:         "Drop request identity",
		ArrangeInput:  reqtype.Drop,
		ExpectedInput: []string{"Drop", "true", "false"},
	},
}

// requestLogicalGroupTestCases
// Expected order: isCreateLogically, isDropLogically, isCrudOnly, isReadOrEdit
var requestLogicalGroupTestCases = []coretestcases.CaseV1{
	{
		Title:         "Create is logically a create operation",
		ArrangeInput:  reqtype.Create,
		ExpectedInput: []string{"true", "false", "true", "false"},
	},
	{
		Title:         "Read is logically a read/edit operation",
		ArrangeInput:  reqtype.Read,
		ExpectedInput: []string{"false", "false", "true", "true"},
	},
	{
		Title:         "Update is logically a read/edit and CRUD operation",
		ArrangeInput:  reqtype.Update,
		ExpectedInput: []string{"false", "false", "true", "true"},
	},
	{
		Title:         "Delete is logically a CRUD operation",
		ArrangeInput:  reqtype.Delete,
		ExpectedInput: []string{"false", "false", "true", "false"},
	},
	{
		Title:         "Drop is logically a drop operation",
		ArrangeInput:  reqtype.Drop,
		ExpectedInput: []string{"false", "true", "false", "false"},
	},
	{
		Title:         "CreateOrUpdate is logically a create operation",
		ArrangeInput:  reqtype.CreateOrUpdate,
		ExpectedInput: []string{"true", "false", "false", "false"},
	},
	{
		Title:         "Append is not a CRUD operation",
		ArrangeInput:  reqtype.Append,
		ExpectedInput: []string{"false", "false", "false", "false"},
	},
}

// requestHttpTestCases
// Expected order: isGet, isPost, isPut, isDelete, isPatch
var requestHttpTestCases = []coretestcases.CaseV1{
	{
		Title:         "GetHttp matches only GET",
		ArrangeInput:  reqtype.GetHttp,
		ExpectedInput: []string{"true", "false", "false", "false", "false"},
	},
	{
		Title:         "PostHttp matches only POST",
		ArrangeInput:  reqtype.PostHttp,
		ExpectedInput: []string{"false", "true", "false", "false", "false"},
	},
	{
		Title:         "PutHttp matches only PUT",
		ArrangeInput:  reqtype.PutHttp,
		ExpectedInput: []string{"false", "false", "true", "false", "false"},
	},
	{
		Title:         "DeleteHttp matches only DELETE",
		ArrangeInput:  reqtype.DeleteHttp,
		ExpectedInput: []string{"false", "false", "false", "true", "false"},
	},
	{
		Title:         "PatchHttp matches only PATCH",
		ArrangeInput:  reqtype.PatchHttp,
		ExpectedInput: []string{"false", "false", "false", "false", "true"},
	},
	{
		Title:         "Create is not an HTTP method",
		ArrangeInput:  reqtype.Create,
		ExpectedInput: []string{"false", "false", "false", "false", "false"},
	},
}
