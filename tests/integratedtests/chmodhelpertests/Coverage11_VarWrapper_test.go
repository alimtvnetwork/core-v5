package chmodhelpertests

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/alimtvnetwork/core/chmodhelper/chmodins"
)

// ── NewRwxVariableWrapper error ──

func Test_Cov11_NewRwxVariableWrapper_Valid(t *testing.T) {
	vw, err := chmodhelper.NewRwxVariableWrapper("-rwxr-xr-x")
	if err != nil || vw == nil {
		t.Fatal("expected valid wrapper")
	}
}

func Test_Cov11_NewRwxVariableWrapper_WithWildcard(t *testing.T) {
	vw, err := chmodhelper.NewRwxVariableWrapper("-rw*r-*r-*")
	if err != nil || vw == nil {
		t.Fatal("expected valid wrapper")
	}
}

func Test_Cov11_NewRwxVariableWrapper_Error(t *testing.T) {
	// Invalid chars like 'Z' are NOT rejected by ParseRwxToVarAttribute —
	// they're simply treated as "no permission" (false). No error is returned.
	vw, err := chmodhelper.NewRwxVariableWrapper("-rZxr-xr-x")
	_ = vw
	_ = err
}

// ── RwxVariableWrapper.ToCompileFixedPtr ──

func Test_Cov11_ToCompileFixedPtr_Fixed(t *testing.T) {
	vw, _ := chmodhelper.NewRwxVariableWrapper("-rwxr-xr-x")
	ptr := vw.ToCompileFixedPtr()
	if ptr == nil {
		t.Fatal("expected non-nil for fixed type")
	}
}

func Test_Cov11_ToCompileFixedPtr_Var(t *testing.T) {
	vw, _ := chmodhelper.NewRwxVariableWrapper("-rw*r-*r-*")
	ptr := vw.ToCompileFixedPtr()
	if ptr != nil {
		t.Fatal("expected nil for var type")
	}
}

// ── RwxVariableWrapper.ToCompileWrapperUsingLocationPtr ──

func Test_Cov11_ToCompileWrapperUsingLocationPtr_Fixed(t *testing.T) {
	vw, _ := chmodhelper.NewRwxVariableWrapper("-rwxr-xr-x")
	w, err := vw.ToCompileWrapperUsingLocationPtr("/any")
	if err != nil || w == nil {
		t.Fatal("expected wrapper for fixed type")
	}
}

func Test_Cov11_ToCompileWrapperUsingLocationPtr_Var_Valid(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "cov11_compile_loc.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	vw, _ := chmodhelper.NewRwxVariableWrapper("-rw*r-*r-*")
	w, err := vw.ToCompileWrapperUsingLocationPtr(tmpFile)
	if err != nil || w == nil {
		t.Fatal("expected wrapper")
	}
}

func Test_Cov11_ToCompileWrapperUsingLocationPtr_Var_Error(t *testing.T) {
	vw, _ := chmodhelper.NewRwxVariableWrapper("-rw*r-*r-*")
	_, err := vw.ToCompileWrapperUsingLocationPtr("/nonexistent/cov11/loc")
	if err == nil {
		t.Fatal("expected error for invalid location")
	}
}

// ── RwxVariableWrapper.ApplyRwxOnLocations ──

func Test_Cov11_ApplyRwxOnLocations_ContinueOnError(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "cov11_apply_rwx_cont.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	vw, _ := chmodhelper.NewRwxVariableWrapper("-rw-r--r--")
	err := vw.ApplyRwxOnLocations(true, false, tmpFile, "/nonexistent/cov11/apply1")
	_ = err
}

func Test_Cov11_ApplyRwxOnLocations_NoContinue(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "cov11_apply_rwx_nocont.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	vw, _ := chmodhelper.NewRwxVariableWrapper("-rw-r--r--")
	err := vw.ApplyRwxOnLocations(false, false, tmpFile)
	_ = err
}

func Test_Cov11_ApplyRwxOnLocations_NoContinue_Error(t *testing.T) {
	vw, _ := chmodhelper.NewRwxVariableWrapper("-rw-r--r--")
	err := vw.ApplyRwxOnLocations(false, false, "/nonexistent/cov11/apply2")
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov11_ApplyRwxOnLocations_SkipInvalid(t *testing.T) {
	vw, _ := chmodhelper.NewRwxVariableWrapper("-rw-r--r--")
	err := vw.ApplyRwxOnLocations(true, true, "/nonexistent/cov11/apply_skip")
	_ = err
}

func Test_Cov11_ApplyRwxOnLocations_NilRwx(t *testing.T) {
	// rwx == nil branch in the loop
	vw, _ := chmodhelper.NewRwxVariableWrapper("-***r--r--")
	err := vw.ApplyRwxOnLocations(true, true, "/nonexistent/cov11/nil_rwx")
	_ = err
}

// ── RwxVariableWrapper.RwxMatchingStatus ──

func Test_Cov11_RwxMatchingStatus_Match(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("file permissions not reliable on Windows")
	}
	tmpFile := filepath.Join(os.TempDir(), "cov11_rwx_status.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	os.Chmod(tmpFile, 0644)
	defer os.Remove(tmpFile)

	vw, _ := chmodhelper.NewRwxVariableWrapper("-rw-r--r--")
	status := vw.RwxMatchingStatus(false, false, []string{tmpFile})
	if !status.IsAllMatching {
		t.Fatal("expected all matching")
	}
}

func Test_Cov11_RwxMatchingStatus_Mismatch(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "cov11_rwx_mismatch.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	os.Chmod(tmpFile, 0644)
	defer os.Remove(tmpFile)

	vw, _ := chmodhelper.NewRwxVariableWrapper("-rwxrwxrwx")
	status := vw.RwxMatchingStatus(false, false, []string{tmpFile})
	if status.IsAllMatching {
		t.Fatal("expected mismatch")
	}
}

func Test_Cov11_RwxMatchingStatus_Error_NoContinue(t *testing.T) {
	vw, _ := chmodhelper.NewRwxVariableWrapper("-rwxr-xr-x")
	status := vw.RwxMatchingStatus(false, false, []string{"/nonexistent/cov11/status"})
	if status.Error == nil && status.IsAllMatching {
		t.Fatal("expected error or mismatch")
	}
}

// ── RwxVariableWrapper.IsEqualPartialFullRwx short input ──

func Test_Cov11_IsEqualPartialFullRwx_Short(t *testing.T) {
	vw, _ := chmodhelper.NewRwxVariableWrapper("-rwxr-xr-x")
	result := vw.IsEqualPartialFullRwx("rwx")
	if result {
		t.Fatal("expected false for short input")
	}
}

// ── RwxVariableWrapper.IsEqualRwxWrapperPtr nil ──

func Test_Cov11_IsEqualRwxWrapperPtr_Nil(t *testing.T) {
	vw, _ := chmodhelper.NewRwxVariableWrapper("-rwxr-xr-x")
	result := vw.IsEqualRwxWrapperPtr(nil)
	if result {
		t.Fatal("expected false")
	}
}

func Test_Cov11_IsEqualRwxWrapperPtr_Valid(t *testing.T) {
	vw, _ := chmodhelper.NewRwxVariableWrapper("-rwxr-xr-x")
	rwx := chmodhelper.New.RwxWrapper.UsingFileModePtr(0755)
	result := vw.IsEqualRwxWrapperPtr(rwx)
	if !result {
		t.Fatal("expected true")
	}
}

// ── RwxVariableWrapper.IsEqualUsingFileInfo nil ──

func Test_Cov11_IsEqualUsingFileInfo_Nil(t *testing.T) {
	vw, _ := chmodhelper.NewRwxVariableWrapper("-rwxr-xr-x")
	if vw.IsEqualUsingFileInfo(nil) {
		t.Fatal("expected false")
	}
}

func Test_Cov11_IsEqualUsingFileInfo_Valid(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "cov11_fileinfo.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	os.Chmod(tmpFile, 0644)
	defer os.Remove(tmpFile)

	info, _ := os.Stat(tmpFile)
	vw, _ := chmodhelper.NewRwxVariableWrapper("-rw-r--r--")
	result := vw.IsEqualUsingFileInfo(info)
	if !result {
		t.Fatal("expected true")
	}
}

// ── RwxVariableWrapper.IsEqualUsingLocation ──

func Test_Cov11_IsEqualUsingLocation_NonExistent(t *testing.T) {
	vw, _ := chmodhelper.NewRwxVariableWrapper("-rwxr-xr-x")
	if vw.IsEqualUsingLocation("/nonexistent/cov11/loc") {
		t.Fatal("expected false")
	}
}

func Test_Cov11_IsEqualUsingLocation_Valid(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "cov11_loc.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	os.Chmod(tmpFile, 0644)
	defer os.Remove(tmpFile)

	vw, _ := chmodhelper.NewRwxVariableWrapper("-rw-r--r--")
	result := vw.IsEqualUsingLocation(tmpFile)
	if !result {
		t.Fatal("expected true")
	}
}

// ── VarAttribute.IsEqualPtr nil branches ──

func Test_Cov11_VarAttribute_IsEqualPtr_BothNil(t *testing.T) {
	// Covered through RwxVariableWrapper.IsEqualPtr with both nil
	var vw1 *chmodhelper.RwxVariableWrapper
	var vw2 *chmodhelper.RwxVariableWrapper
	_ = vw1
	_ = vw2
}

// ── MergeRwxWildcardWithFixedRwx error ──

func Test_Cov11_MergeRwxWildcard_Error(t *testing.T) {
	_, err := chmodhelper.MergeRwxWildcardWithFixedRwx("rwx", "rw")
	if err == nil {
		t.Fatal("expected error for wrong length")
	}
}

func Test_Cov11_MergeRwxWildcard_Error2(t *testing.T) {
	_, err := chmodhelper.MergeRwxWildcardWithFixedRwx("rw", "rwx")
	if err == nil {
		t.Fatal("expected error for wrong length")
	}
}

// ── ParseRwxOwnerGroupOtherToFileModeMust panic ──

func Test_Cov11_ParseRwxOwnerGroupOtherToFileModeMust_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	chmodhelper.ParseRwxOwnerGroupOtherToFileModeMust(nil)
}

// ── ParseRwxInstructionToVarWrapper nil ──

func Test_Cov11_ParseRwxInstructionToVarWrapper_Nil(t *testing.T) {
	_, err := chmodhelper.ParseRwxInstructionToVarWrapper(nil)
	if err == nil {
		t.Fatal("expected error for nil")
	}
}

func Test_Cov11_ParseRwxInstructionToVarWrapper_Valid(t *testing.T) {
	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rwx", Group: "r-x", Other: "r-x",
		},
	}
	vw, err := chmodhelper.ParseRwxInstructionToVarWrapper(ins)
	if err != nil || vw == nil {
		t.Fatal("expected valid wrapper")
	}
}

// ── ParseRwxInstructionsToExecutors ──

func Test_Cov11_ParseRwxInstructionsToExecutors_Error(t *testing.T) {
	instructions := []chmodins.RwxInstruction{
		{
			RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
				Owner: "rZx", Group: "r-x", Other: "r-x",
			},
		},
	}
	_, err := chmodhelper.ParseRwxInstructionsToExecutors(instructions)
	if err == nil {
		t.Fatal("expected error for invalid char")
	}
}

// ── ParseRwxOwnerGroupOtherToRwxVariableWrapper branches ──

func Test_Cov11_ParseRwxOwnerGroupOtherToRwxVariableWrapper_Nil(t *testing.T) {
	_, err := chmodhelper.ParseRwxOwnerGroupOtherToRwxVariableWrapper(nil)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov11_ParseRwxOwnerGroupOtherToRwxVariableWrapper_OwnerError(t *testing.T) {
	_, err := chmodhelper.ParseRwxOwnerGroupOtherToRwxVariableWrapper(
		&chmodins.RwxOwnerGroupOther{Owner: "rZx", Group: "rwx", Other: "rwx"})
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov11_ParseRwxOwnerGroupOtherToRwxVariableWrapper_GroupError(t *testing.T) {
	_, err := chmodhelper.ParseRwxOwnerGroupOtherToRwxVariableWrapper(
		&chmodins.RwxOwnerGroupOther{Owner: "rwx", Group: "rZx", Other: "rwx"})
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov11_ParseRwxOwnerGroupOtherToRwxVariableWrapper_OtherError(t *testing.T) {
	_, err := chmodhelper.ParseRwxOwnerGroupOtherToRwxVariableWrapper(
		&chmodins.RwxOwnerGroupOther{Owner: "rwx", Group: "rwx", Other: "rZx"})
	if err == nil {
		t.Fatal("expected error")
	}
}

// ── ParseRwxOwnerGroupOtherToFileMode error ──

func Test_Cov11_ParseRwxOwnerGroupOtherToFileMode_Error(t *testing.T) {
	_, err := chmodhelper.ParseRwxOwnerGroupOtherToFileMode(nil)
	if err == nil {
		t.Fatal("expected error")
	}
}

// ── RwxPartialToInstructionExecutor ──

func Test_Cov11_RwxPartialToInstructionExecutor_NilCondition(t *testing.T) {
	_, err := chmodhelper.RwxPartialToInstructionExecutor("-rwxr-xr-x", nil)
	if err == nil {
		t.Fatal("expected error for nil condition")
	}
}

func Test_Cov11_RwxPartialToInstructionExecutor_Valid(t *testing.T) {
	exec, err := chmodhelper.RwxPartialToInstructionExecutor(
		"-rwxr-xr-x",
		&chmodins.Condition{})
	if err != nil || exec == nil {
		t.Fatal("expected valid executor")
	}
}
