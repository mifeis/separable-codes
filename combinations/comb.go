package combinations

import (
	"fmt"
	"sync"
)

const WORDS = 8
const GROUP = 3

var combs [][GROUP]int

func Init() [WORDS]int {
	var c [WORDS]int

	for i := 0; i < WORDS; i++ {
		c[i] = i + 1
	}
	fmt.Println("Array:", c)
	return c
}

func List(c [WORDS]int) int {
	groups := make(chan [GROUP]int)
	exit := make(chan bool)
	cases := make(chan int, 100)
	var total int
	wg := sync.WaitGroup{}
	//	var groups chan [GROUP]int

	go GetGroups(c, groups, exit)

	for {
		select {
		case g := <-groups:
			wg.Add(1)
			go getComb(c, g, cases, &wg)
		case <-exit:
			wg.Wait()
			for i := 0; i < len(cases); {
				num := <-cases
				total += num
			}
			return total
		}
	}
}

func GetGroups(c [WORDS]int, groups chan [GROUP]int, exit chan bool) {
	var slice [GROUP]int
	for i := 0; i < WORDS; i++ {
		slice[0] = c[i]
		for j := i + 1; j < WORDS; j++ {
			slice[1] = c[j]
			for k := j + 1; k < WORDS; k++ {
				slice[2] = c[k]
				fmt.Println("...Sending slice...", slice)
				groups <- slice
			}
		}
	}
	exit <- true
}

func getComb(c [WORDS]int, g [GROUP]int, cases chan int, wg *sync.WaitGroup) {
	var colInOriginal bool
	var passed bool
	slice := [GROUP]int{g[0], g[1], g[2]}

	//	{1,2,3} -> {4,5,6},{4,5,7},{4,5,8},{4,6,7},{4,6,8},{4,7,8},{5,6,7},{5,6,8},{5,7,8},{6,7,8}
	//	{4,6,7} -> {1,2,3},{1,2,5},{1,2,8},{1,3,5},{1,3,8},{1,5,8},{2,3,5},{2,3,8},{2,5,8},{3,5,8}
	//	{2,3,7} -> {1,4,5},{1,4,6},{1,4,8},{1,5,6},{1,5,8},{1,6,8},{4,5,6},{4,5,8},{4,6,8},{5,6,8}

	k := 0

	for i := 0; i < GROUP; i++ {
		for j := k; j < WORDS; j++ {
			colInOriginal = false
			if slice[i] != 0 {
				k++
				break //ó j=WORDS
			} else {
				//comprobamos que no este ya la columna en el primer grupo. Por ejemplo: {2,3,7}
				for _, v := range g {
					if c[j] == v {
						colInOriginal = true
					}
				}
				if !colInOriginal {
					if i == GROUP-1 {
						//comprobamos q no este ya contado el caso
						if !checkIn(slice, combs) {
							if k == WORDS-1 {
								passed = true
							} else {
								passed = false
							}
							slice[i] = c[j]
							combs = append(combs, slice)
							cases <- 1
							//una vex encontrado un caso buscar los dmas
						} else {
							passed = true
						}
					} else {
						slice[i] = c[j]
					}
				}
			}
		}
		if passed {
			passed = false
			k = 0
			for o := GROUP; o <= i; o-- {
				i--
				k--
			}
		}
	}

	wg.Done()
}

func checkIn(slice [GROUP]int, combs [][GROUP]int) bool {
	var groupDone int

	for _, w := range combs {
		groupDone = 0
		for l, x := range w {
			if x == slice[l] {
				groupDone++
			}
		}
		if groupDone == GROUP {
			return true
		}
	}
	return false
}
