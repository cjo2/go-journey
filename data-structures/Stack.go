package data_structures

// Stack There's nothing special about this at all. Internally, we're just using a slice.
type Stack struct {
	Items []int
}

func (s *Stack) Push(i int) {
	s.Items = append(s.Items, i)
}

func (s *Stack) Pop() int {
	popped := s.Items[len(s.Items)-1]
	s.Items = s.Items[:len(s.Items)-1]
	return popped
}
