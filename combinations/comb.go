package combinations

import (
	"fmt"
	"sort"
	"sync"
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
	//	id := rand.Intn(100)
	slice := []int{g[0], g[1], g[2]}
	sort.Ints(slice)

	/*	for i := slice[0]; i < WORDS; i++ {
			//		slice[0] = c[i]
			for j := i + 1; j < WORDS; j++ {
				//			slice[1] = c[j]
				for k := j + 1; k < WORDS; k++ {
					//				slice[2] = c[k]
					cases <- 1
				}
			}
		}
	*/
	wg.Done()
}
