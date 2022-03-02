package coreinterface

type BasicSlicerContractsBinder interface {
	BasicSlicer
	AsBasicSliceContractsBinder() BasicSlicerContractsBinder
}

type StandardSlicerContractsBinder interface {
	StandardSlicer
	AsStandardSlicerContractsBinder() StandardSlicerContractsBinder
}

type DynamicLinqWithPagingContractsBinder interface {
	DynamicLinqWithPaging
	AsDynamicLinqWithPagingContractsBinder() DynamicLinqWithPagingContractsBinder
}

type DynamicLinqContractsBinder interface {
	DynamicLinq
	AsDynamicLinqContractsBinder() DynamicLinqContractsBinder
}

type SimpleValidInvalidCheckerContractsBinder interface {
	SimpleValidInvalidChecker
	AsSimpleValidInvalidChecker() SimpleValidInvalidChecker
}

type JsonBytesStringerContractsBinder interface {
	JsonByter
	JsonCombineStringer
	AsJsonBytesStringerContractsBinder() JsonBytesStringerContractsBinder
}

type CountStateTrackerBinder interface {
	CountStateTracker
	AsCountStateTrackerBinder() CountStateTrackerBinder
}

type KeyValueStringDefinerBinder interface {
	KeyValueStringDefiner
	AsKeyValueStringDefinerBinder() KeyValueStringDefinerBinder
}

type KeyAnyValueDefinerBinder interface {
	KeyAnyValueDefiner
	AsKeyAnyValueDefinerBinder() KeyAnyValueDefinerBinder
}

type KeyStringValuesCollectionDefinerBinder interface {
	KeyStringValuesCollectionDefiner
	AsKeyStringValuesCollectionDefinerBinder() KeyStringValuesCollectionDefinerBinder
}

type KeyAnyValuesCollectionDefinerBinder interface {
	KeyAnyValuesCollectionDefiner
	AsKeyAnyValuesCollectionDefinerBinder() KeyAnyValuesCollectionDefinerBinder
}

type AttributesBinder interface {
	Length() int
	HasAnyItem() bool
	Payloads() []byte
	Capacity() int
	AnyKeyValMap() map[string]interface{}
	Hashmap() map[string]string
	CompiledError() error
	HasError() bool
	IsSafeValid() bool
	IsInvalid() bool
	IsValid() bool
	MustBeEmptyError()
	HandleErr()
	DeserializeDynamicPayloads(
		unmarshalToPointer interface{},
	) error
	IsEmptyError() bool
	DynamicBytesLength() int
	StringKeyValuePairsLength() int
	AnyKeyValuePairsLength() int
	IsEmpty() bool
	HasItems() bool
	IsErrorMessageEmpty() bool
	AuthType() string
	ResourceName() string
	HasStringKeyValuePairs() bool
	HasAnyKeyValuePairs() bool
	HasDynamicPayloads() bool
	DynamicPayloadsDeserialize(
		unmarshallingPointer interface{},
	) error
	DynamicPayloadsDeserializeMust(
		unmarshallingPointer interface{},
	)
	AddOrUpdateString(
		key, value string,
	) (isNewlyAdded bool)
	AddOrUpdateAnyItem(
		key string,
		anyItem interface{},
	) (isNewlyAdded bool)
	String() string
	JsonModelAny() interface{}
	AttachOrAppendError(
		err error,
	) AttributesBinder
	Clear()
	Dispose()
}
