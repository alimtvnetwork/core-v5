package regexnewtests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/regexnew"
)

// Cover: CreateLockIf, CreateMustLockIf, CreateApplicableLock, IsMatchFailed,
// MatchError, MatchErrorLock, MatchUsingFuncErrorLock, MatchUsingCustomizeErrorFuncLock,
// NewMustLock, newCreator methods, lazyRegexMap methods, LazyRegex methods

func Test_CreateLockIf_WithLock_Cov2(t *testing.T) {
	r, err := regexnew.CreateLockIf(true, `^\d+$`)
	if err != nil || r == nil {
		t.Error("expected valid regex")
	}
}

func Test_CreateLockIf_WithoutLock_Cov2(t *testing.T) {
	r, err := regexnew.CreateLockIf(false, `^\d+$`)
	if err != nil || r == nil {
		t.Error("expected valid regex")
	}
}

func Test_CreateLockIf_Invalid_Cov2(t *testing.T) {
	_, err := regexnew.CreateLockIf(true, `[invalid`)
	if err == nil {
		t.Error("expected error")
	}
}

func Test_CreateMustLockIf_WithLock_Cov2(t *testing.T) {
	r := regexnew.CreateMustLockIf(true, `^\d+$`)
	if r == nil {
		t.Error("expected valid regex")
	}
}

func Test_CreateMustLockIf_WithoutLock_Cov2(t *testing.T) {
	r := regexnew.CreateMustLockIf(false, `^\d+$`)
	if r == nil {
		t.Error("expected valid regex")
	}
}

func Test_CreateApplicableLock_Valid_Cov2(t *testing.T) {
	r, err, isApplicable := regexnew.CreateApplicableLock(`^\d+$`)
	if err != nil || r == nil || !isApplicable {
		t.Error("expected applicable")
	}
}

func Test_CreateApplicableLock_Invalid_Cov2(t *testing.T) {
	_, err, isApplicable := regexnew.CreateApplicableLock(`[invalid`)
	if err == nil || isApplicable {
		t.Error("expected not applicable")
	}
}

func Test_IsMatchFailed_Cov2(t *testing.T) {
	if regexnew.IsMatchFailed(`^\d+$`, "123") {
		t.Error("should match")
	}
	if !regexnew.IsMatchFailed(`^\d+$`, "abc") {
		t.Error("should not match")
	}
}

func Test_MatchError_Match_Cov2(t *testing.T) {
	err := regexnew.MatchError(`^\d+$`, "123")
	if err != nil {
		t.Error("should match")
	}
}

func Test_MatchError_NoMatch_Cov2(t *testing.T) {
	err := regexnew.MatchError(`^\d+$`, "abc")
	if err == nil {
		t.Error("should error")
	}
}

func Test_MatchError_InvalidPattern_Cov2(t *testing.T) {
	err := regexnew.MatchError(`[invalid`, "abc")
	if err == nil {
		t.Error("should error")
	}
}

func Test_MatchErrorLock_Match_Cov2(t *testing.T) {
	err := regexnew.MatchErrorLock(`^\d+$`, "123")
	if err != nil {
		t.Error("should match")
	}
}

func Test_MatchErrorLock_NoMatch_Cov2(t *testing.T) {
	err := regexnew.MatchErrorLock(`^\d+$`, "abc")
	if err == nil {
		t.Error("should error")
	}
}

func Test_MatchUsingFuncErrorLock_Match_Cov2(t *testing.T) {
	matchFunc := func(r *regexp.Regexp, s string) bool {
		return r.MatchString(s)
	}
	err := regexnew.MatchUsingFuncErrorLock(`^\d+$`, "123", matchFunc)
	if err != nil {
		t.Error("should match")
	}
}

func Test_MatchUsingFuncErrorLock_NoMatch_Cov2(t *testing.T) {
	matchFunc := func(r *regexp.Regexp, s string) bool {
		return r.MatchString(s)
	}
	err := regexnew.MatchUsingFuncErrorLock(`^\d+$`, "abc", matchFunc)
	if err == nil {
		t.Error("should error")
	}
}

func Test_MatchUsingCustomizeErrorFuncLock_Match_Cov2(t *testing.T) {
	matchFunc := func(r *regexp.Regexp, s string) bool {
		return r.MatchString(s)
	}
	err := regexnew.MatchUsingCustomizeErrorFuncLock(`^\d+$`, "123", matchFunc, nil)
	if err != nil {
		t.Error("should match")
	}
}

func Test_MatchUsingCustomizeErrorFuncLock_NoMatch_NilCustomize_Cov2(t *testing.T) {
	matchFunc := func(r *regexp.Regexp, s string) bool {
		return r.MatchString(s)
	}
	err := regexnew.MatchUsingCustomizeErrorFuncLock(`^\d+$`, "abc", matchFunc, nil)
	if err == nil {
		t.Error("should error")
	}
}

func Test_MatchUsingCustomizeErrorFuncLock_NoMatch_WithCustomize_Cov2(t *testing.T) {
	matchFunc := func(r *regexp.Regexp, s string) bool {
		return r.MatchString(s)
	}
	customErr := func(pattern, term string, err error, r *regexp.Regexp) error {
		return err
	}
	// valid pattern but no match → customizeErrFunc called with nil err
	result := regexnew.MatchUsingCustomizeErrorFuncLock(`^\d+$`, "abc", matchFunc, customErr)
	// err from compiler is nil, so customErr returns nil
	_ = result
}

func Test_NewMustLock_Cov2(t *testing.T) {
	r := regexnew.NewMustLock(`^\d+$`)
	if r == nil {
		t.Error("expected valid regex")
	}
}

func Test_NewCreator_Default_Cov2(t *testing.T) {
	r, err := regexnew.New.Default(`^\d+$`)
	if err != nil || r == nil {
		t.Error("expected valid regex")
	}
}

func Test_NewCreator_DefaultLock_Cov2(t *testing.T) {
	r, err := regexnew.New.DefaultLock(`^\d+$`)
	if err != nil || r == nil {
		t.Error("expected valid regex")
	}
}

func Test_NewCreator_DefaultLockIf_Cov2(t *testing.T) {
	r, err := regexnew.New.DefaultLockIf(true, `^\d+$`)
	if err != nil || r == nil {
		t.Error("expected valid regex")
	}
}

func Test_NewCreator_DefaultApplicableLock_Cov2(t *testing.T) {
	r, err, isApplicable := regexnew.New.DefaultApplicableLock(`^\d+$`)
	if err != nil || r == nil || !isApplicable {
		t.Error("expected applicable")
	}
}

func Test_NewCreator_Lazy_Cov2(t *testing.T) {
	lazy := regexnew.New.Lazy(`^\d+$`)
	if lazy == nil {
		t.Error("expected non-nil")
	}
}

func Test_NewCreator_LazyLock_Cov2(t *testing.T) {
	lazy := regexnew.New.LazyLock(`^\d+$`)
	if lazy == nil {
		t.Error("expected non-nil")
	}
}

func Test_LazyRegex_FullString_Cov2(t *testing.T) {
	lazy := regexnew.New.Lazy(`^\d+$`)
	s := lazy.FullString()
	if s == "" {
		t.Error("expected non-empty")
	}
}

func Test_LazyRegex_IsMatchBytes_Cov2(t *testing.T) {
	lazy := regexnew.New.Lazy(`^\d+$`)
	if !lazy.IsMatchBytes([]byte("123")) {
		t.Error("should match")
	}
	if lazy.IsMatchBytes([]byte("abc")) {
		t.Error("should not match")
	}
}

func Test_LazyRegex_IsFailedMatch_Cov2(t *testing.T) {
	lazy := regexnew.New.Lazy(`^\d+$`)
	if lazy.IsFailedMatch("123") {
		t.Error("should not fail")
	}
	if !lazy.IsFailedMatch("abc") {
		t.Error("should fail")
	}
}

func Test_LazyRegex_IsFailedMatchBytes_Cov2(t *testing.T) {
	lazy := regexnew.New.Lazy(`^\d+$`)
	if lazy.IsFailedMatchBytes([]byte("123")) {
		t.Error("should not fail")
	}
	if !lazy.IsFailedMatchBytes([]byte("abc")) {
		t.Error("should fail")
	}
}

func Test_LazyRegex_FirstMatchLine_Valid_Cov2(t *testing.T) {
	lazy := regexnew.New.Lazy(`(\d+)`)
	match, isInvalid := lazy.FirstMatchLine("abc123def")
	if isInvalid || match != "123" {
		t.Errorf("expected 123, got %s, invalid=%v", match, isInvalid)
	}
}

func Test_LazyRegex_FirstMatchLine_NoMatch_Cov2(t *testing.T) {
	lazy := regexnew.New.Lazy(`(\d+)`)
	match, isInvalid := lazy.FirstMatchLine("abcdef")
	if !isInvalid || match != "" {
		t.Error("expected invalid match")
	}
}

func Test_LazyRegex_MatchError_Cov2(t *testing.T) {
	lazy := regexnew.New.Lazy(`^\d+$`)
	if err := lazy.MatchError("123"); err != nil {
		t.Error("should match")
	}
	if err := lazy.MatchError("abc"); err == nil {
		t.Error("should error")
	}
}

func Test_LazyRegex_MatchUsingFuncError_Cov2(t *testing.T) {
	lazy := regexnew.New.Lazy(`^\d+$`)
	matchFunc := func(r *regexp.Regexp, s string) bool {
		return r.MatchString(s)
	}
	if err := lazy.MatchUsingFuncError("123", matchFunc); err != nil {
		t.Error("should match")
	}
	if err := lazy.MatchUsingFuncError("abc", matchFunc); err == nil {
		t.Error("should error")
	}
}

func Test_LazyRegex_MustBeSafe_Cov2(t *testing.T) {
	lazy := regexnew.New.Lazy(`^\d+$`)
	// should not panic
	lazy.MustBeSafe()
}

func Test_LazyRegex_HasError_Valid_Cov2(t *testing.T) {
	lazy := regexnew.New.Lazy(`^\d+$`)
	if lazy.HasError() {
		t.Error("should not have error")
	}
}

func Test_LazyRegex_HasAnyIssues_Valid_Cov2(t *testing.T) {
	lazy := regexnew.New.Lazy(`^\d+$`)
	if lazy.HasAnyIssues() {
		t.Error("should not have issues")
	}
}

func Test_LazyRegex_IsInvalid_Valid_Cov2(t *testing.T) {
	lazy := regexnew.New.Lazy(`^\d+$`)
	if lazy.IsInvalid() {
		t.Error("should not be invalid")
	}
}

func Test_LazyRegex_CompiledError_Cov2(t *testing.T) {
	lazy := regexnew.New.Lazy(`^\d+$`)
	if lazy.CompiledError() != nil {
		t.Error("should not have error")
	}
}

func Test_LazyRegex_Error_Cov2(t *testing.T) {
	lazy := regexnew.New.Lazy(`^\d+$`)
	if lazy.Error() != nil {
		t.Error("should not have error")
	}
}

func Test_LazyRegex_NilReceiver_Cov2(t *testing.T) {
	var lazy *regexnew.LazyRegex
	if !lazy.IsNull() {
		t.Error("should be null")
	}
	if lazy.IsDefined() {
		t.Error("should not be defined")
	}
	if !lazy.IsUndefined() {
		t.Error("should be undefined")
	}
	if lazy.IsCompiled() {
		t.Error("should not be compiled")
	}
	if lazy.String() != "" {
		t.Error("should be empty")
	}
	if lazy.Pattern() != "" {
		t.Error("should be empty")
	}
	if lazy.FullString() != "" {
		t.Error("should be empty")
	}
	if !lazy.HasAnyIssues() {
		t.Error("nil should have issues")
	}
	if !lazy.IsInvalid() {
		t.Error("nil should be invalid")
	}
	if lazy.IsApplicable() {
		t.Error("nil should not be applicable")
	}
	if err := lazy.OnRequiredCompiled(); err == nil {
		t.Error("expected error for nil")
	}
}

func Test_LazyRegex_OnRequiredCompiledMust_NilPanic_Cov2(t *testing.T) {
	var lazy *regexnew.LazyRegex
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic")
		}
	}()
	lazy.OnRequiredCompiledMust()
}

func Test_LazyRegex_CompileMust_Valid_Cov2(t *testing.T) {
	lazy := regexnew.New.Lazy(`^\d+$`)
	r := lazy.CompileMust()
	if r == nil {
		t.Error("expected non-nil")
	}
}

func Test_NewLazyRegexCreator_TwoLock_Cov2(t *testing.T) {
	first, second := regexnew.New.LazyRegex.TwoLock(`^\d+$`, `^[a-z]+$`)
	if first == nil || second == nil {
		t.Error("expected non-nil")
	}
}

func Test_NewLazyRegexCreator_ManyUsingLock_Cov2(t *testing.T) {
	m := regexnew.New.LazyRegex.ManyUsingLock(`^\d+$`, `^[a-z]+$`)
	if len(m) != 2 {
		t.Error("expected 2")
	}
}

func Test_NewLazyRegexCreator_ManyUsingLock_Empty_Cov2(t *testing.T) {
	m := regexnew.New.LazyRegex.ManyUsingLock()
	if len(m) != 0 {
		t.Error("expected 0")
	}
}

func Test_NewLazyRegexCreator_AllPatternsMap_Cov2(t *testing.T) {
	m := regexnew.New.LazyRegex.AllPatternsMap()
	if m == nil {
		t.Error("expected non-nil")
	}
}

func Test_NewLazyRegexCreator_NewLockIf_Cov2(t *testing.T) {
	r := regexnew.New.LazyRegex.NewLockIf(true, `^\d+$`)
	if r == nil {
		t.Error("expected non-nil")
	}
	r = regexnew.New.LazyRegex.NewLockIf(false, `^[a-z]+$`)
	if r == nil {
		t.Error("expected non-nil")
	}
}
