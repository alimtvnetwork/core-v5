package corejson

type JsonerBinder interface {
	Jsoner
	AsJsonerBinder() JsonerBinder
}
