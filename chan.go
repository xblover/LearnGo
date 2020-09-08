package main

import "fmt"

func testChannel() {
	c := make(chan int)
	defer close(c)

	go func() {
		c <- 3 + 4
	}()

	fmt.Println(<-c)
}
