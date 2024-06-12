package workbench

import "math/rand"

func QuickSort(list []int) []int {
	if len(list) < 2 {
		return list
	}

	r := rand.Intn(len(list))
	var same, less, greater []int
	pivot := list[r]
	for _, v := range list {
		if v > pivot {
			greater = append(greater, v)
		} else if v < pivot {
			less = append(less, v)
		} else {
			same = append(same, v)
		}
	}

	sorted := append(QuickSort(less), same...)
	sorted = append(sorted, QuickSort(greater)...)

	return sorted
}
