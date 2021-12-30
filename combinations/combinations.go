package combinations

import "github.com/mifeis/Separable-Codes/lib"

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

//treure els tipus automaticaments
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

	switch t {
	//Tipus {1,2,3}|{4,5,6}, {1,2,3}|{5,6,7}, {4,7,8}|{1,2,3}, ...
	case 1:
		for i := 0; i < len(remaining); i++ {
			slice[0] = remaining[i]
			for j := i + 1; j < len(remaining); j++ {
				slice[1] = remaining[j]
				for k := j + 1; k < len(remaining); k++ {
					slice[2] = remaining[k]
					c = lib.Combi{
						Rows: slice,
					}
					combins = append(combins, c)
				}
			}
		}

	//Tipus {1,2,3}|{1,4,5}, {1,2,3}|{4,2,5}, {1,2,3}|{4,5,3}
	case 2:
		for l := 0; l < len(g); l++ {
			slice[0] = g[l]
			for i := 0; i < len(remaining); i++ {
				slice[1] = remaining[i]
				for j := i + 1; j < len(remaining); j++ {
					slice[2] = remaining[j]
					c = lib.Combi{
						Rows: slice,
					}
					combins = append(combins, c)
				}
			}
		}

	//Tipus {1,2,3}|{1,2,5}, {1,2,3}|{1,4,3}, {1,2,3}|{4,2,3}
	case 3:
		for l := 0; l < len(g); l++ {
			slice[0] = g[l]
			for i := l + 1; i < len(g); i++ {
				slice[1] = g[i]
				for j := 0; j < len(remaining); j++ {
					slice[2] = remaining[j]
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
