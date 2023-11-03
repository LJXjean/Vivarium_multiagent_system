package main

import (
	"fmt"
	"sync"
)

func GoT() {
	defer fmt.Print("Ophélie Winter ")
	defer fmt.Print("is ")
	defer fmt.Print("Coming ")
}

// D'après *A Tour of Go*
func GoF() func() int {
	var a int = 0
	var b int = 1
	return func() int {
		var c int
		c = a
		a = b
		b = c + b
		return c
	}
}
func GoGo() {
	wg := new(sync.WaitGroup)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Print(i, " ")
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println()
}
func main() {
	// Part 1
	GoT()
	fmt.Println()
	// Part 2
	f := GoF()
	for i := 0; i < 6; i++ {
		fmt.Print(f(), " ")
	}
	fmt.Println()
	// Part 3
	GoGo()
}
