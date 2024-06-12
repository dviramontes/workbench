package workbench

func TwoSum(list []int, target int) []int {
	m := make(map[int]int)

	for i, v := range list {
		compliment := target - v
		if j, ok := m[compliment]; ok {
			return []int{j, i}
		} else {
			m[v] = i
		}
	}

	return nil
}
