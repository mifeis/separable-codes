package lib

import "fmt"

func LogType(tipus int) {
	fmt.Println("\n--------------------------------------------------------------------------------")
	fmt.Println("\t\t\t\t\t\t\t\t\tType", tipus)
	fmt.Println("--------------------------------------------------------------------------------")
}

func LogCombinations(g Combi, combs []Combi) int {
	var total int
	fmt.Println()
	for _, c := range combs {
		fmt.Println("\t\t", g.Rows[:], "|", c.Rows)
		total++
	}
	return total
}
