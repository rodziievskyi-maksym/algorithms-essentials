package two_pointers

import (
	"fmt"
	"unicode"
)

func ValidPalindrome2() {
	//Given a string s, return true if the s can be palindrome after deleting at most one character from it.

	//Input:
	//s := "abcwwca"
	//Output: true

	//Input:
	s2 := "atbbga"
	//Output: true
	//Explanation: You could delete the character 'c'.

	//fmt.Println(validPalindrome(s))
	fmt.Println(validPalindrome(s2))
}

func validPalindrome(s string) bool {
	if len(s) <= 2 {
		return true
	}

	leftIndex := 0
	rightIndex := len(s) - 1

	for leftIndex < rightIndex {
		left, right := runeUtil(rune(s[leftIndex]), rune(s[rightIndex]))

		if left == right {
			leftIndex++
			rightIndex--
		} else {
			return validPalindromeUtil(s, leftIndex+1, rightIndex) || validPalindromeUtil(s, leftIndex, rightIndex-1)
		}
	}

	return true
}

func validPalindromeUtil(s string, leftIndex, rightIndex int) bool {
	for leftIndex < rightIndex {
		left, right := runeUtil(rune(s[leftIndex]), rune(s[rightIndex]))

		if left == right {
			leftIndex++
			rightIndex--
		} else {
			return false
		}
	}

	return true
}

func runeUtil(left, right rune) (rune, rune) {
	return unicode.ToLower(left), unicode.ToLower(right)
}
