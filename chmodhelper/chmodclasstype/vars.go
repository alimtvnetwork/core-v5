package chmodclasstype

import (
	"github.com/alimtvnetwork/core/coreimpl/enumimpl"
	"github.com/alimtvnetwork/core/internal/reflectinternal"
)

var (
	Ranges = [...]string{
		Invalid:    "Invalid",
		All:        "All",
		Owner:      "Owner",
		Group:      "Group",
		Other:      "Other",
		OwnerGroup: "OwnerGroup",
		GroupOther: "GroupOther",
		OwnerOther: "OwnerOther",
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
		reflectinternal.TypeName(Invalid),
		Ranges[:])
)
