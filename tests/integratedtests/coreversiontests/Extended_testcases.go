package coreversiontests

type versionCreationCase struct {
	name          string
	input         string
	expectedMajor int
	expectedMinor int
	expectedPatch int
}

var versionCreationCases = []versionCreationCase{
	{"full version", "v1.2.3", 1, 2, 3},
	{"without v prefix", "1.2.3", 1, 2, 3},
	{"major only", "v5", 5, 0, 0},
	{"major.minor", "v1.2", 1, 2, 0},
	{"with build", "v1.2.3.4", 1, 2, 3},
	{"empty", "", -1, -1, -1},
	{"just v", "v", -1, -1, -1},
	{"spaces", "  v1.2.3  ", 1, 2, 3},
}
