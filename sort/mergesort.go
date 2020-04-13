package merge

func MergeSort(a []int) {
	tmp := make([]int, len(a))
	mergeSort(a, 0, len(a)-1, tmp)
}

func mergeSort(a []int, start, end int, tmp []int) {
	if start < end {
		mid := (start + end) / 2
		mergeSort(a, start, mid, tmp)
		mergeSort(a, mid+1, end, tmp)
		merge(a, start, mid, end, tmp)
	}
}

func merge(a []int, start, mid, end int, tmp []int) {
	i := start
	j := mid + 1
	k := 0
	for ; i <= mid && j <= end; k++ {
		if a[i] <= a[j] {
			tmp[k] = a[i]
			i++
		} else {
			tmp[k] = a[j]
			j++
		}
	}
	for ; i <= mid; k++ {
		tmp[k] = a[i]
		i++
	}
	for ; j <= end; k++ {
		tmp[k] = a[j]
		j++
	}
	k = 0
	for start <= end {
		a[start] = tmp[k]
		start++
		k++
	}
}
