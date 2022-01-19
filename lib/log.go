package lib

import (
	"fmt"
)

func LogTipus(k1 int, k2 int) {
	/*	fmt.Println("\n--------------------------------------------------------------------------------")
		fmt.Println("\t\t\t\t\t\t\t\tmax length:", k)
		fmt.Println("--------------------------------------------------------------------------------")
	*/
	fmt.Println(k1, "x", k2, "->")
}

func LogFavs(favs int, nofavs int) {
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println("Total favorable cases:", favs)
	fmt.Println("Total desfavorable cases:", nofavs)
	fmt.Println("--------------------------------------------------------------------------------")
}

func LogDeps(total int) {
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println("Total dependent pairs:", total)
	fmt.Println("--------------------------------------------------------------------------------")
}

func LogCombinations(arraymap []Map, reps int) int {
	var res int

	//	fmt.Println("\t\t\t\t\t\tCombinations for", reps, "ELEMENT REPETITIONS")
	fmt.Println(reps, "element repetitions:")

	arraymaps := Sort(arraymap)
	for k, am := range arraymaps {
		total := make(map[int]int)
		for _, m := range am {
			for k2 := range m.Seconds {
				for _, _ = range m.Seconds[k2] {
					//					fmt.Println("\t\t", m.First, "|", s)
					total[k2]++
				}
			}
			//			fmt.Println()
		}
		var all int
		for l, v := range total {
			if l == k {
				v = v / 2
			}
			all += v
			LogTipus(k, l)
			fmt.Println(v)
		}

		//		fmt.Println("Total:", all)
		fmt.Println("Total", all)
		res += all
	}
	return res
}
