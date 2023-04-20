# Stack

```go
type Stack struct {
	data   []int
	length int
}

func NewStack() *Stack {
	return &Stack{}
}

// Add adds data to the stack.
func (s *Stack) Add(data int) {
	s.data = append(s.data, data)
	s.length++
}

// Pop removes the last item from the stack and returns it.
// If the stack is empty, an error is returned.
func (s *Stack) Pop() (int, error) {
	if s.length == 0 {
		return 0, errors.New("Stack is empty")
	}

	data := s.data[s.length-1]
	s.data = s.data[:s.length-1]
	s.length--
	return data, nil
}

```