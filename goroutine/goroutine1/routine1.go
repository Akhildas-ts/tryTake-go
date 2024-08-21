package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)

	go func() {

		defer close(ch)
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		fmt.Println("done")
	}()

	time.Sleep(10 * time.Second)

}
