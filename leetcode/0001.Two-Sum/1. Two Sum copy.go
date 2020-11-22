package leetcode

func twoSum1(num []int, target int) []int {
	m := make(map[int]int)
	for i, a := range num {
		if idx, ok := m[target-a]; ok {
			return []int{idx, i}
		}
		m[a] = i
	}
	return nil
}


