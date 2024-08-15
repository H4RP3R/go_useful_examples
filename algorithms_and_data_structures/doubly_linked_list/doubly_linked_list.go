package doublylinkedlist

// IntNode - узел списка
type IntNode struct {
	Value int
	Next  *IntNode
	Prev  *IntNode
}

// NewIntNode - конструктор типа IntNode
func NewIntNode(value int) *IntNode {
	return &IntNode{value, nil, nil}
}

// DoubleLinkedList - список
type DoubleLinkedList struct {
	size int
	Head *IntNode
	Tail *IntNode
}

// NewDoubleLinkedList - конструктор списка
func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{0, nil, nil}
}
