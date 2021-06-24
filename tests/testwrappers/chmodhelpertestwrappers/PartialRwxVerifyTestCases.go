package chmodhelpertestwrappers

type PartialRwxVerify struct {
	Header                                string
	PartialRwxInput1, FullRwxVerifyInput2 string
	IsMatchesExpectation                  bool
}

var PartialRwxVerifyTestCases = []PartialRwxVerify{
	{
		Header:               "Same input returns true.",
		PartialRwxInput1:     "-rwx-*-r*x",
		FullRwxVerifyInput2:  "-rwx-*-r*x",
		IsMatchesExpectation: true,
	},
	{
		Header: "Same [-rwx---r*x] comparing " +
			"with [-rwx-*-r*x] returns false.",
		PartialRwxInput1:     "-rwx---r*x",
		FullRwxVerifyInput2:  "-rwx-*-r*x",
		IsMatchesExpectation: false,
	},
	{
		Header: "Same [-rwx-*-r*x] comparing with " +
			"[-rwx-w-r*x] returns true.",
		PartialRwxInput1:     "-rwx-*-r*x",
		FullRwxVerifyInput2:  "-rwx-w-r*x",
		IsMatchesExpectation: true,
	},
	{
		Header: "Same [-rwx-*-] or [-rwx-*-***] (not givens ones are wildcard) " +
			"comparing with [-rwx-w--*x] returns true.",
		PartialRwxInput1:     "-rwx-*-",
		FullRwxVerifyInput2:  "-rwx-w--*x",
		IsMatchesExpectation: true,
	},
	{
		Header: "Same [-rwxr*-] or [-rwxr*-***] (not givens ones are wildcard) " +
			"comparing with [-rwx-w--*x] returns false.",
		PartialRwxInput1:     "-rwxr*-",
		FullRwxVerifyInput2:  "-rwx-w--*x",
		IsMatchesExpectation: false,
	},
}
