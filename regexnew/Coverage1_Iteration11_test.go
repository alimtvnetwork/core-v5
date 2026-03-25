package regexnew

import (
	"errors"
	"regexp"
	"strings"
	"testing"
)

func Test_I11_RX_LazyRegexMap_StateAndLockMethods(t *testing.T) {
	var nilMap *lazyRegexMap

	if !nilMap.IsEmpty() || !nilMap.IsEmptyLock() {
		t.Fatal("nil map should be empty")
	}

	if nilMap.HasAnyItem() || nilMap.HasAnyItemLock() {
		t.Fatal("nil map should not have items")
	}

	if nilMap.Length() != 0 || nilMap.LengthLock() != 0 {
		t.Fatal("nil map length should be zero")
	}

	m := &lazyRegexMap{items: map[string]*LazyRegex{}}
	if !m.IsEmpty() {
		t.Fatal("new map should be empty")
	}

	first, existed := m.CreateOrExisting(`^i11\\d+$`)
	if existed || first == nil {
		t.Fatal("expected creation path")
	}

	if !m.Has(`^i11\\d+$`) || !m.HasLock(`^i11\\d+$`) {
		t.Fatal("expected map has key")
	}

	if !m.HasAnyItem() || !m.HasAnyItemLock() {
		t.Fatal("expected map has item")
	}

	if m.Length() == 0 || m.LengthLock() == 0 {
		t.Fatal("expected non-zero length")
	}

	second, existedAgain := m.CreateOrExisting(`^i11\\d+$`)
	if !existedAgain || first != second {
		t.Fatal("expected existing pointer on second call")
	}

	third, existedLock := m.CreateOrExistingLock(`^i11-lock$`)
	if existedLock || third == nil {
		t.Fatal("expected lock create path")
	}

	fourth, existedLockIf := m.CreateOrExistingLockIf(true, `^i11-lock-if$`)
	if existedLockIf || fourth == nil {
		t.Fatal("expected lock-if create path")
	}

	fifth, existedNoLockIf := m.CreateOrExistingLockIf(false, `^i11-no-lock-if$`)
	if existedNoLockIf || fifth == nil {
		t.Fatal("expected no-lock-if create path")
	}
}

func Test_I11_RX_LazyRegexMap_CreateLazyRegex_CustomCompiler(t *testing.T) {
	m := &lazyRegexMap{items: map[string]*LazyRegex{}}
	custom := m.createLazyRegex(`^z+$`, func(pattern string) (*regexp.Regexp, error) {
		return regexp.Compile(pattern)
	})

	if custom == nil || custom.pattern != `^z+$` {
		t.Fatal("expected custom lazy regex")
	}

	compiled, err := custom.Compile()
	if err != nil || compiled == nil {
		t.Fatal("expected custom compiler success")
	}
}

func Test_I11_RX_PrettyJson_And_MatchErrorBranches(t *testing.T) {
	if prettyJson(nil) != "" {
		t.Fatal("nil prettyJson should be empty")
	}

	marshalFail := struct {
		Fn func()
	}{Fn: func() {}}
	if prettyJson(marshalFail) != "" {
		t.Fatal("marshal-fail prettyJson should be empty")
	}

	errCompile := regExMatchValidationError("", "abc", errors.New("bad-regex"), nil)
	if errCompile == nil || !strings.Contains(errCompile.Error(), "compile failed") {
		t.Fatal("expected compile-failed message")
	}

	errNilRegex := regExMatchValidationError("^x$", "abc", nil, nil)
	if errNilRegex == nil || !strings.Contains(errNilRegex.Error(), "invalid cannot match") {
		t.Fatal("expected nil-regex message")
	}

	errNoMatch := regExMatchValidationError("^x$", "abc", nil, regexp.MustCompile(`^x$`))
	if errNoMatch == nil || !strings.Contains(errNoMatch.Error(), "doesn't match") {
		t.Fatal("expected mismatch message")
	}
}

func Test_I11_RX_LazyRegex_IsMatchBytes_InvalidPatternBranch(t *testing.T) {
	lr := &LazyRegex{
		pattern:  "[",
		compiler: CreateLock,
	}

	if lr.IsMatchBytes([]byte("anything")) {
		t.Fatal("invalid pattern should not match bytes")
	}

	if !lr.IsFailedMatchBytes([]byte("anything")) {
		t.Fatal("invalid pattern should fail match bytes")
	}
}
