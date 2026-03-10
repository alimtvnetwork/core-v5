package chmodhelpertestwrappers

import (
	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/alimtvnetwork/core/chmodhelper/chmodclasstype"
)

var SingleRwxTestCases = []chmodhelper.SingleRwx{
	{
		Rwx:       "r-x",
		ClassType: chmodclasstype.Other,
	},
	{
		Rwx:       "---",
		ClassType: chmodclasstype.Other,
	},
	{
		Rwx:       "--x",
		ClassType: chmodclasstype.Other,
	},
	{
		Rwx:       "r-x",
		ClassType: chmodclasstype.Other,
	},
}
