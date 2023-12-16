package aukast

import "golang.org/x/tools/go/packages"

var (
	globalLoadMode = packages.NeedSyntax | packages.NeedTypes | packages.NeedTypesInfo
	New            = newCreator{}
	astUtil        = utils{}
)
