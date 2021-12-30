package combinations

import (
	"fmt"

	"github.com/mifeis/Separable-Codes/lib_aux"
)

//probar a retornar index i juntar groups0,1,2
//Funció que retorna els elements disjunts en grups de GROUP elements de l'array inicial
func List(initial []int, t int) map[lib_aux.Combi]([]lib_aux.Combi) {
	var combins []lib_aux.Combi
	arraymap := make(map[lib_aux.Combi][]lib_aux.Combi)

	//First combinations of GROUP elements
	groups := GetGroups(initial, nil, 1)
	for _, g := range groups {
		var list []lib_aux.Combi
		combins = GetGroups(initial, g.Group[:], t)
		fmt.Println("Combinations:")
		for _, v := range combins {
			fmt.Println("\t\t\t\t\t", g.Group[:], "|", v.Group)
			list = append(list, v)
		}
		arraymap[g] = list
	}
	return arraymap
}

func GetGroups(initial []int, firstgroup []int, t int) []lib_aux.Combi {
	var comb lib_aux.Combi
	var combins []lib_aux.Combi
	var remaining []int
	slice := [lib_aux.GROUP]int{}

	if firstgroup != nil {
		//Remaining array
		remaining = lib_aux.RemoveSlice(initial, firstgroup)
		fmt.Println("\nRemaining array from", firstgroup, ":", remaining)
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
					comb = lib_aux.Combi{
						Group: slice,
					}
					combins = append(combins, comb)
				}
			}
		}

	//Tipus {1,2,3}|{1,4,5}, {1,2,3}|{4,2,5}, {1,2,3}|{4,5,3}
	case 2:
		for l := 0; l < len(firstgroup); l++ {
			slice[0] = firstgroup[l]
			for i := 0; i < len(remaining); i++ {
				slice[1] = remaining[i]
				for j := i + 1; j < len(remaining); j++ {
					slice[2] = remaining[j]
					comb = lib_aux.Combi{
						Group: slice,
					}
					combins = append(combins, comb)
				}
			}
		}

	//Tipus {1,2,3}|{1,2,5}, {1,2,3}|{1,4,3}, {1,2,3}|{4,2,3}
	case 3:
		for l := 0; l < len(firstgroup); l++ {
			slice[0] = firstgroup[l]
			for i := l + 1; i < len(firstgroup); i++ {
				slice[1] = firstgroup[i]
				for j := 0; j < len(remaining); j++ {
					slice[2] = remaining[j]
					comb = lib_aux.Combi{
						Group: slice,
					}
					combins = append(combins, comb)
				}
			}
		}

	}
	return combins
}
