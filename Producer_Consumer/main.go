package main

import "fmt"

func Consumer(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

func Producer(ch chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}

func main() {
	ch := make(chan int)

	go Producer(ch)
	Consumer(ch)
}


// package main

// import (
// 	"fmt"
// )

// func Producer(in <-chan int, out chan<- int) {
// 	val := <-in
// 	val += 1
// 	fmt.Println("Producer processed:", val)
// 	out <- val
// }

// func Consumer(ch <-chan int) {
// 	val := <-ch
// 	fmt.Println("Consumer received:", val)
// }

// func main() {
// 	input := make(chan int)
// 	output := make(chan int)


// 	go Producer(input, output)
// 	go Consumer(output)

// 	// Send initial value
// 	input <- 12

// 	// Wait (optional) — or block main from exiting too early
// 	var dummy string
// 	fmt.Scanln(&dummy)
// }
