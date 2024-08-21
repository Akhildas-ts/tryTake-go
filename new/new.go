package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {


	var n Person

	n.name = "arun "
	n.age = 24

	p := new(Person)

	p.name = "akhil"
	p.age = 18


	
	

	fmt.Printf("name: %s,  age : %d",n.name,n.age)


}