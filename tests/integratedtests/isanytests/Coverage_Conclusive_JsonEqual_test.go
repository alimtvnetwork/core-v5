package isanytests

import (
	"testing"

	"github.com/alimtvnetwork/core/isany"
)

func Test_Cov_Conclusive_BothReflectNull(t *testing.T) {
	var p1 *string
	var p2 *string
	isEqual, isConclusive := isany.Conclusive(p1, p2)
	if !isEqual || !isConclusive {
		t.Error("both nil ptr should be equal conclusive")
	}
}

func Test_Cov_Conclusive_OneReflectNull(t *testing.T) {
	var p1 *string
	s := "hello"
	isEqual, isConclusive := isany.Conclusive(p1, &s)
	if isEqual || !isConclusive {
		t.Error("one nil should be not equal but conclusive")
	}
}

func Test_Cov_Conclusive_DiffTypes(t *testing.T) {
	isEqual, isConclusive := isany.Conclusive(42, "hello")
	if isEqual || !isConclusive {
		t.Error("diff types should be not equal but conclusive")
	}
}

func Test_Cov_Conclusive_Inconclusive(t *testing.T) {
	a := "hello"
	b := "world"
	isEqual, isConclusive := isany.Conclusive(&a, &b)
	if isEqual || isConclusive {
		t.Error("same type diff values should be inconclusive")
	}
}

func Test_Cov_JsonEqual_IntEqual(t *testing.T) {
	if !isany.JsonEqual(42, 42) {
		t.Error("expected equal")
	}
	if isany.JsonEqual(42, 43) {
		t.Error("expected not equal")
	}
}

func Test_Cov_JsonEqual_JsonMarshal(t *testing.T) {
	a := map[string]int{"a": 1}
	b := map[string]int{"a": 1}
	if !isany.JsonEqual(a, b) {
		t.Error("expected equal")
	}
}
