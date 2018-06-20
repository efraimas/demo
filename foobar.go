package main

import (
	"fmt"
	"os"
)

//var y = "ok"

func test(r []int) int {
	c := r[0]
	for _, i := range r {
		if i < c {
			c = i
		}
	}
	return c
}

func test2(r ...int) int {
	c := r[0]
	for _, i := range r {
		if i < c {
			c = i
		}
	}
	return c
}

func main() {
	myfunc := func() int {
		return 0
	}
	fmt.Println(os.Args[1])
	x := []int{48, 96, 86, 68, 57, 34, 83, 27, 19, 97, 3, 9, 17}
	fmt.Println(test(x))
	n, e := os.Hostname()
	fmt.Println(n)
	//fmt.Println(e)
}
