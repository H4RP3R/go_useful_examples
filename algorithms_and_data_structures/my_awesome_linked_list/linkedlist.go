package myawesomelinkedlist

type Node[T any] struct {
	Value T
	next  *Node[T]
}

type LinkedList[T any] struct {
	Head *Node[T]
	size int
}

// Add appends the specified item to the end of this list.
func (l *LinkedList[T]) Add(item T) {
	if l.size == 0 {
		l.Head = &Node[T]{Value: item, next: nil}
	} else {
		var lastNode *Node[T]
		node := l.Head
		for node != nil {
			lastNode = node
			node = node.next
		}
		lastNode.next = &Node[T]{Value: item, next: nil}
	}
	l.size++
}

// AddFirst inserts the specified element at the beginning of this list.
func (l *LinkedList[T]) AddFirst(item T) {
	if l.size == 0 {
		l.Head = &Node[T]{Value: item, next: nil}
	} else {
		newNode := &Node[T]{Value: item, next: l.Head}
		l.Head = newNode
	}
	l.size++
}

// Poll retrieves and removes the head (first element) of this list.
func (l *LinkedList[T]) Poll() (T, bool) {
	var value T
	if l.size == 0 {
		return value, false
	}

	value = l.Head.Value
	l.Head = l.Head.next
	l.size--

	return value, true
}
