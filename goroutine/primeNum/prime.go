package main

import (
	"fmt"
	"math"
)


// frist prime numbers with go routine 
// in these case we get a number "n" so we need to find range "N" of prime number's with go rouinte by divide work divide into 2 parts 


func goPrime(n int) {

	half := n / 2

	var result = make([]int, 0)
	var result2 = make([]int, 0)
	fristCh := make(chan int)
	secondCh := make(chan int)

	go func() {

		defer close(fristCh)

		for i := 2; i < half; i++ {
			

			ok, val := isPrime(i)
			if ok {

				fristCh <- val

			}
		}

	}()

	for fVal := range fristCh {

		result = append(result, fVal)
	}

	for _, val := range result {
		fmt.Println("frist half prime number :",val)

	}

	go func() {

		defer close(secondCh)

		for i := half; i <= n; i++ {

            
			ok, val := isPrime(i)
			if ok {
                  
				secondCh <- val

			}

		}

	}()

	for sVal := range secondCh {

		result2 = append(result2, sVal)
	}

	for _,val  := range result2 {

		fmt.Println("seond half prime number : ", val)
	}

}

func isPrime(n int) (bool, int) {
	if n < 2 {
		return false, 0
	}

	
	// fmt.Println("square root of %d", n, ": ", p)
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false, 0
		}
	}
	return true, n
}

func main() {

	goPrime(7)

}
