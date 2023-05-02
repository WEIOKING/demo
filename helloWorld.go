package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var count = make(chan int)

func main() {
	for i := 0; i < 10; i++ {
		go hello()
		count <- i
		wg.Add(1)
	}
	wg.Wait()
	ticker := time.NewTicker(time.Second)
	i := 5
	for {
		<-ticker.C
		fmt.Println(i)
		i--
		if i == 0 {
			break
		}
	}
	ticker.Stop()
}

func hello() {
	fmt.Println("hello world!!! ", <-count)
	wg.Done()
}
