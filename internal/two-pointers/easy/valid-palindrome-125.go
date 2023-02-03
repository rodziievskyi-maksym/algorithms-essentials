package two_pointers

import (
	"fmt"
	"unicode"
)

//https://leetcode.com/problems/valid-palindrome/

func ValidPalindrome() {
	/*A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing
	all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.*/

	//Given a string s, return true if it is a palindrome, or false otherwise

	/*
		Input: s = "A man, a plan, a canal: Panama"
		Output: true
		Explanation: "amanaplanacanalpanama" is a palindrome.
	*/

	/*
		Input: s = " "
		Output: true
		Explanation: s is an empty string "" after removing non-alphanumeric characters.
		Since an empty string reads the same forward and backward, it is a palindrome.
	*/

	s := "A man, a plan, a canal: Panama"
	empty := " "
	tricky := "0P"
	fmt.Println('0')
	fmt.Println('O')
	fmt.Println(isPalindrome(s))
	fmt.Println(isPalindrome(empty))
	fmt.Println(isPalindrome(tricky))
}

func isPalindrome(s string) bool {
	// Time: O(n)
	// Space: O(1)
	if len(s) < 2 {
		return true
	}

	//two pointers
	left := 0
	right := len(s) - 1
	arr := []rune(s)

	for left <= right {
		leftChar := unicode.ToLower(arr[left])
		rightChar := unicode.ToLower(arr[right])

		if !isAlphanumeric(leftChar) {
			left++
			continue
		}

		if !isAlphanumeric(rightChar) {
			right--
			continue
		}

		if leftChar != rightChar {
			return false
		}

		left++
		right--
	}

	return true
}

func isAlphanumeric(c rune) bool {
	return 'a' <= c && c <= 'z' || c >= '0' && c <= '9'

}
