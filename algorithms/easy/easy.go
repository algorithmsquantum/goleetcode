package easy

import (
	"math"
)

// https://leetcode.com/problems/two-sum/description/
// Given an array of integers, return indices of the two numbers such that they add up to a specific target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// Example:
// Given nums = [2, 7, 11, 15], target = 9,
// Because nums[0] + nums[1] = 2 + 7 = 9,
// return [0, 1].
func twoSum(nums []int, target int) []int {
	n := len(nums)
	m := make(map[int]int, n)

	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			return []int{i, j}
		}
		m[v] = i
	}

	return nil
}

// https://leetcode.com/problems/reverse-integer/description/
// Given a 32-bit signed integer, reverse digits of an integer.
// Example 1:
// Input: 123
// Output:  321
// Example 2:
// Input: -123
// Output: -321
// Example 3:
// Input: 120
// Output: 21
// Note:
// Assume we are dealing with an environment which could only hold integers within the 32-bit signed integer range.
// For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
func reverse0(x int) int {
	posi := true
	if x < 0 {
		posi = false
		x = -x
	}

	a := make([]int, 0)
	for {
		rem := x % 10
		a = append(a, rem)
		if cons := x / 10; cons == 0 {
			break
		} else {
			x = cons
		}
	}

	x = 0
	for i := range a {
		v := a[i]
		for j := 0; j < (len(a) - 1 - i); j++ {
			v = v * 10
		}
		x += v
	}

	if x < math.MinInt32 || x > math.MaxInt32 {
		return 0
	}
	if !posi {
		return -x
	}
	return x
}

func reverse(x int) int {
	cons, rem, result := x, 0, 0 // 商, 余数, 结果
	for cons != 0 {
		rem = cons % 10
		result = result*10 + rem
		cons = cons / 10
		if result > math.MaxInt32 || result < math.MinInt32 {
			return 0
		}
	}
	return result
}

// https://leetcode.com/problems/palindrome-number/description/
// Determine whether an integer is a palindrome. Do this without extra space.
// Some hints:
// Could negative integers be palindromes? (ie, -1)
// If you are thinking of converting the integer to string, note the restriction of using extra space.
// You could also try reversing an integer. However, if you have solved the problem "Reverse Integer", you know that the reversed integer might overflow. How would you handle such case?
// There is a more generic way of solving this problem.
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	cons, rem, result := x, 0, 0 // 商, 余数, 结果
	for cons != 0 {
		rem = cons % 10
		result = result*10 + rem
		cons = cons / 10
		if result > math.MaxInt32 || result < math.MinInt32 {
			return false
		}
	}
	return result == x
}
