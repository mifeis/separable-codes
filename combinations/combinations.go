package combinations

import (
	"github.com/mifeis/Separable-Codes/lib"
	"gonum.org/v1/gonum/stat/combin"
)

//Funció que retorna els elements disjunts en grups de GROUP elements de l'array inicial
func List(initial []int, t int) map[[lib.GROUP]int][][lib.GROUP]int {
	var combins [][lib.GROUP]int
	arraymap := make(map[[lib.GROUP]int][][lib.GROUP]int)

	//First combinations of GROUP elements
	groups := GetGroups(initial, [lib.GROUP]int{}, 1)
	for _, g := range groups {
		var list [][lib.GROUP]int
		combins = GetGroups(lib.RemoveSlice(initial, g[:]), g, t)
		list = append(list, combins...)
		arraymap[g] = list
	}
	return arraymap
}

func GetGroups(remaining []int, g [lib.GROUP]int, t int) [][lib.GROUP]int {
	var combins [][lib.GROUP]int
	var slice [lib.GROUP]int
	in := combin.Combinations(lib.GROUP, t-1)
	for p := 0; p < len(in); p++ {
		var init int
		//valors no disjunts
		for r := 0; r < t-1; r++ {
			slice[r] = g[in[p][r]]
			init++
		}
		//valors disjunts
		indexes := combin.Combinations(len(remaining), lib.GROUP-(t-1))
		for _, index := range indexes { //{0,1,2},{0,1,3}...
			for i, v := range index {
				slice[init+i] = remaining[v]
			}
			combins = append(combins, slice)
		}
	}
	return combins
}
