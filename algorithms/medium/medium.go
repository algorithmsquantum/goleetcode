package medium

// https://leetcode.com/problems/add-two-numbers/description/

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
