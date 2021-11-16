package combinations

import (
	"fmt"
	"math/rand"
	"sync"
)

const WORDS = 8 //5 SECONDS, 11% CPU
const GROUP = 3

type Combin struct {
	Id    int
	Group [GROUP]int
}

func Init() []int {
	var c []int

	for i := 0; i < WORDS; i++ {
		c = append(c, i+1)
	}
	fmt.Println("Array:", c)
	return c
}

func List(c []int) int {
	combins := make(chan Combin)
	exit := make(chan bool)
	cases := make(chan int, 10000)
	wg := sync.WaitGroup{}
	var total int

	go GetGroups(true, c, combins, exit)

	for {
		select {
		case comb := <-combins:
			wg.Add(1)
			go func(c []int, g [GROUP]int, id int, cases chan int, wg *sync.WaitGroup) {
				slices := make(chan Combin)
				stop := make(chan bool)
				var total int

				remaining := RemoveSlice(c[:], g[:])
				fmt.Println("remaining array", id, remaining)

				go GetGroups(false, remaining, slices, stop)
				for {
					select {
					case /*s := */ <-slices:
						total++
						//						fmt.Println("slice from", id, "num", total, ":", s.Group)
					case <-stop:
						cases <- total
						wg.Done()
						return
					}
				}
			}(c, comb.Group, comb.Id, cases, &wg)

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

func GetGroups(first bool, c []int, combins chan Combin, exit chan bool) {
	var comb Combin
	slice := [GROUP]int{}
	for i := 0; i < len(c); i++ {
		slice[0] = c[i]
		for j := i + 1; j < len(c); j++ {
			slice[1] = c[j]
			for k := j + 1; k < len(c); k++ {
				slice[2] = c[k]
				if first {
					comb = Combin{
						Group: slice,
						Id:    rand.Intn(1000),
					}
					fmt.Println("...Sending slice", comb.Id, "...", comb.Group)
				} else {
					comb = Combin{
						Group: slice,
					}
				}
				combins <- comb
			}
		}
	}

	exit <- true
}

func RemoveSlice(original []int, g []int) []int {
	/*	k := 1

		for _, elem := range g {
			for j := k; j < len(original); j++ {
				if elem != original[j-1] {
					remaining = append(remaining, original[j-1])
					k = elem + 1
				} else {
					k = elem + 1
					break
				}
			}
		}

		for i := k; i <= WORDS; i++ {
			remaining = append(remaining, original[i-1])
		}
	*/

	/*for i, elem := range g {
		for _, v := range original {
			if elem == v {
				remaining = RemoveIndex(remaining, elem-i-1)
			}
		}
	}*/
	var remaining []int
	remaining = append(remaining, original[:]...)

	//cambiar: https://yourbasic.org/golang/delete-element-slice/-->trunccar

	for i, elem := range g {
		RemoveIndex(remaining, elem-i-1)
	}

	return remaining[:WORDS-GROUP]

}

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

/*func removeSlice(new []int, g []int) []int {
	remaining := make([]int, WORDS-GROUP)

	for i, elem := range g {
		for _, v := range new {
			if v != new[elem-i-1] {
				new = append(new, v)
			}
		}
	}

	for i := 0; i < (WORDS - GROUP); i++ {
		remaining[i] = new[i]
	}

	return remaining
}*/
