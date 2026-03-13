package pathinternaltests

import (
	"os"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/internal/pathinternal"
)

// ── Clean ──

func Test_Cov2_Clean_Empty(t *testing.T) {
	actual := args.Map{"result": pathinternal.Clean("")}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "Clean empty", actual)
}

func Test_Cov2_Clean_Valid(t *testing.T) {
	result := pathinternal.Clean("/tmp/test/../other")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Clean valid", actual)
}

// ── Join ──

func Test_Cov2_Join(t *testing.T) {
	result := pathinternal.Join("a", "b", "c")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Join", actual)
}

// ── JoinTemp ──

func Test_Cov2_JoinTemp(t *testing.T) {
	result := pathinternal.JoinTemp("subdir", "file.txt")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "JoinTemp", actual)
}

// ── ParentDir ──

func Test_Cov2_ParentDir_Empty(t *testing.T) {
	actual := args.Map{"result": pathinternal.ParentDir("")}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "ParentDir empty", actual)
}

func Test_Cov2_ParentDir_Valid(t *testing.T) {
	result := pathinternal.ParentDir("/tmp/test/file.txt")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ParentDir valid", actual)
}

// ── Relative ──

func Test_Cov2_Relative(t *testing.T) {
	result := pathinternal.Relative("/tmp", "/tmp/sub/file.txt")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Relative", actual)
}

// ── IsPathExists ──

func Test_Cov2_IsPathExists_True(t *testing.T) {
	actual := args.Map{"exists": pathinternal.IsPathExists(os.TempDir())}
	expected := args.Map{"exists": true}
	expected.ShouldBeEqual(t, 0, "IsPathExists true", actual)
}

func Test_Cov2_IsPathExists_False(t *testing.T) {
	actual := args.Map{"exists": pathinternal.IsPathExists("/nonexistent/path/xyz")}
	expected := args.Map{"exists": false}
	expected.ShouldBeEqual(t, 0, "IsPathExists false", actual)
}

// ── GetTemp ──

func Test_Cov2_GetTemp(t *testing.T) {
	result := pathinternal.GetTemp()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetTemp", actual)
}

// ── RemoveDirIf ──

func Test_Cov2_RemoveDirIf_FalseCondition(t *testing.T) {
	err := pathinternal.RemoveDirIf(false, "/tmp/nonexist", "test")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "RemoveDirIf false condition", actual)
}

func Test_Cov2_RemoveDirIf_NonExistentDir(t *testing.T) {
	err := pathinternal.RemoveDirIf(true, "/tmp/nonexistent_dir_xyz_test", "test")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "RemoveDirIf non-existent dir", actual)
}
