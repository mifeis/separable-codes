package combinations

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/mifeis/Separable-Codes/lib_aux"
)

//Funció que retorna els elements disjunts en grups de GROUP elements de l'array inicial
//Tipus {1,2,3}|{4,5,6}, {1,2,3}|{5,6,7}, {4,7,8}|{1,2,3}, ...

func List0(c []int) map[lib_aux.Combi][]lib_aux.Combi {
	var total int
	arraymap := make(map[lib_aux.Combi][]lib_aux.Combi)
	groups := GetGroups0(true, c)

	for _, g := range groups {
		var list []lib_aux.Combi
		remaining := lib_aux.RemoveSlice(c, g.Group[:])
		fmt.Println("Remaining array from", g.Id, ":", remaining)

		combins := GetGroups0(false, remaining)
		fmt.Println("Combinations:")
		for _, v := range combins {
			total++
			fmt.Println(total, "-", g.Group[:], "|", v.Group)
			list = append(list, v)
		}
		arraymap[g] = list
	}
	return arraymap
}

//Funció que es crida per tornar les combinacions de l'array passat per argument
func GetGroups0(first bool, c []int) []lib_aux.Combi {
	var comb lib_aux.Combi
	var combins []lib_aux.Combi
	slice := [lib_aux.GROUP]int{}
	for i := 0; i < len(c); i++ {
		slice[0] = c[i]
		for j := i + 1; j < len(c); j++ {
			slice[1] = c[j]
			for k := j + 1; k < len(c); k++ {
				slice[2] = c[k]
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

func GetFavsNofavs0(c []int) (int, int) {
	var favs, nofavs int
	arraymap := List0(c)

	defaultvalues := lib_aux.GetDefaultValues()

	//first groups
	for k, combs := range arraymap {
		//array of combinations of the first group
		for _, comb := range combs {
			log.Println(k.Group, "|", comb.Group)
			for i := 0; i < len(defaultvalues); i++ {
				k.Value = defaultvalues[i]
				for j := 0; j < len(defaultvalues); j++ {
					comb.Value = defaultvalues[j]
					if lib_aux.Separable(k.Value, comb.Value) {
						//if lib_aux.Separable(defaultvalues[i], defaultvalues[j]) {
						favs++
					} else {
						nofavs++
					}
				}
			}
			break
		}
		break
	}

	return favs, nofavs
}
