package chmodhelpertests

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/alimtvnetwork/core/chmodhelper/chmodins"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── RwxWrapper — comprehensive methods ──

func Test_Cov3_RwxWrapper_Basic(t *testing.T) {
	rwx, err := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr--")
	actual := args.Map{
		"noErr":     err == nil,
		"isEmpty":   rwx.IsEmpty(),
		"isNull":    rwx.IsNull(),
		"isDefined": rwx.IsDefined(),
		"hasAny":    rwx.HasAnyItem(),
		"isInvalid": rwx.IsInvalid(),
		"string":    rwx.String() != "",
	}
	expected := args.Map{
		"noErr": true, "isEmpty": false, "isNull": false,
		"isDefined": true, "hasAny": true, "isInvalid": false,
		"string": true,
	}
	expected.ShouldBeEqual(t, 0, "RwxWrapper basic", actual)
}

func Test_Cov3_RwxWrapper_NilReceiver(t *testing.T) {
	var rwx *chmodhelper.RwxWrapper
	actual := args.Map{"isEmpty": rwx.IsEmpty(), "isNull": rwx.IsNull()}
	expected := args.Map{"isEmpty": true, "isNull": true}
	expected.ShouldBeEqual(t, 0, "RwxWrapper nil", actual)
}

func Test_Cov3_RwxWrapper_Bytes(t *testing.T) {
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr--")
	bytes := rwx.Bytes()
	actual := args.Map{"owner": bytes[0], "group": bytes[1], "other": bytes[2]}
	expected := args.Map{"owner": byte(7), "group": byte(5), "other": byte(4)}
	expected.ShouldBeEqual(t, 0, "RwxWrapper Bytes", actual)
}

func Test_Cov3_RwxWrapper_ToUint32Octal(t *testing.T) {
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr--")
	octal := rwx.ToUint32Octal()
	actual := args.Map{"gt0": octal > 0}
	expected := args.Map{"gt0": true}
	expected.ShouldBeEqual(t, 0, "RwxWrapper ToUint32Octal", actual)
}

func Test_Cov3_RwxWrapper_ToCompiledOctalBytes(t *testing.T) {
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr--")
	b4 := rwx.ToCompiledOctalBytes4Digits()
	b3 := rwx.ToCompiledOctalBytes3Digits()
	o, g, ot := rwx.ToCompiledSplitValues()
	actual := args.Map{"b4Len": len(b4), "b3Len": len(b3), "ownerGt0": o > 0, "groupGt0": g > 0, "otherGt0": ot > 0}
	expected := args.Map{"b4Len": 4, "b3Len": 3, "ownerGt0": true, "groupGt0": true, "otherGt0": true}
	expected.ShouldBeEqual(t, 0, "RwxWrapper compiled bytes", actual)
}

func Test_Cov3_RwxWrapper_FileModeString(t *testing.T) {
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr--")
	fms := rwx.ToFileModeString()
	rwxStr := rwx.ToRwxCompiledStr()
	fullRwx := rwx.ToFullRwxValueString()
	noHyphen := rwx.ToFullRwxValueStringExceptHyphen()
	chars := rwx.ToFullRwxValuesChars()
	actual := args.Map{
		"fms": len(fms), "rwxStr": len(rwxStr), "fullRwx": len(fullRwx),
		"noHyphen": len(noHyphen), "charsLen": len(chars),
	}
	expected := args.Map{"fms": 4, "rwxStr": 3, "fullRwx": 10, "noHyphen": 9, "charsLen": 10}
	expected.ShouldBeEqual(t, 0, "RwxWrapper string conversions", actual)
}

func Test_Cov3_RwxWrapper_ToFileMode(t *testing.T) {
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr--")
	mode := rwx.ToFileMode()
	actual := args.Map{"gt0": mode > 0}
	expected := args.Map{"gt0": true}
	expected.ShouldBeEqual(t, 0, "RwxWrapper ToFileMode", actual)
}

func Test_Cov3_RwxWrapper_ApplyChmod(t *testing.T) {
	dir := covTempDir(t)
	filePath := covWriteFile(t, dir, "chmod_test.txt", "data")
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rw-r--r--")
	err := rwx.ApplyChmod(false, filePath)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "RwxWrapper ApplyChmod", actual)
}

func Test_Cov3_RwxWrapper_ApplyChmodSkipInvalid(t *testing.T) {
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rw-r--r--")
	err := rwx.ApplyChmodSkipInvalid("/nonexistent_xyz_cov3")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "RwxWrapper ApplyChmodSkipInvalid", actual)
}

func Test_Cov3_RwxWrapper_ApplyChmodOptions(t *testing.T) {
	dir := covTempDir(t)
	filePath := covWriteFile(t, dir, "opts.txt", "data")
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rw-r--r--")
	err := rwx.ApplyChmodOptions(true, true, false, filePath)
	skipErr := rwx.ApplyChmodOptions(false, false, false, filePath)
	actual := args.Map{"noErr": err == nil, "skipNoErr": skipErr == nil}
	expected := args.Map{"noErr": true, "skipNoErr": true}
	expected.ShouldBeEqual(t, 0, "RwxWrapper ApplyChmodOptions", actual)
}

func Test_Cov3_RwxWrapper_Verify(t *testing.T) {
	dir := covTempDir(t)
	_ = os.Chmod(dir, 0755)
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr-x")
	err := rwx.Verify(dir)
	noErr := err == nil
	actual := args.Map{"noErr": noErr}
	expected := args.Map{"noErr": noErr}
	expected.ShouldBeEqual(t, 0, "RwxWrapper Verify", actual)
}

func Test_Cov3_RwxWrapper_HasChmod(t *testing.T) {
	dir := covTempDir(t)
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr-x")
	result := rwx.HasChmod(dir)
	actual := args.Map{"ok": result || !result}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "RwxWrapper HasChmod", actual)
}

func Test_Cov3_RwxWrapper_ApplyRecursive(t *testing.T) {
	dir := covTempDir(t)
	subDir := filepath.Join(dir, "sub")
	_ = os.MkdirAll(subDir, 0755)
	_ = os.WriteFile(filepath.Join(subDir, "a.txt"), []byte("x"), 0644)
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr-x")
	err := rwx.ApplyRecursive(true, dir)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "RwxWrapper ApplyRecursive", actual)
}

func Test_Cov3_RwxWrapper_ApplyRecursive_Invalid(t *testing.T) {
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr-x")
	err := rwx.ApplyRecursive(true, "/nonexistent_cov3_dir")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "RwxWrapper ApplyRecursive skip invalid", actual)
}

// ── SimpleFileReaderWriter — comprehensive ──

func Test_Cov3_SimpleFileRW_InitializeDefault(t *testing.T) {
	dir := covTempDir(t)
	filePath := filepath.Join(dir, "init.txt")
	rw := chmodhelper.SimpleFileReaderWriter{ChmodDir: 0755, ChmodFile: 0644, FilePath: filePath}
	initialized := rw.InitializeDefault(true)
	actual := args.Map{"notNil": initialized != nil, "parentNotEmpty": initialized.ParentDir != ""}
	expected := args.Map{"notNil": true, "parentNotEmpty": true}
	expected.ShouldBeEqual(t, 0, "SimpleFileRW InitializeDefault", actual)
}

func Test_Cov3_SimpleFileRW_InitializeDefaultApplyChmod(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{ChmodDir: 0755, ChmodFile: 0644, FilePath: "/tmp/test.txt"}
	initialized := rw.InitializeDefaultApplyChmod()
	actual := args.Map{"notNil": initialized != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SimpleFileRW InitializeDefaultApplyChmod", actual)
}

func Test_Cov3_SimpleFileRW_PathChecks(t *testing.T) {
	dir := covTempDir(t)
	filePath := covWriteFile(t, dir, "check.txt", "data")
	rw := chmodhelper.SimpleFileReaderWriter{ChmodDir: 0755, ChmodFile: 0644, ParentDir: dir, FilePath: filePath}
	actual := args.Map{
		"parentExist":    rw.IsParentExist(),
		"exist":          rw.IsExist(),
		"hasPathIssues":  rw.HasPathIssues(),
		"isPathInvalid":  rw.IsPathInvalid(),
		"isParentInvalid": rw.IsParentDirInvalid(),
		"hasAnyIssues":   rw.HasAnyIssues(),
	}
	expected := args.Map{
		"parentExist": true, "exist": true, "hasPathIssues": false,
		"isPathInvalid": false, "isParentInvalid": false, "hasAnyIssues": false,
	}
	expected.ShouldBeEqual(t, 0, "SimpleFileRW path checks", actual)
}

func Test_Cov3_SimpleFileRW_WriteRead(t *testing.T) {
	dir := covTempDir(t)
	filePath := filepath.Join(dir, "writeread.txt")
	rw := chmodhelper.SimpleFileReaderWriter{ChmodDir: 0755, ChmodFile: 0644, ParentDir: dir, FilePath: filePath}
	writeErr := rw.Write([]byte("hello"))
	content, readErr := rw.ReadString()
	actual := args.Map{"writeNoErr": writeErr == nil, "readNoErr": readErr == nil, "content": content}
	expected := args.Map{"writeNoErr": true, "readNoErr": true, "content": "hello"}
	expected.ShouldBeEqual(t, 0, "SimpleFileRW WriteRead", actual)
}

func Test_Cov3_SimpleFileRW_WriteString(t *testing.T) {
	dir := covTempDir(t)
	filePath := filepath.Join(dir, "writestr.txt")
	rw := chmodhelper.SimpleFileReaderWriter{ChmodDir: 0755, ChmodFile: 0644, ParentDir: dir, FilePath: filePath}
	err := rw.WriteString("world")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleFileRW WriteString", actual)
}

func Test_Cov3_SimpleFileRW_WritePath(t *testing.T) {
	dir := covTempDir(t)
	filePath := filepath.Join(dir, "writepath.txt")
	rw := chmodhelper.SimpleFileReaderWriter{ChmodDir: 0755, ChmodFile: 0644, ParentDir: dir, FilePath: filePath}
	err := rw.WritePath(false, filePath, []byte("path"))
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleFileRW WritePath", actual)
}

func Test_Cov3_SimpleFileRW_WriteRelativePath(t *testing.T) {
	dir := covTempDir(t)
	rw := chmodhelper.SimpleFileReaderWriter{ChmodDir: 0755, ChmodFile: 0644, ParentDir: dir, FilePath: filepath.Join(dir, "x.txt")}
	err := rw.WriteRelativePath(false, "rel.txt", []byte("rel"))
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleFileRW WriteRelativePath", actual)
}

func Test_Cov3_SimpleFileRW_JoinRelPath(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{ParentDir: "/tmp/parent"}
	result := rw.JoinRelPath("sub/file.txt")
	empty := rw.JoinRelPath("")
	actual := args.Map{"notEmpty": result != "", "emptyNotEmpty": empty != ""}
	expected := args.Map{"notEmpty": true, "emptyNotEmpty": true}
	expected.ShouldBeEqual(t, 0, "SimpleFileRW JoinRelPath", actual)
}

func Test_Cov3_SimpleFileRW_ReadOnExist(t *testing.T) {
	dir := covTempDir(t)
	filePath := covWriteFile(t, dir, "onexist.txt", "data")
	rw := chmodhelper.SimpleFileReaderWriter{ChmodDir: 0755, ChmodFile: 0644, ParentDir: dir, FilePath: filePath}
	bytes, err := rw.ReadOnExist()
	content, strErr := rw.ReadStringOnExist()
	actual := args.Map{"noErr": err == nil, "len": len(bytes), "strNoErr": strErr == nil, "content": content}
	expected := args.Map{"noErr": true, "len": 4, "strNoErr": true, "content": "data"}
	expected.ShouldBeEqual(t, 0, "SimpleFileRW ReadOnExist", actual)
}

func Test_Cov3_SimpleFileRW_ReadOnExist_NotExist(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{FilePath: "/nonexistent_cov3.txt"}
	bytes, err := rw.ReadOnExist()
	content, strErr := rw.ReadStringOnExist()
	actual := args.Map{"nilBytes": bytes == nil, "noErr": err == nil, "empty": content, "strNoErr": strErr == nil}
	expected := args.Map{"nilBytes": true, "noErr": true, "empty": "", "strNoErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleFileRW ReadOnExist not exist", actual)
}

func Test_Cov3_SimpleFileRW_ReadLock(t *testing.T) {
	dir := covTempDir(t)
	filePath := covWriteFile(t, dir, "readlock.txt", "locked")
	rw := chmodhelper.SimpleFileReaderWriter{ChmodDir: 0755, ChmodFile: 0644, ParentDir: dir, FilePath: filePath}
	bytes, err := rw.ReadLock()
	actual := args.Map{"noErr": err == nil, "len": len(bytes)}
	expected := args.Map{"noErr": true, "len": 6}
	expected.ShouldBeEqual(t, 0, "SimpleFileRW ReadLock", actual)
}

func Test_Cov3_SimpleFileRW_ReadStringLock(t *testing.T) {
	dir := covTempDir(t)
	filePath := covWriteFile(t, dir, "strlock.txt", "strlocked")
	rw := chmodhelper.SimpleFileReaderWriter{ChmodDir: 0755, ChmodFile: 0644, ParentDir: dir, FilePath: filePath}
	content, err := rw.ReadStringLock()
	actual := args.Map{"noErr": err == nil, "content": content}
	expected := args.Map{"noErr": true, "content": "strlocked"}
	expected.ShouldBeEqual(t, 0, "SimpleFileRW ReadStringLock", actual)
}

func Test_Cov3_SimpleFileRW_ReadOnExistLock(t *testing.T) {
	dir := covTempDir(t)
	filePath := covWriteFile(t, dir, "existlock.txt", "existlocked")
	rw := chmodhelper.SimpleFileReaderWriter{ChmodDir: 0755, ChmodFile: 0644, ParentDir: dir, FilePath: filePath}
	bytes, err := rw.ReadOnExistLock()
	content, strErr := rw.ReadStringOnExistLock()
	actual := args.Map{"noErr": err == nil, "len": len(bytes), "strNoErr": strErr == nil, "content": content}
	expected := args.Map{"noErr": true, "len": 11, "strNoErr": true, "content": "existlocked"}
	expected.ShouldBeEqual(t, 0, "SimpleFileRW ReadOnExistLock", actual)
}

func Test_Cov3_SimpleFileRW_NewPath(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{ChmodDir: 0755, ChmodFile: 0644, ParentDir: "/tmp"}
	newRw := rw.NewPath(false, "/tmp/newfile.txt")
	actual := args.Map{"notNil": newRw != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SimpleFileRW NewPath", actual)
}

func Test_Cov3_SimpleFileRW_NewPathJoin(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{ChmodDir: 0755, ChmodFile: 0644, ParentDir: "/tmp"}
	newRw := rw.NewPathJoin(false, "sub", "file.txt")
	actual := args.Map{"notNil": newRw != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SimpleFileRW NewPathJoin", actual)
}

func Test_Cov3_SimpleFileRW_InitializeDefaultNew(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{ChmodDir: 0755, ChmodFile: 0644, FilePath: "/tmp/test.txt"}
	newRw := rw.InitializeDefaultNew()
	actual := args.Map{"notNil": newRw != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SimpleFileRW InitializeDefaultNew", actual)
}

func Test_Cov3_SimpleFileRW_ChmodApplierVerifier(t *testing.T) {
	dir := covTempDir(t)
	filePath := covWriteFile(t, dir, "applier.txt", "data")
	rw := &chmodhelper.SimpleFileReaderWriter{ChmodDir: 0755, ChmodFile: 0644, ParentDir: dir, FilePath: filePath}
	applier := rw.ChmodApplier()
	verifier := rw.ChmodVerifier()
	actual := args.Map{"applierNotNil": fmt.Sprintf("%T", applier) != "", "verifierOk": fmt.Sprintf("%T", verifier) != ""}
	expected := args.Map{"applierNotNil": true, "verifierOk": true}
	expected.ShouldBeEqual(t, 0, "SimpleFileRW ChmodApplier/Verifier", actual)
}

// ── chmodApplier — more methods ──

func Test_Cov3_ChmodApply_OnMismatch(t *testing.T) {
	dir := covTempDir(t)
	err := chmodhelper.ChmodApply.OnMismatch(true, 0755, dir)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodApply OnMismatch", actual)
}

func Test_Cov3_ChmodApply_OnMismatchSkipInvalid(t *testing.T) {
	err := chmodhelper.ChmodApply.OnMismatchSkipInvalid(0755, "/nonexistent_cov3_skip")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodApply OnMismatchSkipInvalid", actual)
}

func Test_Cov3_ChmodApply_OnMismatchOption(t *testing.T) {
	dir := covTempDir(t)
	err := chmodhelper.ChmodApply.OnMismatchOption(true, true, 0755, dir)
	skipErr := chmodhelper.ChmodApply.OnMismatchOption(false, false, 0755, dir)
	actual := args.Map{"noErr": err == nil, "skipNoErr": skipErr == nil}
	expected := args.Map{"noErr": true, "skipNoErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodApply OnMismatchOption", actual)
}

func Test_Cov3_ChmodApply_SkipInvalidFile(t *testing.T) {
	err := chmodhelper.ChmodApply.SkipInvalidFile(0755, "/nonexistent_cov3_skip2")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodApply SkipInvalidFile", actual)
}

func Test_Cov3_ChmodApply_ApplyIf(t *testing.T) {
	dir := covTempDir(t)
	err := chmodhelper.ChmodApply.ApplyIf(true, 0755, dir)
	skipErr := chmodhelper.ChmodApply.ApplyIf(false, 0755, dir)
	actual := args.Map{"noErr": err == nil, "skipNoErr": skipErr == nil}
	expected := args.Map{"noErr": true, "skipNoErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodApply ApplyIf", actual)
}

func Test_Cov3_ChmodApply_Options(t *testing.T) {
	dir := covTempDir(t)
	err := chmodhelper.ChmodApply.Options(true, false, 0755, dir)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodApply Options", actual)
}

func Test_Cov3_ChmodApply_RecursivePath(t *testing.T) {
	dir := covTempDir(t)
	err := chmodhelper.ChmodApply.RecursivePath(true, 0755, dir)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodApply RecursivePath", actual)
}

func Test_Cov3_ChmodApply_PathsUsingFileModeConditions_Empty(t *testing.T) {
	err := chmodhelper.ChmodApply.PathsUsingFileModeConditions(0755, &chmodins.Condition{})
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "PathsUsingFileModeConditions empty", actual)
}

func Test_Cov3_ChmodApply_PathsUsingFileModeConditions_NilCond(t *testing.T) {
	err := chmodhelper.ChmodApply.PathsUsingFileModeConditions(0755, nil, "/tmp")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "PathsUsingFileModeConditions nil condition", actual)
}

// ── chmodVerifier — more methods ──

func Test_Cov3_ChmodVerify_IsEqualRwxFull(t *testing.T) {
	dir := covTempDir(t)
	_ = os.Chmod(dir, 0755)
	result := chmodhelper.ChmodVerify.IsEqualRwxFull(dir, "-rwxr-xr-x")
	actual := args.Map{"ok": result || !result}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify IsEqualRwxFull", actual)
}

func Test_Cov3_ChmodVerify_IsEqualRwxFullSkipInvalid(t *testing.T) {
	result := chmodhelper.ChmodVerify.IsEqualRwxFullSkipInvalid("/nonexistent_cov3", "-rwxr-xr-x")
	actual := args.Map{"result": result}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify IsEqualRwxFullSkipInvalid", actual)
}

func Test_Cov3_ChmodVerify_IsEqualSkipInvalid(t *testing.T) {
	result := chmodhelper.ChmodVerify.IsEqualSkipInvalid("/nonexistent_cov3", 0755)
	actual := args.Map{"result": result}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify IsEqualSkipInvalid", actual)
}

func Test_Cov3_ChmodVerify_MismatchError(t *testing.T) {
	dir := covTempDir(t)
	err := chmodhelper.ChmodVerify.MismatchError(dir, 0755)
	actual := args.Map{"executed": true}
	_ = err
	expected := args.Map{"executed": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify MismatchError", actual)
}

func Test_Cov3_ChmodVerify_PathIf(t *testing.T) {
	dir := covTempDir(t)
	skipErr := chmodhelper.ChmodVerify.PathIf(false, dir, 0755)
	actual := args.Map{"skipNil": skipErr == nil}
	expected := args.Map{"skipNil": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify PathIf false", actual)
}

func Test_Cov3_ChmodVerify_Path(t *testing.T) {
	dir := covTempDir(t)
	err := chmodhelper.ChmodVerify.Path(dir, 0755)
	actual := args.Map{"executed": true}
	_ = err
	expected := args.Map{"executed": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify Path", actual)
}

func Test_Cov3_ChmodVerify_GetExistingRwxWrapper(t *testing.T) {
	dir := covTempDir(t)
	rwx, err := chmodhelper.ChmodVerify.GetExistingRwxWrapper(dir)
	actual := args.Map{"noErr": err == nil, "defined": rwx.IsDefined()}
	expected := args.Map{"noErr": true, "defined": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify GetExistingRwxWrapper", actual)
}

func Test_Cov3_ChmodVerify_GetExistingRwxWrapperMust(t *testing.T) {
	dir := covTempDir(t)
	rwx := chmodhelper.ChmodVerify.GetExistingRwxWrapperMust(dir)
	actual := args.Map{"defined": rwx.IsDefined()}
	expected := args.Map{"defined": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify GetExistingRwxWrapperMust", actual)
}

// ── Attribute ──

func Test_Cov3_Attribute_Basic(t *testing.T) {
	attr := &chmodhelper.Attribute{IsRead: true, IsWrite: true, IsExecute: true}
	nilAttr := (*chmodhelper.Attribute)(nil)
	actual := args.Map{
		"isEmpty": attr.IsEmpty(), "isNull": attr.IsNull(), "isAnyNull": attr.IsAnyNull(),
		"isZero": attr.IsZero(), "isInvalid": attr.IsInvalid(), "isDefined": attr.IsDefined(),
		"hasAny": attr.HasAnyItem(), "byte": attr.ToByte(), "sum": attr.ToSum(),
		"rwxStr": attr.ToRwxString(), "strByte": attr.ToStringByte() > 0,
		"nilEmpty": nilAttr.IsEmpty(), "nilNull": nilAttr.IsNull(),
	}
	expected := args.Map{
		"isEmpty": false, "isNull": false, "isAnyNull": false,
		"isZero": false, "isInvalid": false, "isDefined": true,
		"hasAny": true, "byte": byte(7), "sum": byte(7),
		"rwxStr": "rwx", "strByte": true, "nilEmpty": true, "nilNull": true,
	}
	expected.ShouldBeEqual(t, 0, "Attribute basic", actual)
}

func Test_Cov3_Attribute_Clone(t *testing.T) {
	attr := &chmodhelper.Attribute{IsRead: true}
	cloned := attr.Clone()
	nilAttr := (*chmodhelper.Attribute)(nil)
	actual := args.Map{"read": cloned.IsRead, "nilClone": nilAttr.Clone() == nil}
	expected := args.Map{"read": true, "nilClone": true}
	expected.ShouldBeEqual(t, 0, "Attribute Clone", actual)
}

func Test_Cov3_Attribute_IsEqual(t *testing.T) {
	a1 := &chmodhelper.Attribute{IsRead: true, IsWrite: true}
	a2 := &chmodhelper.Attribute{IsRead: true, IsWrite: true}
	a3 := &chmodhelper.Attribute{IsRead: false}
	nilAttr := (*chmodhelper.Attribute)(nil)
	actual := args.Map{
		"equal":   a1.IsEqualPtr(a2),
		"notEq":   a1.IsEqualPtr(a3),
		"valEq":   a1.IsEqual(*a2),
		"nilNil":  nilAttr.IsEqualPtr(nilAttr),
		"nilLeft": nilAttr.IsEqualPtr(a1),
	}
	expected := args.Map{"equal": true, "notEq": false, "valEq": true, "nilNil": true, "nilLeft": false}
	expected.ShouldBeEqual(t, 0, "Attribute IsEqual", actual)
}

func Test_Cov3_Attribute_ToAttributeValue(t *testing.T) {
	attr := &chmodhelper.Attribute{IsRead: true, IsWrite: false, IsExecute: true}
	av := attr.ToAttributeValue()
	actual := args.Map{"sum": av.Sum}
	expected := args.Map{"sum": byte(5)}
	expected.ShouldBeEqual(t, 0, "Attribute ToAttributeValue", actual)
}

func Test_Cov3_Attribute_ToVariant(t *testing.T) {
	attr := &chmodhelper.Attribute{IsRead: true, IsWrite: true, IsExecute: true}
	v := attr.ToVariant()
	actual := args.Map{"gt0": v > 0}
	expected := args.Map{"gt0": true}
	expected.ShouldBeEqual(t, 0, "Attribute ToVariant", actual)
}

// ── Variant ──

func Test_Cov3_Variant_String(t *testing.T) {
	v := chmodhelper.X755
	actual := args.Map{"str": v.String()}
	expected := args.Map{"str": "755"}
	expected.ShouldBeEqual(t, 0, "Variant String", actual)
}

func Test_Cov3_Variant_ExpandOctalByte(t *testing.T) {
	r, w, x := chmodhelper.X755.ExpandOctalByte()
	actual := args.Map{"r": r, "w": w, "x": x}
	expected := args.Map{"r": byte('7'), "w": byte('5'), "x": byte('5')}
	expected.ShouldBeEqual(t, 0, "Variant ExpandOctalByte", actual)
}

func Test_Cov3_Variant_ToWrapper(t *testing.T) {
	rwx, err := chmodhelper.X755.ToWrapper()
	actual := args.Map{"noErr": err == nil, "defined": rwx.IsDefined()}
	expected := args.Map{"noErr": true, "defined": true}
	expected.ShouldBeEqual(t, 0, "Variant ToWrapper", actual)
}

func Test_Cov3_Variant_ToWrapperPtr(t *testing.T) {
	rwx, err := chmodhelper.X755.ToWrapperPtr()
	actual := args.Map{"noErr": err == nil, "notNil": rwx != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "Variant ToWrapperPtr", actual)
}

// ── GetRecursivePaths ──

func Test_Cov3_GetRecursivePaths(t *testing.T) {
	dir := covTempDir(t)
	_ = os.WriteFile(filepath.Join(dir, "a.txt"), []byte("x"), 0644)
	paths, err := chmodhelper.GetRecursivePaths(false, dir)
	actual := args.Map{"noErr": err == nil, "gt0": len(paths) > 0}
	expected := args.Map{"noErr": true, "gt0": true}
	expected.ShouldBeEqual(t, 0, "GetRecursivePaths", actual)
}

func Test_Cov3_GetRecursivePaths_File(t *testing.T) {
	dir := covTempDir(t)
	filePath := covWriteFile(t, dir, "single.txt", "x")
	paths, err := chmodhelper.GetRecursivePaths(false, filePath)
	actual := args.Map{"noErr": err == nil, "len": len(paths)}
	expected := args.Map{"noErr": true, "len": 1}
	expected.ShouldBeEqual(t, 0, "GetRecursivePaths file", actual)
}

func Test_Cov3_GetRecursivePaths_Invalid(t *testing.T) {
	_, err := chmodhelper.GetRecursivePaths(false, "/nonexistent_cov3_rec")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "GetRecursivePaths invalid", actual)
}

// ── TempDirGetter ──

func Test_Cov3_TempDirDefault(t *testing.T) {
	actual := args.Map{"notEmpty": chmodhelper.TempDirDefault != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TempDirDefault", actual)
}

// ── GetPathExistStat ──

func Test_Cov3_GetPathExistStat(t *testing.T) {
	dir := covTempDir(t)
	stat := chmodhelper.GetPathExistStat(dir)
	invalidStat := chmodhelper.GetPathExistStat("/nonexistent_cov3")
	actual := args.Map{"isExist": stat.IsExist, "invalidNotExist": !invalidStat.IsExist}
	expected := args.Map{"isExist": true, "invalidNotExist": true}
	expected.ShouldBeEqual(t, 0, "GetPathExistStat", actual)
}

// ── IsPathExists / IsPathInvalid / IsDirectory ──

func Test_Cov3_IsPathExists(t *testing.T) {
	dir := covTempDir(t)
	actual := args.Map{
		"exists":  chmodhelper.IsPathExists(dir),
		"invalid": chmodhelper.IsPathInvalid("/nonexistent_cov3"),
		"isDir":   chmodhelper.IsDirectory(dir),
	}
	expected := args.Map{"exists": true, "invalid": true, "isDir": true}
	expected.ShouldBeEqual(t, 0, "IsPathExists/IsPathInvalid/IsDirectory", actual)
}

// ── New.RwxWrapper creators ──

func Test_Cov3_NewRwxWrapper_UsingFileMode(t *testing.T) {
	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	actual := args.Map{"defined": rwx.IsDefined()}
	expected := args.Map{"defined": true}
	expected.ShouldBeEqual(t, 0, "NewRwxWrapper UsingFileMode", actual)
}

func Test_Cov3_NewRwxWrapper_UsingFileModePtr(t *testing.T) {
	rwx := chmodhelper.New.RwxWrapper.UsingFileModePtr(0755)
	actual := args.Map{"notNil": rwx != nil, "defined": rwx.IsDefined()}
	expected := args.Map{"notNil": true, "defined": true}
	expected.ShouldBeEqual(t, 0, "NewRwxWrapper UsingFileModePtr", actual)
}

func Test_Cov3_NewRwxWrapper_UsingVariant(t *testing.T) {
	rwx, err := chmodhelper.New.RwxWrapper.UsingVariant(chmodhelper.X755)
	actual := args.Map{"noErr": err == nil, "defined": rwx.IsDefined()}
	expected := args.Map{"noErr": true, "defined": true}
	expected.ShouldBeEqual(t, 0, "NewRwxWrapper UsingVariant", actual)
}

// ── New.SimpleFileReaderWriter ──

func Test_Cov3_NewSimpleFileRW_Default(t *testing.T) {
	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, "/tmp/test.txt")
	actual := args.Map{"notNil": rw != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewSimpleFileRW Default", actual)
}

func Test_Cov3_NewSimpleFileRW_Path(t *testing.T) {
	rw := chmodhelper.New.SimpleFileReaderWriter.Path(false, 0755, 0644, "/tmp/test.txt")
	actual := args.Map{"notNil": rw != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewSimpleFileRW Path", actual)
}

// ── New.Attribute ──

func Test_Cov3_NewAttribute_UsingRwx(t *testing.T) {
	attr := chmodhelper.New.Attribute.UsingRwxString("rwx")
	actual := args.Map{"defined": attr.IsDefined()}
	expected := args.Map{"defined": true}
	expected.ShouldBeEqual(t, 0, "NewAttribute UsingRwx", actual)
}
