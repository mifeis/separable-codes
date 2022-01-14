package combinations

import (
	"log"

	"github.com/mifeis/Separable-Codes/lib"
	"gonum.org/v1/gonum/stat/combin"
)

//Funció que retorna els elements disjunts en grups de GROUP elements de l'array inicial
func List(initial []int, t int) []lib.Map {
	var combins map[int][]int
	arraymap := []lib.Map{}
	//First combinations of GROUP elements
	groups := GetGroups(initial, []int{}, 1)
	for _, g := range groups {
		var list [][]int
		combins = GetGroups(lib.RemoveSlice(initial, g[:]), g, t)
		for i := range combins {
			list = append(list, combins[i])
		}
		//		list = append(list, combins...)
		m := lib.Map{
			First:   g,
			Seconds: list,
		}
		arraymap = append(arraymap, m)
	}
	return arraymap
}

//comptar els casos incomplerts i els complerts (done)
func GetGroups(remaining []int, g []int, t int) map[int][]int {
	combins := make(map[int][]int, 1000)
	in := combin.Combinations(lib.GROUP, t-1)
	for p := 0; p < len(in); p++ {
		var slice [lib.GROUP]int
		var init int
		//valors no disjunts
		for r := 0; r < t-1; r++ {
			slice[r] = g[in[p][r]]
			init++
		}
		//valors disjunts
		//Casos complerts:
		indexes := combin.Combinations(len(remaining), lib.GROUP-(t-1))
		log.Println(len(indexes))
		for o, index := range indexes {
			var slice [lib.GROUP]int
			for i, v := range index {
				slice[init+i] = remaining[v]
			}
			combins[o] = slice[:]
			log.Println("----------POST---------")
			log.Println(combins)
		}
		//*Comentar marcel
		//treure if i tenir en compte els casos de 2 i 3 grups p 1 i 3 (al revés)
		/*		if g[0] != 0 {
					//comentar
					//Casos incomplerts:
					for r := init; r < lib.GROUP; r++ {
						slice[r] = 0
					}
					for l := 0; l < lib.GROUP-1; l++ {
						indexes = combin.Combinations(len(remaining), l+1)
						for _, index := range indexes {
							for i, v := range index {
								slice[init+i] = remaining[v]
							}
							combins = append(combins, slice)
						}
					}
				}
		*/
	}
	return combins
}
