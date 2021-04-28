package corejson

type JsonParseSelfInjector interface {
	JsonParseSelfInject(jsonResult *Result) error
	AsJsonParseSelfInjector() JsonParseSelfInjector
}
