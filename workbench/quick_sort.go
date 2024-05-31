package workbench

func QuickSort(list []int) (sorted []int) {
	if len(list) < 2 {
		return list
	}

	pivot := list[0]

	var lessThan, greaterThan, same []int

	for _, v := range list {
		if v < pivot {
			lessThan = append(lessThan, v)
		} else if v > pivot {
			greaterThan = append(greaterThan, v)
		} else {
			same = append(same, v)
		}
	}

	l := QuickSort(lessThan)
	g := QuickSort(greaterThan)

	sorted = append(l, same...)
	sorted = append(sorted, g...)

	return sorted
}
