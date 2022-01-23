package test

import (
	"fmt"
	"log"

	"github.com/mifeis/Separable-Codes/lib"
)

func Disjunts() {
	if lib.WORDS < 2*lib.GROUP {
		log.Fatal("num of words is small")
	}

	var tipusdisjunts, tipusnodisjunts int
	for i := lib.GROUP; i > 0; i-- {
		for j := i; j > 0; j-- {
			tipusdisjunts++
		}
	}

	for i := lib.GROUP; i > 0; i-- {
		for j := i; j > 0; j-- {
			for k := j; k > 0; k-- {
				tipusnodisjunts++
			}
		}
		tipusnodisjunts--
	}

	fmt.Println("Total de combinacions diferents per", lib.GROUP, "elements de grup DISJUNTES:", tipusdisjunts)
	fmt.Println("Total de tipus de combinacions diferents per", lib.GROUP, "elements de grup NO DISJUNTES:", tipusnodisjunts)

}
