package arrays_and_hashing

import "fmt"

//https://leetcode.com/problems/group-anagrams/

//Given an array of strings strs, group the anagrams together. You can return the answer in any order.
//An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

//Input: strs = ["eat","tea","tan","ate","nat","bat"]
//Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

func GroupAnagrams() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println("Answer:", groupAnagrams(strs))

	fmt.Println('n' - 'a' + 1)
	fmt.Println('a' - 'a' + 1)
	fmt.Println('t' - 'a' + 1)

	mapa := map[[10]int]string{}
	k := [10]int{}
	k = k
	mapa[k] = "here"
	k2 := [10]int{}
	mapa[k2] = "there"

	fmt.Println(mapa)

}

func groupAnagrams(strs []string) [][]string {
	// we create map where keys equal to array of integers no more than 26 due to alphabets letters amount.
	// and values are slice of strings
	// [1,2,4,4,5,6]:["ten","eat","net"]
	mp := map[[26]int][]string{}
	// looping our slice with words
	for _, word := range strs {
		k := [26]int{}
		// looping through each words
		for i := 0; i < len(word); i++ {
			// here we write into k array each letter on a word minus character 'a' which helps us to map 'a' -> 0
			// 'b' -> 1 .... 'z' -> 25 and than we just add 1
			k[word[i]-'a'] += 1
		}
		// store iterated word with corresponding set of numbers which means ascii character code
		// i.e: [5,1,20]:["eat", "tea", "eat"],
		//      [14,1,20]:["tan", "nat"],
		mp[k] = append(mp[k], word)
	}
	var result [][]string
	for _, value := range mp {
		result = append(result, value)
	}
	return result
}
