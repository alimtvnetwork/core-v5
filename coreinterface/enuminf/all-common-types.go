package enuminf

import (
	"gitlab.com/evatix-go/core/internal/internalinterface"
	"gitlab.com/evatix-go/core/internal/internalinterface/internalenuminf"
)

type EnvironmentFlagTyper interface {
	BasicEnumer
	internalenuminf.EnvironmentFlagTyper
}

type EnvironmentTyper interface {
	BasicEnumer
	internalenuminf.EnvironmentTyper
}

type EnvironmentOptioner interface {
	EnvTyper() EnvironmentTyper
	FlagTyper() EnvironmentFlagTyper
}

type CrudTyper interface {
	BasicEnumer
	internalenuminf.CrudTyper
}

type LinuxTyper interface {
	BasicEnumer
	internalenuminf.LinuxTyper
}

type LoggerTyper interface {
	BasicEnumer
	internalinterface.LoggerTyper
}

type LogLevelTyper interface {
	BasicEnumer
	internalenuminf.LogLevelTyper
}

type CompareMethodsTyper interface {
	BasicEnumer
	internalenuminf.CompareMethodsTyper
}
