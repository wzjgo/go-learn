package sort

func QuickSort(a []int) {
	quickSort(a, 0, len(a)-1)
}

func quickSort(a []int, start, end int) {
	if end > start {
		key := a[start]
		i := start
		j := end
		for i <= j {
			for a[i] < key {
				i++
			}
			for a[j] > key {
				j--
			}
			if i <= j {
				a[i], a[j] = a[j], a[i]
				i++
				j--
			}

		}
		if start < j {
			quickSort(a, start, j)
		}
		if end > i {
			quickSort(a, i, end)
		}

	}
}
