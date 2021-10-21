package combinations

import (
	"fmt"
)

const WORDS = 8
const GROUP = 3

func Init() [WORDS]int {
	var c [WORDS]int

	for i := 0; i < WORDS; i++ {
		c[i] = i
	}
	fmt.Println("Array:", c)
	return c
}

func List(c [WORDS]int) int {
	var exit chan bool
	var totalCases int
	groups := make(chan [GROUP]int, 1000)
	//	var groups chan [GROUP]int

	go getGroup(c, groups, exit)

	for {
		select {
		case <-exit:
			return totalCases
		case g := <-groups:
			fmt.Println(g)
			totalCases++
		}
	}
	/*		for i := 0; i < GROUP; i++ {
				for j := i; j < WORDS; j++ {
					slice2[j] = j + 1

				}
			}
	*/
	//	return totalCases
}

func getGroup(c [WORDS]int, groups chan [GROUP]int, exit chan bool) {
	var slice [GROUP]int
	for i := 0; i < WORDS; i++ {
		//i := 0
		slice[0] = c[i]
		for j := i + 1; j < WORDS-1; j++ {
			//j := i + 1
			slice[1] = c[j]
			for k := j + 1; k < WORDS-2; k++ {
				//k := j + 1
				slice[2] = c[k]
				groups <- slice
			}
		}
	}
	exit <- true
}
