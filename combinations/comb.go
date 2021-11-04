package combinations

import (
	"fmt"
	"sync"
)

const WORDS = 8
const GROUP = 3

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
	wg := sync.WaitGroup{}
	var total int

	go GetGroups(c, groups, exit)

	for {
		select {
		case g := <-groups:
			wg.Add(1)
			go getCombs(c, g, cases, &wg)
		case <-exit:
			fmt.Println("waiting")
			wg.Wait()
			for i := 0; i < len(cases); {
				num := <-cases
				total += num
			}
			fmt.Println("done")
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
				break
			}
			break
		}
		break
	}
	exit <- true
}

func checkIn(slice [GROUP]int, combs [][GROUP]int) bool {
	var coinc int

	for _, w := range combs {
		coinc = 0
		for l, x := range w {
			if x == slice[l] {
				coinc++
			}
		}
		if coinc == GROUP {
			return true
		}
	}
	return false
}

func search(c [WORDS]int, g [GROUP]int, slices chan [GROUP]int, stop chan bool) {
	var combs [][GROUP]int
	var colInOriginal bool
	var top bool
	limit := GROUP - 1
	slice := [GROUP]int{}

	//	{1,2,3} -> {4,5,6},{4,5,7},{4,5,8},{4,6,7},{4,6,8},{4,7,8},{5,6,7},{5,6,8},{5,7,8},{6,7,8}
	//	{4,6,7} -> {1,2,3},{1,2,5},{1,2,8},{1,3,5},{1,3,8},{1,5,8},{2,3,5},{2,3,8},{2,5,8},{3,5,8}
	//	{2,3,7} -> {1,4,5},{1,4,6},{1,4,8},{1,5,6},{1,5,8},{1,6,8},{4,5,6},{4,5,8},{4,6,8},{5,6,8}

	//	for {
	k := 0
	for i := 0; i < GROUP; {
		fmt.Println("slice when done", slice)
		if top {
			fmt.Println("limit", limit)
			//limit = 2
			//i=2
			for o := GROUP; o > limit; o-- {
				i--
				k = slice[i] + 1
				slice[i] = 0
			}
			top = false
			fmt.Println("done", top)
		}
		fmt.Println("i: ", i)
		for j := k; j < WORDS; j++ {
			fmt.Println("k: ", k)
			colInOriginal = false
			if slice[i] != 0 {
				fmt.Println("slice[i]!=0")
				i++
				break
			} else {
				//comprobamos que no este ya la columna en el mismo grupo. Por ejemplo: {2,3,7}
				for _, v := range g {
					if c[j] == v {
						colInOriginal = true
						break
					}
				}
				fmt.Println("colInOriginal", colInOriginal)
				if !colInOriginal {
					switch i {
					case GROUP - 1:
						fmt.Println("case i==GROUP-1")
						if c[j] > slice[i-1] {
							fmt.Println("c[j] > slice[i-1]")
							slice[i] = c[j]
							if !checkIn(slice, combs) {
								fmt.Println("!checkIn")
								slice[i] = c[j]
								combs = append(combs, slice)
								slices <- slice
							} else {
								fmt.Println("checkIn. slice[i]=0")
								slice[i] = 0
							}
							if j == WORDS-1 {
								top = true
								limit = i
								fmt.Println("done", top)
							}
							slice[i] = 0
						}
					//comprobamos q no este ya contado el caso
					case GROUP - 2:
						fmt.Println("case 2==GROUP-2")
						if c[j] > slice[i-1] {
							fmt.Println("c[j] > slice[i-1]")
							slice[i] = c[j]
							if j == WORDS-1 {
								limit = i
								fmt.Println("limit", limit)
							}
						}
						//					case GROUP - 3:
					default:
						fmt.Println("default case")
						slice[i] = c[j]
						if j == WORDS-1 {
							limit = i
							fmt.Println("limit", limit)
						}
					}
				}
			}
			if i == 0 && j == WORDS-1 {
				stop <- true
			}
		}
	}
}

func getCombs(c [WORDS]int, g [GROUP]int, cases chan int, wg *sync.WaitGroup) {
	slices := make(chan [GROUP]int)
	stop := make(chan bool)
	var total int

	go search(c, g, slices, stop)
	for {
		select {
		case s := <-slices:
			total++
			fmt.Println("slice", total, s[:])
		case <-stop:
			cases <- total
			wg.Done()
			return
		}
	}
}
