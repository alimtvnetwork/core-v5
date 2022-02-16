package loggerinf

import (
	"fmt"

	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coreinterface/entityinf"
	"gitlab.com/evatix-go/core/coreinterface/enuminf"
	"gitlab.com/evatix-go/core/coreinterface/errcoreinf"
)

type SingleLogger interface {
	LoggerTyperGetter
	Stack() MetaAttributesStacker
	StackTitle(title string) MetaAttributesStacker

	On(isLog bool) SingleLogger
	StackSkip(stackSkipIndex int) SingleLogger
	OnString(input, expected string) SingleLogger

	Title(message string) SingleLogger
	TitleAttr(message, attr string) SingleLogger
	Log(message string) SingleLogger
	LogAttr(message, attr string) SingleLogger
	Str(title, val string) SingleLogger
	Strings(title string, values []string) SingleLogger
	StringsSpread(title string, values ...string) SingleLogger
	Stringer(title string, stringer fmt.Stringer) SingleLogger
	Stringers(title string, stringers ...fmt.Stringer) SingleLogger
	Byte(title string, val byte) SingleLogger
	Bytes(title string, values []byte) SingleLogger
	Hex(title string, val []byte) SingleLogger
	RawJson(title string, rawJson []byte) SingleLogger
	Err(err error) SingleLogger
	AnErr(title string, err error) SingleLogger
	ErrWithType(title string, errType errcoreinf.BasicErrorTyper, err error) SingleLogger

	DefaultStackTraces() SingleLogger
	ErrWithTypeTraces(title string, errType errcoreinf.BasicErrorTyper, err error) SingleLogger
	ErrorsWithTypeTraces(title string, errType errcoreinf.BasicErrorTyper, errorItems ...error) SingleLogger
	StackTraces(stackSkipIndex int, title string) SingleLogger
	OnErrStackTraces(err error) SingleLogger
	OnErrWrapperOrCollectionStackTraces(errWrapperOrCollection errcoreinf.BaseErrorOrCollectionWrapper) SingleLogger

	FullStringer(
		fullStringer errcoreinf.FullStringer,
	) SingleLogger

	FullStringerTitle(
		title string,
		fullStringer errcoreinf.FullStringer,
	) SingleLogger
	FullTraceAsAttr(
		title string,
		attrFullStringWithTraces errcoreinf.FullStringWithTracesGetter,
	) SingleLogger

	BasicErrWrapper(errWrapperOrCollection errcoreinf.BasicErrWrapper) SingleLogger
	BaseRawErrCollectionDefiner(errWrapperOrCollection errcoreinf.BaseRawErrCollectionDefiner) SingleLogger
	BaseErrorWrapperCollectionDefiner(errWrapperOrCollection errcoreinf.BaseErrorWrapperCollectionDefiner) SingleLogger
	ErrWrapperOrCollection(errWrapperOrCollection errcoreinf.BaseErrorOrCollectionWrapper) SingleLogger
	RawErrCollection(title string, err errcoreinf.BaseRawErrCollectionDefiner) SingleLogger
	CompiledBasicErrWrapper(compiler errcoreinf.CompiledBasicErrWrapper) SingleLogger

	Namer(namer enuminf.Namer) SingleLogger
	NamerTitle(title string, namer enuminf.Namer) SingleLogger

	Enum(title string, enum enuminf.BasicEnumer) SingleLogger
	Enums(title string, enums ...enuminf.BasicEnumer) SingleLogger
	OnlyEnum(enum enuminf.BasicEnumer) SingleLogger
	OnlyEnums(enums ...enuminf.BasicEnumer) SingleLogger
	OnlyError(err error) SingleLogger
	OnlyString(value string) SingleLogger
	OnlyStrings(values ...string) SingleLogger
	OnlyIntegers(values ...int) SingleLogger
	OnlyBooleans(values ...bool) SingleLogger
	OnlyBytes(rawBytes []byte) SingleLogger
	OnlyRawJson(rawBytes []byte) SingleLogger
	OnlyBytesErr(rawBytes []byte, err error) SingleLogger
	OnlyAnyItems(values ...interface{}) SingleLogger
	Bool(title string, isResult bool) SingleLogger
	Booleans(title string, isResults ...bool) SingleLogger
	AnyJsonLog(anyItem interface{}) SingleLogger
	Any(anyItem interface{}) SingleLogger
	AnyIf(isLog bool, anyItem interface{}) SingleLogger
	AnyItems(anyItems ...interface{}) SingleLogger
	AnyItemsIf(isLog bool, anyItems ...interface{}) SingleLogger

	Jsoner(jsoner corejson.Jsoner) SingleLogger
	Jsoners(jsoners ...corejson.Jsoner) SingleLogger
	JsonerTitle(title string, jsoner corejson.Jsoner) SingleLogger

	Serializer(serializer Serializer) SingleLogger
	Serializers(serializers ...Serializer) SingleLogger
	SerializerFunc(serializerFunc func() ([]byte, error)) SingleLogger
	SerializerFunctions(serializerFunctions ...func() ([]byte, error)) SingleLogger

	StandardTaskEntityDefiner(entity entityinf.StandardTaskEntityDefiner) SingleLogger
	TaskEntityDefiner(entity entityinf.TaskEntityDefiner) SingleLogger

	StandardTaskEntityDefinerTitle(title string, entity entityinf.StandardTaskEntityDefiner) SingleLogger
	TaskEntityDefinerTitle(title string, entity entityinf.TaskEntityDefiner) SingleLogger

	LogModel(model SingleLogModeler) SingleLogger
	LogModelTitle(title string, model SingleLogModeler) SingleLogger

	Int(title string, i int) SingleLogger
	Integers(title string, integerItems ...int) SingleLogger

	Fmt(format string, v ...interface{}) SingleLogger
	AttrFmt(title string, attrFormat string, attrValues ...interface{}) SingleLogger

	RawPayloadsGetter(payloadsGetter RawPayloadsGetter) SingleLogger
	RawPayloadsGetterTitle(title string, payloadsGetter RawPayloadsGetter) SingleLogger

	Logger() StandardLogger
}
