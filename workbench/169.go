package workbench

func MajorityElement(nums []int) int {
	seen := make(map[int]int)
	for _, n := range nums {
		seen[n]++
		if seen[n] > len(nums)/2 {
			return n
		}
	}
	return 0
}
