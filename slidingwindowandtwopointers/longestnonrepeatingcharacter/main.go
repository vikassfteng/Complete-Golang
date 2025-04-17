package main

import (
	"fmt"
)

func longestSubstring(s string) string {
	sMap := make(map[rune]int)

	start, maxLen := 0, 0
	resultStart := 0
	for i, ch := range s {
		// if the character is already in the map and its index is greater than or equal to start
		// then we need to update the start index to the next index after the last occurrence of the character
		// this ensures that we only keep unique characters in the current substring
		fmt.Println("checking------", ch, i, start)
		if lastIdx, found := sMap[ch]; found && lastIdx >= start {
			fmt.Println("found", ch, lastIdx, start)
			start = lastIdx + 1
		}
		//if the character is not in the map, we add it to the map
		sMap[ch] = i
		//update the max length if the current substring is longer than the previous max length
		//we also update the result start index to the current start index
		//this ensures that we always have the longest unique substring
		//if the current substring is longer than the previous max length
		//we update the max length and the result start index
		//this ensures that we always have the longest unique substring
		if i-start+1 > maxLen {
			maxLen = i - start + 1
			resultStart = start
		}
	}
	//return the longest unique substring
	//we use the result start index and the max length to get the substring
	return s[resultStart : resultStart+maxLen]

}

func main() {
	s := "abcabcabcdabc"
	fmt.Printf("longest unique substring in a given string is %s", longestSubstring(s))
}
