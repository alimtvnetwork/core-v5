package corejson

type JsonSimpleContractsBinder interface {
	Jsoner
	JsonParseSelfInjector
	AsJsonSimpleContractsBinder() JsonSimpleContractsBinder
}
