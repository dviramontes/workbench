package workbench

func RecursiveSum(list []int) int {
	if len(list) == 0 {
		return 0
	}

	head := list[0]
	tail := list[1:]

	return head + RecursiveSum(tail)
}
