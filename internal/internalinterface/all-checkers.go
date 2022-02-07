package internalinterface

type IsIdentifierEqualer interface {
	IsIdentifier(identifier string) bool
}

type IsIdEqualer interface {
	IsId(id string) bool
}

type IsIdUnsignedIntegerEqualer interface {
	IsId(id uint) bool
}

type HasErrorChecker interface {
	HasError() bool
}
