package main

import (
	"container/heap"
	"fmt"
)

// В основе нашей кучи будет обычный срез
// В куче будем хранить целые числа
type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

// Определяем логику определения приоритета элементов кучи
// в нашем случае приоритет у того элемента, который просто
// меньше по значению
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

// Здесь определяем логику перестановки элементов кучи - простое
// копирование в нашем случае
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Логика добавления нового элемента в кучу.
// Так как у нас срез в основе кучи -  то простое добавление в
// конец среза
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Логика извлечения - обычная для списка, основанного на срезе
// с изменением размера среза
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	nums := IntHeap{3, 1, 4, 5, 1, 1, 2, 6}
	heap.Init(&nums)
	fmt.Println(nums)
}
