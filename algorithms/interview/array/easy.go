package array

// Given a sorted array, remove the duplicates in-place such that each element appear only once and return the new length.
// Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/727/
func removeDuplicates(nums []int) int {
	var length int
	for _, v := range nums {
		if length == 0 || v != nums[length-1] {
			nums[length] = v
			length++
		}
	}
	return length
}
