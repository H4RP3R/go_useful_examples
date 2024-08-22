package main

func quickSort(ar []int) []int {
	if len(ar) < 2 {
		return ar
	}

	less := []int{}
	greater := []int{}
	pivot := ar[0]
	for _, n := range ar[1:] {
		if n < pivot {
			less = append(less, n)
		} else {
			greater = append(greater, n)
		}
	}

	return append(append(quickSort(less), pivot), quickSort(greater)...)
}
