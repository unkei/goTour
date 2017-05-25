package main

import (
	"fmt"
	"time"
)

func sendLater(c chan int) {
	time.Sleep(3000 * time.Millisecond)
	c <- 2
}

func main() {
	ch := make(chan int, 2)
	ch <- 1
//	ch <- 2
	go sendLater(ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

