package codestack

import (
	"github.com/alimtvnetwork/core/internal/reflectinternal"
)

var (
	NameOf   = currentNameOf{}
	New      = newCreator{}
	StacksTo = stacksTo{}
	File     = fileGetter{}
	Dir      = dirGetter{}

	getFuncEverything = reflectinternal.GetFunc.All
)
