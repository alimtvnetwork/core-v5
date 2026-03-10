package loggerinf

import (
	"github.com/alimtvnetwork/core/coreinterface/enuminf"
	"github.com/alimtvnetwork/core/coreinterface/pathextendinf"
	"github.com/alimtvnetwork/core/internal/internalinterface"
)

type AttrPersistentLogger interface {
	internalinterface.IdStringerWithNamer
	LogPathInfo() pathextendinf.PathInfoer
	IsRotating() bool
	IsDbLogger() bool
	IsFileLogger() bool

	DynamicConfig() any
	ConfigReflectSetTo(toPointer any) error

	// PersistentLoggerTyper
	//
	//  Which type of persistent logger
	PersistentLoggerTyper() enuminf.BasicEnumer
}
