package chmodhelpertests

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/alimtvnetwork/core/chmodhelper/chmodins"
	"github.com/alimtvnetwork/core/coretests/args"
)

func tempDir(t *testing.T) string {
	t.Helper()
	d := t.TempDir()
	return d
}

// ── CreateDirWithFiles ──

func Test_Cov6_CreateDirWithFiles(t *testing.T) {
	dir := filepath.Join(tempDir(t), "sub")
	err := chmodhelper.CreateDirWithFiles(false, 0755, &chmodhelper.DirWithFiles{Dir: dir, Files: []string{"a.txt"}})
	actual := args.Map{"noErr": err == nil, "exists": chmodhelper.IsPathExists(filepath.Join(dir, "a.txt"))}
	expected := args.Map{"noErr": true, "exists": true}
	expected.ShouldBeEqual(t, 0, "CreateDirWithFiles returns non-empty -- with args", actual)
}

func Test_Cov6_CreateDirWithFiles_NoFiles(t *testing.T) {
	dir := filepath.Join(tempDir(t), "empty")
	err := chmodhelper.CreateDirWithFiles(false, 0755, &chmodhelper.DirWithFiles{Dir: dir, Files: []string{}})
	actual := args.Map{"noErr": err == nil, "dirExists": chmodhelper.IsPathExists(dir)}
	expected := args.Map{"noErr": true, "dirExists": true}
	expected.ShouldBeEqual(t, 0, "CreateDirWithFiles_NoFiles returns non-empty -- with args", actual)
}

func Test_Cov6_CreateDirWithFiles_RemoveBefore(t *testing.T) {
	dir := filepath.Join(tempDir(t), "rmdir")
	_ = os.MkdirAll(dir, 0755)
	err := chmodhelper.CreateDirWithFiles(true, 0755, &chmodhelper.DirWithFiles{Dir: dir, Files: []string{"b.txt"}})
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "CreateDirWithFiles_Remove returns non-empty -- with args", actual)
}

// ── CreateDirsWithFiles ──

func Test_Cov6_CreateDirsWithFiles(t *testing.T) {
	base := tempDir(t)
	d1 := chmodhelper.DirWithFiles{Dir: filepath.Join(base, "d1"), Files: []string{"f1.txt"}}
	d2 := chmodhelper.DirWithFiles{Dir: filepath.Join(base, "d2"), Files: []string{"f2.txt"}}
	err := chmodhelper.CreateDirsWithFiles(false, 0755, d1, d2)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "CreateDirsWithFiles returns non-empty -- with args", actual)
}

func Test_Cov6_CreateDirsWithFiles_Empty(t *testing.T) {
	err := chmodhelper.CreateDirsWithFiles(false, 0755)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "CreateDirsWithFiles_Empty returns empty -- with args", actual)
}

// ── CreateDirFilesWithRwxPermission ──

func Test_Cov6_CreateDirFilesWithRwxPermission(t *testing.T) {
	dir := filepath.Join(tempDir(t), "rwx")
	perm := chmodhelper.DirFilesWithRwxPermission{
		DirWithFiles: chmodhelper.DirWithFiles{Dir: dir, Files: []string{"x.txt"}},
		ApplyRwx:     chmodins.RwxOwnerGroupOther{Owner: "rwx", Group: "r-x", Other: "r-x"},
	}
	err := chmodhelper.CreateDirFilesWithRwxPermission(false, &perm)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "CreateDirFilesWithRwxPermission returns non-empty -- with args", actual)
}

// ── CreateDirFilesWithRwxPermissions ──

func Test_Cov6_CreateDirFilesWithRwxPermissions_Nil(t *testing.T) {
	err := chmodhelper.CreateDirFilesWithRwxPermissions(false, nil)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "CreateDirFilesWithRwxPermissions_Nil returns nil -- with args", actual)
}

func Test_Cov6_CreateDirFilesWithRwxPermissions_Empty(t *testing.T) {
	err := chmodhelper.CreateDirFilesWithRwxPermissions(false, []chmodhelper.DirFilesWithRwxPermission{})
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "CreateDirFilesWithRwxPermissions_Empty returns empty -- with args", actual)
}

func Test_Cov6_CreateDirFilesWithRwxPermissions_Valid(t *testing.T) {
	dir := filepath.Join(tempDir(t), "multi")
	perms := []chmodhelper.DirFilesWithRwxPermission{
		{
			DirWithFiles: chmodhelper.DirWithFiles{Dir: dir, Files: []string{"m.txt"}},
			ApplyRwx:     chmodins.RwxOwnerGroupOther{Owner: "rwx", Group: "r-x", Other: "r-x"},
		},
	}
	err := chmodhelper.CreateDirFilesWithRwxPermissions(false, perms)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "CreateDirFilesWithRwxPermissions_Valid returns non-empty -- with args", actual)
}

func Test_Cov6_CreateDirFilesWithRwxPermissionsMust_Valid(t *testing.T) {
	dir := filepath.Join(tempDir(t), "must")
	perms := []chmodhelper.DirFilesWithRwxPermission{
		{
			DirWithFiles: chmodhelper.DirWithFiles{Dir: dir, Files: []string{}},
			ApplyRwx:     chmodins.RwxOwnerGroupOther{Owner: "rwx", Group: "r-x", Other: "r-x"},
		},
	}
	// Should not panic
	chmodhelper.CreateDirFilesWithRwxPermissionsMust(false, perms)
}

// ── DirFilesWithRwxPermission methods ──

func Test_Cov6_DirFilesRwxPermission_GetPaths(t *testing.T) {
	perm := chmodhelper.DirFilesWithRwxPermission{
		DirWithFiles: chmodhelper.DirWithFiles{Dir: "/tmp/test", Files: []string{"a.txt", "b.txt"}},
	}
	paths := perm.GetPaths()
	actual := args.Map{"len": len(paths)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DirFilesRwxPermission_GetPaths returns correct value -- with args", actual)
}

func Test_Cov6_DirFilesRwxPermission_CreatePaths(t *testing.T) {
	dir := filepath.Join(tempDir(t), "create")
	perm := chmodhelper.DirFilesWithRwxPermission{
		DirWithFiles: chmodhelper.DirWithFiles{Dir: dir, Files: []string{}},
		ApplyRwx:     chmodins.RwxOwnerGroupOther{Owner: "rwx", Group: "r-x", Other: "r-x"},
	}
	err := perm.CreatePaths(false)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "DirFilesRwxPermission_CreatePaths returns correct value -- with args", actual)
}

func Test_Cov6_DirFilesRwxPermission_CreateUsingFileMode(t *testing.T) {
	dir := filepath.Join(tempDir(t), "fmode")
	perm := chmodhelper.DirFilesWithRwxPermission{
		DirWithFiles: chmodhelper.DirWithFiles{Dir: dir, Files: []string{}},
	}
	err := perm.CreateUsingFileMode(false, 0755)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "DirFilesRwxPermission_CreateUsingFileMode returns correct value -- with args", actual)
}

// ── DirWithFiles.CreatePaths ──

func Test_Cov6_DirWithFiles_CreatePaths(t *testing.T) {
	dir := filepath.Join(tempDir(t), "dwf")
	dwf := chmodhelper.DirWithFiles{Dir: dir, Files: []string{"z.txt"}}
	err := dwf.CreatePaths(false, 0755)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "DirWithFiles_CreatePaths returns non-empty -- with args", actual)
}

// ── DirFilesWithContent ──

func Test_Cov6_DirFilesWithContent_GetPaths(t *testing.T) {
	dfc := chmodhelper.DirFilesWithContent{
		Dir:   "/tmp/test",
		Files: []chmodhelper.FileWithContent{{RelativePath: "a.txt"}, {RelativePath: "b.txt"}},
	}
	paths := dfc.GetPaths()
	actual := args.Map{"len": len(paths)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DirFilesWithContent_GetPaths returns non-empty -- with args", actual)
}

func Test_Cov6_DirFilesWithContent_Create(t *testing.T) {
	dir := filepath.Join(tempDir(t), "dfc")
	dfc := chmodhelper.DirFilesWithContent{
		Dir:         dir,
		Files:       []chmodhelper.FileWithContent{{RelativePath: "c.txt", FileMode: 0644, Content: []string{"hello"}}},
		DirFileMode: 0755,
	}
	err := dfc.Create(false)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "DirFilesWithContent_Create returns non-empty -- with args", actual)
}

// ── FileWithContent ──

func Test_Cov6_FileWithContent_Methods(t *testing.T) {
	fc := chmodhelper.FileWithContent{RelativePath: "sub/file.txt", FileMode: 0644, Content: []string{"line1", "line2"}}
	actual := args.Map{
		"absPath":  fc.AbsPath("/root") != "",
		"toString": fc.ContentToString() != "",
		"toBytes":  len(fc.ContentToBytes()) > 0,
	}
	expected := args.Map{
		"absPath":  true,
		"toString": true,
		"toBytes":  true,
	}
	expected.ShouldBeEqual(t, 0, "FileWithContent_Methods returns non-empty -- with args", actual)
}

func Test_Cov6_FileWithContent_ReadWrite(t *testing.T) {
	dir := tempDir(t)
	fc := chmodhelper.FileWithContent{RelativePath: "rw.txt", FileMode: 0644, Content: []string{"test content"}}
	writeErr := fc.Write(dir)
	actual := args.Map{"writeOk": writeErr == nil}
	expected := args.Map{"writeOk": true}
	expected.ShouldBeEqual(t, 0, "FileWithContent_Write returns non-empty -- with args", actual)

	readBytes, readErr := fc.Read(dir)
	actual2 := args.Map{"readOk": readErr == nil, "hasContent": len(readBytes) > 0}
	expected2 := args.Map{"readOk": true, "hasContent": true}
	expected2.ShouldBeEqual(t, 0, "FileWithContent_Read returns non-empty -- with args", actual2)
}

// ── GetPathExistStats ──

func Test_Cov6_GetPathExistStats_Empty(t *testing.T) {
	stats := chmodhelper.GetPathExistStats(false)
	actual := args.Map{"len": len(stats)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "GetPathExistStats_Empty returns empty -- with args", actual)
}

func Test_Cov6_GetPathExistStats_ValidOnly(t *testing.T) {
	dir := tempDir(t)
	f := filepath.Join(dir, "exist.txt")
	_ = os.WriteFile(f, []byte("x"), 0644)
	stats := chmodhelper.GetPathExistStats(true, f, "/nonexistent/path/xyz")
	actual := args.Map{"len": len(stats)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "GetPathExistStats_ValidOnly returns non-empty -- with args", actual)
}

func Test_Cov6_GetPathExistStats_All(t *testing.T) {
	dir := tempDir(t)
	f := filepath.Join(dir, "e.txt")
	_ = os.WriteFile(f, []byte("x"), 0644)
	stats := chmodhelper.GetPathExistStats(false, f, "/nonexistent/xyz")
	actual := args.Map{"len": len(stats)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "GetPathExistStats_All returns correct value -- with args", actual)
}

// ── GetExistingChmodRwxWrapperMustPtr ──

func Test_Cov6_GetExistingChmodRwxWrapperMustPtr(t *testing.T) {
	dir := tempDir(t)
	f := filepath.Join(dir, "chmod.txt")
	_ = os.WriteFile(f, []byte("x"), 0644)
	wrapper := chmodhelper.GetExistingChmodRwxWrapperMustPtr(f)
	actual := args.Map{"notNil": wrapper != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "GetExistingChmodRwxWrapperMustPtr returns correct value -- with args", actual)
}

// ── GetFilesChmodRwxFullMapDirect ──

func Test_Cov6_GetFilesChmodRwxFullMapDirect(t *testing.T) {
	dir := tempDir(t)
	f := filepath.Join(dir, "map.txt")
	_ = os.WriteFile(f, []byte("x"), 0644)
	hm, err := chmodhelper.GetFilesChmodRwxFullMapDirect(f)
	actual := args.Map{"noErr": err == nil, "notNil": hm != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "GetFilesChmodRwxFullMapDirect returns correct value -- with args", actual)
}

// ── GetFilteredExistsFilesInfosOnly ──

func Test_Cov6_GetFilteredExistsFilesInfosOnly_Empty(t *testing.T) {
	result := chmodhelper.GetFilteredExistsFilesInfosOnly()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "GetFilteredExistsFilesInfosOnly_Empty returns empty -- with args", actual)
}

func Test_Cov6_GetFilteredExistsFilesInfosOnly_Mixed(t *testing.T) {
	dir := tempDir(t)
	f := filepath.Join(dir, "info.txt")
	_ = os.WriteFile(f, []byte("x"), 0644)
	result := chmodhelper.GetFilteredExistsFilesInfosOnly(f, "/nonexistent/xyz")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "GetFilteredExistsFilesInfosOnly_Mixed returns correct value -- with args", actual)
}

// ── GetRecursivePathsContinueOnError ──

func Test_Cov6_GetRecursivePathsContinueOnError_File(t *testing.T) {
	dir := tempDir(t)
	f := filepath.Join(dir, "single.txt")
	_ = os.WriteFile(f, []byte("x"), 0644)
	paths, err := chmodhelper.GetRecursivePathsContinueOnError(f)
	actual := args.Map{"noErr": err == nil, "len": len(paths)}
	expected := args.Map{"noErr": true, "len": 1}
	expected.ShouldBeEqual(t, 0, "GetRecursivePaths_File returns correct value -- with args", actual)
}

func Test_Cov6_GetRecursivePathsContinueOnError_Dir(t *testing.T) {
	dir := tempDir(t)
	sub := filepath.Join(dir, "sub")
	_ = os.MkdirAll(sub, 0755)
	_ = os.WriteFile(filepath.Join(sub, "a.txt"), []byte("x"), 0644)
	paths, err := chmodhelper.GetRecursivePathsContinueOnError(dir)
	actual := args.Map{"noErr": err == nil, "hasItems": len(paths) > 0}
	expected := args.Map{"noErr": true, "hasItems": true}
	expected.ShouldBeEqual(t, 0, "GetRecursivePaths_Dir returns correct value -- with args", actual)
}

func Test_Cov6_GetRecursivePathsContinueOnError_NonExist(t *testing.T) {
	_, err := chmodhelper.GetRecursivePathsContinueOnError("/nonexistent/path/xyz123")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "GetRecursivePaths_NonExist returns correct value -- with args", actual)
}

// ── SimpleFileReaderWriter ──

func Test_Cov6_SimpleFileReaderWriter_ReadWriteLock(t *testing.T) {
	dir := tempDir(t)
	f := filepath.Join(dir, "lock.json")
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:               0755,
		ChmodFile:              0644,
		ParentDir:              dir,
		FilePath:               f,
		IsMustChmodApplyOnFile: false,
	}

	type data struct{ Name string }
	result := &data{}
	genFunc := func() (any, error) {
		return &data{Name: "generated"}, nil
	}

	err := rw.ReadWriteLock(result, genFunc)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ReadWriteLock returns correct value -- with args", actual)
}

func Test_Cov6_SimpleFileReaderWriter_GetSetLock(t *testing.T) {
	dir := tempDir(t)
	f := filepath.Join(dir, "getset.json")
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:               0755,
		ChmodFile:              0644,
		ParentDir:              dir,
		FilePath:               f,
		IsMustChmodApplyOnFile: false,
	}

	type data struct{ Val int }
	result := &data{}
	err := rw.GetSetLock(result, func() (any, error) {
		return &data{Val: 42}, nil
	})
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "GetSetLock returns correct value -- with args", actual)
}

func Test_Cov6_SimpleFileReaderWriter_CacheGetSetLock(t *testing.T) {
	dir := tempDir(t)
	f := filepath.Join(dir, "cache.json")
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:               0755,
		ChmodFile:              0644,
		ParentDir:              dir,
		FilePath:               f,
		IsMustChmodApplyOnFile: false,
	}

	type data struct{ X string }
	result := &data{}
	err := rw.CacheGetSetLock(result, func() (any, error) {
		return &data{X: "cached"}, nil
	})
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "CacheGetSetLock returns correct value -- with args", actual)
}

// ── DirFilesWithContent.GetFilesChmodMap ──

func Test_Cov6_DirFilesWithContent_GetFilesChmodMap(t *testing.T) {
	dir := tempDir(t)
	f := filepath.Join(dir, "chm.txt")
	_ = os.WriteFile(f, []byte("x"), 0644)
	dfc := chmodhelper.DirFilesWithContent{
		Dir:   dir,
		Files: []chmodhelper.FileWithContent{{RelativePath: "chm.txt", FileMode: 0644}},
	}
	hm := dfc.GetFilesChmodMap()
	actual := args.Map{"notNil": hm != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DirFilesWithContent_GetFilesChmodMap returns non-empty -- with args", actual)
}

// ── DirFilesWithRwxPermission.GetFilesChmodMap ──

func Test_Cov6_DirFilesRwxPermission_GetFilesChmodMap(t *testing.T) {
	dir := tempDir(t)
	f := filepath.Join(dir, "rwxmap.txt")
	_ = os.WriteFile(f, []byte("x"), 0644)
	perm := chmodhelper.DirFilesWithRwxPermission{
		DirWithFiles: chmodhelper.DirWithFiles{Dir: dir, Files: []string{"rwxmap.txt"}},
	}
	hm := perm.GetFilesChmodMap()
	actual := args.Map{"notNil": hm != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DirFilesRwxPermission_GetFilesChmodMap returns correct value -- with args", actual)
}
