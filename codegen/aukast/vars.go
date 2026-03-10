package aukast

import (
	"github.com/alimtvnetwork/core/coreutils/stringutil"
	"golang.org/x/tools/go/packages"
)

var (
	globalLoadMode           = packages.NeedSyntax | packages.NeedTypes | packages.NeedTypesInfo
	New                      = newCreator{}
	astUtil                  = utils{}
	replaceToSingleSpaceFunc = stringutil.ReplaceTemplate.ReplaceWhiteSpacesToSingle
	substringEndsFunc        = stringutil.SafeSubstringEnds
)
