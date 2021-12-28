package combinations

import (
	"fmt"
	"math/rand"

	"github.com/mifeis/Separable-Codes/lib_aux"
)

//probar a retornar index i juntar groups0,1,2
//Funció que retorna els elements disjunts en grups de GROUP elements de l'array inicial
func List(c []int, t int) map[lib_aux.Combi]([]lib_aux.Combi) {
	var index int
	var combins []lib_aux.Combi
	arraymap := make(map[lib_aux.Combi][]lib_aux.Combi)

	//First combinations of GROUP elements
	groups := GetGroups0(true, c)
	for _, g := range groups {
		var list []lib_aux.Combi
		//Remaining array
		remaining := lib_aux.RemoveSlice(c, g.Group[:])
		fmt.Println("Remaining array from", g.Id, ":", remaining)
		switch t {
		case 0:
			combins = GetGroups0(false, remaining)
		case 1:
			combins = GetGroups1(false, g.Group[:], remaining)
		case 2:
			combins = GetGroups2(false, g.Group[:], remaining)
		}
		fmt.Println("Combinations:")
		for _, v := range combins {
			index++
			fmt.Println(index, "-", g.Group[:], "|", v.Group)
			list = append(list, v)
		}
		arraymap[g] = list
	}
	return arraymap
}

//Funció que es crida per tornar les combinacions de l'array passat per argument
//Tipus {1,2,3}|{4,5,6}, {1,2,3}|{5,6,7}, {4,7,8}|{1,2,3}, ...
func GetGroups0(first bool, remaining []int) []lib_aux.Combi {
	var comb lib_aux.Combi
	var combins []lib_aux.Combi
	slice := [lib_aux.GROUP]int{}
	for i := 0; i < len(remaining); i++ {
		slice[0] = remaining[i]
		for j := i + 1; j < len(remaining); j++ {
			slice[1] = remaining[j]
			for k := j + 1; k < len(remaining); k++ {
				slice[2] = remaining[k]
				if first {
					comb = lib_aux.Combi{
						Group: slice,
						Id:    rand.Intn(1000),
					}
					fmt.Println("...Sending slice", comb.Id, "... ->", comb.Group)
				} else {
					comb = lib_aux.Combi{
						Group: slice,
					}
				}
				combins = append(combins, comb)
			}
		}
	}
	return combins
}

//Tipus {1,2,3}|{1,4,5}, {1,2,3}|{4,2,5}, {1,2,3}|{4,5,3}
func GetGroups1(first bool, g []int, remaining []int) []lib_aux.Combi {
	var comb lib_aux.Combi
	var combins []lib_aux.Combi
	slice := [lib_aux.GROUP]int{}

	for l := 0; l < len(g); l++ {
		slice[0] = g[l]
		for i := 0; i < len(remaining); i++ {
			slice[1] = remaining[i]
			for j := i + 1; j < len(remaining); j++ {
				slice[2] = remaining[j]
				if first {
					comb = lib_aux.Combi{
						Group: slice,
						Id:    rand.Intn(1000),
					}
					fmt.Println("...Sending slice", comb.Id, "...", comb.Group)
				} else {
					comb = lib_aux.Combi{
						Group: slice,
					}
				}
				combins = append(combins, comb)
			}
		}
	}
	return combins
}

//Tipus {1,2,3}|{1,2,5}, {1,2,3}|{1,4,3}, {1,2,3}|{4,2,3}
func GetGroups2(first bool, g []int, remaining []int) []lib_aux.Combi {
	var comb lib_aux.Combi
	var combins []lib_aux.Combi
	slice := [lib_aux.GROUP]int{}

	for l := 0; l < len(g); l++ {
		slice[0] = g[l]
		for i := l + 1; i < len(g); i++ {
			slice[1] = g[i]
			for j := 0; j < len(remaining); j++ {
				slice[2] = remaining[j]
				if first {
					comb = lib_aux.Combi{
						Group: slice,
						Id:    rand.Intn(1000),
					}
					fmt.Println("...Sending slice", comb.Id, "...", comb.Group)
				} else {
					comb = lib_aux.Combi{
						Group: slice,
					}
				}
				combins = append(combins, comb)
			}
		}
	}
	return combins
}
