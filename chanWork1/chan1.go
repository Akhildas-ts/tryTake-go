package main
import "fmt"              

func genteor() <-chan int {

	c := make(chan int)

	go func() {             

		for i := 0; i < 10; i++ {

			c <- i
		}

		close(c)
	}()

	return c

}

// func reciver(c <-chan int) {

// 	for val := range c {

// 		fmt.Println(val)
// 	}

// }

func main() {





	ch := genteor()

	for val := range ch{{

		fmt.Println(val)
	}}

	fmt.Println(ch)

	// reciver(ch)

}
