package codestack

import "gitlab.com/auk-go/core/internal/reflectinternal"

var (
	NameOf            = currentNameOf{}
	getFuncEverything = reflectinternal.GetFunc.All
)
