package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/mifeis/Separable-Codes/combinations"
	"github.com/mifeis/Separable-Codes/lib"
)

//Test que retorna el numero de combinacions dependents

func TestDep(t *testing.T) {
	if lib.WORDS < 2*lib.GROUP {
		log.Fatal("num of words must be smaller than 2 * group elements")
	}
	initial := lib.Init(0, lib.WORDS)
	fmt.Println("\n...Getting dependence")
	deps := getDeps(initial)
	lib.LogDeps(deps)
}

func getDeps(initial []int) int {
	var depsdiv, deps int
	var array []map[int][]lib.Map
	var arraypairs [][]int

	//arraymaps[0]: 0 repetitions
	//arraymaps[1]: 1 repetitions
	//arraymaps[2]: 2 repetitions

	for i := 0; i < lib.REPS; i++ {
		arraymap := combinations.List(initial, i)
		array = append(array, lib.Sort(arraymap))
	}

	for _, arraymaps := range array {
		//Set a combination
		for _, am := range arraymaps {
			for _, m := range am {
				for _, ss := range m.Seconds {
					//					fmt.Println("first group key:", k1)
					//					fmt.Println("second group key:", k2)
					for _, s := range ss {
						pair := append(m.First, s...)
						//						fmt.Println(pair)
						arraypairs = append(arraypairs, pair)
					}
				}
			}
		}
	}
	//Falta comprobar
	for o, v := range arraypairs {
		for i := 0; i < len(arraypairs); i++ {
			if lib.Dependent(v, arraypairs[(o+i)%len(arraypairs)]) {
				//for l:=0;l<lib.GROUP;l++ en canvi de if if if
				// +o fer arraypairs com a map amb clau de num de reps
				if len(v) == 2*lib.GROUP {
					depsdiv++
				} else {
					//calcular mides i possibles dobles countings
					deps++
				}
			}
		}
	}

	return deps + (depsdiv / 2)
}
