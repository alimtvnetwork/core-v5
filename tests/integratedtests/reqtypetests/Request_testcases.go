package reqtypetests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/reqtype"
)

// requestIdentityTestCases
// Expected order: name, isValid, isInvalid
var requestIdentityTestCases = []coretestcases.CaseV1{
	{
		Title:        "Invalid request identity",
		ArrangeInput: reqtype.Invalid,
		ExpectedInput: args.Three[string, string, string]{
			First:  "Invalid", // name
			Second: "false",   // isValid
			Third:  "true",    // isInvalid
		},
	},
	{
		Title:        "Create request identity",
		ArrangeInput: reqtype.Create,
		ExpectedInput: args.Three[string, string, string]{
			First:  "Create", // name
			Second: "true",   // isValid
			Third:  "false",  // isInvalid
		},
	},
	{
		Title:        "Read request identity",
		ArrangeInput: reqtype.Read,
		ExpectedInput: args.Three[string, string, string]{
			First:  "Read",  // name
			Second: "true",  // isValid
			Third:  "false", // isInvalid
		},
	},
	{
		Title:        "Update request identity",
		ArrangeInput: reqtype.Update,
		ExpectedInput: args.Three[string, string, string]{
			First:  "Update", // name
			Second: "true",   // isValid
			Third:  "false",  // isInvalid
		},
	},
	{
		Title:        "Delete request identity",
		ArrangeInput: reqtype.Delete,
		ExpectedInput: args.Three[string, string, string]{
			First:  "Delete", // name
			Second: "true",   // isValid
			Third:  "false",  // isInvalid
		},
	},
	{
		Title:        "Drop request identity",
		ArrangeInput: reqtype.Drop,
		ExpectedInput: args.Three[string, string, string]{
			First:  "Drop",  // name
			Second: "true",  // isValid
			Third:  "false", // isInvalid
		},
	},
}

// requestLogicalGroupTestCases
// Expected order: isCreateLogically, isDropLogically, isCrudOnly, isReadOrEdit
var requestLogicalGroupTestCases = []coretestcases.CaseV1{
	{
		Title:        "Create is logically a create operation",
		ArrangeInput: reqtype.Create,
		ExpectedInput: args.Four[string, string, string, string]{
			First:  "true",  // isCreateLogically
			Second: "false", // isDropLogically
			Third:  "true",  // isCrudOnly
			Fourth: "false", // isReadOrEdit
		},
	},
	{
		Title:        "Read is logically a read/edit operation",
		ArrangeInput: reqtype.Read,
		ExpectedInput: args.Four[string, string, string, string]{
			First:  "false", // isCreateLogically
			Second: "false", // isDropLogically
			Third:  "true",  // isCrudOnly
			Fourth: "true",  // isReadOrEdit
		},
	},
	{
		Title:        "Update is logically a read/edit and CRUD operation",
		ArrangeInput: reqtype.Update,
		ExpectedInput: args.Four[string, string, string, string]{
			First:  "false", // isCreateLogically
			Second: "false", // isDropLogically
			Third:  "true",  // isCrudOnly
			Fourth: "true",  // isReadOrEdit
		},
	},
	{
		Title:        "Delete is logically a CRUD operation",
		ArrangeInput: reqtype.Delete,
		ExpectedInput: args.Four[string, string, string, string]{
			First:  "false", // isCreateLogically
			Second: "false", // isDropLogically
			Third:  "true",  // isCrudOnly
			Fourth: "false", // isReadOrEdit
		},
	},
	{
		Title:        "Drop is logically a drop operation",
		ArrangeInput: reqtype.Drop,
		ExpectedInput: args.Four[string, string, string, string]{
			First:  "false", // isCreateLogically
			Second: "true",  // isDropLogically
			Third:  "false", // isCrudOnly
			Fourth: "false", // isReadOrEdit
		},
	},
	{
		Title:        "CreateOrUpdate is logically a create operation",
		ArrangeInput: reqtype.CreateOrUpdate,
		ExpectedInput: args.Four[string, string, string, string]{
			First:  "true",  // isCreateLogically
			Second: "false", // isDropLogically
			Third:  "false", // isCrudOnly
			Fourth: "false", // isReadOrEdit
		},
	},
	{
		Title:        "Append is not a CRUD operation",
		ArrangeInput: reqtype.Append,
		ExpectedInput: args.Four[string, string, string, string]{
			First:  "false", // isCreateLogically
			Second: "false", // isDropLogically
			Third:  "false", // isCrudOnly
			Fourth: "false", // isReadOrEdit
		},
	},
}

// requestHttpTestCases
// Expected order: isGet, isPost, isPut, isDelete, isPatch
var requestHttpTestCases = []coretestcases.CaseV1{
	{
		Title:        "GetHttp matches only GET",
		ArrangeInput: reqtype.GetHttp,
		ExpectedInput: args.Five[string, string, string, string, string]{
			First:  "true",  // isGet
			Second: "false", // isPost
			Third:  "false", // isPut
			Fourth: "false", // isDelete
			Fifth:  "false", // isPatch
		},
	},
	{
		Title:        "PostHttp matches only POST",
		ArrangeInput: reqtype.PostHttp,
		ExpectedInput: args.Five[string, string, string, string, string]{
			First:  "false", // isGet
			Second: "true",  // isPost
			Third:  "false", // isPut
			Fourth: "false", // isDelete
			Fifth:  "false", // isPatch
		},
	},
	{
		Title:        "PutHttp matches only PUT",
		ArrangeInput: reqtype.PutHttp,
		ExpectedInput: args.Five[string, string, string, string, string]{
			First:  "false", // isGet
			Second: "false", // isPost
			Third:  "true",  // isPut
			Fourth: "false", // isDelete
			Fifth:  "false", // isPatch
		},
	},
	{
		Title:        "DeleteHttp matches only DELETE",
		ArrangeInput: reqtype.DeleteHttp,
		ExpectedInput: args.Five[string, string, string, string, string]{
			First:  "false", // isGet
			Second: "false", // isPost
			Third:  "false", // isPut
			Fourth: "true",  // isDelete
			Fifth:  "false", // isPatch
		},
	},
	{
		Title:        "PatchHttp matches only PATCH",
		ArrangeInput: reqtype.PatchHttp,
		ExpectedInput: args.Five[string, string, string, string, string]{
			First:  "false", // isGet
			Second: "false", // isPost
			Third:  "false", // isPut
			Fourth: "false", // isDelete
			Fifth:  "true",  // isPatch
		},
	},
	{
		Title:        "Create is not an HTTP method",
		ArrangeInput: reqtype.Create,
		ExpectedInput: args.Five[string, string, string, string, string]{
			First:  "false", // isGet
			Second: "false", // isPost
			Third:  "false", // isPut
			Fourth: "false", // isDelete
			Fifth:  "false", // isPatch
		},
	},
}
