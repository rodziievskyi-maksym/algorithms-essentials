package main

import (
	"fmt"
	bit_manipulations "github.com/max-rodziyevsky/algorithms-essentials/internal/bit-manipulations/easy"
)

func main() {
	//if err := run(); err != nil {
	//	log.Fatalf("application returned error: %s", err)
	//}

	run()
}

func run() {
	printProcessingMessage()

	//easy
	//arrays_and_hashing_easy.ContainsDuplicate()
	//arrays_and_hashing_easy.ValidAnagram()
	//arrays_and_hashing_easy.TwoSums()
	//arrays_and_hashing_easy.MissingNumber()
	//
	//// Medium
	//arrays_and_hashing_medium.GroupAnagrams()
	//
	//stack.ValidParentheses()

	bit_manipulations.MissingNumber()

}

func printProcessingMessage() {
	fmt.Println("=================================")
	fmt.Println("Main function is processing...")
	fmt.Println("=================================\n")
}
