package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortWords(s string) string {
	ch := strings.Split(s, "")
	sort.Strings(ch)
	return strings.Join(ch, "")
}

func main() {
	words := []string{"eat", "tea", "Tan", "ate", "Nat", "bat"}

	smap := make(map[string][]string)

	for _, word := range words {
		lowercase := strings.ToLower(word)
		srted := sortWords(lowercase)
		smap[srted] = append(smap[srted], word)
	}
	for _, word := range smap {
		fmt.Println(word)
	}
}
