package combinations

import (
	"github.com/mifeis/Separable-Codes/lib"
	"gonum.org/v1/gonum/stat/combin"
)

//Funció que retorna els elements disjunts en grups de GROUP elements de l'array inicial
func List(initial []int, t int) []lib.Map {
	var combins map[int][]int
	arraymap := []lib.Map{}
	//First combinations of GROUP elements
	groups := getCombins(initial, []int{}, 1)
	for _, g := range groups {
		var list [][]int
		combins = getCombins(lib.RemoveSlice(initial, g[:]), g, t)
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
func getCombins(remaining []int, g []int, t int) map[int][]int {
	var key int
	combins := make(map[int][]int, 1000)
	//argum GROUP en cmombinations cambiara depenent del index del primer grup (passat per argument desde main)
	in := combin.Combinations(lib.GROUP, t-1)
	for p := 0; p < len(in); p++ {
		var slice1 []int
		//valors no disjunts
		for r := 0; r < t-1; r++ {
			slice1 = append(slice1, g[in[p][r]])
		}
		//valors disjunts
		//Casos complerts:
		indexes := combin.Combinations(len(remaining), lib.GROUP-(t-1))
		for _, index := range indexes {
			var slice2 []int
			slice2 = append(slice2, slice1...)
			for _, v := range index {
				slice2 = append(slice2, remaining[v])
			}
			combins[key] = slice2[:]
			key++
		}
		//DEV
		//treure if i tenir en compte els casos de 2 i 3 grups p 1 i 3 (al revés)
		/*		if g[0] != 0 {
					//comentar
					//Casos incomplerts:
					for r := init; r < lib.GROUP; r++ {
						slice[r] = 0
					}
					for l := 0; l < lib.GROUP-1; l++ {
						indexes = combin.Combinations(len(remaining), l+1)
						for _, index := range indexes {
							for i, v := range index {
								slice[init+i] = remaining[v]
							}
							combins = append(combins, slice)
						}
					}
				}
		*/

	}

	return combins
}
