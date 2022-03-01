package coreinterface

type KeyValueStringDefiner interface {
	VariableNameStringGetter
	ExplicitValueStringGetter
	ValueStringGetter
	IsVariableNameEqual(name string) bool
	IsValueEqual(name string) bool
	IsEqualKeyValueStringDefiner(right KeyValueStringDefiner) bool

	CoreDefiner
}

type KeyAnyValueDefiner interface {
	VariableNameStringGetter
	ValueAnyGetter
	ExplicitValueStringGetter
	IsVariableNameEqual(name string) bool
	IsAnyValueEqual(right interface{}) bool
	IsEqualKeyAnyValueDefiner(right KeyAnyValueDefiner) bool

	CoreDefiner
}

type KeyStringValuesCollectionDefiner interface {
	KeyValueStringDefiners() []KeyValueStringDefiner
	HashmapGetter
	KeysHashsetGetter
	StringsGetter
	HasKeyChecker

	IsEqualKeyStringValuesCollectionDefiner(
		right KeyStringValuesCollectionDefiner,
	) bool
	CoreDefiner
}

type KeyAnyValuesCollectionDefiner interface {
	KeyValueStringDefiners() []KeyAnyValueDefiner
	HashmapGetter
	KeysHashsetGetter
	StringsGetter
	MapStringAnyGetter
	HasKeyChecker

	IsEqualKeyAnyValuesCollectionDefiner(
		right KeyAnyValuesCollectionDefiner,
	) bool

	CoreDefiner
}
