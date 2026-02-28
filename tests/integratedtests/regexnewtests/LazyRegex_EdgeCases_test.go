package regexnewtests

import (
	"fmt"
	"sync"
	"testing"

	"gitlab.com/auk-go/core/regexnew"
)

func Test_NilLazyRegex_IsNull(t *testing.T) {
	// Arrange
	var lazy *regexnew.LazyRegex

	// Act & Assert
	if !lazy.IsNull() {
		t.Error("nil LazyRegex should return true for IsNull")
	}
}

func Test_NilLazyRegex_IsUndefined(t *testing.T) {
	// Arrange
	var lazy *regexnew.LazyRegex

	// Act & Assert
	if !lazy.IsUndefined() {
		t.Error("nil LazyRegex should return true for IsUndefined")
	}
}

func Test_NilLazyRegex_IsDefined(t *testing.T) {
	// Arrange
	var lazy *regexnew.LazyRegex

	// Act & Assert
	if lazy.IsDefined() {
		t.Error("nil LazyRegex should return false for IsDefined")
	}
}

func Test_NilLazyRegex_IsCompiled(t *testing.T) {
	// Arrange
	var lazy *regexnew.LazyRegex

	// Act & Assert
	if lazy.IsCompiled() {
		t.Error("nil LazyRegex should return false for IsCompiled")
	}
}

func Test_NilLazyRegex_String(t *testing.T) {
	// Arrange
	var lazy *regexnew.LazyRegex

	// Act
	result := lazy.String()

	// Assert
	if result != "" {
		t.Errorf("nil LazyRegex String() should return empty, got %q", result)
	}
}

func Test_NilLazyRegex_Pattern(t *testing.T) {
	// Arrange
	var lazy *regexnew.LazyRegex

	// Act
	result := lazy.Pattern()

	// Assert
	if result != "" {
		t.Errorf("nil LazyRegex Pattern() should return empty, got %q", result)
	}
}

func Test_NilLazyRegex_HasAnyIssues(t *testing.T) {
	// Arrange
	var lazy *regexnew.LazyRegex

	// Act & Assert
	if !lazy.HasAnyIssues() {
		t.Error("nil LazyRegex should return true for HasAnyIssues")
	}
}

func Test_NilLazyRegex_IsInvalid(t *testing.T) {
	// Arrange
	var lazy *regexnew.LazyRegex

	// Act & Assert
	if !lazy.IsInvalid() {
		t.Error("nil LazyRegex should return true for IsInvalid")
	}
}

func Test_NilLazyRegex_OnRequiredCompiled(t *testing.T) {
	// Arrange
	var lazy *regexnew.LazyRegex

	// Act
	err := lazy.OnRequiredCompiled()

	// Assert
	if err == nil {
		t.Error("nil LazyRegex OnRequiredCompiled should return error")
	}
}

func Test_NilLazyRegex_FullString(t *testing.T) {
	// Arrange
	var lazy *regexnew.LazyRegex

	// Act
	result := lazy.FullString()

	// Assert
	if result != "" {
		t.Errorf("nil LazyRegex FullString() should return empty, got %q", result)
	}
}

func Test_EmptyPattern_IsUndefined(t *testing.T) {
	// Arrange
	lazy := regexnew.New.Lazy("")

	// Act & Assert
	if !lazy.IsUndefined() {
		t.Error("empty pattern LazyRegex should be undefined")
	}
}

func Test_EmptyPattern_IsNotApplicable(t *testing.T) {
	// Arrange
	lazy := regexnew.New.Lazy("")

	// Act & Assert
	if lazy.IsApplicable() {
		t.Error("empty pattern LazyRegex should not be applicable")
	}
}

func Test_EmptyPattern_IsMatch_ReturnsFalse(t *testing.T) {
	// Arrange
	lazy := regexnew.New.Lazy("")

	// Act & Assert
	if lazy.IsMatch("anything") {
		t.Error("empty pattern LazyRegex IsMatch should return false")
	}
}

func Test_EmptyPattern_IsFailedMatch_ReturnsFalse(t *testing.T) {
	// Arrange
	lazy := regexnew.New.Lazy("")

	// Act
	// IsFailedMatch on undefined returns true (regEx == nil)
	result := lazy.IsFailedMatch("anything")

	// Assert
	if !result {
		t.Error("empty pattern LazyRegex IsFailedMatch should return true")
	}
}

func Test_EmptyPattern_Compile_ReturnsError(t *testing.T) {
	// Arrange
	lazy := regexnew.New.Lazy("")

	// Act
	regex, err := lazy.Compile()

	// Assert
	if err == nil {
		t.Error("empty pattern Compile should return error")
	}
	if regex != nil {
		t.Error("empty pattern Compile should return nil regex")
	}
}

func Test_InvalidPattern_ConcurrentAccess(t *testing.T) {
	// Arrange
	invalidPatterns := []string{"[bad", "(unclosed", "*invalid", "(?P<>bad)"}
	goroutineCount := 50
	wg := sync.WaitGroup{}
	wg.Add(goroutineCount)
	errors := make(chan string, goroutineCount)

	// Act
	for i := 0; i < goroutineCount; i++ {
		go func(index int) {
			defer wg.Done()

			p := invalidPatterns[index%len(invalidPatterns)]
			lazy := regexnew.New.LazyLock(p)

			if lazy == nil {
				errors <- fmt.Sprintf("goroutine %d: LazyLock returned nil for invalid pattern %s", index, p)
				return
			}

			if lazy.IsApplicable() {
				errors <- fmt.Sprintf("goroutine %d: invalid pattern %s should not be applicable", index, p)
			}

			if lazy.IsMatch("test") {
				errors <- fmt.Sprintf("goroutine %d: invalid pattern %s should not match", index, p)
			}

			if !lazy.HasAnyIssues() {
				errors <- fmt.Sprintf("goroutine %d: invalid pattern %s should have issues", index, p)
			}
		}(i)
	}

	wg.Wait()
	close(errors)

	// Assert
	for errMsg := range errors {
		t.Error(errMsg)
	}
}

func Test_InvalidPattern_ConcurrentCompileError(t *testing.T) {
	// Arrange
	pattern := "[broken"
	goroutineCount := 50
	wg := sync.WaitGroup{}
	wg.Add(goroutineCount)
	errors := make(chan string, goroutineCount)

	// Act
	for i := 0; i < goroutineCount; i++ {
		go func(index int) {
			defer wg.Done()

			lazy := regexnew.New.LazyLock(pattern)
			regex, err := lazy.Compile()

			if err == nil {
				errors <- fmt.Sprintf("goroutine %d: expected compile error", index)
			}

			if regex != nil {
				errors <- fmt.Sprintf("goroutine %d: expected nil regex", index)
			}

			if !lazy.HasError() {
				errors <- fmt.Sprintf("goroutine %d: HasError should be true", index)
			}
		}(i)
	}

	wg.Wait()
	close(errors)

	// Assert
	for errMsg := range errors {
		t.Error(errMsg)
	}
}

func Test_MixedValidInvalid_ConcurrentAccess(t *testing.T) {
	// Arrange
	patterns := []string{`\d+`, "[bad", `[a-z]+`, "(unclosed"}
	goroutineCount := 80
	wg := sync.WaitGroup{}
	wg.Add(goroutineCount)
	errors := make(chan string, goroutineCount)

	// Act
	for i := 0; i < goroutineCount; i++ {
		go func(index int) {
			defer wg.Done()

			p := patterns[index%len(patterns)]
			lazy := regexnew.New.LazyLock(p)
			isValid := (index%len(patterns))%2 == 0 // even indices are valid

			if isValid {
				if !lazy.IsApplicable() {
					errors <- fmt.Sprintf("goroutine %d: valid pattern %s should be applicable", index, p)
				}
			} else {
				if lazy.IsApplicable() {
					errors <- fmt.Sprintf("goroutine %d: invalid pattern %s should not be applicable", index, p)
				}
			}
		}(i)
	}

	wg.Wait()
	close(errors)

	// Assert
	for errMsg := range errors {
		t.Error(errMsg)
	}
}
