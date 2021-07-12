package issetter

import (
	"gitlab.com/evatix-go/core/simplewrap"
)

var (
	values = []string{"Uninitialized", "True", "False", "Unset", "Set", "Wildcard"}

	jsonValuesMap = map[string]Value{
		simplewrap.WithDoubleQuote("0"):        Uninitialized,
		simplewrap.WithDoubleQuote(""):         Uninitialized,
		simplewrap.WithDoubleQuote("-"):        Uninitialized,
		simplewrap.WithDoubleQuote("-1"):       Uninitialized,
		simplewrap.WithDoubleQuote("1"):        True,
		simplewrap.WithDoubleQuote("yes"):      True,
		simplewrap.WithDoubleQuote("Yes"):      True,
		simplewrap.WithDoubleQuote("true"):     True,
		simplewrap.WithDoubleQuote("True"):     True,
		simplewrap.WithDoubleQuote("no"):       False,
		simplewrap.WithDoubleQuote("No"):       False,
		simplewrap.WithDoubleQuote("Nop"):      False,
		simplewrap.WithDoubleQuote("None"):     False,
		simplewrap.WithDoubleQuote("false"):    False,
		simplewrap.WithDoubleQuote("False"):    False,
		simplewrap.WithDoubleQuote("set"):      Set,
		simplewrap.WithDoubleQuote("Set"):      Set,
		simplewrap.WithDoubleQuote("Unset"):    Unset,
		simplewrap.WithDoubleQuote("unset"):    Unset,
		simplewrap.WithDoubleQuote("*"):        Wildcard,
		simplewrap.WithDoubleQuote("%"):        Wildcard,
		simplewrap.WithDoubleQuote("Wildcard"): Wildcard,
		simplewrap.WithDoubleQuote("WildCard"): Wildcard,
		simplewrap.WithDoubleQuote("wildcard"): Wildcard, // all small
	}

	valuesToJsonBytesMap = map[Value][]byte{
		Uninitialized: jsonBytes("Uninitialized"),
		True:          jsonBytes("True"),
		False:         jsonBytes("False"),
		Unset:         jsonBytes("Unset"),
		Set:           jsonBytes("Set"),
		Wildcard:      jsonBytes("Wildcard"),
	}

	valuesToNameMap = map[Value]string{
		Uninitialized: "Uninitialized",
		True:          "True",
		False:         "False",
		Unset:         "Unset",
		Set:           "Set",
		Wildcard:      "Wildcard",
	}
)
