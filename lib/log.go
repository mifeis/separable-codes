package lib

import (
	"fmt"
)

func LogTipus(tipus int) {
	fmt.Println("\n--------------------------------------------------------------------------------")
	fmt.Println("\t\t\t\t\t\t\t\t\tType", tipus)
	fmt.Println("--------------------------------------------------------------------------------")
}

func LogCombinations(m Map) int {
	var total int
	fmt.Println()
	for _, c := range m.Seconds {
		fmt.Println("\t\t", m.First, "|", c)
		total++
	}
	return total
}
