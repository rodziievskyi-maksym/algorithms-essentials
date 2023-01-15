package arrays_and_hashing

import (
	"fmt"
)

func ValidAnagram() {
	stringOne := "ac"
	stringTwo := "ac"

	fmt.Printf("Valid Anagram: %t\n", isAnagram(stringOne, stringTwo))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// we need HashMap for this problem
	hashMap := make(map[rune]int)

	for _, char := range s {
		hashMap[char]++ // map[97: 1, 99: 1]
	}

	for _, charT := range t {
		if _, ok := hashMap[charT]; !ok {
			return false
		}

		if hashMap[charT] != 0 {
			hashMap[charT]-- // map[97: 1, 99: 0,]
		}
	}

	for _, value := range hashMap {
		if value > 0 {
			return false
		}
	}

	return true
}
