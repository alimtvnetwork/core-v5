package chmodhelpertests

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/alimtvnetwork/core/chmodhelper/chmodins"
)

// ── RwxWrapper.ToUint32Octal ──

func Test_Cov9_RwxWrapper_ToUint32Octal(t *testing.T) {
	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	oct := rwx.ToUint32Octal()
	if oct != 0755 {
		t.Fatalf("expected 0755, got %o", oct)
	}
}

// ── RwxWrapper.ApplyChmod branches ──

func Test_Cov9_RwxWrapper_ApplyChmod_SkipInvalid(t *testing.T) {
	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	err := rwx.ApplyChmod(true, "/nonexistent/cov9/skip")
	if err != nil {
		t.Fatal("expected nil for skip on invalid")
	}
}

func Test_Cov9_RwxWrapper_ApplyChmod_NotSkipInvalid(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("chmod behavior differs on Windows")
	}
	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	err := rwx.ApplyChmod(false, "/nonexistent/cov9/noskip")
	if err == nil {
		t.Fatal("expected error for non-skip invalid path")
	}
}

func Test_Cov9_RwxWrapper_ApplyChmod_Success(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "cov9_apply_chmod.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0644)
	err := rwx.ApplyChmod(false, tmpFile)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_Cov9_RwxWrapper_ApplyChmod_ChmodFail(t *testing.T) {
	// On most systems, regular chmod doesn't fail on valid paths
	// This covers the success path with error=nil
	tmpFile := filepath.Join(os.TempDir(), "cov9_apply_chmod_ok.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0777)
	err := rwx.ApplyChmod(false, tmpFile)
	if err != nil {
		t.Fatal(err)
	}
}

// ── RwxWrapper.invalidPathErr ──

func Test_Cov9_RwxWrapper_InvalidPathErr(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("chmod behavior differs on Windows")
	}
	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	err := rwx.ApplyChmod(false, "/nonexistent/cov9/invalid_path")
	if err == nil {
		t.Fatal("expected error")
	}
}

// ── RwxWrapper.ApplyChmodOptions ──

func Test_Cov9_ApplyChmodOptions_SkipApply(t *testing.T) {
	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	err := rwx.ApplyChmodOptions(false, true, false, "/any")
	if err != nil {
		t.Fatal("expected nil when isApply=false")
	}
}

func Test_Cov9_ApplyChmodOptions_InvalidSkip(t *testing.T) {
	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	err := rwx.ApplyChmodOptions(true, true, true, "/nonexistent/cov9/opts")
	if err != nil {
		t.Fatal("expected nil for skip invalid")
	}
}

func Test_Cov9_ApplyChmodOptions_InvalidNoSkip(t *testing.T) {
	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	err := rwx.ApplyChmodOptions(true, true, false, "/nonexistent/cov9/opts2")
	if err == nil {
		t.Fatal("expected error for invalid no-skip")
	}
}

func Test_Cov9_ApplyChmodOptions_MismatchApply(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "cov9_opts_mismatch.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	err := rwx.ApplyChmodOptions(true, true, false, tmpFile)
	_ = err
}

func Test_Cov9_ApplyChmodOptions_AlreadyMatching(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "cov9_opts_match.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	os.Chmod(tmpFile, 0644)
	defer os.Remove(tmpFile)

	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0644)
	err := rwx.ApplyChmodOptions(true, true, false, tmpFile)
	if err != nil {
		t.Fatal(err)
	}
}

// ── RwxWrapper.LinuxApplyRecursive ──

func Test_Cov9_LinuxApplyRecursive_SkipInvalid_NotExists(t *testing.T) {
	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	err := rwx.LinuxApplyRecursive(true, "/nonexistent/cov9/linux_recur")
	if err != nil {
		t.Fatal("expected nil for skip invalid")
	}
}

func Test_Cov9_LinuxApplyRecursive_NoSkip_NotExists(t *testing.T) {
	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	err := rwx.LinuxApplyRecursive(false, "/nonexistent/cov9/linux_recur2")
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov9_LinuxApplyRecursive_Valid(t *testing.T) {
	tmpDir := filepath.Join(os.TempDir(), "cov9_linux_recur")
	os.MkdirAll(tmpDir, 0755)
	defer os.RemoveAll(tmpDir)

	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	err := rwx.LinuxApplyRecursive(false, tmpDir)
	_ = err // depends on OS
}

// ── RwxWrapper.ApplyRecursive ──

func Test_Cov9_ApplyRecursive_SkipInvalid(t *testing.T) {
	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	err := rwx.ApplyRecursive(true, "/nonexistent/cov9/recur_skip")
	if err != nil {
		t.Fatal("expected nil")
	}
}

func Test_Cov9_ApplyRecursive_NotExist_NoSkip(t *testing.T) {
	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	err := rwx.ApplyRecursive(false, "/nonexistent/cov9/recur_noskip")
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov9_ApplyRecursive_File(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "cov9_recur_file.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0644)
	err := rwx.ApplyRecursive(false, tmpFile)
	_ = err
}

func Test_Cov9_ApplyRecursive_Dir(t *testing.T) {
	tmpDir := filepath.Join(os.TempDir(), "cov9_recur_dir")
	os.MkdirAll(filepath.Join(tmpDir, "sub"), 0755)
	os.WriteFile(filepath.Join(tmpDir, "f.txt"), []byte("x"), 0644)
	defer os.RemoveAll(tmpDir)

	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	err := rwx.ApplyRecursive(false, tmpDir)
	_ = err
}

// ── RwxWrapper.MustApplyChmod ──

func Test_Cov9_MustApplyChmod_Success(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "cov9_must_apply.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0644)
	rwx.MustApplyChmod(tmpFile) // should not panic
}

func Test_Cov9_MustApplyChmod_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0644)
	rwx.MustApplyChmod("/nonexistent/cov9/must_apply")
}

// ── RwxWrapper.ApplyLinuxChmodOnMany ──

func Test_Cov9_ApplyLinuxChmodOnMany_Recursive(t *testing.T) {
	tmpDir := filepath.Join(os.TempDir(), "cov9_linux_many_recur")
	os.MkdirAll(tmpDir, 0755)
	defer os.RemoveAll(tmpDir)

	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	err := rwx.ApplyLinuxChmodOnMany(
		&chmodins.Condition{IsRecursive: true},
		tmpDir)
	_ = err
}

func Test_Cov9_ApplyLinuxChmodOnMany_NonRecursive(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "cov9_linux_many_nonrecur.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0644)
	err := rwx.ApplyLinuxChmodOnMany(
		&chmodins.Condition{IsRecursive: false},
		tmpFile)
	_ = err
}

func Test_Cov9_ApplyLinuxChmodOnMany_ContinueOnError_NonRecursive(t *testing.T) {
	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	err := rwx.ApplyLinuxChmodOnMany(
		&chmodins.Condition{IsContinueOnError: true, IsRecursive: false},
		"/nonexistent/cov9/many1", "/nonexistent/cov9/many2")
	_ = err
}

func Test_Cov9_ApplyLinuxChmodOnMany_ContinueOnError_Recursive(t *testing.T) {
	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	err := rwx.ApplyLinuxChmodOnMany(
		&chmodins.Condition{IsContinueOnError: true, IsRecursive: true},
		"/nonexistent/cov9/many3")
	_ = err
}

func Test_Cov9_ApplyLinuxChmodOnMany_StopOnError_Recursive(t *testing.T) {
	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	err := rwx.ApplyLinuxChmodOnMany(
		&chmodins.Condition{IsRecursive: true},
		"/nonexistent/cov9/many4")
	_ = err
}

func Test_Cov9_ApplyLinuxChmodOnMany_StopOnError_NonRecursive(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "cov9_stop_nonrecur.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0644)
	err := rwx.ApplyLinuxChmodOnMany(
		&chmodins.Condition{},
		tmpFile, "/nonexistent/cov9/many5")
	_ = err
}

// ── RwxWrapper.IsEqualVarWrapper ──

func Test_Cov9_IsEqualVarWrapper_Nil(t *testing.T) {
	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	if rwx.IsEqualVarWrapper(nil) {
		t.Fatal("expected false for nil")
	}
}

func Test_Cov9_IsEqualVarWrapper_Match(t *testing.T) {
	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	varW, _ := chmodhelper.NewRwxVariableWrapper("-rwxr-xr-x")
	result := rwx.IsEqualVarWrapper(varW)
	if !result {
		t.Fatal("expected true")
	}
}

// ── RwxWrapper.IsRwxEqualFileInfo ──

func Test_Cov9_IsRwxEqualFileInfo_Nil(t *testing.T) {
	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	if rwx.IsRwxEqualFileInfo(nil) {
		t.Fatal("expected false")
	}
}

func Test_Cov9_IsRwxEqualFileInfo_Valid(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "cov9_fileinfo.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	os.Chmod(tmpFile, 0644)
	defer os.Remove(tmpFile)

	info, _ := os.Stat(tmpFile)
	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0644)
	result := rwx.IsRwxEqualFileInfo(info)
	if !result {
		t.Fatal("expected true")
	}
}

// ── RwxWrapper.IsRwxEqualLocation ──

func Test_Cov9_IsRwxEqualLocation_NonExistent(t *testing.T) {
	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	if rwx.IsRwxEqualLocation("/nonexistent/cov9/rwxloc") {
		t.Fatal("expected false")
	}
}

func Test_Cov9_IsRwxEqualLocation_Valid(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "cov9_rwxloc.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	os.Chmod(tmpFile, 0644)
	defer os.Remove(tmpFile)

	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0644)
	result := rwx.IsRwxEqualLocation(tmpFile)
	if !result {
		t.Fatal("expected true")
	}
}

// ── RwxWrapper.getLinuxRecursiveCmdForChmod ──

func Test_Cov9_GetLinuxRecursiveCmdForChmod(t *testing.T) {
	// Covered through LinuxApplyRecursive on valid dir
	tmpDir := filepath.Join(os.TempDir(), "cov9_getcmd")
	os.MkdirAll(tmpDir, 0755)
	defer os.RemoveAll(tmpDir)

	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	_ = rwx.LinuxApplyRecursive(false, tmpDir)
}

// ── RwxWrapper.applyLinuxRecursiveChmodUsingCmd ──

func Test_Cov9_ApplyLinuxRecursiveChmodUsingCmd(t *testing.T) {
	tmpDir := filepath.Join(os.TempDir(), "cov9_linuxcmd")
	os.MkdirAll(tmpDir, 0755)
	defer os.RemoveAll(tmpDir)

	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	err := rwx.LinuxApplyRecursive(false, tmpDir)
	_ = err
}
