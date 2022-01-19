package combinations

import (
	"github.com/mifeis/Separable-Codes/lib"
	"gonum.org/v1/gonum/stat/combin"
)

//Funció que retorna els elements disjunts en grups de GROUP elements de l'array inicial
func List(initial []int, reps int) []lib.Map {
	var combins map[int][]int
	arraymap := []lib.Map{}

	//First combinations of GROUP elements
	groups := getCombins(initial, []int{}, reps, 0)
	for _, g := range groups {
		if reps <= len(g) {
			list := make(map[int][][]int)
			combins = getCombins(lib.RemoveSlice(initial, g[:]), g, 0, reps)
			for _, c := range combins {
				list[len(c)] = append(list[len(c)], c)
			}
			m := lib.Map{
				First:   g,
				Seconds: list,
			}
			arraymap = append(arraymap, m)
		}
	}
	return arraymap
}

//compta els casos incomplerts i complerts
func getCombins(remaining []int, g []int, init int, reps int) map[int][]int {
	var key int
	maxLen := lib.GROUP - len(g)
	combins := make(map[int][]int, 1000)

	if len(g) != 0 {
		maxLen = len(g)
	}

	in := combin.Combinations(maxLen, reps)
	for p := 0; p < len(in); p++ {
		var slice1 []int
		//valors no disjunts
		for r := 0; r < reps; r++ {
			slice1 = append(slice1, g[in[p][r]])
		}
		if len(slice1) != 0 {
			combins[key] = slice1[:]
			key++
		}

		//valors disjunts
		for l := maxLen - reps; l > init; l-- {
			indexes := combin.Combinations(len(remaining), l)
			for _, index := range indexes {
				var slice2 []int
				slice2 = append(slice2, slice1...)
				for _, v := range index {
					slice2 = append(slice2, remaining[v])
				}
				combins[key] = slice2[:]
				key++
			}
		}
	}
	return combins
}
