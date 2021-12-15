package combinations

import (
	"fmt"
	"math/rand"

	"github.com/mifeis/Separable-Codes/lib_aux"
	"github.com/mifeis/Separable-Codes/lib_main"
)

//Funció que retorna els elements disjunts en grups de GROUP elements de l'array inicial
//Tipus {1,2,3}|{4,5,6}, {1,2,3}|{5,6,7}, {4,7,8}|{1,2,3}, ...

func List0(c []int) map[[lib_main.GROUP]int][][lib_main.GROUP]int {
	var total int
	arraymap := make(map[[lib_main.GROUP]int][][lib_main.GROUP]int)
	groups := GetGroups0(true, c)

	for _, g := range groups {
		var list [][lib_main.GROUP]int
		remaining := lib_aux.RemoveSlice(c, g.Group[:])
		fmt.Println("Remaining array from", g.Id, ":", remaining)

		combins := GetGroups0(false, remaining)
		fmt.Println("Combinations:")
		for _, v := range combins {
			total++
			fmt.Println(total, "-", g.Group[:], "|", v.Group)
			list = append(list, v.Group)
		}
		arraymap[g.Group] = list
	}
	return arraymap
}

//Funció que es crida per tornar les combinacions de l'array passat per argument
func GetGroups0(first bool, c []int) []lib_aux.Combin {
	var comb lib_aux.Combin
	var combins []lib_aux.Combin
	slice := [lib_main.GROUP]int{}
	for i := 0; i < len(c); i++ {
		slice[0] = c[i]
		for j := i + 1; j < len(c); j++ {
			slice[1] = c[j]
			for k := j + 1; k < len(c); k++ {
				slice[2] = c[k]
				if first {
					comb = lib_aux.Combin{
						Group: slice,
						Id:    rand.Intn(1000),
					}
					fmt.Println("...Sending slice", comb.Id, "...", comb.Group)
				} else {
					comb = lib_aux.Combin{
						Group: slice,
					}
				}
				combins = append(combins, comb)
			}
		}
	}
	return combins
}
