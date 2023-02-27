package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func myFunction(first chan<- int) {
	first <- 234
	fmt.Println("myfucntion called")
	wg.Done()
}

func function2(value <-chan int) {
	got := <-value
	fmt.Println("function2 called value is ", got)
	wg.Done()
}

func main() {
	wg.Add(2)

	mychanel := make(chan int)
	fmt.Println("hello strated")

	go myFunction(mychanel)
	//got := <-mychanel
	go function2(mychanel)
	//mychanel <- got
	//<-mychanel
	wg.Wait()
	fmt.Println("end function")

}
