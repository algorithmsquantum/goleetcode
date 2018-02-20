package medium

import (
	"math"
	"unicode"
)

// https://leetcode.com/problems/add-two-numbers/description/
// You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order and each of their nodes contain a single digit.
// Add the two numbers and return it as a linked list.
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 0 -> 8

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	p, q, curr := l1, l2, dummyHead
	carry := 0
	for p != nil || q != nil {
		var x, y int
		if p != nil {
			x = p.Val
		}
		if q != nil {
			y = q.Val
		}
		sum := carry + x + y
		carry = sum / 10
		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next
		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
	}
	if carry > 0 {
		curr.Next = &ListNode{Val: carry}
	}
	return dummyHead.Next
}

// https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
// Given a string, find the length of the longest substring without repeating characters.
// Examples:
// Given "abcabcbb", the answer is "abc", which the length is 3.
// Given "bbbbb", the answer is "b", with the length of 1.
// Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
func lengthOfLongestSubstring(s string) int {
	sr := []rune(s)
	n := len(sr)
	m := make(map[rune]bool, n)
	var maxLen, i, j int

	for i < n && j < n {
		if _, ok := m[sr[j]]; !ok {
			m[sr[j]] = true
			j++
			subLen := j - i
			if subLen > maxLen {
				maxLen = subLen
			}
		} else {
			delete(m, sr[i])
			i++
		}
	}

	return maxLen
}

// https://leetcode.com/problems/string-to-integer-atoi/description/
// Implement atoi to convert a string to an integer.
// Hint: Carefully consider all possible input cases. If you want a challenge, please do not see below and ask yourself what are the possible input cases.
// Notes: It is intended for this problem to be specified vaguely (ie, no given input specs). You are responsible to gather all the input requirements up front.
// Requirements for atoi:
// The function first discards as many whitespace characters as necessary until the first non-whitespace character is found. Then, starting from this character, takes an optional initial plus or minus sign followed by as many numerical digits as possible, and interprets them as a numerical value.
// The string can contain additional characters after those that form the integral number, which are ignored and have no effect on the behavior of this function.
// If the first sequence of non-whitespace characters in str is not a valid integral number, or if no such sequence exists because either str is empty or it contains only whitespace characters, no conversion is performed.
// If no valid conversion could be performed, a zero value is returned. If the correct value is out of the range of representable values, INT_MAX (2147483647) or INT_MIN (-2147483648) is returned.
func myAtoi(str string) int { // nolint: gocyclo
	nums, ops := make([]int, 0), make([]rune, 0) // 数字, 操作符
	for _, v := range str {
		if v == '-' || v == '+' {
			if len(nums) != 0 {
				break
			}
			if len(ops) == 1 {
				return 0
			}
			ops = append(ops, v)
		} else if unicode.IsSpace(v) {
			if len(ops) > 0 && len(nums) == 0 {
				return 0
			}
			if len(nums) == 0 {
				continue
			} else {
				break
			}
		} else if v-'0' >= 0 && v-'0' <= 9 {
			nums = append(nums, int(v-'0'))
		} else {
			if len(nums) == 0 {
				return 0
			}
			if len(nums) > 0 {
				break
			}
		}
	}

	pos := 1 // 正/负
	if len(ops) > 0 && ops[0] == '-' {
		pos = -1
	}
	result := 0
	over := false // 是否溢出
	for i := range nums {
		v := 1
		for j := i; j < len(nums)-1; j++ {
			if math.MaxInt32/10 < 10 {
				over = true
				break
			}
			v *= 10
		}
		if math.MaxInt32/v < nums[i] {
			over = true
			break
		}
		result += nums[i] * v
	}

	if over {
		if pos > 0 {
			return math.MaxInt32
		}
		return math.MinInt32
	}

	result *= pos
	if result > math.MaxInt32 {
		result = math.MaxInt32
	}
	if result < math.MinInt32 {
		result = math.MinInt32
	}
	return result
}
