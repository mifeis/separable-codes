package combinations

/*
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
	slice := [GROUP]int{1, 2, 8}
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

func search(c [WORDS]int, g [GROUP]int, slices chan [GROUP]int, stop chan bool) {
	var colInOriginal bool
	var slice [GROUP]int
	var top bool

	//	{1,2,8} -> {3,4,5},{3,4,6},{3,4,7},{3,5,6},{3,5,7},{3,6,7},{4,5,6},{4,5,7},{4,6,7},{5,6,7}
	//	{4,6,7} -> {1,2,3},{1,2,5},{1,2,8},{1,3,5},{1,3,8},{1,5,8},{2,3,5},{2,3,8},{2,5,8},{3,5,8}
	//	{2,3,7} -> {1,4,5},{1,4,6},{1,4,8},{1,5,6},{1,5,8},{1,6,8},{4,5,6},{4,5,8},{4,6,8},{5,6,8}

	k := 0
	for i := 0; i < GROUP; {
		if top {
			for o := GROUP; o > GROUP-1; o-- {
				i--
				k = slice[i]
				slice[i] = 0
			}
			top = false
		}

		for j := k; j < WORDS; j++ {
			colInOriginal = false
			if slice[i] != 0 {
				k = slice[i]
				i++
				break
			} else {
				for _, v := range g {
					if c[j] == v {
						colInOriginal = true
						break
					}
				}
				if colInOriginal {
					if j == WORDS-1 {
						if i == 0 {
							stop <- true
							return
						} else {
							slice[i] = 0
							top = true
							break
						}
					}
				} else {
					switch i {
					case GROUP - 1:
						slice[i] = c[j]
						slices <- slice
						if j == WORDS-1 {
							top = true
						}
						slice[i] = 0
					case GROUP - 2:
						slice[i] = c[j]
						if j == WORDS-1 {
							slice[i] = 0
							top = true
							break
						}
					case GROUP - 3:
						slice[i] = c[j]
						if j == WORDS-1 {
							stop <- true
							return
						}
					}
				}
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
*/
