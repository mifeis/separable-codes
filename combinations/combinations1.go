package combinations

import (
	"fmt"
	"math/rand"

	"github.com/mifeis/Separable-Codes/lib"
)

//Tipus {1,2,3}|{1,4,5}, {1,2,3}|{4,2,5}, {1,2,3}|{4,5,3}

func List1(c []int) int {
	var total int

	groups := GetGroups0(true, c)

	for _, g := range groups {
		remaining := lib.RemoveSlice(c, g.Group[:])
		fmt.Println("remaining array ", g.Id, ": ", remaining)

		combins := GetGroups1(false, g.Group[:], remaining)
		for _, v := range combins {
			total++
			fmt.Println("slice from", g.Id, "num", total, ":", v.Group)
		}
	}
	return total
}

func GetGroups1(first bool, g []int, remaining []int) []Combin {
	var comb Combin
	var combins []Combin
	slice := [GROUP]int{}

	for l := 0; l < len(g); l++ {
		slice[0] = g[l]
		for i := 0; i < len(remaining); i++ {
			slice[1] = remaining[i]
			for j := i + 1; j < len(remaining); j++ {
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
