package main

import (
	"fmt"
)

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	start, end := 0, 0

	for i := 0; i < len(s); i++ {
		// Odd-length palindrome
		l1, r1 := expandAroundCenter(s, i, i)
		if r1-l1 > end-start {
			start, end = l1, r1
		}

		// Even-length palindrome
		l2, r2 := expandAroundCenter(s, i, i+1)
		if r2-l2 > end-start {
			start, end = l2, r2
		}
	}

	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	// return last valid bounds
	return left + 1, right - 1
}

func main() {
	input := "cbbd"
	result := longestPalindrome(input)
	fmt.Println("Longest Palindromic Substring:", result)
}
