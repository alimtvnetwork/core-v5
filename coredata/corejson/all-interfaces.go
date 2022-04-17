package corejson

type bytesSerializer interface {
	Serialize() ([]byte, error)
}

type bytesDeserializer interface {
	Deserialize(toPtr interface{}) error
}

type JsonStringBinder interface {
	JsonStringer
	PrettyJsonStringer
	AsJsonStringBinder() JsonStringBinder
}

type JsonAnyModeler interface {
	JsonModelAny() interface{}
}

type JsonContractsBinder interface {
	SimpleJsoner
	AsJsonContractsBinder() JsonContractsBinder
}

type JsonStringer interface {
	JsonString() string
}

type SimpleJsonBinder interface {
	Jsoner
	AsSimpleJsonBinder() SimpleJsonBinder
}

type JsonMarshaller interface {
	// MarshalJSON
	//
	//  alias for Serialize (from any to json)
	MarshalJSON() (jsonBytes []byte, parsedErr error)
	// UnmarshalJSON
	//
	//  alias for Deserialize (from json to any)
	UnmarshalJSON(rawJsonBytes []byte) error
}

type JsonParseSelfInjector interface {
	JsonParseSelfInject(jsonResult *Result) error
}

type PrettyJsonStringer interface {
	PrettyJsonString() string
}

type SimpleJsoner interface {
	Jsoner
	JsonParseSelfInjector
}
