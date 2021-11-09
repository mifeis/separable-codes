package combinations

import (
	"fmt"
	"math/rand"
	"sync"
)

const WORDS = 8
const GROUP = 3

func Init() []int {
	var c []int

	for i := 0; i < WORDS; i++ {
		c = append(c, i+1)
	}
	fmt.Println("Array:", c)
	return c
}

func List(c []int) int {
	groups := make(chan [GROUP + 1]int)
	exit := make(chan bool)
	cases := make(chan int, 100)
	wg := sync.WaitGroup{}
	var total int

	go GetGroups1(c, groups, exit)

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

func GetGroups1(c []int, groups chan [GROUP + 1]int, exit chan bool) {
	slice := [GROUP + 1]int{}
	for i := 0; i < len(c); i++ {
		slice[0] = c[i]
		for j := i + 1; j < len(c); j++ {
			slice[1] = c[j]
			for k := j + 1; k < len(c); k++ {
				slice[2] = c[k]
				slice[3] = rand.Intn(100)
				fmt.Println("...Sending slice ", slice[3], "...", slice)
				groups <- slice
			}
		}
	}

	exit <- true
}

func GetGroups2(remaining []int, slices chan [GROUP]int, stop chan bool) {
	slice := [GROUP]int{}
	for i := 0; i < len(remaining); i++ {
		slice[0] = remaining[i]
		for j := i + 1; j < len(remaining); j++ {
			slice[1] = remaining[j]
			for k := j + 1; k < len(remaining); k++ {
				slice[2] = remaining[k]
				slices <- slice
			}
		}
	}

	stop <- true
}

func RemoveIndex(s []int, index int) []int {
	var remaining []int
	for _, v := range s {
		if v != s[index] {
			remaining = append(remaining, v)
		}
	}
	return remaining
	//	return append(s[:index], s[index+1:]...)
}

func removeSlice(new []int, g []int) []int {
	remaining := make([]int, WORDS-GROUP)

	//{1,2,3}->{4,5,6,7,8}
	for i, v := range g {
		new = RemoveIndex(new, v-i-1)
	}

	for i := 0; i < (WORDS - GROUP); i++ {
		remaining[i] = new[i]
	}

	return remaining
}

func getCombs(c []int, g [GROUP + 1]int, cases chan int, wg *sync.WaitGroup) {
	slices := make(chan [GROUP]int)
	stop := make(chan bool)
	new := c
	var total int

	id := g[3]
	remaining := removeSlice(new, g[:3])
	fmt.Println("remaining array: ", id, remaining)

	go GetGroups2(remaining, slices, stop)
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
