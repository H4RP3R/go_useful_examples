package main

func bubbleSort(ar []int) {
	sorted := false

	for !sorted {
		sorted = true
		for i := 0; i < len(ar)-1; i++ {
			if ar[i] > ar[i+1] {
				ar[i], ar[i+1] = ar[i+1], ar[i]
				sorted = false
			}
		}
	}
}

func bubbleSortRecursive(ar []int) []int {
	if len(ar) == 1 {
		return ar
	}

	sorted := true
	for i := 0; i < len(ar)-1; i++ {
		if ar[i] > ar[i+1] {
			ar[i], ar[i+1] = ar[i+1], ar[i]
			sorted = false
		}
	}

	if sorted {
		return ar
	}

	return bubbleSortRecursive(ar)
}

func bubbleSortReversed(ar []int) {
	sorted := false

	for !sorted {
		sorted = true
		for i := 0; i < len(ar)-1; i++ {
			if ar[i] < ar[i+1] {
				ar[i], ar[i+1] = ar[i+1], ar[i]
				sorted = false
			}
		}
	}
}
