package loggerinf

import (
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coreinterface"
	"github.com/alimtvnetwork/core/coreinterface/enuminf"
	"github.com/alimtvnetwork/core/coreinterface/errcoreinf"
	"github.com/alimtvnetwork/core/coreinterface/serializerinf"
	"github.com/alimtvnetwork/core/internal/internalinterface"
)

type SingleLogModeler interface {
	internalinterface.IdentifierGetter
	PersistentIdGetter
	ParentPersistentIdGetter

	HasParentChecker
	HasModelChecker
	hasErrorChecker

	coreinterface.CategoryRevealer
	coreinterface.CategoryRevealerGetter

	FilterTyper() enuminf.BasicEnumer
	LevelTyper() enuminf.LogLevelTyper
	LogTyper() enuminf.LoggerTyper
	BasicErrorTyper() errcoreinf.BasicErrorTyper
	ModelTyper() enuminf.BasicEnumer
	EntityTypeName() string

	ModelBytesGetter

	LogMessageGetter
	CompiledAttributesGetter
	CallerGetter

	SpecificValuerGetter
	ErrorAsBasicErrWrapperGetter
	ReflectSetter
	serializerinf.Deserializer

	corejson.JsonContractsBinder
}
