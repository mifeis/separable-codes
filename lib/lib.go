package lib

const WORDS = 16
const GROUP = 3

func RemoveSlice(original []int, g []int) []int {
	var remaining []int
	remaining = append(remaining, original[:]...)

	for i, elem := range g {
		RemoveIndex(remaining, elem-i-1)
	}

	return remaining[:WORDS-GROUP]
}

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}
