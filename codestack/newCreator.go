package codestack

import "runtime"

type newCreator struct{}

func (it newCreator) Default() Trace {
	return it.Create(defaultInternalSkip)
}

func (it newCreator) SkipOne() Trace {
	return it.Create(Skip2)
}

func (it newCreator) Ptr(skipIndex int) *Trace {
	trace := it.Create(skipIndex + defaultInternalSkip)

	return &trace
}

func (it newCreator) Create(skipIndex int) Trace {
	pc, file, line, isOkay := runtime.Caller(skipIndex + defaultInternalSkip)
	funcInfo := runtime.FuncForPC(pc)
	fullFuncName := funcInfo.Name()

	fullMethodSignature, packageName, methodName := NameOf.All(fullFuncName)

	return Trace{
		SkipIndex:         skipIndex,
		PackageName:       packageName,
		MethodName:        methodName,
		PackageMethodName: fullMethodSignature,
		FilePath:          file,
		Line:              line,
		IsOkay:            isOkay,
	}
}

func (it newCreator) CollectionByCap(capacity int) *TraceCollection {
	slice := make([]Trace, 0, capacity)

	return &TraceCollection{
		slice,
	}
}

func (it newCreator) Collection() *TraceCollection {
	return it.CollectionByCap(DefaultStackCount + 5)
}

func (it newCreator) CollectionUsing(
	isClone bool,
	traces ...Trace,
) *TraceCollection {
	if traces == nil {
		return it.EmptyTraces()
	}

	if !isClone {
		return &TraceCollection{
			traces,
		}
	}

	slice := it.CollectionByCap(len(traces))

	return slice.Adds(traces...)
}

func (it newCreator) EmptyTraces() *TraceCollection {
	return it.CollectionByCap(0)
}
