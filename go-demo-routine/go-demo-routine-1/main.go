package main

import (
	"fmt"
	"sync"
)

/**
 * golang 多个协程并发执行
**/
func Foo(idx int, wg *sync.WaitGroup) {
	fmt.Printf("test go routines idx -> %d\n", idx)
	wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go Foo(i, wg)
	}
	wg.Wait()
}
