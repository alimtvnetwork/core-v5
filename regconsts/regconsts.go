package regconsts

//goland:noinspection ALL
const (
	RegExForEachWordsWithDollarSymbol       = "\\$(\\w+)(\\d*)"  // Selects a full word that starts with a "$" symbol
	EachWordsWithinPercentSymbol            = "\\%(\\w+)(\\d*)%" // Selects a full word that is within two "%" symbol
	NginxBlockStart                         = "\\{\\s*(\\#(\\s)*.*)"
	NginxBlockEnding                        = "\\}\\s*(\\#(\\s)*.*)"
	NginxEmptyBlockOnlyBraces               = "\\{\\}"
	NginxEmptyBlockWithSpaces               = "\\{\\s*\\}"
	NginxEmptyBlockWithCommentAndSpacesOnly = "\\{\\s*\\#+\\s*[aA0-zZ9]*\\s*\\n*\\}"
	IpFormat                                = "(\\d{1,3}\\.{1}){3}\\d{1,3}"                                              // 255.255.255.255
	IpWithSubnetPrefixFormat                = "(\\d{1,3}\\.{1}){3}\\d{1,3}\\/\\d{1,2}"                                   // 255.255.255.255/0
	IpWithPortFormat                        = "(\\d{1,3}\\.{1}){3}\\d{1,3}\\:\\d{1,6}"                                   // 127.0.0.1:8080
	NginxEmptyBlock                         = "(\\{\\}|\\{\\s*\\}|\\{\\s*\\#+\\s*[aA0-zZ9]*.*[aA0-zZ9]*\\.*\\s*\\n*\\})" // Reference : https://regexr.com/5jgh0
)
