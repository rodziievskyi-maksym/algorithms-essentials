package main

import (
	"fmt"
	arrays_and_hashing "github.com/max-rodziyevsky/algorithms-essentials/internal/algorithms-essentials/arrays-and-hashing"
	"github.com/max-rodziyevsky/algorithms-essentials/internal/algorithms-essentials/stack"
)

func main() {
	//if err := run(); err != nil {
	//	log.Fatalf("application returned error: %s", err)
	//}

	run()
}

func run() {
	printProcessingMessage()

	//arrays_and_hashing.ContainsDuplicate()
	//arrays_and_hashing.ValidAnagram()
	//arrays_and_hashing.TwoSums()

	// Medium
	arrays_and_hashing.GroupAnagrams()

	stack.ValidParentheses()

}

func printProcessingMessage() {
	fmt.Println("=================================")
	fmt.Println("Main function is processing...")
	fmt.Println("=================================\n")
}
