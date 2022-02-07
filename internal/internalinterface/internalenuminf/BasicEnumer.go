package internalenuminf

type BasicEnumer interface {
	enumNameStinger
	nameValuer
	IsNameEqualer
	IsAnyNameOfChecker
	toNumberStringer
	IsValidInvalidChecker
	BasicEnumValuer
}
