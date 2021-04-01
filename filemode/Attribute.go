package filemode

import (
	"gitlab.com/evatix-go/core/conditional"
	"gitlab.com/evatix-go/core/constants"
)

type Attribute struct {
	IsRead    bool
	IsWrite   bool
	IsExecute bool
}

func (attribute Attribute) ToAttributeValue() AttributeValue {
	read, write, exe, sum := attribute.ToSpecificBytes()

	return AttributeValue{
		Read:    read,
		Write:   write,
		Execute: exe,
		Sum:     sum,
	}
}

func (attribute Attribute) ToSpecificBytes() (read, write, exe, sum byte) {
	read = conditional.Byte(attribute.IsRead, readValue, constants.Zero)
	write = conditional.Byte(attribute.IsWrite, writeValue, constants.Zero)
	exe = conditional.Byte(attribute.IsExecute, executeValue, constants.Zero)

	return read, write, exe, read + write + exe
}

func (attribute Attribute) ToByte() byte {
	r := conditional.Byte(attribute.IsRead, readValue, constants.Zero)
	w := conditional.Byte(attribute.IsWrite, writeValue, constants.Zero)
	e := conditional.Byte(attribute.IsExecute, executeValue, constants.Zero)

	return r + w + e
}

func (attribute Attribute) ToSum() byte {
	return attribute.ToByte()
}

func (attribute Attribute) ToRwx() [3]byte {
	return [3]byte{
		conditional.Byte(attribute.IsRead, readChar, constants.HyphenChar),
		conditional.Byte(attribute.IsWrite, writeChar, constants.HyphenChar),
		conditional.Byte(attribute.IsExecute, executeChar, constants.HyphenChar),
	}
}

func (attribute Attribute) ToRwxString() string {
	rwxBytes := attribute.ToRwx()

	return string(rwxBytes[:])
}

func (attribute Attribute) ToVariant() AttrVariant {
	b := attribute.ToByte()

	return AttrVariant(b)
}

func (attribute Attribute) ToChar() byte {
	return attribute.ToByte() + constants.ZeroChar
}
