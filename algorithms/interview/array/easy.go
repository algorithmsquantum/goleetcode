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

// Given an array of integers, every element appears twice except for one. Find that single one.
// Note: Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory ?
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/549/
func singleNumber(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		if _, ok := m[v]; !ok {
			m[v] = v
		} else {
			delete(m, v)
		}
	}
	for k := range m {
		return k
	}
	return 0
}
