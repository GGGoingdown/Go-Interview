# Queue

```go
type Queue struct {
	elements []int
	size     int
}

func NewQueue() *Queue {
	return &Queue{
		elements: make([]int, 0),
		size:     0,
	}
}

// This function adds an element to the end of the queue.
func (q *Queue) Enqueue(element int) {
	q.elements = append(q.elements, element)
	q.size++
}

// Dequeue removes the element at the front of the queue and returns it.
// It returns an error if the queue is empty.
func (q *Queue) Dequeue() (int, error) {
	if q.size == 0 {
		return 0, errors.New("Queue is empty")
	}
	element := q.elements[0]
	q.elements = q.elements[1:]
	q.size--
	return element, nil
}

```