# LinkedList

### Singly linked list
```go
type Node struct {
	Val  int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) append(val int) {
	
	if l.Head == nil {
		l.Head = &Node{Val: val}
		return
	}
	curNode := l.Head
	for curNode.Next != nil {
		curNode = curNode.Next
	}
	curNode.Next = &Node{Val: val}

}

func (l *LinkedList) prepend(val int) {
	newNode := &Node{Val: val}
	newNode.Next = l.Head
	l.Head = newNode
}

func (l *LinkedList) search(val int) bool {
	curNode := l.Head
	for curNode != nil {
		if curNode.Val == val {
			return true
		}
		curNode = curNode.Next
	}
	return false
}

func (l *LinkedList) delete(val int) {
	if l.Head == nil {
		return
	}
	if l.Head.Val == val {
		l.Head = l.Head.Next
		return
	}
	curNode := l.Head
	for curNode.Next != nil {
		if curNode.Next.Val == val {
			curNode.Next = curNode.Next.Next
			return
		}
		curNode = curNode.Next
	}
}
```

### Doubly linked list
```go

type Node struct {
	Val  int
	Next *Node
	Pre  *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) append(val int) {
	// check if the list is empty
	if l.Tail == nil {
		// if the list is empty, create the first node
		l.Tail = &Node{Val: val}
		// set the head to the first node
		l.Head = l.Tail
		// return from the function
		return
	}
	// get the last node from the list
	curNode := l.Tail
	// create a new node and add it to the end of the list
	curNode.Next = &Node{Val: val, Pre: curNode}
	// set the last node to the new node
	l.Tail = curNode.Next
}

func (l *LinkedList) prepend(val int) {
	// If the list is empty, create a new node and set it to the head and the tail
	if l.Head == nil {
		l.Head = &Node{Val: val}
		l.Tail = l.Head
		return
	}

	// If the list is not empty, create a new node and set it as the head
	curHead := l.Head
	newHead := &Node{Val: val, Next: curHead}
	curHead.Pre = newHead
	l.Head = newHead
}

func (l *LinkedList) search(val int) bool {
	curNode := l.Head
	for curNode != nil {
		if curNode.Val == val {
			return true
		}
		curNode = curNode.Next
	}
	return false
}

func (l *LinkedList) delete(val int) {
	if l.Head == nil {
		return
	}
	// check if the head is the value
	if l.Head.Val == val {
		l.Head = l.Head.Next // move the head to the next node
		if l.Head != nil {
			l.Head.Pre = nil
		}
		return
	}
	// 1. Get the head node of the linked list
	curNode := l.Head
	// 2. Loop through the linked list
	for curNode.Next != nil {
		// 3. If the next node's value is the value to be deleted
		if curNode.Next.Val == val {
			// 4. Skip the next node
			curNode.Next = curNode.Next.Next
			// 5. If the next node is not the tail node, set the previous node of the next node to the current node
			if curNode.Next.Next != nil {
				curNode.Next.Next.Pre = curNode
			}
			// 6. Exit the loop
			return
		}
		// 7. Move to the next node
		curNode = curNode.Next
	}
}

```