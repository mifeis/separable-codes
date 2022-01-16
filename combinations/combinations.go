package combinations

import (
	"github.com/mifeis/Separable-Codes/lib"
	"gonum.org/v1/gonum/stat/combin"
)

//Funció que retorna els elements disjunts en grups de GROUP elements de l'array inicial
func List(initial []int, k int, reps int) []lib.Map {
	var combins map[int][]int
	arraymap := []lib.Map{}

	//First combinations of GROUP elements
	groups := getCombins(initial, []int{}, k, 0)
	for _, g := range groups {
		var list [][]int
		combins = getCombins(lib.RemoveSlice(initial, g[:]), g, k, reps)
		for i := range combins {
			list = append(list, combins[i])
		}
		m := lib.Map{
			First:   g,
			Seconds: list,
		}
		arraymap = append(arraymap, m)
	}
	return arraymap
}

//compta els casos incomplerts i complerts
func getCombins(remaining []int, g []int, k int, reps int) map[int][]int {
	var key int
	combins := make(map[int][]int, 1000)
	//argum GROUP en cmombinations cambiara depenent del index del primer grup (passat per argument desde main)
	in := combin.Combinations(lib.GROUP, reps) //group-k?
	for p := 0; p < len(in); p++ {
		var slice1 []int
		//valors no disjunts
		for r := 0; r < reps; r++ {
			slice1 = append(slice1, g[in[p][r]])
		}
		//valors disjunts
		//Casos complerts:
		indexes := combin.Combinations(len(remaining), lib.GROUP-(reps)-k)
		for _, index := range indexes {
			var slice2 []int
			slice2 = append(slice2, slice1...)
			for _, v := range index {
				slice2 = append(slice2, remaining[v])
			}
			combins[key] = slice2[:]
			key++
		}
		//treure if i tenir en compte els casos de 2 i 3 grups p 1 i 3 (al revés)
		if len(g) != 0 {
			for l := 0; l < lib.GROUP-1-k; l++ {
				indexes = combin.Combinations(len(remaining), l+1)
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

	}
	return combins
}
