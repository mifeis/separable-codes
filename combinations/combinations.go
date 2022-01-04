package combinations

import (
	"github.com/mifeis/Separable-Codes/lib"
)

//Funció que retorna els elements disjunts en grups de GROUP elements de l'array inicial
func List(initial []int, t int) map[lib.Combi]([]lib.Combi) {
	var combins []lib.Combi
	arraymap := make(map[lib.Combi][]lib.Combi)

	//First combinations of GROUP elements
	groups := GetGroups(initial, nil, 1)
	for _, g := range groups {
		var list []lib.Combi
		combins = GetGroups(initial, g.Rows[:], t)
		list = append(list, combins...)
		arraymap[g] = list
	}
	return arraymap
}

//treure els tipus automaticament
func GetGroups(initial []int, g []int, t int) []lib.Combi {
	var c lib.Combi
	var combins []lib.Combi
	var remaining []int
	slice := [lib.GROUP]int{}
	if g != nil {
		remaining = lib.RemoveSlice(initial, g) //Remaining array
	} else {
		remaining = initial
	}

	if t == 1 {
		//Casos disjunts
		//Tipus {1,2,3}|{4,5,6}, {1,2,3}|{5,6,7}, {4,7,8}|{1,2,3}, ...
		for i := 0; i < len(remaining); i++ {
			for j := i + 1; j < len(remaining); j++ {
				for k := j + 1; k < len(remaining); k++ {
					slice[0] = remaining[i]
					slice[1] = remaining[j]
					slice[2] = remaining[k]
					c = lib.Combi{
						Rows: slice,
					}
					combins = append(combins, c)
				}
			}
		}

	} else {
		//Casos NO disjunts
		//[1,2,3]-> 3
		//[1,2,3,4]-> 6
		//[1,2,3,4,5]-> 10
		for p := 0; p < lib.GROUP; p++ {
			init := 0
			for r := init; r < t-1; r++ {
				slice[r] = g[(r+p)%lib.GROUP]
				init++
			}
			switch t {
			case 2:
				//Tipus {1,2,3}|{1,4,5}, {1,2,3}|{4,2,5}, {1,2,3}|{4,5,3}
				for i := 0; i < len(remaining); i++ {
					for j := i + 1; j < len(remaining); j++ {
						slice[init] = remaining[i]
						slice[init+1] = remaining[j]
						c = lib.Combi{
							Rows: slice,
						}
						combins = append(combins, c)
					}
				}
			case 3:
				//Tipus {1,2,3}|{1,2,5}, {1,2,3}|{1,4,3}, {1,2,3}|{4,2,3}
				for j := 0; j < len(remaining); j++ {
					slice[init] = remaining[j]
					c = lib.Combi{
						Rows: slice,
					}
					combins = append(combins, c)
				}
			}
		}
	}
	return combins
}
