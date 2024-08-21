package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)
	go readCh(ch)
	ch <- 10
	
	time.Sleep(3 *time.Second)
	

}


func readCh(ch chan int) {

	 new := <- ch
	 fmt.Println(new)

}