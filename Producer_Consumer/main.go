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
