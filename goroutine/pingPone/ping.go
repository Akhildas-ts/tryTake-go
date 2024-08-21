package main

import (
	"fmt"
)

// ping pong with go routine and channel

func pingPong(exchange int) {

	for exchange > 0 {

		chPing := make(chan string)
		chPong := make(chan string)
         
		go func() {

			chPing <- "ping"

			// fmt.Println("ping")

		}()
		fmt.Println(<-chPing)

		go func() {

			// fmt.Println("pong")

			chPong <- "pong"
		}()

		fmt.Println(<-chPong)
		// time.Sleep(2 * time.Second)
		exchange--
		println()
	}

}

func main() {

	pingPong(8)

}
