package corefuncs

import (
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corepayload"
	"github.com/alimtvnetwork/core/coreinterface/errcoreinf"
	"github.com/alimtvnetwork/core/coreinterface/serializerinf"
)

type (
	ExecFunc                    func()
	StringerActionFunc          func() (result string)
	StringerWithErrorActionFunc func() (result string, err error)
	ActionFunc                  func()
	IsApplyFunc                 func() (isSuccess bool)
	IsBooleanFunc               func() bool
	InOutFunc                   func(input any) (output any)
	InOutErrFunc                func(input any) (output any, err error)
	SerializeOutputFunc         func(input any) (serializedBytes []byte, err error)
	SerializerVoidFunc          func() (serializedBytes []byte, err error)
	InActionReturnsErrFunc      func(input any) (err error)
	NamedActionFunc             func(name string)
	ActionReturnsErrorFunc      func() error
	IsSuccessFunc               func() (isSuccess bool)
	IsFailureFunc               func() (isFailed bool)
	// ResultDelegatingFunc
	//
	// resultDelegatedTo can be unmarshal or marshal or reflect set
	ResultDelegatingFunc           func(resultDelegatedTo any) error
	NextReturnErrWrapperFunc       func(nextAction ActionReturnsErrorFunc) error
	NextVoidActionFunc             func(nextAction ExecFunc)
	PayloadProcessorFunc           func(payloads []byte) (err error)
	PayloadToBasicErrProcessorFunc func(payloads []byte) (basicError errcoreinf.BasicErrWrapper)
	SimpleBytesResultProcessorFunc func(simpleBytes serializerinf.SimpleBytesResulter) (basicError errcoreinf.BasicErrWrapper)
	ErrorToBasicError              func(
		errorTyper errcoreinf.BaseErrorTyper,
		err error,
	) (basicError errcoreinf.BasicErrWrapper)
	BaseJsonResultProcessorFunc          func(baseJsonResulter serializerinf.BaseJsonResulter) (basicError errcoreinf.BasicErrWrapper)
	JsonResulterProcessorFunc            func(result serializerinf.JsonResulter) (basicError errcoreinf.BasicErrWrapper)
	JsonResultProcessorFunc              func(result *corejson.Result) (basicError errcoreinf.BasicErrWrapper)
	PayloadWrapperProcessorFunc          func(payloadWrapper *corepayload.PayloadWrapper) (basicError errcoreinf.BasicErrWrapper)
	MultiPayloadsProcessorFunc           func(multiPayloads ...[]byte) (err error)
	BytesCollectionPayloadsProcessorFunc func(collectionOfBytes *corejson.BytesCollection) (err error)
	PayloadToPayloadWrapperFunc          func(payloads []byte) (payloadWrapper *corepayload.PayloadWrapper, err error)
	NextPayloadProcessorLinkerFunc       func(nextLinkerFunc PayloadProcessorFunc) error
)
