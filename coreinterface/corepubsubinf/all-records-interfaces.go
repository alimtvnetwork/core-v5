package corepubsubinf

import (
	"github.com/alimtvnetwork/core/coreinterface"
	"github.com/alimtvnetwork/core/coreinterface/enuminf"
	"github.com/alimtvnetwork/core/coreinterface/errcoreinf"
	"github.com/alimtvnetwork/core/coreinterface/loggerinf"
	"github.com/alimtvnetwork/core/coreinterface/pathextendinf"
	"github.com/alimtvnetwork/core/internal/internalinterface"
)

type IdAsStringer interface {
	internalinterface.IdAsStringer
}

type SubscriptionMainRecorder interface {
	IdAsStringer
	TableName() string

	IsEmpty() bool

	pathextendinf.PathExtenderGetter

	HasRecordError() bool
	SetRecordError() bool
	IsArchivedRecord() bool
	IsCompletedRecord() bool
	IsMigratedRecord() bool
	CompletionTyper() enuminf.CompletionStateTyper

	// DefaultDelayMillis
	//
	//  Chmod delay in milliseconds
	DefaultDelayMillis() int
}

type BaseLogModeler interface {
	enuminf.LoggerTyperGetter
	enuminf.EventTyperGetter
	errcoreinf.BasicErrorTyperGetter
	errcoreinf.ErrorStringGetter
	coreinterface.StackTracesBytesGetter
	coreinterface.JsonErrorBytesGetter
	IsEmpty() bool
	LogMessage() string
}

type CommunicateModeler interface {
	BaseLogModeler() BaseLogModeler
	PersistentId() uint
	IdAsStringer
	TableName() string

	SetCallerFileLineUsingStackSkip(
		stackSkip int,
	)

	loggerinf.SingleLogModeler
}

type SubscriptionRecorder interface {
	MainRecord() SubscriptionMainRecorder
	CommunicateRecord() CommunicateModeler
}
