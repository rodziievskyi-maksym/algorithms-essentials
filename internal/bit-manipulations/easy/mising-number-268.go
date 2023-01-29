package bit_manipulations

import "fmt"

//https://leetcode.com/problems/missing-number/

func MissingNumber() {
	//Given an array nums containing n distinct numbers in the range [0, n], return the only number in the range that is missing from the array.

	/*
		Input: nums = [3,0,1]
		Output: 2
		Explanation: n = 3 since there are 3 numbers, so all numbers are in the range [0,3]. 2 is the missing number in the range since it does not appear in nums.
	*/

	fmt.Println(missingNumber([]int{3, 0, 1}))

}

func missingNumber(nums []int) int {
	// 101 ^ 01 = 111 = 7
	res := 0
	for i := 0; i < len(nums); i++ {
		//0 1 3
		//2 2 0
		//0 3 1 = 3 ^ 1 -> 011 ^ 001 = 010 ->2
		fmt.Println(res, i+1, nums[i])
		res = res ^ (i + 1) ^ nums[i]
	}
	return res
}
