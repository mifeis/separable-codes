package combinations

import (
	"fmt"
	"math/rand"

	"github.com/mifeis/Separable-Codes/lib"
)

//Tipus {1,2,3}|{1,2,5}, {1,2,3}|{1,4,3}, {1,2,3}|{4,2,3}
func List2(c []int) int {
	var total int

	groups := GetGroups0(true, c)

	for _, g := range groups {
		remaining := lib.RemoveSlice(c, g.Group[:])
		fmt.Println("remaining array ", g.Id, ": ", remaining)

		combins := GetGroups2(false, g.Group[:], remaining)
		for _, v := range combins {
			total++
			fmt.Println("slice from", g.Id, "num", total, ":", v.Group)
		}
	}
	return total
}

//Tipus {1,2,3}|{1,2,5}, {1,2,3}|{1,4,3}, {1,2,3}|{4,2,3}
func GetGroups2(first bool, g []int, remaining []int) []Combin {
	var comb Combin
	var combins []Combin
	slice := [lib.GROUP]int{}

	for l := 0; l < len(g); l++ {
		slice[0] = g[l]
		for i := l + 1; i < len(g); i++ {
			slice[1] = g[i]
			for j := 0; j < len(remaining); j++ {
				slice[2] = remaining[j]
				if first {
					comb = Combin{
						Group: slice,
						Id:    rand.Intn(1000),
					}
					fmt.Println("...Sending slice ", comb.Id, "...", comb.Group)
				} else {
					comb = Combin{
						Group: slice,
					}
				}
				combins = append(combins, comb)
			}
		}
	}
	return combins
}
