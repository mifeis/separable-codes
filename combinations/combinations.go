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
	groups := GetGroups(initial, []int{}, 1)
	for _, g := range groups {
		var list [][]int
		combins = GetGroups(lib.RemoveSlice(initial, g[:]), g, t)
		for i := range combins {
			list = append(list, combins[i])
		}
		//		list = append(list, combins...)
		m := lib.Map{
			First:   g,
			Seconds: list,
		}
		arraymap = append(arraymap, m)
	}
	return arraymap
}

//comptar els casos incomplerts i els complerts (done)
func GetGroups(remaining []int, g []int, t int) map[int][]int {
	var key int
	combins := make(map[int][]int, 1000)
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
	}

	return combins
}
