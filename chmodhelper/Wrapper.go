package chmodhelper

import (
	"bytes"
	"os"
	"os/exec"
	"strconv"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/constants/bitsize"
	"gitlab.com/evatix-go/core/internal/fsinternal"
	"gitlab.com/evatix-go/core/internal/messages"
	"gitlab.com/evatix-go/core/msgtype"
)

type Wrapper struct {
	Owner, Group, Other Attribute
}

func (wrapper Wrapper) Verify(location string) error {
	return VerifyChmod(location, wrapper.ToHyphenedRwx())
}

func (wrapper Wrapper) VerifyPaths(location *[]string, isContinueOnError bool) error {
	return VerifyChmodPaths(
		location,
		wrapper.ToHyphenedRwx(),
		isContinueOnError)
}

func (wrapper Wrapper) HasChmod(location string) bool {
	return IsChmod(location, wrapper.ToHyphenedRwx())
}

// Bytes return rwx, (Owner)(Group)(Other) byte values under 1-7
func (wrapper Wrapper) Bytes() [3]byte {
	// # https://play.golang.org/p/dX-wsvJmFie
	owner := wrapper.Owner.ToSum()
	group := wrapper.Group.ToSum()
	other := wrapper.Other.ToSum()

	allBytes := [3]byte{owner, group, other}

	return allBytes
}

func (wrapper Wrapper) ToUint32Octal() uint32 {
	// # https://play.golang.org/p/dX-wsvJmFie
	str := wrapper.ToFileModeString()

	// # https://bit.ly/35aBepk
	octal, err := strconv.ParseUint(str, bitsize.Of8, bitsize.Of32)

	if err != nil {
		msgtype.
			MeaningFulErrorHandle(
				msgtype.PathChmodConvertFailedMessage,
				"ToUint32Octal",
				err)
	}

	return uint32(octal)
}

// Chars return 0rwx, '0'(Owner + '0')(Group + '0')(Other + '0')
// eg. 0777, 0555, 0755 NOT 0rwx
func (wrapper Wrapper) Chars() [4]byte {
	// # https://play.golang.org/p/dX-wsvJmFie
	owner := wrapper.Owner.ToChar()
	group := wrapper.Group.ToChar()
	other := wrapper.Other.ToChar()

	allBytes := [4]byte{constants.ZeroChar, owner, group, other}

	return allBytes
}

// ToFileModeString 4 digit string 0rwx, example 0777
func (wrapper Wrapper) ToFileModeString() string {
	// # https://play.golang.org/p/dX-wsvJmFie
	allBytes := wrapper.Chars()

	return string(allBytes[:])
}

// ToModeStr 3 digit string, example 777
func (wrapper Wrapper) ToModeStr() string {
	// # https://play.golang.org/p/dX-wsvJmFie
	allBytes := wrapper.Chars()

	return string(allBytes[1:])
}

// ToHyphenedRwx returns "-rwxrwxrwx"
func (wrapper Wrapper) ToHyphenedRwx() string {
	owner := wrapper.Owner.ToRwxString()
	group := wrapper.Group.ToRwxString()
	other := wrapper.Other.ToRwxString()

	// # https://ss64.com/bash/chmod.html, needs to be 10 always
	return constants.Hyphen + owner + group + other
}

func (wrapper Wrapper) ToHyphenedRwxChars() []byte {
	str := wrapper.ToHyphenedRwx()
	chars := []byte(str)

	return chars
}

func (wrapper Wrapper) String() string {
	// # https://ss64.com/bash/chmod.html, needs to be 10 always
	return wrapper.ToHyphenedRwx()
}

func (wrapper Wrapper) ToFileMode() os.FileMode {
	// # https://play.golang.org/p/dX-wsvJmFie
	octalUint32 := wrapper.ToUint32Octal()

	return os.FileMode(octalUint32)
}

func (wrapper Wrapper) ApplyChmod(
	fileOrDirectoryPath string,
	isSkipOnNonExist bool,
) error {
	isFileExist := fsinternal.IsPathExists(fileOrDirectoryPath)

	if isSkipOnNonExist && !isFileExist {
		return nil
	}

	if !isSkipOnNonExist && !isFileExist {
		return msgtype.
			PathInvalidErrorMessage.
			Error(
				messages.PathNotExist, fileOrDirectoryPath)
	}

	err := os.Chmod(fileOrDirectoryPath, wrapper.ToFileMode())

	if err != nil {
		return msgtype.
			PathChmodApplyMessage.
			Error(err.Error(), fileOrDirectoryPath)
	}

	return nil
}

// UnixApplyRecursive skip if it is a non dir path
func (wrapper Wrapper) UnixApplyRecursive(
	dirPath string,
	isSkipOnNonExist bool,
) error {
	isFileExist := fsinternal.IsPathExists(dirPath)

	if isSkipOnNonExist && !isFileExist {
		return nil
	}

	if !isSkipOnNonExist && !isFileExist {
		return msgtype.
			PathInvalidErrorMessage.
			Error(
				"Path doesn't exist", dirPath)
	}

	isDir := fsinternal.IsDirectory(dirPath)

	if isDir {
		return wrapper.applyUnixRecursiveChmodUsingCmd(
			dirPath)
	}

	return nil
}

func (wrapper Wrapper) applyUnixRecursiveChmodUsingCmd(dirPath string) error {
	cmd := wrapper.getUnixRecursiveCmdForChmod(dirPath)

	if cmd == nil {
		return msgtype.
			FailedToCreateCmd.Error(
			constants.BashCommandline,
			dirPath)
	}

	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()

	if err != nil {
		return msgtype.
			FailedToCreateCmd.Error(
			constants.ChmodCommand,
			err.Error()+constants.NewLineUnix+stderr.String())
	}

	return nil
}

func (wrapper Wrapper) getUnixRecursiveCmdForChmod(dirPath string) *exec.Cmd {
	instructionLine := constants.ChmodCommand +
		constants.Space +
		constants.RecursiveCommandFlag +
		constants.Space +
		wrapper.ToModeStr() +
		constants.Space +
		dirPath

	return exec.Command(
		constants.BinShellCmd,
		constants.NonInteractiveFlag,
		instructionLine)
}

func (wrapper Wrapper) MustApplyChmod(fileOrDirectoryPath string) {
	err := os.Chmod(fileOrDirectoryPath, wrapper.ToFileMode())

	if err != nil {
		panic(err)
	}
}
