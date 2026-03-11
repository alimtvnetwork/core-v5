package regexnewtests

var params = struct {
	// Input keys
	pattern      string
	compareInput string
	content      string
	comparing    string
	input        string
	isLock       string
	customizer   string

	// Expected keys — state
	patternResult string
	stringResult  string
	isDefined     string
	isUndefined   string
	isApplicable  string
	isCompiled    string
	regexNotNil   string
	regexIsNil    string

	// Expected keys — matching
	isMatch            string
	isFailedMatch      string
	isMatchBytes       string
	isFailedMatchBytes string
	isFailed           string
	firstMatch         string
	isInvalidMatch     string

	// Expected keys — error / validation
	hasError      string
	isInvalid     string
	isNoError     string
	isCustomError string

	// Expected keys — misc
	isNotEmpty string
	panicked   string
}{
	// Input keys
	pattern:      "pattern",
	compareInput: "compareInput",
	content:      "content",
	comparing:    "comparing",
	input:        "input",
	isLock:       "isLock",
	customizer:   "customizer",

	// Expected keys — state
	patternResult: "patternResult",
	stringResult:  "stringResult",
	isDefined:     "isDefined",
	isUndefined:   "isUndefined",
	isApplicable:  "isApplicable",
	isCompiled:    "isCompiled",
	regexNotNil:   "regexNotNil",
	regexIsNil:    "regexIsNil",

	// Expected keys — matching
	isMatch:            "isMatch",
	isFailedMatch:      "isFailedMatch",
	isMatchBytes:       "isMatchBytes",
	isFailedMatchBytes: "isFailedMatchBytes",
	isFailed:           "isFailed",
	firstMatch:         "firstMatch",
	isInvalidMatch:     "isInvalidMatch",

	// Expected keys — error / validation
	hasError:      "hasError",
	isInvalid:     "isInvalid",
	isNoError:     "isNoError",
	isCustomError: "isCustomError",

	// Expected keys — misc
	isNotEmpty: "isNotEmpty",
	panicked:   "panicked",
}
