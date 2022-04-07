package chmodhelper

import "gitlab.com/evatix-go/core/internal/osconstsinternal"

type tempDirGetter struct{}

func (it tempDirGetter) TempDefault() string {
	return TempDirDefault
}

func (it tempDirGetter) TempPermanent() string {
	if osconstsinternal.IsWindows {
		return osconstsinternal.WindowsPermanentTemp
	}

	// unix
	return osconstsinternal.LinuxPermanentTemp
}

func (it tempDirGetter) TempOption(isPermanent bool) string {
	if isPermanent {
		return it.TempPermanent()
	}

	return it.TempDefault()
}
