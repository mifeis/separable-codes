package lib

import (
	"fmt"
	"log"
)

func LogTipus(tipus int, k int) {
	fmt.Println("\n--------------------------------------------------------------------------------")
	fmt.Println("\t\t\t\t\t\t\t\tFirst group length:", k)
	fmt.Println("\t\t\t\t\t\tCombinations for", tipus, "element repetitions")
	fmt.Println("--------------------------------------------------------------------------------")

	log.Print("First group length: ", k)
	log.Print("Combinations for ", tipus, " element repetitions")
}

func LogCombinations(arraymap []Map) int {
	var total int

	fmt.Println("Combinations:")
	for _, m := range arraymap {
		fmt.Println()
		for _, c := range m.Seconds {
			fmt.Println("\t\t", m.First, "|", c)
			total++
		}
	}
	fmt.Println("Total cases:", total/2)
	log.Println(total / 2)

	return total
}
