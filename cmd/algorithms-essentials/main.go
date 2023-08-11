package main

import (
	"fmt"
	two_pointers "github.com/max-rodziyevsky/algorithms-essentials/internal/two-pointers/easy"
)

func main() {
	run()
}

func run() {
	printProcessingMessage()

	//easy arrays and hashing
	//arrays_and_hashing_easy.ContainsDuplicate()
	//arrays_and_hashing_easy.ValidAnagram()
	//arrays_and_hashing_easy.TwoSums()
	//arrays_and_hashing_easy.MissingNumber()

	// Medium arrays and hashing
	//arrays_and_hashing_medium.GroupAnagrams()

	//stack.ValidParentheses()

	//bit_manipulations.MissingNumber()

	// easy two pointers
	//two_pointers.ValidPalindrome()
	two_pointers.ValidPalindrome2()
}

func printProcessingMessage() {
	fmt.Println("=================================")
	fmt.Println("Main function is processing...")
	fmt.Println("=================================\n")
}
