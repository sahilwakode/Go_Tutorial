package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang -> go.dev")
	myCh := make(chan int)
	// myCh <- 5
	// fmt.Println(<-myCh)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		// fmt.Println(<-myCh)
		// fmt.Println(<-myCh)
		val, isChanopen := <-myCh
		fmt.Println(isChanopen)
		fmt.Println(val)
		wg.Done()
	}(myCh, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		// myCh <- 5
		// myCh <- 6
		close(myCh)
		wg.Done()
	}(myCh, wg)

	wg.Wait()

}
