package errcoreinf

import (
	"reflect"

	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coreinterface"
	"gitlab.com/evatix-go/core/coreinterface/enuminf"
)

type ShouldBeMessager interface {
	Title() string
	Actual() interface{}
	Expected() interface{}
	GenericErrorCompiler

	ConcatNew() ShouldBeMessager
	ModifyMergeNew(another ShouldBeMessager) ShouldBeMessager
}

type ShouldBeChainCollectionDefiner interface {
	GenericErrorCompiler

	ListShouldBeChainCollectionDefiner() []ShouldBeMessager
	Strings() []string
}

type ShouldBeChainer interface {
	On(isCollect bool) ShouldBeChainer
	OnString(actual, expected string) ShouldBeChainer

	IsCompleted() bool
	IsFrozen() bool
	IsAddPossible() bool

	Title(title string) ShouldBeChainer

	JsonerShouldBe(
		title string,
		actual, expected corejson.Jsoner,
	) ShouldBeChainer

	StringShouldBeDefined(
		title string,
		actual string,
	) ShouldBeChainer

	IntegerShouldBeDefined(
		title string,
		actual int,
	) ShouldBeChainer

	ShouldBeEmptyString(
		title string,
		actual string,
	) ShouldBeChainer

	ShouldBeEmptyInteger(
		title string,
		actual int,
	) ShouldBeChainer

	ShouldBeEmptyByte(
		title string,
		actual int,
	) ShouldBeChainer

	ShouldBeFalse(
		title string,
		actual bool,
	) ShouldBeChainer

	ShouldBeTrue(
		title string,
		actual bool,
	) ShouldBeChainer

	JsonResultShouldBe(
		title string,
		actual, expected *corejson.Result,
	) ShouldBeChainer

	IntegerShouldBeGreater(
		title string,
		actual, expected int,
	) ShouldBeChainer

	IntegerShouldBeGreaterOrEqual(
		title string,
		actual, expected int,
	) ShouldBeChainer

	BytesShouldBe(
		title string,
		actual, expected []byte,
	) ShouldBeChainer

	TypeShouldBe(
		title string,
		actual, expected reflect.Type,
	) ShouldBeChainer

	ShouldBeEmptyError(
		title string,
		actual error,
	) ShouldBeChainer

	ShouldBeEmptyBaseErr(
		title string,
		actual BaseErrorOrCollectionWrapper,
	) ShouldBeChainer

	PointerShouldBe(
		title string,
		actual, expected interface{},
	) ShouldBeChainer

	IntegerShouldBe(
		title string,
		actual, expected int,
	) ShouldBeChainer

	ByteShouldBe(
		title string,
		actual, expected byte,
	) ShouldBeChainer

	ChainerShouldBeEmpty(
		title string,
		actual ShouldBeChainer,
	) ShouldBeChainer

	LengthShouldBe(
		title string,
		actual coreinterface.LengthGetter,
		expected int,
	) ShouldBeChainer

	ShouldBeHaveItem(
		title string,
		actual coreinterface.LengthGetter,
	) ShouldBeChainer

	StringShouldContain(
		title string,
		actual string,
		contains string,
	) ShouldBeChainer

	SimpleEnumShouldBe(
		title string,
		actual, expected enuminf.SimpleEnumer,
	) ShouldBeChainer

	BasicEnumShouldBe(
		title string,
		actual, expected enuminf.BasicEnumer,
	) ShouldBeChainer

	BasicEnumShouldBeInvalid(
		title string,
		actual enuminf.BasicEnumer,
	) ShouldBeChainer

	BasicEnumShouldBeDefined(
		title string,
		actual enuminf.BasicEnumer,
	) ShouldBeChainer

	BooleanShouldBe(
		title string,
		actual, expected bool,
	) ShouldBeChainer

	StringShouldBe(
		title string,
		actual, expected string,
	) ShouldBeChainer

	StringShouldBeOptions(
		compareTyper enuminf.CompareMethodsTyper,
		title string,
		actual, expected string,
	) ShouldBeChainer

	StringsShouldBe(
		title string,
		actual, expected []string,
	)

	StringsShouldBeOptions(
		compareTyper enuminf.CompareMethodsTyper,
		title string,
		actual, expected []string,
	) ShouldBeChainer

	DistinctStringsShouldBeOptions(
		compareTyper enuminf.CompareMethodsTyper,
		title string,
		actual, expected []string,
	) ShouldBeChainer

	DistinctOrderStringsShouldBeOptions(
		compareTyper enuminf.CompareMethodsTyper,
		title string,
		actual, expected []string,
	) ShouldBeChainer

	OrderStringsShouldBeOptions(
		compareTyper enuminf.CompareMethodsTyper,
		title string,
		actual, expected []string,
	) ShouldBeChainer

	MapStringAnyShouldBe(
		title string,
		actual, expected map[string]interface{},
	) ShouldBeChainer

	AnyShouldBe(
		title string,
		actual, expected interface{},
	) ShouldBeChainer

	AnyShouldBeOn(
		isCollect bool,
		title string,
		actual, expected interface{},
	) ShouldBeChainer

	AnyShouldBeRegardlessOn(
		isCollect bool,
		title string,
		actual, expected interface{},
	) ShouldBeChainer

	AnyShouldBeOption(
		isRegardless bool,
		title string,
		actual, expected interface{},
	) ShouldBeChainer

	AnyShouldBeOptionOn(
		isCollect bool,
		isRegardless bool,
		title string,
		actual, expected interface{},
	) ShouldBeChainer

	AnyShouldBeRegardless(
		title string,
		actual, expected interface{},
	) ShouldBeChainer

	AnyShouldBeUsingFunc(
		title string,
		actual, expected interface{},
		compareFunc func(actual, expected interface{}) (isMatch bool),
	) ShouldBeChainer

	AnyShouldBeHaveNoPanic(
		title string,
		actual, expected interface{},
		recoverPanicCompareFunc func(actual, expected interface{}) (isMatch bool),
	) ShouldBeChainer

	Compile() BaseErrorOrCollectionWrapper
	CompileErr() error
	MustBeEmptier
	GenericErrorCompiler

	HandleError()
	CompileString() string
	CompileJson() string
	IsCollected() bool

	Append(anotherItems ...ShouldBeChainer) ShouldBeChainer

	LogOnIssues() (logged string)

	Strings() []string
}
