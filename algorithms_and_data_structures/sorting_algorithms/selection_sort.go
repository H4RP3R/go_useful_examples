package main

func selectionSort(ar []int) {
	if len(ar) < 2 {
		return
	}

	for i := 0; i < len(ar); i++ {
		minIdx := i
		for j := i + 1; j < len(ar); j++ {
			if ar[j] < ar[minIdx] {
				minIdx = j
			}
		}
		ar[i], ar[minIdx] = ar[minIdx], ar[i]
	}
}

func selectionSortByMax(ar []int) {
	if len(ar) < 2 {
		return
	}

	for i := len(ar) - 1; i > 0; i-- {
		maxIdx := i
		for j := i - 1; j >= 0; j-- {
			if ar[j] > ar[maxIdx] {
				maxIdx = j
			}
		}
		ar[i], ar[maxIdx] = ar[maxIdx], ar[i]
	}
}

// Когда вы ищете одновременно максимальное и минимальное число, в конце итерации вам потребуется совершить две
// операции обмена: переместить минимальное и максимальное число. Учтите, что после первой операции может быть
// перемещено число с индексом, который будет использоваться во второй операции.
func selectionSortBidirectional(ar []int) {
	n := len(ar)

	for i := 0; i < n/2; i++ {
		minIdx := i
		maxIdx := i

		for j := i + 1; j < n-i; j++ {
			if ar[j] < ar[minIdx] {
				minIdx = j
			}
			if ar[j] > ar[maxIdx] {
				maxIdx = j
			}
		}

		if minIdx != i {
			ar[i], ar[minIdx] = ar[minIdx], ar[i]
		}

		if maxIdx == i {
			maxIdx = minIdx
		}

		if maxIdx != n-i-1 {
			ar[n-i-1], ar[maxIdx] = ar[maxIdx], ar[n-i-1]
		}
	}
}
