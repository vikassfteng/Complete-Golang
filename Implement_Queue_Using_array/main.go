package main

import (
	"fmt"
)

const MAX = 5

type Queue struct {
	arr               [MAX]int
	size, rear, front int
}

func (q *Queue) Enqueue(val int) {
	if q.size == MAX {
		fmt.Println("Overflow")
	}
	q.arr[q.rear] = val
	q.rear = (q.rear + 1) % MAX
	q.size++
}

func (q *Queue) Dequeu() {
	if q.size == 0 {
		fmt.Println("empty")
	}
	// val := q.arr[q.front]
	q.front = (q.front + 1) % MAX
	q.size--

}

func main() {
	q := Queue{}
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	q.Enqueue(40)
	q.Dequeu()
	fmt.Println("Array1", q.arr)
	q.Dequeu()
	fmt.Println("Array2", q.arr)
	q.Dequeu()
	fmt.Println("Array3", q.arr)
	q.Enqueue(50)
	fmt.Println("Array4", q.arr)
	q.Enqueue(60)
	fmt.Println("Array5", q.arr)
}
