package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	n := 0
	p2 := 0
	p1 := 1
	return func() int {
		var p int
		switch n {
		case 0:
			p = 0
		case 1:
			p = 1
		default:
			p = p2 + p1
			p2 = p1
			p1 = p
		}
		n++
		return p
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

