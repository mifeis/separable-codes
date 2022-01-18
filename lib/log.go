package lib

import (
	"fmt"
	"log"
)

func LogTipus(k int) {
	fmt.Println("\n--------------------------------------------------------------------------------")
	fmt.Println("\t\t\t\t\t\t\t\tmax length:", k)
	fmt.Println("--------------------------------------------------------------------------------")

	log.Print("max length: ", k)
}

func LogCombinations(arraymap []Map, reps int) int {
	var res int
	arraymaps := make(map[int][]Map)

	fmt.Println("\t\t\t\t\t\tCombinations for", reps, "ELEMENT REPETITIONS")
	log.Print(reps, " element repetitions:")

	//ordena els arraymaps segons el tamany del primer grup
	for _, m := range arraymap {
		arraymaps[len(m.First)] = append(arraymaps[len(m.First)], m)
	}

	for k, am := range arraymaps {
		LogTipus(k)
		total := make(map[int]int)
		//{1,2,3}	{4,5,6}{4,5}{4}...
		//{1,2,4}	{4,5}{6}{3,6,7}{7}...
		//{1,2,6}	{4}{7}{8}{4,5,7}...
		for _, m := range am {
			//{1,2,3}	{4,5,6}{4,5}{4}...
			for k2 := range m.Seconds {
				for _, s := range m.Seconds[k2] {
					fmt.Println("\t\t", m.First, "|", s)
					total[k2]++
				}
			}
			fmt.Println()
		}
		var all int
		for l, v := range total {
			if (l == k) && (k != reps) { //si contabilitzem {1},{1} exemple
				v = v / 2
			}
			all += v
		}
		log.Println(all)
		res += all
	}
	return res
}
