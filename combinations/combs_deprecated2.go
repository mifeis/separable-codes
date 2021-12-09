package combinations

import (
	"fmt"
	"math/rand"
	"sync"

	"github.com/mifeis/Separable-Codes/lib"
)

func List(c []int) int {
	combins := make(chan lib.Combin)
	exit := make(chan bool)
	cases := make(chan int, 10000)
	wg := sync.WaitGroup{}
	var total int

	go GetGroups(true, c, combins, exit)

	for {
		select {
		case comb := <-combins:
			wg.Add(1)
			go func(c []int, g [lib.GROUP]int, id int, cases chan int, wg *sync.WaitGroup) {
				slices := make(chan lib.Combin)
				stop := make(chan bool)
				var total int

				remaining := lib.RemoveSlice(c[:], g[:])
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

func GetGroups(first bool, c []int, combins chan lib.Combin, exit chan bool) {
	var comb lib.Combin
	slice := [lib.GROUP]int{}
	for i := 0; i < len(c); i++ {
		slice[0] = c[i]
		for j := i + 1; j < len(c); j++ {
			slice[1] = c[j]
			for k := j + 1; k < len(c); k++ {
				slice[2] = c[k]
				if first {
					comb = lib.Combin{
						Group: slice,
						Id:    rand.Intn(1000),
					}
					fmt.Println("...Sending slice", comb.Id, "...", comb.Group)
				} else {
					comb = lib.Combin{
						Group: slice,
					}
				}
				combins <- comb
			}
		}
	}

	exit <- true
}
