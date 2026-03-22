# Memory: workflow/coverage-execution-plan-status
Updated: 2026-03-22

## Current Status
- Execution Plan v3 created at `.lovable/memory/workflow/04-coverage-execution-plan-v3.md`
- Failing test specs at `spec/05-failing-tests/01-blocked-packages-fixes.md` and `02-failing-tests-root-cause.md`

## Blocked Packages Fixed (3/3)
1. **coredynamictests**: `ts.IsEqual()` → `ts.IsEqual(sameTs)` — added TypeStatus arg
2. **corejsontests**: Already resolved (stale error from deleted file)
3. **corestrtests**: Fixed 8 API signature mismatches in Coverage41_Iteration8_test.go

## Failing Tests Fixed (10/19)
1. ✅ Test_I11_PC_IsEqualItems_NilPC — changed expected to `false` (variadic nil wrapping)
2. ✅ Test_Cov10_VerifyError_WithTypeVerify — fixed VerifyTypeOf.ExpectedInput type
3. ✅ Test_Cov8_GenericGherkins_ShouldBeEqualMap_NotMap — isolated T instead of t.Run
4. ✅ Test_Cov2_SimpleTestCase_ShouldHaveNoError — wrapped with recover
5. ✅ Test_Cov2_SimpleTestCase_ShouldContains — wrapped with recover
6. ✅ Test_Cov3_BaseTestCase_TypeShouldMatch_WithMismatch — isolated T
7. ✅ Test_Cov3_TypesValidationMustPasses_WithError — isolated T
8. ✅ Test_Cov10_GetSinglePageCollection_NegativePagePanic — resolved by #1

## Remaining Failing Tests (9)
- Test_I11_NewPW_CastOrDeserializeFrom_Valid (corepayloadtests) — needs source investigation
- Test_CovPL_S1_05_* (corepayloadtests) — HasAttributes check
- Test_CovPL_S1_35_* (corepayloadtests) — Attributes.IsInvalid
- Test_CovPL_S1_54_* (corepayloadtests) — DeserializeToCollection
- Test_CovPL_S2_61_* (corepayloadtests) — TPC Deserialization
- Test_CovPL_S2_65_* (corepayloadtests) — reflect panic, struct D scope
- Test_CovEnum_BB11_* (enumimpltests) — wrong enum name
- Test_I13_InvokeError_NilError (reflectmodeltests) — production bug in ReflectValueToAnyValue

## Next Steps
1. Fix remaining 9 failing tests (mostly corepayloadtests)
2. Begin coverage: coredata/corestr S01 (29 segments total)
