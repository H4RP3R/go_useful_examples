package main

func mergeSort(ar []int) []int {
	if len(ar) < 2 {
		return ar
	}

	mid := len(ar) / 2
	left := mergeSort(ar[:mid])
	right := mergeSort(ar[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	merged := []int{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			merged = append(merged, left[i])
			i++
		} else {
			merged = append(merged, right[j])
			j++
		}
	}

	merged = append(merged, left[i:]...)
	merged = append(merged, right[j:]...)

	return merged
}
