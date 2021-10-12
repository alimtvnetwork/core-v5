package corejson

type JsonStandardContractsBinder interface {
	Jsoner
	JsonParseSelfInjector
	JsonMarshaller
	AsJsonStandardContractsBinder() JsonStandardContractsBinder
}
