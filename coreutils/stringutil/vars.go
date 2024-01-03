package stringutil

var (
	ReplaceTemplate       = replaceTemplate{}
	WhiteSpaceReplacerMap = []KeyValReplacer{
		{
			Key:   "\t",
			Value: "",
		},
		{
			Key:   "\n",
			Value: "",
		},
		{
			Key:   "\r",
			Value: "",
		},
		{
			Key:   "  ",
			Value: "",
		},
		{
			Key:   " ",
			Value: "",
		},
	}
)
