// Пример реализации статической очереди, содержащей целые числа, и метода добавления нового
// элемента-льготника в начало очереди, основанной на динамическом массиве.

package main

import "fmt"

// ErrQueueOverFlow - ошибка при попытке добавления элемента в
// полностью заполненную очередь
var ErrQueueOverFlow = fmt.Errorf("ошибка переполнения очереди")

// IntDoubleQueue - двухсторонняя очередь элементов, хранящих целые
// числа
type IntDoubleQueue struct {
	data      []int // структура, непосредственно хранящая элементы
	tailIndex int   // индекс элемента, являющегося в настоящий
	// момент хвостом очереди
	size int // размер очереди
}

// QueueHead - метод добавления элемента в начало очереди
func (q *IntDoubleQueue) QueueHead(el int) error {
	if q.tailIndex >= q.size-1 {
		return ErrQueueOverFlow
	}
	for i := q.tailIndex; i >= 0; i-- {
		q.data[i+1] = q.data[i]
	}
	q.data[0] = el
	q.tailIndex++
	return nil
}
