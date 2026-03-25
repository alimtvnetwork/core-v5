# Completed: Diagnostic Formatting Improvements

## Completed: 2026-03-11

### What was done
1. MapMismatchError header uses 4-space indented 3-line format with leading newline
2. LineDiff actual/expected labels align colons at column 21
3. args.Map ExpectedInput compiles to strings before assertion (prevents PrettyJSON fallthrough)
4. Map diagnostic entries use Go-literal format with tab indentation
5. Separator headers (`============================>`) restored for visual structure

### Files Modified
- `errcore/MapMismatchError.go`
- `errcore/LineDiff.go`
- `coretests/coretestcases/CaseV1MapAssertions.go`
