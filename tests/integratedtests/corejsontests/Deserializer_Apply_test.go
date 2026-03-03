package corejsontests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/coreutils/stringutil"
	"gitlab.com/auk-go/core/errcore"
)

func Test_FromTo(t *testing.T) {
	type Example struct {
		A       string
		B       int
		SomeMap map[string]string
	}

	exampleFrom := &Example{
		A:       "Something",
		B:       1,
		SomeMap: map[string]string{},
	}

	exampleTo := &Example{}

	err := corejson.Deserialize.FromTo(
		exampleFrom,
		exampleTo)

	errcore.HandleErr(err)

	to := stringutil.AnyToStringNameField(exampleTo)
	from := stringutil.AnyToStringNameField(exampleFrom)

	// Assert
	actLines := []string{fmt.Sprintf("%v", to == from)}
	expected := []string{"true"}

	errcore.PrintLineDiff(0, "corejson.Deserializer.FromTo - should match from to casting", actLines, expected)

	if to != from {
		t.Errorf("corejson.Deserializer.FromTo: got %q, want %q", to, from)
	}
}
