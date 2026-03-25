package pathextendinf

import "github.com/alimtvnetwork/core/internal/internalinterface/internalpathextender"

type PathInfoer interface {
	IsPathChecker
	internalpathextender.PathInfoer
}
