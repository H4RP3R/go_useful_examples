package myawesomelinkedlist

import (
	"testing"
)

func CompareLinkedLists[T comparable](l1, l2 LinkedList[T]) bool {
	if l1.size != l2.size {
		return false
	}

	if l1.size == 0 && l2.size == 0 {
		return true
	}

	item1 := l1.Head
	item2 := l2.Head
	for item1 != nil {
		if item1.Value != item2.Value {
			return false
		}
		item1 = item1.next
		item2 = item2.next
	}

	return true
}

func TestLinkedListAddFirstInt(t *testing.T) {
	var tests = []struct {
		name string
		item int
		want LinkedList[int]
	}{
		{
			"add first int 1",
			1,
			LinkedList[int]{Head: &Node[int]{Value: 1, next: nil}, size: 1},
		},
		{
			"add first int 2",
			2,
			LinkedList[int]{Head: &Node[int]{Value: 2, next: &Node[int]{Value: 1,
				next: nil}}, size: 2},
		},
		{
			"add first int 3",
			3,
			LinkedList[int]{Head: &Node[int]{Value: 3,
				next: &Node[int]{Value: 2,
					next: &Node[int]{Value: 1,
						next: nil}}}, size: 3},
		},
	}

	var list LinkedList[int]
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list.AddFirst(tt.item)
			listsAreEqual := CompareLinkedLists(tt.want, list)
			if !listsAreEqual {
				t.Error("lists aren't equal")
			}
		})
	}
}

func TestLinkedListAddString(t *testing.T) {
	var tests = []struct {
		name string
		item string
		want LinkedList[string]
	}{
		{
			"add string 1",
			"a",
			LinkedList[string]{Head: &Node[string]{Value: "a", next: nil}, size: 1},
		},
		{
			"add string 2",
			"bb",
			LinkedList[string]{Head: &Node[string]{Value: "a", next: &Node[string]{Value: "bb",
				next: nil}}, size: 2},
		},
		{
			"add first int 3",
			"ccc",
			LinkedList[string]{Head: &Node[string]{Value: "a",
				next: &Node[string]{Value: "bb",
					next: &Node[string]{Value: "ccc",
						next: nil}}}, size: 3},
		},
	}

	var list LinkedList[string]
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list.Add(tt.item)
			listsAreEqual := CompareLinkedLists(tt.want, list)
			if !listsAreEqual {
				t.Error("lists aren't equal")
			}
		})
	}
}

func TestLinkedListPollInt(t *testing.T) {
	type want struct {
		value int
		ok    bool
		size  int
	}

	tests := []struct {
		name string
		want want
	}{
		{"poll 1", want{
			value: 11,
			ok:    true,
			size:  2,
		}},
	}

	var list LinkedList[int]
	list.Add(11)
	list.Add(22)
	list.Add(33)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, ok := list.Poll()
			if res != tt.want.value {
				t.Errorf("want value: %v, got value: %v", tt.want.value, res)
			}
			if ok != tt.want.ok {
				t.Errorf("want ok: %v, got ok: %v", tt.want.ok, ok)
			}
			if list.size != tt.want.size {
				t.Errorf("want size: %v, got size: %v", tt.want.size, list.size)
			}
		})
	}
}
