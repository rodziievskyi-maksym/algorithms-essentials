package arrays_and_hashing

//https://leetcode.com/problems/missing-number/

func MissingNumber() {
	//Given an array nums containing n distinct numbers in the range [0, n], return the only number in the range that is missing from the array.

	/*
		Input: nums = [3,0,1]
		Output: 2
		Explanation: n = 3 since there are 3 numbers, so all numbers are in the range [0,3]. 2 is the missing number in the range since it does not appear in nums.
	*/

	missingNumber([]int{3, 0, 1})
}

func missingNumber(nums []int) int {
	var total, length int

	length = len(nums)
	for _, num := range nums {
		if num <= length {
			total += num
		}
	}

	// Tricky solving cuz wen can use here arithmetics progression formula -> Sn = (a1 + an) * n / 2 -> result is sum of all numbers in array
	// what left is just extract sum of current array without missing number from formula result's array
	// We take a length of array of adding + 1 to make it "full" with missing number
	return length*(length+1)/2 - total

}
