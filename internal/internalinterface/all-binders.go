package internalinterface

type PathExtenderBinder interface {
	PathExtender
	AsPathExtender() PathExtender
}

type LogTypeCheckerBinder interface {
	LogTypeChecker
	AsLogTypeChecker() LogTypeChecker
}
