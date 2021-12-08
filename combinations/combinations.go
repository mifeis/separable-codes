package combinations

import (
	"fmt"
	"math/rand"

	"github.com/mifeis/Separable-Codes/lib"
)

//Estructura que conté el grup de GROUP elements i un random id
//per saber de quina combinació es tracta i fer mes entendible l'arxiu resultant

type Combin struct {
	Id    int
	Group [lib.GROUP]int
}

//funció que inicialitza i retorna l'array a combinar: {1,2,3,4,5,6,7,8,...}
func Init() []int {
	var c []int

	for i := 0; i < lib.WORDS; i++ {
		c = append(c, i+1)
	}
	fmt.Println("Array:", c)
	return c
}

//Funció que retorna els elements disjunts en grups de GROUP elements de l'array inicial
//Tipus {1,2,3}|{4,5,6}, {1,2,3}|{5,6,7}, {4,7,8}|{1,2,3}, ...

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

//Funció que es crida per tornar les combinacions de l'array passat per argument
func GetGroups0(first bool, c []int) []Combin {
	var comb Combin
	var combins []Combin
	slice := [lib.GROUP]int{}
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
