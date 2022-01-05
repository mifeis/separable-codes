package lib

import "fmt"

func LogType(tipus int) {
	fmt.Println("\n--------------------------------------------------------------------------------")
	fmt.Println("\t\t\t\t\t\t\t\t\tType", tipus)
	fmt.Println("--------------------------------------------------------------------------------")
}

func LogCombinations(g []int, combs [][GROUP]int) int {
	var total int
	fmt.Println()
	for _, c := range combs {
		fmt.Println("\t\t", g, "|", c)
		total++
	}
	return total
}
