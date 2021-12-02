package combinations

import (
	"fmt"
	"math/rand"

	"github.com/mifeis/Separable-Codes/lib"
)

const (
	WORDS = 8
	GROUP = 3
)

func List0(c []int) int {
	var total int

	groups := GetGroups0(true, c)

	for _, g := range groups {
		remaining := lib.RemoveSlice(c, g.Group[:])
		fmt.Println("remaining array ", g.Id, ": ", remaining)

		combins := GetGroups0(false, remaining)
		for _, v := range combins {
			total++
			fmt.Println("slice from", g.Id, "num", total, ":", v.Group)
		}
	}
	return total
}

func GetGroups0(first bool, c []int) []Combin {
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
