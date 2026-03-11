# Suggestions Tracker

## Active Suggestions

### 1. Continue Test Title Audit (Remaining Packages)
**Priority:** Medium
**Status:** Pending
Remaining packages need title audit to match `"{Function} returns {Result} -- {Input Context}"` convention. See plan for full list.

### 2. Nil Receiver Coverage Audit
**Priority:** Medium
**Status:** Pending
Systematically migrate nil receiver test cases from CaseV1 to CaseNilSafe pattern across all packages.

### 3. Deep Clone Bug Investigation
**Priority:** High
**Status:** Pending
`corepayload.Attributes.Clone(deep=true)` returns error. Investigate root cause in production code.

### 4. Test Runner Hardening
**Priority:** Low
**Status:** Pending
Review all test runners for:
- Unconditional map key insertion (like the `containsName: false` pattern)
- Value vs pointer type assertions
- Incorrect independence/equality check logic

### 5. Diagnostic Output Regression Tests
**Priority:** Low
**Status:** Pending
Create snapshot tests for diagnostic output formatting to prevent regressions.

## Completed Suggestions (Moved)

See `.lovable/memory/suggestions/completed/` for implemented suggestions.
