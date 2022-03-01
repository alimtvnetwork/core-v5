package corepubsubinf

import (
	"io"
	"sync"

	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coreinterface"
	"gitlab.com/evatix-go/core/coreinterface/enuminf"
	"gitlab.com/evatix-go/core/coreinterface/errcoreinf"
	"gitlab.com/evatix-go/core/coreinterface/loggerinf"
)

type GenericSubscriber interface {
	OnStart(
		subscribers ...StartFunc,
	) *sync.WaitGroup

	OnComplete(
		subscribers ...CompletionFunc,
	) *sync.WaitGroup

	OnStartComplete(
		startFunc StartFunc,
		completeFunc CompletionFunc,
	) *sync.WaitGroup

	CategoryAnyItem(
		subscribers ...CategoryNameAnyItemSubscriptionFunc,
	) GenericSubscriber

	LogTyperCategoryAnyItem(
		logTyper enuminf.LoggerTyper,
		subscribers ...CategoryNameAnyItemSubscriptionFunc,
	) *sync.WaitGroup

	SingleLogModeler(
		subscriberFunc func(modeler loggerinf.SingleLogModeler),
	) *sync.WaitGroup

	LogTyperSingleLogModeler(
		logTyper enuminf.LoggerTyper,
		subscriberFunc DirectSingleLogModelerSubscribeFunc,
	) *sync.WaitGroup

	JsonResultFunc(
		subscriberFunc JsonResultSubscribeFunc,
	) *sync.WaitGroup

	MessageSubscriberFunc(
		subscriberFunc StringSubscribeFunc,
	) *sync.WaitGroup

	JsonBytesSubscriberFunc(
		subscriberFunc ModelJsonSubscribeFunc,
	) *sync.WaitGroup

	MapAnySubscriberFunc(
		subscriberFunc func(
			category coreinterface.CategoryRevealer,
			mapAny map[string]interface{},
		),
	) *sync.WaitGroup

	Info() GenericSubscriber
	Debug() GenericSubscriber
	Warn() GenericSubscriber
	Error() GenericSubscriber
	Failure() GenericSubscriber

	OnDebug() GenericSubscriber
	OnVerbose() GenericSubscriber

	OnFlag(name, value string) GenericSubscriber
	OnAnyFlag(name string, value interface{}) GenericSubscriber
	OnFlagEnabled(name string) GenericSubscriber
	OnFlagDisabled(name string) GenericSubscriber
	StackSkip(index int) GenericSubscriber

	coreinterface.IsCompletedLockUnlockChecker

	OnMatch(isCondition bool) GenericSubscriber
	OnErr(err error) GenericSubscriber
	OnString(message string) GenericSubscriber
	OnConditionFunc(isSubscribed func() bool) GenericSubscriber

	WaitAll(waitGroups ...*sync.WaitGroup) errcoreinf.BasicErrWrapper
}

type DirectSubscriber interface {
	BasicErrorWrapper(
		basicErrorWrapperSubscribeFunc DirectBasicErrorSubscribeFunc,
	) *sync.WaitGroup

	BaseErrorOrCollectionWrapper(
		subscriberFunc DirectBaseErrorOrCollectionWrapperSubscribeFunc,
	) *sync.WaitGroup

	JsonResultError(
		subscriberFunc DirectJsonResultSubscribeFunc,
	) *sync.WaitGroup

	String(
		messageFunc DirectStringSubscribeFunc,
	) *sync.WaitGroup

	AnyItem(
		subscribedFunc DirectAnyItemSubscribeFunc,
	) *sync.WaitGroup

	Bytes(
		subscribedFunc DirectBytesSubscribeFunc,
	) *sync.WaitGroup

	JsonBytes(
		subscribedFunc DirectModelJsonSubscribeFunc,
	) *sync.WaitGroup

	HashmapSubscriberFunc(
		subscriberFunc HashmapSubscribeFunc,
	) *sync.WaitGroup

	JsonResult(
		subscribedFunc DirectJsonResultSubscribeFunc,
	) *sync.WaitGroup
}

type FilterSubscriber interface {
	FilterText() string

	Filter(
		simpleFunc SimpleCompletionFunc,
	) *sync.WaitGroup

	SkipFilter(
		simpleFunc SimpleCompletionFunc,
	) *sync.WaitGroup

	CategoryFilter(
		simpleFunc SimpleCompletionFunc,
	) *sync.WaitGroup

	DirectSubscriber
}

type GenericPublisher interface {
	Message(
		category coreinterface.CategoryRevealer,
		message string,
	) GenericPublisher

	Boolean(
		category coreinterface.CategoryRevealer,
		isResult bool,
	) GenericPublisher

	SimpleModeler(communicate CommunicateModeler) GenericPublisher

	JsonResult(
		jsonResult *corejson.Result,
	) GenericPublisher

	CategoryMessage(
		categoryName,
		message string,
	) GenericPublisher

	AnyItem(
		categoryName string,
		anyItem interface{},
	) GenericPublisher

	AnyItemDirect(
		anyItem interface{},
	) GenericPublisher

	BytesDirect(
		rawBytes []byte,
	) GenericPublisher

	BooleanDirect(
		isResult bool,
	) GenericPublisher

	Jsoner(
		jsoner corejson.Jsoner,
	) GenericPublisher

	FilterJsoner(
		filterText string,
		jsoner corejson.Jsoner,
	) GenericPublisher

	FilterMessage(
		filterText,
		message string,
	) GenericPublisher

	FilterMetaCompiler(
		filterText,
		title string,
		compiler loggerinf.MetaAttributesCompiler,
	) GenericPublisher

	FilterJsonResult(
		filterText,
		jsonResult *corejson.Result,
	) GenericPublisher

	LogTyperAnyItem(
		logTyper enuminf.LoggerTyper,
		anyItem interface{},
	) GenericPublisher

	Fmt(
		format string,
		v ...interface{},
	) GenericPublisher

	FilterFmt(
		filter,
		format string,
		v ...interface{},
	) GenericPublisher

	LogTyperAnyItemCategory(
		logTyper enuminf.LoggerTyper,
		categoryName string,
		anyItem interface{},
	) GenericPublisher

	All() LogTyperPublisher
	Info() LogTyperPublisher
	Debug() LogTyperPublisher
	Error() LogTyperPublisher
	Warn() LogTyperPublisher
	Failure() LogTyperPublisher

	OnDebug() LogTyperPublisher
	OnVerbose() LogTyperPublisher
	OnMatcherFunc(
		logTyper enuminf.LoggerTyper,
		matcherFunc MatcherFunc,
	) LogTyperPublisher

	Write(p []byte) (n int, err error)
	AsWriter() io.Writer
	AsWriterByLogTyper(logTyper enuminf.LoggerTyper) io.Writer
	AsWriterByLogTyperFilterText(logTyper enuminf.LoggerTyper, filterText string) io.Writer

	LogTyper(
		logTyper enuminf.LoggerTyper,
	) LogTyperPublisher

	CompletePublisher
	coreinterface.IsCompletedLockUnlockChecker
}

type LogTyperPublisher interface {
	LogTyper() enuminf.LoggerTyper
	FilterText() string
	Message(
		category coreinterface.CategoryRevealer,
		message string,
	) LogTyperPublisher

	Boolean(
		category coreinterface.CategoryRevealer,
		isResult bool,
	) GenericPublisher

	MetaStacker() loggerinf.MetaAttributesStacker

	DirectMessage(
		message string,
	) LogTyperPublisher

	JsonResult(
		jsonResult *corejson.Result,
	) LogTyperPublisher

	CategoryMessage(
		categoryName,
		message string,
	) LogTyperPublisher

	SimpleModeler(
		communicate CommunicateModeler,
	) LogTyperPublisher

	AnyItem(
		categoryName string,
		anyItem interface{},
	) LogTyperPublisher

	AnyItemDirect(
		anyItem interface{},
	) LogTyperPublisher

	BytesDirect(
		rawBytes []byte,
	) LogTyperPublisher

	BooleanDirect(
		isResult bool,
	) LogTyperPublisher

	Jsoner(
		jsoner corejson.Jsoner,
	) LogTyperPublisher

	FilterJsoner(
		filterText string,
		jsoner corejson.Jsoner,
	) LogTyperPublisher

	FilterMessage(
		filterText,
		message string,
	) LogTyperPublisher

	FilterMetaCompiler(
		filterText,
		title string,
		compiler loggerinf.MetaAttributesCompiler,
	) LogTyperPublisher

	FilterJsonResult(
		filterText string,
		jsonResult *corejson.Result,
	) LogTyperPublisher

	FilterAnyItem(
		filterText string,
		anyItem interface{},
	) LogTyperPublisher

	FilterCategoryAnyItem(
		filterText, categoryName string,
		anyItem interface{},
	) LogTyperPublisher

	Fmt(
		format string,
		v ...interface{},
	) LogTyperPublisher

	FilterFmt(
		filter,
		format string,
		v ...interface{},
	) LogTyperPublisher

	Hashmap(
		categoryName string,
		hashmap map[string]string,
	) LogTyperPublisher

	HashmapFilter(
		filter, categoryName string,
		hashmap map[string]string,
	) LogTyperPublisher

	DirectHashmap(
		hashmap map[string]string,
	) LogTyperPublisher

	StackSkip(stackSkip int) LogTyperPublisher

	AnErr(err error) LogTyperPublisher
	Err(title string, err error) LogTyperPublisher
	BaseErrOrCollection(baseErrOrCollection errcoreinf.BaseErrorOrCollectionWrapper) LogTyperPublisher
	BasicErrWrapper(basicErrWrapper errcoreinf.BasicErrWrapper) LogTyperPublisher
	BaseErrorWrapperCollectionDefiner(baseErrOrCollection errcoreinf.BaseErrorWrapperCollectionDefiner) LogTyperPublisher
	BaseRawErrCollectionDefiner(baseErrOrCollection errcoreinf.BaseRawErrCollectionDefiner) LogTyperPublisher

	OnMatch(isMatch bool) LogTyperPublisher
	OnMatcherFunc(matcherFunc MatcherFunc) LogTyperPublisher

	Write(p []byte) (n int, err error)
	AsWriter() io.Writer

	CompletePublisher
	Publisher() GenericPublisher
}

type CompletePublisher interface {
	Complete(completionTyper enuminf.CompletionStateTyper)
	CompleteUsingErr(err error)
	CompleteUsingErrWithTitle(title string, err error)
	CompleteUsingBaseErrOrCollection(baseErrOrCollection errcoreinf.BaseErrorOrCollectionWrapper)
	CompleteUsingBasicErrWrapper(basicErrWrapper errcoreinf.BasicErrWrapper)
	CompleteUsingBaseErrorWrapperCollectionDefiner(baseErrOrCollection errcoreinf.BaseErrorWrapperCollectionDefiner)
	CompleteUsingBaseRawErrCollectionDefiner(baseErrOrCollection errcoreinf.BaseRawErrCollectionDefiner)
}
