package loggerinf

import (
	"fmt"

	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/coreinterface"
	"gitlab.com/auk-go/core/coreinterface/entityinf"
	"gitlab.com/auk-go/core/coreinterface/enuminf"
	"gitlab.com/auk-go/core/coreinterface/errcoreinf"
)

type MetaAttributesWithoutTileStacker interface {
	enuminf.LoggerTyperGetter

	Title() string
	IsSilent() bool

	On(isLog bool) MetaAttributesWithoutTileStacker
	Attr(attr string) MetaAttributesWithoutTileStacker
	Str(val string) MetaAttributesWithoutTileStacker
	Strings(stringItems ...string) MetaAttributesWithoutTileStacker
	StandardSlicer(standardSlice coreinterface.StandardSlicer) MetaAttributesWithoutTileStacker
	Stringer(stringer fmt.Stringer) MetaAttributesWithoutTileStacker
	Stringers(stringers ...fmt.Stringer) MetaAttributesWithoutTileStacker
	Byte(singleByteValue byte) MetaAttributesWithoutTileStacker
	Bytes(values []byte) MetaAttributesWithoutTileStacker
	Hex(hexValues []byte) MetaAttributesWithoutTileStacker
	RawJson(rawJsonBytes []byte) MetaAttributesWithoutTileStacker
	Error(err error) MetaAttributesWithoutTileStacker
	MapAny(mapAny map[string]any) MetaAttributesWithoutTileStacker
	MapIntegerAny(mapAny map[int]any) MetaAttributesWithoutTileStacker

	JsonResult(json *corejson.Result) MetaAttributesWithoutTileStacker
	JsonResultItems(jsons ...*corejson.Result) MetaAttributesWithoutTileStacker

	Err(err error) MetaAttributesWithoutTileStacker

	DefaultStackTraces() MetaAttributesWithoutTileStacker

	ErrWithTypeTraces(errType errcoreinf.BasicErrorTyper, err error) MetaAttributesWithoutTileStacker
	ErrorsWithTypeTraces(errType errcoreinf.BasicErrorTyper, errorItems ...error) MetaAttributesWithoutTileStacker
	StackTraces(stackSkipIndex int) MetaAttributesWithoutTileStacker
	OnErrStackTraces(err error) MetaAttributesWithoutTileStacker
	OnErrWrapperOrCollectionStackTraces(
		errWrapperOrCollection errcoreinf.BaseErrorOrCollectionWrapper,
	) MetaAttributesWithoutTileStacker

	FullStringer(
		fullStringer errcoreinf.FullStringer,
	) MetaAttributesWithoutTileStacker

	FullStringerTitle(
		title string,
		fullStringer errcoreinf.FullStringer,
	) MetaAttributesWithoutTileStacker
	FullTraceAsAttr(
		title string,
		attrFullStringWithTraces errcoreinf.FullStringWithTracesGetter,
	) MetaAttributesWithoutTileStacker

	BasicErrWrapper(errWrapperOrCollection errcoreinf.BasicErrWrapper) MetaAttributesWithoutTileStacker
	BaseRawErrCollectionDefiner(errWrapperOrCollection errcoreinf.BaseRawErrCollectionDefiner) MetaAttributesWithoutTileStacker
	BaseErrorWrapperCollectionDefiner(errWrapperOrCollection errcoreinf.BaseErrorWrapperCollectionDefiner) MetaAttributesWithoutTileStacker
	ErrWrapperOrCollection(errWrapperOrCollection errcoreinf.BaseErrorOrCollectionWrapper) MetaAttributesWithoutTileStacker
	RawErrCollection(err errcoreinf.BaseRawErrCollectionDefiner) MetaAttributesWithoutTileStacker
	CompiledBasicErrWrapper(compiler errcoreinf.CompiledBasicErrWrapper) MetaAttributesWithoutTileStacker

	Namer(namer enuminf.Namer) MetaAttributesWithoutTileStacker

	SimpleEnum(enum enuminf.SimpleEnumer) MetaAttributesWithoutTileStacker
	SimpleEnums(enums ...enuminf.SimpleEnumer) MetaAttributesWithoutTileStacker
	Enum(enum enuminf.BasicEnumer) MetaAttributesWithoutTileStacker
	Enums(enums ...enuminf.BasicEnumer) MetaAttributesWithoutTileStacker
	OnlyEnum(enum enuminf.BasicEnumer) MetaAttributesWithoutTileStacker
	OnlyEnums(enums ...enuminf.BasicEnumer) MetaAttributesWithoutTileStacker

	OnlyString(value string) MetaAttributesWithoutTileStacker
	OnlyStrings(values ...string) MetaAttributesWithoutTileStacker
	OnlyIntegers(values ...int) MetaAttributesWithoutTileStacker
	OnlyBooleans(values ...bool) MetaAttributesWithoutTileStacker
	OnlyBytes(rawBytes []byte) MetaAttributesWithoutTileStacker
	OnlyRawJson(rawBytes []byte) MetaAttributesWithoutTileStacker
	OnlyBytesErr(rawBytes []byte, err error) MetaAttributesWithoutTileStacker

	// OnlyAnyItems
	//
	//  Convert any values to json
	OnlyAnyItems(values ...any) MetaAttributesWithoutTileStacker

	// OnlyAnyItemsString
	//
	//  Convert any values to string
	OnlyAnyItemsString(values ...any) MetaAttributesWithoutTileStacker
	// OnlyAnyItemsJson
	//
	//  Convert any values to json then compile
	OnlyAnyItemsJson(values ...any) MetaAttributesWithoutTileStacker

	Bool(isResult bool) MetaAttributesWithoutTileStacker
	Booleans(isResults ...bool) MetaAttributesWithoutTileStacker

	// Any
	//
	//  Convert any item to json
	Any(anyItem any) MetaAttributesWithoutTileStacker
	// AnyIf
	//
	//  Convert any item to json
	AnyIf(isLog bool, anyItem any) MetaAttributesWithoutTileStacker
	// AnyItems
	//
	//  Convert any item to json
	AnyItems(anyItems ...any) MetaAttributesWithoutTileStacker
	// AnyItemsIf
	//
	//  Convert any item to json
	AnyItemsIf(isLog bool, anyItems ...any) MetaAttributesWithoutTileStacker

	AnyItemsJson(title string, anyItems ...any) MetaAttributesWithoutTileStacker
	AnyItemsString(title string, anyItems ...any) MetaAttributesWithoutTileStacker

	Jsoner(jsoner corejson.Jsoner) MetaAttributesWithoutTileStacker
	Jsoners(jsoners ...corejson.Jsoner) MetaAttributesWithoutTileStacker
	JsonerTitle(jsoner corejson.Jsoner) MetaAttributesWithoutTileStacker
	JsonerIf(isLog bool, jsoner corejson.Jsoner) MetaAttributesWithoutTileStacker
	JsonersIf(isLog bool, jsoners ...corejson.Jsoner) MetaAttributesWithoutTileStacker

	Serializer(serializer Serializer) MetaAttributesWithoutTileStacker
	Serializers(serializers ...Serializer) MetaAttributesWithoutTileStacker
	SerializerFunc(serializerFunc func() ([]byte, error)) MetaAttributesWithoutTileStacker
	SerializerFunctions(serializerFunctions ...func() ([]byte, error)) MetaAttributesWithoutTileStacker

	StandardTaskEntityDefiner(entity entityinf.StandardTaskEntityDefiner) MetaAttributesWithoutTileStacker
	TaskEntityDefiner(entity entityinf.TaskEntityDefiner) MetaAttributesWithoutTileStacker

	StandardTaskEntityDefinerTitle(entity entityinf.StandardTaskEntityDefiner) MetaAttributesWithoutTileStacker
	TaskEntityDefinerTitle(entity entityinf.TaskEntityDefiner) MetaAttributesWithoutTileStacker

	LoggerModel(loggerModel SingleLogModeler) MetaAttributesWithoutTileStacker
	LoggerModelTitle(loggerModel SingleLogModeler) MetaAttributesWithoutTileStacker

	Int(i int) MetaAttributesWithoutTileStacker
	Integers(integerItems ...int) MetaAttributesWithoutTileStacker
	Fmt(format string, v ...any) MetaAttributesWithoutTileStacker
	FmtIf(isLog bool, format string, v ...any) MetaAttributesWithoutTileStacker

	RawPayloadsGetter(payloadsGetter RawPayloadsGetter) MetaAttributesWithoutTileStacker
	RawPayloadsGetterIf(isLog bool, payloadsGetter RawPayloadsGetter) MetaAttributesWithoutTileStacker

	Inject(others ...MetaAttributesWithoutTileStacker) MetaAttributesWithoutTileStacker
	ConcatNew(others ...MetaAttributesWithoutTileStacker) MetaAttributesWithoutTileStacker
	coreinterface.Clearer

	Items() map[string]any

	GetAsStrings() []string
	HasKey(name string) bool
	GetVal(keyName string) (val any)

	MetaAttributesCompiler
	coreinterface.StandardSlicerContractsBinder
}
