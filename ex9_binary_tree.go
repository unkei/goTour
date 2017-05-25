package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	//	fmt.Println("Left=", t.Left, ", Value=", t.Value, ", Right=", t.Right)
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := 0; i < 10; i++ {
		if <-ch1 != <-ch2 {
			return false
		}
	}
	return true
}

func PrintTree(hdr string, ch chan int) {
	fmt.Print(hdr, "=")
	for i := 0; i < 10; i++ {
		if i > 0 {
			fmt.Print(",")
		}
		fmt.Print(<-ch)
	}
	fmt.Println()
}

func main() {
	ch := make(chan int, 10)
	go Walk(tree.New(1), ch)
	PrintTree("tree.New(1)", ch)

	go Walk(tree.New(2), ch)
	PrintTree("tree.New(2)", ch)

	fmt.Print("Same(tree.New(1), tree.New(1))=", Same(tree.New(1), tree.New(1)), "\n")
	fmt.Print("Same(tree.New(1), tree.New(2))=", Same(tree.New(1), tree.New(2)), "\n")

}

