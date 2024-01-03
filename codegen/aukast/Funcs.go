package aukast

type (
	AstWithBreakFilterFunc func(elem *AstElem) (isTake, isBreak bool)
)
