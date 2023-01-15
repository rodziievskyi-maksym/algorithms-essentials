package arrays_and_hashing

import "fmt"

func TwoSums() {
	/*
		Input: nums = [2,7,11,15], target = 9
		Output: [0,1]
		Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
	*/

	//fmt.Printf("Two Sums: %v\n", twoSum([]int{3, 2, 4}, 6))
	fmt.Printf("Two Sums: %v\n", twoSum2([]int{3, 2, 4}, 6))

}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		num := nums[i] // 3
		for j := i + 1; j < len(nums); j++ {
			if (num + nums[j]) == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

func twoSum2(nums []int, target int) []int {
	hashMap := make(map[int]int, len(nums))

	for i, num := range nums {
		result := target - num
		if _, ok := hashMap[result]; ok {
			return []int{hashMap[result], i}
		}
		hashMap[num] = i
	}

	return nil
}
