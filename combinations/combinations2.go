package combinations

import (
	"fmt"
	"math/rand"

	"github.com/mifeis/Separable-Codes/lib"
)

func ListSeq2(c []int) int {
	var total int

	groups := GetGroupsSeq(true, c)

	for _, g := range groups {
		new := c

		remaining := lib.RemoveSlice(new, g.Group[:])
		fmt.Println("remaining array ", g.Id, ": ", remaining)

		combins := GetGroupsSeq(false, remaining)
		for _, v := range combins {
			total++
			fmt.Println("slice from", g.Id, "num", total, ":", v.Group)
		}
	}
	return total
}

func GetGroupsSeq2(first bool, c []int) []Combin {
	var comb Combin
	var combins []Combin
	slice := [GROUP]int{}
	for i := 0; i < len(c); i++ {
		slice[0] = c[i]
		for j := i + 1; j < len(c); j++ {
			slice[1] = c[j]
			for k := j + 1; k < len(c); k++ {
				slice[2] = c[k]
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
