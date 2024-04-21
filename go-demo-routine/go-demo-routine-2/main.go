package main

import "fmt"

/**
 * golang 管道在协程同步上的应用
**/

func main() {
	ch := make(chan int, 1)

	go func() {
		fmt.Printf("This is a foo.\n")
		ch <- 1
	}()
	fmt.Printf("this is main func.\n")
	<-ch
}
