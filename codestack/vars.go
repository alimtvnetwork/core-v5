package codestack

import "gitlab.com/auk-go/core/internal/reflectinternal"

var (
	NameOf   = currentNameOf{}
	New      = newCreator{}
	StacksTo = stacksTo{}

	getFuncEverything = reflectinternal.GetFunc.All
)
