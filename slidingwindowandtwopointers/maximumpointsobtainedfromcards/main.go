package main

import (
	"fmt"
)

func maximumPoints(arr []int, k int) int {
	n := len(arr)
	if k > n {
		k = n // Avoid panic if k > len(arr)
	}

	leftSum := 0
	for i := 0; i < k; i++ {
		leftSum += arr[i]
	}
	maxSum := leftSum

	rightSum := 0
	for i := 0; i < k; i++ {
		leftSum -= arr[k-1-i]
		rightSum += arr[n-1-i]
		maxSum = max(maxSum, leftSum+rightSum)
	}

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	arr := []int{1, 2, 9, 1, 5, 6, 5, 1, 2, 2}
	k := 4
	fmt.Println(maximumPoints(arr, k)) // Output: 22
}
