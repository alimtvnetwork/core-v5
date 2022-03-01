package corepubsubinf

import (
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coreinterface"
	"gitlab.com/evatix-go/core/coreinterface/errcoreinf"
	"gitlab.com/evatix-go/core/coreinterface/loggerinf"
)

type (
	SubscribeFunc func(
		subscriptionRecorder SubscriptionRecorder,
	)

	DirectSingleLogModelerSubscribeFunc func(modeler loggerinf.SingleLogModeler)

	JsonResultSubscribeFunc func(
		categoryRevealer coreinterface.CategoryRevealer,
		jsonResult *corejson.Result,
	)

	BooleanSubscribeFunc func(
		categoryRevealer coreinterface.CategoryRevealer,
		isResult bool,
	)

	HashmapSubscribeFunc func(
		categoryRevealer coreinterface.CategoryRevealer,
		hashmap map[string]string,
	)
	BytesSubscribeFunc func(
		categoryRevealer coreinterface.CategoryRevealer,
		rawBytes []byte,
	)
	ModelJsonSubscribeFunc func(
		categoryRevealer coreinterface.CategoryRevealer,
		modelJson []byte,
	)
	AnyItemSubscribeFunc func(
		categoryRevealer coreinterface.CategoryRevealer,
		anyItem interface{},
	)

	StringSubscribeFunc func(
		categoryRevealer coreinterface.CategoryRevealer,
		message string,
	)

	DirectBytesSubscribeFunc func(
		rawBytes []byte,
	)

	DirectModelJsonSubscribeFunc func(
		modelJson []byte,
	)

	DirectAnyItemSubscribeFunc func(
		anyItem interface{},
	)

	DirectJsonResultSubscribeFunc func(
		jsonResult *corejson.Result,
	)

	DirectStringSubscribeFunc func(
		message string,
	)

	DirectBasicErrorSubscribeFunc func(
		basicErrorWrapper errcoreinf.BasicErrWrapper,
	)

	DirectBaseErrorOrCollectionWrapperSubscribeFunc func(
		basicErrorWrapper errcoreinf.BaseErrorOrCollectionWrapper,
	)

	DirectBooleanSubscribeFunc func(
		isResult bool,
	)

	SimpleSubscribeFunc                 func(communicate CommunicateModeler)
	FilterStringSubscribeFunc           func(communicate CommunicateModeler, currentStringValue string)
	LogSubscriberFunc                   func(logRecord BaseLogModeler)
	StartFunc                           func(subscriptionRecorder SubscriptionRecorder)
	CompletionFunc                      func(subscriptionRecorder SubscriptionRecorder)
	StartEndSubscriptionFunc            func(isStart bool, subscriptionRecorder SubscriptionRecorder)
	SimpleCompletionFunc                func(communicate CommunicateModeler)
	CategoryNameAnyItemSubscriptionFunc func(categoryName string, anyItem interface{})
	MatcherFunc                         func() (isMatch bool)
)
