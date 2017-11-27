package easy

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
