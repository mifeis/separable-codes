package combinations

import (
	"fmt"
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

//casos de 6 paraules i la diferencia multiplica
func GetFavsNofavs0(c []int) (int, int) {
	var favs, nofavs int
	arraymap := List0(c)

	//first groups
	for k, combs := range arraymap {

		//array of combinations of the first group
		for _, comb := range combs {
			for i := 0; i < lib_aux.GROUP*2; i++ {
				for j := 0; j < lib_aux.GROUP*2; j++ {
					//posicio dins l'array
					p := j / 2
					//cas dins l'array (0/1)
					switch j % 2 {
					case 0:
						comb.Value[lib_aux.GROUP-1-p]++
					case 1:
						comb.Value[lib_aux.GROUP-2-p]--
						comb.Value[lib_aux.GROUP-1-p]++
					}

					if lib_aux.Separable(k.Value, comb.Value) {
						favs++
					} else {
						nofavs++
					}
				}
				//fer el same q amb comb
				k.Value[lib_aux.GROUP-1-i]++
				comb.Value = [lib_aux.GROUP]int{0, 0, 0}
			}
		}

	}

	return len(arraymap), 0
}
