package combinations

import (
	"fmt"
	"math/rand"

	"github.com/mifeis/Separable-Codes/lib"
)

//Funció que retorna els elements disjunts en grups de GROUP elements de l'array inicial
//Tipus {1,2,3}|{4,5,6}, {1,2,3}|{5,6,7}, {4,7,8}|{1,2,3}, ...

func List0(c []int) [][]int {
	var total int
	var list [][]int
	groups := GetGroups0(true, c)

	for _, g := range groups {
		remaining := lib.RemoveSlice(c, g.Group[:])
		fmt.Println("Remaining array from", g.Id, ":", remaining)

		combins := GetGroups0(false, remaining)
		fmt.Println("Combinations:")
		for _, v := range combins {
			total++
			//			fmt.Println("slice from", g.Group[:], "with id:", g.Id, "num", total, ":", v.Group)
			fmt.Println(total, "-", g.Group[:], "|", v.Group)
			groups := append(g.Group[:], v.Group[:]...)
			list = append(list, groups)
		}
	}
	return list
}

//Funció que es crida per tornar les combinacions de l'array passat per argument
func GetGroups0(first bool, c []int) []lib.Combin {
	var comb lib.Combin
	var combins []lib.Combin
	slice := [lib.GROUP]int{}
	for i := 0; i < len(c); i++ {
		slice[0] = c[i]
		for j := i + 1; j < len(c); j++ {
			slice[1] = c[j]
			for k := j + 1; k < len(c); k++ {
				slice[2] = c[k]
				if first {
					comb = lib.Combin{
						Group: slice,
						Id:    rand.Intn(1000),
					}
					fmt.Println("...Sending slice", comb.Id, "...", comb.Group)
				} else {
					comb = lib.Combin{
						Group: slice,
					}
				}
				combins = append(combins, comb)
			}
		}
	}
	return combins
}
