package main

import (
	"fmt"
	"sync"
)

func calculate(number int, resultChan chan map[int]int, wg *sync.WaitGroup) {
	defer wg.Done()
	square := number * number
	result := map[int]int{number: square}
	resultChan <- result
}
func main() {
	numbers := []int{1, 2, 3, 4, 5}
	var wg sync.WaitGroup

	resultChan := make(chan map[int]int)

	go func() {
		wg.Wait()
		close(resultChan)
	}()
	for _, num := range numbers {
		wg.Add(1)
		go calculate(num, resultChan, &wg)
	}

	for result := range resultChan {
		for k, v := range result {
			fmt.Printf("square of %d is %d\n", k, v)
		}
	}

	// if we directly iterate over the channel
	// for num := range resultChan {
	// 	fmt.Println(num)
	// }
	// output = map[5:25]
	// map[2:4]
	// map[1:1]
	// map[4:16]
	// map[3:9]

}

// this is the solution using struct

// type SquareResult struct {
// 	Number int
// 	Square int
// }

// func calculateSquare(resultChan chan SquareResult, number int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	square := number * number
// 	resultChan <- SquareResult{Number: number, Square: square}
// }

// func main() {
// 	numbers := []int{1, 2, 3, 4, 5}
// 	resultChan := make(chan SquareResult)
// 	var wg sync.WaitGroup

// 	go func() {
// 		wg.Wait()
// 		close(resultChan)
// 	}()

// 	for _, num := range numbers {
// 		wg.Add(1)
// 		go calculateSquare(resultChan, num, &wg)
// 	}

// 	for result := range resultChan {
// 		fmt.Printf("Square of %d is %d\n", result.Number, result.Square)
// 	}
// }
