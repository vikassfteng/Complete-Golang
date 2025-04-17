package main

import "fmt"

func Median(num1 []int, num2 []int) float64 {
	mergedArray := merge(num1, num2)
	lenOfMergedArray := len(mergedArray)
	if lenOfMergedArray%2 == 1 {
		return float64(mergedArray[lenOfMergedArray/2])
	} else {
		return float64(((mergedArray[lenOfMergedArray/2-1]) + (mergedArray[lenOfMergedArray/2])) / 2)
	}
}
func merge(num1 []int, num2 []int) []int {
	m := len(num1)
	n := len(num2)
	merged := make([]int, 0, m+n)

	i, j := 0, 0
	for i < m && j < n {
		if num1[i] < num2[j] {
			merged = append(merged, num1[i])
			i++
		} else {
			merged = append(merged, num2[j])
			j++
		}
	}
	for i < m {
		merged = append(merged, num1[i])
		i++
	}
	for j < n {
		merged = append(merged, num2[j])
		j++
	}
	return merged
}

func main() {
	num1 := []int{1, 3, 5}
	num2 := []int{2, 6, 9}
	output := Median(num1, num2)
	fmt.Print(output)
}

// Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

// Example 1:
// Input: nums1 = [1,3], nums2 = [2]
// Output: 2.00000
// Explanation: merged array = [1,2,3] and median is 2.

// Example 2:
// Input: nums1 = [1,2], nums2 = [3,4]
// Output: 2.50000
// Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
