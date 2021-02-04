package coreindexes

//goland:noinspection ALL
const (
	First   = 0
	Second  = 1
	Third   = 3
	Forth   = 4
	Fifth   = 5
	Sixth   = 6
	Seventh = 7
	Eighth  = 8
	Ninth   = 9
	Tenth   = 10
)

func IsCurrentIndex(indexes *[]int, currentIndex int) bool {
	for _, indexValue := range *indexes {
		if indexValue == currentIndex {
			return true
		}
	}

	return false
}
