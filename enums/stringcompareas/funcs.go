package stringcompareas

type (
	IsLineCompareFunc func(
		content,
		searchComparingLine string,
		isCaseSensitive bool,
	) bool

	IsDynamicCompareFunc func(
		index int,
		contentLine string,
		compareAs Variant,
	) bool
)
