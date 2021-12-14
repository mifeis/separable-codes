package combinations

import (
	"fmt"
	"math/rand"

	"github.com/mifeis/Separable-Codes/lib"
)

//Tipus {1,2,3}|{1,2,5}, {1,2,3}|{1,4,3}, {1,2,3}|{4,2,3}
func List2(c []int) map[[lib.GROUP]int][][lib.GROUP]int {
	var total int
	arraymap := make(map[[lib.GROUP]int][][lib.GROUP]int)
	groups := GetGroups0(true, c)

	for _, g := range groups {
		var list [][lib.GROUP]int
		remaining := lib.RemoveSlice(c, g.Group[:])
		fmt.Println("Remaining array from", g.Id, ":", remaining)

		combins := GetGroups2(false, g.Group[:], remaining)
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

//Tipus {1,2,3}|{1,2,5}, {1,2,3}|{1,4,3}, {1,2,3}|{4,2,3}
func GetGroups2(first bool, g []int, remaining []int) []lib.Combin {
	var comb lib.Combin
	var combins []lib.Combin
	slice := [lib.GROUP]int{}

	for l := 0; l < len(g); l++ {
		slice[0] = g[l]
		for i := l + 1; i < len(g); i++ {
			slice[1] = g[i]
			for j := 0; j < len(remaining); j++ {
				slice[2] = remaining[j]
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
