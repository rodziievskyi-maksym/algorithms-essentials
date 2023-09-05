package main

import (
	"fmt"
	arrays_and_hashing "github.com/max-rodziyevsky/algorithms-essentials/internal/arrays-and-hashing/easy"
)

func main() {
	run()
}

func run() {
	printProcessingMessage()

	//easy arrays and hashing
	//arrays_and_hashing.ContainsDuplicate()
	//arrays_and_hashing.ValidAnagram()
	//arrays_and_hashing.TwoSums()
	//arrays_and_hashing.MissingNumber()
	arrays_and_hashing.ReplaceElements()

	// Medium arrays and hashing
	//arrays_and_hashing.GroupAnagrams()

	//stack.ValidParentheses()

	//bit_manipulations.MissingNumber()

	// easy two pointers
	//two_pointers.ValidPalindrome()
	//two_pointers.ValidPalindrome2()
}

func printProcessingMessage() {
	fmt.Println("=================================")
	fmt.Println("Main function is processing...")
	fmt.Println("=================================\n")
}
