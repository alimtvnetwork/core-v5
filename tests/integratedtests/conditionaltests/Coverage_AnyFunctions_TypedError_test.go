package conditionaltests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/conditional"
)

func Test_Cov_AnyFunctionsExecuteResults_False(t *testing.T) {
	result := conditional.AnyFunctionsExecuteResults(
		false,
		nil,
		[]func() (any, bool, bool){
			func() (any, bool, bool) { return "b", true, false },
		},
	)
	if len(result) != 1 {
		t.Errorf("expected 1 got %d", len(result))
	}
}

func Test_Cov_TypedErrorFunctionsExecuteResults_False(t *testing.T) {
	results, err := conditional.TypedErrorFunctionsExecuteResults[int](
		false,
		nil,
		[]func() (int, error){
			func() (int, error) { return 42, nil },
		},
	)
	if err != nil || len(results) != 1 || results[0] != 42 {
		t.Error("expected 42")
	}
}

func Test_Cov_TypedErrorFunctionsExecuteResults_WithError(t *testing.T) {
	results, err := conditional.TypedErrorFunctionsExecuteResults[int](
		true,
		[]func() (int, error){
			func() (int, error) { return 0, errors.New("fail") },
			nil,
			func() (int, error) { return 1, nil },
		},
		nil,
	)
	if err == nil {
		t.Error("expected error")
	}
	if len(results) != 1 {
		t.Errorf("expected 1 got %d", len(results))
	}
}

func Test_Cov_TypedErrorFunctionsExecuteResults_Empty(t *testing.T) {
	results, err := conditional.TypedErrorFunctionsExecuteResults[int](
		true,
		nil,
		nil,
	)
	if err != nil || results != nil {
		t.Error("expected nil")
	}
}
