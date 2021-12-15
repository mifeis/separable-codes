package combinations

import (
	"fmt"
	"math/rand"

	"github.com/mifeis/Separable-Codes/lib_aux"
	"github.com/mifeis/Separable-Codes/lib_main"
)

//Tipus {1,2,3}|{1,4,5}, {1,2,3}|{4,2,5}, {1,2,3}|{4,5,3}
func List1(c []int) map[[lib_main.GROUP]int][][lib_main.GROUP]int {
	var total int
	arraymap := make(map[[lib_main.GROUP]int][][lib_main.GROUP]int)
	groups := GetGroups0(true, c)

	for _, g := range groups {
		var list [][lib_main.GROUP]int
		remaining := lib_aux.RemoveSlice(c, g.Group[:])
		fmt.Println("Remaining array from", g.Id, ":", remaining)

		combins := GetGroups1(false, g.Group[:], remaining)
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

func GetGroups1(first bool, g []int, remaining []int) []lib_aux.Combin {
	var comb lib_aux.Combin
	var combins []lib_aux.Combin
	slice := [lib_main.GROUP]int{}

	for l := 0; l < len(g); l++ {
		slice[0] = g[l]
		for i := 0; i < len(remaining); i++ {
			slice[1] = remaining[i]
			for j := i + 1; j < len(remaining); j++ {
				slice[2] = remaining[j]
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
