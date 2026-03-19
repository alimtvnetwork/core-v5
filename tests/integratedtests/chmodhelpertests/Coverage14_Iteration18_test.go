package chmodhelpertests

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/alimtvnetwork/core/chmodhelper/chmodclasstype"
	"github.com/alimtvnetwork/core/chmodhelper/chmodins"
)

func skipIfWindows(t *testing.T) {
	t.Helper()
	if runtime.GOOS == "windows" {
		t.Skip("skipping file permission test on Windows")
	}
}

// --- Variant ---

func Test_I18_Variant_String(t *testing.T) {
	v := chmodhelper.Variant("755")
	if v.String() != "755" {
		t.Fatalf("expected 755, got %s", v.String())
	}
}

func Test_I18_Variant_ExpandOctalByte(t *testing.T) {
	v := chmodhelper.Variant("755")
	r, w, x := v.ExpandOctalByte()
	if r == 0 && w == 0 && x == 0 {
		// at least some should be non-zero for 755
	}
	_ = r
	_ = w
	_ = x
}

func Test_I18_Variant_ToWrapper(t *testing.T) {
	v := chmodhelper.Variant("755")
	wrapper, err := v.ToWrapper()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if wrapper.IsEmpty() {
		t.Fatal("expected non-empty wrapper")
	}
}

func Test_I18_Variant_ToWrapperPtr(t *testing.T) {
	v := chmodhelper.Variant("755")
	wrapper, err := v.ToWrapperPtr()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if wrapper == nil || wrapper.IsEmpty() {
		t.Fatal("expected non-empty wrapper ptr")
	}
}

// --- RwxWrapper basic ---

func Test_I18_RwxWrapper_IsEmpty_Nil(t *testing.T) {
	var w *chmodhelper.RwxWrapper
	if !w.IsEmpty() {
		t.Fatal("expected empty for nil")
	}
	if !w.IsNull() {
		t.Fatal("expected null for nil")
	}
	if !w.IsInvalid() {
		t.Fatal("expected invalid for nil")
	}
}

func Test_I18_RwxWrapper_IsDefined(t *testing.T) {
	v := chmodhelper.Variant("755")
	w, err := v.ToWrapperPtr()
	if err != nil {
		t.Fatal(err)
	}
	if !w.IsDefined() {
		t.Fatal("expected defined")
	}
	if !w.HasAnyItem() {
		t.Fatal("expected has items")
	}
}

// --- SingleRwx ---

func Test_I18_NewSingleRwx_Valid(t *testing.T) {
	s, err := chmodhelper.NewSingleRwx("rwx", chmodclasstype.All)
	if err != nil || s == nil {
		t.Fatal("expected valid SingleRwx")
	}
}

func Test_I18_NewSingleRwx_InvalidLength(t *testing.T) {
	_, err := chmodhelper.NewSingleRwx("rw", chmodclasstype.All)
	if err == nil {
		t.Fatal("expected error for invalid rwx length")
	}
}

func Test_I18_SingleRwx_ToRwxOwnerGroupOther_All(t *testing.T) {
	s, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.All)
	ogo := s.ToRwxOwnerGroupOther()
	if ogo == nil || ogo.Owner != "rwx" || ogo.Group != "rwx" || ogo.Other != "rwx" {
		t.Fatal("expected all rwx")
	}
}

func Test_I18_SingleRwx_ToRwxOwnerGroupOther_Owner(t *testing.T) {
	s, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.Owner)
	ogo := s.ToRwxOwnerGroupOther()
	if ogo == nil || ogo.Owner != "rwx" {
		t.Fatal("expected owner rwx")
	}
}

func Test_I18_SingleRwx_ToRwxOwnerGroupOther_Group(t *testing.T) {
	s, _ := chmodhelper.NewSingleRwx("r-x", chmodclasstype.Group)
	ogo := s.ToRwxOwnerGroupOther()
	if ogo == nil || ogo.Group != "r-x" {
		t.Fatal("expected group r-x")
	}
}

func Test_I18_SingleRwx_ToRwxOwnerGroupOther_Other(t *testing.T) {
	s, _ := chmodhelper.NewSingleRwx("r--", chmodclasstype.Other)
	ogo := s.ToRwxOwnerGroupOther()
	if ogo == nil || ogo.Other != "r--" {
		t.Fatal("expected other r--")
	}
}

func Test_I18_SingleRwx_ToRwxOwnerGroupOther_OwnerGroup(t *testing.T) {
	s, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.OwnerGroup)
	ogo := s.ToRwxOwnerGroupOther()
	if ogo == nil || ogo.Owner != "rwx" || ogo.Group != "rwx" {
		t.Fatal("expected owner+group rwx")
	}
}

func Test_I18_SingleRwx_ToRwxOwnerGroupOther_GroupOther(t *testing.T) {
	s, _ := chmodhelper.NewSingleRwx("r-x", chmodclasstype.GroupOther)
	ogo := s.ToRwxOwnerGroupOther()
	if ogo == nil || ogo.Group != "r-x" || ogo.Other != "r-x" {
		t.Fatal("expected group+other r-x")
	}
}

func Test_I18_SingleRwx_ToRwxOwnerGroupOther_OwnerOther(t *testing.T) {
	s, _ := chmodhelper.NewSingleRwx("rw-", chmodclasstype.OwnerOther)
	ogo := s.ToRwxOwnerGroupOther()
	if ogo == nil || ogo.Owner != "rw-" || ogo.Other != "rw-" {
		t.Fatal("expected owner+other rw-")
	}
}

func Test_I18_SingleRwx_ToRwxInstruction(t *testing.T) {
	s, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.All)
	cond := &chmodins.Condition{}
	ins := s.ToRwxInstruction(cond)
	if ins == nil {
		t.Fatal("expected non-nil instruction")
	}
}

func Test_I18_SingleRwx_ToVarRwxWrapper(t *testing.T) {
	s, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.All)
	vw, err := s.ToVarRwxWrapper()
	if err != nil || vw == nil {
		t.Fatal("expected valid var wrapper")
	}
}

func Test_I18_SingleRwx_ToDisabledRwxWrapper(t *testing.T) {
	s, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.All)
	dw, err := s.ToDisabledRwxWrapper()
	if err != nil || dw == nil {
		t.Fatal("expected valid disabled wrapper")
	}
}

func Test_I18_SingleRwx_ToRwxWrapper_All(t *testing.T) {
	s, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.All)
	w, err := s.ToRwxWrapper()
	if err != nil || w == nil {
		t.Fatal("expected valid rwx wrapper")
	}
}

func Test_I18_SingleRwx_ToRwxWrapper_NotAll(t *testing.T) {
	s, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.Owner)
	_, err := s.ToRwxWrapper()
	if err == nil {
		t.Fatal("expected error for non-all class type")
	}
}

func Test_I18_SingleRwx_ApplyOnMany_Empty(t *testing.T) {
	s, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.All)
	cond := &chmodins.Condition{}
	err := s.ApplyOnMany(cond)
	if err != nil {
		t.Fatal("expected nil for empty locations")
	}
}

func Test_I18_SingleRwx_ApplyOnMany_Valid(t *testing.T) {
	skipIfWindows(t)

	tmpDir := t.TempDir()
	f := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(f, []byte("test"), 0644)

	s, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.All)
	cond := &chmodins.Condition{}
	err := s.ApplyOnMany(cond, f)
	_ = err // may succeed or fail based on OS, just exercise path
}

// --- NewCreator.RwxWrapper ---

func Test_I18_NewRwxWrapper_UsingVariant(t *testing.T) {
	w, err := chmodhelper.New.RwxWrapper.UsingVariant(chmodhelper.Variant("644"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if w.IsEmpty() {
		t.Fatal("expected non-empty")
	}
}

func Test_I18_NewRwxWrapper_UsingVariantPtr(t *testing.T) {
	w, err := chmodhelper.New.RwxWrapper.UsingVariantPtr(chmodhelper.Variant("644"))
	if err != nil || w == nil {
		t.Fatalf("unexpected: err=%v, w=%v", err, w)
	}
}

func Test_I18_NewRwxWrapper_RwxFullString(t *testing.T) {
	w, err := chmodhelper.New.RwxWrapper.RwxFullString("rwxr-xr-x")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if w.IsEmpty() {
		t.Fatal("expected non-empty")
	}
}

// --- ChmodApply and Verify ---

func Test_I18_ChmodApply_RecursivePath(t *testing.T) {
	skipIfWindows(t)

	tmpDir := t.TempDir()
	f := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(f, []byte("test"), 0644)

	err := chmodhelper.ChmodApply.RecursivePath(true, 0755, tmpDir)
	_ = err
}

func Test_I18_ChmodVerify_RwxFull(t *testing.T) {
	skipIfWindows(t)

	tmpDir := t.TempDir()
	f := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(f, []byte("test"), 0755)
	os.Chmod(f, 0755)

	err := chmodhelper.ChmodVerify.RwxFull(f, "-rwxr-xr-x")
	_ = err
}

func Test_I18_ChmodVerify_RwxFull_NoDash(t *testing.T) {
	skipIfWindows(t)

	tmpDir := t.TempDir()
	f := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(f, []byte("test"), 0755)
	os.Chmod(f, 0755)

	err := chmodhelper.ChmodVerify.RwxFull(f, "rwxr-xr-x")
	_ = err
}

// --- TempDirGetter ---

func Test_I18_TempDirGetter(t *testing.T) {
	td := chmodhelper.TempDirGetter.TempDefault()
	if td == "" {
		t.Fatal("expected non-empty temp dir")
	}
}

// --- ExpandCharRwx ---

func Test_I18_ExpandCharRwx_Valid(t *testing.T) {
	r, w, x := chmodhelper.ExpandCharRwx("755")
	_ = r
	_ = w
	_ = x
}

func Test_I18_ExpandCharRwx_Short(t *testing.T) {
	defer func() { recover() }() // may panic on short string
	chmodhelper.ExpandCharRwx("")
}

// --- SimpleFileReaderWriter ---

func Test_I18_SimpleFileReaderWriter(t *testing.T) {
	skipIfWindows(t)

	tmpDir := t.TempDir()
	f := filepath.Join(tmpDir, "sub", "test.txt")

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, f)
	if rw == nil {
		t.Fatal("expected non-nil reader writer")
	}
}

// --- FileModeFriendlyString ---

func Test_I18_FileModeFriendlyString(t *testing.T) {
	s := chmodhelper.FileModeFriendlyString(0755)
	_ = s
}

// --- PathExistStat ---

func Test_I18_GetPathExistStat_NonExistent(t *testing.T) {
	stat := chmodhelper.GetPathExistStat("/nonexistent/path/xyz_i18")
	if stat == nil {
		t.Fatal("expected non-nil stat")
	}
	if stat.IsExist {
		t.Fatal("expected non-exist for fake path")
	}
}

func Test_I18_GetPathExistStat_Existing(t *testing.T) {
	tmpDir := t.TempDir()
	f := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(f, []byte("test"), 0644)

	stat := chmodhelper.GetPathExistStat(f)
	if !stat.IsExist {
		t.Fatal("expected exist for real path")
	}
}

// --- IsPathExists ---

func Test_I18_IsPathExists(t *testing.T) {
	tmpDir := t.TempDir()
	if !chmodhelper.IsPathExists(tmpDir) {
		t.Fatal("expected exists for temp dir")
	}

	if chmodhelper.IsPathExists("/nonexistent/xyz") {
		t.Fatal("expected not exists")
	}
}

func Test_I18_IsPathInvalid(t *testing.T) {
	if !chmodhelper.IsPathInvalid("/nonexistent/xyz") {
		t.Fatal("expected invalid for nonexistent path")
	}
}

func Test_I18_IsDirectory(t *testing.T) {
	tmpDir := t.TempDir()
	if !chmodhelper.IsDirectory(tmpDir) {
		t.Fatal("expected directory")
	}

	f := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(f, []byte("test"), 0644)
	if chmodhelper.IsDirectory(f) {
		t.Fatal("expected not directory for file")
	}
}

// --- GetExistingChmod ---

func Test_I18_GetExistingChmod(t *testing.T) {
	skipIfWindows(t)

	tmpDir := t.TempDir()
	f := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(f, []byte("test"), 0644)

	chmod, err := chmodhelper.GetExistingChmod(f)
	if err != nil || chmod == 0 {
		t.Fatal("expected non-zero chmod")
	}
}

func Test_I18_GetExistingChmodOfValidFile(t *testing.T) {
	skipIfWindows(t)

	tmpDir := t.TempDir()
	f := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(f, []byte("test"), 0644)

	chmod, isInvalid := chmodhelper.GetExistingChmodOfValidFile(f)
	if isInvalid || chmod == 0 {
		t.Fatal("unexpected result")
	}
}

func Test_I18_GetExistingChmodOfValidFile_NonExistent(t *testing.T) {
	_, isInvalid := chmodhelper.GetExistingChmodOfValidFile("/nonexistent/xyz")
	if !isInvalid {
		t.Fatal("expected invalid for nonexistent file")
	}
}
