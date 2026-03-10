package corejsontests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coreutils/stringutil"
	"github.com/alimtvnetwork/core/errcore"
)

func Test_FromTo(t *testing.T) {
	tc := deserializerFromToTestCase

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
		exampleTo,
	)

	errcore.HandleErr(err)

	to := stringutil.AnyToStringNameField(exampleTo)
	from := stringutil.AnyToStringNameField(exampleFrom)

	// Assert
	actLines := []string{fmt.Sprintf("%v", to == from)}

	tc.ShouldBeEqual(t, 0, actLines...)
}
