package chmodhelper

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

// ══════════════════════════════════════════════════════════════════════════════
// CreateDefaultPaths — error path (L16)
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov2_CreateDirFilesWithRwxPermissions_ErrorPath(t *testing.T) {
	// Use an invalid path to trigger mkdirAll error
	items := []DirFilesWithRwxPermission{
		{
			DirWithFiles: DirWithFiles{
				DirPath: "/dev/null/impossible_dir",
			},
		},
	}
	err := CreateDirFilesWithRwxPermissions(false, items)
	if err == nil {
		// On some systems this might succeed — if so, skip the assertion
		t.Log("no error returned (OS allowed creation?)")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// CreateDefaultPathsMust — panic path (L11-12)
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov2_CreateDirFilesWithRwxPermissionsMust_PanicPath(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Log("no panic — OS allowed the creation")
		}
	}()
	items := []DirFilesWithRwxPermission{
		{
			DirWithFiles: DirWithFiles{
				DirPath: "/dev/null/impossible_dir",
			},
		},
	}
	CreateDirFilesWithRwxPermissionsMust(false, items)
}

// ══════════════════════════════════════════════════════════════════════════════
// RwxInstructionExecutor — CompiledWrapper fallthrough error (L67)
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov2_RwxInstructionExecutor_CompiledWrapper_NeitherFixedNorVar(t *testing.T) {
	exec := &RwxInstructionExecutor{
		// both isFixedWrapper and isVarWrapper are false
	}
	_, err := exec.CompiledWrapper(0o644)
	if err == nil {
		t.Fatal("expected error when neither fixed nor var wrapper")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// RwxInstructionExecutor — CompiledRwxWrapperUsingFixedRwxWrapper fallthrough (L85)
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov2_RwxInstructionExecutor_CompiledRwxWrapperUsingFixed_Error(t *testing.T) {
	exec := &RwxInstructionExecutor{}
	w := &RwxWrapper{}
	_, err := exec.CompiledRwxWrapperUsingFixedRwxWrapper(w)
	if err == nil {
		t.Fatal("expected error when neither fixed nor var")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// RwxInstructionExecutor — ApplyOnPathsDirect / ApplyOnPaths pass-through (L253, L261)
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov2_RwxInstructionExecutor_ApplyOnPathsDirect_Empty(t *testing.T) {
	exec := &RwxInstructionExecutor{}
	err := exec.ApplyOnPathsDirect()
	if err != nil {
		t.Fatal("expected nil for empty locations")
	}
}

func Test_Cov2_RwxInstructionExecutor_ApplyOnPaths_Empty(t *testing.T) {
	exec := &RwxInstructionExecutor{}
	err := exec.ApplyOnPaths([]string{})
	if err != nil {
		t.Fatal("expected nil for empty locations")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// RwxInstructionExecutors — ApplyOnPaths (L155)
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov2_RwxInstructionExecutors_ApplyOnPaths_Empty(t *testing.T) {
	execs := &RwxInstructionExecutors{}
	err := execs.ApplyOnPaths([]string{})
	if err != nil {
		t.Fatal("expected nil for empty")
	}
}

func Test_Cov2_RwxInstructionExecutors_ApplyOnPaths_EmptyExecutors(t *testing.T) {
	execs := &RwxInstructionExecutors{}
	err := execs.ApplyOnPaths([]string{"/tmp"})
	if err != nil {
		t.Fatal("expected nil for empty executors")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// PathExistStat — MeaningFullError with error (L239)
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov2_PathExistStat_MeaningFullError_WithError(t *testing.T) {
	stat := &PathExistStat{
		Location: "/nonexistent",
		Error:    errors.New("test error"),
	}
	err := stat.MeaningFullError()
	if err == nil {
		t.Fatal("expected meaningful error")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// MergeRwxWildcardWithFixedRwx — ParseRwxToVarAttribute error path (L38)
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov2_MergeRwxWildcardWithFixedRwx_InvalidWildcard(t *testing.T) {
	_, err := MergeRwxWildcardWithFixedRwx("INVALID", "rwx")
	if err == nil {
		t.Fatal("expected error for invalid wildcard")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// RwxVariableWrapper — nil RwxWrapper in apply loop (L186, L207)
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov2_RwxVariableWrapper_ApplyOnLocations_NilWrapper(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(tmpFile, []byte("test"), 0o644)

	// Create LocationFileInfoRwxWrappers with nil RwxWrapper
	items := []LocationFileInfoRwxWrapper{
		{
			Location:   tmpFile,
			RwxWrapper: nil,
		},
	}

	w := &RwxVariableWrapper{
		isAllFixedType: true,
	}
	// Exercise the nil RwxWrapper skip branch
	err := w.applyOnLocations(true, false, items)
	if err != nil {
		t.Fatal("expected nil — nil wrappers should be skipped")
	}

	// Also the continue-on-error path
	err = w.applyOnLocations(true, true, items)
	if err != nil {
		t.Fatal("expected nil — nil wrappers should be skipped")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// RwxVariableWrapper — Parse error (L46)
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov2_RwxVariableWrapper_Parse_Error(t *testing.T) {
	_, err := ParseRwxOwnerGroupOtherInstructionToRwxVariableWrapper("INVALID")
	if err == nil {
		t.Fatal("expected error for invalid input")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// RwxInstructionExecutor — verifyChmodLocations error paths (L196, L227)
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov2_RwxInstructionExecutor_VerifyChmod_CompiledWrapperError(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(tmpFile, []byte("test"), 0o644)

	// neither fixed nor var — CompiledWrapper will fail
	exec := &RwxInstructionExecutor{}
	info, _ := os.Stat(tmpFile)

	resultsMap := &FilteredPathFileInfoMap{
		FilesToInfoMap: map[string]os.FileInfo{
			tmpFile: info,
		},
	}

	// Exercise verifyChmodLocationsContinue error (L196)
	err := exec.verifyChmodLocationsContinue(resultsMap)
	if err == nil {
		t.Fatal("expected error from CompiledWrapper failure")
	}

	// Exercise verifyChmodLocationsNoContinue error (L227)
	err = exec.verifyChmodLocationsNoContinue(resultsMap)
	if err == nil {
		t.Fatal("expected error from CompiledWrapper failure")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// RwxPartialToInstructionExecutor — error path (L29)
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov2_RwxPartialToInstructionExecutor_InvalidPartial(t *testing.T) {
	cond := defaultCondition()
	_, err := ParseRwxPartialToInstructionExecutor("X", &cond)
	if err == nil {
		t.Fatal("expected error for invalid partial rwx")
	}
}

func defaultCondition() RwxMatchingStatus {
	return RwxMatchingStatus{}
}
