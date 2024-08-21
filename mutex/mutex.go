package main

import (
	"fmt"
	"sync"
)

var Commonparam = 0

func increment(wg *sync.WaitGroup, m *sync.Mutex) {

	defer wg.Done()

	// m.Lock()
	Commonparam += 1
	fmt.Println(Commonparam)

	// m.Unlock()

}

func main() {

  var w sync.WaitGroup

	var m sync.Mutex


	for i:=1; i<=100;i++ {
   
		w.Add(1)
		  go increment(&w,&m)

	}

	w.Wait()

	fmt.Println(Commonparam)

}
