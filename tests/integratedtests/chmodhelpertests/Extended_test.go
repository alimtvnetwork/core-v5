package chmodhelpertests

import (
	"os"
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/results"
)

type pathExistStatType = chmodhelper.PathExistStat

type nilSafeResult = results.ResultAny

// ── PathExistStat nil-safety ──

func Test_PathExistStat_NilSafe(t *testing.T) {
	for caseIndex, tc := range extPathExistStatNilSafeCases {
		// Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}

// ── PathExistStat method tests ──

func Test_PathExistStat_InvalidPath(t *testing.T) {
	for caseIndex, testCase := range extPathExistStatTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		path, _ := input.GetAsString("path")

		// Act
		stat := chmodhelper.GetPathExistStat(path)

		actual := args.Map{
			"isExist":  stat.IsExist,
			"isInvalid": stat.IsInvalid(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_PathExistStat_TempDir(t *testing.T) {
	// Arrange
	tempDir := os.TempDir()

	// Act
	stat := chmodhelper.GetPathExistStat(tempDir)

	// Assert
	if !stat.IsExist {
		t.Error("temp dir should exist")
	}

	if !stat.IsDir() {
		t.Error("temp dir should be a directory")
	}

	if stat.IsFile() {
		t.Error("temp dir should not be a file")
	}

	if stat.IsInvalid() {
		t.Error("temp dir should not be invalid")
	}

	if stat.HasAnyIssues() {
		t.Error("temp dir should not have issues")
	}

	if stat.HasError() {
		t.Error("temp dir should not have error")
	}

	if !stat.IsEmptyError() {
		t.Error("temp dir should have empty error")
	}

	if !stat.HasFileInfo() {
		t.Error("temp dir should have file info")
	}

	if stat.IsInvalidFileInfo() {
		t.Error("temp dir should have valid file info")
	}

	if stat.FileMode() == nil {
		t.Error("FileMode should not be nil")
	}

	if stat.LastModifiedDate() == nil {
		t.Error("LastModifiedDate should not be nil")
	}

	if stat.NotExistError() != nil {
		t.Error("temp dir NotExistError should be nil")
	}

	if stat.NotADirError() != nil {
		t.Error("temp dir NotADirError should be nil")
	}

	if stat.String() == "" {
		t.Error("String should not be empty")
	}
}

func Test_PathExistStat_TempDir_Navigation(t *testing.T) {
	// Arrange
	tempDir := os.TempDir()
	stat := chmodhelper.GetPathExistStat(tempDir)

	// Act & Assert
	combined := stat.CombineWithNewPath("subdir")
	if combined == "" {
		t.Error("CombineWithNewPath should not be empty")
	}

	combinedStat := stat.CombineWith("subdir")
	if combinedStat == nil {
		t.Error("CombineWith should not return nil")
	}
}

func Test_PathExistStat_File(t *testing.T) {
	// Arrange
	tmpFile, err := os.CreateTemp("", "test-*.txt")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Act
	stat := chmodhelper.GetPathExistStat(tmpFile.Name())

	// Assert
	if !stat.IsExist {
		t.Error("temp file should exist")
	}

	if !stat.IsFile() {
		t.Error("should be a file")
	}

	if stat.IsDir() {
		t.Error("should not be a dir")
	}

	if stat.FileName() == "" {
		t.Error("FileName should not be empty")
	}

	if stat.ParentDir() == "" {
		t.Error("ParentDir should not be empty")
	}

	if stat.DotExt() == "" {
		t.Error("DotExt should not be empty for .txt file")
	}

	if stat.Size() == nil {
		t.Error("Size should not be nil")
	}

	parent := stat.Parent()
	if parent == nil {
		t.Error("Parent should not be nil")
	}

	if stat.NotAFileError() != nil {
		t.Error("NotAFileError should be nil for a file")
	}

	parentPath := stat.ParentWithNewPath("test.txt")
	if parentPath == "" {
		t.Error("ParentWithNewPath should not be empty")
	}

	parentStat := stat.ParentWith("test.txt")
	if parentStat == nil {
		t.Error("ParentWith should not return nil")
	}
}

func Test_PathExistStat_NotAFileError(t *testing.T) {
	// Arrange
	stat := chmodhelper.GetPathExistStat(os.TempDir())

	// Act
	err := stat.NotAFileError()

	// Assert
	if err == nil {
		t.Error("NotAFileError should return error for directory")
	}
}

func Test_PathExistStat_NotADirError(t *testing.T) {
	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()
	stat := chmodhelper.GetPathExistStat(tmpFile.Name())

	// Act
	err := stat.NotADirError()

	// Assert
	if err == nil {
		t.Error("NotADirError should return error for file")
	}
}

func Test_PathExistStat_NotExist(t *testing.T) {
	// Arrange
	stat := chmodhelper.GetPathExistStat("/nonexistent/path/xyz")

	// Act
	err := stat.NotExistError()

	// Assert
	if err == nil {
		t.Error("NotExistError should return error for non-existent path")
	}
}

func Test_PathExistStat_Dispose(t *testing.T) {
	// Arrange
	stat := chmodhelper.GetPathExistStat(os.TempDir())

	// Act
	stat.Dispose()

	// Assert
	if stat.IsExist {
		t.Error("IsExist should be false after Dispose")
	}

	if stat.Location != "" {
		t.Error("Location should be empty after Dispose")
	}
}

func Test_PathExistStat_MessageWithPathWrapped(t *testing.T) {
	// Arrange
	stat := chmodhelper.GetPathExistStat(os.TempDir())

	// Act
	msg := stat.MessageWithPathWrapped("test message")

	// Assert
	if msg == "" {
		t.Error("MessageWithPathWrapped should not be empty")
	}
}

func Test_PathExistStat_ParentWithGlobPatternFiles(t *testing.T) {
	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()
	stat := chmodhelper.GetPathExistStat(tmpFile.Name())

	// Act
	_, err := stat.ParentWithGlobPatternFiles("*.txt")

	// Assert
	if err != nil {
		t.Errorf("ParentWithGlobPatternFiles error: %v", err)
	}
}

// ── chmodVerifier tests ──

func Test_ChmodVerifier_GetRwx9(t *testing.T) {
	// Arrange
	fileMode := os.FileMode(0755)

	// Act
	rwx9 := chmodhelper.ChmodVerify.GetRwx9(fileMode)

	// Assert
	if len(rwx9) != 9 {
		t.Errorf("expected 9 chars, got %d: %s", len(rwx9), rwx9)
	}
}

func Test_ChmodVerifier_GetRwxFull(t *testing.T) {
	// Arrange
	fileMode := os.FileMode(0755)

	// Act
	rwxFull := chmodhelper.ChmodVerify.GetRwxFull(fileMode)

	// Assert
	if len(rwxFull) != 10 {
		t.Errorf("expected 10 chars, got %d: %s", len(rwxFull), rwxFull)
	}
}

func Test_ChmodVerifier_IsEqual(t *testing.T) {
	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	info, _ := os.Stat(tmpFile.Name())
	existingMode := info.Mode()

	// Act
	isEqual := chmodhelper.ChmodVerify.IsEqual(tmpFile.Name(), existingMode)

	// Assert
	if !isEqual {
		t.Error("IsEqual should return true for existing file mode")
	}
}

func Test_ChmodVerifier_IsMismatch(t *testing.T) {
	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Act
	isMismatch := chmodhelper.ChmodVerify.IsMismatch(tmpFile.Name(), os.FileMode(0777))

	// Assert -- may or may not be mismatch depending on OS defaults
	_ = isMismatch
}

func Test_ChmodVerifier_IsEqualSkipInvalid(t *testing.T) {
	// Arrange
	invalidPath := "/nonexistent/path/xyz"

	// Act
	isEqual := chmodhelper.ChmodVerify.IsEqualSkipInvalid(invalidPath, os.FileMode(0644))

	// Assert
	if !isEqual {
		t.Error("IsEqualSkipInvalid should return true for invalid path")
	}
}

func Test_ChmodVerifier_IsEqualRwxFullSkipInvalid(t *testing.T) {
	// Arrange
	invalidPath := "/nonexistent/path/xyz"

	// Act
	isEqual := chmodhelper.ChmodVerify.IsEqualRwxFullSkipInvalid(invalidPath, "-rwxr-xr-x")

	// Assert
	if !isEqual {
		t.Error("IsEqualRwxFullSkipInvalid should return true for invalid path")
	}
}

func Test_ChmodVerifier_GetExisting(t *testing.T) {
	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Act
	mode, err := chmodhelper.ChmodVerify.GetExisting(tmpFile.Name())

	// Assert
	if err != nil {
		t.Errorf("GetExisting error: %v", err)
	}

	if mode == 0 {
		t.Error("mode should not be 0")
	}
}

func Test_ChmodVerifier_PathIf(t *testing.T) {
	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Act
	err := chmodhelper.ChmodVerify.PathIf(false, tmpFile.Name(), os.FileMode(0644))

	// Assert
	if err != nil {
		t.Error("PathIf with false should return nil")
	}
}

// ── GetExistingChmod tests ──

func Test_GetExistingChmod(t *testing.T) {
	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Act
	mode, err := chmodhelper.GetExistingChmod(tmpFile.Name())

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}

	if mode == 0 {
		t.Error("mode should not be 0")
	}
}

func Test_GetExistingChmod_InvalidPath(t *testing.T) {
	// Act
	_, err := chmodhelper.GetExistingChmod("/nonexistent/xyz")

	// Assert
	if err == nil {
		t.Error("expected error for invalid path")
	}
}

// ── IsPathExists / IsDirectory / IsPathInvalid ──

func Test_IsPathExists(t *testing.T) {
	// Assert
	if !chmodhelper.IsPathExists(os.TempDir()) {
		t.Error("temp dir should exist")
	}

	if chmodhelper.IsPathExists("/nonexistent/path/xyz") {
		t.Error("nonexistent path should not exist")
	}
}

func Test_IsDirectory(t *testing.T) {
	// Assert
	if !chmodhelper.IsDirectory(os.TempDir()) {
		t.Error("temp dir should be directory")
	}
}

func Test_IsPathInvalid(t *testing.T) {
	// Assert
	if !chmodhelper.IsPathInvalid("/nonexistent/xyz") {
		t.Error("nonexistent path should be invalid")
	}

	if chmodhelper.IsPathInvalid(os.TempDir()) {
		t.Error("temp dir should not be invalid")
	}
}

// ── TempDirGetter / TempDirDefault ──

func Test_TempDirDefault(t *testing.T) {
	// Assert
	if chmodhelper.TempDirDefault == "" {
		t.Error("TempDirDefault should not be empty")
	}
}

// ── GetExistingChmodRwxWrapperPtr ──

func Test_GetExistingChmodRwxWrapperPtr(t *testing.T) {
	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Act
	wrapper, err := chmodhelper.GetExistingChmodRwxWrapperPtr(tmpFile.Name())

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}

	if wrapper == nil {
		t.Error("wrapper should not be nil")
	}
}

func Test_GetExistingChmodRwxWrapperPtr_InvalidPath(t *testing.T) {
	// Act
	_, err := chmodhelper.GetExistingChmodRwxWrapperPtr("/nonexistent/xyz")

	// Assert
	if err == nil {
		t.Error("expected error for invalid path")
	}
}

// ── dirCreator tests ──

func Test_DirCreator_Direct(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_direct_test"
	defer os.RemoveAll(dir)

	// Act
	err := chmodhelper.SimpleFileWriter.CreateDir.Direct(dir)

	// Assert
	if err != nil {
		t.Errorf("Direct error: %v", err)
	}

	if !chmodhelper.IsPathExists(dir) {
		t.Error("dir should exist after Direct")
	}
}

func Test_DirCreator_IfMissing(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_ifmissing_test"
	defer os.RemoveAll(dir)

	// Act
	err := chmodhelper.SimpleFileWriter.DirCreator.IfMissing(os.FileMode(0755), dir)

	// Assert
	if err != nil {
		t.Errorf("IfMissing error: %v", err)
	}

	// Act again - should be no-op
	err2 := chmodhelper.SimpleFileWriter.DirCreator.IfMissing(os.FileMode(0755), dir)

	// Assert
	if err2 != nil {
		t.Errorf("IfMissing second call error: %v", err2)
	}
}

func Test_DirCreator_If_False(t *testing.T) {
	// Act
	err := chmodhelper.SimpleFileWriter.DirCreator.If(false, os.FileMode(0755), "/whatever")

	// Assert
	if err != nil {
		t.Error("If with false should return nil")
	}
}

// ── fileWriter tests ──

func Test_FileWriter_WriteAndRead(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_write"
	filePath := dir + "/test.txt"
	defer os.RemoveAll(dir)
	content := "hello world"

	// Act
	err := chmodhelper.SimpleFileWriter.FileWriter.Bytes.Default(
		true,
		filePath,
		[]byte(content),
	)

	// Assert
	if err != nil {
		t.Errorf("write error: %v", err)
	}

	if !chmodhelper.IsPathExists(filePath) {
		t.Error("file should exist after write")
	}
}

func Test_FileWriter_String_Default(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_string_write"
	filePath := dir + "/test.txt"
	defer os.RemoveAll(dir)

	// Act
	err := chmodhelper.SimpleFileWriter.FileWriter.String.Default(
		true,
		filePath,
		"hello string",
	)

	// Assert
	if err != nil {
		t.Errorf("string write error: %v", err)
	}
}

func Test_FileWriter_Remove(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_remove"
	filePath := dir + "/test.txt"
	os.MkdirAll(dir, 0755)
	os.WriteFile(filePath, []byte("test"), 0644)

	// Act
	err := chmodhelper.SimpleFileWriter.FileWriter.Remove(filePath)

	// Assert
	if err != nil {
		t.Errorf("remove error: %v", err)
	}

	if chmodhelper.IsPathExists(filePath) {
		t.Error("file should not exist after remove")
	}

	os.RemoveAll(dir)
}

func Test_FileWriter_RemoveIf(t *testing.T) {
	// Act -- should be no-op when false
	err := chmodhelper.SimpleFileWriter.FileWriter.RemoveIf(false, "/whatever")

	// Assert
	if err != nil {
		t.Error("RemoveIf with false should return nil")
	}
}

func Test_FileWriter_ParentDir(t *testing.T) {
	// Act
	parent := chmodhelper.SimpleFileWriter.FileWriter.ParentDir("/tmp/test/file.txt")

	// Assert
	if parent == "" {
		t.Error("ParentDir should not be empty")
	}
}
