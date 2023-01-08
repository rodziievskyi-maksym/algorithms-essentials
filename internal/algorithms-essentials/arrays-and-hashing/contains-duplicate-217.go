package arrays_and_hashing

import "fmt"

func ContainsDuplicate() {
	/*
		Given an integer array nums, return true if any value appears at least twice in the array,
		and return false if every element is distinct.
	*/

	result := containsDuplicate([]int{1, 2, 3, 4, 5, 2, 4})
	fmt.Printf("Result: %t\n", result)
}

func containsDuplicate(nums []int) bool {
	// hashset
	set := NewSet()
	for _, num := range nums {
		// looping nums we check in hashset for existence, if yes return true which means the duplicate
		// otherwise add to hashset and check next value
		if set.Has(num) {
			return true
		}
		set.Add(num)
	}

	return false
}

type Set struct {
	// empty struct occupies zero bytes in storage
	// Basically an empty strut can be used every time you are only interested in a property of data
	// structure rather than its value.
	items map[int]struct{}
}

func NewSet() *Set {
	return &Set{items: make(map[int]struct{})}
}

func (s *Set) Has(val int) bool {
	_, ok := s.items[val]
	return ok
}

func (s *Set) Add(val int) {
	s.items[val] = struct{}{}
}
