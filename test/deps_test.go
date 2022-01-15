package main

import (
	"log"
	"testing"

	"github.com/mifeis/Separable-Codes/lib"
)

//Test que retorna el numero de combinacions dependents

func TestDependence(t *testing.T) {
	if lib.WORDS < 2*lib.GROUP {
		log.Fatal("num of words must be smaller than 2 * group elements")
	}
	//falta desenvolupar
}
