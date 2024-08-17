package main

import (
	"container/ring"
	"fmt"
)

func main() {
	r := ring.New(5) // Создаем кольцо размером 5 элементов
	n := r.Len()
	for i := 0; i < n; i++ { // Инициализируем элементы кольца значениями
		r.Value = i
		r = r.Next()
	}

	// Перемещаемся вперёд по кольцу и печатаем значение каждого элемента
	for j := 0; j < n; j++ {
		fmt.Println(r.Value)
		r = r.Next()
	}
	// Перемещаемся назад по кольцу и печатаем значение каждого элемента
	for k := n; k > 0; k-- {
		r = r.Prev()
		fmt.Println(r.Value)
	}

	for j := 0; j < 2; j++ { // перемещаемся вперед по кольцу на два элемента
		r = r.Next()
	}
	newRing := r.Next()
	r = r.Next()
	fmt.Printf("%d %d", newRing.Value, r.Value) // 3 3
}
