package main

import (
	"testing"

	"github.com/mifeis/Separable-Codes/combinations"
	"github.com/mifeis/Separable-Codes/lib"
)

const (
	WORDS = lib.WORDS
	GROUP = lib.GROUP
)

func TestMatch(t *testing.T) {

	var c []int

	groups := make(chan combinations.Combin)
	exit := make(chan bool)

	c = combinations.Init()
	go combinations.GetGroups(true, c, groups, exit)

	for {
		select {
		case g := <-groups:
			go match(t, g.Group)
		case <-exit:
			return
		}
	}
}

func match(t *testing.T, g [GROUP]int) {
	match := [GROUP]int{6, 2, 1}
	var m int

	for i := 0; i < GROUP; i++ {
		if g[0] == match[i] {
			m++
		} else if g[1] == match[i] {
			m++
		} else if g[2] == match[i] {
			m++
		}
	}
	if m != 3 {
		t.SkipNow()
	}
}
