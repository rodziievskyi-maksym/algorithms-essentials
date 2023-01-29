package main

import (
	"fmt"
	arrays_and_hashing_easy "github.com/max-rodziyevsky/algorithms-essentials/internal/arrays-and-hashing/easy"
	arrays_and_hashing_medium "github.com/max-rodziyevsky/algorithms-essentials/internal/arrays-and-hashing/medium"
	stack "github.com/max-rodziyevsky/algorithms-essentials/internal/stack/easy"
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
	arrays_and_hashing_easy.ContainsDuplicate()
	arrays_and_hashing_easy.ValidAnagram()
	arrays_and_hashing_easy.TwoSums()
	arrays_and_hashing_easy.MissingNumber()

	// Medium
	arrays_and_hashing_medium.GroupAnagrams()

	stack.ValidParentheses()

}

func printProcessingMessage() {
	fmt.Println("=================================")
	fmt.Println("Main function is processing...")
	fmt.Println("=================================\n")
}
