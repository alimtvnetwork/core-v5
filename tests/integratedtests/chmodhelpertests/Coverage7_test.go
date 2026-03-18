package chmodhelpertests

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/alimtvnetwork/core/chmodhelper/chmodins"
)

// ── ParseRwxOwnerGroupOtherToFileModeMust ──

func Test_Cov7_ParseRwxOwnerGroupOtherToFileModeMust(t *testing.T) {
	rwx := &chmodins.RwxOwnerGroupOther{
		Owner: "rwx",
		Group: "r-x",
		Other: "r-x",
	}
	mode := chmodhelper.ParseRwxOwnerGroupOtherToFileModeMust(rwx)
	if mode == 0 {
		t.Fatal("expected non-zero mode")
	}
}

// ── ParseBaseRwxInstructionsToExecutors ──

func Test_Cov7_ParseBaseRwxInstructionsToExecutors_Nil(t *testing.T) {
	_, err := chmodhelper.ParseBaseRwxInstructionsToExecutors(nil)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov7_ParseBaseRwxInstructionsToExecutors_Valid(t *testing.T) {
	base := &chmodins.BaseRwxInstructions{
		RwxInstructions: []chmodins.RwxInstruction{
			{
				RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
					Owner: "rwx",
					Group: "r-x",
					Other: "r-x",
				},
				Condition: chmodins.Condition{},
			},
		},
	}
	executors, err := chmodhelper.ParseBaseRwxInstructionsToExecutors(base)
	if err != nil {
		t.Fatal(err)
	}
	if executors == nil {
		t.Fatal("expected non-nil")
	}
}

// ── GetFilesChmodRwxFullMap ──

func Test_Cov7_GetFilesChmodRwxFullMap_Empty(t *testing.T) {
	hm, err := chmodhelper.GetFilesChmodRwxFullMap(nil)
	if err != nil || hm == nil {
		t.Fatal("unexpected")
	}
}

func Test_Cov7_GetFilesChmodRwxFullMap_Valid(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "test.txt")
	_ = os.WriteFile(f, []byte("x"), 0644)
	hm, err := chmodhelper.GetFilesChmodRwxFullMap([]string{f})
	if err != nil || hm == nil {
		t.Fatal("unexpected")
	}
}

func Test_Cov7_GetFilesChmodRwxFullMap_Invalid(t *testing.T) {
	hm, err := chmodhelper.GetFilesChmodRwxFullMap([]string{"/nonexistent/path/xyz123"})
	if err == nil {
		t.Fatal("expected error")
	}
	_ = hm
}

// ── SimpleFileReaderWriter additional methods ──

func Test_Cov7_SimpleFileReaderWriter_InitializeDefault(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "init.txt")
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		FilePath:  f,
	}
	initialized := rw.InitializeDefault(true)
	if initialized.ParentDir == "" {
		t.Fatal("expected parent dir")
	}
}

func Test_Cov7_SimpleFileReaderWriter_InitializeDefaultApplyChmod(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "init2.txt")
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		FilePath:  f,
	}
	initialized := rw.InitializeDefaultApplyChmod()
	if initialized == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_Cov7_SimpleFileReaderWriter_IsExistAndParent(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "exist.txt")
	_ = os.WriteFile(f, []byte("x"), 0644)
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	if !rw.IsExist() {
		t.Fatal("expected exist")
	}
	if !rw.IsParentExist() {
		t.Fatal("expected parent exist")
	}
	if rw.HasPathIssues() {
		t.Fatal("expected no issues")
	}
	if rw.IsPathInvalid() {
		t.Fatal("expected valid")
	}
	if rw.IsParentDirInvalid() {
		t.Fatal("expected valid parent")
	}
	if rw.HasAnyIssues() {
		t.Fatal("expected no issues")
	}
}

func Test_Cov7_SimpleFileReaderWriter_WriteAndRead(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "wr.txt")
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	err := rw.Write([]byte("hello"))
	if err != nil {
		t.Fatal(err)
	}
	content, err := rw.Read()
	if err != nil || string(content) != "hello" {
		t.Fatal("unexpected")
	}
}

func Test_Cov7_SimpleFileReaderWriter_WriteString(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "ws.txt")
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	err := rw.WriteString("world")
	if err != nil {
		t.Fatal(err)
	}
	content, err := rw.ReadString()
	if err != nil || content != "world" {
		t.Fatal("unexpected")
	}
}

func Test_Cov7_SimpleFileReaderWriter_ReadOnExist(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "roe.txt")
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	// File doesn't exist yet
	bytes, err := rw.ReadOnExist()
	if err != nil || bytes != nil {
		t.Fatal("expected nil nil")
	}
	content, err := rw.ReadStringOnExist()
	if err != nil || content != "" {
		t.Fatal("expected empty")
	}
}

func Test_Cov7_SimpleFileReaderWriter_WritePath(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "wp.txt")
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	err := rw.WritePath(false, f, []byte("test"))
	if err != nil {
		t.Fatal(err)
	}
}

func Test_Cov7_SimpleFileReaderWriter_WriteRelativePath(t *testing.T) {
	dir := t.TempDir()
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  filepath.Join(dir, "dummy.txt"),
	}
	err := rw.WriteRelativePath(false, "rel.txt", []byte("data"))
	if err != nil {
		t.Fatal(err)
	}
}

func Test_Cov7_SimpleFileReaderWriter_JoinRelPath(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{
		ParentDir: "/tmp/base",
	}
	p := rw.JoinRelPath("sub/file.txt")
	if p == "" {
		t.Fatal("expected path")
	}
	p2 := rw.JoinRelPath("")
	if p2 == "" {
		t.Fatal("expected path")
	}
}

func Test_Cov7_SimpleFileReaderWriter_WriteAny(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "any.json")
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	type data struct{ Name string }
	err := rw.WriteAny(&data{Name: "test"})
	if err != nil {
		t.Fatal(err)
	}
}

func Test_Cov7_SimpleFileReaderWriter_WriteAnyLock(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "anylock.json")
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	type data struct{ Val int }
	err := rw.WriteAnyLock(&data{Val: 42})
	if err != nil {
		t.Fatal(err)
	}
}

func Test_Cov7_SimpleFileReaderWriter_ReadLock(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "rl.txt")
	_ = os.WriteFile(f, []byte("locked"), 0644)
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	b, err := rw.ReadLock()
	if err != nil || string(b) != "locked" {
		t.Fatal("unexpected")
	}
}

func Test_Cov7_SimpleFileReaderWriter_ReadStringLock(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "rsl.txt")
	_ = os.WriteFile(f, []byte("locked"), 0644)
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	s, err := rw.ReadStringLock()
	if err != nil || s != "locked" {
		t.Fatal("unexpected")
	}
}

func Test_Cov7_SimpleFileReaderWriter_ReadOnExistLock(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "roel.txt")
	_ = os.WriteFile(f, []byte("exists"), 0644)
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	b, err := rw.ReadOnExistLock()
	if err != nil || string(b) != "exists" {
		t.Fatal("unexpected")
	}
}

func Test_Cov7_SimpleFileReaderWriter_ReadStringOnExistLock(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "rsoel.txt")
	_ = os.WriteFile(f, []byte("exists"), 0644)
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	s, err := rw.ReadStringOnExistLock()
	if err != nil || s != "exists" {
		t.Fatal("unexpected")
	}
}

func Test_Cov7_SimpleFileReaderWriter_String(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: "/tmp",
		FilePath:  "/tmp/test.txt",
	}
	s := rw.String()
	if s == "" {
		t.Fatal("expected non-empty string")
	}
}

func Test_Cov7_SimpleFileReaderWriter_StringFilePath(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: "/tmp",
		FilePath:  "/tmp/test.txt",
	}
	s := rw.StringFilePath("/other/path.txt")
	if s == "" {
		t.Fatal("expected non-empty string")
	}
}

func Test_Cov7_SimpleFileReaderWriter_ChmodApplier(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "ca.txt")
	_ = os.WriteFile(f, []byte("x"), 0644)
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	applier := rw.ChmodApplier()
	_ = applier
}

func Test_Cov7_SimpleFileReaderWriter_ChmodVerifier(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "cv.txt")
	_ = os.WriteFile(f, []byte("x"), 0644)
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	verifier := rw.ChmodVerifier()
	_ = verifier
}

func Test_Cov7_SimpleFileReaderWriter_NewPath(t *testing.T) {
	dir := t.TempDir()
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  filepath.Join(dir, "orig.txt"),
	}
	newRw := rw.NewPath(false, filepath.Join(dir, "new.txt"))
	if newRw == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_Cov7_SimpleFileReaderWriter_NewPathJoin(t *testing.T) {
	dir := t.TempDir()
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  filepath.Join(dir, "orig.txt"),
	}
	newRw := rw.NewPathJoin(false, "sub", "file.txt")
	if newRw == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_Cov7_SimpleFileReaderWriter_InitializeDefaultNew(t *testing.T) {
	dir := t.TempDir()
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  filepath.Join(dir, "idn.txt"),
	}
	newRw := rw.InitializeDefaultNew()
	if newRw == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_Cov7_SimpleFileReaderWriter_Set(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "set.json")
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	type data struct{ X int }
	err := rw.Set(&data{X: 1})
	if err != nil {
		t.Fatal(err)
	}
}

func Test_Cov7_SimpleFileReaderWriter_SetLock(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "setlock.json")
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	type data struct{ X int }
	err := rw.SetLock(&data{X: 2})
	if err != nil {
		t.Fatal(err)
	}
}

func Test_Cov7_SimpleFileReaderWriter_Get(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "get.json")
	_ = os.WriteFile(f, []byte(`{"X":42}`), 0644)
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	type data struct{ X int }
	result := &data{}
	err := rw.Get(result)
	if err != nil || result.X != 42 {
		t.Fatal("unexpected")
	}
}

func Test_Cov7_SimpleFileReaderWriter_GetLock(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "getlock.json")
	_ = os.WriteFile(f, []byte(`{"X":99}`), 0644)
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	type data struct{ X int }
	result := &data{}
	err := rw.GetLock(result)
	if err != nil || result.X != 99 {
		t.Fatal("unexpected")
	}
}

func Test_Cov7_SimpleFileReaderWriter_Expire(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "expire.txt")
	_ = os.WriteFile(f, []byte("x"), 0644)
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	rw.Expire()
	if rw.IsExist() {
		t.Fatal("expected removed")
	}
}

func Test_Cov7_SimpleFileReaderWriter_Serialize(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "ser.txt")
	_ = os.WriteFile(f, []byte("data"), 0644)
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	b, err := rw.Serialize()
	if err != nil || len(b) == 0 {
		t.Fatal("unexpected")
	}
}

func Test_Cov7_SimpleFileReaderWriter_SerializeLock(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "serlock.txt")
	_ = os.WriteFile(f, []byte("data"), 0644)
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	b, err := rw.SerializeLock()
	if err != nil || len(b) == 0 {
		t.Fatal("unexpected")
	}
}

func Test_Cov7_SimpleFileReaderWriter_Deserialize(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "deser.json")
	_ = os.WriteFile(f, []byte(`{"X":10}`), 0644)
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	type data struct{ X int }
	result := &data{}
	err := rw.Deserialize(result)
	if err != nil || result.X != 10 {
		t.Fatal("unexpected")
	}
}

func Test_Cov7_SimpleFileReaderWriter_DeserializeLock(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "deserlock.json")
	_ = os.WriteFile(f, []byte(`{"X":20}`), 0644)
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	type data struct{ X int }
	result := &data{}
	err := rw.DeserializeLock(result)
	if err != nil || result.X != 20 {
		t.Fatal("unexpected")
	}
}

func Test_Cov7_SimpleFileReaderWriter_ReadMust(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "readmust.txt")
	_ = os.WriteFile(f, []byte("must"), 0644)
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	b := rw.ReadMust()
	if string(b) != "must" {
		t.Fatal("unexpected")
	}
}

func Test_Cov7_SimpleFileReaderWriter_ReadStringMust(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "readstrmust.txt")
	_ = os.WriteFile(f, []byte("strmust"), 0644)
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	s := rw.ReadStringMust()
	if s != "strmust" {
		t.Fatal("unexpected")
	}
}
