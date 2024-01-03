package stringutil

var (
	ReplaceTemplate             = replaceTemplate{}
	WhitespaceReplacerKeyValues = []KeyValReplacer{
		{
			Key:   "\t",
			Value: "",
		},
		{
			Key:   "\n",
			Value: "",
		},
		{
			Key:   "\r\n",
			Value: "",
		},
		{
			Key:   "\r",
			Value: "",
		},
		{
			Key:   "   ",
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
