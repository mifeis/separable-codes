package combinations

import (
	"github.com/mifeis/Separable-Codes/lib"
)

//Funció que retorna els elements disjunts en grups de GROUP elements de l'array inicial
func List(initial []int, t int) map[[lib.GROUP]int][][lib.GROUP]int {
	var combins [][lib.GROUP]int
	arraymap := make(map[[lib.GROUP]int][][lib.GROUP]int)

	//First combinations of GROUP elements
	groups := GetGroups(initial, [lib.GROUP]int{}, 1)
	for _, g := range groups {
		var list [][lib.GROUP]int
		combins = GetGroups(initial, g, t)
		list = append(list, combins...)
		arraymap[g] = list
	}
	return arraymap
}

//treure els tipus automaticaments
func GetGroups(initial []int, g [lib.GROUP]int, t int) [][lib.GROUP]int {
	var combins [][lib.GROUP]int
	var remaining []int
	slice := [lib.GROUP]int{}

	if g != [lib.GROUP]int{} {
		remaining = lib.RemoveSlice(initial, g[:]) //Remaining array
	} else {
		remaining = initial
	}

	switch t {
	case 1:
		//Tipus {1,2,3}|{4,5,6}, {1,2,3}|{5,6,7}, {4,7,8}|{1,2,3}, ...
		for i := 0; i < len(remaining); i++ {
			slice[0] = remaining[i]
			for j := i + 1; j < len(remaining); j++ {
				slice[1] = remaining[j]
				for k := j + 1; k < len(remaining); k++ {
					slice[2] = remaining[k]
					combins = append(combins, slice)
				}
			}
		}

	case 2:
		//Tipus {1,2,3}|{1,4,5}, {1,2,3}|{4,2,5}, {1,2,3}|{4,5,3}
		for i := 0; i < len(g); i++ {
			slice[0] = g[i]
			for j := 0; j < len(remaining); j++ {
				slice[1] = remaining[j]
				for k := j + 1; k < len(remaining); k++ {
					slice[2] = remaining[k]
					combins = append(combins, slice)
				}
			}
		}

	case 3:
		//Tipus {1,2,3}|{1,2,5}, {1,2,3}|{1,4,3}, {1,2,3}|{4,2,3}
		for i := 0; i < len(g); i++ {
			slice[0] = g[i]
			for j := i + 1; j < len(g); j++ {
				slice[1] = g[j]
				for k := 0; k < len(remaining); k++ {
					slice[2] = remaining[k]
					combins = append(combins, slice)
				}
			}
		}

	}
	return combins
}
