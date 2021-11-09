package combinations

import (
	"fmt"
	"log"
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

func GetGroups(c []int, groups chan [GROUP]int, exit chan bool) {
	slice := [GROUP]int{}
	for i := 0; i < len(c); i++ {
		slice[0] = c[i]
		for j := i + 1; j < len(c); j++ {
			slice[1] = c[j]
			for k := j + 1; k < len(c); k++ {
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

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func removeSlice(c []int, g [GROUP]int) []int {
	remaining := make([]int, (len(c) - GROUP))

	//{1,2,3}->{4,5,6,7,8}
	for _, v := range g {
		remaining = RemoveIndex(c, v-1)
	}
	log.Println(remaining)
	return remaining
}

func getCombs(c []int, g [GROUP]int, cases chan int, wg *sync.WaitGroup) {
	slices := make(chan [GROUP]int)
	stop := make(chan bool)
	var total int

	c = removeSlice(c, g)

	go GetGroups(c, slices, stop)
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
