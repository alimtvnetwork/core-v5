package corestr

import "fmt"

type utils struct{}

func (it utils) WrapDoubleIfMissing(s string) string {
	if s == "" || s == "\"\"" {
		return "\"\""
	}

	if s[0] == '"' && s[len(s)-1] == '"' {
		return s
	}

	return fmt.Sprintf("\"%s\"", s)
}

func (it utils) WrapSingleIfMissing(s string) string {
	if s == "" || s == "''" {
		return "''"
	}

	if s[0] == '\'' && s[len(s)-1] == '\'' {
		return s
	}

	return fmt.Sprintf("'%s'", s)
}

func (it utils) WrapDouble(s string) string {
	return fmt.Sprintf("\"%s\"", s)
}

func (it utils) WrapSingle(s string) string {
	return fmt.Sprintf("'%s'", s)
}
