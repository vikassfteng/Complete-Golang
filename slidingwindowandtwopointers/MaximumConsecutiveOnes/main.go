package main

import (
	"fmt"
)

func LongestConsecutiveOnes(s []int, k int) int {
	left, right := 0, 0
	maxLen := 0
	zeroes := 0
	n := len(s)
	for right < n {
		if s[right] == 0 {
			zeroes++
		}
		if zeroes > k {
			if s[left] == 0 {
				zeroes--
			}
			left++
		}
		if zeroes <= k {
			currLen := right - left + 1
			if currLen > maxLen {
				maxLen = currLen
			}
		}
		right++
	}
	return maxLen
}

func main() {
	nums := []int{1, 1, 0, 0, 1, 1, 1, 0, 1}
	k := 2
	result := LongestConsecutiveOnes(nums, k)
	fmt.Printf("The longest subarray with at most %d zeros is of length: %d\n", k, result)
}
