package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func randomIntArr(size, min, max int) []int {
	max += 1
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(max-min) + min
	}
	return arr
}

func Test_bubbleSort(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"empty array"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testArr := []int{}
			gotArr := []int{}
			bubbleSort(gotArr)
			sort.Ints(testArr)

			if !reflect.DeepEqual(testArr, gotArr) {
				t.Errorf("slices aren't equal")
			}
		})
	}
	for i := 0; i < 100; i++ {
		size := rand.Intn(150)
		min := 0
		max := rand.Intn(150)
		testName := fmt.Sprintf("arr size: %d, min: %d, max: %d", size, min, max)
		t.Run(testName, func(t *testing.T) {
			testArr := randomIntArr(size, min, max)
			gotArr := make([]int, size)
			copy(gotArr[:], testArr[:])
			sort.Ints(testArr)
			bubbleSort(gotArr)
			if !reflect.DeepEqual(testArr, gotArr) {
				t.Errorf("slices aren't equal")
			}
		})
	}
}

func Test_bubbleSortRecursive(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"empty array"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testArr := []int{}
			gotArr := []int{}
			gotArr = bubbleSortRecursive(gotArr)
			sort.Ints(testArr)

			if !reflect.DeepEqual(testArr, gotArr) {
				t.Errorf("slices aren't equal")
			}
		})
	}
	for i := 0; i < 100; i++ {
		size := rand.Intn(150)
		min := 0
		max := rand.Intn(150)
		testName := fmt.Sprintf("arr size: %d, min: %d, max: %d", size, min, max)
		t.Run(testName, func(t *testing.T) {
			testArr := randomIntArr(size, min, max)
			gotArr := make([]int, size)
			copy(gotArr[:], testArr[:])
			sort.Ints(testArr)
			bubbleSortRecursive(gotArr)
			if !reflect.DeepEqual(testArr, gotArr) {
				t.Errorf("slices aren't equal")
			}
		})
	}
}

func Test_bubbleSortReversed(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"empty array"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testArr := []int{}
			gotArr := []int{}
			bubbleSortReversed(gotArr)
			sort.Ints(testArr)

			if !reflect.DeepEqual(testArr, gotArr) {
				t.Errorf("slices aren't equal")
			}
		})
	}
	for i := 0; i < 100; i++ {
		size := rand.Intn(150)
		min := 0
		max := rand.Intn(150)
		testName := fmt.Sprintf("arr size: %d, min: %d, max: %d", size, min, max)
		t.Run(testName, func(t *testing.T) {
			testArr := randomIntArr(size, min, max)
			gotArr := make([]int, size)
			copy(gotArr[:], testArr[:])
			sort.Sort(sort.Reverse(sort.IntSlice(testArr)))
			bubbleSortReversed(gotArr)
			if !reflect.DeepEqual(testArr, gotArr) {
				t.Errorf("slices aren't equal")
			}
		})
	}
}

func Test_insertionSort(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"empty array"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testArr := []int{}
			gotArr := []int{}
			insertionSort(gotArr)
			sort.Ints(testArr)

			if !reflect.DeepEqual(testArr, gotArr) {
				t.Errorf("slices aren't equal")
			}
		})
	}
	for i := 0; i < 100; i++ {
		size := rand.Intn(150)
		min := 0
		max := rand.Intn(150)
		testName := fmt.Sprintf("arr size: %d, min: %d, max: %d", size, min, max)
		t.Run(testName, func(t *testing.T) {
			testArr := randomIntArr(size, min, max)
			gotArr := make([]int, size)
			copy(gotArr[:], testArr[:])
			sort.Ints(testArr)
			insertionSort(gotArr)
			if !reflect.DeepEqual(testArr, gotArr) {
				t.Errorf("slices aren't equal")
			}
		})
	}
}

func Test_selectionSort(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"empty array"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testArr := []int{}
			gotArr := []int{}
			selectionSort(gotArr)
			sort.Ints(testArr)

			if !reflect.DeepEqual(testArr, gotArr) {
				t.Errorf("slices aren't equal")
			}
		})
	}
	for i := 0; i < 100; i++ {
		size := rand.Intn(150)
		min := 0
		max := rand.Intn(150)
		testName := fmt.Sprintf("arr size: %d, min: %d, max: %d", size, min, max)
		t.Run(testName, func(t *testing.T) {
			testArr := randomIntArr(size, min, max)
			gotArr := make([]int, size)
			copy(gotArr[:], testArr[:])
			sort.Ints(testArr)
			selectionSort(gotArr)
			if !reflect.DeepEqual(testArr, gotArr) {
				t.Errorf("slices aren't equal")
			}
		})
	}
}

func Test_selectionSortByMax(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"empty array"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testArr := []int{}
			gotArr := []int{}
			selectionSortByMax(gotArr)
			sort.Ints(testArr)

			if !reflect.DeepEqual(testArr, gotArr) {
				t.Errorf("slices aren't equal")
			}
		})
	}
	for i := 0; i < 100; i++ {
		size := rand.Intn(150)
		min := 0
		max := rand.Intn(150)
		testName := fmt.Sprintf("arr size: %d, min: %d, max: %d", size, min, max)
		t.Run(testName, func(t *testing.T) {
			testArr := randomIntArr(size, min, max)
			gotArr := make([]int, size)
			copy(gotArr[:], testArr[:])
			sort.Ints(testArr)
			selectionSortByMax(gotArr)
			if !reflect.DeepEqual(testArr, gotArr) {
				t.Errorf("slices aren't equal")
			}
		})
	}
}

func Test_selectionSortBidirectional(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"empty array"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testArr := []int{}
			gotArr := []int{}
			selectionSortBidirectional(gotArr)
			sort.Ints(testArr)

			if !reflect.DeepEqual(testArr, gotArr) {
				t.Errorf("slices aren't equal")
			}
		})
	}
	for i := 0; i < 100; i++ {
		size := rand.Intn(150)
		min := 0
		max := rand.Intn(150)
		testName := fmt.Sprintf("arr size: %d, min: %d, max: %d", size, min, max)
		t.Run(testName, func(t *testing.T) {
			testArr := randomIntArr(size, min, max)
			gotArr := make([]int, size)
			copy(gotArr[:], testArr[:])
			sort.Ints(testArr)
			selectionSortBidirectional(gotArr)
			if !reflect.DeepEqual(testArr, gotArr) {
				t.Errorf("slices aren't equal")
			}
		})
	}
}

func Test_mergeSort(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"empty array"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testArr := []int{}
			gotArr := []int{}
			gotArr = mergeSort(gotArr)
			sort.Ints(testArr)

			if !reflect.DeepEqual(testArr, gotArr) {
				t.Errorf("slices aren't equal")
			}
		})
	}
	for i := 0; i < 100; i++ {
		size := rand.Intn(150)
		min := 0
		max := rand.Intn(150)
		testName := fmt.Sprintf("arr size: %d, min: %d, max: %d", size, min, max)
		t.Run(testName, func(t *testing.T) {
			testArr := randomIntArr(size, min, max)
			gotArr := make([]int, size)
			copy(gotArr[:], testArr[:])
			sort.Ints(testArr)
			gotArr = mergeSort(gotArr)
			if !reflect.DeepEqual(testArr, gotArr) {
				t.Errorf("slices aren't equal")
			}
		})
	}
}

func Test_quickSort(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"empty array"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testArr := []int{}
			gotArr := []int{}
			gotArr = quickSort(gotArr)
			sort.Ints(testArr)

			if !reflect.DeepEqual(testArr, gotArr) {
				t.Errorf("slices aren't equal")
			}
		})
	}
	for i := 0; i < 100; i++ {
		size := rand.Intn(150)
		min := 0
		max := rand.Intn(150)
		testName := fmt.Sprintf("arr size: %d, min: %d, max: %d", size, min, max)
		t.Run(testName, func(t *testing.T) {
			testArr := randomIntArr(size, min, max)
			gotArr := make([]int, size)
			copy(gotArr[:], testArr[:])
			sort.Ints(testArr)
			gotArr = quickSort(gotArr)
			if !reflect.DeepEqual(testArr, gotArr) {
				t.Errorf("slices aren't equal")
			}
		})
	}
}
