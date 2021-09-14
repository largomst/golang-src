package linkedlist

type Node struct {
	val  int
	next *Node
}

type MyLinkedList struct {
	head *Node
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (m *MyLinkedList) Get(index int) int {
	i := 0
	for cur := m.head; cur != nil; cur = cur.next {
		if i == index {
			return cur.val
		}
		i++
	}
	return -1

}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (m *MyLinkedList) AddAtHead(val int) {
	newNode := &Node{val, m.head}
	m.head = newNode
}

/** Append a node of value val to the last element of the linked list. */
func (m *MyLinkedList) AddAtTail(val int) {
	cur := m.head
	if cur != nil {
		for ; cur.next != nil; cur = cur.next {
		}
		newNode := &Node{val, nil}
		cur.next = newNode
	} else {
		m.head = &Node{val, nil}
	}
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (m *MyLinkedList) AddAtIndex(index int, val int) {
	var prev *Node
	cur := m.head
	found := false
	i := 0
	if cur == nil && index == 0 {
		m.head = &Node{val, nil}
		return
	}

	for cur != nil {
		if i == index {
			found = true
			break
		}
		i++
		prev = cur
		cur = cur.next
	}
	if found {
		if prev != nil {
			prev.next = &Node{val, prev.next}
		} else {
			m.head = &Node{val, m.head}
		}
	} else {
		if i == index {
			prev.next = &Node{val, prev.next}
		}
	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (m *MyLinkedList) DeleteAtIndex(index int) {
	var prev *Node
	found := false
	i := 0
	cur := m.head
	for cur != nil {
		if i == index {
			found = true
			break
		}
		i++
		prev = cur
		cur = cur.next

	}
	if found {
		if prev != nil {
			prev.next = cur.next
		} else {
			m.head = cur.next
		}

	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
