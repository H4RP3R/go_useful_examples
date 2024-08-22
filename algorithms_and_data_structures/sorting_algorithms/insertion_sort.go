package main

func insertionSort(ar []int) {
	for i := 1; i < len(ar); i++ {
		for j := i; j > 0; j-- {
			if ar[j] < ar[j-1] {
				ar[j], ar[j-1] = ar[j-1], ar[j]
			} else {
				break
			}
		}
	}
}
