package chmodhelper

// IsPartialMatchVariableAttr
//
//	givenVarAttr can have wildcards "*"
//	 On wildcard present comparison will ignore for that segment.
//
//	Example (will consider this a match):
//	 - givenVarAttr: (rwx : "r*x"),
//	 - rwx         : (rwx : "r-x")
func IsPartialMatchVariableAttr(
	givenVarAttr *VarAttribute,
	rwx string,
) bool {
	r, w, x := ExpandCharRwx(rwx)

	read := givenVarAttr.isRead.ToByteCondition(
		ReadChar,
		NopChar,
		WildcardChar)
	write := givenVarAttr.isWrite.ToByteCondition(
		WriteChar,
		NopChar,
		WildcardChar)
	execute := givenVarAttr.isExecute.ToByteCondition(
		ExecuteChar,
		NopChar,
		WildcardChar,
	)

	isRead := givenVarAttr.isRead.IsWildcard() || (read == r)
	isWrite := givenVarAttr.isWrite.IsWildcard() || (write == w)
	isExecute := givenVarAttr.isExecute.IsWildcard() || (execute == x)

	return isRead &&
		isWrite &&
		isExecute
}
