package pathextendinf

import "github.com/alimtvnetwork/core/coreinterface/errcoreinf"

type ActionExecutor interface {
	HasAnyAction() bool
	IsEmptyActions() bool
	Exec() errcoreinf.BasicErrWrapper
	ExecAll() errcoreinf.BaseErrorWrapperCollectionDefiner
}
