package codestacktests

import (
	"github.com/alimtvnetwork/core/codestack"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/coretests/results"
)

type fileWithLineType = codestack.FileWithLine
type traceType = codestack.Trace
type traceCollectionType = codestack.TraceCollection
type nilResult = results.ResultAny

// ── FileWithLine nil-safety ──

var coverageFileWithLineNilSafeCases = []coretestcases.CaseNilSafe{
	{
		Title:    "FullFilePath on nil returns empty",
		Func:     (*fileWithLineType).FullFilePath,
		Expected: nilResult{Panicked: true},
	},
	{
		Title:    "LineNumber on nil returns 0",
		Func:     (*fileWithLineType).LineNumber,
		Expected: nilResult{Panicked: true},
	},
	{
		Title:    "IsNil on nil returns true",
		Func:     (*fileWithLineType).IsNil,
		Expected: nilResult{Value: "true", Panicked: false},
	},
	{
		Title:    "IsNotNil on nil returns false",
		Func:     (*fileWithLineType).IsNotNil,
		Expected: nilResult{Value: "false", Panicked: false},
	},
	{
		Title:    "String on nil returns empty",
		Func:     (*fileWithLineType).String,
		Expected: nilResult{Value: "", Panicked: false},
	},
	{
		Title:    "JsonModelAny on nil panics",
		Func:     (*fileWithLineType).JsonModelAny,
		Expected: nilResult{Panicked: true},
	},
	{
		Title:    "JsonString on nil panics",
		Func:     (*fileWithLineType).JsonString,
		Expected: nilResult{Panicked: true},
	},
	{
		Title:    "AsFileLiner on nil returns nil",
		Func:     (*fileWithLineType).AsFileLiner,
		Expected: nilResult{Panicked: false},
	},
}

// ── Trace nil-safety ──

var coverageTraceNilSafeCases = []coretestcases.CaseNilSafe{
	{
		Title:    "Trace.IsNil on nil returns true",
		Func:     (*traceType).IsNil,
		Expected: nilResult{Value: "true", Panicked: false},
	},
	{
		Title:    "Trace.IsNotNil on nil returns false",
		Func:     (*traceType).IsNotNil,
		Expected: nilResult{Value: "false", Panicked: false},
	},
	{
		Title:    "Trace.HasIssues on nil returns true",
		Func:     (*traceType).HasIssues,
		Expected: nilResult{Value: "true", Panicked: false},
	},
	{
		Title:    "Trace.String on nil returns empty",
		Func:     (*traceType).String,
		Expected: nilResult{Value: "", Panicked: false},
	},
	{
		Title:    "Trace.Dispose on nil does not panic",
		Func:     (*traceType).Dispose,
		Expected: nilResult{Panicked: false},
	},
	{
		Title:    "Trace.ClonePtr on nil returns nil",
		Func:     (*traceType).ClonePtr,
		Expected: nilResult{Panicked: false},
	},
	{
		Title:    "Trace.JsonModelAny on nil panics",
		Func:     (*traceType).JsonModelAny,
		Expected: nilResult{Panicked: true},
	},
	{
		Title:    "Trace.AsFileLiner on nil returns nil",
		Func:     (*traceType).AsFileLiner,
		Expected: nilResult{Panicked: false},
	},
}
